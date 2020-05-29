// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Side int32

const (
	Side_Ask Side = 0
	Side_Bid Side = 1
)

var Side_name = map[int32]string{
	0: "Ask",
	1: "Bid",
}

var Side_value = map[string]int32{
	"Ask": 0,
	"Bid": 1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}

func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

type Strategy int32

const (
	Strategy_Market Strategy = 0
	Strategy_Limit  Strategy = 1
	Strategy_Follow Strategy = 2
	Strategy_Flex   Strategy = 3
)

var Strategy_name = map[int32]string{
	0: "Market",
	1: "Limit",
	2: "Follow",
	3: "Flex",
}

var Strategy_value = map[string]int32{
	"Market": 0,
	"Limit":  1,
	"Follow": 2,
	"Flex":   3,
}

func (x Strategy) String() string {
	return proto.EnumName(Strategy_name, int32(x))
}

func (Strategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

type Book_State int32

const (
	Book_Pending Book_State = 0
	Book_Paid    Book_State = 1
	Book_Done    Book_State = 2
)

var Book_State_name = map[int32]string{
	0: "Pending",
	1: "Paid",
	2: "Done",
}

var Book_State_value = map[string]int32{
	"Pending": 0,
	"Paid":    1,
	"Done":    2,
}

func (x Book_State) String() string {
	return proto.EnumName(Book_State_name, int32(x))
}

func (Book_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 0}
}

type Order_State int32

const (
	Order_Trading   Order_State = 0
	Order_Filled    Order_State = 1
	Order_Cancelled Order_State = 2
	Order_Rejected  Order_State = 3
	Order_Timeout   Order_State = 4
)

var Order_State_name = map[int32]string{
	0: "Trading",
	1: "Filled",
	2: "Cancelled",
	3: "Rejected",
	4: "Timeout",
}

var Order_State_value = map[string]int32{
	"Trading":   0,
	"Filled":    1,
	"Cancelled": 2,
	"Rejected":  3,
	"Timeout":   4,
}

func (x Order_State) String() string {
	return proto.EnumName(Order_State_name, int32(x))
}

func (Order_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1, 0}
}

