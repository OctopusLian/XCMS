/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-03 11:32:21
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-04 16:59:38
 */
package models

import (
	"runtime"

	"github.com/astaxie/beego/orm"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
)

type CasbinRule struct {
	Id    int    // 自增主键
	PType string // Policy Type - 用于区分 policy和 group(role)
	V0    string // subject
	V1    string // object
	V2    string // action
	V3    string // 这个和下面的字段无用，仅预留位置，如果你的不是
	V4    string // sub, obj, act的话才会用到
	V5    string // 如 sub, obj, act, suf就会用到 V3
}

func init() {
	//TODO
	//orm.RegisterModel(new(CasbinRule))
	// 实际上同步数据库在整个Beego项目中只需要执行一次，如果
	// 您在别的地方已经同步数据库，这里就不用在执行一次 RunSyncdb
	//_ = orm.RunSyncdb("default", false, false)

	//orm.RegisterDataBase("default", "mysql", "root:mysql123@tcp(127.0.0.1:3306)/godb?charset=utf8", 30)

	// 初始化 Casbin
	// 初始化的目的就是获取一个可用的Enforce，这是Casbin访问控制框架的核心，无论什么操作都离不开它。
	//RegisterCasbin()
}

// 注意，这个Enforcer很重要，Casbin使用都是调用这个变量
var Enforcer *casbin.Enforcer

type Adapter struct {
	o orm.Ormer
}

func RegisterCasbin() {
	a := &Adapter{}
	a.o = orm.NewOrm()
	// 这个我不知道干嘛的
	runtime.SetFinalizer(a, finalizer)
	// Enforcer初始化 - 即传入Adapter对象
	Enforcer = casbin.NewEnforcer("conf/casbin.conf", a)
	// Enforcer读取Policy
	err := Enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
}

// finalizer is the destructor for Adapter.
// 这个函数里面啥都没有，就是这样
func finalizer(_ *Adapter) {}

// 注意，方法对应的具体代码要从beego-ORM-Adapter/adapter.go中复制过来
// 这里方法里面使用的orm操作还是要根据自己的
// 实际情况作出调整，不要盲目复制
// func loadPolicyLine(line CasbinRule, model model.Model)     {}
func savePolicyLine(ptype string, rule []string) CasbinRule {
	return CasbinRule{
		PType: ptype,
	}
}

// func (a *Adapter) LoadPolicy(model model.Model) error                         {}
// func (a *Adapter) SavePolicy(model model.Model) error                         {}
// func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error    {}
// func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {}
// func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
// }

func (a *Adapter) SavePolicy(model model.Model) error {
	var lines []CasbinRule
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}
	_, err := a.o.InsertMulti(len(lines), lines)
	return err
}
