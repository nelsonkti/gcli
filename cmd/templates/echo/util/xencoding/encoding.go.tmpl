/**
** @创建时间 : 2022/1/4 17:25
** @作者 : fzy
 */
package xencoding

import (
	"strings"
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

var registered = make(map[string]Codec)

func RegisterCodec(codec Codec) {
	if codec == nil {
		panic("cannot register a nil Codec")
	}
	if codec.Name() == "" {
		panic("cannot register Codec with empty string result for Name()")
	}
	contentSubtype := strings.ToLower(codec.Name())
	registered[contentSubtype] = codec
}

func GetCodec(key string) Codec {
	return registered[key]
}
