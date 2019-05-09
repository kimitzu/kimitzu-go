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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xb6, 0xae, 0xcd, 0xb7, 0x5b, 0x10, 0x0d, 0x22, 0x32, 0x76, 0x90, 0x9d, 0x76,
	0xd0, 0x5f, 0xa0, 0xe2, 0x4d, 0x10, 0xca, 0x4e, 0x5e, 0xa4, 0x6d, 0xc2, 0x0c, 0xc4, 0x24, 0x24,
	0x29, 0x5b, 0xfd, 0x2f, 0xfe, 0x54, 0x41, 0xf2, 0x35, 0xeb, 0xe6, 0xad, 0xcf, 0xf3, 0xf0, 0xd1,
	0x97, 0x16, 0xae, 0xc4, 0x21, 0x08, 0xcd, 0x05, 0xff, 0x50, 0xa6, 0xad, 0x83, 0x34, 0x7a, 0x63,
	0x9d, 0x09, 0x66, 0xf5, 0x93, 0xc1, 0xe2, 0xe5, 0x10, 0x5e, 0x93, 0xa5, 0x0c, 0x0a, 0xeb, 0xe4,
	0x57, 0xed, 0x7a, 0x96, 0x2d, 0xb3, 0x75, 0x5e, 0x1d, 0x91, 0x5e, 0x43, 0xe9, 0x3f, 0xa5, 0xb5,
	0x52, 0xef, 0xd8, 0x04, 0xd3, 0xc8, 0xf1, 0xaa, 0x91, 0x4a, 0xc5, 0x34, 0x1d, 0xae, 0x12, 0xd2,
	0x4b, 0x98, 0x3b, 0x11, 0x3a, 0xa7, 0xd9, 0x0c, 0x43, 0x22, 0x7a, 0x07, 0xa4, 0xe6, 0xdc, 0x09,
	0xef, 0x85, 0x67, 0xf9, 0x72, 0xba, 0x5e, 0xdc, 0x97, 0x9b, 0xc7, 0xc1, 0x54, 0xa7, 0xb4, 0xfa,
	0xcd, 0xa0, 0x48, 0x3a, 0x2e, 0x50, 0x75, 0x90, 0xa1, 0xe3, 0x02, 0xc7, 0x91, 0x6a, 0x64, 0x7a,
	0x03, 0x44, 0x19, 0xbd, 0x1b, 0xe2, 0x04, 0xe3, 0x49, 0xc4, 0x4b, 0xab, 0x3a, 0xff, 0x6c, 0xb8,
	0xc0, 0x81, 0xa4, 0x1a, 0x99, 0xde, 0x02, 0xa4, 0xd7, 0xbd, 0x69, 0x81, 0x2b, 0x49, 0x75, 0x66,
	0xce, 0xfa, 0x76, 0x6f, 0x58, 0xfe, 0xaf, 0x6f, 0xf7, 0x86, 0x52, 0x98, 0xb5, 0x32, 0xf4, 0x6c,
	0x8e, 0x05, 0x9f, 0xe9, 0x05, 0xe4, 0x3e, 0xd4, 0x41, 0xb0, 0x02, 0xe5, 0x00, 0xf1, 0x2b, 0xb5,
	0xa6, 0xd3, 0xc1, 0xf5, 0xac, 0x44, 0x7f, 0xc4, 0x58, 0xbe, 0xa5, 0xc5, 0x79, 0x64, 0x28, 0x09,
	0x9f, 0x66, 0xef, 0x13, 0xdb, 0x34, 0x73, 0xfc, 0x59, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xde, 0x48, 0x03, 0xba, 0xc7, 0x01, 0x00, 0x00,
}
