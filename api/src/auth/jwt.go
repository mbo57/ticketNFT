package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// 署名用秘密鍵（仮）
var signingKey = []byte("secretkey")

var JwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwt.RegisteredClaims)
	},
	SigningKey: signingKey,
}

func CreateJwtToken(id string) string {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //tokenの有効期限
		ID:        id,                                                 // ユーザー識別子
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(signingKey)

	if err != nil {
		return ""
	}
	return token
}

// JWT TokenからClaimsを取り出す関数
func GetClaims(c echo.Context) *jwt.RegisteredClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt.RegisteredClaims)
	return claims
}
