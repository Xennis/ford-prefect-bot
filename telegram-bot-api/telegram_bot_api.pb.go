// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telegram_bot_api.proto

package telegrambotapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Updates struct {
	Ok                   bool      `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Result               []*Update `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Updates) Reset()         { *m = Updates{} }
func (m *Updates) String() string { return proto.CompactTextString(m) }
func (*Updates) ProtoMessage()    {}
func (*Updates) Descriptor() ([]byte, []int) {
	return fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe, []int{0}
}
func (m *Updates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Updates.Unmarshal(m, b)
}
func (m *Updates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Updates.Marshal(b, m, deterministic)
}
func (dst *Updates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Updates.Merge(dst, src)
}
func (m *Updates) XXX_Size() int {
	return xxx_messageInfo_Updates.Size(m)
}
func (m *Updates) XXX_DiscardUnknown() {
	xxx_messageInfo_Updates.DiscardUnknown(m)
}

var xxx_messageInfo_Updates proto.InternalMessageInfo

func (m *Updates) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Updates) GetResult() []*Update {
	if m != nil {
		return m.Result
	}
	return nil
}

type Update struct {
	UpdateId             int32    `protobuf:"varint,1,opt,name=update_id,json=updateId" json:"update_id,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe, []int{1}
}
func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (dst *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(dst, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

func (m *Update) GetUpdateId() int32 {
	if m != nil {
		return m.UpdateId
	}
	return 0
}

func (m *Update) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Message struct {
	MessageId            int32    `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	From                 *User    `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	User                 *Chat    `protobuf:"bytes,3,opt,name=User" json:"User,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
	Text                 string   `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe, []int{2}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Message) GetFrom() *User {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Message) GetUser() *Chat {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Message) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	IsBot                bool     `protobuf:"varint,2,opt,name=is_bot,json=isBot" json:"is_bot,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
	LanguageCode         string   `protobuf:"bytes,6,opt,name=language_code,json=languageCode" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetIsBot() bool {
	if m != nil {
		return m.IsBot
	}
	return false
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

type Chat struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe, []int{4}
}
func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (dst *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(dst, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Chat) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Chat) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Updates)(nil), "telegrambotapi.Updates")
	proto.RegisterType((*Update)(nil), "telegrambotapi.Update")
	proto.RegisterType((*Message)(nil), "telegrambotapi.Message")
	proto.RegisterType((*User)(nil), "telegrambotapi.User")
	proto.RegisterType((*Chat)(nil), "telegrambotapi.Chat")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TelegramBotApiClient is the client API for TelegramBotApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelegramBotApiClient interface {
	ReceiveUpdates(ctx context.Context, in *Update, opts ...grpc.CallOption) (*empty.Empty, error)
}

type telegramBotApiClient struct {
	cc *grpc.ClientConn
}

func NewTelegramBotApiClient(cc *grpc.ClientConn) TelegramBotApiClient {
	return &telegramBotApiClient{cc}
}

func (c *telegramBotApiClient) ReceiveUpdates(ctx context.Context, in *Update, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telegrambotapi.TelegramBotApi/ReceiveUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramBotApiServer is the server API for TelegramBotApi service.
type TelegramBotApiServer interface {
	ReceiveUpdates(context.Context, *Update) (*empty.Empty, error)
}

func RegisterTelegramBotApiServer(s *grpc.Server, srv TelegramBotApiServer) {
	s.RegisterService(&_TelegramBotApi_serviceDesc, srv)
}

func _TelegramBotApi_ReceiveUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Update)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramBotApiServer).ReceiveUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telegrambotapi.TelegramBotApi/ReceiveUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramBotApiServer).ReceiveUpdates(ctx, req.(*Update))
	}
	return interceptor(ctx, in, info, handler)
}

var _TelegramBotApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telegrambotapi.TelegramBotApi",
	HandlerType: (*TelegramBotApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveUpdates",
			Handler:    _TelegramBotApi_ReceiveUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telegram_bot_api.proto",
}

func init() {
	proto.RegisterFile("telegram_bot_api.proto", fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe)
}

var fileDescriptor_telegram_bot_api_19de9b7cd93c3ffe = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xdd, 0xc4, 0xb1, 0xa7, 0xc5, 0x87, 0x55, 0x09, 0x56, 0x0a, 0x52, 0x64, 0x2e, 0x3e,
	0x39, 0x22, 0x7c, 0x01, 0xad, 0x90, 0xc8, 0x01, 0x0e, 0x2b, 0x40, 0xdc, 0xa2, 0x4d, 0x3d, 0x71,
	0x57, 0xb1, 0x3d, 0x96, 0x77, 0x8d, 0xe8, 0x95, 0x5f, 0xe0, 0x17, 0x38, 0xf1, 0x3b, 0xfc, 0x02,
	0x1f, 0x82, 0x76, 0xd7, 0x2e, 0x6a, 0x43, 0x6f, 0x33, 0xef, 0x3d, 0xcf, 0xbe, 0x19, 0x3f, 0x98,
	0x6b, 0xac, 0xb0, 0xec, 0x44, 0xbd, 0xdd, 0x91, 0xde, 0x8a, 0x56, 0xe6, 0x6d, 0x47, 0x9a, 0x58,
	0x3c, 0xe2, 0x3b, 0xd2, 0xa2, 0x95, 0x8b, 0xe7, 0x25, 0x51, 0x59, 0xe1, 0x4a, 0xb4, 0x72, 0x25,
	0x9a, 0x86, 0xb4, 0xd0, 0x92, 0x1a, 0xe5, 0xd4, 0x8b, 0x8b, 0x81, 0xb5, 0xdd, 0xae, 0xdf, 0xaf,
	0xb0, 0x6e, 0xf5, 0xad, 0x23, 0xd3, 0x0d, 0xcc, 0x3e, 0xb5, 0x85, 0xd0, 0xa8, 0x58, 0x0c, 0x3e,
	0x1d, 0x12, 0x6f, 0xe9, 0x65, 0x21, 0xf7, 0xe9, 0xc0, 0x72, 0x08, 0x3a, 0x54, 0x7d, 0xa5, 0x13,
	0x7f, 0x79, 0x92, 0x9d, 0xae, 0xe7, 0xf9, 0xfd, 0x67, 0x73, 0xf7, 0x21, 0x1f, 0x54, 0xe9, 0x17,
	0x08, 0x1c, 0xc2, 0x2e, 0x20, 0xea, 0x6d, 0xb5, 0x95, 0x85, 0x1d, 0x38, 0xe5, 0xa1, 0x03, 0x36,
	0x05, 0x7b, 0x05, 0xb3, 0x1a, 0x95, 0x12, 0x25, 0x26, 0xfe, 0xd2, 0xcb, 0x4e, 0xd7, 0xcf, 0x1e,
	0xce, 0x7d, 0xef, 0x68, 0x3e, 0xea, 0xd2, 0x9f, 0x1e, 0xcc, 0x06, 0x90, 0xbd, 0x00, 0x18, 0xe0,
	0x7f, 0xc3, 0xa3, 0x01, 0xd9, 0x14, 0x2c, 0x83, 0xc9, 0xbe, 0xa3, 0x7a, 0x18, 0x7d, 0x7e, 0x64,
	0x59, 0x61, 0xc7, 0xad, 0xc2, 0x28, 0x4d, 0x97, 0x9c, 0xfc, 0x5f, 0x79, 0x75, 0x23, 0x34, 0xb7,
	0x0a, 0xc6, 0x60, 0x62, 0xbc, 0x27, 0x93, 0xa5, 0x97, 0x45, 0xdc, 0xd6, 0x06, 0xd3, 0xf8, 0x4d,
	0x27, 0x53, 0x87, 0x99, 0x3a, 0xfd, 0xe5, 0xb9, 0x91, 0xe6, 0x92, 0x77, 0xde, 0x7c, 0x59, 0xb0,
	0xa7, 0x10, 0x48, 0x65, 0xfe, 0xa1, 0xb5, 0x15, 0xf2, 0xa9, 0x54, 0x97, 0xa4, 0xcd, 0x2a, 0x7b,
	0xd9, 0x29, 0xbd, 0x6d, 0x44, 0x8d, 0xd6, 0x47, 0xc4, 0x23, 0x8b, 0x7c, 0x10, 0xb5, 0xbd, 0x62,
	0x25, 0x46, 0xd6, 0xbd, 0x1d, 0x1a, 0xc0, 0x92, 0x0b, 0x08, 0x7b, 0x85, 0x9d, 0xe5, 0x9c, 0x87,
	0xbb, 0x9e, 0xbd, 0x84, 0x27, 0x95, 0x68, 0xca, 0xde, 0xdc, 0xe8, 0x9a, 0x0a, 0x4c, 0x02, 0x2b,
	0x38, 0x1b, 0xc1, 0x2b, 0x2a, 0x30, 0xdd, 0xc0, 0xc4, 0xac, 0x78, 0xe4, 0xf5, 0xbe, 0x29, 0xff,
	0xa1, 0x29, 0xb3, 0xf7, 0x6d, 0x3b, 0xba, 0xb5, 0xf5, 0xfa, 0x06, 0xe2, 0x8f, 0xc3, 0xf1, 0x2e,
	0x49, 0xbf, 0x69, 0x25, 0xfb, 0x0c, 0x31, 0xc7, 0x6b, 0x94, 0x5f, 0x71, 0x0c, 0xd7, 0x23, 0xe1,
	0x59, 0xcc, 0x73, 0x97, 0xce, 0x7c, 0x4c, 0x67, 0xfe, 0xd6, 0xa4, 0x33, 0x3d, 0xff, 0xfe, 0xfb,
	0xcf, 0x0f, 0x3f, 0x4e, 0xcf, 0x56, 0x2e, 0x38, 0xea, 0x1d, 0xd1, 0x61, 0x17, 0x58, 0xd5, 0xeb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0x5e, 0x2a, 0x2e, 0x19, 0x03, 0x00, 0x00,
}