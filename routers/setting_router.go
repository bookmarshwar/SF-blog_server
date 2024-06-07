package routers

import (
	"blog_server/api"
)

func (r RouterGroup) SettingRouter() {
	r.GET("", api.ApiGroupAPP.SETTINGAPI.SettingInfo)
}
