package helper

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type MapStrInt map[string]int

func (t MapStrInt) Value() (driver.Value, error) {
	data, err := json.Marshal(t)
	return string(data), err
}

func (t *MapStrInt) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if ok {
		return json.Unmarshal(value, t)
	}
	return fmt.Errorf("can not convert %v to map str int", v)
}

type MapStrInterface map[string]interface{}

func (t MapStrInterface) Value() (driver.Value, error) {
	data, err := json.Marshal(t)
	return string(data), err
}

func (t *MapStrInterface) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if ok {
		return json.Unmarshal(value, t)
	}
	return fmt.Errorf("can not convert %v to map str str", v)
}

type SliceUint64 []uint64

func (t SliceUint64) Value() (driver.Value, error) {
	data, err := json.Marshal(t)
	return string(data), err
}

func (t *SliceUint64) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if ok {
		if len(value) == 0 {
			return nil
		} else {
			return json.Unmarshal(value, t)
		}
	}
	return fmt.Errorf("can not convert %v to slice uint64", v)
}

type SliceString []string

func (t SliceString) Value() (driver.Value, error) {
	data, err := json.Marshal(t)
	return string(data), err
}

func (t *SliceString) Scan(v interface{}) error {
	value, ok := v.([]byte)
	if ok {
		if len(value) == 0 {
			return nil
		} else {
			return json.Unmarshal(value, t)
		}
	}
	return fmt.Errorf("can not convert %v to slice string", v)
}

//为 js 准备的 bigint
type JsBigInt uint64

func (i JsBigInt) MarshalJSON() ([]byte, error) {
	return []byte("\"" + strconv.FormatUint(uint64(i), 10) + "\""), nil
}

func (i *JsBigInt) UnmarshalJSON(data []byte) error {
	val, err := strconv.ParseUint(strings.Trim(string(data), "\""), 10, 64)
	if err != nil {
		return err
	}
	*i = JsBigInt(val)
	return nil
}

func (i JsBigInt) Value() (driver.Value, error) {
	return int64(i), nil
}

func (i *JsBigInt) Scan(v interface{}) error {
	val, ok := v.(int64)
	if ok {
		*i = JsBigInt(val)
		return nil
	} else {
		val, ok := v.([]uint8)
		if ok {
			v, _ := strconv.ParseInt(string(val), 10, 64)
			*i = JsBigInt(v)
			return nil
		}
	}
	return fmt.Errorf("can not convert %v to js big int", v)
}
