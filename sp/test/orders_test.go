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
	resp, err := engine.CallApi("GET", orderUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(*resp))
}

func TestFinance(t *testing.T) {
	engine := sp.NewEngine()
	url := "/finances/2024-06-19/transactions?postedAfter=2024-01-01T00:00:00.000Z"
	resp, err := engine.CallApi("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(*resp))
}
