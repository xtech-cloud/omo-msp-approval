// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/approval/workflow.proto

package approval

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

// 创建的请求
type WorkflowMakeRequest struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mode                 WorkflowMode `protobuf:"varint,2,opt,name=mode,proto3,enum=approval.WorkflowMode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WorkflowMakeRequest) Reset()         { *m = WorkflowMakeRequest{} }
func (m *WorkflowMakeRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowMakeRequest) ProtoMessage()    {}
func (*WorkflowMakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{0}
}

func (m *WorkflowMakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMakeRequest.Unmarshal(m, b)
}
func (m *WorkflowMakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMakeRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowMakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMakeRequest.Merge(m, src)
}
func (m *WorkflowMakeRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowMakeRequest.Size(m)
}
func (m *WorkflowMakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMakeRequest proto.InternalMessageInfo

func (m *WorkflowMakeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowMakeRequest) GetMode() WorkflowMode {
	if m != nil {
		return m.Mode
	}
	return WorkflowMode_WORKFLOW_MODE_INVALID
}

// 列举的请求
type WorkflowListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowListRequest) Reset()         { *m = WorkflowListRequest{} }
func (m *WorkflowListRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowListRequest) ProtoMessage()    {}
func (*WorkflowListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{1}
}

func (m *WorkflowListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowListRequest.Unmarshal(m, b)
}
func (m *WorkflowListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowListRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowListRequest.Merge(m, src)
}
func (m *WorkflowListRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowListRequest.Size(m)
}
func (m *WorkflowListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowListRequest proto.InternalMessageInfo

func (m *WorkflowListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *WorkflowListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 列举的回复
type WorkflowListResponse struct {
	Status               *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*WorkflowEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkflowListResponse) Reset()         { *m = WorkflowListResponse{} }
func (m *WorkflowListResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowListResponse) ProtoMessage()    {}
func (*WorkflowListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{2}
}

func (m *WorkflowListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowListResponse.Unmarshal(m, b)
}
func (m *WorkflowListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowListResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowListResponse.Merge(m, src)
}
func (m *WorkflowListResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowListResponse.Size(m)
}
func (m *WorkflowListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowListResponse proto.InternalMessageInfo

func (m *WorkflowListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *WorkflowListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *WorkflowListResponse) GetEntity() []*WorkflowEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 删除的请求
type WorkflowRemoveRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowRemoveRequest) Reset()         { *m = WorkflowRemoveRequest{} }
func (m *WorkflowRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowRemoveRequest) ProtoMessage()    {}
func (*WorkflowRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{3}
}

func (m *WorkflowRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowRemoveRequest.Unmarshal(m, b)
}
func (m *WorkflowRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowRemoveRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowRemoveRequest.Merge(m, src)
}
func (m *WorkflowRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowRemoveRequest.Size(m)
}
func (m *WorkflowRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowRemoveRequest proto.InternalMessageInfo

func (m *WorkflowRemoveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取的请求
type WorkflowGetRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowGetRequest) Reset()         { *m = WorkflowGetRequest{} }
func (m *WorkflowGetRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowGetRequest) ProtoMessage()    {}
func (*WorkflowGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{4}
}

func (m *WorkflowGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowGetRequest.Unmarshal(m, b)
}
func (m *WorkflowGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowGetRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowGetRequest.Merge(m, src)
}
func (m *WorkflowGetRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowGetRequest.Size(m)
}
func (m *WorkflowGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowGetRequest proto.InternalMessageInfo

func (m *WorkflowGetRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *WorkflowGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 获取的回复
type WorkflowGetResponse struct {
	Status               *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *WorkflowEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             []string        `protobuf:"bytes,3,rep,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WorkflowGetResponse) Reset()         { *m = WorkflowGetResponse{} }
func (m *WorkflowGetResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowGetResponse) ProtoMessage()    {}
func (*WorkflowGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b1463671f157d, []int{5}
}

func (m *WorkflowGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowGetResponse.Unmarshal(m, b)
}
func (m *WorkflowGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowGetResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowGetResponse.Merge(m, src)
}
func (m *WorkflowGetResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowGetResponse.Size(m)
}
func (m *WorkflowGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowGetResponse proto.InternalMessageInfo

func (m *WorkflowGetResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *WorkflowGetResponse) GetEntity() *WorkflowEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *WorkflowGetResponse) GetOperator() []string {
	if m != nil {
		return m.Operator
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowMakeRequest)(nil), "approval.WorkflowMakeRequest")
	proto.RegisterType((*WorkflowListRequest)(nil), "approval.WorkflowListRequest")
	proto.RegisterType((*WorkflowListResponse)(nil), "approval.WorkflowListResponse")
	proto.RegisterType((*WorkflowRemoveRequest)(nil), "approval.WorkflowRemoveRequest")
	proto.RegisterType((*WorkflowGetRequest)(nil), "approval.WorkflowGetRequest")
	proto.RegisterType((*WorkflowGetResponse)(nil), "approval.WorkflowGetResponse")
}

func init() {
	proto.RegisterFile("proto/approval/workflow.proto", fileDescriptor_823b1463671f157d)
}

