package response

type JwtNewTokenData struct {
	Token      string `json:"newtoken"`
	ExpireTime uint64 `json:"newExpireTime"`
}
type JwtNewTokenResponse struct {
	Data *JwtNewTokenData `json:"newTokenData"`
}
