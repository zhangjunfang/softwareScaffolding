package lib

import (
	"os"
	"os/exec"
)

// CreateApp ...
// @desc 创建项目
func CreateApp() bool {
	cmd := exec.Command("bee", "new", HeadToUpper(GetDBName()))
	cmd.Start()
	cmd.Wait()
	return true
}

func Src() string {
	return os.Getenv("GOPATH") + "/src/"
}
