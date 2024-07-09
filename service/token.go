package service

import (
	"blog_server/global"
	"blog_server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *models.UserModel) (string, error) { //创建token
	Jwtc := models.JWTClaims{
		ID: int(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: &jwt.NumericDate{
				Time: time.Unix(time.Now().Unix()-1, 0),
			}, //生效时间
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(time.Now().Unix()+int64(global.CONFIG.JWT.DeadTime), 0),
			}, //无效时间
			Issuer: "folk is fucking stand", //签发者
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Jwtc)
	return t.SignedString(global.CONFIG.JWT.SigningKey)
}
