package clients

import (
	"crypto/rsa"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtBuilderClient interface {
	CreateJwtToken() (string, error)
}

type JwtBuilder struct {
	ClientId          string
	OrganizationId    string
	TechnicalAccount  string
	CertificateSecret []byte
}

func NewJwtBuilder(clientId, organizationId, technicalAccount string, certificateSecret []byte) JwtBuilderClient {
	return &JwtBuilder{ClientId: clientId, OrganizationId: organizationId, TechnicalAccount: technicalAccount, CertificateSecret: certificateSecret}
}

func (builder JwtBuilder) CreateJwtToken() (string, error) {
	tokenExpiryUnixTime := time.Now().Add(time.Minute * 10).Unix()
	audience := fmt.Sprintf("https://ims-na1.adobelogin.com/c/%s", builder.ClientId)

	adobeClaims := jwt.MapClaims{}
	adobeClaims["exp"] = tokenExpiryUnixTime
	adobeClaims["iss"] = builder.OrganizationId
	adobeClaims["sub"] = builder.TechnicalAccount
	adobeClaims["aud"] = audience
	adobeClaims["https://ims-na1.adobelogin.com/s/ent_user_sdk"] = true

	jwtToken := NewWithClaims(adobeClaims)

	privatePemKey, err := ParseRSAPrivateKeyFromPEM(builder.CertificateSecret)

	if err != nil {
		return "", fmt.Errorf("failed to convert certificate secret to PEM key: %s", err.Error())
	}

	signedJwtToken, err := GetSignedString(jwtToken, privatePemKey)

	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token: %s", err.Error())
	}

	return signedJwtToken, nil
}

var NewWithClaims = func(claims jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
}

var ParseRSAPrivateKeyFromPEM = func(certificateSecret []byte) (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM(certificateSecret)
}

var GetSignedString = func(jwtToken *jwt.Token, privatePemKey *rsa.PrivateKey) (string, error) {
	signedJwtToken, err := jwtToken.SignedString(privatePemKey)
	return signedJwtToken, err
}
