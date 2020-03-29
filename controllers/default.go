package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Test"] = "你好"
 	c.TplName = "test.html"

}

type IndexController struct {
	beego.Controller
}

func (c*IndexController)ShowGet()  {
	c.Data["data"] = "王者荣耀"
	c.TplName = "test.html"
}

func (c*IndexController)HandlePost()  {
	c.Data["data"] = "王者荣耀Post"
	c.TplName = "test.html"
}

