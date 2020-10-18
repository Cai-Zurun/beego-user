package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"

	//"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	//"github.com/astaxie/beego/logs"
)

var (
	//o = orm.NewOrm()
	o = GetOrmObject()
)

type User struct {
	Id int	`orm:"pk;auto;unique"`
	Name string	`orm:"unique"`
	Email string	`valid:"Required; Email" orm:"unique"`
	Password string	`valid:"Required" orm:"unique"`
}

func AddUser(u User)  error{
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			logs.Error(err)
			return err
		}
	}
	//if _, err := o.Insert(u) ;err != nil {
	//	fmt.Println(err)
	//	//logs.Error(err)
	//	return err
	//}
	return nil
}
