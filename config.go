package etradeapi

import (
	"github.com/blbgo/general"
)

// Config must be implemented and provided to New
type Config interface {
	ConsumerKey() string
	ConsumerSecret() string
	RequestTokenURL() string
	AccessTokenURL() string
	RenewTokenURL() string
	AuthorizeURL() string
	BaseServerURL() string
	//User() string
	//Password() string
}

type config struct {
	ConsumerKeyValue     string
	ConsumerSecretValue  string
	RequestTokenURLValue string
	AccessTokenURLValue  string
	RenewTokenURLValue   string
	AuthorizeURLValue    string
	BaseServerURLValue   string
	//UserValue            string
	//PasswordValue        string
}

// NewConfig provides a Config implementation based on general.Config
func NewConfig(c general.Config) (Config, error) {
	r := &config{}
	var err error

	r.ConsumerKeyValue, err = c.Value("EtradeAPI", "ConsumerKey")
	if err != nil {
		return nil, err
	}
	r.ConsumerSecretValue, err = c.Value("EtradeAPI", "ConsumerSecret")
	if err != nil {
		return nil, err
	}
	r.RequestTokenURLValue, err = c.Value("EtradeAPI", "RequestTokenURL")
	if err != nil {
		return nil, err
	}
	r.AccessTokenURLValue, err = c.Value("EtradeAPI", "AccessTokenURL")
	if err != nil {
		return nil, err
	}
	r.RenewTokenURLValue, err = c.Value("EtradeAPI", "RenewTokenURL")
	if err != nil {
		return nil, err
	}
	r.AuthorizeURLValue, err = c.Value("EtradeAPI", "AuthorizeURL")
	if err != nil {
		return nil, err
	}
	r.BaseServerURLValue, err = c.Value("EtradeAPI", "BaseServerURL")
	if err != nil {
		return nil, err
	}
	//r.UserValue, err = c.Value("EtradeAPI", "User")
	//if err != nil {
	//	return nil, err
	//}
	//r.PasswordValue, err = c.Value("EtradeAPI", "Password")
	//if err != nil {
	//	return nil, err
	//}

	return r, nil
}

// ConsumerKey method of etradeapi.AuthConfig
func (r *config) ConsumerKey() string {
	return r.ConsumerKeyValue
}

// ConsumerSecret method of etradeapi.AuthConfig
func (r *config) ConsumerSecret() string {
	return r.ConsumerSecretValue
}

// RequestTokenURL method of etradeapi.AuthConfig
func (r *config) RequestTokenURL() string {
	return r.RequestTokenURLValue
}

// AccessTokenURL method of etradeapi.AuthConfig
func (r *config) AccessTokenURL() string {
	return r.AccessTokenURLValue
}

// RenewTokenURL method of etradeapi.AuthConfig
func (r *config) RenewTokenURL() string {
	return r.RenewTokenURLValue
}

// AuthorizeURL method of etradeapi.AuthConfig
func (r *config) AuthorizeURL() string {
	return r.AuthorizeURLValue
}

// BaseServerURL method of etradeapi.AuthConfig
func (r *config) BaseServerURL() string {
	return r.BaseServerURLValue
}

// User method of etradeapi.AuthConfig
//func (r *config) User() string {
//	return r.UserValue
//}

// Password method of etradeapi.AuthConfig
//func (r *config) Password() string {
//	return r.PasswordValue
//}
