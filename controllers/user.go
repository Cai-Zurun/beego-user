package controllers

import (
	"beego-user/models"
	"beego-user/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register()  {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err := models.AddUser(user); err != nil {
		c.Data["json"] = utils.Response(utils.FAIL, "注册失败", err)
		fmt.Println("Debug: 失败")
	} else {
		c.Data["json"] = utils.Response(utils.SUCCESS, "注册成功", nil)
	}
	c.ServeJSON()
}

