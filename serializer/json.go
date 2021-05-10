package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ConvertProtobufToJson(message proto.Message) string {
	marshler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       " ",
		AnyResolver:  nil,
	}
	m, _ := marshler.MarshalToString(message)
	return m
}
