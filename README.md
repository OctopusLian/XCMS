# XCMS内容管理平台  

用Beego框架及其ORM模块，开发一个内容管理平台。项目会使用MVC架构，完成常见的登录、权限管理、菜单管理等模块。其中，内容管理模块采用通用化设计，在无需二次开发的前提下，可实现各种内容管理模块的添加，做到灵活高效。课程地址https://www.imooc.com/learn/1162  

## 涉及技术栈  

- Go  
- [Beego框架](https://beego.vip/)  
- MySQL数据库和ORM  
- 前端

## 运行  

首先默认您已经安装配置好了Go语言相关环境  

克隆项目仓库  
git clone git@github.com:OctopusLian/xcms.git  

执行go mod tidy下载相关包  
go mod tidy  

在mysql数据库中创建`defatlt`和`godb`数据库  

将static/sql/godb.sql放在mysql的godb数据库中执行  

使用bee命令运行  
bee run

打开本地端口号http://localhost:8080  

![](./res/login.png)  

用户名：aa  
密码：123456  

验证成功后即可跳转  
![](./res/ex1.png)  

## 如果您想使用Debug方式来学习的话  

### VSCode下配置  

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "xcms",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/main.go",
        }
    ]
}
```


## 新增模块  

在此基础上，新增[gin + casbin权限管理验证程序](/cmd/casbint/)  