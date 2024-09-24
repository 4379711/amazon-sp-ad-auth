package test

import (
	"amazon-sdk/ad"
	"fmt"
	"testing"
)

func TestPortfolios(t *testing.T) {
	engine := ad.NewEngine()
	resp, err := engine.SignRequest("GET", "/v2/portfolios", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(*resp))
}
