package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type Jwt struct {
	SigningKey []byte
}

func NewJwt() *Jwt {
	return &Jwt{
		[]byte("31231232"),
	}
}

// 生成
func (j *Jwt) GenerateToken(id string) (string, error) {
	// 设置有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claim := Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 时间戳
			Issuer:    "gin-blog",
		},
	}

	tokenClamis := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenClamis.SignedString(j.SigningKey)

	return token, err
}

// 解析
func (j *Jwt) ParseToken(token string) (*Claim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claim); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	//
	return nil, err
}
