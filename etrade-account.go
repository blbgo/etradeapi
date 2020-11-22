package etradeapi

import (
	"fmt"
	"net/url"
	"time"
)

func (r *etrade) Accounts() ([]Account, error) {
	accountListResponse := &AccountListResponse{}
	err := r.MakeGetRequest("/v1/accounts/list", nil, accountListResponse)
	if err != nil {
		return nil, err
	}

	return accountListResponse.AccountListResponse.Accounts.Account, nil
}

func (r *etrade) Transactions(accountIDKey string, start time.Time) ([]Transaction, error) {
	var result []Transaction
	startString := start.Format(timeFormatMMDDYYYY)
	endString := time.Now().Format(timeFormatMMDDYYYY)
	urlValues := url.Values{
		"sortOrder": {"ASC"},
		"startDate": {startString},
		"endDate":   {endString},
		"count":     {"50"},
	}
	count := 50
	lastTransID := "NONE"
	for count == 50 {
		transactionListResponse := &TransactionListResponse{}
		err := r.MakeGetRequest(
			fmt.Sprintf("/v1/accounts/%v/transactions", accountIDKey),
			urlValues,
			transactionListResponse,
		)
		if err != nil {
			return nil, err
		}
		transactions := transactionListResponse.TransactionListResponse.Transaction
		// testing indicates that when paging the next page starts with the first transaction from
		// the previous page.  So lets detect that and drop it from the next page.
		if lastTransID == transactions[0].TransactionID {
			transactions = transactions[1:]
		}
		lastTransID = transactions[len(transactions)-1].TransactionID
		result = append(result, transactions...)
		urlValues["marker"] = []string{transactionListResponse.TransactionListResponse.Marker}
		count = transactionListResponse.TransactionListResponse.TransactionCount
		time.Sleep(time.Second * 2)
	}

	return result, nil
}
