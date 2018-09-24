package utils

import (
	"net/url"
)

// WebDataModel url object for redirect
type WebDataModel struct {
	URL      string
	Code     int64
	Msg      string
	UserInfo interface{}
}

// MakeRedirectURL append redirect url into post result
func MakeRedirectURL(data *map[interface{}]interface{}, code int64, url, msg string) {

	MakeRedirectURLWithUserInfo(data, code, url, msg, nil)
}

// MakeRedirectURLWithUserInfo append redirect url into post result
func MakeRedirectURLWithUserInfo(data *map[interface{}]interface{}, code int64, url, msg string, userInfo interface{}) {
	value := &WebDataModel{URL: url, Code: code, Msg: msg, UserInfo: userInfo}
	(*data)["json"] = value
}

// CovertRequestInputToMap cover request arguents into map
func CovertRequestInputToMap(values url.Values) map[string]interface{} {
	var result = make(map[string]interface{})

	for k, v := range values {
		result[k] = v[0]
	}

	return result
}
