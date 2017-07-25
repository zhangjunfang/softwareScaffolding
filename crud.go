package main

import (
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/conf"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/controller"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/lib"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/model"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/router"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/utils"
	"github.com/zhangjunfang/softwareScaffolding/codeGenerator/vender/view"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", lib.GetDBUrl())
}

func main() {
	if lib.CreateApp() {
		router.Write()
		conf.AddSessionOn()
		conf.Newmain()
		conf.Write()
		controller.Create()
		controller.Write()
		model.Create()
		model.Write()
		view.Create()
		view.Write()
		utils.Create()
		utils.Write()
	} else {
		logs.Error("项目创建失败")
	}
}

// go run github.com/zhangjunfang/softwareScaffolding.go -conn="root:lione520520@tcp(127.0.0.1:3306)/lione_blog"
