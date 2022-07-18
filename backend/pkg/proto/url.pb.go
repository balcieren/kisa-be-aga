// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: url.proto

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

type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShortId   string `protobuf:"bytes,2,opt,name=short_id,json=shortId,proto3" json:"short_id,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Url) GetShortId() string {
	if x != nil {
		return x.ShortId
	}
	return ""
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Url) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateUrlRequest) Reset() {
	*x = CreateUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlRequest) ProtoMessage() {}

func (x *CreateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateUrlRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId string `protobuf:"bytes,1,opt,name=short_id,json=shortId,proto3" json:"short_id,omitempty"`
}

func (x *GetUrlRequest) Reset() {
	*x = GetUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlRequest) ProtoMessage() {}

func (x *GetUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlRequest.ProtoReflect.Descriptor instead.
func (*GetUrlRequest) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{2}
}

func (x *GetUrlRequest) GetShortId() string {
	if x != nil {
		return x.ShortId
	}
	return ""
}

var File_url_proto protoreflect.FileDescriptor

var file_url_proto_rawDesc = []byte{
	0x0a, 0x09, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x22, 0x61, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x24, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2a, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x32, 0x64, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x72, 0x6c, 0x42, 0x13, 0x5a,
	0x11, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_proto_rawDescOnce sync.Once
	file_url_proto_rawDescData = file_url_proto_rawDesc
)

func file_url_proto_rawDescGZIP() []byte {
	file_url_proto_rawDescOnce.Do(func() {
		file_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_proto_rawDescData)
	})
	return file_url_proto_rawDescData
}

var file_url_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_url_proto_goTypes = []interface{}{
	(*Url)(nil),              // 0: entpb.Url
	(*CreateUrlRequest)(nil), // 1: entpb.CreateUrlRequest
	(*GetUrlRequest)(nil),    // 2: entpb.GetUrlRequest
}
var file_url_proto_depIdxs = []int32{
	1, // 0: entpb.UrlService.Create:input_type -> entpb.CreateUrlRequest
	2, // 1: entpb.UrlService.Get:input_type -> entpb.GetUrlRequest
	0, // 2: entpb.UrlService.Create:output_type -> entpb.Url
	0, // 3: entpb.UrlService.Get:output_type -> entpb.Url
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_proto_init() }
func file_url_proto_init() {
	if File_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_url_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlRequest); i {
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
		file_url_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlRequest); i {
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
			RawDescriptor: file_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_proto_goTypes,
		DependencyIndexes: file_url_proto_depIdxs,
		MessageInfos:      file_url_proto_msgTypes,
	}.Build()
	File_url_proto = out.File
	file_url_proto_rawDesc = nil
	file_url_proto_goTypes = nil
	file_url_proto_depIdxs = nil
}
