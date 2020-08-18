package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type JsonReturn struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

func (c *BaseController) ApiJson(msg string,code int,data interface{}) {
     var JsonReturn JsonReturn
     JsonReturn.Msg = msg
     JsonReturn.Code = code
     JsonReturn.Data = data
     c.Data["json"] = JsonReturn
     c.ServeJSON()
     c.StopRun()
}