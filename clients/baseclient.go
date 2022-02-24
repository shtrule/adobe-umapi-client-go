package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	once       sync.Once
	httpClient *http.Client
)

var GetHttpClient = func() *http.Client {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout: time.Minute * 15,
		}
	})
	return httpClient
}

func GetRequest(url, token, clientId string) ([]byte, error) {
	authorizationHeader := fmt.Sprintf("Bearer %v", token)
	rqs, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create new get request, url: %s, error: %s", url, err.Error())
	}

	rqs.Header.Set("Authorization", authorizationHeader)
	rqs.Header.Set("X-Api-Key", clientId)

	rsp, err := GetHttpClient().Do(rqs)

	if err != nil {
		return nil, fmt.Errorf("get request, url: %s, error: %s", url, err.Error())
	}

	responseBody, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body, error: %s", err.Error())
	}

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("creative cloud user management api request was not successful. status code: %v, response: %v", rsp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func GetTokenRequest(exchangeJwtUrl string, bodyData url.Values) ([]byte, error) {
	rsp, err := GetHttpClient().PostForm(exchangeJwtUrl, bodyData)

	if err != nil {
		return nil, fmt.Errorf("failed to get access token. Error: %s", err.Error())
	}

	responseBody, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read the response body. request URL:%s  Error:%s", exchangeJwtUrl, err.Error())
	}

	// https://www.adobe.io/authentication/auth-methods.html#!AdobeDocs/adobeio-auth/master/JWT/JWT.md
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("creative cloud get token request was not successful. status code: %v, response: %v", rsp.StatusCode, string(responseBody))
	}
	return responseBody, nil
}
