package helper

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoMarshal(m protoreflect.ProtoMessage) []byte {

	res, _ := proto.Marshal(m)

	return res
}
