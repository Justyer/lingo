package s2m

import (
	"errors"
	"reflect"
	"strconv"
)

// 根据tag作为map的key进行转换
// 如果没有该tag则不转换该字段
// 转换后的value为string
func KeyByTagToString(obj interface{}, tag string) (map[string]string, error) {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	m := make(map[string]string)
	for i := 0; i < k.NumField(); i++ {
		v_in := v.Field(i).Interface()
		var v_str string
		switch v_in.(type) {
		case int:
			v_str = strconv.Itoa(v_in.(int))
		case string:
			v_str = v_in.(string)
		case float64:
			v_str = strconv.FormatFloat(v_in.(float64), 'f', -1, 64)
		case struct{}:

		default:
			return m, errors.New("map have not supported type")
		}
		if t := k.Field(i).Tag.Get(tag); t != "" {
			m[t] = v_str
		}
	}
	return m, nil
}

// 根据tag
func KeyByTag(obj interface{}, tag string) map[string]interface{} {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	m := make(map[string]interface{})
	for i := 0; i < k.NumField(); i++ {
		if t := k.Field(i).Tag.Get(tag); t != "" {
			m[t] = v.Field(i).Interface()
		}
	}
	return m
}
