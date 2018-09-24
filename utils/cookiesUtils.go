package utils

import (
	"FFQATracking/constants"
	"encoding/base64"
	"log"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// CookiesUtils class for manager Cookies
type CookiesUtils struct {
	Version string
	Data    map[string]string
}

var gCookieInstance *CookiesUtils
var gCookieOnce sync.Once

// CookieInstance cookie instance
func CookieInstance() *CookiesUtils {
	gCookieOnce.Do(func() {
		gCookieInstance = &CookiesUtils{}
	})
	return gCookieInstance
}

// Init initialize the cookie manager
func (cm *CookiesUtils) Init(ctx context.Context, version string) {
	if cm != CookieInstance() {
		log.Fatal("caller should using singleton handler")
	}
	cm.Version = version
}

// Set set value for key into cookie
func (cm *CookiesUtils) Set(ctx *context.Context, key string, value string, life int) {

	if cm != CookieInstance() {
		log.Fatal("caller should using singleton handler")
	}

	if life <= 0 {
		life = constants.MAXINT
	}
	ctx.SetCookie(key, value, life, "/")
}

// Get get value for key from cookie
func (cm *CookiesUtils) Get(ctx *context.Context, key string) string {

	if cm != CookieInstance() {
		log.Fatal("caller should using singleton handler")
	}

	ck, err := ctx.Request.Cookie(key)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

// SetSecret set value for key into cookie
func (cm *CookiesUtils) SetSecret(ctx *context.Context, key string, value string, life int) {

	if cm != CookieInstance() {
		log.Fatal("caller should using singleton handler")
	}

	encodedValue := Base64Encode(value)
	cm.Set(ctx, key, encodedValue, life)
}

// GetSecret get secret value from cookie
func (cm *CookiesUtils) GetSecret(ctx *context.Context, key string) string {

	if cm != CookieInstance() {
		log.Fatal("caller should using singleton handler")
	}

	encodedString := cm.Get(ctx, key)
	decodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return ""
	}
	return string(decodedData[:])
}