type Book struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	State                Book_State           `protobuf:"varint,3,opt,name=state,proto3,enum=fox.atm.service.Book_State" json:"state,omitempty"`
	MerchantId           string               `protobuf:"bytes,4,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	BrokerId             string               `protobuf:"bytes,5,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	TraceId              string               `protobuf:"bytes,6,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	UserId               string               `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Payer                string               `protobuf:"bytes,8,opt,name=payer,proto3" json:"payer,omitempty"`
	SnapshotId           string               `protobuf:"bytes,9,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	ReceiptId            string               `protobuf:"bytes,10,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	Memo                 string               `protobuf:"bytes,11,opt,name=memo,proto3" json:"memo,omitempty"`
	Funds                string               `protobuf:"bytes,12,opt,name=funds,proto3" json:"funds,omitempty"`
	PaySymbol            string               `protobuf:"bytes,13,opt,name=pay_symbol,json=paySymbol,proto3" json:"pay_symbol,omitempty"`
	FillSymbol           string               `protobuf:"bytes,14,opt,name=fill_symbol,json=fillSymbol,proto3" json:"fill_symbol,omitempty"`
	Strategy             Strategy             `protobuf:"varint,15,opt,name=strategy,proto3,enum=fox.atm.service.Strategy" json:"strategy,omitempty"`
	Price                string               `protobuf:"bytes,16,opt,name=price,proto3" json:"price,omitempty"`
	Discount             string               `protobuf:"bytes,17,opt,name=discount,proto3" json:"discount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Book) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Book) GetState() Book_State {
	if m != nil {
		return m.State
	}
	return Book_Pending
}

func (m *Book) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *Book) GetBrokerId() string {
	if m != nil {
		return m.BrokerId
	}
	return ""
}

func (m *Book) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Book) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Book) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *Book) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *Book) GetReceiptId() string {
	if m != nil {
		return m.ReceiptId
	}
	return ""
}

func (m *Book) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Book) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

func (m *Book) GetPaySymbol() string {
	if m != nil {
		return m.PaySymbol
	}
	return ""
}

func (m *Book) GetFillSymbol() string {
	if m != nil {
		return m.FillSymbol
	}
	return ""
}

func (m *Book) GetStrategy() Strategy {
	if m != nil {
		return m.Strategy
	}
	return Strategy_Market
}

func (m *Book) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Book) GetDiscount() string {
	if m != nil {
		return m.Discount
	}
	return ""
}

type Order struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CancelledAt          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=cancelled_at,json=cancelledAt,proto3" json:"cancelled_at,omitempty"`
	UserId               string               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MerchantId           string               `protobuf:"bytes,6,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	State                Order_State          `protobuf:"varint,7,opt,name=state,proto3,enum=fox.atm.service.Order_State" json:"state,omitempty"`
	PaySymbol            string               `protobuf:"bytes,8,opt,name=pay_symbol,json=paySymbol,proto3" json:"pay_symbol,omitempty"`
	FillSymbol           string               `protobuf:"bytes,9,opt,name=fill_symbol,json=fillSymbol,proto3" json:"fill_symbol,omitempty"`
	Symbol               string               `protobuf:"bytes,10,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Side                 Side                 `protobuf:"varint,11,opt,name=side,proto3,enum=fox.atm.service.Side" json:"side,omitempty"`
	Strategy             Strategy             `protobuf:"varint,12,opt,name=strategy,proto3,enum=fox.atm.service.Strategy" json:"strategy,omitempty"`
	Price                string               `protobuf:"bytes,13,opt,name=price,proto3" json:"price,omitempty"`
	Discount             string               `protobuf:"bytes,14,opt,name=discount,proto3" json:"discount,omitempty"`
	Funds                string               `protobuf:"bytes,15,opt,name=funds,proto3" json:"funds,omitempty"`
	FilledFunds          string               `protobuf:"bytes,16,opt,name=filled_funds,json=filledFunds,proto3" json:"filled_funds,omitempty"`
	FilledAmount         string               `protobuf:"bytes,17,opt,name=filled_amount,json=filledAmount,proto3" json:"filled_amount,omitempty"`
	ExtraFilledAmount    string               `protobuf:"bytes,18,opt,name=extra_filled_amount,json=extraFilledAmount,proto3" json:"extra_filled_amount,omitempty"`
	FeeAmount            string               `protobuf:"bytes,19,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Order) GetCancelledAt() *timestamp.Timestamp {
	if m != nil {
		return m.CancelledAt
	}
	return nil
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetMerchantId() string {
	if m != nil {
		return m.MerchantId
	}
	return ""
}

func (m *Order) GetState() Order_State {
	if m != nil {
		return m.State
	}
	return Order_Trading
}

func (m *Order) GetPaySymbol() string {
	if m != nil {
		return m.PaySymbol
	}
	return ""
}

func (m *Order) GetFillSymbol() string {
	if m != nil {
		return m.FillSymbol
	}
	return ""
}

func (m *Order) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Order) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_Ask
}

func (m *Order) GetStrategy() Strategy {
	if m != nil {
		return m.Strategy
	}
	return Strategy_Market
}

func (m *Order) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Order) GetDiscount() string {
	if m != nil {
		return m.Discount
	}
	return ""
}

func (m *Order) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

func (m *Order) GetFilledFunds() string {
	if m != nil {
		return m.FilledFunds
	}
	return ""
}

func (m *Order) GetFilledAmount() string {
	if m != nil {
		return m.FilledAmount
	}
	return ""
}

func (m *Order) GetExtraFilledAmount() string {
	if m != nil {
		return m.ExtraFilledAmount
	}
	return ""
}

func (m *Order) GetFeeAmount() string {
	if m != nil {
		return m.FeeAmount
	}
	return ""
}

// merchant service
type MerchantServiceReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq) Reset()         { *m = MerchantServiceReq{} }
func (m *MerchantServiceReq) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq) ProtoMessage()    {}
func (*MerchantServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *MerchantServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq.Unmarshal(m, b)
}
func (m *MerchantServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq.Merge(m, src)
}
func (m *MerchantServiceReq) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq.Size(m)
}
func (m *MerchantServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq proto.InternalMessageInfo

type MerchantServiceReq_CreateBook struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReceiptId            string   `protobuf:"bytes,3,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_CreateBook) Reset()         { *m = MerchantServiceReq_CreateBook{} }
func (m *MerchantServiceReq_CreateBook) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_CreateBook) ProtoMessage()    {}
func (*MerchantServiceReq_CreateBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2, 0}
}

