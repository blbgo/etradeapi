package etradeapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gomodule/oauth1/oauth"

	"github.com/blbgo/general"
	"github.com/blbgo/record/recordlog"
)

// Etrade provides access to the etrade API
type Etrade interface {
	Status() string
	AccessTokenOK() bool

	AuthURL() (string, error)
	CheckCode(code string) error

	MakeGetRequest(url string, form url.Values, result interface{}) error
	MakePostRequest(url string, form url.Values, body interface{}, result interface{}) error

	EtradeTimeAsGoTime(etradeTime int64) time.Time

	Accounts() ([]Account, error)
	Transactions(accountIDKey string, start time.Time) ([]Transaction, error)
}

type etrade struct {
	recordlog.RecordLog
	general.DumperFactory
	http.RoundTripper
	//httpClient httpinterface.Client
	general.PersistentState
	oauth.Client

	baseServerURL string

	requestCred *oauth.Credentials

	etradeState
}

const stateName = "etradeAuth"
const timeFormatMMDDYYYY = "01022006"

type etradeState struct {
	AccessCred *oauth.Credentials
	AccessTime time.Time
}

// New provides an implementation of the Etrade interface
func New(
	config Config,
	recordLog recordlog.RecordLog,
	dumperFactory general.DumperFactory,
	roundTripper http.RoundTripper,
	//httpClient httpinterface.Client,
	persistentState general.PersistentState,
) Etrade {
	return &etrade{
		RecordLog:     recordLog,
		DumperFactory: dumperFactory,
		//httpClient:      httpClient,
		RoundTripper:    roundTripper,
		PersistentState: persistentState,
		Client: oauth.Client{
			Credentials: oauth.Credentials{
				Token:  config.ConsumerKey(),
				Secret: config.ConsumerSecret(),
			},
			TemporaryCredentialRequestURI: config.RequestTokenURL(),
			ResourceOwnerAuthorizationURI: config.AuthorizeURL(),
			TokenRequestURI:               config.AccessTokenURL(),
			RenewCredentialRequestURI:     config.RenewTokenURL(),
		},
		baseServerURL: config.BaseServerURL(),
	}
}

func (r *etrade) Status() string {
	if r.AccessCred == nil {
		err := r.PersistentState.Retrieve(stateName, &r.etradeState)
		if err != nil {
			return "No access token"
		}
	}
	d := time.Since(r.AccessTime).Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprint("Access token acquired ", int64(h), " hours and ", int64(m), " minutes ago")
}

func (r *etrade) AccessTokenOK() bool {
	if r.AccessCred == nil {
		r.PersistentState.Retrieve(stateName, &r.etradeState)
	}
	return r.AccessCred != nil && time.Since(r.AccessTime) < (time.Minute*110)
}

func (r *etrade) AuthURL() (string, error) {
	var err error
	r.requestCred, err = r.Client.RequestTemporaryCredentials(nil, "oob", nil)
	if err != nil {
		return "", err
	}

	authorizationURL := r.Client.AuthorizationURL(
		r.requestCred,
		url.Values{"key": {r.Client.Credentials.Token}},
	)

	// etrade used paramiter "token" instead of the standard "oauth_token"
	return strings.ReplaceAll(authorizationURL, "oauth_token", "token"), nil
}

func (r *etrade) CheckCode(code string) error {
	var err error
	r.AccessCred, _, err = r.Client.RequestToken(nil, r.requestCred, code)
	if err != nil {
		return err
	}
	r.AccessTime = time.Now()
	return r.PersistentState.Save(stateName, &r.etradeState)
}

func (r *etrade) EtradeTimeAsGoTime(etradeTime int64) time.Time {
	return time.Unix(etradeTime/1000, 0).UTC()
}
