package sp

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
	//北美地区
	case "na":
		return "https://sellingpartnerapi-na.amazon.com"
	//欧洲地区
	case "eu":
		return "https://sellingpartnerapi-eu.amazon.com"
	//日本
	case "jp":
		return "https://sellingpartnerapi-fe.amazon.com"
	//澳大利亚
	case "au":
		return "https://sellingpartnerapi-fe.amazon.com"
	//新加坡
	case "sg":
		return "https://sellingpartnerapi-fe.amazon.com"
	default:
		panic("invalid region")
	}
}

func (a *Auth) GetLwaCodeEndpoint() string {
	switch a.Region {
	//北美地区
	case "na":
		return "https://sellercentral.amazon.com/apps/authorize/consent"
	//欧洲地区
	case "eu":
		return "https://sellercentral-europe.amazon.com/apps/authorize/consent"
	//日本
	case "jp":
		return "https://sellercentral-japan.amazon.com/apps/authorize/consent"
	//澳大利亚
	case "au":
		return "https://sellercentral.amazon.com.au/apps/authorize/consent"
	//新加坡
	case "sg":
		return "https://sellercentral.amazon.sg/apps/authorize/consent"
	default:
		panic("invalid region")
	}
}

func (a *Auth) GetLwaURL() (url string) {
	baseURL := a.GetLwaCodeEndpoint()
	url = fmt.Sprintf("%s?application_id=%s&redirect_uri=%s&state=%s", baseURL, a.AppID, a.RedirectURL, uuid.NewString())
	if a.Beta {
		url = fmt.Sprintf("%s&version=beta", url)
	}
	return url
}
