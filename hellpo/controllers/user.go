package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"hellpo/models"
	"math"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {

	//设置每页展示的条数
	pageSetNum,_ := c.GetInt("limit")
	pageSetNu,_ := c.GetInt("page")
		o:=orm.NewOrm()
	cnt, _ := o.QueryTable("user").Count()

	//	总页数
	pageCount := math.Ceil((float64(cnt) / float64(pageSetNum)))
	//	获取当前页码
	page, err := c.GetInt("pageNum")
	if err != nil {
		page = 1
	}

	var users []*models.User
	_,err = o.QueryTable("user").Limit(pageSetNum, pageSetNum*(pageSetNu-1)).All(&users)
	if err !=nil {
		c.ApiJson("未查到信息",200,"shibai")
	}
	var data  = make(map[string]interface{})
	data["list"] = users
	data["total"] = cnt
	data["pageCount"] = pageCount
	data["page"] = page
	c.ApiJson("ok",200,data)
}

func (c *UserController) Show() {
   id,_:= c.GetInt("id")
    o:=orm.NewOrm()
    user:= models.User{Id:id}
    err:=o.Read(&user)
	if err != nil {
		c.ApiJson("未查到信息",200,"shibai")
	}else {
		c.ApiJson("查到信息",200,user)
	}
}


func (c *UserController) UserUpdate() {
	id,_:= c.GetInt("id")
	username:= c.GetString("username")
	password:= c.GetString("password")
	o:=orm.NewOrm()
	user:=models.User{Id:id}
	if o.Read(&user) == nil{
		user.Username = username
		user.Password = password
		if num,err:= o.Update(&user); err == nil {
			c.ApiJson("查到信息",200,num)
		}
	}

}

func (c *UserController) Del() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	if num, err := o.Delete(&models.User{Id: id}); err == nil {
		c.ApiJson("success",200,num)
	}
}

func (c *UserController) Add() {

	username :=c.GetString("username")
	password := c.GetString("password")
	h := md5.New()
	h.Write([]byte(password)) // 需要加密的字符串为 123456
	password = hex.EncodeToString(h.Sum(nil))
	o := orm.NewOrm()
	var user models.User
	user.Username = username
	user.Password = password
	id,err:=o.Insert(&user)
	if err != nil {
		c.ApiJson("bad",200,"no")
	}else {
		c.ApiJson("success",200,id)
	}


}