package ad

import (
	"amazon-sdk/pkg"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Engine struct {
	*Auth
	HttpClient *http.Client
}

func (engine *Engine) CallApi(method, path string, body io.Reader) (*[]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			pkg.Logger.Error("CallApi错误:\n", err)
		}
	}()
	endpoint := engine.GetDataEndpoint()
	r, err := http.NewRequest(method, endpoint+path, body)
	if err != nil {
		pkg.Logger.Error("请求构造错误:\n", err)
		return nil, err
	}
	err = engine.SignRequest(r)
	if err != nil {
		pkg.Logger.Error("签名错误:\n", err)
		return nil, err
	}

	resp, err := engine.HttpClient.Do(r)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		pkg.Logger.Error("请求错误:\n", err)
		return nil, err
	}
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		pkg.Logger.Error("响应状态码错误:\n", err)
		return nil, errors.New(fmt.Sprintf("Fail to request. %s", string(respBody)))
	}
	return &respBody, nil
}

func (engine *Engine) SignRequest(r *http.Request) error {
	if time.Now().Add(time.Minute).After(engine.ExpiresAt) {
		err := engine.GetAccessTokenFromEndpoint()
		if err != nil {
			pkg.Logger.Error("获取access token错误:\n", err)
			return err
		}
	}
	r.Header.Add("Authorization", "Bearer "+engine.AccessToken)
	r.Header.Add("Amazon-Advertising-API-ClientId", engine.ClientID)
	if engine.ProfileId != "" {
		r.Header.Add("Amazon-Advertising-API-Scope", engine.ProfileId)
	}
	r.Header.Add("content-type", "application/json")
	r.Header.Add("accept", "application/json")
	return nil
}
