package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//dsn = "root:@tcp(localhost:3306)/beego-user?charset=utf8&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", beego.AppConfig.String("db_user"), beego.AppConfig.String("db_pass"), beego.AppConfig.String("db_host"), beego.AppConfig.String("db_port"), beego.AppConfig.String("db_name"))
	MysqlClient, err = gorm.Open("mysql", dsn)
)

func init() {
	if err != nil {
		logs.Error(err)
	}
	MysqlClient.AutoMigrate(&User{})
}
