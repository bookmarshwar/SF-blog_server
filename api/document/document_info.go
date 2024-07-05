package document

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service"
	"fmt"
	"net/http"

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
func (DocumentApi) UpDocument(c *gin.Context) {
	file, err := c.FormFile("markdown")
	if err != nil && err != http.ErrMissingFile {
		res.FailMessage(fmt.Sprintf("上传文件 'file' 时出错：%s", err.Error()), c)
		return
	}
	filePath, err := service.Savefile(file, "./res/md_file/")
	if err != nil {
		res.FailMessage(fmt.Sprintf("上传文件 'file' 时出错：%s", err.Error()), c)
	}
	cover, err := c.FormFile("cover")
	if err != nil {
		res.FailMessage(fmt.Sprintf("上传文件 'file' 时出错：%s", err.Error()), c)
	}
	coverPath, err := service.Savefile(cover, "./res/cover_file/")
	name := c.PostForm("name")
	if name == "" {
		name = "未命名文章"
	}
	briefly := c.PostForm("briefly")
	global.DB.Create(&models.DocumentModel{
		Briefly: briefly,
		Title:   name,
		Cover:   coverPath[1:],
		Path:    filePath[1:],
	})
	res.SucceesMessage("上传成功", c)

}
