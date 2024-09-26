package ad

import (
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

func (engine *Engine) SignRequest(method, path string, body io.Reader) (*[]byte, error) {
	endpoint := engine.GetDataEndpoint()
	r, err := http.NewRequest(method, endpoint+path, body)
	if err != nil {
		return nil, err
	}
	if time.Now().Add(time.Minute).After(engine.ExpiresAt) {
		err := engine.GetAccessTokenFromEndpoint()
		if err != nil {
			return nil, err
		}
	}
	r.Header.Add("Authorization", "Bearer "+engine.AccessToken)
	r.Header.Add("Amazon-Advertising-API-ClientId", engine.ClientID)
	if engine.ProfileId != "" {
		r.Header.Add("Amazon-Advertising-API-Scope", engine.ProfileId)
	}
	r.Header.Add("content-type", "application/json")
	r.Header.Add("accept", "application/json")
	resp, err := engine.HttpClient.Do(r)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Fail to request. %s", string(respBody)))
	}
	return &respBody, nil

}