func (m *MerchantServiceReq_CreateBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Unmarshal(m, b)
}
func (m *MerchantServiceReq_CreateBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_CreateBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_CreateBook.Merge(m, src)
}
func (m *MerchantServiceReq_CreateBook) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_CreateBook.Size(m)
}
func (m *MerchantServiceReq_CreateBook) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_CreateBook.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_CreateBook proto.InternalMessageInfo

func (m *MerchantServiceReq_CreateBook) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *MerchantServiceReq_CreateBook) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MerchantServiceReq_CreateBook) GetReceiptId() string {
	if m != nil {
		return m.ReceiptId
	}
	return ""
}

type MerchantServiceReq_ReadOrder struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_ReadOrder) Reset()         { *m = MerchantServiceReq_ReadOrder{} }
func (m *MerchantServiceReq_ReadOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_ReadOrder) ProtoMessage()    {}
func (*MerchantServiceReq_ReadOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2, 1}
}

func (m *MerchantServiceReq_ReadOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Unmarshal(m, b)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_ReadOrder.Merge(m, src)
}
func (m *MerchantServiceReq_ReadOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_ReadOrder.Size(m)
}
func (m *MerchantServiceReq_ReadOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_ReadOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_ReadOrder proto.InternalMessageInfo

func (m *MerchantServiceReq_ReadOrder) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

type MerchantServiceReq_CancelOrder struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceReq_CancelOrder) Reset()         { *m = MerchantServiceReq_CancelOrder{} }
func (m *MerchantServiceReq_CancelOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceReq_CancelOrder) ProtoMessage()    {}
func (*MerchantServiceReq_CancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2, 2}
}

func (m *MerchantServiceReq_CancelOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Unmarshal(m, b)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceReq_CancelOrder.Merge(m, src)
}
func (m *MerchantServiceReq_CancelOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceReq_CancelOrder.Size(m)
}
func (m *MerchantServiceReq_CancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceReq_CancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceReq_CancelOrder proto.InternalMessageInfo

func (m *MerchantServiceReq_CancelOrder) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

type MerchantServiceResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceResp) Reset()         { *m = MerchantServiceResp{} }
func (m *MerchantServiceResp) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp) ProtoMessage()    {}
func (*MerchantServiceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *MerchantServiceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp.Unmarshal(m, b)
}
func (m *MerchantServiceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp.Merge(m, src)
}
func (m *MerchantServiceResp) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp.Size(m)
}
func (m *MerchantServiceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp proto.InternalMessageInfo

type MerchantServiceResp_CancelOrder struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantServiceResp_CancelOrder) Reset()         { *m = MerchantServiceResp_CancelOrder{} }
func (m *MerchantServiceResp_CancelOrder) String() string { return proto.CompactTextString(m) }
func (*MerchantServiceResp_CancelOrder) ProtoMessage()    {}
func (*MerchantServiceResp_CancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3, 0}
}

