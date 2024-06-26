// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: disk.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 文件系统扩容请求消息
type ExpandVolumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume string `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *ExpandVolumeRequest) Reset() {
	*x = ExpandVolumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandVolumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandVolumeRequest) ProtoMessage() {}

func (x *ExpandVolumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandVolumeRequest.ProtoReflect.Descriptor instead.
func (*ExpandVolumeRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{0}
}

func (x *ExpandVolumeRequest) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

var File_disk_proto protoreflect.FileDescriptor

var file_disk_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6b, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x32, 0x6a, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x58, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x2e, 0x6b, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_disk_proto_rawDescOnce sync.Once
	file_disk_proto_rawDescData = file_disk_proto_rawDesc
)

func file_disk_proto_rawDescGZIP() []byte {
	file_disk_proto_rawDescOnce.Do(func() {
		file_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_disk_proto_rawDescData)
	})
	return file_disk_proto_rawDescData
}

var file_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_disk_proto_goTypes = []interface{}{
	(*ExpandVolumeRequest)(nil), // 0: katacontainers.extends.v1.ExpandVolumeRequest
	(*emptypb.Empty)(nil),       // 1: google.protobuf.Empty
}
var file_disk_proto_depIdxs = []int32{
	0, // 0: katacontainers.extends.v1.ExtendedStatus.ExpandVolume:input_type -> katacontainers.extends.v1.ExpandVolumeRequest
	1, // 1: katacontainers.extends.v1.ExtendedStatus.ExpandVolume:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_disk_proto_init() }
func file_disk_proto_init() {
	if File_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandVolumeRequest); i {
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
			RawDescriptor: file_disk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disk_proto_goTypes,
		DependencyIndexes: file_disk_proto_depIdxs,
		MessageInfos:      file_disk_proto_msgTypes,
	}.Build()
	File_disk_proto = out.File
	file_disk_proto_rawDesc = nil
	file_disk_proto_goTypes = nil
	file_disk_proto_depIdxs = nil
}
