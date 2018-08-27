package utils

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/astaxie/beego"
)

// MD5 digist
func MD5(origin string) string {
	digist := md5.New()
	io.WriteString(digist, origin)
	result := fmt.Sprintf("%s", digist.Sum(nil))

	beego.Debug("MD5: %s -> %s", origin, result)

	return result
}
