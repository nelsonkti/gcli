package helper

import (
    "database/sql/driver"
    "fmt"
    "time"
)

// JSONTimeDate format json time field by myself
type JSONTimeDate struct {
    time.Time
}

// MarshalJSON on JSONTimeDate format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTimeDate) MarshalJSON() ([]byte, error) {
    formatted := fmt.Sprintf("\"%s\"", t.Format(InitDateTime))
    return []byte(formatted), nil
}

// MarshalJSON on JSONTimeDate format Time field with %Y-%m-%d %H:%M:%S
func (t *JSONTimeDate) UnmarshalJSON(data []byte) ([]byte, error) {
    formatted := fmt.Sprintf("\"%s\"", t.Format(InitDateTime))
    return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTimeDate) Value() (driver.Value, error) {
    var zeroTime time.Time
    if t.Time.UnixNano() == zeroTime.UnixNano() {
        return nil, nil
    }
    return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTimeDate) Scan(v interface{}) error {
    value, ok := v.(time.Time)
    if ok {
        *t = JSONTimeDate{Time: value}
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}

//转为毫秒
func (t *JSONTimeDate) UnixMilli() int64 {
    return t.Time.UnixNano() / 1e6
}
