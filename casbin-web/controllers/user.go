package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zct/casbin-web/dao"
	"zct/casbin-web/models"
	"zct/casbin-web/utils"
)

/**
 *@Author  zct
 *@Date  2019-07-04
 *@Description
 */

//注册
type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get()  {
	c.TplName="register.html"
}

func (c *RegisterController) Post()  {
	userName:=c.GetString("user_name")
	userPassword:=c.GetString("user_password")
	user:=models.User{
		Name:userName,
		Password:userPassword,
	}
	created, _, err:=dao.NewDao().ReadOrCreateUser(user)
	if err!=nil{
		c.Ctx.Redirect(302, "/false?des=unknown-error")
	}
	if created{
		c.Ctx.Redirect(302, "/true")
	}else {
		c.Ctx.Redirect(302, "/false?des=already-exits")
	}
}

//登录
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get()  {
	c.TplName="login.html"
}

func (c *LoginController) Post()  {
	//使用session
	w:=c.Ctx.ResponseWriter
	r:=c.Ctx.Request
	sess, _ := utils.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	userName:=c.GetString("user_name")
	userPassword:=c.GetString("user_password")

	user:=models.User{
		Name:userName,
		Password:userPassword,
	}

	err:=dao.NewDao().ReadUser(user)
	if err == orm.ErrNoRows {
		c.Ctx.Redirect(302, "/false?des=wrong-user-info")
	} else if err == orm.ErrMissPK {
		c.Ctx.Redirect(302, "/false?des=can-not-find-primary-key")
	} else {
		err=sess.Set("username", userName)
		c.Ctx.Redirect(302, "/true")
	}
}

//注销
type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get()  {
	w:=c.Ctx.ResponseWriter
	r:=c.Ctx.Request
	sess, _ := utils.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	fmt.Println("before logout:",sess.Get("username"))
	_=sess.Delete("username")
	fmt.Println("after logout:",sess.Get("username"))
	c.Ctx.Redirect(302, "/true")
}

//神秘页面
type MysteryController struct {
	beego.Controller
}

func (c *MysteryController) Get()  {
	w:=c.Ctx.ResponseWriter
	r:=c.Ctx.Request
	sess, _ := utils.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	userName:=sess.Get("username")
	fmt.Println("=======",userName)
	if userName==nil{
		userName=""
	}
	fmt.Println("userName:",userName)
	allow:=utils.NewCasbinTool().Enforcer.Enforce(userName,"mystery.html","view")
	fmt.Println("allow:",allow)
	if allow{
		c.TplName="mystery.html"
	}else {
		c.Ctx.Redirect(302, "/false?des=no-permission")
	}
}

//授权
type GrantController struct {
	beego.Controller
}

func (c *GrantController) Get()  {
	c.TplName="grant.html"
}

func (c *GrantController) Post()  {
	w:=c.Ctx.ResponseWriter
	r:=c.Ctx.Request
	sess, _ := utils.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	userName:=sess.Get("username")
	fmt.Println("=======",userName)
	if userName==nil{
		userName=""
	}
	fmt.Println("userName:",userName)
	allow:=utils.NewCasbinTool().Enforcer.Enforce(userName,"permission","grant")
	fmt.Println("allow:",allow)
	if allow{
		userName:=c.GetString("user_name")
		user:=models.User{
			Name:userName,
		}
		err:=dao.NewDao().ReadUserName(user)
		if err == orm.ErrNoRows {
			c.Ctx.Redirect(302, "/false?des=no-user")
		} else if err == orm.ErrMissPK {
			c.Ctx.Redirect(302, "/false?des=user-can-not-find-primary-key")
		}

		roleName:=c.GetString("role_name")
		role:=models.Role{
			Name:roleName,
		}
		err=dao.NewDao().ReadRole(role)
		if err == orm.ErrNoRows {
			c.Ctx.Redirect(302, "/false?des=no-role")
		} else if err == orm.ErrMissPK {
			c.Ctx.Redirect(302, "/false?des=role-can-not-find-primary-key")
		}

		utils.NewCasbinTool().Enforcer.AddRoleForUser(userName,roleName)
		c.Ctx.Redirect(302, "/true")
	}else {
		c.Ctx.Redirect(302, "/false?des=no-permission")
	}
}

//给角色赋权
type RoleController struct {
	beego.Controller
}

func (c *RoleController) Get()  {
	c.TplName="role.html"
}

func (c *RoleController) Post()  {
	w:=c.Ctx.ResponseWriter
	r:=c.Ctx.Request
	sess, _ := utils.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	userName:=sess.Get("username")
	fmt.Println("=======",userName)
	if userName==nil{
		userName=""
	}
	fmt.Println("userName:",userName)
	allow:=utils.NewCasbinTool().Enforcer.Enforce(userName,"role","permission")
	fmt.Println("allow:",allow)
	if allow{
		roleName:=c.GetString("role_name")
		role:=models.Role{
			Name:roleName,
		}
		err:=dao.NewDao().ReadRole(role)
		if err == orm.ErrNoRows {
			c.Ctx.Redirect(302, "/false?des=no-role")
		} else if err == orm.ErrMissPK {
			c.Ctx.Redirect(302, "/false?des=role-can-not-find-primary-key")
		}

		sourceName:=c.GetString("source_name")
		source:=models.Source{
			Name:sourceName,
		}
		err=dao.NewDao().ReadSource(source)
		if err == orm.ErrNoRows {
			c.Ctx.Redirect(302, "/false?des=no-source")
		} else if err == orm.ErrMissPK {
			c.Ctx.Redirect(302, "/false?des=source-can-not-find-primary-key")
		}

		actionName:=c.GetString("action_name")
		action:=models.Action{
			Name:actionName,
		}
		err=dao.NewDao().ReadAction(action)
		if err == orm.ErrNoRows {
			c.Ctx.Redirect(302, "/false?des=no-action")
		} else if err == orm.ErrMissPK {
			c.Ctx.Redirect(302, "/false?des=action-can-not-find-primary-key")
		}

		utils.NewCasbinTool().Enforcer.AddPolicy(roleName,sourceName,actionName)
		c.Ctx.Redirect(302, "/true")
	}else {
		c.Ctx.Redirect(302, "/false?des=no-permission")
	}
}
