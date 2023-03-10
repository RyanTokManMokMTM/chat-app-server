// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: socket-proto/socket.proto

package socket_message

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar       string `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`             //user avatar path
	FromUserName string `protobuf:"bytes,2,opt,name=fromUserName,proto3" json:"fromUserName,omitempty"` //sender user name
	FromUUID     string `protobuf:"bytes,3,opt,name=fromUUID,proto3" json:"fromUUID,omitempty"`         //sender uuid
	ToUUID       string `protobuf:"bytes,4,opt,name=toUUID,proto3" json:"toUUID,omitempty"`             //receiver uuid
	Content      string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`           //sending content
	ContentType  int32  `protobuf:"varint,6,opt,name=contentType,proto3" json:"contentType,omitempty"`  //sending content type. For example 1: text, 2: file, 3: audio, 4: video....
	Type         int32  `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`                //For example: "heatbeat" for checking server/client health , video call/audio call ->"webrtc"
	MessageType  int32  `protobuf:"varint,8,opt,name=messageType,proto3" json:"messageType,omitempty"`  //1: single 2: group
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_socket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_socket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_socket_proto_socket_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Message) GetFromUserName() string {
	if x != nil {
		return x.FromUserName
	}
	return ""
}

func (x *Message) GetFromUUID() string {
	if x != nil {
		return x.FromUUID
	}
	return ""
}

func (x *Message) GetToUUID() string {
	if x != nil {
		return x.ToUUID
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *Message) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Message) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

var File_socket_proto_socket_proto protoreflect.FileDescriptor

var file_socket_proto_socket_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xeb, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_socket_proto_socket_proto_rawDescOnce sync.Once
	file_socket_proto_socket_proto_rawDescData = file_socket_proto_socket_proto_rawDesc
)

func file_socket_proto_socket_proto_rawDescGZIP() []byte {
	file_socket_proto_socket_proto_rawDescOnce.Do(func() {
		file_socket_proto_socket_proto_rawDescData = protoimpl.X.CompressGZIP(file_socket_proto_socket_proto_rawDescData)
	})
	return file_socket_proto_socket_proto_rawDescData
}

var file_socket_proto_socket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_socket_proto_socket_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: message.Message
}
var file_socket_proto_socket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_socket_proto_socket_proto_init() }
func file_socket_proto_socket_proto_init() {
	if File_socket_proto_socket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_socket_proto_socket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_socket_proto_socket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_socket_proto_socket_proto_goTypes,
		DependencyIndexes: file_socket_proto_socket_proto_depIdxs,
		MessageInfos:      file_socket_proto_socket_proto_msgTypes,
	}.Build()
	File_socket_proto_socket_proto = out.File
	file_socket_proto_socket_proto_rawDesc = nil
	file_socket_proto_socket_proto_goTypes = nil
	file_socket_proto_socket_proto_depIdxs = nil
}
