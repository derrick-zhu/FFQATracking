package helpers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
)

// GetCurrentDirectory - 获取程序运行路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// AbosolutePath - get abosolute file path for given file name
func AbosolutePath(relateFilePath string) string {

	if len(relateFilePath) == 0 {
		return relateFilePath
	}
	execPath := utils.SelfDir()
	paths := []string{
		execPath,
		relateFilePath,
	}
	result := strings.Join(paths, "/")
	return result
}
