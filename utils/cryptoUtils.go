package utils

import (
	"crypto/md5"
	"encoding/base64"
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

// Base64Encode encode string with base64
func Base64Encode(origin string) string {
	result := base64.StdEncoding.EncodeToString([]byte(origin))

	beego.Debug(fmt.Sprintf("`Base64Encode`: %s -> %s", origin, result))
	return result
}

// Base64Decode decode encoded string with base64
func Base64Decode(encoded string) string {
	result, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		beego.Error(fmt.Sprintf("decode error: %s", err))
		return ""
	}
	beego.Debug("`Base64Decode`: %s -> %s", encoded, result)
	return string(result)
}
