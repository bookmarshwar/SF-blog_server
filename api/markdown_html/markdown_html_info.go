package markdownhtml

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"

	"github.com/gin-gonic/gin"
)

var mds = `# header

Sample text.

[link](http://example.com)
`

type MarkdownReq struct {
	Md string `json:"md"`
}

func (MarkdownhtmlApi) MarkdownHtmlInfo(c *gin.Context) {
	var markdownReq MarkdownReq
	err := c.ShouldBindJSON(&markdownReq)
	if err != nil {
		global.LOG.Error(err)
		res.FailCode(1101, c)
		return
	}
	html := service.MarkdownToHtml([]byte(markdownReq.Md))
	res.Succees(map[string]string{
		"html": string(html),
	}, "成功", c)
}
