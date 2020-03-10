// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/click_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
// Queries including ClickView must have a filter limiting the results to one
// day and can be requested for dates back to 90 days before the time of the
// request.
type ClickView struct {
	// The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Google Click ID.
	Gclid *wrappers.StringValue `protobuf:"bytes,2,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Page number in search results where the ad was shown.
	PageNumber *wrappers.Int64Value `protobuf:"bytes,5,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// The associated ad.
	AdGroupAd            *wrappers.StringValue `protobuf:"bytes,7,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickView) Reset()         { *m = ClickView{} }
func (m *ClickView) String() string { return proto.CompactTextString(m) }
func (*ClickView) ProtoMessage()    {}
func (*ClickView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eb15a7d43c843e9, []int{0}
}

func (m *ClickView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickView.Unmarshal(m, b)
}
func (m *ClickView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickView.Marshal(b, m, deterministic)
}
func (m *ClickView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickView.Merge(m, src)
}
func (m *ClickView) XXX_Size() int {
	return xxx_messageInfo_ClickView.Size(m)
}
func (m *ClickView) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickView.DiscardUnknown(m)
}

var xxx_messageInfo_ClickView proto.InternalMessageInfo

func (m *ClickView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ClickView) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if m != nil {
		return m.AreaOfInterest
	}
	return nil
}

func (m *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if m != nil {
		return m.LocationOfPresence
	}
	return nil
}

func (m *ClickView) GetPageNumber() *wrappers.Int64Value {
	if m != nil {
		return m.PageNumber
	}
	return nil
}

func (m *ClickView) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickView)(nil), "google.ads.googleads.v3.resources.ClickView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/click_view.proto", fileDescriptor_7eb15a7d43c843e9)
}

var fileDescriptor_7eb15a7d43c843e9 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0xae, 0x55, 0x76, 0xaa, 0x45, 0x82, 0x87, 0x58, 0x8b, 0xb4, 0x95, 0x42, 0x0f,
	0x3a, 0x81, 0x8d, 0x78, 0x88, 0x5e, 0x52, 0x0f, 0x4b, 0x45, 0xda, 0x25, 0x62, 0x04, 0x59, 0x08,
	0xb3, 0x99, 0xb7, 0x43, 0x30, 0x99, 0x09, 0x33, 0xc9, 0xee, 0x41, 0xfa, 0x65, 0x3c, 0xfa, 0x51,
	0xfc, 0x28, 0xbd, 0x79, 0xf6, 0x22, 0x93, 0x99, 0xcc, 0x0a, 0x52, 0x15, 0x6f, 0x6f, 0xe7, 0xfd,
	0xff, 0xbf, 0xf9, 0xbf, 0xb7, 0x13, 0x34, 0x65, 0x42, 0xb0, 0x0a, 0x42, 0x42, 0x55, 0x68, 0x4a,
	0x5d, 0xad, 0xa3, 0x50, 0x82, 0x12, 0x9d, 0x2c, 0x40, 0x85, 0x45, 0x55, 0x16, 0x9f, 0xf2, 0x75,
	0x09, 0x1b, 0xdc, 0x48, 0xd1, 0x0a, 0xff, 0xc8, 0x08, 0x31, 0xa1, 0x0a, 0x3b, 0x0f, 0x5e, 0x47,
	0xd8, 0x79, 0xf6, 0xa3, 0x9b, 0xb0, 0x85, 0xa8, 0x6b, 0xc1, 0x2d, 0xb3, 0x12, 0x05, 0x69, 0x4b,
	0xc1, 0x0d, 0x77, 0xff, 0xe1, 0x60, 0x6a, 0x4a, 0x77, 0xbd, 0x6d, 0x3d, 0xb6, 0xad, 0xfe, 0xd7,
	0xb2, 0x5b, 0x85, 0x1b, 0x49, 0x9a, 0x06, 0xa4, 0xb2, 0xfd, 0x83, 0x5f, 0xac, 0x84, 0x73, 0xd1,
	0xf6, 0x5c, 0xdb, 0x3d, 0xfe, 0x3e, 0x46, 0x93, 0xd7, 0xfa, 0xc6, 0xac, 0x84, 0x8d, 0xff, 0x04,
	0xdd, 0x1b, 0xe8, 0x39, 0x27, 0x35, 0x04, 0xde, 0xa1, 0x77, 0x3a, 0x49, 0xef, 0x0e, 0x87, 0x17,
	0xa4, 0x06, 0x7f, 0x8a, 0x76, 0x58, 0x51, 0x95, 0x34, 0x18, 0x1d, 0x7a, 0xa7, 0xbb, 0xd3, 0x03,
	0x3b, 0x28, 0x1e, 0x02, 0xe0, 0x77, 0xad, 0x2c, 0x39, 0xcb, 0x48, 0xd5, 0x41, 0x6a, 0xa4, 0xfe,
	0x07, 0x74, 0x9f, 0x48, 0x20, 0xb9, 0x58, 0xe5, 0x25, 0x6f, 0x41, 0x82, 0x6a, 0x83, 0x71, 0x6f,
	0x7f, 0x86, 0x6f, 0x5a, 0x99, 0xd9, 0x07, 0xee, 0xd3, 0xbd, 0xb5, 0xeb, 0x48, 0xf7, 0x34, 0xe6,
	0x72, 0x75, 0x6e, 0x21, 0x7e, 0x8e, 0x1e, 0x0c, 0xab, 0xd2, 0xf0, 0x46, 0x82, 0x02, 0x5e, 0x40,
	0x70, 0xeb, 0x7f, 0xe0, 0xfe, 0x80, 0xba, 0x5c, 0xcd, 0x2d, 0xc8, 0x7f, 0x85, 0x76, 0x1b, 0xc2,
	0x20, 0xe7, 0x5d, 0xbd, 0x04, 0x19, 0xec, 0xf4, 0xdc, 0x47, 0xbf, 0xcd, 0x7c, 0xce, 0xdb, 0x17,
	0xcf, 0xcd, 0xc8, 0x48, 0xeb, 0x2f, 0x7a, 0xb9, 0x76, 0x13, 0x9a, 0x33, 0x29, 0xba, 0x26, 0x27,
	0x34, 0xb8, 0xf3, 0x0f, 0x1b, 0x9b, 0x10, 0x3a, 0xd3, 0xfa, 0x84, 0xc6, 0xef, 0xaf, 0x93, 0x14,
	0x1d, 0x6f, 0x73, 0xdb, 0xaa, 0x29, 0x95, 0xce, 0x1f, 0x6e, 0xff, 0xb7, 0xa7, 0x45, 0xa7, 0x5a,
	0x51, 0x83, 0x54, 0xe1, 0xe7, 0xa1, 0xbc, 0x32, 0x2f, 0x49, 0xf7, 0xf5, 0xa9, 0x7b, 0xa9, 0x57,
	0x67, 0x3f, 0x3c, 0x74, 0x52, 0x88, 0x1a, 0xff, 0xf5, 0xad, 0x9e, 0xed, 0xb9, 0x2b, 0xe6, 0x3a,
	0xea, 0xdc, 0xfb, 0xf8, 0xc6, 0x9a, 0x98, 0xa8, 0x08, 0x67, 0x58, 0x48, 0x16, 0x32, 0xe0, 0xfd,
	0x20, 0xe1, 0x36, 0xde, 0x1f, 0xbe, 0x99, 0x97, 0xae, 0xfa, 0x32, 0x1a, 0xcf, 0x92, 0xe4, 0xeb,
	0xe8, 0x68, 0x66, 0x90, 0x09, 0x55, 0xd8, 0x94, 0xba, 0xca, 0x22, 0x9c, 0x0e, 0xca, 0x6f, 0x83,
	0x66, 0x91, 0x50, 0xb5, 0x70, 0x9a, 0x45, 0x16, 0x2d, 0x9c, 0xe6, 0x7a, 0x74, 0x62, 0x1a, 0x71,
	0x9c, 0x50, 0x15, 0xc7, 0x4e, 0x15, 0xc7, 0x59, 0x14, 0xc7, 0x4e, 0xb7, 0xbc, 0xdd, 0x87, 0x8d,
	0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x56, 0xa3, 0xdc, 0xdf, 0x03, 0x00, 0x00,
}