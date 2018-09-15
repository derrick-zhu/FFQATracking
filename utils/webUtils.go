package utils

import (
	"net/url"
)

// WebDataModel url object for redirect
type WebDataModel struct {
	URL  string
	Code int64
	Msg  string
}

// MakeRedirectURL append redirect url into post result
func MakeRedirectURL(data *map[interface{}]interface{}, code int64, url, msg string) {

	value := &WebDataModel{URL: url, Code: code, Msg: msg}
	(*data)["json"] = value
}

func CovertRequestInputToMap(values url.Values) map[string]interface{} {
	var result = make(map[string]interface{})

	for k, v := range values {
		result[k] = v[0]
	}

	return result
}
