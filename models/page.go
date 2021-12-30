/*
 * @Description:
 * @Author: neozhang
 * @Date: 2021-12-30 21:56:24
 * @LastEditors: neozhang
 * @LastEditTime: 2021-12-30 23:03:37
 */
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init() {
	orm.RegisterDataBase("default", "mysql", db)
	orm.RegisterModel(new(Page))
}

func GetPage() Page {
	// rtn := Page{
	// 	Website: "hellobeego.com",
	// 	Email:   "model@beego.com",
	// }
	// return rtn

	o := orm.NewOrm()
	p := Page{
		Id: 1,
	}
	err := o.Read(&p)
	if err != nil {
		fmt.Println(err)
	}

	return p
}
