package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func Str2I64(s string, defaultValue ...int64) int64 {
	var ret int64
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return ret
}

func Str2Float64(s string, defaultValue ...float64) float64 {
	var ret float64
	ret, err := strconv.ParseFloat(s, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return ret
}

func GetJsonStr(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func Obj2Map(v interface{}) map[string]interface{} {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf(err.Error())
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(b, &ret)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return ret
}

func RoundFloat(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}
