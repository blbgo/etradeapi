package etradeapi

// Product represents a specific equity, option, or other instrument
type Product struct {
	Symbol       string
	SecurityType string
	ExpiryYear   int32
	ExpiryMonth  int32
	ExpiryDay    int32
	StrikePrice  float32
	CallPut      string

	// in doc but not seen in real response
	SecuritySubType string
	ExpiryType      string
}
