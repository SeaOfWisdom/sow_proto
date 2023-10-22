// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: statistics-srv/statistics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogType int32

const (
	LogType_UNKNOWN        LogType = 0
	LogType_PURCHASED_WORK LogType = 1
	LogType_CLAIM_REWARDS  LogType = 2
)

// Enum value maps for LogType.
var (
	LogType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PURCHASED_WORK",
		2: "CLAIM_REWARDS",
	}
	LogType_value = map[string]int32{
		"UNKNOWN":        0,
		"PURCHASED_WORK": 1,
		"CLAIM_REWARDS":  2,
	}
)

func (x LogType) Enum() *LogType {
	p := new(LogType)
	*p = x
	return p
}

func (x LogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogType) Descriptor() protoreflect.EnumDescriptor {
	return file_statistics_srv_statistics_proto_enumTypes[0].Descriptor()
}

func (LogType) Type() protoreflect.EnumType {
	return &file_statistics_srv_statistics_proto_enumTypes[0]
}

func (x LogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogType.Descriptor instead.
func (LogType) EnumDescriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{0}
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Since   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Until   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=until,proto3" json:"until,omitempty"`
	LogType LogType                `protobuf:"varint,4,opt,name=log_type,json=logType,proto3,enum=pp.contractor.LogType" json:"log_type,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SearchRequest) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *SearchRequest) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

func (x *SearchRequest) GetLogType() LogType {
	if x != nil {
		return x.LogType
	}
	return LogType_UNKNOWN
}

type Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalEarnings     float64        `protobuf:"fixed64,1,opt,name=total_earnings,json=totalEarnings,proto3" json:"total_earnings,omitempty"`
	EarningsPeriod    float64        `protobuf:"fixed64,2,opt,name=earnings_period,json=earningsPeriod,proto3" json:"earnings_period,omitempty"`
	EarningsLastMonth float64        `protobuf:"fixed64,3,opt,name=earnings_last_month,json=earningsLastMonth,proto3" json:"earnings_last_month,omitempty"`
	EarningsLastWeek  float64        `protobuf:"fixed64,4,opt,name=earnings_last_week,json=earningsLastWeek,proto3" json:"earnings_last_week,omitempty"`
	Transactions      []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Statistics) Reset() {
	*x = Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistics) ProtoMessage() {}

func (x *Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistics.ProtoReflect.Descriptor instead.
func (*Statistics) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *Statistics) GetTotalEarnings() float64 {
	if x != nil {
		return x.TotalEarnings
	}
	return 0
}

func (x *Statistics) GetEarningsPeriod() float64 {
	if x != nil {
		return x.EarningsPeriod
	}
	return 0
}

func (x *Statistics) GetEarningsLastMonth() float64 {
	if x != nil {
		return x.EarningsLastMonth
	}
	return 0
}

func (x *Statistics) GetEarningsLastWeek() float64 {
	if x != nil {
		return x.EarningsLastWeek
	}
	return 0
}

func (x *Statistics) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash        string                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	BlockNumber uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Type        LogType                `protobuf:"varint,3,opt,name=type,proto3,enum=pp.contractor.LogType" json:"type,omitempty"`
	From        string                 `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To          string                 `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	WorkId      *string                `protobuf:"bytes,6,opt,name=work_id,json=workId,proto3,oneof" json:"work_id,omitempty"`
	Amount      float64                `protobuf:"fixed64,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Transaction) GetType() LogType {
	if x != nil {
		return x.Type
	}
	return LogType_UNKNOWN
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetWorkId() string {
	if x != nil && x.WorkId != nil {
		return *x.WorkId
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type TransactionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber     uint64                 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	TransactionHash string                 `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	Type            LogType                `protobuf:"varint,3,opt,name=type,proto3,enum=pp.contractor.LogType" json:"type,omitempty"`
	From            string                 `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To              string                 `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	WorkId          string                 `protobuf:"bytes,6,opt,name=work_id,json=workId,proto3" json:"work_id,omitempty"`
	Amount          string                 `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransactionLog) Reset() {
	*x = TransactionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLog) ProtoMessage() {}

func (x *TransactionLog) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLog.ProtoReflect.Descriptor instead.
func (*TransactionLog) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionLog) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TransactionLog) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TransactionLog) GetType() LogType {
	if x != nil {
		return x.Type
	}
	return LogType_UNKNOWN
}

