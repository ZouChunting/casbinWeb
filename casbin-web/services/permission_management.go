package services

import (
	"zct/casbin-web/utils"
)

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description  权限管理中间件
 */

type PermissionManager struct {
	casbinTool *utils.CasbinTool
}

func NewPermissionManager() *PermissionManager {
	return &PermissionManager{
		utils.NewCasbinTool(),
	}
}

//授权
func (pm PermissionManager) GrantPolicy(sub string,obj string,act string) bool {
	flag:=pm.casbinTool.Enforcer.AddPolicy(sub,obj,act)
	return flag
}

//为用户添加角色
func (pm PermissionManager) GrantRoleForUser(user string,role string) bool {
	flag:=pm.casbinTool.Enforcer.AddRoleForUser(user,role)
	return flag
}

//鉴权
func (pm PermissionManager) Authentication(sub string,obj string,act string) bool {
	flag:=pm.casbinTool.Enforcer.Enforce(sub,obj,act)
	return flag
}

