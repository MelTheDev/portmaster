// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: reporter.proto

package proto

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

type ReportConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *ReportConnectionRequest) Reset() {
	*x = ReportConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportConnectionRequest) ProtoMessage() {}

func (x *ReportConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportConnectionRequest.ProtoReflect.Descriptor instead.
func (*ReportConnectionRequest) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{0}
}

func (x *ReportConnectionRequest) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

type ReportConnectionRespose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportConnectionRespose) Reset() {
	*x = ReportConnectionRespose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportConnectionRespose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportConnectionRespose) ProtoMessage() {}

func (x *ReportConnectionRespose) ProtoReflect() protoreflect.Message {
	mi := &file_reporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportConnectionRespose.ProtoReflect.Descriptor instead.
func (*ReportConnectionRespose) Descriptor() ([]byte, []int) {
	return file_reporter_proto_rawDescGZIP(), []int{1}
}

var File_reporter_proto protoreflect.FileDescriptor

var file_reporter_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73,
	0x65, 0x32, 0x9a, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x2e, 0x73, 0x61, 0x66,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_reporter_proto_rawDescOnce sync.Once
	file_reporter_proto_rawDescData = file_reporter_proto_rawDesc
)

func file_reporter_proto_rawDescGZIP() []byte {
	file_reporter_proto_rawDescOnce.Do(func() {
		file_reporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_reporter_proto_rawDescData)
	})
	return file_reporter_proto_rawDescData
}

var file_reporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_reporter_proto_goTypes = []interface{}{
	(*ReportConnectionRequest)(nil), // 0: safing.portmaster.plugin.proto.ReportConnectionRequest
	(*ReportConnectionRespose)(nil), // 1: safing.portmaster.plugin.proto.ReportConnectionRespose
	(*Connection)(nil),              // 2: safing.portmaster.plugin.proto.Connection
}
var file_reporter_proto_depIdxs = []int32{
	2, // 0: safing.portmaster.plugin.proto.ReportConnectionRequest.connection:type_name -> safing.portmaster.plugin.proto.Connection
	0, // 1: safing.portmaster.plugin.proto.ReporterService.ReportConnection:input_type -> safing.portmaster.plugin.proto.ReportConnectionRequest
	1, // 2: safing.portmaster.plugin.proto.ReporterService.ReportConnection:output_type -> safing.portmaster.plugin.proto.ReportConnectionRespose
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reporter_proto_init() }
func file_reporter_proto_init() {
	if File_reporter_proto != nil {
		return
	}
	file_network_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_reporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportConnectionRequest); i {
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
		file_reporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportConnectionRespose); i {
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
			RawDescriptor: file_reporter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reporter_proto_goTypes,
		DependencyIndexes: file_reporter_proto_depIdxs,
		MessageInfos:      file_reporter_proto_msgTypes,
	}.Build()
	File_reporter_proto = out.File
	file_reporter_proto_rawDesc = nil
	file_reporter_proto_goTypes = nil
	file_reporter_proto_depIdxs = nil
}
