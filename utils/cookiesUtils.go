package utils

import (
	"FFQATracking/constants"
	"encoding/base64"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// CookiesUtils class for manager Cookies
type CookiesUtils struct {
	Version string
	Data    map[string]string
}

var cookieInstance *CookiesUtils
var cookieOnce sync.Once

// CookieInstance cookie instance
func CookieInstance() *CookiesUtils {
	cookieOnce.Do(func() {
		cookieInstance = &CookiesUtils{}
	})
	return cookieInstance
}

// Init initialize the cookie manager
func (cm *CookiesUtils) Init(ctx context.Context, version string) {
	cm.Version = version
}

// Set set value for key into cookie
func (cm *CookiesUtils) Set(ctx *context.Context, key string, value string, life int) {
	if life <= 0 {
		life = constants.MAXINT
	}
	beego.Info("key: " + key + ", value: " + value + ", life: " + strconv.Itoa(life))
	ctx.SetCookie(key, value, life, "/")
}

// Get get value for key from cookie
func (cm *CookiesUtils) Get(ctx *context.Context, key string) string {
	beego.Info("Cookie Get value for:" + key)
	ck, err := ctx.Request.Cookie(key)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

// SetSecret set value for key into cookie
func (cm *CookiesUtils) SetSecret(ctx *context.Context, key string, value string, life int) {

	encodedValue := Base64Encode(value)
	cm.Set(ctx, key, encodedValue, life)
}

// GetSecret get secret value from cookie
func (cm *CookiesUtils) GetSecret(ctx *context.Context, key string) string {

	encodedString := cm.Get(ctx, key)
	decodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return ""
	}
	return string(decodedData[:])
}
