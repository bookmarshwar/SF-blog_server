package routers

import (
	"blog_server/api"
)

func (r RouterGroup) UserRouter() {
	r.Register()
	r.Login()
}
func (r RouterGroup) Register() {
	r.POST("/register", api.ApiGroupAPP.USERAPI.Register)
}
func (r RouterGroup) Login() {
	r.POST("/login", api.ApiGroupAPP.USERAPI.Loggin)
}
