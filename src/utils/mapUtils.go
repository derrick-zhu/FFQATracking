package utils

import (
	"reflect"
	"strings"
)

// Struct2Map struct mapping to map
func Struct2Map(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		if IsEmpty(v.Field(i)) == false {

			data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
		}
	}

	return data
}

// IsEmpty if any object is empty or not
func IsEmpty(a interface{}) bool {

	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {

		v = v.Elem()
	}

	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
