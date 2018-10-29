package helpers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

/* 获取程序运行路径 */
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
