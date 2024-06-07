package routers

import (
	"blog_server/api"
)

func (r RouterGroup) MarkdownhtmRouter() {
	r.GET("/md_html", api.ApiGroupAPP.MARKDOWNHTMLAPI.MarkdownHtmlInfo)
}
