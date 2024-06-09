package document

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (DocumentApi) Document(c *gin.Context) {
	fmt.Println("sb")
	global.DB.Create(&models.DocumentModel{
		Briefly: "后端是这样的,前端只需要乖乖吃后端写的屎就行了....",
		Title:   "论前后端关系",
		Cover:   "assets/image/2abffe2e804bac277d1c8c002f33f2ec.jpg",
		Path:    "0",
	})
	fmt.Println("sb2")
	res.SucceesMessage("成功", c)
}
func (DocumentApi) GetList(c *gin.Context) {
	var documentList []models.DocumentModel
	result := global.DB.Order("created_at DESC").Limit(10).Find(&documentList)
	if result.Error != nil {
		// 处理查询错误，比如数据库连接问题或查询失败
		global.LOG.Error("Error retrieving top 10 articles: %v", result.Error)
	}
	res.SucceesWithData(documentList, c)
}
func (DocumentApi) PostInfo(c *gin.Context) {

}
