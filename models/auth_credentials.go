package models

type AuthCredentials struct {
	AccessToken     string
	RefreshToken    string
	AccessTokenExp  int64
	RefreshTokenExp int64
}
