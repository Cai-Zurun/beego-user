package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisPool = newPool()
)

func newPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (c redis.Conn, err error) {
			return redis.Dial("tcp", beego.AppConfig.String("redis_host"), redis.DialPassword(beego.AppConfig.String("redis_pass")))
		},
		MaxIdle:     3,
		IdleTimeout: 60 * time.Second,
	}
}

func init()  {
	//logs
	logs.SetLogger(logs.AdapterFile,`{"filename":"error.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true)
}
