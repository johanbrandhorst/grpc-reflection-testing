// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/example.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	proto/example.proto

It has these top-level messages:
	User
	Void
*/
package example

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *User) Validate() error {
	if !(this.Id > 0) {
		return go_proto_validators.FieldError("Id", fmt.Errorf(`ID must a positive integer`))
	}
	return nil
}
func (this *Void) Validate() error {
	return nil
}
