package lib

import (
	"os/exec"
)

// GoFmt ...
// @desc 代码格式化
func GoFmt(filename string) {
	cmd := exec.Command("go", "fmt", filename)
	cmd.Start()
	cmd.Wait()
}
