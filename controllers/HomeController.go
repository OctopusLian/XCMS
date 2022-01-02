/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-02 09:06:33
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-02 09:06:34
 */
package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"

	//查询数据并渲染
	//c.TplName = "user/index.html"
	c.setTpl("home/index.html")
}
