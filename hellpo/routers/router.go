package routers

import (
	"hellpo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{},"get:Get;post:Post")
    beego.Router("user/update",&controllers.UserController{},"get:Show;post:UserUpdate")
    beego.Router("user/delete",&controllers.UserController{},"get:Del")
    beego.Router("user/add",&controllers.UserController{},"post:Add")

    //文章
	beego.Router("/article", &controllers.ArticleController{},"get:Get;post:Post")
	beego.Router("/article/add", &controllers.ArticleController{},"post:Add")
	beego.Router("/article/update", &controllers.ArticleController{},"get:Update")


    //登陆接口
	beego.Router("/user/login", &controllers.LoginController{},"post:Login")

    //上传文件
	beego.Router("/upload", &controllers.UploadController{},"post:Upload")

    //权限
	beego.Router("/rule", &controllers.RuleController{},"get:Get;post:Post")
}
