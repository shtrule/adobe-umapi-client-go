package clients

import (
	"encoding/json"
	"fmt"
	"net/url"

	cc "snowsoftware.com/saas/adobe-creative-cloud-data-collector/models"
)

type AdobeAuthClient interface {
	GetToken() (cc.Token, error)
}

type AuthClient struct {
	ClientId     string
	ClientSecret string
	JwtToken     string
}

func NewAuthClient(clientId, clientSecret, jwtToken string) *AuthClient {
	return &AuthClient{ClientId: clientId, ClientSecret: clientSecret, JwtToken: jwtToken}
}

func (client AuthClient) GetToken() (cc.Token, error) {
	bodyData := url.Values{}

	bodyData.Add("client_id", client.ClientId)
	bodyData.Add("client_secret", client.ClientSecret)
	bodyData.Add("jwt_token", client.JwtToken)

	exchangeJwtUrl := "https://ims-na1.adobelogin.com/ims/exchange/v1/jwt/"

	responseBody, err := GetTokenRequest(exchangeJwtUrl, bodyData)
	if err != nil {
		return cc.Token{}, err
	}
	var token cc.Token
	err = json.Unmarshal(responseBody, &token)

	if err != nil {
		return cc.Token{}, fmt.Errorf("failed to unmarshal response body to token: %s", err.Error())
	}

	return token, nil
}
