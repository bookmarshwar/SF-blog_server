package markdownhtml

import (
	"blog_server/models/res"
	"blog_server/service"

	"github.com/gin-gonic/gin"
)

var mds = `# header

Sample text.

[link](http://example.com)
`

func (MarkdownhtmlApi) MarkdownHtmlInfo(c *gin.Context) {
	md := service.MarkdownToHtml([]byte(mds))
	res.Succees(map[string]string{
		"html": string(md),
	}, "成功", c)
}
