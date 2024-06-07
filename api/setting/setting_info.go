package setting

import (
	"blog_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfo(c *gin.Context) {
	res.SucceesMessage("成功", c)
}
