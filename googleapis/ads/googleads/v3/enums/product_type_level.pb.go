// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/product_type_level.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing the level of the type of a product offer.
type ProductTypeLevelEnum_ProductTypeLevel int32

const (
	// Not specified.
	ProductTypeLevelEnum_UNSPECIFIED ProductTypeLevelEnum_ProductTypeLevel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductTypeLevelEnum_UNKNOWN ProductTypeLevelEnum_ProductTypeLevel = 1
	// Level 1.
	ProductTypeLevelEnum_LEVEL1 ProductTypeLevelEnum_ProductTypeLevel = 7
	// Level 2.
	ProductTypeLevelEnum_LEVEL2 ProductTypeLevelEnum_ProductTypeLevel = 8
	// Level 3.
	ProductTypeLevelEnum_LEVEL3 ProductTypeLevelEnum_ProductTypeLevel = 9
	// Level 4.
	ProductTypeLevelEnum_LEVEL4 ProductTypeLevelEnum_ProductTypeLevel = 10
	// Level 5.
	ProductTypeLevelEnum_LEVEL5 ProductTypeLevelEnum_ProductTypeLevel = 11
)

var ProductTypeLevelEnum_ProductTypeLevel_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	7:  "LEVEL1",
	8:  "LEVEL2",
	9:  "LEVEL3",
	10: "LEVEL4",
	11: "LEVEL5",
}

var ProductTypeLevelEnum_ProductTypeLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LEVEL1":      7,
	"LEVEL2":      8,
	"LEVEL3":      9,
	"LEVEL4":      10,
	"LEVEL5":      11,
}

func (x ProductTypeLevelEnum_ProductTypeLevel) String() string {
	return proto.EnumName(ProductTypeLevelEnum_ProductTypeLevel_name, int32(x))
}

func (ProductTypeLevelEnum_ProductTypeLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b23b5c7ad7feb34b, []int{0, 0}
}

// Level of the type of a product offer.
type ProductTypeLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductTypeLevelEnum) Reset()         { *m = ProductTypeLevelEnum{} }
func (m *ProductTypeLevelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductTypeLevelEnum) ProtoMessage()    {}
func (*ProductTypeLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b23b5c7ad7feb34b, []int{0}
}

func (m *ProductTypeLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductTypeLevelEnum.Unmarshal(m, b)
}
func (m *ProductTypeLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductTypeLevelEnum.Marshal(b, m, deterministic)
}
func (m *ProductTypeLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductTypeLevelEnum.Merge(m, src)
}
func (m *ProductTypeLevelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductTypeLevelEnum.Size(m)
}
func (m *ProductTypeLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductTypeLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductTypeLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ProductTypeLevelEnum_ProductTypeLevel", ProductTypeLevelEnum_ProductTypeLevel_name, ProductTypeLevelEnum_ProductTypeLevel_value)
	proto.RegisterType((*ProductTypeLevelEnum)(nil), "google.ads.googleads.v3.enums.ProductTypeLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/product_type_level.proto", fileDescriptor_b23b5c7ad7feb34b)
}

var fileDescriptor_b23b5c7ad7feb34b = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0xff, 0xe9, 0x0f, 0xad, 0xa6, 0x0b, 0xc3, 0xa0, 0x1b, 0xb1, 0x8b, 0xf6, 0x01, 0x12,
	0x34, 0xea, 0x22, 0xae, 0xa6, 0x3a, 0x96, 0x62, 0x19, 0x07, 0xb4, 0x23, 0xc8, 0x40, 0x19, 0x9b,
	0x10, 0x0a, 0xd3, 0x24, 0x34, 0xd3, 0x42, 0xf7, 0x3e, 0x89, 0x4b, 0x1f, 0xc5, 0x47, 0xe9, 0x53,
	0xc8, 0x24, 0x4e, 0x16, 0x05, 0xdd, 0x84, 0x8f, 0xdc, 0x7b, 0x0e, 0xf7, 0x1c, 0x70, 0x2d, 0x94,
	0x12, 0x25, 0xc7, 0x05, 0x33, 0xd8, 0x61, 0x4d, 0x1b, 0x82, 0xb9, 0x5c, 0x2f, 0x0d, 0xd6, 0x2b,
	0xc5, 0xd6, 0xf3, 0x6a, 0x56, 0x6d, 0x35, 0x9f, 0x95, 0x7c, 0xc3, 0x4b, 0xa4, 0x57, 0xaa, 0x52,
	0x61, 0xcf, 0x2d, 0xa3, 0x82, 0x19, 0xe4, 0x75, 0x68, 0x43, 0x90, 0xd5, 0x9d, 0x9e, 0x35, 0xb6,
	0x7a, 0x81, 0x0b, 0x29, 0x55, 0x55, 0x54, 0x0b, 0x25, 0x8d, 0x13, 0x0f, 0xde, 0x03, 0x70, 0x9c,
	0x3a, 0xe7, 0xe7, 0xad, 0xe6, 0x93, 0xda, 0x37, 0x96, 0xeb, 0xe5, 0xa0, 0x04, 0x70, 0xff, 0x3f,
	0x3c, 0x02, 0xdd, 0x69, 0xf2, 0x94, 0xc6, 0xb7, 0xe3, 0xfb, 0x71, 0x7c, 0x07, 0xff, 0x85, 0x5d,
	0xd0, 0x99, 0x26, 0x0f, 0xc9, 0xe3, 0x4b, 0x02, 0x83, 0x10, 0x80, 0xf6, 0x24, 0xce, 0xe2, 0xc9,
	0x39, 0xec, 0x78, 0xbe, 0x80, 0x07, 0x9e, 0x09, 0x3c, 0xf4, 0x7c, 0x09, 0x81, 0xe7, 0x2b, 0xd8,
	0x1d, 0xee, 0x02, 0xd0, 0x9f, 0xab, 0x25, 0xfa, 0x33, 0xca, 0xf0, 0x64, 0xff, 0xa2, 0xb4, 0xce,
	0x90, 0x06, 0xaf, 0xc3, 0x1f, 0x9d, 0x50, 0x65, 0x21, 0x05, 0x52, 0x2b, 0x81, 0x05, 0x97, 0x36,
	0x61, 0x53, 0xa5, 0x5e, 0x98, 0x5f, 0x9a, 0xbd, 0xb1, 0xef, 0x47, 0xeb, 0xff, 0x28, 0x8a, 0x3e,
	0x5b, 0xbd, 0x91, 0xb3, 0x8a, 0x98, 0x41, 0x0e, 0x6b, 0xca, 0x08, 0xaa, 0x5b, 0x31, 0x5f, 0xcd,
	0x3c, 0x8f, 0x98, 0xc9, 0xfd, 0x3c, 0xcf, 0x48, 0x6e, 0xe7, 0xbb, 0x56, 0xdf, 0x7d, 0x52, 0x1a,
	0x31, 0x43, 0xa9, 0xdf, 0xa0, 0x34, 0x23, 0x94, 0xda, 0x9d, 0xb7, 0xb6, 0x3d, 0x8c, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0xca, 0x80, 0x46, 0xf1, 0x01, 0x00, 0x00,
}