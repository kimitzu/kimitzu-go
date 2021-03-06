// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Coupon struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Coupon) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OrderRespApi struct {
	Contract                   *RicardianContract   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	State                      OrderState           `protobuf:"varint,2,opt,name=state,proto3,enum=OrderState" json:"state,omitempty"`
	Read                       bool                 `protobuf:"varint,3,opt,name=read,proto3" json:"read,omitempty"`
	Funded                     bool                 `protobuf:"varint,4,opt,name=funded,proto3" json:"funded,omitempty"`
	UnreadChatMessages         uint64               `protobuf:"varint,5,opt,name=unreadChatMessages,proto3" json:"unreadChatMessages,omitempty"`
	PaymentAddressTransactions []*TransactionRecord `protobuf:"bytes,6,rep,name=paymentAddressTransactions,proto3" json:"paymentAddressTransactions,omitempty"`
	RefundAddressTransaction   *TransactionRecord   `protobuf:"bytes,7,opt,name=refundAddressTransaction,proto3" json:"refundAddressTransaction,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *OrderRespApi) Reset()         { *m = OrderRespApi{} }
func (m *OrderRespApi) String() string { return proto.CompactTextString(m) }
func (*OrderRespApi) ProtoMessage()    {}
func (*OrderRespApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *OrderRespApi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRespApi.Unmarshal(m, b)
}
func (m *OrderRespApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRespApi.Marshal(b, m, deterministic)
}
func (m *OrderRespApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRespApi.Merge(m, src)
}
func (m *OrderRespApi) XXX_Size() int {
	return xxx_messageInfo_OrderRespApi.Size(m)
}
func (m *OrderRespApi) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRespApi.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRespApi proto.InternalMessageInfo

func (m *OrderRespApi) GetContract() *RicardianContract {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *OrderRespApi) GetState() OrderState {
	if m != nil {
		return m.State
	}
	return OrderState_PENDING
}

func (m *OrderRespApi) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *OrderRespApi) GetFunded() bool {
	if m != nil {
		return m.Funded
	}
	return false
}

func (m *OrderRespApi) GetUnreadChatMessages() uint64 {
	if m != nil {
		return m.UnreadChatMessages
	}
	return 0
}

func (m *OrderRespApi) GetPaymentAddressTransactions() []*TransactionRecord {
	if m != nil {
		return m.PaymentAddressTransactions
	}
	return nil
}

func (m *OrderRespApi) GetRefundAddressTransaction() *TransactionRecord {
	if m != nil {
		return m.RefundAddressTransaction
	}
	return nil
}

type CaseRespApi struct {
	Timestamp                      *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BuyerContract                  *RicardianContract   `protobuf:"bytes,2,opt,name=buyerContract,proto3" json:"buyerContract,omitempty"`
	VendorContract                 *RicardianContract   `protobuf:"bytes,3,opt,name=vendorContract,proto3" json:"vendorContract,omitempty"`
	BuyerContractValidationErrors  []string             `protobuf:"bytes,4,rep,name=buyerContractValidationErrors,proto3" json:"buyerContractValidationErrors,omitempty"`
	VendorContractValidationErrors []string             `protobuf:"bytes,5,rep,name=vendorContractValidationErrors,proto3" json:"vendorContractValidationErrors,omitempty"`
	State                          OrderState           `protobuf:"varint,6,opt,name=state,proto3,enum=OrderState" json:"state,omitempty"`
	Read                           bool                 `protobuf:"varint,7,opt,name=read,proto3" json:"read,omitempty"`
	BuyerOpened                    bool                 `protobuf:"varint,8,opt,name=buyerOpened,proto3" json:"buyerOpened,omitempty"`
	Claim                          string               `protobuf:"bytes,9,opt,name=claim,proto3" json:"claim,omitempty"`
	UnreadChatMessages             uint64               `protobuf:"varint,10,opt,name=unreadChatMessages,proto3" json:"unreadChatMessages,omitempty"`
	Resolution                     *DisputeResolution   `protobuf:"bytes,11,opt,name=resolution,proto3" json:"resolution,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}             `json:"-"`
	XXX_unrecognized               []byte               `json:"-"`
	XXX_sizecache                  int32                `json:"-"`
}

func (m *CaseRespApi) Reset()         { *m = CaseRespApi{} }
func (m *CaseRespApi) String() string { return proto.CompactTextString(m) }
func (*CaseRespApi) ProtoMessage()    {}
func (*CaseRespApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *CaseRespApi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaseRespApi.Unmarshal(m, b)
}
func (m *CaseRespApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaseRespApi.Marshal(b, m, deterministic)
}
func (m *CaseRespApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaseRespApi.Merge(m, src)
}
func (m *CaseRespApi) XXX_Size() int {
	return xxx_messageInfo_CaseRespApi.Size(m)
}
func (m *CaseRespApi) XXX_DiscardUnknown() {
	xxx_messageInfo_CaseRespApi.DiscardUnknown(m)
}

