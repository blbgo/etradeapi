package etradeapi

// TransactionListResponse is the root object returned by "/v1/accounts/{accountKey}/transactions"
// endpoint
type TransactionListResponse struct {
	TransactionListResponse Transactions
}

// Transactions is the TransactionListResponse fields value
type Transactions struct {
	Next             string
	Marker           string
	PageMarkers      string
	MoreTransactions bool
	TransactionCount int
	TotalCount       int
	Transaction      []Transaction
}

// Transaction is the array element of the Transaction fields value
type Transaction struct {
	TransactionID   string
	AccountID       string
	TransactionDate int64
	PostDate        int64
	Amount          float64
	Description     string
	Description2    string // not in doc but showed up in a transfer type transaction
	TransactionType string
	Memo            string
	ImageFlag       bool
	InstType        string
	StoreID         int
	Brokerage       Brokerage
	DetailsURI      string
}

// Brokerage is the Brokerage fields value
type Brokerage struct {
	Product            Product
	Quantity           float64
	Price              float64
	SettlementCurrency string
	PaymentCurrency    string
	Fee                float64
	DisplaySymbol      string
	SettlementDate     int64
}
