package proto

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	//"github.com/magicsea/ganet/config"
)

//json pb
var protoType string = "json"

//json pb 默认json
func SetProtoType(t string) {
	protoType = t
}

func Marshal(m proto.Message) ([]byte, error) {
	if protoType == "json" {
		return json.Marshal(m)
	}

	return proto.Marshal(m)
}
func Unmarshal(raw []byte, m proto.Message) error {
	if protoType == "json" {
		return json.Unmarshal(raw, m)
	}

	return proto.Unmarshal(raw, m)
}
