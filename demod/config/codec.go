/**
** @创建时间 : 2022/1/6 09:23
** @作者 : fzy
 */
package config

import (
	"demod/util/xencoding"
	"demod/util/xfile"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func unmarshal(info *xfile.FileInfo, target *map[string]interface{}) error {
	if codec := xencoding.GetCodec(info.Format); codec != nil {
		return codec.Unmarshal(info.Data, &target)
	}

	return errors.New(fmt.Sprintf("unsupported key: %s format: %s", info.Name, info.Format))
}

func marshalJSON(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}

func unmarshalJSON(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, m)
	}

	return json.Unmarshal(data, v)
}
