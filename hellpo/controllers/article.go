package controllers

import (
	"github.com/astaxie/beego/orm"
	"hellpo/models"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) Get() {
	//设置每页展示的条数
	pageSetNum,_ := c.GetInt("limit")
	pageSetNu,_ := c.GetInt("page")
	o:=orm.NewOrm()
	cnt, _ := o.QueryTable("article").Count()

	var articles []*models.Article
	_,err := o.QueryTable("article").Limit(pageSetNum, pageSetNum*(pageSetNu - 1)).All(&articles)
	if err !=nil {
		c.ApiJson("未查到信息",200,"shibai")
	}
	var data  = make(map[string]interface{})
	data["list"] = articles
	data["total"] = cnt

	c.ApiJson("ok",200,data)


}


func (c *ArticleController) Add() {
	id,_:= c.GetInt("id")
	title := c.GetString("title")
	author:=c.GetString("author")
	star,_:=c.GetInt("star")
	content:=c.GetString("content")
	comment:=c.GetString("comment")
	img := c.GetString("image_uri")
	o:= orm.NewOrm()
	if id !=0 {
		articles:=models.Article{Id:id}
	     o.Read(&articles)
			articles.ImgUrl =img
			articles.Title = title
			articles.Content = content
			articles.Comment = comment
			if _,err := o.Update(&articles);err ==nil {
				c.ApiJson("更新成功",200,"qqq")
			}else {
				c.ApiJson("更新失败",200,"失败")
			}

	}else {

		var article models.Article
		article.Comment = comment
		article.Content = content
		//article.Create = time1
		article.Star = star

		article.Title = title
		article.Author = author
		article.ImgUrl = img

		_,err := o.Insert(&article)
		if err !=nil {
			c.ApiJson("新增失败",200,"失败")
		}else {
			c.ApiJson("新增成功",200,"ok")
		}
	}


}


func (c *ArticleController) Update() {
   id,_ := c.GetInt("id")
   o:= orm.NewOrm()
   article:=models.Article{Id:id}
	if err := o.Read(&article);err !=nil {
		c.ApiJson("未查询到信息",200,"失败")
	}else {
		c.ApiJson("查到文章信息",200,article)
	}
	//if err !=nil {
	//	c.ApiJson("新增失败",200,"失败")
	//}else {
	//	c.ApiJson("新增成功",200,article)
	//}

}