func (x *TransactionLog) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransactionLog) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransactionLog) GetWorkId() string {
	if x != nil {
		return x.WorkId
	}
	return ""
}

func (x *TransactionLog) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TransactionLog) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type TransactionLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*TransactionLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *TransactionLogs) Reset() {
	*x = TransactionLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionLogs) ProtoMessage() {}

func (x *TransactionLogs) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionLogs.ProtoReflect.Descriptor instead.
func (*TransactionLogs) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionLogs) GetLogs() []*TransactionLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type WorkIDsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkIDs []string               `protobuf:"bytes,1,rep,name=workIDs,proto3" json:"workIDs,omitempty"`
	Since   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty"`
	Until   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *WorkIDsList) Reset() {
	*x = WorkIDsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkIDsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkIDsList) ProtoMessage() {}

func (x *WorkIDsList) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkIDsList.ProtoReflect.Descriptor instead.
func (*WorkIDsList) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{5}
}

func (x *WorkIDsList) GetWorkIDs() []string {
	if x != nil {
		return x.WorkIDs
	}
	return nil
}

func (x *WorkIDsList) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *WorkIDsList) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

type PurchasedWorkStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TxHash    string                 `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *PurchasedWorkStatistics) Reset() {
	*x = PurchasedWorkStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchasedWorkStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchasedWorkStatistics) ProtoMessage() {}

func (x *PurchasedWorkStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchasedWorkStatistics.ProtoReflect.Descriptor instead.
func (*PurchasedWorkStatistics) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{6}
}

func (x *PurchasedWorkStatistics) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *PurchasedWorkStatistics) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetPurchasedWorkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purchases map[string]float32 `protobuf:"bytes,1,rep,name=purchases,proto3" json:"purchases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *GetPurchasedWorkResponse) Reset() {
	*x = GetPurchasedWorkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_srv_statistics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPurchasedWorkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPurchasedWorkResponse) ProtoMessage() {}

func (x *GetPurchasedWorkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_srv_statistics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPurchasedWorkResponse.ProtoReflect.Descriptor instead.
func (*GetPurchasedWorkResponse) Descriptor() ([]byte, []int) {
	return file_statistics_srv_statistics_proto_rawDescGZIP(), []int{7}
}

func (x *GetPurchasedWorkResponse) GetPurchases() map[string]float32 {
	if x != nil {
		return x.Purchases
	}
	return nil
}

var File_statistics_srv_statistics_proto protoreflect.FileDescriptor

var file_statistics_srv_statistics_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x73, 0x72, 0x76,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05,
	0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x31,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0e, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x2c, 0x0a, 0x12, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x3e,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x90,
	0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x64, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x44, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x31, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69,
	0x6c, 0x22, 0x6b, 0x0a, 0x17, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0xae,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x57,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x3d, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x55, 0x52, 0x43, 0x48,
	0x41, 0x53, 0x45, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4c, 0x41, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x53, 0x10, 0x02, 0x32, 0x86,
	0x02, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x4f, 0x0a,
	0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x57,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x57, 0x6f,
	0x72, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statistics_srv_statistics_proto_rawDescOnce sync.Once
	file_statistics_srv_statistics_proto_rawDescData = file_statistics_srv_statistics_proto_rawDesc
)

func file_statistics_srv_statistics_proto_rawDescGZIP() []byte {
	file_statistics_srv_statistics_proto_rawDescOnce.Do(func() {
		file_statistics_srv_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistics_srv_statistics_proto_rawDescData)
	})
	return file_statistics_srv_statistics_proto_rawDescData
}

