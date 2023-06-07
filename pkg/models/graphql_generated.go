// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type FiatOpenAccountResponse struct {
	ClientID string `json:"clientID"`
	Currency string `json:"currency"`
}

type FiatPaginatedTxDetailsRequest struct {
	Currency   string  `json:"currency"`
	PageSize   *string `json:"pageSize,omitempty"`
	PageCursor *string `json:"pageCursor,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
	Month      *string `json:"month,omitempty"`
	Year       *string `json:"year,omitempty"`
}
