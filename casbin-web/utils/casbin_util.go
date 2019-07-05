package utils

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

import (
	"github.com/astaxie/beego"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

const CASBINMODEL = ` 
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub)  && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`

type CasbinTool struct {
	Enforcer *casbin.Enforcer
}

func NewCasbinTool() *CasbinTool {
	/**
	 *@描述  casbin工具初始化
	 *@参数
	 *@返回值  &CasbinTool{}
	 *@创建人  zct
	 *@创建时间  2019-07-04
	 *@修改人和其它信息
	 */
	casbinEnforcer,err:=casbin.NewEnforcerSafe(casbin.NewModel(CASBINMODEL),gormadapter.NewAdapter(beego.AppConfig.String("driverName"),beego.AppConfig.String("dataSourceName")))
	if err!=nil{
		panic(err)
	}
	err=casbinEnforcer.LoadPolicy()
	if err!=nil{
		panic(err)
	}
	return &CasbinTool{
		casbinEnforcer,
	}
}
