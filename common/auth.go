package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type App struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Beta         bool   //sp应用使用
	AppID        string //sp应用使用

}

type Seller struct {
	SellerName string
	SellerId   string
	Region     string
	ProfileId  string //广告对应一个店铺
}

type Auth struct {
	*App
	*Seller
	*Token
}

func (c *App) String() string {
	return fmt.Sprintf("AppID: %s\nClientID: %s\nClientSecret: %s\nRedirectURL: %s", c.AppID, c.ClientID, c.ClientSecret, c.RedirectURL)
}

func (t *Token) String() string {
	return fmt.Sprintf("AccessToken: %s\nRefreshToken: %s\nTokenType: %s\nExpiresIn: %d\nExpiresAt: %s\n", t.AccessToken, t.RefreshToken, t.TokenType, t.ExpiresIn, t.ExpiresAt.Format("2006-01-02 15:04:05"))
}

func (a *Auth) GetLwaTokenEndpoint() string {
	switch a.Region {
	//北美地区
	case "na":
		return "https://api.amazon.com/auth/o2/token"
	//欧洲地区
	case "eu":
		return "https://api.amazon.co.uk/auth/o2/token"
	//远东-日本
	case "jp":
		return "https://api.amazon.co.jp/auth/o2/token"
	//远东-澳大利亚
	case "au":
		return "https://api.amazon.co.jp/auth/o2/token"
	//远东-新加坡
	case "sg":
		return "https://api.amazon.co.jp/auth/o2/token"
	default:
		panic("invalid region")
	}
}

func (a *Auth) GetAccessTokenFromEndpoint() error {
	endpoint := a.GetLwaTokenEndpoint()
	reqBody, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": a.RefreshToken,
		"client_id":     a.ClientID,
		"client_secret": a.ClientSecret,
	})
	resp, err := http.Post(
		endpoint,
		"application/json",
		bytes.NewBuffer(reqBody))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	return a.ParseToken(resp)
}

func (a *Auth) ParseToken(resp *http.Response) error {
	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		errorStr := buf.String()
		return errors.New(fmt.Sprintf("Fail to generate token. %s", errorStr))
	}
	decoder := json.NewDecoder(resp.Body)
	token := new(Token)
	err := decoder.Decode(token)
	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	a.Token = token
	return err
}

func (a *Auth) GetRefreshToken(code string) error {
	baseURL := a.GetLwaTokenEndpoint()
	resp, err := http.PostForm(baseURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {a.RedirectURL},
		"client_id":     {a.ClientID},
		"client_secret": {a.ClientSecret}})
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	return a.ParseToken(resp)
}
