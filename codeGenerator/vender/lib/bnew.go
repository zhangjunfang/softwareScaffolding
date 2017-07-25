package lib

import (
	"os"
	"os/exec"
)

// CreateApp ...
// @desc create project
func CreateApp() bool {
	cmd := exec.Command("bee", "new", HeadToUpper(GetDBName()))
	cmd.Start()
	cmd.Wait()
	return true
}

func Src() string {
	return os.Getenv("GOPATH") + "/src/"
}
