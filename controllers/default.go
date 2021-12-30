/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 21:45:08
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 22:28:31
 */
package controllers

import (
	"xcms/models"

	"github.com/astaxie/beego"
	//"github.com/go-swagger/go-swagger/examples/external-types/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var pg models.Page
	pg = models.GetPage()
	c.Data["Website"] = pg.Website
	c.Data["Email"] = pg.Email
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl" //模板文件名
}
