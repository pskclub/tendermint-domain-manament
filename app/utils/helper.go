package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func IsEmpty(x interface{}) bool {
	if x == nil {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func StructToString(v interface{}) (string, error) {
	value, err := json.MarshalIndent(v, "", " ")
	return string(value), err
}

func LogStruct(v interface{}) {
	value, _ := StructToString(v)
	fmt.Println(value)
}

func MapToStruct(src interface{}, to interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &to)
}
