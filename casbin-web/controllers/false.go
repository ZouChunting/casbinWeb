package controllers

import "github.com/astaxie/beego"

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

type FalseController struct {
	beego.Controller
}

func (c *FalseController) Get()  {
	c.Data["des"] = c.GetString("des")
	c.TplName="false.html"
}