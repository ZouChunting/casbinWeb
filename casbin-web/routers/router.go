package routers

import (
	"github.com/astaxie/beego"
	"zct/casbin-web/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/true", &controllers.TrueController{})
	beego.Router("/false", &controllers.FalseController{})
	beego.Router("/user/register", &controllers.RegisterController{})
	beego.Router("/user/login", &controllers.LoginController{})
	beego.Router("/user/logout", &controllers.LogoutController{})
	beego.Router("/mystery", &controllers.MysteryController{})
	beego.Router("/service/grant", &controllers.GrantController{})
	beego.Router("/service/role", &controllers.RoleController{})
}
