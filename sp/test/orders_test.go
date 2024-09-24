package test

import (
	"amazon-sdk/sp"
	"fmt"
	"testing"
)

func TestGetOrder(t *testing.T) {
	engine := sp.NewEngine()
	var orderId = "112-1244590-3672238"
	orderUrl := "/orders/v0/orders/" + orderId
	resp, err := engine.SignRequest("GET", orderUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(*resp))
}