var xxx_messageInfo_CaseRespApi proto.InternalMessageInfo

func (m *CaseRespApi) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *CaseRespApi) GetBuyerContract() *RicardianContract {
	if m != nil {
		return m.BuyerContract
	}
	return nil
}

func (m *CaseRespApi) GetVendorContract() *RicardianContract {
	if m != nil {
		return m.VendorContract
	}
	return nil
}

func (m *CaseRespApi) GetBuyerContractValidationErrors() []string {
	if m != nil {
		return m.BuyerContractValidationErrors
	}
	return nil
}

func (m *CaseRespApi) GetVendorContractValidationErrors() []string {
	if m != nil {
		return m.VendorContractValidationErrors
	}
	return nil
}

func (m *CaseRespApi) GetState() OrderState {
	if m != nil {
		return m.State
	}
	return OrderState_PENDING
}

func (m *CaseRespApi) GetRead() bool {
	if m != nil {
		return m.Read
	}
	return false
}

func (m *CaseRespApi) GetBuyerOpened() bool {
	if m != nil {
		return m.BuyerOpened
	}
	return false
}

func (m *CaseRespApi) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *CaseRespApi) GetUnreadChatMessages() uint64 {
	if m != nil {
		return m.UnreadChatMessages
	}
	return 0
}

func (m *CaseRespApi) GetResolution() *DisputeResolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

type TransactionRecord struct {
	Txid                 string               `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Value                int64                `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Confirmations        uint32               `protobuf:"varint,3,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	Height               uint32               `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransactionRecord) Reset()         { *m = TransactionRecord{} }
func (m *TransactionRecord) String() string { return proto.CompactTextString(m) }
func (*TransactionRecord) ProtoMessage()    {}
func (*TransactionRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *TransactionRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRecord.Unmarshal(m, b)
}
func (m *TransactionRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRecord.Marshal(b, m, deterministic)
}
func (m *TransactionRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRecord.Merge(m, src)
}
func (m *TransactionRecord) XXX_Size() int {
	return xxx_messageInfo_TransactionRecord.Size(m)
}
func (m *TransactionRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRecord proto.InternalMessageInfo

func (m *TransactionRecord) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *TransactionRecord) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransactionRecord) GetConfirmations() uint32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *TransactionRecord) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TransactionRecord) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PeerAndProfile struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Profile              *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerAndProfile) Reset()         { *m = PeerAndProfile{} }
func (m *PeerAndProfile) String() string { return proto.CompactTextString(m) }
func (*PeerAndProfile) ProtoMessage()    {}
func (*PeerAndProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *PeerAndProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAndProfile.Unmarshal(m, b)
}
func (m *PeerAndProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAndProfile.Marshal(b, m, deterministic)
}
func (m *PeerAndProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAndProfile.Merge(m, src)
}
func (m *PeerAndProfile) XXX_Size() int {
	return xxx_messageInfo_PeerAndProfile.Size(m)
}
func (m *PeerAndProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAndProfile.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAndProfile proto.InternalMessageInfo

func (m *PeerAndProfile) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerAndProfile) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type PeerAndProfileWithID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerId               string   `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Profile              *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerAndProfileWithID) Reset()         { *m = PeerAndProfileWithID{} }
func (m *PeerAndProfileWithID) String() string { return proto.CompactTextString(m) }
func (*PeerAndProfileWithID) ProtoMessage()    {}
func (*PeerAndProfileWithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *PeerAndProfileWithID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAndProfileWithID.Unmarshal(m, b)
}
func (m *PeerAndProfileWithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAndProfileWithID.Marshal(b, m, deterministic)
}
func (m *PeerAndProfileWithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAndProfileWithID.Merge(m, src)
}
func (m *PeerAndProfileWithID) XXX_Size() int {
	return xxx_messageInfo_PeerAndProfileWithID.Size(m)
}
func (m *PeerAndProfileWithID) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAndProfileWithID.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAndProfileWithID proto.InternalMessageInfo

func (m *PeerAndProfileWithID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerAndProfileWithID) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *PeerAndProfileWithID) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type RatingWithID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RatingId             string   `protobuf:"bytes,2,opt,name=ratingId,proto3" json:"ratingId,omitempty"`
	Rating               *Rating  `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RatingWithID) Reset()         { *m = RatingWithID{} }
func (m *RatingWithID) String() string { return proto.CompactTextString(m) }
func (*RatingWithID) ProtoMessage()    {}
func (*RatingWithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *RatingWithID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatingWithID.Unmarshal(m, b)
}
func (m *RatingWithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatingWithID.Marshal(b, m, deterministic)
}
func (m *RatingWithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatingWithID.Merge(m, src)
}
func (m *RatingWithID) XXX_Size() int {
	return xxx_messageInfo_RatingWithID.Size(m)
}
func (m *RatingWithID) XXX_DiscardUnknown() {
	xxx_messageInfo_RatingWithID.DiscardUnknown(m)
}

