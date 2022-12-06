package util

import (
	"reflect"
	"strconv"
)

func InterfaceToFloat64(v interface{}) float64 {
	if v != nil {
		if reflect.TypeOf(v).Kind() == reflect.Float64 {
			return reflect.ValueOf(v).Float()
		} else if reflect.TypeOf(v).Kind() == reflect.String {
			if s, err := strconv.ParseFloat(v.(string), 64); err == nil {
				return s
			}
		}
	}

	return 0
}

func InterfaceToString(v interface{}) string {
	if v != nil && reflect.TypeOf(v).Kind() == reflect.String {
		return reflect.ValueOf(v).String()
	}
	return ""
}

func InterfaceToInt(v interface{}) int64 {
	if v != nil {
		if (reflect.TypeOf(v).Kind() == reflect.Int64) || reflect.TypeOf(v).Kind() == reflect.Int32 || reflect.TypeOf(v).Kind() == reflect.Int {
			return reflect.ValueOf(v).Int()
		} else if reflect.TypeOf(v).Kind() == reflect.String {
			if i, err := strconv.Atoi(v.(string)); err == nil {
				return int64(i)
			}
		}
	}
	return 0
}

func InterfaceToBool(v interface{}) bool {
	if v != nil && reflect.TypeOf(v).Kind() == reflect.Bool {
		return reflect.ValueOf(v).Bool()
	}
	return false
}

func IsExistField(v map[string]interface{}, key string) bool {
	if v, ok := v[key]; ok && v != nil {
		return true
	}
	return false
}
