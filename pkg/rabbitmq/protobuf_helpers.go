package rabbitmq

import (
	"google.golang.org/protobuf/proto"
)

// MarshalProto serializes a Protobuf message
func MarshalProto(message proto.Message) ([]byte, error) {
	return proto.Marshal(message)
}

// UnmarshalProto deserializes a Protobuf message
func UnmarshalProto(data []byte, message proto.Message) error {
	return proto.Unmarshal(data, message)
}
