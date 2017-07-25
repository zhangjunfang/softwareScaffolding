package conf

import (
	"os"

	"BeegoCURD/vender/lib"

	"github.com/astaxie/beego/logs"
)

// AddSessionOn ...
// @desc 开启项目session功能
func AddSessionOn() {
	f, err := os.OpenFile(lib.Src()+lib.HeadToUpper(lib.GetDBName())+"/conf/app.conf", os.O_WRONLY, 0766)
	if err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}
	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte("sessionOn = true"), n)
}
