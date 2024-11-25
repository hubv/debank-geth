// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/store.proto

package pb

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

type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *BlockInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	NoCache bool       `protobuf:"varint,2,opt,name=no_cache,json=noCache,proto3" json:"no_cache,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockRequest) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetBlockRequest) GetNoCache() bool {
	if x != nil {
		return x.NoCache
	}
	return false
}

type BlockChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *BlockInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Chunk []byte     `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *BlockChunk) Reset() {
	*x = BlockChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockChunk) ProtoMessage() {}

func (x *BlockChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockChunk.ProtoReflect.Descriptor instead.
func (*BlockChunk) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{1}
}

func (x *BlockChunk) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BlockChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type PutBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutBlockReply) Reset() {
	*x = PutBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBlockReply) ProtoMessage() {}

func (x *PutBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBlockReply.ProtoReflect.Descriptor instead.
func (*PutBlockReply) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{2}
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetFileReply) Reset() {
	*x = GetFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileReply) ProtoMessage() {}

func (x *GetFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileReply.ProtoReflect.Descriptor instead.
func (*GetFileReply) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PutFileRequest) Reset() {
	*x = PutFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileRequest) ProtoMessage() {}

func (x *PutFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileRequest.ProtoReflect.Descriptor instead.
func (*PutFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{5}
}

func (x *PutFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PutFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutFileReply) Reset() {
	*x = PutFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutFileReply) ProtoMessage() {}

func (x *PutFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutFileReply.ProtoReflect.Descriptor instead.
func (*PutFileReply) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{6}
}

type ListHeaderStartAtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env            string `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	ChainId        string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Role           string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	BlockNum       int64  `protobuf:"varint,4,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	CountNum       int64  `protobuf:"varint,5,opt,name=count_num,json=countNum,proto3" json:"count_num,omitempty"`
	AfterMsgOffset int64  `protobuf:"varint,6,opt,name=after_msg_offset,json=afterMsgOffset,proto3" json:"after_msg_offset,omitempty"`
}

func (x *ListHeaderStartAtRequest) Reset() {
	*x = ListHeaderStartAtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHeaderStartAtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHeaderStartAtRequest) ProtoMessage() {}

func (x *ListHeaderStartAtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHeaderStartAtRequest.ProtoReflect.Descriptor instead.
func (*ListHeaderStartAtRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{7}
}

func (x *ListHeaderStartAtRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *ListHeaderStartAtRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ListHeaderStartAtRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ListHeaderStartAtRequest) GetBlockNum() int64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *ListHeaderStartAtRequest) GetCountNum() int64 {
	if x != nil {
		return x.CountNum
	}
	return 0
}

func (x *ListHeaderStartAtRequest) GetAfterMsgOffset() int64 {
	if x != nil {
		return x.AfterMsgOffset
	}
	return 0
}

type ListHeaderStartAtReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*BlockInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *ListHeaderStartAtReply) Reset() {
	*x = ListHeaderStartAtReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHeaderStartAtReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHeaderStartAtReply) ProtoMessage() {}

func (x *ListHeaderStartAtReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHeaderStartAtReply.ProtoReflect.Descriptor instead.
func (*ListHeaderStartAtReply) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{8}
}

func (x *ListHeaderStartAtReply) GetInfos() []*BlockInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type RemoveFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*BlockInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *RemoveFilesRequest) Reset() {
	*x = RemoveFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFilesRequest) ProtoMessage() {}

func (x *RemoveFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFilesRequest.ProtoReflect.Descriptor instead.
func (*RemoveFilesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveFilesRequest) GetInfos() []*BlockInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type RemoveFilesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveFilesReply) Reset() {
	*x = RemoveFilesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_store_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFilesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFilesReply) ProtoMessage() {}

func (x *RemoveFilesReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_store_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFilesReply.ProtoReflect.Descriptor instead.
func (*RemoveFilesReply) Descriptor() ([]byte, []int) {
	return file_pkg_pb_store_proto_rawDescGZIP(), []int{10}
}

var File_pkg_pb_store_proto protoreflect.FileDescriptor

var file_pkg_pb_store_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x45, 0x0a,
	0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x22, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x38, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3d, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x39, 0x0a, 0x12, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xe7, 0x02, 0x0a, 0x07, 0x53, 0x33,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x75,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x07, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x44, 0x65, 0x46, 0x69, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x78, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_pb_store_proto_rawDescOnce sync.Once
	file_pkg_pb_store_proto_rawDescData = file_pkg_pb_store_proto_rawDesc
)

