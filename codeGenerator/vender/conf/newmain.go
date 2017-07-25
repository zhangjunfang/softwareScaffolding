package conf

import (
	"BeegoCURD/vender/lib"

	"os"

	"fmt"

	"github.com/astaxie/beego/logs"
)

func Newmain() {
	f, err := os.OpenFile(lib.Src()+lib.HeadToUpper(lib.GetDBName())+"/main.go", os.O_RDWR, 0766)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}
	_, err = f.WriteString(
		fmt.Sprintf(`
		package main

		import (
			_ "%s/routers"

			"github.com/astaxie/beego"
			"github.com/astaxie/beego/orm"
			_ "github.com/go-sql-driver/mysql"
		)

		func init() {
			orm.RegisterDataBase("default", "mysql", "%s")
		}
		
		func main() {
			beego.Run()
		}
		`, lib.HeadToUpper(lib.GetDBName()), lib.GetDBUrl()))
	f.Close()
	lib.GoFmt(lib.Src() + lib.HeadToUpper(lib.GetDBName()) + "/main.go")
}
