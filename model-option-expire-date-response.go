package etradeapi

// OptionExpireDateResponse is the root object returned by "/v1/market/optionexpiredate" endpoint
type OptionExpireDateResponse struct {
	OptionExpireDateResponse ExpirationDates
}

// ExpirationDates is the OptionExpireDateResponse fields value
type ExpirationDates struct {
	ExpirationDate []ExpirationDate
}

// ExpirationDate is the array element of the ExpirationDate fields value
type ExpirationDate struct {
	Year       int32
	Month      int32
	Day        int32
	ExpiryType string
}
