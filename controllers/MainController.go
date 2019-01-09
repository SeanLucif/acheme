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
	c.TplName = "index.html"
}

func (c *MainController)Login()  {
	if c.GetString("school") == "" {
		c.Ctx.WriteString("请输入学号")
		return
	}else if c.GetString("school") == "" {
		c.Ctx.WriteString("请输入学号")
		return
	} else if c.GetString("school") == "" {
		c.Ctx.WriteString("请输入学号")
		return
	}
	c.Ctx.WriteString("登录成功")
	return
}
