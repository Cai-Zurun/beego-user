
package routers

import (
	"github.com/astaxie/beego"
	"beego-user/controllers"
)

func init() {
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/user/password/update", &controllers.UserController{}, "post:UpdatePassword")
	beego.Router("/user/list/get", &controllers.UserController{}, "get:GetUserList")
	beego.Router("/user/email/send", &controllers.UserController{}, "post:SendEmail")
	beego.Router("/user/password/reset", &controllers.UserController{}, "post:ResetPassword")
}
