// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: header.proto

package header

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

// 压缩类型，Raw不压缩数据
type Compress int32

const (
	Compress_Raw    Compress = 0
	Compress_Gzip   Compress = 1
	Compress_Snappy Compress = 2
	Compress_Zlib   Compress = 3
)

// Enum value maps for Compress.
var (
	Compress_name = map[int32]string{
		0: "Raw",
		1: "Gzip",
		2: "Snappy",
		3: "Zlib",
	}
	Compress_value = map[string]int32{
		"Raw":    0,
		"Gzip":   1,
		"Snappy": 2,
		"Zlib":   3,
	}
)

func (x Compress) Enum() *Compress {
	p := new(Compress)
	*p = x
	return p
}

func (x Compress) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Compress) Descriptor() protoreflect.EnumDescriptor {
	return file_header_proto_enumTypes[0].Descriptor()
}

func (Compress) Type() protoreflect.EnumType {
	return &file_header_proto_enumTypes[0]
}

func (x Compress) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Compress.Descriptor instead.
func (Compress) EnumDescriptor() ([]byte, []int) {
	return file_header_proto_rawDescGZIP(), []int{0}
}

// 请求头信息
type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 压缩类型
	CompressType Compress `protobuf:"varint,1,opt,name=compress_type,json=compressType,proto3,enum=header.Compress" json:"compress_type,omitempty"`
	// 远程调用过程方法名称
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// rpc调用序号
	Id uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 数据长度
	RequestLen uint32 `protobuf:"varint,4,opt,name=request_len,json=requestLen,proto3" json:"request_len,omitempty"`
	// 检验和
	Checksum uint32 `protobuf:"varint,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_header_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetCompressType() Compress {
	if x != nil {
		return x.CompressType
	}
	return Compress_Raw
}

func (x *RequestHeader) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RequestHeader) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RequestHeader) GetRequestLen() uint32 {
	if x != nil {
		return x.RequestLen
	}
	return 0
}

func (x *RequestHeader) GetChecksum() uint32 {
	if x != nil {
		return x.Checksum
	}
	return 0
}

// 响应头信息
type ResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 压缩类型
	CompressType Compress `protobuf:"varint,1,opt,name=compress_type,json=compressType,proto3,enum=header.Compress" json:"compress_type,omitempty"`
	// rpc调用序号
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// 远程调用过程发生错误信息，若为空则不发生错误
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// 数据长度
	ResponseLen uint32 `protobuf:"varint,4,opt,name=response_len,json=responseLen,proto3" json:"response_len,omitempty"`
	// 检验和
	Checksum uint32 `protobuf:"varint,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_header_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeader) ProtoMessage() {}

func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_header_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeader.ProtoReflect.Descriptor instead.
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_header_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseHeader) GetCompressType() Compress {
	if x != nil {
		return x.CompressType
	}
	return Compress_Raw
}

func (x *ResponseHeader) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponseHeader) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ResponseHeader) GetResponseLen() uint32 {
	if x != nil {
		return x.ResponseLen
	}
	return 0
}

func (x *ResponseHeader) GetChecksum() uint32 {
	if x != nil {
		return x.Checksum
	}
	return 0
}

var File_header_proto protoreflect.FileDescriptor

var file_header_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x75, 0x6d, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x2a, 0x33, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x07, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x7a, 0x69, 0x70,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6e, 0x61, 0x70, 0x70, 0x79, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x5a, 0x6c, 0x69, 0x62, 0x10, 0x03, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_header_proto_rawDescOnce sync.Once
	file_header_proto_rawDescData = file_header_proto_rawDesc
)

func file_header_proto_rawDescGZIP() []byte {
	file_header_proto_rawDescOnce.Do(func() {
		file_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_header_proto_rawDescData)
	})
	return file_header_proto_rawDescData
}

var file_header_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_header_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_header_proto_goTypes = []interface{}{
	(Compress)(0),          // 0: header.Compress
	(*RequestHeader)(nil),  // 1: header.RequestHeader
	(*ResponseHeader)(nil), // 2: header.ResponseHeader
}
var file_header_proto_depIdxs = []int32{
	0, // 0: header.RequestHeader.compress_type:type_name -> header.Compress
	0, // 1: header.ResponseHeader.compress_type:type_name -> header.Compress
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_header_proto_init() }
func file_header_proto_init() {
	if File_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_header_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHeader); i {
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
			RawDescriptor: file_header_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_header_proto_goTypes,
		DependencyIndexes: file_header_proto_depIdxs,
		EnumInfos:         file_header_proto_enumTypes,
		MessageInfos:      file_header_proto_msgTypes,
	}.Build()
	File_header_proto = out.File
	file_header_proto_rawDesc = nil
	file_header_proto_goTypes = nil
	file_header_proto_depIdxs = nil
}