var file_statistics_srv_statistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_statistics_srv_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_statistics_srv_statistics_proto_goTypes = []interface{}{
	(LogType)(0),                     // 0: pp.contractor.LogType
	(*SearchRequest)(nil),            // 1: pp.contractor.SearchRequest
	(*Statistics)(nil),               // 2: pp.contractor.Statistics
	(*Transaction)(nil),              // 3: pp.contractor.Transaction
	(*TransactionLog)(nil),           // 4: pp.contractor.TransactionLog
	(*TransactionLogs)(nil),          // 5: pp.contractor.TransactionLogs
	(*WorkIDsList)(nil),              // 6: pp.contractor.WorkIDsList
	(*PurchasedWorkStatistics)(nil),  // 7: pp.contractor.PurchasedWorkStatistics
	(*GetPurchasedWorkResponse)(nil), // 8: pp.contractor.GetPurchasedWorkResponse
	nil,                              // 9: pp.contractor.GetPurchasedWorkResponse.PurchasesEntry
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_statistics_srv_statistics_proto_depIdxs = []int32{
	10, // 0: pp.contractor.SearchRequest.since:type_name -> google.protobuf.Timestamp
	10, // 1: pp.contractor.SearchRequest.until:type_name -> google.protobuf.Timestamp
	0,  // 2: pp.contractor.SearchRequest.log_type:type_name -> pp.contractor.LogType
	3,  // 3: pp.contractor.Statistics.transactions:type_name -> pp.contractor.Transaction
	0,  // 4: pp.contractor.Transaction.type:type_name -> pp.contractor.LogType
	10, // 5: pp.contractor.Transaction.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 6: pp.contractor.TransactionLog.type:type_name -> pp.contractor.LogType
	10, // 7: pp.contractor.TransactionLog.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 8: pp.contractor.TransactionLogs.logs:type_name -> pp.contractor.TransactionLog
	10, // 9: pp.contractor.WorkIDsList.since:type_name -> google.protobuf.Timestamp
	10, // 10: pp.contractor.WorkIDsList.until:type_name -> google.protobuf.Timestamp
	10, // 11: pp.contractor.PurchasedWorkStatistics.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 12: pp.contractor.GetPurchasedWorkResponse.purchases:type_name -> pp.contractor.GetPurchasedWorkResponse.PurchasesEntry
	1,  // 13: pp.contractor.StatisticService.GetStatistics:input_type -> pp.contractor.SearchRequest
	5,  // 14: pp.contractor.StatisticService.StoreBatchTransaction:input_type -> pp.contractor.TransactionLogs
	6,  // 15: pp.contractor.StatisticService.GetPurchasedWork:input_type -> pp.contractor.WorkIDsList
	2,  // 16: pp.contractor.StatisticService.GetStatistics:output_type -> pp.contractor.Statistics
	11, // 17: pp.contractor.StatisticService.StoreBatchTransaction:output_type -> google.protobuf.Empty
	8,  // 18: pp.contractor.StatisticService.GetPurchasedWork:output_type -> pp.contractor.GetPurchasedWorkResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_statistics_srv_statistics_proto_init() }
func file_statistics_srv_statistics_proto_init() {
	if File_statistics_srv_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistics_srv_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionLogs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkIDsList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchasedWorkStatistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_statistics_srv_statistics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPurchasedWorkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_statistics_srv_statistics_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_statistics_srv_statistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statistics_srv_statistics_proto_goTypes,
		DependencyIndexes: file_statistics_srv_statistics_proto_depIdxs,
		EnumInfos:         file_statistics_srv_statistics_proto_enumTypes,
		MessageInfos:      file_statistics_srv_statistics_proto_msgTypes,
	}.Build()
	File_statistics_srv_statistics_proto = out.File
	file_statistics_srv_statistics_proto_rawDesc = nil
	file_statistics_srv_statistics_proto_goTypes = nil
	file_statistics_srv_statistics_proto_depIdxs = nil
}
