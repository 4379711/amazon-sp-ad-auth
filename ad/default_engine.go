package ad

import (
	"amazon-sdk/common"
	"net/http"
	"time"
)

func NewEngine() *Engine {
	token := common.Token{
		RefreshToken: "Af5-KLl7aAbsKejhyhXomOZ0B_LGLYS-8dZXON9BxbyxL6UFJeeZoVL9A1O4_LO9-MSRgwgZcjtSeprx0pq121_hmHUZdN6oJ7njFhNvEnU70NWgmCeToix7exEpbh-gbIzU-zpYllWMGSsoabqssEJwul5juwDgABEmr8se3pXZZfkCuSp5nbkDAbAQ-DnbWevaHNpLnTT7ik_qGBWjLqse8IjxhdY8hmLtjIejmWJ4mV-pJSXEHNSL6pwIMuTa4Xgi5LixNrF5tW811aO-WU1amWhX60GRwM5UzRc4szJk3WIgyPMW4b4s9AVw_E1tIfm-EQB0RhvrrA5-Wdl8psFNRwyrGEZOsWqCTmeIKDx0S8_ao5zQT2bW32HJ39nqesRiImyMPXpJ3aLtQGNYv4e1hTMs1YWSaPiR",
	}
	seller := common.Seller{
		Region:     "na",
		SellerName: "Rili Shop",
		SellerId:   "A9R02BM67",
		ProfileId:  "1305508",
	}

	credentials := common.App{
		ClientID:     "amzn1.a9c09dda1c6121c76",
		ClientSecret: "6f3808641692c55c0c703b6ac478b4f18cc038394",
		RedirectURL:  "https://www.zj.com/return",
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
