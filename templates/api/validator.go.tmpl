package api

import (
	"fmt"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

const KeyID = "app-key-id"
const Issuer = "app-issuer"
const Audience = "app-audience"
const PermissionsClaim = "perm"
const UserInfoClaim = "user_info"

type Authenticator struct {
	PrivateKey string
}

var _ JWSValidator = (*Authenticator)(nil)

func NewAppAuthenticator(privateKey string) *Authenticator {
	return &Authenticator{PrivateKey: privateKey}
}

func (f *Authenticator) ValidateJWS(jwsString string) (jwt.Token, error) {
	return jwt.Parse(
		[]byte(jwsString),
		jwt.WithVerify(
			jwa.HS256,
			[]byte(f.PrivateKey),
		),
		jwt.WithAudience(Audience),
		jwt.WithIssuer(Issuer),
	)
}

func (f *Authenticator) SignToken(t jwt.Token) ([]byte, error) {
	hdr := jws.NewHeaders()
	return jwt.Sign(t, jwa.HS256, []byte(f.PrivateKey), jwt.WithHeaders(hdr))
}

func (f *Authenticator) GenerateJWS(info interface{}, permissions interface{}) (string, error) {
	t := jwt.New()
	err := t.Set(jwt.IssuerKey, Issuer)
	if err != nil {
		return "", fmt.Errorf("setting issuer: %w", err)
	}
	err = t.Set(jwt.AudienceKey, Audience)
	if err != nil {
		return "", fmt.Errorf("setting audience: %w", err)
	}
	err = t.Set(UserInfoClaim, info)
	if err != nil {
		return "", fmt.Errorf("setting info: %w", err)
	}

	if permissions == nil {
		permissions = []string{}
	}
	err = t.Set(PermissionsClaim, permissions)
	if err != nil {
		return "", fmt.Errorf("setting permissions: %w", err)
	}

	byteResult, err := f.SignToken(t)

	return string(byteResult), err
}
