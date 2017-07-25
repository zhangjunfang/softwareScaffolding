package controller

import (
	"BeegoCURD/vender/lib"
	"os"

	"github.com/astaxie/beego/logs"
)

func Create() {
	for _, v := range lib.GetAllTables() {
		s := string(v.TableName)
		f, err := os.Create(lib.Src() + lib.HeadToUpper(lib.GetDBName()) + "/controllers/" + lib.HeadToUpper(s) + ".go")
		defer f.Close()
		if err != nil {
			logs.Error(err.Error())
		}
	}
}
