package etradeapi

// QuoteResponse is the root object returned by "/v1/market/quote" endpoint
type QuoteResponse struct {
	QuoteResponse Quote
}

// Quote is the QuoteResponse fields value
type Quote struct {
	QuoteData []QuoteData

	// in doc but not seen in real response
	Messages Messages
}

// QuoteData is the array element of the QuoteData fields value
type QuoteData struct {
	All         AllQuoteDetails
	Product     Product
	DateTime    string
	DateTimeUTC int64
	QuoteStatus string
	AhFlag      string

	// in doc but not seen in real response
	ErrorMessage   string
	TimeZone       string
	DstFlag        bool
	HasMiniOptions bool
}

// Messages is the Messages fields value
type Messages struct {
	Message []Message
}

// Message is the array element of the Message fields value
type Message struct {
	Description string
	Code        int32
	Type        string
}

// AllQuoteDetails is the All fields value
type AllQuoteDetails struct {
	AdjustedFlag          bool
	Ask                   float64
	AskSize               int64
	AskTime               string
	Bid                   float64
	BidExchange           string
	BidSize               int64
	BidTime               string
	ChangeClose           float64
	ChangeClosePercentage float64
	CompanyName           string
	DaysToExpiration      int64
	DirLast               string
	Dividend              float64
	Eps                   float64
	EstEarnings           float64
	ExDividendDate        int64
	High                  float64
	High52                float64
	LastTrade             float64
	Low                   float64
	Low52                 float64
	Open                  float64
	OpenInterest          int64
	OptionStyle           string
	PreviousClose         float64
	PreviousDayVolume     int64
	PrimaryExchange       string
	SymbolDescription     string
	TotalVolume           int64
	Upc                   int64
	CashDeliverable       float64
	MarketCap             float64
	SharesOutstanding     float64
	NextEarningDate       string
	Beta                  float64
	Yield                 float64
	DeclaredDividend      float64
	DividendPayableDate   int64
	Pe                    float64
	Week52LowDate         int64
	Week52HiDate          int64
	IntrinsicValue        float64
	TimePremium           float64
	OptionMultiplier      float64
	ContractSize          float64
	ExpirationDate        int64
	TimeOfLastTrade       int64
	AverageVolume         int64

	// in doc but not seen in real response
	AnnualDividend          float64
	AskExchange             string
	HighAsk                 float64
	HighBid                 float64
	LowAsk                  float64
	LowBid                  float64
	NumberOfTrades          int64
	optionUnderlier         string
	optionUnderlierExchange string
	TodayClose              float64
	Volume10Day             int64
	//optionDeliverableList []OptionDeliverable
	MarketCloseBidSize int64
	MarketCloseAskSize int64
	MarketCloseVolume  int64
	//EhQuote ExtendedHourQuoteDetail
	OptionPreviousBidPrice float64
	OptionPreviousAskPrice float64
	OsiKey                 string
}
