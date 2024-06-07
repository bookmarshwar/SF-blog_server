package routers

import (
	"blog_server/api"
)

func (r RouterGroup) MarkdownhtmRouter() {
	r.POST("/md_html", api.ApiGroupAPP.MARKDOWNHTMLAPI.MarkdownHtmlInfo)
}
