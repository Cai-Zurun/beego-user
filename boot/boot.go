package boot

import (
	"github.com/astaxie/beego/logs"
)



func init()  {
	//logs
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/error.log"}`)
	logs.EnableFuncCallDepth(true)
}
