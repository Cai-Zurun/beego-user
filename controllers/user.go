package controllers

import (
	"beego-user/models"
	"beego-user/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

const (
	SALT = "fdsafagfdgv43532ju76jM"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Register()  {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Password = utils.Md5(user.Password + SALT)
	if err := models.AddUser(user); err != nil {
		logs.Error(err)
		u.Data["json"] = utils.Response(utils.FAIL, "注册失败", err.Error())
	} else {
		u.Data["json"] = utils.Response(utils.SUCCESS, "注册成功", nil)
	}
	u.ServeJSON()
}

func (u *UserController) Login()  {
	email := u.GetString("email")
	password := u.GetString("password")
	password = utils.Md5(password + SALT)
	if user, err := models.ValidateUser(email, password); err != nil {
		logs.Error(err)
		u.Data["json"] = utils.Response(utils.FAIL, "登录失败", err.Error())
	} else {
		if token, err := utils.GenerateToken(user.Email); err != nil {
			logs.Error(err)
			u.Data["json"] = utils.Response(utils.FAIL, "登录失败", err.Error())
		} else {
			res := map[string]string{
				"token" : token,
			}
			u.Data["json"] = utils.Response(utils.SUCCESS, "登录成功", res)
		}
	}
	u.ServeJSON()
}
func (u *UserController) UpdatePassword()  {
	email := u.GetString("email")
	password := u.GetString("password")
	fmt.Println("测试：",email, password)
	newPassword := u.GetString("newPassword")
	password = utils.Md5(password + SALT)
	fmt.Println(password)
	if user, err := models.ValidateUser(email, password); err != nil {
		logs.Error(err)
		u.Data["json"] = utils.Response(utils.FAIL, "修改密码失败", err.Error())
	} else {
		user.Password = utils.Md5(newPassword + SALT)
		if err := models.UpdateUser(user); err != nil {
			logs.Error(err)
			u.Data["json"] = utils.Response(utils.FAIL, "修改密码失败", err.Error())
		}
		u.Data["json"] = utils.Response(utils.FAIL, "修改密码成功", nil)
	}
	u.ServeJSON()
}

func (u *UserController) SendEmail()  {
	email := u.GetString("email")
	if _, err := models.ValidateEmail(email); err != nil {
		u.Data["json"] = utils.Response(utils.FAIL, "用户不存在", err.Error())
	} else {
		subject := "重置密码"
		captcha := utils.CreateCaptcha()
		body := "您的验证码是：" + captcha
		if err := utils.SendMail(email, subject, body); err != nil {
			logs.Error(err)
			u.Data["json"] = utils.Response(utils.FAIL, "邮件发送失败", err.Error())
		} else {
			if _, err := utils.RedisPool.Get().Do("SETEX", email, 60, captcha); err != nil{
				logs.Error(err)
				u.Data["json"] = utils.Response(utils.FAIL, "Redis异常", err.Error())
			}
			u.Data["json"] = utils.Response(utils.SUCCESS, "邮件发送成功", nil)
		}
	}
	u.ServeJSON()
}

func (u *UserController) ResetPassword()  {
	email := u.GetString("email")
	captcha := u.GetString("captcha")
	newPassword := u.GetString("newPassword")
	if realCaptcha, err := redis.String(utils.RedisPool.Get().Do("GET", email)); err != nil{
		logs.Error(err)
		u.Data["json"] = utils.Response(utils.FAIL, "Redis异常", err.Error())
	} else {
		if realCaptcha != "" {
			fmt.Println(captcha)
			fmt.Println(realCaptcha)
			if captcha != realCaptcha {
				u.Data["json"] = utils.Response(utils.FAIL, "验证码错误", nil)
			} else {
				user, _ := models.ValidateEmail(email)
				user.Password = utils.Md5(newPassword + SALT)
				if err := models.UpdateUser(user); err != nil {
					logs.Error(err)
					u.Data["json"] = utils.Response(utils.FAIL, "重置密码失败", err.Error())
				}
				u.Data["json"] = utils.Response(utils.FAIL, "重置密码成功", nil)
			}
		} else {
			u.Data["json"] = utils.Response(utils.FAIL, "验证码过期", nil)
		}
	}
	u.ServeJSON()
}

func (u *UserController) GetUserList()  {
	if users, err := models.GetUserList(); err != nil {
		logs.Error(err)
		u.Data["json"] = utils.Response(utils.FAIL, "获取用户名列表失败", err.Error())
	} else {
		userNames := make([]string, len(users))
		for i := 0; i < len(users); i++ {
			userNames[i] = users[i].Name
		}
		u.Data["json"] = utils.Response(utils.SUCCESS, "获取用户名列表成功", userNames)
	}
	u.ServeJSON()
}
