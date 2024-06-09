package api

import (
	"blog_server/api/document"
	markdownhtml "blog_server/api/markdown_html"
	"blog_server/api/setting"
)

type ApiGroup struct {
	SETTINGAPI      setting.SettingApi
	MARKDOWNHTMLAPI markdownhtml.MarkdownhtmlApi
	DOCUMENTAPI     document.DocumentApi
}

var ApiGroupAPP = new(ApiGroup)
