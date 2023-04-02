// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package postgres

import (
	"database/sql/driver"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Currency string

const (
	CurrencyAED     Currency = "AED"
	CurrencyAFN     Currency = "AFN"
	CurrencyALL     Currency = "ALL"
	CurrencyAMD     Currency = "AMD"
	CurrencyANG     Currency = "ANG"
	CurrencyAOA     Currency = "AOA"
	CurrencyARS     Currency = "ARS"
	CurrencyAUD     Currency = "AUD"
	CurrencyAWG     Currency = "AWG"
	CurrencyAZN     Currency = "AZN"
	CurrencyBAM     Currency = "BAM"
	CurrencyBBD     Currency = "BBD"
	CurrencyBDT     Currency = "BDT"
	CurrencyBGN     Currency = "BGN"
	CurrencyBHD     Currency = "BHD"
	CurrencyBIF     Currency = "BIF"
	CurrencyBMD     Currency = "BMD"
	CurrencyBND     Currency = "BND"
	CurrencyBOB     Currency = "BOB"
	CurrencyBRL     Currency = "BRL"
	CurrencyBSD     Currency = "BSD"
	CurrencyBTN     Currency = "BTN"
	CurrencyBWP     Currency = "BWP"
	CurrencyBYN     Currency = "BYN"
	CurrencyBZD     Currency = "BZD"
	CurrencyCAD     Currency = "CAD"
	CurrencyCDF     Currency = "CDF"
	CurrencyCHF     Currency = "CHF"
	CurrencyCLP     Currency = "CLP"
	CurrencyCNY     Currency = "CNY"
	CurrencyCOP     Currency = "COP"
	CurrencyCRC     Currency = "CRC"
	CurrencyCUC     Currency = "CUC"
	CurrencyCUP     Currency = "CUP"
	CurrencyCVE     Currency = "CVE"
	CurrencyCZK     Currency = "CZK"
	CurrencyDJF     Currency = "DJF"
	CurrencyDKK     Currency = "DKK"
	CurrencyDOP     Currency = "DOP"
	CurrencyDZD     Currency = "DZD"
	CurrencyEGP     Currency = "EGP"
	CurrencyERN     Currency = "ERN"
	CurrencyETB     Currency = "ETB"
	CurrencyEUR     Currency = "EUR"
	CurrencyFJD     Currency = "FJD"
	CurrencyFKP     Currency = "FKP"
	CurrencyGBP     Currency = "GBP"
	CurrencyGEL     Currency = "GEL"
	CurrencyGGP     Currency = "GGP"
	CurrencyGHS     Currency = "GHS"
	CurrencyGIP     Currency = "GIP"
	CurrencyGMD     Currency = "GMD"
	CurrencyGNF     Currency = "GNF"
	CurrencyGTQ     Currency = "GTQ"
	CurrencyGYD     Currency = "GYD"
	CurrencyHKD     Currency = "HKD"
	CurrencyHNL     Currency = "HNL"
	CurrencyHRK     Currency = "HRK"
	CurrencyHTG     Currency = "HTG"
	CurrencyHUF     Currency = "HUF"
	CurrencyIDR     Currency = "IDR"
	CurrencyILS     Currency = "ILS"
	CurrencyIMP     Currency = "IMP"
	CurrencyINR     Currency = "INR"
	CurrencyIQD     Currency = "IQD"
	CurrencyIRR     Currency = "IRR"
	CurrencyISK     Currency = "ISK"
	CurrencyJEP     Currency = "JEP"
	CurrencyJMD     Currency = "JMD"
	CurrencyJOD     Currency = "JOD"
	CurrencyJPY     Currency = "JPY"
	CurrencyKES     Currency = "KES"
	CurrencyKGS     Currency = "KGS"
	CurrencyKHR     Currency = "KHR"
	CurrencyKMF     Currency = "KMF"
	CurrencyKPW     Currency = "KPW"
	CurrencyKRW     Currency = "KRW"
	CurrencyKWD     Currency = "KWD"
	CurrencyKYD     Currency = "KYD"
	CurrencyKZT     Currency = "KZT"
	CurrencyLAK     Currency = "LAK"
	CurrencyLBP     Currency = "LBP"
	CurrencyLKR     Currency = "LKR"
	CurrencyLRD     Currency = "LRD"
	CurrencyLSL     Currency = "LSL"
	CurrencyLYD     Currency = "LYD"
	CurrencyMAD     Currency = "MAD"
	CurrencyMDL     Currency = "MDL"
	CurrencyMGA     Currency = "MGA"
	CurrencyMKD     Currency = "MKD"
	CurrencyMMK     Currency = "MMK"
	CurrencyMNT     Currency = "MNT"
	CurrencyMOP     Currency = "MOP"
	CurrencyMRU     Currency = "MRU"
	CurrencyMUR     Currency = "MUR"
	CurrencyMVR     Currency = "MVR"
	CurrencyMWK     Currency = "MWK"
	CurrencyMXN     Currency = "MXN"
	CurrencyMYR     Currency = "MYR"
	CurrencyMZN     Currency = "MZN"
	CurrencyNAD     Currency = "NAD"
	CurrencyNGN     Currency = "NGN"
	CurrencyNIO     Currency = "NIO"
	CurrencyNOK     Currency = "NOK"
	CurrencyNPR     Currency = "NPR"
	CurrencyNZD     Currency = "NZD"
	CurrencyOMR     Currency = "OMR"
	CurrencyPAB     Currency = "PAB"
	CurrencyPEN     Currency = "PEN"
	CurrencyPGK     Currency = "PGK"
	CurrencyPHP     Currency = "PHP"
	CurrencyPKR     Currency = "PKR"
	CurrencyPLN     Currency = "PLN"
	CurrencyPYG     Currency = "PYG"
	CurrencyQAR     Currency = "QAR"
	CurrencyRON     Currency = "RON"
	CurrencyRSD     Currency = "RSD"
	CurrencyRUB     Currency = "RUB"
	CurrencyRWF     Currency = "RWF"
	CurrencySAR     Currency = "SAR"
	CurrencySBD     Currency = "SBD"
	CurrencySCR     Currency = "SCR"
	CurrencySDG     Currency = "SDG"
	CurrencySEK     Currency = "SEK"
	CurrencySGD     Currency = "SGD"
	CurrencySHP     Currency = "SHP"
	CurrencySLL     Currency = "SLL"
	CurrencySOS     Currency = "SOS"
	CurrencySPL     Currency = "SPL"
	CurrencySRD     Currency = "SRD"
	CurrencySTN     Currency = "STN"
	CurrencySVC     Currency = "SVC"
	CurrencySYP     Currency = "SYP"
	CurrencySZL     Currency = "SZL"
	CurrencyTHB     Currency = "THB"
	CurrencyTJS     Currency = "TJS"
	CurrencyTMT     Currency = "TMT"
	CurrencyTND     Currency = "TND"
	CurrencyTOP     Currency = "TOP"
	CurrencyTRY     Currency = "TRY"
	CurrencyTTD     Currency = "TTD"
	CurrencyTVD     Currency = "TVD"
	CurrencyTWD     Currency = "TWD"
	CurrencyTZS     Currency = "TZS"
	CurrencyUAH     Currency = "UAH"
	CurrencyUGX     Currency = "UGX"
	CurrencyUSD     Currency = "USD"
	CurrencyUYU     Currency = "UYU"
	CurrencyUZS     Currency = "UZS"
	CurrencyVEF     Currency = "VEF"
	CurrencyVND     Currency = "VND"
	CurrencyVUV     Currency = "VUV"
	CurrencyWST     Currency = "WST"
	CurrencyXAF     Currency = "XAF"
	CurrencyXCD     Currency = "XCD"
	CurrencyXDR     Currency = "XDR"
	CurrencyXOF     Currency = "XOF"
	CurrencyXPF     Currency = "XPF"
	CurrencyYER     Currency = "YER"
	CurrencyZAR     Currency = "ZAR"
	CurrencyZMW     Currency = "ZMW"
	CurrencyZWD     Currency = "ZWD"
	CurrencyDEPOSIT Currency = "DEPOSIT"
	CurrencyCRYPTO  Currency = "CRYPTO"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported scan type for Currency: %T", src)
	}
	return nil
}

