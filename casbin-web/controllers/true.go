package controllers

import "github.com/astaxie/beego"

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

type TrueController struct {
	beego.Controller
}

func (c *TrueController) Get()  {
	c.TplName="true.html"
}