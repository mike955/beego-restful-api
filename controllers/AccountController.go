package controllers

import (
	"encoding/json"
	// "github.com/astaxie/beego"
	"beego-restful-api/models"
	// "fmt"
)

type AccountController struct {
	BaseController
}

func (c *AccountController) Index() {
	
}

func (c *AccountController) Login() {

}

func (this *AccountController) Register() {
	account := models.NewAccount()
	json.Unmarshal(this.Ctx.Input.RequestBody, &account)
	_ ,_ = models.AddAccount(account)
	this.JsonResopnse(0,"ok")
}

func (c *AccountController) Captcha() {

}