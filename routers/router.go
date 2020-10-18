
package routers

import (
	"github.com/astaxie/beego"
	"beego-user/controllers"
)

func init() {
	beego.Router("/register", &controllers.UserController{}, "post:Register")
}
