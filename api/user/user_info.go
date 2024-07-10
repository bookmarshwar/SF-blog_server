package user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (UserApi) Loggin(c *gin.Context) {
	var user models.UserModel
	c.BindJSON(&user)
	db := global.DB
	var TrueUser models.UserModel
	db.Where("user_name =?", user.UserName).First(&TrueUser)
	if TrueUser.ID == 0 {
		res.FailMessage("用户不存在", c)
		return
	}
	if user.Password == TrueUser.Password {
		token, error := service.CreateToken(&TrueUser)
		if error != nil {
			res.FailMessage(fmt.Sprintln("蜜汁错误%s", error.Error()), c)
			return
		}
		res.SucceesWithData(map[string]string{
			"token": token,
		}, c)
		return
	}
	res.FailMessage("密码错误", c)
}
func (UserApi) Register(c *gin.Context) {
	db := global.DB
	var user models.UserModel
	c.BindJSON(&user)
	if !checkUser(user) {
		res.FailMessage("格式错误", c)
	} else {
		db.Create(&user)
		res.SucceesMessage("注册成功", c)
	}
}
func (UserApi) GetUser(c *gin.Context) {

}
func checkUser(user models.UserModel) bool {
	return true
}
