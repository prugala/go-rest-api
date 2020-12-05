package model

type JwtTokenResponse struct {
	Code   string `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}
