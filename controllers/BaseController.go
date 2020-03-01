package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}


func (this *BaseController) JsonResopnse(errno int, msg string, data ...interface{}){
	response := make(map[string]interface{}, 3)
	
	response["errno"] = errno
	response["msg"] = msg

	if len(data) > 0 && data[0] != nil {
		response["data"] = data[0]
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	this.Data["json"] = response;
	this.ServeJSON()
}