type NullCurrency struct {
	Currency Currency
	Valid    bool // Valid is true if Currency is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrency) Scan(value interface{}) error {
	if value == nil {
		ns.Currency, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currency.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrency) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Currency), nil
}

func (e Currency) Valid() bool {
	switch e {
	case CurrencyAED,
		CurrencyAFN,
		CurrencyALL,
		CurrencyAMD,
		CurrencyANG,
		CurrencyAOA,
		CurrencyARS,
		CurrencyAUD,
		CurrencyAWG,
		CurrencyAZN,
		CurrencyBAM,
		CurrencyBBD,
		CurrencyBDT,
		CurrencyBGN,
		CurrencyBHD,
		CurrencyBIF,
		CurrencyBMD,
		CurrencyBND,
		CurrencyBOB,
		CurrencyBRL,
		CurrencyBSD,
		CurrencyBTN,
		CurrencyBWP,
		CurrencyBYN,
		CurrencyBZD,
		CurrencyCAD,
		CurrencyCDF,
		CurrencyCHF,
		CurrencyCLP,
		CurrencyCNY,
		CurrencyCOP,
		CurrencyCRC,
		CurrencyCUC,
		CurrencyCUP,
		CurrencyCVE,
		CurrencyCZK,
		CurrencyDJF,
		CurrencyDKK,
		CurrencyDOP,
		CurrencyDZD,
		CurrencyEGP,
		CurrencyERN,
		CurrencyETB,
		CurrencyEUR,
		CurrencyFJD,
		CurrencyFKP,
		CurrencyGBP,
		CurrencyGEL,
		CurrencyGGP,
		CurrencyGHS,
		CurrencyGIP,
		CurrencyGMD,
		CurrencyGNF,
		CurrencyGTQ,
		CurrencyGYD,
		CurrencyHKD,
		CurrencyHNL,
		CurrencyHRK,
		CurrencyHTG,
		CurrencyHUF,
		CurrencyIDR,
		CurrencyILS,
		CurrencyIMP,
		CurrencyINR,
		CurrencyIQD,
		CurrencyIRR,
		CurrencyISK,
		CurrencyJEP,
		CurrencyJMD,
		CurrencyJOD,
		CurrencyJPY,
		CurrencyKES,
		CurrencyKGS,
		CurrencyKHR,
		CurrencyKMF,
		CurrencyKPW,
		CurrencyKRW,
		CurrencyKWD,
		CurrencyKYD,
		CurrencyKZT,
		CurrencyLAK,
		CurrencyLBP,
		CurrencyLKR,
		CurrencyLRD,
		CurrencyLSL,
		CurrencyLYD,
		CurrencyMAD,
		CurrencyMDL,
		CurrencyMGA,
		CurrencyMKD,
		CurrencyMMK,
		CurrencyMNT,
		CurrencyMOP,
		CurrencyMRU,
		CurrencyMUR,
		CurrencyMVR,
		CurrencyMWK,
		CurrencyMXN,
		CurrencyMYR,
		CurrencyMZN,
		CurrencyNAD,
		CurrencyNGN,
		CurrencyNIO,
		CurrencyNOK,
		CurrencyNPR,
		CurrencyNZD,
		CurrencyOMR,
		CurrencyPAB,
		CurrencyPEN,
		CurrencyPGK,
		CurrencyPHP,
		CurrencyPKR,
		CurrencyPLN,
		CurrencyPYG,
		CurrencyQAR,
		CurrencyRON,
		CurrencyRSD,
		CurrencyRUB,
		CurrencyRWF,
		CurrencySAR,
		CurrencySBD,
		CurrencySCR,
		CurrencySDG,
		CurrencySEK,
		CurrencySGD,
		CurrencySHP,
		CurrencySLL,
		CurrencySOS,
		CurrencySPL,
		CurrencySRD,
		CurrencySTN,
		CurrencySVC,
		CurrencySYP,
		CurrencySZL,
		CurrencyTHB,
		CurrencyTJS,
		CurrencyTMT,
		CurrencyTND,
		CurrencyTOP,
		CurrencyTRY,
		CurrencyTTD,
		CurrencyTVD,
		CurrencyTWD,
		CurrencyTZS,
		CurrencyUAH,
		CurrencyUGX,
		CurrencyUSD,
		CurrencyUYU,
		CurrencyUZS,
		CurrencyVEF,
		CurrencyVND,
		CurrencyVUV,
		CurrencyWST,
		CurrencyXAF,
		CurrencyXCD,
		CurrencyXDR,
		CurrencyXOF,
		CurrencyXPF,
		CurrencyYER,
		CurrencyZAR,
		CurrencyZMW,
		CurrencyZWD,
		CurrencyDEPOSIT,
		CurrencyCRYPTO:
		return true
	}
	return false
}

type FiatAccount struct {
	Currency  Currency           `json:"currency"`
	Balance   decimal.Decimal    `json:"balance"`
	LastTx    decimal.Decimal    `json:"lastTx"`
	LastTxTs  pgtype.Timestamptz `json:"lastTxTs"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	ClientID  uuid.UUID          `json:"clientID"`
}

type FiatJournal struct {
	Currency     Currency           `json:"currency"`
	Amount       decimal.Decimal    `json:"amount"`
	TransactedAt pgtype.Timestamptz `json:"transactedAt"`
	ClientID     uuid.UUID          `json:"clientID"`
	TxID         uuid.UUID          `json:"txID"`
}

type User struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	ClientID  uuid.UUID `json:"clientID"`
	IsDeleted bool      `json:"isDeleted"`
}
