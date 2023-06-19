package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/surahman/FTeX/pkg/common"
	"github.com/surahman/FTeX/pkg/constants"
	graphql_generated "github.com/surahman/FTeX/pkg/graphql/generated"
	"github.com/surahman/FTeX/pkg/models"
	modelsPostgres "github.com/surahman/FTeX/pkg/models/postgres"
	"github.com/surahman/FTeX/pkg/validator"
	"go.uber.org/zap"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input *modelsPostgres.UserAccount) (*models.JWTAuthResponse, error) {
	var (
		authToken *models.JWTAuthResponse
		err       error
		httpMsg   string
		payload   any
	)

	if authToken, httpMsg, _, payload, err = common.HTTPRegisterUser(r.auth, r.db, r.logger, input); err != nil {
		return nil, fmt.Errorf("%s: %s", httpMsg, payload)
	}

	return authToken, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, input models.HTTPDeleteUserRequest) (string, error) {
	var (
		clientID    uuid.UUID
		err         error
		userAccount modelsPostgres.User
	)

	if err = validator.ValidateStruct(&input); err != nil {
		return "", fmt.Errorf("validation %w", err)
	}

	// Validate the JWT and extract the clientID. Compare the clientID against the deletion request login
	// credentials.
	if clientID, _, err = AuthorizationCheck(ctx, r.auth, r.logger, r.authHeaderKey); err != nil {
		return "", errors.New("authorization failure")
	}

	// Get user account information to validate against.
	if userAccount, err = r.db.UserGetInfo(clientID); err != nil {
		r.logger.Warn("failed to read user record during an account deletion request",
			zap.String("clientID", clientID.String()), zap.Error(err))

		return "", errors.New("please retry your request later")
	}

	// Check Username and password.
	if userAccount.Username != input.Username {
		return "", errors.New("invalid deletion request")
	}

	if err = r.auth.CheckPassword(userAccount.Password, input.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Check the confirmation message.
	if fmt.Sprintf(constants.DeleteUserAccountConfirmation(), userAccount.Username) != input.Confirmation {
		return "", errors.New("incorrect or incomplete deletion request confirmation")
	}

	// Check to make sure the account is not already deleted.
	if userAccount.IsDeleted {
		r.logger.Warn("attempt to delete an already deleted user account",
			zap.String("username", userAccount.Username))

		return "", errors.New("user account is already deleted")
	}

	// Mark account as deleted.
	if err = r.db.UserDelete(clientID); err != nil {
		r.logger.Warn("failed to mark a user record as deleted",
			zap.String("username", userAccount.Username), zap.Error(err))

		return "", errors.New("please retry your request later")
	}

	return "account successfully deleted", nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input modelsPostgres.UserLoginCredentials) (*models.JWTAuthResponse, error) {
	var (
		err            error
		authToken      *models.JWTAuthResponse
		clientID       uuid.UUID
		hashedPassword string
	)

	if err = validator.ValidateStruct(&input); err != nil {
		return nil, fmt.Errorf("validation %w", err)
	}

	if clientID, hashedPassword, err = r.db.UserCredentials(input.Username); err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err = r.auth.CheckPassword(hashedPassword, input.Password); err != nil {
		return nil, errors.New("invalid username or password")
	}

	if authToken, err = r.auth.GenerateJWT(clientID); err != nil {
		r.logger.Error("failure generating JWT during login", zap.Error(err))

		return nil, errors.New("please retry your request later")
	}

	return authToken, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context) (*models.JWTAuthResponse, error) {
	var (
		err         error
		freshToken  *models.JWTAuthResponse
		clientID    uuid.UUID
		accountInfo modelsPostgres.User
		expiresAt   int64
	)

	// Validate the JWT and extract the clientID. Compare the clientID against the deletion request login
	// credentials.
	if clientID, expiresAt, err = AuthorizationCheck(ctx, r.auth, r.logger, r.authHeaderKey); err != nil {
		return freshToken, errors.New("authorization failure")
	}

	if accountInfo, err = r.db.UserGetInfo(clientID); err != nil {
		r.logger.Warn("failed to read user record for a valid JWT",
			zap.String("username", accountInfo.Username), zap.Error(err))

		return nil, errors.New("please retry your request later")
	}

	if accountInfo.IsDeleted {
		r.logger.Warn("attempt to refresh a JWT for a deleted user", zap.String("clientID", accountInfo.Username))

		return nil, errors.New("invalid token")
	}

	// Do not refresh tokens that are outside the refresh threshold. Tokens could expire during the execution of
	// this handler but expired ones would be rejected during token validation. Thus, it is not necessary to
	// re-check expiration.
	if expiresAt-time.Now().Unix() > r.auth.RefreshThreshold() {
		return nil, fmt.Errorf("JWT is still valid for more than %d seconds", r.auth.RefreshThreshold())
	}

	if freshToken, err = r.auth.GenerateJWT(clientID); err != nil {
		r.logger.Error("failure generating JWT during token refresh", zap.Error(err))

		return nil, errors.New("please retry your request later")
	}

	return freshToken, nil
}

// Mutation returns graphql_generated.MutationResolver implementation.
func (r *Resolver) Mutation() graphql_generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
