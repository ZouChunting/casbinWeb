package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

//用户表
type User struct {
	Id int
	Name string
	Password string
}

//资源表
type Source struct {
	Id int
	Name string
}

//动作表
type Action struct {
	Id int
	Name string
}

//角色表
type Role struct {
	Id int
	Name string
}

func init(){
	err:=orm.RegisterDataBase("default", "mysql", "root:root1234@tcp(127.0.0.1:3306)/casbin?charset=utf8", 30)
	if err!=nil{
		panic(err)
	}
	orm.RegisterModel(new(User),new(Source),new(Action),new(Role))
	err=orm.RunSyncdb("default", false, true)
	if err!=nil{
		panic(err)
	}
}