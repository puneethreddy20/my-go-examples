// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	AuthCode
	Action
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthCode struct {
	AuthCode string `protobuf:"bytes,1,opt,name=auth_code,json=authCode" json:"auth_code,omitempty"`
}

func (m *AuthCode) Reset()                    { *m = AuthCode{} }
func (m *AuthCode) String() string            { return proto.CompactTextString(m) }
func (*AuthCode) ProtoMessage()               {}
func (*AuthCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthCode) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

type Action struct {
	Action string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Action) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Action) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthCode)(nil), "main.AuthCode")
	proto.RegisterType((*Action)(nil), "main.Action")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc,
	0x53, 0x52, 0xe7, 0xe2, 0x70, 0x2c, 0x2d, 0xc9, 0x70, 0xce, 0x4f, 0x49, 0x15, 0x92, 0xe6, 0xe2,
	0x4c, 0x2c, 0x2d, 0xc9, 0x88, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0xe2, 0x48, 0x84, 0x4a, 0x2a, 0x99, 0x70, 0xb1, 0x39, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x09, 0x89,
	0x71, 0xb1, 0x25, 0x82, 0x59, 0x50, 0x35, 0x50, 0x9e, 0x90, 0x10, 0x17, 0x4b, 0x52, 0x7e, 0x4a,
	0xa5, 0x04, 0x13, 0x58, 0x14, 0xcc, 0x4e, 0x62, 0x03, 0xdb, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x64, 0xec, 0x0b, 0xbb, 0x7d, 0x00, 0x00, 0x00,
}