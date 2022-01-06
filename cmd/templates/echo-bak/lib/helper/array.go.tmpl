package helper

import (
    "reflect"
    "strings"
)

func InterfaceSlice(slice interface{}) []interface{} {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("InterfaceSlice() given a non-slice type")
    }

    ret := make([]interface{}, s.Len())

    for i := 0; i < s.Len(); i++ {
        ret[i] = s.Index(i).Interface()
    }

    return ret
}

func ArrayDiffString(a, b []string) []string {
    mb := make(map[string]struct{}, len(b))
    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []string
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}

func ArrayDiffUint32(a, b []uint32) []uint32 {
    mb := make(map[uint32]struct{}, len(b))
    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []uint32
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}

func ArrayDiffUint64(a, b []uint64) []uint64 {
    mb := make(map[uint64]struct{}, len(b))
    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []uint64
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}

// 数组去重
func RemoveDuplicateString(a []string) []string {
    ret := make([]string, 0, len(a))
    temp := map[string]struct{}{}
    for _, item := range a {
        itemMd5 := Md5(strings.Trim(item, " "))
        if _, ok := temp[itemMd5]; !ok {
            temp[itemMd5] = struct{}{}
            ret = append(ret, item)
        }
    }
    return ret
}

func RemoveSlice(a []uint64, b uint64) []uint64 {

    var newData []uint64

    for _, value := range a {
        if value != b {
            newData = append(newData, value)
        }
    }

    return newData
}

func RemoveSliceByString(a []string, b string) []string {

    var newData []string

    for _, value := range a {
        if value != b {
            newData = append(newData, value)
        }
    }

    return newData
}