var xxx_messageInfo_RatingWithID proto.InternalMessageInfo

func (m *RatingWithID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RatingWithID) GetRatingId() string {
	if m != nil {
		return m.RatingId
	}
	return ""
}

func (m *RatingWithID) GetRating() *Rating {
	if m != nil {
		return m.Rating
	}
	return nil
}

func init() {
	proto.RegisterType((*Coupon)(nil), "Coupon")
	proto.RegisterType((*OrderRespApi)(nil), "OrderRespApi")
	proto.RegisterType((*CaseRespApi)(nil), "CaseRespApi")
	proto.RegisterType((*TransactionRecord)(nil), "TransactionRecord")
	proto.RegisterType((*PeerAndProfile)(nil), "PeerAndProfile")
	proto.RegisterType((*PeerAndProfileWithID)(nil), "PeerAndProfileWithID")
	proto.RegisterType((*RatingWithID)(nil), "RatingWithID")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0xf3, 0x9f, 0x49, 0x13, 0x84, 0x55, 0xa1, 0x55, 0x24, 0x68, 0x59, 0x71, 0xe0, 0xb4,
	0x45, 0xe5, 0x82, 0xb8, 0x95, 0x14, 0xa4, 0x4a, 0x40, 0x2b, 0x53, 0x81, 0x04, 0x27, 0x27, 0x76,
	0x12, 0x4b, 0x89, 0xbd, 0xb2, 0xbd, 0x15, 0x7d, 0x19, 0xde, 0x82, 0x07, 0xe2, 0x4d, 0xf0, 0xdf,
	0xa6, 0x59, 0xd2, 0x6d, 0xc5, 0x29, 0x9e, 0x99, 0x6f, 0xbe, 0x99, 0x9d, 0xf9, 0x26, 0xd0, 0x27,
	0x39, 0xcf, 0x72, 0x25, 0x8d, 0x1c, 0x3f, 0x9a, 0x49, 0x61, 0x14, 0x99, 0x19, 0x1d, 0x1d, 0xfb,
	0x52, 0x51, 0xa6, 0x4a, 0x6b, 0x68, 0x7f, 0xe6, 0x7c, 0xc5, 0xa2, 0x79, 0xb8, 0x90, 0x72, 0xb1,
	0x62, 0xc7, 0xde, 0x9a, 0x16, 0xf3, 0x63, 0xc3, 0xd7, 0x4c, 0x1b, 0xb2, 0xce, 0x03, 0x20, 0x7d,
	0x05, 0x9d, 0x89, 0x2c, 0x72, 0x29, 0x10, 0x82, 0xd6, 0x92, 0xe8, 0x65, 0xb2, 0x77, 0xb4, 0xf7,
	0xb2, 0x8f, 0xfd, 0xdb, 0xf9, 0x66, 0x92, 0xb2, 0xa4, 0x11, 0x7c, 0xee, 0x9d, 0xfe, 0x69, 0xc0,
	0xfe, 0x85, 0x2b, 0x89, 0x99, 0xce, 0x4f, 0x73, 0x8e, 0x32, 0xe8, 0x95, 0x3d, 0xf9, 0xe4, 0xc1,
	0x09, 0xca, 0x30, 0x9f, 0x11, 0x45, 0x39, 0x11, 0x93, 0x18, 0xc1, 0x1b, 0x0c, 0x7a, 0x0e, 0x6d,
	0xdb, 0x81, 0x09, 0xac, 0xa3, 0x93, 0x41, 0xe6, 0xd9, 0xbe, 0x38, 0x17, 0x0e, 0x11, 0x57, 0x57,
	0x31, 0x42, 0x93, 0xa6, 0x45, 0xf4, 0xb0, 0x7f, 0xa3, 0x27, 0xd0, 0x99, 0x17, 0x82, 0x32, 0x9a,
	0xb4, 0xbc, 0x37, 0x5a, 0xb6, 0x3c, 0x2a, 0x84, 0x43, 0x4c, 0x96, 0xc4, 0x7c, 0x62, 0x5a, 0x93,
	0x05, 0xd3, 0x49, 0xdb, 0x62, 0x5a, 0xf8, 0x8e, 0x08, 0xc2, 0x30, 0xce, 0xc9, 0xcd, 0x9a, 0x09,
	0x73, 0x4a, 0xa9, 0xb2, 0xde, 0x2b, 0x45, 0x84, 0xb6, 0x8d, 0x71, 0x29, 0x74, 0xd2, 0x39, 0x6a,
	0xfa, 0x0f, 0xd8, 0x72, 0x62, 0x36, 0xb3, 0x23, 0xc6, 0xf7, 0x64, 0xa1, 0xcf, 0x90, 0x28, 0xe6,
	0xfa, 0xd9, 0x0d, 0x26, 0xdd, 0x38, 0x92, 0x5d, 0xc6, 0xda, 0x9c, 0xf4, 0x57, 0x0b, 0x06, 0x13,
	0xa2, 0x59, 0x39, 0xe2, 0x37, 0xd0, 0xdf, 0x2c, 0x2e, 0xce, 0x78, 0x9c, 0x85, 0xd5, 0x66, 0xe5,
	0x6a, 0xb3, 0xab, 0x12, 0x81, 0x6f, 0xc1, 0x36, 0x73, 0x38, 0x2d, 0x6e, 0x98, 0x2a, 0xf7, 0xe0,
	0x87, 0x7e, 0xf7, 0x86, 0xaa, 0x40, 0xf4, 0x16, 0x46, 0xd7, 0x4c, 0x50, 0x79, 0x9b, 0xda, 0xac,
	0x4d, 0xfd, 0x07, 0x89, 0xce, 0xe0, 0x69, 0x85, 0xec, 0x2b, 0x59, 0x71, 0x4a, 0xdc, 0xa7, 0xbd,
	0x57, 0x4a, 0x2a, 0x6d, 0x57, 0xd8, 0xb4, 0x82, 0xba, 0x1f, 0x84, 0x3e, 0xc0, 0xb3, 0x2a, 0xef,
	0x0e, 0x4d, 0xdb, 0xd3, 0x3c, 0x80, 0xba, 0x15, 0x5c, 0xe7, 0x41, 0xc1, 0x75, 0xb7, 0x04, 0x77,
	0x04, 0x03, 0xdf, 0xdf, 0x45, 0xce, 0x84, 0x55, 0x5d, 0xcf, 0x87, 0xb6, 0x5d, 0xe8, 0x00, 0xda,
	0xb3, 0x15, 0xe1, 0xeb, 0xa4, 0xef, 0xef, 0x23, 0x18, 0x35, 0x82, 0x84, 0x5a, 0x41, 0x9e, 0x00,
	0xd8, 0xfd, 0xcb, 0x55, 0xe1, 0xe5, 0x32, 0x88, 0x43, 0x3e, 0xe3, 0x3a, 0x2f, 0x8c, 0x53, 0x40,
	0x8c, 0xe0, 0x2d, 0x54, 0xfa, 0x7b, 0x0f, 0x1e, 0xef, 0x08, 0xca, 0x7d, 0x85, 0xf9, 0xc9, 0x69,
	0x79, 0xc2, 0xee, 0xed, 0x7a, 0xbc, 0x26, 0xab, 0x22, 0x5c, 0x5b, 0x13, 0x07, 0x03, 0xbd, 0x80,
	0xa1, 0xbd, 0xc7, 0x39, 0x57, 0x6b, 0x12, 0x74, 0xef, 0x76, 0x3b, 0xc4, 0x55, 0xa7, 0x3b, 0xb9,
	0x25, 0xe3, 0x8b, 0xa5, 0xf1, 0x27, 0x37, 0xc4, 0xd1, 0xaa, 0xca, 0xb1, 0xfd, 0x1f, 0x72, 0x4c,
	0x3f, 0xc2, 0xe8, 0x92, 0x31, 0x75, 0x2a, 0xe8, 0x65, 0xf8, 0x9f, 0x72, 0x35, 0x72, 0xeb, 0x39,
	0x2f, 0xbb, 0x8e, 0x16, 0x4a, 0xa1, 0x1b, 0xff, 0xca, 0xa2, 0x64, 0x7b, 0x59, 0x4c, 0xc1, 0x65,
	0x20, 0x9d, 0xc2, 0x41, 0x95, 0xed, 0x1b, 0x37, 0xcb, 0xf3, 0x33, 0x34, 0x82, 0xc6, 0x66, 0x0a,
	0xf6, 0xb5, 0x55, 0xa3, 0x51, 0x57, 0xa3, 0x59, 0x57, 0xe3, 0x07, 0xec, 0x63, 0x3b, 0x0e, 0xb1,
	0xa8, 0xe1, 0x1e, 0x43, 0x4f, 0xf9, 0xf8, 0x86, 0x7d, 0x63, 0xa3, 0x43, 0xe8, 0x84, 0x77, 0xa4,
	0xef, 0x66, 0x81, 0x0a, 0x47, 0xf7, 0xbb, 0xd6, 0xf7, 0x46, 0x3e, 0x9d, 0x76, 0xfc, 0xcc, 0x5e,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x96, 0xc3, 0x9b, 0xe6, 0x05, 0x00, 0x00,
}