func (m *MerchantServiceResp_CancelOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Unmarshal(m, b)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Marshal(b, m, deterministic)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantServiceResp_CancelOrder.Merge(m, src)
}
func (m *MerchantServiceResp_CancelOrder) XXX_Size() int {
	return xxx_messageInfo_MerchantServiceResp_CancelOrder.Size(m)
}
func (m *MerchantServiceResp_CancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantServiceResp_CancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantServiceResp_CancelOrder proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("fox.atm.service.Side", Side_name, Side_value)
	proto.RegisterEnum("fox.atm.service.Strategy", Strategy_name, Strategy_value)
	proto.RegisterEnum("fox.atm.service.Book_State", Book_State_name, Book_State_value)
	proto.RegisterEnum("fox.atm.service.Order_State", Order_State_name, Order_State_value)
	proto.RegisterType((*Book)(nil), "fox.atm.service.Book")
	proto.RegisterType((*Order)(nil), "fox.atm.service.Order")
	proto.RegisterType((*MerchantServiceReq)(nil), "fox.atm.service.MerchantServiceReq")
	proto.RegisterType((*MerchantServiceReq_CreateBook)(nil), "fox.atm.service.MerchantServiceReq.CreateBook")
	proto.RegisterType((*MerchantServiceReq_ReadOrder)(nil), "fox.atm.service.MerchantServiceReq.ReadOrder")
	proto.RegisterType((*MerchantServiceReq_CancelOrder)(nil), "fox.atm.service.MerchantServiceReq.CancelOrder")
	proto.RegisterType((*MerchantServiceResp)(nil), "fox.atm.service.MerchantServiceResp")
	proto.RegisterType((*MerchantServiceResp_CancelOrder)(nil), "fox.atm.service.MerchantServiceResp.CancelOrder")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0x3f, 0x27, 0x4e, 0x62, 0x4f, 0xfe, 0x9c, 0xbb, 0x47, 0x8b, 0xeb, 0x52, 0xf5, 0x08, 0x08,
	0x85, 0x4a, 0xf8, 0x20, 0x08, 0xa1, 0x3e, 0xf0, 0x90, 0x2b, 0x3a, 0x29, 0x88, 0x8a, 0xca, 0xe9,
	0x13, 0x2f, 0xd1, 0x9e, 0x77, 0x72, 0x35, 0x67, 0x67, 0xcd, 0x7a, 0x03, 0x97, 0x4f, 0xc3, 0x17,
	0xe0, 0x7b, 0x20, 0xbe, 0x15, 0xda, 0x5d, 0xfb, 0x2e, 0x8e, 0x7b, 0x4a, 0x85, 0xfa, 0x14, 0xcf,
	0xfc, 0x7e, 0x33, 0x3b, 0x3b, 0xfb, 0x9b, 0x09, 0x0c, 0x0b, 0x14, 0x7f, 0x24, 0x31, 0x86, 0xb9,
	0xe0, 0x92, 0x93, 0xe3, 0x15, 0xbf, 0x09, 0xa9, 0xcc, 0xc2, 0xd2, 0x1d, 0x3c, 0xbb, 0xe2, 0xfc,
	0x2a, 0xc5, 0x33, 0x0d, 0x5f, 0x6e, 0x56, 0x67, 0x32, 0xc9, 0xb0, 0x90, 0x34, 0xcb, 0x4d, 0xc4,
	0xf8, 0x5f, 0x1b, 0xec, 0x73, 0xce, 0xaf, 0xc9, 0x08, 0x5a, 0x09, 0xf3, 0xad, 0x53, 0x6b, 0xe2,
	0x46, 0xad, 0x84, 0x91, 0x17, 0x00, 0xb1, 0x40, 0x2a, 0x91, 0x2d, 0xa9, 0xf4, 0x5b, 0xa7, 0xd6,
	0xa4, 0x3f, 0x0d, 0x42, 0x93, 0x2e, 0xac, 0xd2, 0x85, 0x6f, 0xaa, 0x74, 0x91, 0x5b, 0xb2, 0x67,
	0x92, 0x7c, 0x03, 0x9d, 0x42, 0x52, 0x89, 0x7e, 0xfb, 0xd4, 0x9a, 0x8c, 0xa6, 0x4f, 0xc2, 0xbd,
	0xaa, 0x42, 0x75, 0x60, 0xb8, 0x50, 0x94, 0xc8, 0x30, 0xc9, 0x33, 0xe8, 0x67, 0x28, 0xe2, 0xb7,
	0x74, 0x2d, 0x97, 0x09, 0xf3, 0x6d, 0x5d, 0x06, 0x54, 0xae, 0x39, 0x23, 0x4f, 0xc0, 0xbd, 0x14,
	0xfc, 0x1a, 0x85, 0x82, 0x3b, 0x1a, 0x76, 0x8c, 0x63, 0xce, 0xc8, 0x63, 0x70, 0xa4, 0xa0, 0x31,
	0x2a, 0xac, 0xab, 0xb1, 0x9e, 0xb6, 0xe7, 0x8c, 0x7c, 0x0c, 0xbd, 0x4d, 0x61, 0xa2, 0x7a, 0x1a,
	0xe9, 0x2a, 0x73, 0xce, 0xc8, 0x47, 0xd0, 0xc9, 0xe9, 0x16, 0x85, 0xef, 0x68, 0xb7, 0x31, 0x54,
	0x1d, 0xc5, 0x9a, 0xe6, 0xc5, 0x5b, 0xae, 0xeb, 0x70, 0x4d, 0x1d, 0x95, 0x6b, 0xce, 0xc8, 0x53,
	0x00, 0x81, 0x31, 0x26, 0xb9, 0xc6, 0x41, 0xe3, 0x6e, 0xe9, 0x99, 0x33, 0x42, 0xc0, 0xce, 0x30,
	0xe3, 0x7e, 0x5f, 0x03, 0xfa, 0x5b, 0x9d, 0xb4, 0xda, 0xac, 0x59, 0xe1, 0x0f, 0xcc, 0x49, 0xda,
	0x50, 0x89, 0x72, 0xba, 0x5d, 0x16, 0xdb, 0xec, 0x92, 0xa7, 0xfe, 0xd0, 0x24, 0xca, 0xe9, 0x76,
	0xa1, 0x1d, 0xaa, 0x90, 0x55, 0x92, 0xa6, 0x15, 0x3e, 0x32, 0x85, 0x28, 0x57, 0x49, 0xf8, 0x0e,
	0x9c, 0x42, 0x0a, 0x2a, 0xf1, 0x6a, 0xeb, 0x1f, 0xeb, 0x3e, 0x3f, 0x6e, 0xf4, 0x79, 0x51, 0x12,
	0xa2, 0x5b, 0xaa, 0xbe, 0xb6, 0x48, 0x62, 0xf4, 0xbd, 0xf2, 0xda, 0xca, 0x20, 0x01, 0x38, 0x2c,
	0x29, 0x62, 0xbe, 0x59, 0x4b, 0xff, 0x81, 0x69, 0x6e, 0x65, 0x8f, 0x27, 0xd0, 0xd1, 0x4f, 0x45,
	0xfa, 0xd0, 0x7b, 0x8d, 0x6b, 0x96, 0xac, 0xaf, 0xbc, 0x23, 0xe2, 0x80, 0xfd, 0x9a, 0x26, 0xcc,
	0xb3, 0xd4, 0xd7, 0x8f, 0x7c, 0x8d, 0x5e, 0x6b, 0xfc, 0x57, 0x17, 0x3a, 0xbf, 0x08, 0x86, 0xe2,
	0x43, 0x8a, 0xe9, 0x05, 0xc0, 0x26, 0x67, 0x55, 0x68, 0xfb, 0x70, 0x68, 0xc9, 0x9e, 0x49, 0xf2,
	0x03, 0x0c, 0x62, 0xba, 0x8e, 0x31, 0x4d, 0x4d, 0xb0, 0x7d, 0x30, 0xb8, 0x7f, 0xcb, 0x9f, 0xc9,
	0x5d, 0xe9, 0x74, 0x6a, 0xd2, 0xd9, 0x13, 0x6b, 0xb7, 0x21, 0xd6, 0x69, 0x35, 0x00, 0x3d, 0xfd,
	0x30, 0x9f, 0x34, 0x1e, 0x46, 0x77, 0xa9, 0x3e, 0x01, 0x75, 0x3d, 0x38, 0x07, 0xf4, 0xe0, 0x36,
	0xf4, 0xf0, 0x08, 0xba, 0x25, 0x66, 0x44, 0x59, 0x5a, 0xe4, 0x4b, 0xb0, 0x8b, 0x84, 0xa1, 0x56,
	0xe4, 0x68, 0xfa, 0xb0, 0xa9, 0x91, 0x84, 0x61, 0xa4, 0x29, 0x35, 0x49, 0x0d, 0xfe, 0x87, 0xa4,
	0x86, 0xf7, 0x49, 0x6a, 0x54, 0x97, 0xd4, 0xdd, 0x44, 0x1c, 0xef, 0x4e, 0xc4, 0xa7, 0x30, 0x50,
	0xf7, 0x41, 0xb6, 0x34, 0xa0, 0x51, 0x68, 0xdf, 0xf8, 0x2e, 0x34, 0xe5, 0x33, 0x18, 0x96, 0x14,
	0x9a, 0xed, 0x88, 0xb5, 0x8c, 0x9b, 0x69, 0x1f, 0x09, 0xe1, 0x04, 0x6f, 0xa4, 0xa0, 0xcb, 0x3a,
	0x95, 0x68, 0xea, 0x03, 0x0d, 0x5d, 0xec, 0xf2, 0x9f, 0x02, 0xac, 0x10, 0x2b, 0xda, 0x89, 0xe9,
	0xfc, 0x0a, 0xd1, 0xc0, 0xe3, 0x9f, 0x76, 0xf4, 0xff, 0x46, 0xd0, 0x52, 0xff, 0x00, 0x5d, 0x93,
	0xc4, 0xb3, 0xc8, 0x10, 0xdc, 0x97, 0x95, 0x6e, 0xbc, 0x16, 0x19, 0x80, 0x13, 0xe1, 0x6f, 0x18,
	0x4b, 0x64, 0x5e, 0x5b, 0x47, 0x25, 0x19, 0xf2, 0x8d, 0xf4, 0xec, 0xf1, 0x3f, 0x16, 0x90, 0x57,
	0xa5, 0x4e, 0x16, 0xa6, 0xa3, 0x11, 0xfe, 0x1e, 0x2c, 0x01, 0x5e, 0x6a, 0xc1, 0xeb, 0x4d, 0xbc,
	0xbb, 0xcd, 0xac, 0x7b, 0xb7, 0x59, 0xab, 0x26, 0xc9, 0xfa, 0x5a, 0x6a, 0xef, 0xad, 0xa5, 0xe0,
	0x0b, 0x70, 0x23, 0xa4, 0xcc, 0x0c, 0xe7, 0xfd, 0xf9, 0x83, 0x09, 0xf4, 0xcd, 0x4d, 0x0e, 0x31,
	0xc7, 0x9f, 0xc3, 0x49, 0xe3, 0x22, 0x45, 0x1e, 0x0c, 0x6b, 0x09, 0x9e, 0xfb, 0x60, 0x2b, 0x7d,
	0x91, 0x1e, 0xb4, 0x67, 0xc5, 0xb5, 0x77, 0xa4, 0x3e, 0xce, 0xd5, 0xd6, 0x78, 0xfe, 0x3d, 0x38,
	0x95, 0x94, 0x54, 0x2f, 0x5f, 0x51, 0x71, 0x8d, 0xd2, 0x3b, 0x22, 0x2e, 0x74, 0x7e, 0x4e, 0xb2,
	0x44, 0x7a, 0x96, 0x6e, 0x31, 0x4f, 0x53, 0xfe, 0xa7, 0xd7, 0x52, 0x4b, 0xe6, 0x22, 0xc5, 0x1b,
	0xaf, 0x3d, 0xfd, 0xbb, 0x05, 0xc7, 0x7b, 0x27, 0x93, 0x45, 0xad, 0x7f, 0x61, 0x43, 0xb4, 0xcd,
	0x96, 0x87, 0x77, 0xfc, 0xe0, 0xe1, 0x3b, 0xff, 0x9f, 0x48, 0xb4, 0xdb, 0xb3, 0xaf, 0xde, 0x27,
	0xe7, 0x2d, 0x3d, 0x78, 0xf4, 0xee, 0x89, 0x27, 0x79, 0xbd, 0xbf, 0x67, 0xef, 0x55, 0xe9, 0x5d,
	0x40, 0xf0, 0xf5, 0xe1, 0x80, 0x22, 0xdf, 0x8d, 0x38, 0xef, 0xfd, 0xda, 0x31, 0x7b, 0xae, 0xab,
	0x7f, 0xbe, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xfc, 0xa5, 0x68, 0x32, 0x08, 0x00, 0x00,
}
