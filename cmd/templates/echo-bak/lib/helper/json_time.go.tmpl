package helper

import (
    "database/sql/driver"
    "fmt"
    "strconv"
    "time"
)

// JSONTime format json time field by myself
type JSONTime struct {
    time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
    //formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
    //return []byte(formatted), nil
    //格式化秒
    ts := t.UnixNano() / 1e6
    return []byte(strconv.FormatInt(ts, 10)), nil
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t *JSONTime) UnmarshalJSON(data []byte) error {
    //formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
    //return []byte(formatted), nil
    //格式化秒
    val, err := strconv.ParseInt(string(data), 10, 64)
    if err != nil {
        return err
    }
    t.Time = time.Unix(0, val*1e6)
    return nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
    var zeroTime time.Time
    if t.Time.UnixNano() == zeroTime.UnixNano() {
        return nil, nil
    }
    return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
    value, ok := v.(time.Time)
    if ok {
        *t = JSONTime{Time: value}
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}

//转为毫秒
func (t *JSONTime) UnixMilli() int64 {
    return t.Time.UnixNano() / 1e6
}
