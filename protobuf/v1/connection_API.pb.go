// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/connection_API.proto

package org_apache_geode_internal_protocol_protobuf_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HandshakeRequest struct {
	// Credentials used to authenticate with the security service
	// Not required if the server does not have a security service enabled
	Credentials map[string]string `protobuf:"bytes,1,rep,name=credentials" json:"credentials,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Set the serialization format for values. This should match the getID of a registered
	// org.apache.geode.protocol.serialization.ValueSerializer on the server side.
	//
	// If set any EncodedValue can be sent using EncodedValue.customObjectResult, and EncodedValues
	// sent by the server may also come back as EncodedValue.customObjectResult.
	//
	// If not set the server will attempt to convert non primative objects to JSON and
	// send them as EncodedValue.jsonResult.
	ValueFormat string `protobuf:"bytes,2,opt,name=valueFormat" json:"valueFormat,omitempty"`
}

func (m *HandshakeRequest) Reset()                    { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()               {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HandshakeRequest) GetCredentials() map[string]string {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *HandshakeRequest) GetValueFormat() string {
	if m != nil {
		return m.ValueFormat
	}
	return ""
}

type HandshakeResponse struct {
	Authenticated bool `protobuf:"varint,1,opt,name=authenticated" json:"authenticated,omitempty"`
}

func (m *HandshakeResponse) Reset()                    { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string            { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()               {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *HandshakeResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

type DisconnectClientRequest struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *DisconnectClientRequest) Reset()                    { *m = DisconnectClientRequest{} }
func (m *DisconnectClientRequest) String() string            { return proto.CompactTextString(m) }
func (*DisconnectClientRequest) ProtoMessage()               {}
func (*DisconnectClientRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DisconnectClientRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type DisconnectClientResponse struct {
}

func (m *DisconnectClientResponse) Reset()                    { *m = DisconnectClientResponse{} }
func (m *DisconnectClientResponse) String() string            { return proto.CompactTextString(m) }
func (*DisconnectClientResponse) ProtoMessage()               {}
func (*DisconnectClientResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func init() {
	proto.RegisterType((*HandshakeRequest)(nil), "org.apache.geode.internal.protocol.protobuf.v1.HandshakeRequest")
	proto.RegisterType((*HandshakeResponse)(nil), "org.apache.geode.internal.protocol.protobuf.v1.HandshakeResponse")
	proto.RegisterType((*DisconnectClientRequest)(nil), "org.apache.geode.internal.protocol.protobuf.v1.DisconnectClientRequest")
	proto.RegisterType((*DisconnectClientResponse)(nil), "org.apache.geode.internal.protocol.protobuf.v1.DisconnectClientResponse")
}

func init() { proto.RegisterFile("v1/connection_API.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x8b, 0x45, 0x27, 0x08, 0x71, 0x11, 0x1b, 0x7a, 0x0a, 0xc1, 0x43, 0x4f, 0x2b,
	0xd1, 0x8b, 0x7a, 0x10, 0xa4, 0x2a, 0x7a, 0xd3, 0xfc, 0x01, 0x99, 0x6e, 0xc6, 0x26, 0x34, 0xce,
	0xc6, 0xdd, 0x4d, 0xa0, 0x3f, 0xd9, 0x7f, 0x21, 0x4d, 0x22, 0xc6, 0x7a, 0xf2, 0x36, 0x33, 0xcc,
	0xfb, 0xde, 0xe3, 0xc1, 0xb4, 0x49, 0xce, 0x94, 0x66, 0x26, 0xe5, 0x0a, 0xcd, 0xaf, 0xb7, 0xcf,
	0x4f, 0xb2, 0x32, 0xda, 0x69, 0x21, 0xb5, 0x59, 0x49, 0xac, 0x50, 0xe5, 0x24, 0x57, 0xa4, 0x33,
	0x92, 0x05, 0x3b, 0x32, 0x8c, 0x65, 0xf7, 0xa0, 0x74, 0x3f, 0x2c, 0xeb, 0x37, 0xd9, 0x24, 0xf1,
	0xa7, 0x07, 0xc1, 0x23, 0x72, 0x66, 0x73, 0x5c, 0x53, 0x4a, 0x1f, 0x35, 0x59, 0x27, 0x2c, 0xf8,
	0xca, 0x50, 0x46, 0xec, 0x0a, 0x2c, 0x6d, 0xe8, 0x45, 0xe3, 0xb9, 0x7f, 0xfe, 0xf2, 0x4f, 0xb4,
	0xdc, 0xc5, 0xca, 0xc5, 0x0f, 0xf3, 0x9e, 0x9d, 0xd9, 0xa4, 0x43, 0x17, 0x11, 0x81, 0xdf, 0x60,
	0x59, 0xd3, 0x83, 0x36, 0xef, 0xe8, 0xc2, 0x51, 0xe4, 0xcd, 0x0f, 0xd2, 0xe1, 0x69, 0x76, 0x03,
	0xc1, 0x2e, 0x42, 0x04, 0x30, 0x5e, 0xd3, 0x26, 0xf4, 0xda, 0xef, 0xed, 0x28, 0x8e, 0x61, 0xaf,
	0x15, 0xf5, 0x84, 0x6e, 0xb9, 0x1e, 0x5d, 0x7a, 0xf1, 0x15, 0x1c, 0x0d, 0x32, 0xd9, 0x4a, 0xb3,
	0x25, 0x71, 0x0a, 0x87, 0x58, 0xbb, 0x7c, 0x0b, 0x55, 0xe8, 0x28, 0x6b, 0x51, 0xfb, 0xe9, 0xef,
	0x63, 0x9c, 0xc0, 0xf4, 0xae, 0xb0, 0x7d, 0xe3, 0x8b, 0xb2, 0x20, 0x76, 0xdf, 0x65, 0x9d, 0xc0,
	0xc4, 0x10, 0x5a, 0xcd, 0x7d, 0x88, 0x7e, 0x8b, 0x67, 0x10, 0xfe, 0x95, 0x74, 0xa6, 0xcb, 0x49,
	0xdb, 0xd3, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xd5, 0x5a, 0x0e, 0xc7, 0x01, 0x00,
	0x00,
}
