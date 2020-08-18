package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"hellpo/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
    username := c.GetString("username")
    password := c.GetString("password")

	h := md5.New()
	h.Write([]byte(password)) // 需要加密的字符串为 123456
	 password = hex.EncodeToString(h.Sum(nil))


	o:= orm.NewOrm()
	 user := models.User{Username: username,Password: password}
	 //user.Username = username
	 //user.Password = password
	if err:= o.Read(&user,"Username","Password");err==nil {
		c.ApiJson("查到信息",200,user)
	}else {
		c.ApiJson("bad",200,"no")
	}

}