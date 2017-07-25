package conf

import (
	"BeegoCURD/vender/lib"
	"os"

	"fmt"

	"github.com/astaxie/beego/logs"
)

func Write() {
	f, err := os.OpenFile(lib.Src()+lib.HeadToUpper(lib.GetDBName())+"/controllers/default.go", os.O_RDWR, 0766)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}
	f.WriteString(fmt.Sprintf(`
	package controllers

	import (
		"Pj/utils"
		"github.com/astaxie/beego"
	)

	type MainController struct {
		beego.Controller
	}

	func (c *MainController) Get() {
		// Data
		var tpl utils.Tpl
		tpl.Title = "BeegoCURD"
		tpl.Success = c.GetSession("success")
		tpl.Danger = c.GetSession("danger")
		c.DelSession("success")
		c.DelSession("danger")
		c.Data["Tpl"] = tpl
		c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.Layout = "public/tpl.html"
		c.TplName = "index.tpl"
	}
	`))
	f.Close()
	lib.GoFmt(lib.Src() + lib.HeadToUpper(lib.GetDBName()) + "/controllers/default.go")
}
