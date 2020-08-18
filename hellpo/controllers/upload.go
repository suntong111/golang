package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hellpo/models"
	"path"
	"time"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Upload() {
	f, h,_ := c.GetFile("file")
	defer f.Close()
	//1.限定格式
	//拿到文件的后缀名
	fileExt := path.Ext(h.Filename)
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg"{
		beego.Info("上传文件格式错误")
		return
	}
	//2.限制大小
	if h.Size > 500000 {
		beego.Info("上传文件过大")
		return
	}
	beego.Info(fileExt)
	//3.对文件重命名避免重复+时间戳 2006-01-02 15:04:05 go语言诞生时间 可正常格式化时间
	//随机数也可
	fileName := time.Now().Format("2006-01-02 15-04-05")

	//     “.”必须
	FN :=c.SaveToFile("file", "./static/img/"+fileName+h.Filename)
	beego.Info("FN",FN)

	o:=orm.NewOrm()
	var uploads models.Upload

	uploads.File = "static/img/"+fileName+h.Filename

	if _,err:=o.Insert(&uploads);err ==nil{
		c.ApiJson("ok",200,uploads)
	}else {
		c.ApiJson("新增失败",200,"失败")
	}




}