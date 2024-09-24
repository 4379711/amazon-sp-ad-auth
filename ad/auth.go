package ad

import (
	"amazon-sdk/common"
	"fmt"
	"github.com/google/uuid"
)

type Auth struct {
	*common.Auth
}

func (a *Auth) GetDataEndpoint() string {
	switch a.Region {
	//欧洲地区
	case "eu":
		return "https://advertising-api-eu.amazon.com"
	//日本
	case "jp":
		return "https://advertising-api-fe.amazon.com"
	//澳大利亚
	case "au":
		return "https://advertising-api-fe.amazon.com"
	//新加坡
	case "sg":
		return "https://advertising-api-fe.amazon.com"
	//北美地区
	case "na":
		return "https://advertising-api.amazon.com"
	default:
		panic("invalid region")
	}
}

func (a *Auth) GetLwaCodeEndpoint() string {
	switch a.Region {
	//欧洲地区
	case "eu":
		return "https://eu.account.amazon.com/ap/oa"
	//日本
	case "jp":
		return "https://apac.account.amazon.com/ap/oa"
	//澳大利亚
	case "au":
		return "https://apac.account.amazon.com/ap/oa"
	//新加坡
	case "sg":
		return "https://apac.account.amazon.com/ap/oa"
	//北美地区
	case "na":
		return "https://www.amazon.com/ap/oa"
	default:
		panic("invalid region")
	}
}

func (a *Auth) GetLwaURL() (url string) {
	baseURL := a.GetLwaCodeEndpoint()
	url = fmt.Sprintf("%s?client_id=%s&scope=advertising::campaign_management&response_type=code&redirect_uri=%s&state=%s", baseURL, a.ClientID, a.RedirectURL, uuid.NewString())
	return url
}
