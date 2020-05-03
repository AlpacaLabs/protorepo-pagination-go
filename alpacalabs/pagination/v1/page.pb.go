// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/pagination/v1/page.proto

package v1

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

type Sort int32

const (
	Sort_SORT_UNSPECIFIED Sort = 0
	Sort_SORT_ASC         Sort = 1
	Sort_SORT_DESC        Sort = 2
)

var Sort_name = map[int32]string{
	0: "SORT_UNSPECIFIED",
	1: "SORT_ASC",
	2: "SORT_DESC",
}

var Sort_value = map[string]int32{
	"SORT_UNSPECIFIED": 0,
	"SORT_ASC":         1,
	"SORT_DESC":        2,
}

func (x Sort) String() string {
	return proto.EnumName(Sort_name, int32(x))
}

func (Sort) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89006c53d509bbb7, []int{0}
}

type SortClause struct {
	FieldName            string   `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Sort                 Sort     `protobuf:"varint,2,opt,name=sort,proto3,enum=alpacalabs.pagination.v1.Sort" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortClause) Reset()         { *m = SortClause{} }
func (m *SortClause) String() string { return proto.CompactTextString(m) }
func (*SortClause) ProtoMessage()    {}
func (*SortClause) Descriptor() ([]byte, []int) {
	return fileDescriptor_89006c53d509bbb7, []int{0}
}

func (m *SortClause) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortClause.Unmarshal(m, b)
}
func (m *SortClause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortClause.Marshal(b, m, deterministic)
}
func (m *SortClause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortClause.Merge(m, src)
}
func (m *SortClause) XXX_Size() int {
	return xxx_messageInfo_SortClause.Size(m)
}
func (m *SortClause) XXX_DiscardUnknown() {
	xxx_messageInfo_SortClause.DiscardUnknown(m)
}

var xxx_messageInfo_SortClause proto.InternalMessageInfo

func (m *SortClause) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *SortClause) GetSort() Sort {
	if m != nil {
		return m.Sort
	}
	return Sort_SORT_UNSPECIFIED
}