var fileDescriptor_823b1463671f157d = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x4e, 0xe3, 0x30,
	0x14, 0x85, 0x27, 0x3f, 0x13, 0xb5, 0xb7, 0xd2, 0x68, 0xe4, 0xe9, 0x94, 0x28, 0x50, 0xa8, 0xb2,
	0x8a, 0x40, 0x6a, 0x51, 0xd8, 0xb2, 0x40, 0x50, 0xd4, 0x0d, 0x6c, 0x8c, 0x10, 0x6b, 0x43, 0x5c,
	0x51, 0x35, 0x8d, 0x43, 0xec, 0xb4, 0xe2, 0x0d, 0xd8, 0xf0, 0x10, 0xbc, 0x29, 0x8a, 0x13, 0xc7,
	0x2d, 0xfd, 0x41, 0xec, 0x72, 0x7d, 0x8e, 0x3f, 0xf9, 0x9e, 0x7b, 0x03, 0xdd, 0x34, 0x63, 0x82,
	0x0d, 0x48, 0x9a, 0x66, 0x6c, 0x4e, 0xe2, 0xc1, 0x82, 0x65, 0xd3, 0x71, 0xcc, 0x16, 0x7d, 0x79,
	0x8e, 0x1a, 0x4a, 0xf0, 0xf6, 0xbf, 0x18, 0xf9, 0x33, 0xc9, 0x68, 0x54, 0xda, 0xfc, 0x7b, 0xf8,
	0xf7, 0x50, 0x5d, 0xbc, 0x25, 0x53, 0x8a, 0xe9, 0x4b, 0x4e, 0xb9, 0x40, 0x08, 0xec, 0x84, 0xcc,
	0xa8, 0x6b, 0xf4, 0x8c, 0xa0, 0x89, 0xe5, 0x37, 0x3a, 0x06, 0x7b, 0xc6, 0x22, 0xea, 0x9a, 0x3d,
	0x23, 0xf8, 0x13, 0x76, 0xfa, 0x0a, 0xd8, 0xaf, 0x01, 0x2c, 0xa2, 0x58, 0x7a, 0xfc, 0x2b, 0x8d,
	0xbd, 0x99, 0x70, 0xa1, 0xb0, 0x1d, 0x70, 0xd8, 0x78, 0xcc, 0xa9, 0x90, 0x60, 0x0b, 0x57, 0x15,
	0x6a, 0xc3, 0xef, 0x27, 0x96, 0x27, 0x42, 0xb2, 0x2d, 0x5c, 0x16, 0xfe, 0x9b, 0x01, 0xed, 0x55,
	0x0a, 0x4f, 0x59, 0xc2, 0x29, 0x0a, 0xc0, 0xe1, 0x82, 0x88, 0x9c, 0x4b, 0x4c, 0x2b, 0xfc, 0xab,
	0xdf, 0x72, 0x27, 0xcf, 0x71, 0xa5, 0x17, 0x60, 0xc1, 0x04, 0x89, 0x25, 0xd8, 0xc6, 0x65, 0x81,
	0x4e, 0xc1, 0xa1, 0x89, 0x98, 0x88, 0x57, 0xd7, 0xea, 0x59, 0x41, 0x2b, 0x74, 0xd7, 0x7b, 0xb9,
	0x96, 0x3a, 0xae, 0x7c, 0xfe, 0x09, 0xfc, 0x57, 0x0a, 0xa6, 0x33, 0x36, 0x5f, 0x0e, 0x2a, 0xcf,
	0x27, 0x91, 0x0a, 0xaa, 0xf8, 0xf6, 0xcf, 0x01, 0x29, 0xf3, 0x88, 0x8a, 0x1d, 0xce, 0x3a, 0x66,
	0x53, 0xc7, 0xec, 0xbf, 0x1b, 0x3a, 0x3b, 0x79, 0xfd, 0xc7, 0x4d, 0xeb, 0xf6, 0x4c, 0xe9, 0xfc,
	0xb6, 0x3d, 0xe4, 0x41, 0x83, 0xa5, 0x34, 0x23, 0x82, 0x65, 0x32, 0x92, 0x26, 0xae, 0xeb, 0xf0,
	0xc3, 0x84, 0x86, 0xba, 0x86, 0x2e, 0xc0, 0x2e, 0xd6, 0x04, 0x75, 0x37, 0x4c, 0x5f, 0xaf, 0x8f,
	0xb7, 0xa7, 0xe5, 0xcb, 0x98, 0x24, 0x53, 0xd5, 0x84, 0xff, 0x0b, 0x8d, 0xc0, 0x2e, 0x66, 0xb9,
	0x89, 0xb0, 0xb4, 0x29, 0xde, 0xe1, 0x36, 0xb9, 0x06, 0x0d, 0xc1, 0x29, 0x47, 0x81, 0x8e, 0xd6,
	0xbd, 0x2b, 0x43, 0xda, 0xf5, 0x9c, 0x21, 0x58, 0x23, 0x2a, 0xd0, 0xc1, 0x3a, 0x42, 0x8f, 0xce,
	0xeb, 0x6e, 0x51, 0x15, 0xe5, 0xd1, 0x91, 0x3f, 0xd3, 0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x98, 0x0f, 0xbd, 0x94, 0x03, 0x00, 0x00,
}
