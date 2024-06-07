package api

import (
	markdownhtml "blog_server/api/markdown_html"
	"blog_server/api/setting"
)

type ApiGroup struct {
	SETTINGAPI      setting.SettingApi
	MARKDOWNHTMLAPI markdownhtml.MarkdownhtmlApi
}

var ApiGroupAPP = new(ApiGroup)
