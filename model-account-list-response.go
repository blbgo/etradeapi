package etradeapi

// AccountListResponse is the root object returned by "/v1/accounts/list" endpoint
type AccountListResponse struct {
	AccountListResponse Accounts
}

// Accounts is the AccountListResponse fields value
type Accounts struct {
	Accounts AccountsInner
}

// AccountsInner is the Accounts fields value
type AccountsInner struct {
	Account []Account
}

// Account is the array element of the Account fields value
type Account struct {
	AccountID       string // user friendly ID shown on UIs example "64983239"
	AccountIDKey    string // internal ID needed for other calls example "Oy0YQdGsas2IN-_pVwGAew"
	AccountMode     string // "MARGIN" or "CASH"
	AccountDesc     string
	AccountName     string
	AccountType     string
	InstitutionType string
	AccountStatus   string
	ClosedDate      int64
}
