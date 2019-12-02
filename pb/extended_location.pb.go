// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extended_location.proto

package pb

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

type ExtLocation struct {
	Primary              int32      `protobuf:"varint,1,opt,name=primary,proto3" json:"primary,omitempty"`
	Shipping             int32      `protobuf:"varint,2,opt,name=shipping,proto3" json:"shipping,omitempty"`
	Billing              int32      `protobuf:"varint,3,opt,name=billing,proto3" json:"billing,omitempty"`
	Return               int32      `protobuf:"varint,4,opt,name=return,proto3" json:"return,omitempty"`
	Addresses            []*Address `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExtLocation) Reset()         { *m = ExtLocation{} }
func (m *ExtLocation) String() string { return proto.CompactTextString(m) }
func (*ExtLocation) ProtoMessage()    {}
func (*ExtLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e5b30b37c9f3a59, []int{0}
}

func (m *ExtLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtLocation.Unmarshal(m, b)
}
func (m *ExtLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtLocation.Marshal(b, m, deterministic)
}
func (m *ExtLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtLocation.Merge(m, src)
}
func (m *ExtLocation) XXX_Size() int {
	return xxx_messageInfo_ExtLocation.Size(m)
}
func (m *ExtLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ExtLocation proto.InternalMessageInfo

func (m *ExtLocation) GetPrimary() int32 {
	if m != nil {
		return m.Primary
	}
	return 0
}

func (m *ExtLocation) GetShipping() int32 {
	if m != nil {
		return m.Shipping
	}
	return 0
}

func (m *ExtLocation) GetBilling() int32 {
	if m != nil {
		return m.Billing
	}
	return 0
}

func (m *ExtLocation) GetReturn() int32 {
	if m != nil {
		return m.Return
	}
	return 0
}

func (m *ExtLocation) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Address struct {
	Latitude             string   `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            string   `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PlusCode             string   `protobuf:"bytes,3,opt,name=plusCode,proto3" json:"plusCode,omitempty"`
	AddressOne           string   `protobuf:"bytes,4,opt,name=addressOne,proto3" json:"addressOne,omitempty"`
	AddressTwo           string   `protobuf:"bytes,5,opt,name=addressTwo,proto3" json:"addressTwo,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode              string   `protobuf:"bytes,9,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e5b30b37c9f3a59, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *Address) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *Address) GetPlusCode() string {
	if m != nil {
		return m.PlusCode
	}
	return ""
}

func (m *Address) GetAddressOne() string {
	if m != nil {
		return m.AddressOne
	}
	return ""
}

func (m *Address) GetAddressTwo() string {
	if m != nil {
		return m.AddressTwo
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtLocation)(nil), "ExtLocation")
	proto.RegisterType((*Address)(nil), "Address")
}

func init() { proto.RegisterFile("extended_location.proto", fileDescriptor_1e5b30b37c9f3a59) }

var fileDescriptor_1e5b30b37c9f3a59 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x99, 0x76, 0x7e, 0x4f, 0x77, 0xe1, 0xe3, 0x33, 0x88, 0x48, 0xe9, 0x42, 0xba, 0xea,
	0x42, 0xaf, 0x40, 0xc5, 0x9d, 0x20, 0x0c, 0xae, 0xdc, 0xc8, 0xfc, 0x84, 0x1a, 0x88, 0x49, 0x48,
	0x32, 0xd8, 0x7a, 0x2f, 0x5e, 0xaa, 0xe0, 0xe4, 0x24, 0xd3, 0xce, 0x6e, 0x9e, 0xf7, 0xe1, 0x65,
	0x5e, 0x4e, 0xe0, 0x82, 0x1d, 0x1c, 0x93, 0x3d, 0xeb, 0xdf, 0x85, 0xea, 0x1a, 0xc7, 0x95, 0xdc,
	0x69, 0xa3, 0x9c, 0xda, 0xfc, 0x24, 0xb0, 0x7a, 0x3a, 0xb8, 0xe7, 0x98, 0x12, 0x0a, 0x85, 0x36,
	0xfc, 0xb3, 0x31, 0x47, 0x9a, 0xac, 0x93, 0x6d, 0x56, 0x4f, 0x48, 0x2e, 0xa1, 0xb4, 0x1f, 0x5c,
	0x6b, 0x2e, 0xf7, 0x74, 0x81, 0xea, 0xc4, 0xbe, 0xd5, 0x72, 0x21, 0xbc, 0x5a, 0x86, 0x56, 0x44,
	0xf2, 0x1f, 0x72, 0xc3, 0xdc, 0x60, 0x24, 0x4d, 0x51, 0x44, 0x22, 0x37, 0x50, 0x35, 0x7d, 0x6f,
	0x98, 0xb5, 0xcc, 0xd2, 0x6c, 0xbd, 0xdc, 0xae, 0x6e, 0xcb, 0xdd, 0x7d, 0x48, 0xea, 0xb3, 0xda,
	0xfc, 0x26, 0x50, 0xc4, 0xd8, 0x2f, 0x10, 0xe3, 0x4a, 0x37, 0xf4, 0x0c, 0xc7, 0x55, 0xf5, 0x89,
	0xc9, 0x15, 0x54, 0x42, 0xc9, 0x7d, 0x90, 0x0b, 0x94, 0xe7, 0xc0, 0x37, 0xb5, 0x18, 0xec, 0xa3,
	0x1a, 0xe5, 0x32, 0x34, 0x27, 0x26, 0xd7, 0x00, 0xf1, 0x77, 0x2f, 0x92, 0xe1, 0xca, 0xaa, 0x9e,
	0x25, 0x33, 0xff, 0xfa, 0xa5, 0xc6, 0xa9, 0x73, 0x3f, 0x26, 0x84, 0x40, 0xda, 0x71, 0x77, 0xa4,
	0x39, 0x1a, 0xfc, 0x26, 0xff, 0x20, 0xb3, 0xae, 0x71, 0x8c, 0x16, 0x18, 0x06, 0xf0, 0x57, 0xea,
	0xd4, 0x20, 0xdd, 0x78, 0xdb, 0x12, 0xf3, 0x09, 0xbd, 0xf9, 0xe6, 0x1a, 0xe7, 0x55, 0xc1, 0x44,
	0x7c, 0x48, 0xdf, 0x16, 0xba, 0x6d, 0x73, 0x7c, 0xac, 0xbb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xde, 0x48, 0x03, 0xba, 0xc7, 0x01, 0x00, 0x00,
}
