// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: zenia/authz/v1/authz_service.proto

package authz

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object    *Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Relation  string  `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	SubjectId string  `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// optional but highly recommend
	Zookie string `protobuf:"bytes,4,opt,name=zookie,proto3" json:"zookie,omitempty"`
	// follow subjectsets
	Expand bool `protobuf:"varint,5,opt,name=expand,proto3" json:"expand,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zenia_authz_v1_authz_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zenia_authz_v1_authz_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_zenia_authz_v1_authz_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *CheckRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *CheckRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *CheckRequest) GetZookie() string {
	if x != nil {
		return x.Zookie
	}
	return ""
}

func (x *CheckRequest) GetExpand() bool {
	if x != nil {
		return x.Expand
	}
	return false
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowed bool `protobuf:"varint,1,opt,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zenia_authz_v1_authz_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zenia_authz_v1_authz_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_zenia_authz_v1_authz_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetAllowed() bool {
	if x != nil {
		return x.Allowed
	}
	return false
}

var File_zenia_authz_v1_authz_service_proto protoreflect.FileDescriptor

var file_zenia_authz_v1_authz_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x7a, 0x65, 0x6e, 0x69, 0x61,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x7a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x32, 0x7e,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x1c, 0x2e, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x20, 0xca, 0x41,
	0x1d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62,
	0x69, 0x6e, 0x62, 0x72, 0x61, 0x65, 0x6d, 0x65, 0x72, 0x2f, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x7a, 0x65, 0x6e, 0x69, 0x61, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zenia_authz_v1_authz_service_proto_rawDescOnce sync.Once
	file_zenia_authz_v1_authz_service_proto_rawDescData = file_zenia_authz_v1_authz_service_proto_rawDesc
)

func file_zenia_authz_v1_authz_service_proto_rawDescGZIP() []byte {
	file_zenia_authz_v1_authz_service_proto_rawDescOnce.Do(func() {
		file_zenia_authz_v1_authz_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_zenia_authz_v1_authz_service_proto_rawDescData)
	})
	return file_zenia_authz_v1_authz_service_proto_rawDescData
}

var file_zenia_authz_v1_authz_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_zenia_authz_v1_authz_service_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),  // 0: zenia.authz.v1.CheckRequest
	(*CheckResponse)(nil), // 1: zenia.authz.v1.CheckResponse
	(*Object)(nil),        // 2: zenia.authz.v1.Object
}
var file_zenia_authz_v1_authz_service_proto_depIdxs = []int32{
	2, // 0: zenia.authz.v1.CheckRequest.object:type_name -> zenia.authz.v1.Object
	0, // 1: zenia.authz.v1.AuthorizationService.Check:input_type -> zenia.authz.v1.CheckRequest
	1, // 2: zenia.authz.v1.AuthorizationService.Check:output_type -> zenia.authz.v1.CheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zenia_authz_v1_authz_service_proto_init() }
func file_zenia_authz_v1_authz_service_proto_init() {
	if File_zenia_authz_v1_authz_service_proto != nil {
		return
	}
	file_zenia_authz_v1_acl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zenia_authz_v1_authz_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_zenia_authz_v1_authz_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
			RawDescriptor: file_zenia_authz_v1_authz_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zenia_authz_v1_authz_service_proto_goTypes,
		DependencyIndexes: file_zenia_authz_v1_authz_service_proto_depIdxs,
		MessageInfos:      file_zenia_authz_v1_authz_service_proto_msgTypes,
	}.Build()
	File_zenia_authz_v1_authz_service_proto = out.File
	file_zenia_authz_v1_authz_service_proto_rawDesc = nil
	file_zenia_authz_v1_authz_service_proto_goTypes = nil
	file_zenia_authz_v1_authz_service_proto_depIdxs = nil
}
