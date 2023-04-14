package postgres

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/afero"
	"github.com/surahman/FTeX/pkg/constants"
	"github.com/surahman/FTeX/pkg/logger"
	"go.uber.org/zap"
)

// Mock Postgres SQLC Querier interface stub generation. This is local to the Postgres package.
//go:generate mockgen -destination=querier_mocks.go -package=postgres github.com/surahman/FTeX/pkg/postgres Querier

// Mock Postgres interface stub generation.
//go:generate mockgen -destination=../mocks/mock_postgres.go -package=mocks github.com/surahman/FTeX/pkg/postgres Postgres

// Postgres is the interface through which the queries will be executed on the database.
type Postgres interface {
}

// Check to ensure the Postgres interface has been implemented.
var _ Postgres = &postgresImpl{}

// postgresImpl contains objects required to interface with the database.
type postgresImpl struct {
	conf    config
	pool    *pgxpool.Pool
	logger  *logger.Logger
	queries *Queries
	Query   Querier
}

// NewPostgres will create a new Postgres configuration by loading it.
func NewPostgres(fs *afero.Fs, logger *logger.Logger) (Postgres, error) {
	if fs == nil || logger == nil {
		return nil, errors.New("nil file system of logger supplied")
	}

	return newPostgresImpl(fs, logger)
}

// newPostgresImpl will create a new postgresImpl configuration and load it from disk.
func newPostgresImpl(fs *afero.Fs, logger *logger.Logger) (*postgresImpl, error) {
	if fs == nil || logger == nil {
		return nil, errors.New("nil file system or logger supplied")
	}

	postgres := &postgresImpl{conf: newConfig(), logger: logger}
	if err := postgres.conf.Load(*fs); err != nil {
		postgres.logger.Error("failed to load Postgres configurations from disk", zap.Error(err))

		return nil, err
	}

	return postgres, nil
}

// Open will start a database connection pool and establish a connection.
func (p *postgresImpl) Open() error {
	var err error
	if err = p.verifySession(); err == nil {
		return errors.New("connection is already established to Postgres")
	}

	var pgxConfig *pgxpool.Config

	if pgxConfig, err = pgxpool.ParseConfig(fmt.Sprintf(constants.GetPostgresDSN(),
		p.conf.Authentication.Username,
		p.conf.Authentication.Password,
		p.conf.Connection.Host,
		p.conf.Connection.Port,
		p.conf.Connection.Database,
		p.conf.Connection.Timeout)); err != nil {
		msg := "failed to parse Postgres DSN"
		p.logger.Error(msg, zap.Error(err))

		return fmt.Errorf(msg+"%w", err)
	}

	pgxConfig.MaxConns = p.conf.Pool.MaxConns
	pgxConfig.MinConns = p.conf.Pool.MinConns
	pgxConfig.HealthCheckPeriod = p.conf.Pool.HealthCheckPeriod

	if p.pool, err = pgxpool.NewWithConfig(context.Background(), pgxConfig); err != nil {
		msg := "failed to configure Postgres connection"
		p.logger.Error(msg, zap.Error(err))

		return fmt.Errorf(msg+"%w", err)
	}

	// Binary Exponential Backoff connection to Postgres. The lazy connection can be opened via a ping to the database.
	if err = p.createSessionRetry(); err != nil {
		return err
	}

	// Setup SQLC DBTX interface.
	p.queries = New(p.pool)
	p.Query = p.queries

	return nil
}

// verifySession will check to see if a session is established.
func (p *postgresImpl) verifySession() error {
	if p.pool == nil || p.pool.Ping(context.Background()) != nil {
		return errors.New("no session established")
	}

	return nil
}

// createSessionRetry will attempt to open the connection using binary exponential back-off.
// Stop on the first success or fail after the last one.
func (p *postgresImpl) createSessionRetry() (err error) {
	for attempt := 1; attempt <= p.conf.Connection.MaxConnAttempts; attempt++ {
		waitTime := time.Duration(math.Pow(2, float64(attempt))) * time.Second
		p.logger.Info(fmt.Sprintf("Attempting connection to Postgres database in %s...", waitTime),
			zap.String("attempt", strconv.Itoa(attempt)))
		time.Sleep(waitTime)

		if err = p.pool.Ping(context.Background()); err == nil {
			return nil
		}
	}
	p.logger.Error("unable to establish connection to Postgres database", zap.Error(err))

	return
}

// Close will close the database connection pool.
func (p *postgresImpl) Close() (err error) {
	if err = p.verifySession(); err != nil {
		msg := "no established Postgres connection to close"
		p.logger.Error(msg)

		return errors.New(msg)
	}

	p.pool.Close()

	return
}

// Healthcheck will run a ping on the database to ascertain health.
func (p *postgresImpl) Healthcheck() error {
	var err error
	if err = p.verifySession(); err != nil {
		return err
	}

	if err = p.pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("postgres cluster ping failed: %w", err)
	}

	return nil
}
