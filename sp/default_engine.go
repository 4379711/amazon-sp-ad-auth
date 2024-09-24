package sp

import (
	"amazon-sdk/common"
	"net/http"
	"time"
)

func NewEngine() *Engine {
	token := common.Token{
		RefreshToken: "Atzr|IwEBIEaJdOLREKopls63cxDrofSeNTzmP0gAmD3hcofyCoWEHxk3XRpI8oTsCBA-ws3Zwd4XaHnMQiVWipG1VWxQhnD-j04mko4SVoxAPCZTfD0GyHRHqCguvNwztDWEZFYURz6ZALDQPVx5o1L1kPv0hnOssFJivIcuiG5_6hoXEgi5WRMQt7oBgO9ShHTjr-2p-cnIHyStKy7o9FIY3-wc4JyFzWPk19aa3WDC7QbOHLdUFcG2Zg8X-KR3KaqiloFXU",
	}

	seller := common.Seller{
		Region:     "na",
		SellerName: "Rili Shop",
		SellerId:   "A9R00C",
	}

	credentials := common.App{
		AppID:        "amzn1.sp.solution.691fc4bb-c5a13ef4cb99",
		ClientID:     "amzn1.application-oa2-clien04b1e34144753ec",
		ClientSecret: "amzn1.oa2-cs.v1.76373d4e7ae3486b84ba44c0bc73534c5d88",
		RedirectURL:  "https://baidu.com",
		Beta:         true,
	}
	commonAuth := common.Auth{
		Seller: &seller,
		App:    &credentials,
		Token:  &token,
	}

	auth := Auth{&commonAuth}

	client := http.DefaultClient
	client.Timeout = time.Second * 180

	return &Engine{Auth: &auth, HttpClient: client}
}
