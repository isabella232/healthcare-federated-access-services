// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consents/v1/consents.proto

// Package v1 provides protocol buffer versions of OAuth Consents API for
// listing and revoking OAuth consents.
package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// OAuth Consent.
type Consent struct {
	// Name of the OAuth Consent.
	// Format: `consents/{consent_id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifies the user who gave the OAuth consent.
	// E.g. subject or account number
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Identifies the client for which the OAuth consent was given.
	Client string `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// Identifies the items for which the OAuth consent was given.
	// E.g. JTI of a Visa JWT.
	Items []string `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Identifies the scopes for which the OAuth consent was given.
	Scopes []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Identifies the resources for which the OAuth consent was given.
	Resouces []string `protobuf:"bytes,6,rep,name=resouces,proto3" json:"resouces,omitempty"`
	// Time at which OAuth consent was first given.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time at which consent was last updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Consent) Reset()         { *m = Consent{} }
func (m *Consent) String() string { return proto.CompactTextString(m) }
func (*Consent) ProtoMessage()    {}
func (*Consent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{0}
}

func (m *Consent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent.Unmarshal(m, b)
}
func (m *Consent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent.Marshal(b, m, deterministic)
}
func (m *Consent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent.Merge(m, src)
}
func (m *Consent) XXX_Size() int {
	return xxx_messageInfo_Consent.Size(m)
}
func (m *Consent) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent.DiscardUnknown(m)
}

var xxx_messageInfo_Consent proto.InternalMessageInfo

func (m *Consent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consent) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Consent) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *Consent) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Consent) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Consent) GetResouces() []string {
	if m != nil {
		return m.Resouces
	}
	return nil
}

func (m *Consent) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Consent) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type GetConsentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConsentRequest) Reset()         { *m = GetConsentRequest{} }
func (m *GetConsentRequest) String() string { return proto.CompactTextString(m) }
func (*GetConsentRequest) ProtoMessage()    {}
func (*GetConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{1}
}

