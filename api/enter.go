package api

import (
	"blog_server/api/document"
	markdownhtml "blog_server/api/markdown_html"
	"blog_server/api/setting"
	"blog_server/api/user"
)

type ApiGroup struct {
	SETTINGAPI      setting.SettingApi
	MARKDOWNHTMLAPI markdownhtml.MarkdownhtmlApi
	DOCUMENTAPI     document.DocumentApi
	USERAPI         user.UserApi
}

var ApiGroupAPP = new(ApiGroup)
