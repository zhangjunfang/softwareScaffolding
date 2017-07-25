package lib

import "github.com/astaxie/beego/orm"
import "fmt"

type Table struct {
	TableName string `orm:"column(table_name)"`
}

// GetAllTables ...
// @return []slice
// @desc 获取数据库中所有表名
func GetAllTables() []Table {
	var tablenames []Table
	o := orm.NewOrm()
	o.Raw(fmt.Sprintf("select distinct table_name from information_schema.columns where table_schema='%s'", GetDBName())).QueryRows(&tablenames)
	return tablenames
}
