package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string	`gorm:"unique"`
	Email string	`valid:"Required; Email" gorm:"not null; unique"`
	Password string	`valid:"Required" gorm:"not null; unique"`
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
	if err := MysqlClient.Create(&u).Error ;err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func ValidateUser(email, password string) (User, error){
	var user User
	if err := MysqlClient.Table("users").Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func ValidateEmail(email string) (User, error) {
	var user User
	if err := MysqlClient.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(u User) error {
	if err := MysqlClient.Table("users").Save(&u).Error ;err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetUserList() ([]User, error) {
	var users []User
	if err := MysqlClient.Table("users").Find(&users).Error ;err != nil {
		logs.Error(err)
		return nil, err
	}
	return users, nil
}