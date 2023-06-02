package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type DummyTokenConfig struct {
	TokenOnboarding string `split_words:"true" json:"DUMMY_TOKEN_ONBOARDING"`
	TokenCustomer   string `split_words:"true" json:"DUMMY_TOKEN_CUSTOMER"`
	TokenTransfers  string `split_words:"true" json:"DUMMY_TOKEN_TRANSFERS"`
	TokenExpired    string `split_words:"true" json:"DUMMY_TOKEN_EXPIRED"`

	//SANDBOX:true Token
	TokenOnboardingSandbox string `split_words:"true" json:"DUMMY_TOKEN_ONBOARDING_SANDBOX"`
	TokenCustomerSandbox   string `split_words:"true" json:"DUMMY_TOKEN_CUSTOMER_SANDBOX"`
	TokenTransfersSandbox  string `split_words:"true" json:"DUMMY_TOKEN_TRANSFERS_SANDBOX"`
}

var DummyToken *DummyTokenConfig

func loadDummyTokenConfig() {
	DummyToken = &DummyTokenConfig{}
	err := envconfig.Process("dummy", DummyToken)
	if err != nil {
		log.Fatal(err.Error())
	}
}