func file_pkg_pb_store_proto_rawDescGZIP() []byte {
	file_pkg_pb_store_proto_rawDescOnce.Do(func() {
		file_pkg_pb_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_store_proto_rawDescData)
	})
	return file_pkg_pb_store_proto_rawDescData
}

var file_pkg_pb_store_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_store_proto_goTypes = []interface{}{
	(*GetBlockRequest)(nil),          // 0: pb.GetBlockRequest
	(*BlockChunk)(nil),               // 1: pb.BlockChunk
	(*PutBlockReply)(nil),            // 2: pb.PutBlockReply
	(*GetFileRequest)(nil),           // 3: pb.GetFileRequest
	(*GetFileReply)(nil),             // 4: pb.GetFileReply
	(*PutFileRequest)(nil),           // 5: pb.PutFileRequest
	(*PutFileReply)(nil),             // 6: pb.PutFileReply
	(*ListHeaderStartAtRequest)(nil), // 7: pb.ListHeaderStartAtRequest
	(*ListHeaderStartAtReply)(nil),   // 8: pb.ListHeaderStartAtReply
	(*RemoveFilesRequest)(nil),       // 9: pb.RemoveFilesRequest
	(*RemoveFilesReply)(nil),         // 10: pb.RemoveFilesReply
	(*BlockInfo)(nil),                // 11: pb.BlockInfo
}
var file_pkg_pb_store_proto_depIdxs = []int32{
	11, // 0: pb.GetBlockRequest.info:type_name -> pb.BlockInfo
	11, // 1: pb.BlockChunk.info:type_name -> pb.BlockInfo
	11, // 2: pb.ListHeaderStartAtReply.infos:type_name -> pb.BlockInfo
	11, // 3: pb.RemoveFilesRequest.infos:type_name -> pb.BlockInfo
	0,  // 4: pb.S3Proxy.GetBlock:input_type -> pb.GetBlockRequest
	1,  // 5: pb.S3Proxy.PutBlock:input_type -> pb.BlockChunk
	3,  // 6: pb.S3Proxy.GetFile:input_type -> pb.GetFileRequest
	5,  // 7: pb.S3Proxy.PutFile:input_type -> pb.PutFileRequest
	7,  // 8: pb.S3Proxy.ListHeaderStartAt:input_type -> pb.ListHeaderStartAtRequest
	9,  // 9: pb.S3Proxy.RemoveFiles:input_type -> pb.RemoveFilesRequest
	1,  // 10: pb.S3Proxy.GetBlock:output_type -> pb.BlockChunk
	2,  // 11: pb.S3Proxy.PutBlock:output_type -> pb.PutBlockReply
	4,  // 12: pb.S3Proxy.GetFile:output_type -> pb.GetFileReply
	6,  // 13: pb.S3Proxy.PutFile:output_type -> pb.PutFileReply
	8,  // 14: pb.S3Proxy.ListHeaderStartAt:output_type -> pb.ListHeaderStartAtReply
	10, // 15: pb.S3Proxy.RemoveFiles:output_type -> pb.RemoveFilesReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_store_proto_init() }
func file_pkg_pb_store_proto_init() {
	if File_pkg_pb_store_proto != nil {
		return
	}
	file_pkg_pb_block_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
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
		file_pkg_pb_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockChunk); i {
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
		file_pkg_pb_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutBlockReply); i {
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
		file_pkg_pb_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_pkg_pb_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileReply); i {
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
		file_pkg_pb_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileRequest); i {
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
		file_pkg_pb_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutFileReply); i {
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
		file_pkg_pb_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHeaderStartAtRequest); i {
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
		file_pkg_pb_store_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHeaderStartAtReply); i {
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
		file_pkg_pb_store_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFilesRequest); i {
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
		file_pkg_pb_store_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFilesReply); i {
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
			RawDescriptor: file_pkg_pb_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_store_proto_goTypes,
		DependencyIndexes: file_pkg_pb_store_proto_depIdxs,
		MessageInfos:      file_pkg_pb_store_proto_msgTypes,
	}.Build()
	File_pkg_pb_store_proto = out.File
	file_pkg_pb_store_proto_rawDesc = nil
	file_pkg_pb_store_proto_goTypes = nil
	file_pkg_pb_store_proto_depIdxs = nil
}
