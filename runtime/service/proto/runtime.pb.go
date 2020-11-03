// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime.proto

package stack_rpc_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Service struct {
	// name of the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// version of the service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// git url of the source
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// service metadata
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateOptions struct {
	// command to pass in
	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
	// environment to pass in
	Env []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// output to send to
	Output               string   `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOptions) Reset()         { *m = CreateOptions{} }
func (m *CreateOptions) String() string { return proto.CompactTextString(m) }
func (*CreateOptions) ProtoMessage()    {}
func (*CreateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{1}
}

func (m *CreateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOptions.Unmarshal(m, b)
}
func (m *CreateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOptions.Marshal(b, m, deterministic)
}
func (m *CreateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOptions.Merge(m, src)
}
func (m *CreateOptions) XXX_Size() int {
	return xxx_messageInfo_CreateOptions.Size(m)
}
func (m *CreateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOptions proto.InternalMessageInfo

func (m *CreateOptions) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CreateOptions) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *CreateOptions) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type CreateRequest struct {
	Service              *Service       `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Options              *CreateOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateRequest) GetOptions() *CreateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type ReadOptions struct {
	// service name
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// version of the service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// type of service
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{4}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

func (m *ReadOptions) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ReadOptions) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ReadOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ReadRequest struct {
	Options              *ReadOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{5}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReadResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{6}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{9}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{10}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{11}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{12}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "stack.rpc.runtime.Service")
	proto.RegisterMapType((map[string]string)(nil), "stack.rpc.runtime.Service.MetadataEntry")
	proto.RegisterType((*CreateOptions)(nil), "stack.rpc.runtime.CreateOptions")
	proto.RegisterType((*CreateRequest)(nil), "stack.rpc.runtime.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "stack.rpc.runtime.CreateResponse")
	proto.RegisterType((*ReadOptions)(nil), "stack.rpc.runtime.ReadOptions")
	proto.RegisterType((*ReadRequest)(nil), "stack.rpc.runtime.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "stack.rpc.runtime.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "stack.rpc.runtime.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "stack.rpc.runtime.DeleteResponse")
	proto.RegisterType((*UpdateRequest)(nil), "stack.rpc.runtime.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "stack.rpc.runtime.UpdateResponse")
	proto.RegisterType((*ListRequest)(nil), "stack.rpc.runtime.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "stack.rpc.runtime.ListResponse")
}

func init() { proto.RegisterFile("runtime.proto", fileDescriptor_86e2dd377c869464) }

var fileDescriptor_86e2dd377c869464 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0x38, 0xc4, 0xed, 0x18, 0xa3, 0x32, 0x42, 0xc8, 0xca, 0x01, 0x82, 0x4f, 0x39,
	0xf9, 0x10, 0x10, 0xaa, 0xca, 0x91, 0x06, 0x84, 0x04, 0x8a, 0xe4, 0xaa, 0x0f, 0xb0, 0x38, 0x73,
	0xb0, 0x5a, 0x7b, 0x8d, 0x77, 0x1d, 0x29, 0x47, 0x5e, 0x81, 0x47, 0xe3, 0x89, 0xd0, 0xfe, 0x4b,
	0xec, 0x62, 0xf7, 0xd2, 0xde, 0x76, 0x3c, 0xbb, 0xdf, 0xfe, 0xbe, 0xf9, 0x56, 0x86, 0xa8, 0x69,
	0x2b, 0x59, 0x94, 0x94, 0xd6, 0x0d, 0x97, 0x1c, 0x5f, 0x0a, 0xc9, 0xf2, 0xdb, 0xb4, 0xa9, 0xf3,
	0xd4, 0x36, 0x92, 0xbf, 0x1e, 0x04, 0xd7, 0xd4, 0xec, 0x8a, 0x9c, 0x10, 0x61, 0x5a, 0xb1, 0x92,
	0x62, 0x6f, 0xe1, 0x2d, 0xcf, 0x32, 0xbd, 0xc6, 0x18, 0x82, 0x1d, 0x35, 0xa2, 0xe0, 0x55, 0x3c,
	0xd1, 0x9f, 0x5d, 0x89, 0xaf, 0x61, 0x26, 0x78, 0xdb, 0xe4, 0x14, 0xfb, 0xba, 0x61, 0x2b, 0xbc,
	0x82, 0xd3, 0x92, 0x24, 0xdb, 0x32, 0xc9, 0xe2, 0xe9, 0xc2, 0x5f, 0x86, 0xab, 0x65, 0xfa, 0xdf,
	0xbd, 0xa9, 0xbd, 0x33, 0xfd, 0x61, 0xb7, 0xae, 0x2b, 0xd9, 0xec, 0xb3, 0xc3, 0xc9, 0xf9, 0x27,
	0x88, 0x7a, 0x2d, 0x3c, 0x07, 0xff, 0x96, 0xf6, 0x96, 0x4d, 0x2d, 0xf1, 0x15, 0x3c, 0xdb, 0xb1,
	0xbb, 0x96, 0x2c, 0x98, 0x29, 0x2e, 0x27, 0x17, 0x5e, 0x72, 0x0d, 0xd1, 0xe7, 0x86, 0x98, 0xa4,
	0x4d, 0x2d, 0x0b, 0x5e, 0x09, 0xe5, 0x22, 0xe7, 0x65, 0xc9, 0xaa, 0x6d, 0xec, 0x2d, 0x7c, 0xe5,
	0xc2, 0x96, 0x4a, 0x96, 0xaa, 0x5d, 0x3c, 0xd1, 0x5f, 0xd5, 0x52, 0xf9, 0xe2, 0xad, 0xac, 0x5b,
	0xe9, 0x7c, 0x99, 0x2a, 0xf9, 0xed, 0x39, 0xd5, 0x8c, 0x7e, 0xb5, 0x24, 0x24, 0x7e, 0x80, 0x40,
	0x18, 0x1b, 0x1a, 0x2b, 0x5c, 0xcd, 0xc7, 0x8d, 0x66, 0x6e, 0x2b, 0x5e, 0x42, 0xc0, 0x0d, 0x96,
	0x06, 0x0f, 0x57, 0x8b, 0x81, 0x53, 0x3d, 0xfc, 0xcc, 0x1d, 0x48, 0xce, 0xe1, 0x85, 0x43, 0x10,
	0x35, 0xaf, 0x04, 0x25, 0x37, 0x10, 0x66, 0xc4, 0xb6, 0x1d, 0xa3, 0x5d, 0xa4, 0xb3, 0xe3, 0xb5,
	0xe3, 0x41, 0x22, 0x4c, 0xe5, 0xbe, 0x76, 0x31, 0xea, 0x75, 0xf2, 0xd5, 0xc8, 0x3a, 0xa7, 0x17,
	0x47, 0x66, 0xe3, 0xf4, 0xcd, 0x00, 0x73, 0x87, 0xe3, 0x48, 0xfc, 0x05, 0x9e, 0x1b, 0x21, 0xc3,
	0x8b, 0x1f, 0xe1, 0xd4, 0x12, 0x09, 0x1d, 0xc5, 0xc3, 0x43, 0x3b, 0xec, 0x4d, 0xd6, 0x10, 0x5d,
	0xd1, 0x1d, 0x3d, 0x72, 0xf8, 0x6a, 0x80, 0x4e, 0xc6, 0x0e, 0x70, 0x0d, 0xd1, 0x4d, 0xbd, 0x65,
	0x4f, 0x20, 0xec, 0x64, 0xac, 0x70, 0x04, 0xe1, 0xf7, 0x42, 0x48, 0x2b, 0xab, 0x06, 0x61, 0xca,
	0xc7, 0x0d, 0x62, 0xf5, 0xc7, 0x87, 0x20, 0x33, 0x5d, 0xdc, 0xc0, 0xcc, 0x3c, 0x07, 0x1c, 0x7f,
	0x43, 0xf6, 0xfe, 0xf9, 0xbb, 0x07, 0x76, 0x58, 0xe2, 0x13, 0xfc, 0x06, 0x53, 0x95, 0x16, 0x8e,
	0xc5, 0xeb, 0xc4, 0xde, 0x8e, 0xf6, 0x0f, 0x52, 0x1b, 0x98, 0x99, 0x49, 0x0f, 0xb2, 0xf5, 0xb2,
	0x1c, 0x64, 0xbb, 0x17, 0x93, 0x16, 0x34, 0x13, 0x1e, 0x14, 0xec, 0x65, 0x38, 0x28, 0x78, 0x2f,
	0x1e, 0x6d, 0x56, 0x25, 0x32, 0x68, 0xb6, 0x93, 0xdc, 0xa0, 0xd9, 0x6e, 0x94, 0xc9, 0xc9, 0xcf,
	0x99, 0xfe, 0xbf, 0xbe, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x62, 0xdb, 0x01, 0xf3, 0x70, 0x05,
	0x00, 0x00,
}