func (m *GetConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConsentRequest.Unmarshal(m, b)
}
func (m *GetConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConsentRequest.Marshal(b, m, deterministic)
}
func (m *GetConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConsentRequest.Merge(m, src)
}
func (m *GetConsentRequest) XXX_Size() int {
	return xxx_messageInfo_GetConsentRequest.Size(m)
}
func (m *GetConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConsentRequest proto.InternalMessageInfo

func (m *GetConsentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteConsentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConsentRequest) Reset()         { *m = DeleteConsentRequest{} }
func (m *DeleteConsentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConsentRequest) ProtoMessage()    {}
func (*DeleteConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{2}
}

func (m *DeleteConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConsentRequest.Unmarshal(m, b)
}
func (m *DeleteConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConsentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConsentRequest.Merge(m, src)
}
func (m *DeleteConsentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConsentRequest.Size(m)
}
func (m *DeleteConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConsentRequest proto.InternalMessageInfo

func (m *DeleteConsentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListConsentsRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConsentsRequest) Reset()         { *m = ListConsentsRequest{} }
func (m *ListConsentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListConsentsRequest) ProtoMessage()    {}
func (*ListConsentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{3}
}

func (m *ListConsentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsentsRequest.Unmarshal(m, b)
}
func (m *ListConsentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsentsRequest.Marshal(b, m, deterministic)
}
func (m *ListConsentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsentsRequest.Merge(m, src)
}
func (m *ListConsentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListConsentsRequest.Size(m)
}
func (m *ListConsentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsentsRequest proto.InternalMessageInfo

func (m *ListConsentsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListConsentsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListConsentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConsentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListConsentsResponse struct {
	Consents             []*Consent `protobuf:"bytes,1,rep,name=consents,proto3" json:"consents,omitempty"`
	NextPageToken        string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListConsentsResponse) Reset()         { *m = ListConsentsResponse{} }
func (m *ListConsentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListConsentsResponse) ProtoMessage()    {}
func (*ListConsentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{4}
}

func (m *ListConsentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsentsResponse.Unmarshal(m, b)
}
func (m *ListConsentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsentsResponse.Marshal(b, m, deterministic)
}
func (m *ListConsentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsentsResponse.Merge(m, src)
}
func (m *ListConsentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListConsentsResponse.Size(m)
}
func (m *ListConsentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsentsResponse proto.InternalMessageInfo

func (m *ListConsentsResponse) GetConsents() []*Consent {
	if m != nil {
		return m.Consents
	}
	return nil
}

func (m *ListConsentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Consent)(nil), "consents.v1.Consent")
	proto.RegisterType((*GetConsentRequest)(nil), "consents.v1.GetConsentRequest")
	proto.RegisterType((*DeleteConsentRequest)(nil), "consents.v1.DeleteConsentRequest")
	proto.RegisterType((*ListConsentsRequest)(nil), "consents.v1.ListConsentsRequest")
	proto.RegisterType((*ListConsentsResponse)(nil), "consents.v1.ListConsentsResponse")
}

func init() { proto.RegisterFile("proto/consents/v1/consents.proto", fileDescriptor_0bb2e176a5b7c192) }

var fileDescriptor_0bb2e176a5b7c192 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x8f, 0x26, 0x93, 0x16, 0xc4, 0x12, 0x82, 0xe5, 0x52, 0xe1, 0xfa, 0x00, 0x51,
	0xa4, 0xd8, 0x34, 0x1c, 0x11, 0x17, 0x0a, 0xea, 0x85, 0x43, 0x65, 0x8a, 0x90, 0xb8, 0x44, 0x1b,
	0x67, 0x92, 0xac, 0xb0, 0xbd, 0xc6, 0xbb, 0x8e, 0xa0, 0x88, 0x03, 0xdc, 0x39, 0xf1, 0xd3, 0xf8,
	0x03, 0x1c, 0xf8, 0x21, 0x68, 0x77, 0x6d, 0x2b, 0xa1, 0x11, 0x70, 0x9b, 0xf7, 0xe6, 0x8d, 0x67,
	0xe6, 0xed, 0xae, 0xc1, 0xcd, 0x72, 0x2e, 0x79, 0x10, 0xf1, 0x54, 0x60, 0x2a, 0x45, 0xb0, 0x3e,
	0xad, 0x63, 0x5f, 0xa7, 0x48, 0xaf, 0xc6, 0xeb, 0x53, 0xe7, 0xde, 0x92, 0xf3, 0x65, 0x8c, 0x01,
	0xcd, 0x58, 0x40, 0xd3, 0x94, 0x4b, 0x2a, 0x19, 0x4f, 0x4b, 0xa9, 0x73, 0x54, 0x66, 0x35, 0x9a,
	0x15, 0x8b, 0x00, 0x93, 0x4c, 0x7e, 0x2c, 0x93, 0xf7, 0xff, 0x4c, 0x4a, 0x96, 0xa0, 0x90, 0x34,
	0xc9, 0x8c, 0xc0, 0xfb, 0xb6, 0x07, 0xfb, 0x67, 0xa6, 0x17, 0x21, 0xd0, 0x4c, 0x69, 0x82, 0xb6,
	0xe5, 0x5a, 0xc3, 0x6e, 0xa8, 0x63, 0xc5, 0x15, 0x02, 0x73, 0x7b, 0xcf, 0x70, 0x2a, 0x26, 0x03,
	0x68, 0x47, 0x31, 0xc3, 0x54, 0xda, 0x0d, 0xcd, 0x96, 0x88, 0xf4, 0xa1, 0xc5, 0x24, 0x26, 0xc2,
	0x6e, 0xba, 0x8d, 0x61, 0x37, 0x34, 0x40, 0xa9, 0x45, 0xc4, 0x33, 0x14, 0x76, 0x4b, 0xd3, 0x25,
	0x22, 0x0e, 0x74, 0x72, 0x14, 0xbc, 0x88, 0x50, 0xd8, 0x6d, 0x9d, 0xa9, 0x31, 0x79, 0x02, 0xbd,
	0x28, 0x47, 0x2a, 0x71, 0xaa, 0xe6, 0xb5, 0xf7, 0x5d, 0x6b, 0xd8, 0x9b, 0x38, 0xbe, 0x59, 0xc6,
	0xaf, 0x96, 0xf1, 0x2f, 0xab, 0x65, 0x42, 0x30, 0x72, 0x45, 0xa8, 0xe2, 0x22, 0x9b, 0xd7, 0xc5,
	0x9d, 0x7f, 0x17, 0x1b, 0xb9, 0x22, 0xbc, 0x87, 0x70, 0xeb, 0x1c, 0x65, 0xe9, 0x48, 0x88, 0xef,
	0x0b, 0x14, 0x3b, 0x8d, 0xf1, 0x46, 0xd0, 0x7f, 0x8e, 0x31, 0x4a, 0xfc, 0x0f, 0xed, 0x17, 0x0b,
	0x6e, 0xbf, 0x64, 0xa2, 0xfa, 0xac, 0xa8, 0xb4, 0x03, 0x68, 0x67, 0x34, 0x57, 0x46, 0x1a, 0x75,
	0x89, 0x14, 0xbf, 0x60, 0xb1, 0xac, 0x6d, 0x2f, 0x11, 0x39, 0x82, 0x6e, 0x46, 0x97, 0x38, 0x15,
	0xec, 0x0a, 0xb5, 0xf7, 0xad, 0xb0, 0xa3, 0x88, 0x57, 0xec, 0x0a, 0xc9, 0x31, 0x80, 0x4e, 0x4a,
	0xfe, 0x0e, 0x53, 0xbb, 0xa9, 0x0b, 0xb5, 0xfc, 0x52, 0x11, 0x5e, 0x06, 0xfd, 0xed, 0x11, 0x44,
	0xa6, 0x42, 0xf2, 0x08, 0x3a, 0xd5, 0x5d, 0xb3, 0x2d, 0xb7, 0x31, 0xec, 0x4d, 0xfa, 0xfe, 0xc6,
	0xe5, 0xf3, 0xab, 0xf5, 0x6a, 0x15, 0x79, 0x00, 0x37, 0x53, 0xfc, 0x20, 0xa7, 0x1b, 0xdd, 0xcc,
	0x98, 0x87, 0x8a, 0xbe, 0xa8, 0x3a, 0x4e, 0x7e, 0x5a, 0xd0, 0xa9, 0xda, 0x11, 0x06, 0x87, 0x5b,
	0x76, 0x91, 0x93, 0xad, 0x2e, 0xbb, 0xac, 0x74, 0x06, 0xd7, 0xce, 0xec, 0x85, 0xba, 0xda, 0xde,
	0xf1, 0xd7, 0x1f, 0xbf, 0xbe, 0xef, 0xdd, 0x1d, 0xdd, 0x51, 0x0f, 0xe7, 0x93, 0x72, 0xf8, 0x69,
	0xfd, 0x94, 0x46, 0x9f, 0x09, 0x83, 0x83, 0xcd, 0x4d, 0x89, 0xbb, 0xd5, 0x69, 0xc7, 0x39, 0x38,
	0x27, 0x7f, 0x51, 0x18, 0x9b, 0xbc, 0xbe, 0xee, 0x79, 0x83, 0x1c, 0x6c, 0x3e, 0xd6, 0x67, 0x6f,
	0xde, 0xbe, 0x5e, 0x32, 0xb9, 0x2a, 0x66, 0x7e, 0xc4, 0x93, 0xe0, 0x5c, 0x4f, 0x7b, 0x16, 0xf3,
	0x62, 0x7e, 0x11, 0x53, 0xb9, 0xe0, 0x79, 0x12, 0xac, 0x90, 0xc6, 0x72, 0x15, 0xd1, 0x1c, 0xc7,
	0x0b, 0x9c, 0x63, 0x4e, 0x25, 0xce, 0xc7, 0x34, 0x8a, 0x50, 0x88, 0xb1, 0xc0, 0x7c, 0xcd, 0x22,
	0x14, 0xc1, 0xb5, 0xdf, 0xc1, 0xac, 0xad, 0xa9, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x94, 0x5b, 0xd0, 0x2a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsentsClient is the client API for Consents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsentsClient interface {
	// Deletes the specified OAuth Consent.
	DeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists the OAuth Consents.
	ListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*ListConsentsResponse, error)
}

type consentsClient struct {
	cc *grpc.ClientConn
}

func NewConsentsClient(cc *grpc.ClientConn) ConsentsClient {
	return &consentsClient{cc}
}

func (c *consentsClient) DeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/DeleteConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentsClient) ListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*ListConsentsResponse, error) {
	out := new(ListConsentsResponse)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/ListConsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsentsServer is the server API for Consents service.
type ConsentsServer interface {
	// Deletes the specified OAuth Consent.
	DeleteConsent(context.Context, *DeleteConsentRequest) (*empty.Empty, error)
	// Lists the OAuth Consents.
	ListConsents(context.Context, *ListConsentsRequest) (*ListConsentsResponse, error)
}

// UnimplementedConsentsServer can be embedded to have forward compatible implementations.
type UnimplementedConsentsServer struct {
}

func (*UnimplementedConsentsServer) DeleteConsent(ctx context.Context, req *DeleteConsentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsent not implemented")
}
func (*UnimplementedConsentsServer) ListConsents(ctx context.Context, req *ListConsentsRequest) (*ListConsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConsents not implemented")
}

func RegisterConsentsServer(s *grpc.Server, srv ConsentsServer) {
	s.RegisterService(&_Consents_serviceDesc, srv)
}

func _Consents_DeleteConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).DeleteConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/DeleteConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).DeleteConsent(ctx, req.(*DeleteConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consents_ListConsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConsentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).ListConsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/ListConsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).ListConsents(ctx, req.(*ListConsentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Consents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consents.v1.Consents",
	HandlerType: (*ConsentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteConsent",
			Handler:    _Consents_DeleteConsent_Handler,
		},
		{
			MethodName: "ListConsents",
			Handler:    _Consents_ListConsents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consents/v1/consents.proto",
}