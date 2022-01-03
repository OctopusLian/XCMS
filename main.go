/*
 * @Description:项目的总入口
 * @Author: neozhang
 * @Date: 2021-12-30 21:45:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 11:56:03
 */
package main

import (
	_ "xcms/routers"

	"github.com/astaxie/beego"
)

func main() {
	//beego.InsertFilter("/v1/user/*", beego.BeforeRouter, filters.NewAuthorizer(models.Enforcer))
	beego.Run()
}
