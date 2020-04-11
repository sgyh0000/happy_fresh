package log

import "github.com/astaxie/beego/logs"

func init() {
	_ = logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
