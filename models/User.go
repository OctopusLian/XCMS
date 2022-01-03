/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 11:43:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-03 11:43:47
 */
package models

type User struct {
	// 用户模型
	Id       int    `orm:"auto;pk" description:"用户序号" json:"uid"`
	Username string `orm:"unique" description:"用户名" json:"username"`
	Password string `description:"用户密码" json:"password"`
	Role     *Role  `orm:"rel(fk);null" description:"角色" json:"Role"`
}

// TODO:各种ORM查询方法的实现
