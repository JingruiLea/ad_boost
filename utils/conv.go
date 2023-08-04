package utils

import "strconv"

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
