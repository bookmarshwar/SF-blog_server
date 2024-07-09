package routers

import (
	"blog_server/api"
)

func (r RouterGroup) UserRouter() {
	r.Register()
}
func (r RouterGroup) Register() {
	r.POST("/register", api.ApiGroupAPP.USERAPI.Register)
}
