/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 21:45:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 09:08:55
 */
package controllers

import (
	"fmt"
	"xcms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/go-swagger/go-swagger/examples/external-types/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	//session使用
	c.SetSession("username", "leixiaotian")
	user := c.GetSession("username")
	logs.Informational("user you loged in")
	fmt.Println(user)

	//更新邮箱
	models.UpdatePage()

	//查询数据并渲染
	m := models.GetPage()
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email
	c.TplName = "index.tpl"
}
