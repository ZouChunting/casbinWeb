package dao

import (
	"github.com/astaxie/beego/orm"
	"log"
	"zct/casbin-web/models"
)

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

type Dao struct {
	o orm.Ormer
}

func NewDao() *Dao {
	o:=orm.NewOrm()
	return &Dao{
		o,
	}
}

//user表
func (dao Dao) ReadUser(user models.User) error {
	err:=dao.o.Read(&user,"Name","Password")
	return err
}

func (dao Dao) ReadUserName(user models.User) error {
	err:=dao.o.Read(&user,"Name")
	return err
}

func (dao Dao) ReadOrCreateUser(user models.User) (bool,int64,error) {
	created, id, err:=dao.o.ReadOrCreate(&user,"Name")
	return created, id, err
}

func (dao Dao) AddUser(user models.User) {
	_,err:=dao.o.Insert(&user)
	if err!=nil{
		log.Print(err)
	}
}

//role表
func (dao Dao) ReadRole(role models.Role) error {
	err:=dao.o.Read(&role,"Name")
	return err
}

//source表
func (dao Dao) ReadSource(source models.Source) error {
	err:=dao.o.Read(&source,"Name")
	return err
}

//action表
func (dao Dao) ReadAction(action models.Action) error {
	err:=dao.o.Read(&action,"Name")
	return err
}