// Code generated by protoc-gen-go.
// source: mailinglist.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	mailinglist.proto

It has these top-level messages:
	SendEmailRequest
	SubscribeRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SendEmailRequest struct {
	SenderName    string `protobuf:"bytes,1,opt,name=sender_name,json=senderName" json:"sender_name,omitempty"`
	SenderEmail   string `protobuf:"bytes,2,opt,name=sender_email,json=senderEmail" json:"sender_email,omitempty"`
	Subject       string `protobuf:"bytes,3,opt,name=subject" json:"subject,omitempty"`
	Body          string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	ReceiverEmail string `protobuf:"bytes,5,opt,name=receiver_email,json=receiverEmail" json:"receiver_email,omitempty"`
}

func (m *SendEmailRequest) Reset()                    { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string            { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()               {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendEmailRequest) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *SendEmailRequest) GetSenderEmail() string {
	if m != nil {
		return m.SenderEmail
	}
	return ""
}

func (m *SendEmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendEmailRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendEmailRequest) GetReceiverEmail() string {
	if m != nil {
		return m.ReceiverEmail
	}
	return ""
}

type SubscribeRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubscribeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*SendEmailRequest)(nil), "appscode.mailinglist.v1beta1.SendEmailRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "appscode.mailinglist.v1beta1.SubscribeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MailingList service

type MailingListClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type mailingListClient struct {
	cc *grpc.ClientConn
}

func NewMailingListClient(cc *grpc.ClientConn) MailingListClient {
	return &mailingListClient{cc}
}

func (c *mailingListClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/SendEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.mailinglist.v1beta1.MailingList/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MailingList service

type MailingListServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*appscode_dtypes.VoidResponse, error)
	Subscribe(context.Context, *SubscribeRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterMailingListServer(s *grpc.Server, srv MailingListServer) {
	s.RegisterService(&_MailingList_serviceDesc, srv)
}

func _MailingList_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingList_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.mailinglist.v1beta1.MailingList/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MailingList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.mailinglist.v1beta1.MailingList",
	HandlerType: (*MailingListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _MailingList_SendEmail_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _MailingList_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailinglist.proto",
}

func init() { proto.RegisterFile("mailinglist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xe2, 0x40,
	0x18, 0xc6, 0xd3, 0x2e, 0x2c, 0xcb, 0xb0, 0xbb, 0x61, 0x27, 0x7b, 0x68, 0x1a, 0x76, 0xc5, 0x2a,
	0x09, 0xf1, 0xd0, 0xa6, 0x9a, 0x78, 0xf0, 0x88, 0xf1, 0xa6, 0x86, 0x40, 0xe2, 0xc1, 0x0b, 0x99,
	0xb6, 0x6f, 0xc8, 0x10, 0x3a, 0x53, 0x3b, 0x03, 0x09, 0x57, 0xee, 0x9e, 0xbc, 0xfb, 0x05, 0x3c,
	0x7b, 0xf5, 0x4b, 0xf8, 0x15, 0xfc, 0x20, 0xa6, 0x33, 0x6d, 0x41, 0x82, 0x72, 0x69, 0x3a, 0xef,
	0x9f, 0xe7, 0xfd, 0xbd, 0xcf, 0x0c, 0xfa, 0x13, 0x13, 0x3a, 0xa5, 0x6c, 0x3c, 0xa5, 0x42, 0xba,
	0x49, 0xca, 0x25, 0xc7, 0x2d, 0x92, 0x24, 0x22, 0xe4, 0x11, 0xb8, 0xeb, 0xb9, 0xb9, 0x1f, 0x80,
	0x24, 0xbe, 0xdd, 0x1a, 0x73, 0x3e, 0x9e, 0x82, 0x47, 0x12, 0xea, 0x11, 0xc6, 0xb8, 0x24, 0x92,
	0x72, 0x26, 0x74, 0xaf, 0xfd, 0xbf, 0xe8, 0xfd, 0x24, 0x7f, 0x50, 0x6a, 0x87, 0x3c, 0x56, 0x35,
	0x91, 0x5c, 0x24, 0x20, 0x3c, 0xf5, 0xd5, 0x45, 0xce, 0x93, 0x81, 0x9a, 0x43, 0x60, 0xd1, 0x45,
	0x36, 0x7f, 0x00, 0x77, 0x33, 0x10, 0x12, 0xef, 0xa1, 0x86, 0x00, 0x16, 0x41, 0x3a, 0x62, 0x24,
	0x06, 0xcb, 0x68, 0x1b, 0xdd, 0xfa, 0x00, 0xe9, 0xd0, 0x35, 0x89, 0x01, 0xef, 0xa3, 0x9f, 0x79,
	0x01, 0x64, 0x7d, 0x96, 0xa9, 0x2a, 0xf2, 0x26, 0x25, 0x85, 0x2d, 0x54, 0x13, 0xb3, 0x60, 0x02,
	0xa1, 0xb4, 0xbe, 0xa9, 0x6c, 0x71, 0xc4, 0x18, 0x55, 0x02, 0x1e, 0x2d, 0xac, 0x8a, 0x0a, 0xab,
	0x7f, 0xdc, 0x41, 0xbf, 0x53, 0x08, 0x81, 0xce, 0x4b, 0xc9, 0xaa, 0xca, 0xfe, 0x2a, 0xa2, 0x4a,
	0xd4, 0xe9, 0xa2, 0xe6, 0x70, 0x16, 0x88, 0x30, 0xa5, 0x01, 0x14, 0xb0, 0x7f, 0x51, 0x55, 0x77,
	0x68, 0x4c, 0x7d, 0x38, 0x7e, 0x31, 0x51, 0xe3, 0x4a, 0x5b, 0x7a, 0x49, 0x85, 0xc4, 0xf7, 0x06,
	0xaa, 0x97, 0x7b, 0x62, 0xd7, 0xfd, 0xca, 0x77, 0x77, 0xd3, 0x10, 0xfb, 0xdf, 0xaa, 0x5e, 0x7b,
	0xe8, 0xde, 0x70, 0x1a, 0x0d, 0x40, 0x24, 0x9c, 0x09, 0x70, 0xfc, 0xe5, 0xb3, 0x65, 0xfe, 0x30,
	0x96, 0xaf, 0x6f, 0x0f, 0x66, 0xc7, 0x69, 0x7b, 0xa3, 0x0f, 0x37, 0x93, 0x09, 0x79, 0xb9, 0xb4,
	0x37, 0x11, 0x9c, 0x9d, 0x19, 0x47, 0xf8, 0x31, 0xe3, 0x29, 0x56, 0xd9, 0xc9, 0xb3, 0xb1, 0xf3,
	0x2e, 0x9e, 0xde, 0x1a, 0xcf, 0xa9, 0xed, 0x6f, 0xe1, 0xc9, 0x27, 0x94, 0x58, 0xa2, 0x98, 0x50,
	0x00, 0xf6, 0xce, 0xd1, 0x61, 0xc8, 0xe3, 0xd5, 0x1c, 0x92, 0xd0, 0x6d, 0x6c, 0xbd, 0xe6, 0x9a,
	0xcb, 0xfd, 0xec, 0x49, 0xf5, 0x8d, 0xdb, 0x5a, 0x9e, 0x0c, 0xbe, 0xab, 0x47, 0x76, 0xf2, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x11, 0xaf, 0xba, 0xfa, 0x02, 0x00, 0x00,
}