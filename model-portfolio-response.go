package etradeapi

// PortfolioResponse is the root object returned by "/v1/accounts/{accountKey}/portfolio" endpoint
type PortfolioResponse struct {
	PortfolioResponse AccountPortfolios
}

// AccountPortfolios is the PortfolioResponse fields value
type AccountPortfolios struct {
	AccountPortfolio []AccountPortfolio
}

// AccountPortfolio is the array element of the AccountPortfolio fields value
type AccountPortfolio struct {
	AccountID  string
	Position   []Position
	TotalPages int
}

// Position is the array element of the Position fields value
type Position struct {
	PositionID        int64
	SymbolDescription string
	DateAcquired      int64
	PricePaid         float32
	Commissions       float32
	OtherFees         float32
	Quantity          float32
	PositionIndicator string
	PositionType      string
	DaysGain          float32
	DaysGainPct       float32
	MarketValue       float32
	TotalCost         float32
	TotalGain         float32
	TotalGainPct      float32
	PctOfPortfolio    float32
	CostPerShare      float32
	TodayCommissions  float32
	TodayFees         float32
	TodayPricePaid    float32
	TodayQuantity     float32
	LotsDetails       string
	QuoteDetails      string
	Product           Product
	Quick             QuickView
	Performance       Performance
	Fundamental       Fundamental
}

// QuickView is the Quick fields value
type QuickView struct {
	LastTrade     float32
	LastTradeTime int64
	Change        float32
	ChangePct     float32
	Volume        int64

	// in doc but not seen in real response
	QuoteStatus             string
	sevenDayCurrentYield    float64
	AnnualTotalReturn       float64
	WeightedAverageMaturity float64
}

// Performance is the Performance fields value
type Performance struct {
	Change    float32
	ChangePct float32
	LastTrade float32

	// in doc but not seen in real response
	DaysGain      float32
	TotalGain     float32
	TotalGainPct  float32
	MarketValue   float32
	QuoteStatus   string
	LastTradeTime int64
}

// Fundamental is the Fundamental fields value
type Fundamental struct {
	LastTrade     float32
	LastTradeTime int64
	Change        float32
	ChangePct     float32
	PeRatio       float32
	Eps           float32
	Dividend      float32
	DivYield      float32
	MarketCap     float64
	Week52Range   string

	// in doc but not seen in real response
	quoteStatus string
}
