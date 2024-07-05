package routers

import "blog_server/api"

func (r RouterGroup) DocumentRouter() {
	r.Document()
	r.DocumentGetInfo()
	r.DocumentUP()
}
func (r RouterGroup) Document() {
	r.GET("/document", api.ApiGroupAPP.DOCUMENTAPI.Document)
}

func (r RouterGroup) DocumentGetInfo() {
	r.GET("/getdocumentlist", api.ApiGroupAPP.DOCUMENTAPI.GetList)
}
func (r RouterGroup) DocumentUP() {

	r.POST("/updoucument", api.ApiGroupAPP.DOCUMENTAPI.UpDocument)

}
