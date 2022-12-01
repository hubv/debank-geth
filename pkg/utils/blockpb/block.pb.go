// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: pkg/utils/pb/block.proto

package blockpb

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

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNum  int64  `protobuf:"varint,1,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	BlockHash string `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	LastTime  int64  `protobuf:"varint,3,opt,name=last_time,json=lastTime,proto3" json:"last_time,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_utils_pb_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_utils_pb_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_pkg_utils_pb_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockInfo) GetBlockNum() int64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *BlockInfo) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockInfo) GetLastTime() int64 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

type BatchItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchItem) Reset() {
	*x = BatchItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_utils_pb_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchItem) ProtoMessage() {}

func (x *BatchItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_utils_pb_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchItem.ProtoReflect.Descriptor instead.
func (*BatchItem) Descriptor() ([]byte, []int) {
	return file_pkg_utils_pb_block_proto_rawDescGZIP(), []int{1}
}

func (x *BatchItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BatchItem) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlockData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info       *BlockInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	BatchItems []*BatchItem `protobuf:"bytes,2,rep,name=batch_items,json=batchItems,proto3" json:"batch_items,omitempty"`
}

func (x *BlockData) Reset() {
	*x = BlockData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_utils_pb_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockData) ProtoMessage() {}

func (x *BlockData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_utils_pb_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockData.ProtoReflect.Descriptor instead.
func (*BlockData) Descriptor() ([]byte, []int) {
	return file_pkg_utils_pb_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockData) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BlockData) GetBatchItems() []*BatchItem {
	if x != nil {
		return x.BatchItems
	}
	return nil
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info       *BlockInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	BatchItems []*BatchItem `protobuf:"bytes,2,rep,name=batch_items,json=batchItems,proto3" json:"batch_items,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_utils_pb_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_utils_pb_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_pkg_utils_pb_block_proto_rawDescGZIP(), []int{3}
}

func (x *BlockHeader) GetInfo() *BlockInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BlockHeader) GetBatchItems() []*BatchItem {
	if x != nil {
		return x.BatchItems
	}
	return nil
}

var File_pkg_utils_pb_block_proto protoreflect.FileDescriptor

var file_pkg_utils_pb_block_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x70, 0x62, 0x22, 0x64, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x09, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x09, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x33, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x6a, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x13, 0x5a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_utils_pb_block_proto_rawDescOnce sync.Once
	file_pkg_utils_pb_block_proto_rawDescData = file_pkg_utils_pb_block_proto_rawDesc
)

func file_pkg_utils_pb_block_proto_rawDescGZIP() []byte {
	file_pkg_utils_pb_block_proto_rawDescOnce.Do(func() {
		file_pkg_utils_pb_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_utils_pb_block_proto_rawDescData)
	})
	return file_pkg_utils_pb_block_proto_rawDescData
}

var file_pkg_utils_pb_block_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_utils_pb_block_proto_goTypes = []interface{}{
	(*BlockInfo)(nil),   // 0: blockpb.BlockInfo
	(*BatchItem)(nil),   // 1: blockpb.BatchItem
	(*BlockData)(nil),   // 2: blockpb.BlockData
	(*BlockHeader)(nil), // 3: blockpb.BlockHeader
}
var file_pkg_utils_pb_block_proto_depIdxs = []int32{
	0, // 0: blockpb.BlockData.info:type_name -> blockpb.BlockInfo
	1, // 1: blockpb.BlockData.batch_items:type_name -> blockpb.BatchItem
	0, // 2: blockpb.BlockHeader.info:type_name -> blockpb.BlockInfo
	1, // 3: blockpb.BlockHeader.batch_items:type_name -> blockpb.BatchItem
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_utils_pb_block_proto_init() }
func file_pkg_utils_pb_block_proto_init() {
	if File_pkg_utils_pb_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_utils_pb_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
		file_pkg_utils_pb_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchItem); i {
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
		file_pkg_utils_pb_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockData); i {
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
		file_pkg_utils_pb_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
			RawDescriptor: file_pkg_utils_pb_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_utils_pb_block_proto_goTypes,
		DependencyIndexes: file_pkg_utils_pb_block_proto_depIdxs,
		MessageInfos:      file_pkg_utils_pb_block_proto_msgTypes,
	}.Build()
	File_pkg_utils_pb_block_proto = out.File
	file_pkg_utils_pb_block_proto_rawDesc = nil
	file_pkg_utils_pb_block_proto_goTypes = nil
	file_pkg_utils_pb_block_proto_depIdxs = nil
}
