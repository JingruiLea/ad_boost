package utils

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
)

type KV struct {
	Key   string
	Value interface{}
}

type SortedList []*KV

// MarshalJSON 实现json序列化方法
func (s SortedList) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	for i, kv := range s {
		if i != 0 {
			buf.WriteString(",") // 在项之间添加逗号，但不在第一项之前添加
		}

		// 对键进行JSON序列化
		key, err := jsoniter.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)

		buf.WriteString(":")

		// 对值进行JSON序列化
		value, err := jsoniter.Marshal(kv.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(value)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
