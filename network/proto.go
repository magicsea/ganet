package network

import (
	"encoding/json"

	"github.com/gogo/protobuf/proto"
	"github.com/magicsea/ganet/config"
)

func Marshal(m proto.Message) ([]byte, error) {
	if config.IsJsonProto() {
		return json.Marshal(m)
	}

	return proto.Marshal(m)
}
func Unmarshal(raw []byte,m proto.Message) error {
	if config.IsJsonProto() {
		return json.Unmarshal(raw,m)
	}

	return proto.Unmarshal(raw,m)
}
