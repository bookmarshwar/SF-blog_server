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
	return t.SignedString([]byte(global.CONFIG.JWT.SigningKey))
}
func ParseToken(token string) (*models.JWTClaims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*models.JWTClaims); ok && tokenClaims.Valid { // Valid()验证基于时间的声明
			return claims, nil
		}
	}
	return nil, err
}
