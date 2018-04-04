package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// ToStr interface to string
func SetValueFromStr(value reflect.Value, s string) error {
	switch value.Interface().(type) {
	case bool:
		val, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		value.SetBool(val)
	case float32:
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		value.SetFloat(val)
	case float64:
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		value.SetFloat(val)
	case int, int32:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		value.SetInt(val)
	case int8:
		val, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return err
		}
		value.SetInt(val)
	case int16:
		val, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return err
		}
		value.SetInt(val)
	case int64:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(val)
	case uint, uint32:
		val, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return err
		}
		value.SetUint(val)
	case uint8:
		val, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return err
		}
		value.SetUint(val)
	case uint16:
		val, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return err
		}
		value.SetUint(val)
	case uint64:
		val, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return err
		}
		value.SetUint(val)
	case string:
		value.SetString(s)
	case []byte:
		value.SetBytes([]byte(s))
	case []int32:
		var val []int32
		var err = json.Unmarshal([]byte(s), &val)
		if err != nil {
			return err
		}
		value.Set(reflect.ValueOf(val))
	default:
		return fmt.Errorf("unkown-type :%v", reflect.TypeOf(value))
	}
	return nil
}
