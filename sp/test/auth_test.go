package test

import (
	"amazon-sdk/sp"
	"fmt"
	"testing"
)


func TestAuthUrl(t *testing.T) {
	engine := sp.NewEngine()
	lwaUrl := engine.GetLwaURL()
	fmt.Println(lwaUrl)
}
func TestRefreshToken(t *testing.T) {
	engine := sp.NewEngine()
	err := engine.GetRefreshToken("ANGhsekcomsiSqgssXyS")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", engine.Token)
}
func TestAccessToken(t *testing.T) {
	engine := sp.NewEngine()
	err := engine.GetAccessTokenFromEndpoint()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(engine.Token)
}
