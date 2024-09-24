package test

import (
	"amazon-sdk/ad"
	"fmt"
	"testing"
)

func TestAuthUrl(t *testing.T) {
	engine := ad.NewEngine()
	lwaUrl := engine.GetLwaURL()
	fmt.Println(lwaUrl)
}
func TestRefreshToken(t *testing.T) {
	engine := ad.NewEngine()
	err := engine.GetRefreshToken("ANvcywiLlkqsluMsTLyy")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", engine.Token)
}
func TestAccessToken(t *testing.T) {
	engine := ad.NewEngine()
	err := engine.GetAccessTokenFromEndpoint()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(engine.Token)
}
