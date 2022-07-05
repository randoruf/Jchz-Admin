package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"jchz-admin/global"
	"time"
)

// Claims 自定义载荷
type Claims struct {
	UID                  string
	BufferTime           int64
	jwt.RegisteredClaims //jwt 自带载荷
}

type JWT struct {
	SecretKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.JA_CONFIG.TokenConfig.AccessSecret),
	}
}

// BuildClaims 构建 payload
func (J *JWT) BuildClaims(uid string, ttl uint64) Claims {
	return Claims{
		UID:        uid,
		BufferTime: global.JA_CONFIG.TokenConfig.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                     //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                     //生效时间
		}}
}

// CreateToken 生成token
func (J *JWT) CreateToken(userID string) (string, error) {
	claims := J.BuildClaims(userID, global.JA_CONFIG.TokenConfig.AccessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(J.SecretKey)
	return tokenString, err
}

// secret 用于获取密钥
func (J *JWT) secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return J.SecretKey, nil
	}
}

// ParseToken 解析 token，返回 payload 部分
func (J *JWT) ParseToken(tokenss string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenss, &Claims{}, J.secret())

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims *Claims) (string, error) {
	v, err, _ := global.JA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims.UID)
	})
	return v.(string), err
}