// CursorRequest provides metadata for cursor pagination.
// See https://developer.twitter.com/en/docs/basics/cursoring.
type CursorRequest struct {
	// Cursor is inclusive.
	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Count is how many elements the client wishes to fetch.
	// Endpoints should set a maximum number of elements that can be returned.
	Count                int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	SortClauses          []*SortClause `protobuf:"bytes,3,rep,name=sort_clauses,json=sortClauses,proto3" json:"sort_clauses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CursorRequest) Reset()         { *m = CursorRequest{} }
func (m *CursorRequest) String() string { return proto.CompactTextString(m) }
func (*CursorRequest) ProtoMessage()    {}
func (*CursorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89006c53d509bbb7, []int{1}
}

func (m *CursorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorRequest.Unmarshal(m, b)
}
func (m *CursorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorRequest.Marshal(b, m, deterministic)
}
func (m *CursorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorRequest.Merge(m, src)
}
func (m *CursorRequest) XXX_Size() int {
	return xxx_messageInfo_CursorRequest.Size(m)
}
func (m *CursorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CursorRequest proto.InternalMessageInfo

func (m *CursorRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *CursorRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CursorRequest) GetSortClauses() []*SortClause {
	if m != nil {
		return m.SortClauses
	}
	return nil
}

type CursorResponse struct {
	// PreviousCursor is the cursor the client sent in their CursorRequest.
	PreviousCursor string `protobuf:"bytes,1,opt,name=previous_cursor,json=previousCursor,proto3" json:"previous_cursor,omitempty"`
	// NextCursor is the cursor clients should use to request the next page.
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// Count is the page size, the number of elements in the page.
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CursorResponse) Reset()         { *m = CursorResponse{} }
func (m *CursorResponse) String() string { return proto.CompactTextString(m) }
func (*CursorResponse) ProtoMessage()    {}
func (*CursorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89006c53d509bbb7, []int{2}
}

func (m *CursorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorResponse.Unmarshal(m, b)
}
func (m *CursorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorResponse.Marshal(b, m, deterministic)
}
func (m *CursorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorResponse.Merge(m, src)
}
func (m *CursorResponse) XXX_Size() int {
	return xxx_messageInfo_CursorResponse.Size(m)
}
func (m *CursorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CursorResponse proto.InternalMessageInfo

func (m *CursorResponse) GetPreviousCursor() string {
	if m != nil {
		return m.PreviousCursor
	}
	return ""
}

func (m *CursorResponse) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *CursorResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("alpacalabs.pagination.v1.Sort", Sort_name, Sort_value)
	proto.RegisterType((*SortClause)(nil), "alpacalabs.pagination.v1.SortClause")
	proto.RegisterType((*CursorRequest)(nil), "alpacalabs.pagination.v1.CursorRequest")
	proto.RegisterType((*CursorResponse)(nil), "alpacalabs.pagination.v1.CursorResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/pagination/v1/page.proto", fileDescriptor_89006c53d509bbb7)
}

var fileDescriptor_89006c53d509bbb7 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x8a, 0xda, 0x40,
	0x18, 0x6d, 0x12, 0x95, 0xfa, 0xf9, 0x17, 0x06, 0x29, 0xb9, 0xe8, 0x8f, 0xd8, 0x42, 0xa5, 0xe0,
	0x84, 0xd8, 0xab, 0x5e, 0x6a, 0xd4, 0x22, 0x88, 0x0d, 0x49, 0xeb, 0x45, 0x11, 0xc2, 0x98, 0x4e,
	0xdd, 0x40, 0x92, 0xc9, 0x66, 0x12, 0xd9, 0x27, 0xd8, 0x07, 0xd9, 0xcb, 0xbd, 0xdf, 0x97, 0xd8,
	0xa7, 0x5a, 0x32, 0x51, 0xe3, 0x2e, 0xb8, 0x77, 0x39, 0xe7, 0x3b, 0x73, 0xce, 0xf9, 0xc8, 0x07,
	0x9f, 0x49, 0x10, 0x13, 0x8f, 0x04, 0x64, 0xcb, 0xf5, 0x98, 0xec, 0xfc, 0x88, 0xa4, 0x3e, 0x8b,
	0xf4, 0xbd, 0x91, 0x23, 0x8a, 0xe3, 0x84, 0xa5, 0x0c, 0x69, 0xa5, 0x08, 0x97, 0x22, 0xbc, 0x37,
	0xfa, 0x2e, 0x80, 0xc3, 0x92, 0xd4, 0x0c, 0x48, 0xc6, 0x29, 0xfa, 0x00, 0xf0, 0xdf, 0xa7, 0xc1,
	0x3f, 0x37, 0x22, 0x21, 0xd5, 0xa4, 0x9e, 0x34, 0xa8, 0xdb, 0x75, 0xc1, 0xac, 0x48, 0x48, 0xd1,
	0x08, 0x2a, 0x9c, 0x25, 0xa9, 0x26, 0xf7, 0xa4, 0x41, 0x7b, 0xf4, 0x11, 0x5f, 0x72, 0xc5, 0xb9,
	0xa5, 0x2d, 0xb4, 0xfd, 0x5b, 0x09, 0x5a, 0x66, 0x96, 0x70, 0x96, 0xd8, 0xf4, 0x3a, 0xa3, 0x3c,
	0x45, 0xef, 0xa0, 0xe6, 0x09, 0xe2, 0x10, 0x70, 0x40, 0xa8, 0x0b, 0x55, 0x8f, 0x65, 0x51, 0x61,
	0x5f, 0xb5, 0x0b, 0x80, 0x7e, 0x42, 0x33, 0xf7, 0x71, 0x3d, 0xd1, 0x90, 0x6b, 0x4a, 0x4f, 0x19,
	0x34, 0x46, 0x5f, 0x5e, 0xcf, 0x2e, 0xd6, 0xb1, 0x1b, 0xfc, 0xf4, 0xcd, 0xfb, 0x31, 0xb4, 0x8f,
	0x3d, 0x78, 0xcc, 0x22, 0x4e, 0xd1, 0x57, 0xe8, 0xc4, 0x09, 0xdd, 0xfb, 0x2c, 0xe3, 0xee, 0xb3,
	0x46, 0xed, 0x23, 0x5d, 0x3c, 0x40, 0x9f, 0xa0, 0x11, 0xd1, 0x9b, 0xf4, 0x28, 0x92, 0x85, 0x08,
	0x72, 0xca, 0x7c, 0x51, 0x5d, 0x39, 0xab, 0xfe, 0xed, 0x07, 0x54, 0xf2, 0x32, 0xa8, 0x0b, 0xaa,
	0xf3, 0xcb, 0xfe, 0xed, 0xfe, 0x59, 0x39, 0xd6, 0xcc, 0x5c, 0xcc, 0x17, 0xb3, 0xa9, 0xfa, 0x06,
	0x35, 0xe1, 0xad, 0x60, 0xc7, 0x8e, 0xa9, 0x4a, 0xa8, 0x05, 0x75, 0x81, 0xa6, 0x33, 0xc7, 0x54,
	0xe5, 0xc9, 0x83, 0x04, 0xef, 0x3d, 0x16, 0x5e, 0xdc, 0x72, 0xd2, 0xb1, 0x4e, 0xd0, 0xca, 0x7f,
	0xb1, 0x25, 0xfd, 0x9d, 0xef, 0xfc, 0xf4, 0x2a, 0xdb, 0x62, 0x8f, 0x85, 0xfa, 0x58, 0xbc, 0x5b,
	0x8a, 0xa3, 0xc8, 0xc7, 0x09, 0x8d, 0xd9, 0xb0, 0x74, 0x18, 0xee, 0x98, 0x7e, 0xe9, 0x6e, 0xee,
	0x64, 0x65, 0xbc, 0xb4, 0xee, 0x65, 0xad, 0xf4, 0xc0, 0x65, 0x18, 0x5e, 0x1b, 0x8f, 0xe7, 0xa3,
	0x4d, 0x39, 0xda, 0xac, 0x8d, 0x6d, 0x4d, 0xa4, 0x7d, 0x7f, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xb2, 0x57, 0x5d, 0x96, 0x02, 0x00, 0x00,
}
