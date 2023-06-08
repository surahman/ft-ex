// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/surahman/FTeX/pkg/models"
	models1 "github.com/surahman/FTeX/pkg/models/postgres"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	FiatAccount() FiatAccountResolver
	FiatDepositResponse() FiatDepositResponseResolver
	FiatExchangeOfferResponse() FiatExchangeOfferResponseResolver
	FiatExchangeTransferResponse() FiatExchangeTransferResponseResolver
	FiatJournal() FiatJournalResolver
	FiatTransactionsPaginated() FiatTransactionsPaginatedResolver
	Mutation() MutationResolver
	PriceQuote() PriceQuoteResolver
	Query() QueryResolver
	CryptoOfferRequest() CryptoOfferRequestResolver
	FiatDepositRequest() FiatDepositRequestResolver
	FiatExchangeOfferRequest() FiatExchangeOfferRequestResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	CryptoOpenAccountResponse struct {
		ClientID func(childComplexity int) int
		Ticker   func(childComplexity int) int
	}

	FiatAccount struct {
		Balance   func(childComplexity int) int
		ClientID  func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		Currency  func(childComplexity int) int
		LastTx    func(childComplexity int) int
		LastTxTs  func(childComplexity int) int
	}

	FiatBalancesPaginated struct {
		AccountBalances func(childComplexity int) int
		Links           func(childComplexity int) int
	}

	FiatDepositResponse struct {
		Balance     func(childComplexity int) int
		ClientID    func(childComplexity int) int
		Currency    func(childComplexity int) int
		LastTx      func(childComplexity int) int
		TxID        func(childComplexity int) int
		TxTimestamp func(childComplexity int) int
	}

	FiatExchangeOfferResponse struct {
		DebitAmount func(childComplexity int) int
		Expires     func(childComplexity int) int
		OfferID     func(childComplexity int) int
		PriceQuote  func(childComplexity int) int
	}

	FiatExchangeTransferResponse struct {
		DestinationReceipt func(childComplexity int) int
		SourceReceipt      func(childComplexity int) int
	}

	FiatJournal struct {
		Amount       func(childComplexity int) int
		ClientID     func(childComplexity int) int
		Currency     func(childComplexity int) int
		TransactedAt func(childComplexity int) int
		TxID         func(childComplexity int) int
	}

	FiatOpenAccountResponse struct {
		ClientID func(childComplexity int) int
		Currency func(childComplexity int) int
	}

	FiatTransactionsPaginated struct {
		Links        func(childComplexity int) int
		Transactions func(childComplexity int) int
	}

	JWTAuthResponse struct {
		Expires   func(childComplexity int) int
		Threshold func(childComplexity int) int
		Token     func(childComplexity int) int
	}

	Links struct {
		NextPage   func(childComplexity int) int
		PageCursor func(childComplexity int) int
	}

	Mutation struct {
		DeleteUser           func(childComplexity int, input models.HTTPDeleteUserRequest) int
		DepositFiat          func(childComplexity int, input models.HTTPDepositCurrencyRequest) int
		ExchangeOfferFiat    func(childComplexity int, input models.HTTPExchangeOfferRequest) int
		ExchangeTransferFiat func(childComplexity int, offerID string) int
		LoginUser            func(childComplexity int, input models1.UserLoginCredentials) int
		OfferCrypto          func(childComplexity int, input models.HTTPCryptoOfferRequest) int
		OpenCrypto           func(childComplexity int, ticker string) int
		OpenFiat             func(childComplexity int, currency string) int
		RefreshToken         func(childComplexity int) int
		RegisterUser         func(childComplexity int, input *models1.UserAccount) int
	}

	PriceQuote struct {
		Amount         func(childComplexity int) int
		ClientID       func(childComplexity int) int
		DestinationAcc func(childComplexity int) int
		Rate           func(childComplexity int) int
		SourceAcc      func(childComplexity int) int
	}

	Query struct {
		BalanceAllFiat            func(childComplexity int, pageCursor *string, pageSize *int32) int
		BalanceFiat               func(childComplexity int, currencyCode string) int
		Healthcheck               func(childComplexity int) int
		TransactionDetailsAllFiat func(childComplexity int, input models.FiatPaginatedTxDetailsRequest) int
		TransactionDetailsFiat    func(childComplexity int, transactionID string) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "CryptoOpenAccountResponse.clientID":
		if e.complexity.CryptoOpenAccountResponse.ClientID == nil {
			break
		}

		return e.complexity.CryptoOpenAccountResponse.ClientID(childComplexity), true

	case "CryptoOpenAccountResponse.ticker":
		if e.complexity.CryptoOpenAccountResponse.Ticker == nil {
			break
		}

		return e.complexity.CryptoOpenAccountResponse.Ticker(childComplexity), true

	case "FiatAccount.balance":
		if e.complexity.FiatAccount.Balance == nil {
			break
		}

		return e.complexity.FiatAccount.Balance(childComplexity), true

	case "FiatAccount.clientID":
		if e.complexity.FiatAccount.ClientID == nil {
			break
		}

		return e.complexity.FiatAccount.ClientID(childComplexity), true

	case "FiatAccount.createdAt":
		if e.complexity.FiatAccount.CreatedAt == nil {
			break
		}

		return e.complexity.FiatAccount.CreatedAt(childComplexity), true

	case "FiatAccount.currency":
		if e.complexity.FiatAccount.Currency == nil {
			break
		}

		return e.complexity.FiatAccount.Currency(childComplexity), true

	case "FiatAccount.lastTx":
		if e.complexity.FiatAccount.LastTx == nil {
			break
		}

		return e.complexity.FiatAccount.LastTx(childComplexity), true

	case "FiatAccount.lastTxTs":
		if e.complexity.FiatAccount.LastTxTs == nil {
			break
		}

		return e.complexity.FiatAccount.LastTxTs(childComplexity), true

	case "FiatBalancesPaginated.accountBalances":
		if e.complexity.FiatBalancesPaginated.AccountBalances == nil {
			break
		}

		return e.complexity.FiatBalancesPaginated.AccountBalances(childComplexity), true

	case "FiatBalancesPaginated.links":
		if e.complexity.FiatBalancesPaginated.Links == nil {
			break
		}

		return e.complexity.FiatBalancesPaginated.Links(childComplexity), true

	case "FiatDepositResponse.balance":
		if e.complexity.FiatDepositResponse.Balance == nil {
			break
		}

		return e.complexity.FiatDepositResponse.Balance(childComplexity), true

	case "FiatDepositResponse.clientId":
		if e.complexity.FiatDepositResponse.ClientID == nil {
			break
		}

		return e.complexity.FiatDepositResponse.ClientID(childComplexity), true

	case "FiatDepositResponse.currency":
		if e.complexity.FiatDepositResponse.Currency == nil {
			break
		}

		return e.complexity.FiatDepositResponse.Currency(childComplexity), true

	case "FiatDepositResponse.lastTx":
		if e.complexity.FiatDepositResponse.LastTx == nil {
			break
		}

		return e.complexity.FiatDepositResponse.LastTx(childComplexity), true

	case "FiatDepositResponse.txId":
		if e.complexity.FiatDepositResponse.TxID == nil {
			break
		}

		return e.complexity.FiatDepositResponse.TxID(childComplexity), true

	case "FiatDepositResponse.txTimestamp":
		if e.complexity.FiatDepositResponse.TxTimestamp == nil {
			break
		}

		return e.complexity.FiatDepositResponse.TxTimestamp(childComplexity), true

	case "FiatExchangeOfferResponse.debitAmount":
		if e.complexity.FiatExchangeOfferResponse.DebitAmount == nil {
			break
		}

		return e.complexity.FiatExchangeOfferResponse.DebitAmount(childComplexity), true

	case "FiatExchangeOfferResponse.expires":
		if e.complexity.FiatExchangeOfferResponse.Expires == nil {
			break
		}

		return e.complexity.FiatExchangeOfferResponse.Expires(childComplexity), true

	case "FiatExchangeOfferResponse.offerID":
		if e.complexity.FiatExchangeOfferResponse.OfferID == nil {
			break
		}

		return e.complexity.FiatExchangeOfferResponse.OfferID(childComplexity), true

	case "FiatExchangeOfferResponse.priceQuote":
		if e.complexity.FiatExchangeOfferResponse.PriceQuote == nil {
			break
		}

		return e.complexity.FiatExchangeOfferResponse.PriceQuote(childComplexity), true

	case "FiatExchangeTransferResponse.destinationReceipt":
		if e.complexity.FiatExchangeTransferResponse.DestinationReceipt == nil {
			break
		}

		return e.complexity.FiatExchangeTransferResponse.DestinationReceipt(childComplexity), true

	case "FiatExchangeTransferResponse.sourceReceipt":
		if e.complexity.FiatExchangeTransferResponse.SourceReceipt == nil {
			break
		}

		return e.complexity.FiatExchangeTransferResponse.SourceReceipt(childComplexity), true

	case "FiatJournal.amount":
		if e.complexity.FiatJournal.Amount == nil {
			break
		}

		return e.complexity.FiatJournal.Amount(childComplexity), true

	case "FiatJournal.clientID":
		if e.complexity.FiatJournal.ClientID == nil {
			break
		}

		return e.complexity.FiatJournal.ClientID(childComplexity), true

	case "FiatJournal.currency":
		if e.complexity.FiatJournal.Currency == nil {
			break
		}

		return e.complexity.FiatJournal.Currency(childComplexity), true

	case "FiatJournal.transactedAt":
		if e.complexity.FiatJournal.TransactedAt == nil {
			break
		}

		return e.complexity.FiatJournal.TransactedAt(childComplexity), true

	case "FiatJournal.txID":
		if e.complexity.FiatJournal.TxID == nil {
			break
		}

		return e.complexity.FiatJournal.TxID(childComplexity), true

	case "FiatOpenAccountResponse.clientID":
		if e.complexity.FiatOpenAccountResponse.ClientID == nil {
			break
		}

		return e.complexity.FiatOpenAccountResponse.ClientID(childComplexity), true

	case "FiatOpenAccountResponse.currency":
		if e.complexity.FiatOpenAccountResponse.Currency == nil {
			break
		}

		return e.complexity.FiatOpenAccountResponse.Currency(childComplexity), true

	case "FiatTransactionsPaginated.links":
		if e.complexity.FiatTransactionsPaginated.Links == nil {
			break
		}

		return e.complexity.FiatTransactionsPaginated.Links(childComplexity), true

	case "FiatTransactionsPaginated.transactions":
		if e.complexity.FiatTransactionsPaginated.Transactions == nil {
			break
		}

		return e.complexity.FiatTransactionsPaginated.Transactions(childComplexity), true

	case "JWTAuthResponse.expires":
		if e.complexity.JWTAuthResponse.Expires == nil {
			break
		}

		return e.complexity.JWTAuthResponse.Expires(childComplexity), true

	case "JWTAuthResponse.threshold":
		if e.complexity.JWTAuthResponse.Threshold == nil {
			break
		}

		return e.complexity.JWTAuthResponse.Threshold(childComplexity), true

	case "JWTAuthResponse.token":
		if e.complexity.JWTAuthResponse.Token == nil {
			break
		}

		return e.complexity.JWTAuthResponse.Token(childComplexity), true

	case "Links.nextPage":
		if e.complexity.Links.NextPage == nil {
			break
		}

		return e.complexity.Links.NextPage(childComplexity), true

	case "Links.pageCursor":
		if e.complexity.Links.PageCursor == nil {
			break
		}

		return e.complexity.Links.PageCursor(childComplexity), true

	case "Mutation.deleteUser":
		if e.complexity.Mutation.DeleteUser == nil {
			break
		}

		args, err := ec.field_Mutation_deleteUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteUser(childComplexity, args["input"].(models.HTTPDeleteUserRequest)), true

	case "Mutation.depositFiat":
		if e.complexity.Mutation.DepositFiat == nil {
			break
		}

		args, err := ec.field_Mutation_depositFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DepositFiat(childComplexity, args["input"].(models.HTTPDepositCurrencyRequest)), true

	case "Mutation.exchangeOfferFiat":
		if e.complexity.Mutation.ExchangeOfferFiat == nil {
			break
		}

		args, err := ec.field_Mutation_exchangeOfferFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ExchangeOfferFiat(childComplexity, args["input"].(models.HTTPExchangeOfferRequest)), true

	case "Mutation.exchangeTransferFiat":
		if e.complexity.Mutation.ExchangeTransferFiat == nil {
			break
		}

		args, err := ec.field_Mutation_exchangeTransferFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.ExchangeTransferFiat(childComplexity, args["offerID"].(string)), true

	case "Mutation.loginUser":
		if e.complexity.Mutation.LoginUser == nil {
			break
		}

		args, err := ec.field_Mutation_loginUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.LoginUser(childComplexity, args["input"].(models1.UserLoginCredentials)), true

	case "Mutation.offerCrypto":
		if e.complexity.Mutation.OfferCrypto == nil {
			break
		}

		args, err := ec.field_Mutation_offerCrypto_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.OfferCrypto(childComplexity, args["input"].(models.HTTPCryptoOfferRequest)), true

	case "Mutation.openCrypto":
		if e.complexity.Mutation.OpenCrypto == nil {
			break
		}

		args, err := ec.field_Mutation_openCrypto_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.OpenCrypto(childComplexity, args["ticker"].(string)), true

	case "Mutation.openFiat":
		if e.complexity.Mutation.OpenFiat == nil {
			break
		}

		args, err := ec.field_Mutation_openFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.OpenFiat(childComplexity, args["currency"].(string)), true

	case "Mutation.refreshToken":
		if e.complexity.Mutation.RefreshToken == nil {
			break
		}

		return e.complexity.Mutation.RefreshToken(childComplexity), true

	case "Mutation.registerUser":
		if e.complexity.Mutation.RegisterUser == nil {
			break
		}

		args, err := ec.field_Mutation_registerUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RegisterUser(childComplexity, args["input"].(*models1.UserAccount)), true

	case "PriceQuote.amount":
		if e.complexity.PriceQuote.Amount == nil {
			break
		}

		return e.complexity.PriceQuote.Amount(childComplexity), true

	case "PriceQuote.clientID":
		if e.complexity.PriceQuote.ClientID == nil {
			break
		}

		return e.complexity.PriceQuote.ClientID(childComplexity), true

	case "PriceQuote.destinationAcc":
		if e.complexity.PriceQuote.DestinationAcc == nil {
			break
		}

		return e.complexity.PriceQuote.DestinationAcc(childComplexity), true

	case "PriceQuote.rate":
		if e.complexity.PriceQuote.Rate == nil {
			break
		}

		return e.complexity.PriceQuote.Rate(childComplexity), true

	case "PriceQuote.sourceAcc":
		if e.complexity.PriceQuote.SourceAcc == nil {
			break
		}

		return e.complexity.PriceQuote.SourceAcc(childComplexity), true

	case "Query.balanceAllFiat":
		if e.complexity.Query.BalanceAllFiat == nil {
			break
		}

		args, err := ec.field_Query_balanceAllFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.BalanceAllFiat(childComplexity, args["pageCursor"].(*string), args["pageSize"].(*int32)), true

	case "Query.balanceFiat":
		if e.complexity.Query.BalanceFiat == nil {
			break
		}

		args, err := ec.field_Query_balanceFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.BalanceFiat(childComplexity, args["currencyCode"].(string)), true

	case "Query.healthcheck":
		if e.complexity.Query.Healthcheck == nil {
			break
		}

		return e.complexity.Query.Healthcheck(childComplexity), true

	case "Query.transactionDetailsAllFiat":
		if e.complexity.Query.TransactionDetailsAllFiat == nil {
			break
		}

		args, err := ec.field_Query_transactionDetailsAllFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.TransactionDetailsAllFiat(childComplexity, args["input"].(models.FiatPaginatedTxDetailsRequest)), true

	case "Query.transactionDetailsFiat":
		if e.complexity.Query.TransactionDetailsFiat == nil {
			break
		}

		args, err := ec.field_Query_transactionDetailsFiat_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.TransactionDetailsFiat(childComplexity, args["transactionID"].(string)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCryptoOfferRequest,
		ec.unmarshalInputDeleteUserRequest,
		ec.unmarshalInputFiatDepositRequest,
		ec.unmarshalInputFiatExchangeOfferRequest,
		ec.unmarshalInputFiatPaginatedTxDetailsRequest,
		ec.unmarshalInputUserAccount,
		ec.unmarshalInputUserLoginCredentials,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/auth.graphqls", Input: `# JWT Authorization Response.
type JWTAuthResponse {
    token: String!
    expires: Int64!
    threshold: Int64!
}
`, BuiltIn: false},
	{Name: "../schema/crypto.graphqls", Input: `# CryptoOpenAccountResponse is the response returned when opening a Cryptocurrency account.
type CryptoOpenAccountResponse {
    clientID: String!
    ticker: String!
}

# CryptoOfferRequest is the request parameters to purchase or sell a Cryptocurrency.
input CryptoOfferRequest {
    sourceCurrency:         String!
    destinationCurrency:    String!
    sourceAmount:           Float!
    isPurchase:             Boolean!
}

# Requests that might alter the state of data in the database.
extend type Mutation {
    # openFiat is a request to open an account if it does not already exist.
    openCrypto(ticker: String!): CryptoOpenAccountResponse!

    # offerCrypto is a request for a Cryptocurrency purchase/sale quote. The exchange quote provided will expire after a fixed period.
    offerCrypto(input: CryptoOfferRequest!): FiatExchangeOfferResponse!
}
`, BuiltIn: false},
	{Name: "../schema/fiat.graphqls", Input: `# FiatOpenAccountResponse is the response returned
type FiatOpenAccountResponse {
    clientID: String!
    currency: String!
}

# FiatDepositResponse is the response to a Fiat currency deposit from an external source.
type FiatDepositResponse {
    txId: String!
    clientId: String!
    txTimestamp: String!
    balance: String!
    lastTx: String!
    currency: String!
}

# FiatExchangeOfferResponse is an offer to convert a source to destination currency in the source currency amount.
type FiatExchangeOfferResponse {
    priceQuote: PriceQuote!
    debitAmount: Float!
    offerID: String!
    expires: Int64!
}

# FiatExchangeTransferResponse
type FiatExchangeTransferResponse {
    sourceReceipt: FiatDepositResponse!
    destinationReceipt: FiatDepositResponse!
}

# FiatAccount are the Fiat account details associated with a specific Client ID.
type FiatAccount {
    currency:   String!
    balance:    Float!
    lastTx:     Float!
    lastTxTs:   String!
    createdAt:  String!
    clientID:   UUID!
}

# FiatJournal are the Fiat transactional records for a specific transaction.
type FiatJournal {
    currency:       String!
    amount:         Float!
    transactedAt:   String!
    clientID:       UUID!
    txID:           UUID!
}

# Links are links used in responses to retrieve pages of information.
type Links {
    nextPage:   String
    pageCursor: String
}

# FiatBalancesPaginated are all of the Fiat account balances retrieved via pagination.
type FiatBalancesPaginated {
    accountBalances:    [FiatAccount!]!
    links:              Links!
}

# FiatBalancesPaginated are all of the Fiat account balances retrieved via pagination.
type FiatTransactionsPaginated {
    transactions:   [FiatJournal!]!
    links:          Links!
}

# FiatDepositRequest is a request to deposit Fiat currency from an external source.
input FiatDepositRequest {
    amount:     Float!
    currency:   String!
}

# FiatExchangeOfferRequest is a request to exchange Fiat currency from one to another.
input FiatExchangeOfferRequest {
    sourceCurrency:         String!
    destinationCurrency:    String!
    sourceAmount:           Float!
}

# FiatPaginatedTxDetailsRequest request input parameters for all transaction records for a specific currency.
input FiatPaginatedTxDetailsRequest{
    currency:   String!
    pageSize:   String
    pageCursor: String
    timezone:   String
    month:      String
    year:       String
}

# Requests that might alter the state of data in the database.
extend type Mutation {
    # openFiat is a request to open an account if it does not already exist.
    openFiat(currency: String!): FiatOpenAccountResponse!

    # depositFiat is a request to deposit Fiat currency from an external source.
    depositFiat(input: FiatDepositRequest!): FiatDepositResponse!

    # exchangeOfferFiat is a request for an exchange quote. The exchange quote provided will expire after a fixed period.
    exchangeOfferFiat(input: FiatExchangeOfferRequest!): FiatExchangeOfferResponse!

    # exchangeTransferFiat will execute and complete a valid Fiat currency exchange offer.
    exchangeTransferFiat(offerID: String!): FiatExchangeTransferResponse!
}

extend type Query {
    # balanceFiat is a request to retrieve the balance for a specific Fiat currency.
    balanceFiat(currencyCode: String!): FiatAccount!

    # balanceAllFiat is a request to retrieve the balance for a specific Fiat currency.
    balanceAllFiat(pageCursor: String, pageSize: Int32): FiatBalancesPaginated!

    # transactionDetailsFiat is a request to retrieve the details for a specific transaction.
    transactionDetailsFiat(transactionID: String!): [FiatJournal!]!

    # transactionDetailsAllFiat is a request to retrieve the details for a specific transaction.
    transactionDetailsAllFiat(input: FiatPaginatedTxDetailsRequest!): FiatTransactionsPaginated!
}
`, BuiltIn: false},
	{Name: "../schema/healthcheck.graphqls", Input: `type Query {
    healthcheck: String!
}
`, BuiltIn: false},
	{Name: "../schema/redis.graphqls", Input: `# PriceQuote is the quote provided to the end-user requesting a transfer and will be stored in the Redis cache.
type PriceQuote {
    clientID: UUID!
    sourceAcc: String!
    destinationAcc: String!
    rate: Float!
    amount: Float!
}
`, BuiltIn: false},
	{Name: "../schema/scalars.graphqls", Input: `scalar Int32
scalar Int64
scalar UUID
`, BuiltIn: false},
	{Name: "../schema/user.graphqls", Input: `# User account information.
input UserAccount {
    firstname: String!
    lastname: String!
    email: String!
    userLoginCredentials: UserLoginCredentials!
}

# Login credentials.
input UserLoginCredentials {
    username: String!
    password: String!
}

# User account deletion request.
input DeleteUserRequest {
    username: String!
    password: String!
    confirmation: String!
}

# Requests that might alter the state of data in the database.
type Mutation {
    # Send a user registration request and receive a JWT authorization token in response.
    registerUser(input: UserAccount): JWTAuthResponse!

    # Send a user account deletion request.
    deleteUser(input: DeleteUserRequest!): String!

    # Send a user login request And receive a JWT authorization token in response. This has no side effects but is a
    # mutation to force sequential execution. This stops operations such as delete and refresh from being run in
    # parallel with a login.
    loginUser(input: UserLoginCredentials!): JWTAuthResponse!

    # Refreshes a users JWT if it is within the refresh time window.
    refreshToken: JWTAuthResponse!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
