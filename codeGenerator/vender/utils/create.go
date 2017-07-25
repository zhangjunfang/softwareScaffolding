package utils

import (
	"BeegoCURD/vender/lib"
	"os"

	"github.com/astaxie/beego/logs"
)

func Create() {
	os.Mkdir(lib.Src()+lib.HeadToUpper(lib.GetDBName())+"/utils", 0766)
	f, err := os.Create(lib.Src() + lib.HeadToUpper(lib.GetDBName()) + "/utils/log.go")
	f.Close()
	if err != nil {
		logs.Error(err.Error())
	}
	f, err = os.Create(lib.Src() + lib.HeadToUpper(lib.GetDBName()) + "/utils/tpl.go")
	f.Close()
	if err != nil {
		logs.Error(err.Error())
	}
}
