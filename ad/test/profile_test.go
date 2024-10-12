package test

import (
	"amazon-sdk/ad"
	"fmt"
	"testing"
)

func TestProfile(t *testing.T) {
	engine := ad.NewEngine()
	resp, err := engine.CallApi("GET", "/v2/profiles", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(*resp))
}
