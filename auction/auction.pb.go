// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: auction/auction.proto

package auction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BackupDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *BackupDetails) Reset() {
	*x = BackupDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupDetails) ProtoMessage() {}

func (x *BackupDetails) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupDetails.ProtoReflect.Descriptor instead.
func (*BackupDetails) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BackupDetails) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type BackupJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *BackupJoin) Reset() {
	*x = BackupJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupJoin) ProtoMessage() {}

func (x *BackupJoin) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupJoin.ProtoReflect.Descriptor instead.
func (*BackupJoin) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BackupJoin) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BackupJoin) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type BackupLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *BackupLeave) Reset() {
	*x = BackupLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupLeave) ProtoMessage() {}

func (x *BackupLeave) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupLeave.ProtoReflect.Descriptor instead.
func (*BackupLeave) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{2}
}

func (x *BackupLeave) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid    uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{3}
}

func (x *Bid) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Bid) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BidAck) Reset() {
	*x = BidAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidAck) ProtoMessage() {}

func (x *BidAck) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidAck.ProtoReflect.Descriptor instead.
func (*BidAck) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{4}
}

func (x *BidAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bidders    []uint32 `protobuf:"varint,1,rep,packed,name=bidders,proto3" json:"bidders,omitempty"` // Process ids of all bidders.
	Winner     uint32   `protobuf:"varint,2,opt,name=winner,proto3" json:"winner,omitempty"`          // The process id of the winner. 0 means no winner yet.
	HighestBid uint32   `protobuf:"varint,3,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{5}
}

func (x *Outcome) GetBidders() []uint32 {
	if x != nil {
		return x.Bidders
	}
	return nil
}

func (x *Outcome) GetWinner() uint32 {
	if x != nil {
		return x.Winner
	}
	return 0
}

func (x *Outcome) GetHighestBid() uint32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type Election struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *Election) Reset() {
	*x = Election{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Election) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Election) ProtoMessage() {}

func (x *Election) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Election.ProtoReflect.Descriptor instead.
func (*Election) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{6}
}

func (x *Election) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Elected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Elected) Reset() {
	*x = Elected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Elected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Elected) ProtoMessage() {}

func (x *Elected) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Elected.ProtoReflect.Descriptor instead.
func (*Elected) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{7}
}

func (x *Elected) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Elected) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auction_auction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{8}
}

var File_auction_auction_proto protoreflect.FileDescriptor

var file_auction_auction_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x21, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x1f, 0x0a,
	0x0b, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x2f,
	0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x22, 0x0a, 0x06, 0x42, 0x69, 0x64, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x22, 0x1c, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x2f,
	0x0a, 0x07, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0xd2, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3b,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x80, 0x01, 0x0a,
	0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x69, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x15,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42,
	0x69, 0x64, 0x41, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x32,
	0x87, 0x01, 0x0a, 0x0b, 0x52, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a,
	0x0b, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x32, 0xc0, 0x01, 0x0a, 0x0f, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x69, 0x64, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x15, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x69, 0x64,
	0x41, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3d, 0x0a,
	0x0e, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x72, 0x70, 0x68, 0x69,
	0x6c, 0x32, 0x31, 0x30, 0x35, 0x2f, 0x44, 0x49, 0x53, 0x59, 0x53, 0x2d, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x66, 0x69, 0x76, 0x65, 0x3b, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_auction_proto_rawDescOnce sync.Once
	file_auction_auction_proto_rawDescData = file_auction_auction_proto_rawDesc
)

func file_auction_auction_proto_rawDescGZIP() []byte {
	file_auction_auction_proto_rawDescOnce.Do(func() {
		file_auction_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_auction_proto_rawDescData)
	})
	return file_auction_auction_proto_rawDescData
}

var file_auction_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auction_auction_proto_goTypes = []interface{}{
	(*BackupDetails)(nil), // 0: auctionSystem.BackupDetails
	(*BackupJoin)(nil),    // 1: auctionSystem.BackupJoin
	(*BackupLeave)(nil),   // 2: auctionSystem.BackupLeave
	(*Bid)(nil),           // 3: auctionSystem.Bid
	(*BidAck)(nil),        // 4: auctionSystem.BidAck
	(*Outcome)(nil),       // 5: auctionSystem.Outcome
	(*Election)(nil),      // 6: auctionSystem.Election
	(*Elected)(nil),       // 7: auctionSystem.Elected
	(*Void)(nil),          // 8: auctionSystem.Void
}
var file_auction_auction_proto_depIdxs = []int32{
	8,  // 0: auctionSystem.ConnectService.FinishConnect:input_type -> auctionSystem.Void
	1,  // 1: auctionSystem.ConnectService.AddBackup:input_type -> auctionSystem.BackupJoin
	2,  // 2: auctionSystem.ConnectService.RemoveBackup:input_type -> auctionSystem.BackupLeave
	3,  // 3: auctionSystem.AuctionService.SendBid:input_type -> auctionSystem.Bid
	8,  // 4: auctionSystem.AuctionService.GetResult:input_type -> auctionSystem.Void
	6,  // 5: auctionSystem.RingService.SendElection:input_type -> auctionSystem.Election
	7,  // 6: auctionSystem.RingService.SendElected:input_type -> auctionSystem.Elected
	3,  // 7: auctionSystem.FrontendService.SendBid:input_type -> auctionSystem.Bid
	8,  // 8: auctionSystem.FrontendService.GetResult:input_type -> auctionSystem.Void
	7,  // 9: auctionSystem.FrontendService.SetPrimaryNode:input_type -> auctionSystem.Elected
	0,  // 10: auctionSystem.ConnectService.FinishConnect:output_type -> auctionSystem.BackupDetails
	8,  // 11: auctionSystem.ConnectService.AddBackup:output_type -> auctionSystem.Void
	8,  // 12: auctionSystem.ConnectService.RemoveBackup:output_type -> auctionSystem.Void
	4,  // 13: auctionSystem.AuctionService.SendBid:output_type -> auctionSystem.BidAck
	5,  // 14: auctionSystem.AuctionService.GetResult:output_type -> auctionSystem.Outcome
	8,  // 15: auctionSystem.RingService.SendElection:output_type -> auctionSystem.Void
	8,  // 16: auctionSystem.RingService.SendElected:output_type -> auctionSystem.Void
	4,  // 17: auctionSystem.FrontendService.SendBid:output_type -> auctionSystem.BidAck
	5,  // 18: auctionSystem.FrontendService.GetResult:output_type -> auctionSystem.Outcome
	8,  // 19: auctionSystem.FrontendService.SetPrimaryNode:output_type -> auctionSystem.Void
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auction_auction_proto_init() }
func file_auction_auction_proto_init() {
	if File_auction_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auction_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupDetails); i {
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
		file_auction_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupJoin); i {
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
		file_auction_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupLeave); i {
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
		file_auction_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_auction_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidAck); i {
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
		file_auction_auction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_auction_auction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Election); i {
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
		file_auction_auction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Elected); i {
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
		file_auction_auction_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auction_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_auction_auction_proto_goTypes,
		DependencyIndexes: file_auction_auction_proto_depIdxs,
		MessageInfos:      file_auction_auction_proto_msgTypes,
	}.Build()
	File_auction_auction_proto = out.File
	file_auction_auction_proto_rawDesc = nil
	file_auction_auction_proto_goTypes = nil
	file_auction_auction_proto_depIdxs = nil
}
