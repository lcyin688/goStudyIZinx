// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: msg.proto

package msg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgId int32

const (
	MsgId_MSG_Default       MsgId = 0
	MsgId_MSG_CS_Ping       MsgId = 1   // 心跳请求
	MsgId_MSG_SC_Pong       MsgId = 2   // 心跳响应
	MsgId_MSG_CS_Login      MsgId = 101 // 登录请求
	MsgId_MSG_SC_Login      MsgId = 102 // 登录响应
	MsgId_MSG_CS_Register   MsgId = 103 // 注册请求
	MsgId_MSG_SC_Register   MsgId = 104 // 注册响应
	MsgId_MSG_CS_HallInfo   MsgId = 111 // 大厅数据请求
	MsgId_MSG_SC_HallInfo   MsgId = 112 // 大厅数据响应
	MsgId_MSG_CS_CreateRoom MsgId = 121 // 创建房间请求
	MsgId_MSG_SC_CreateRoom MsgId = 122 // 创建房间响应
	MsgId_MSG_CS_JoinRoom   MsgId = 123 // 加入房间请求
	MsgId_MSG_SC_JoinRoom   MsgId = 124 // 加入房间响应
	MsgId_MSG_CS_MatchRoom  MsgId = 125 // 匹配房间请求
	MsgId_MSG_SC_MatchRoom  MsgId = 126 // 匹配房间响应
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0:   "MSG_Default",
		1:   "MSG_CS_Ping",
		2:   "MSG_SC_Pong",
		101: "MSG_CS_Login",
		102: "MSG_SC_Login",
		103: "MSG_CS_Register",
		104: "MSG_SC_Register",
		111: "MSG_CS_HallInfo",
		112: "MSG_SC_HallInfo",
		121: "MSG_CS_CreateRoom",
		122: "MSG_SC_CreateRoom",
		123: "MSG_CS_JoinRoom",
		124: "MSG_SC_JoinRoom",
		125: "MSG_CS_MatchRoom",
		126: "MSG_SC_MatchRoom",
	}
	MsgId_value = map[string]int32{
		"MSG_Default":       0,
		"MSG_CS_Ping":       1,
		"MSG_SC_Pong":       2,
		"MSG_CS_Login":      101,
		"MSG_SC_Login":      102,
		"MSG_CS_Register":   103,
		"MSG_SC_Register":   104,
		"MSG_CS_HallInfo":   111,
		"MSG_SC_HallInfo":   112,
		"MSG_CS_CreateRoom": 121,
		"MSG_SC_CreateRoom": 122,
		"MSG_CS_JoinRoom":   123,
		"MSG_SC_JoinRoom":   124,
		"MSG_CS_MatchRoom":  125,
		"MSG_SC_MatchRoom":  126,
	}
)

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}

func (x MsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type RoomState int32

const (
	RoomState_None   RoomState = 0 // 未开始
	RoomState_Ready  RoomState = 1 // 准备
	RoomState_Draw   RoomState = 2 // 画画
	RoomState_Result RoomState = 3 // 结果
	RoomState_Over   RoomState = 4 // 结束
)

// Enum value maps for RoomState.
var (
	RoomState_name = map[int32]string{
		0: "None",
		1: "Ready",
		2: "Draw",
		3: "Result",
		4: "Over",
	}
	RoomState_value = map[string]int32{
		"None":   0,
		"Ready":  1,
		"Draw":   2,
		"Result": 3,
		"Over":   4,
	}
)

func (x RoomState) Enum() *RoomState {
	p := new(RoomState)
	*p = x
	return p
}

func (x RoomState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomState) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[1].Descriptor()
}

func (RoomState) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[1]
}

func (x RoomState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomState.Descriptor instead.
func (RoomState) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

type CS_Ping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_Ping) Reset() {
	*x = CS_Ping{}
	mi := &file_msg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Ping) ProtoMessage() {}

func (x *CS_Ping) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Ping.ProtoReflect.Descriptor instead.
func (*CS_Ping) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *CS_Ping) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SC_Pong struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     int64                  `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_Pong) Reset() {
	*x = SC_Pong{}
	mi := &file_msg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Pong) ProtoMessage() {}

func (x *SC_Pong) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Pong.ProtoReflect.Descriptor instead.
func (*SC_Pong) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SC_Pong) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PlayerInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HeadId        int32                  `protobuf:"varint,1,opt,name=headId,proto3" json:"headId,omitempty"`
	Account       string                 `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	NickName      string                 `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	mi := &file_msg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerInfo) GetHeadId() int32 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

func (x *PlayerInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *PlayerInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type RoomInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 房间id
	Rid int32 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	// 房间人数
	GameNum int32 `protobuf:"varint,2,opt,name=gameNum,proto3" json:"gameNum,omitempty"`
	// 房间最大座位数
	Max int32 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	// 房间状态
	State int32 `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	// 房间创建时间
	CreateTime int32 `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 开始时间
	StartTime int32 `protobuf:"varint,6,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// 结果时间
	ResultTime int32 `protobuf:"varint,7,opt,name=resultTime,proto3" json:"resultTime,omitempty"`
	// 提示
	Hint string `protobuf:"bytes,8,opt,name=hint,proto3" json:"hint,omitempty"`
	// 单词
	Word string `protobuf:"bytes,9,opt,name=word,proto3" json:"word,omitempty"`
	// 单词索引
	WordIndex int32 `protobuf:"varint,10,opt,name=wordIndex,proto3" json:"wordIndex,omitempty"`
	// 画师
	Painter int32 `protobuf:"varint,11,opt,name=painter,proto3" json:"painter,omitempty"`
	// 房间所有玩家
	MapPlayerInfo map[int32]*PlayerInfo `protobuf:"bytes,12,rep,name=mapPlayerInfo,proto3" json:"mapPlayerInfo,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	mi := &file_msg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *RoomInfo) GetRid() int32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *RoomInfo) GetGameNum() int32 {
	if x != nil {
		return x.GameNum
	}
	return 0
}

func (x *RoomInfo) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *RoomInfo) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *RoomInfo) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *RoomInfo) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *RoomInfo) GetResultTime() int32 {
	if x != nil {
		return x.ResultTime
	}
	return 0
}

func (x *RoomInfo) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *RoomInfo) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *RoomInfo) GetWordIndex() int32 {
	if x != nil {
		return x.WordIndex
	}
	return 0
}

func (x *RoomInfo) GetPainter() int32 {
	if x != nil {
		return x.Painter
	}
	return 0
}

func (x *RoomInfo) GetMapPlayerInfo() map[int32]*PlayerInfo {
	if x != nil {
		return x.MapPlayerInfo
	}
	return nil
}

// *
// 登录请求
type CS_Login struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       string                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ServerId      int32                  `protobuf:"varint,3,opt,name=serverId,proto3" json:"serverId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_Login) Reset() {
	*x = CS_Login{}
	mi := &file_msg_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Login) ProtoMessage() {}

func (x *CS_Login) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Login.ProtoReflect.Descriptor instead.
func (*CS_Login) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *CS_Login) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CS_Login) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CS_Login) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

// *
// 登录响应
type SC_Login struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	PlayerInfo    *PlayerInfo            `protobuf:"bytes,3,opt,name=PlayerInfo,proto3" json:"PlayerInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_Login) Reset() {
	*x = SC_Login{}
	mi := &file_msg_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Login) ProtoMessage() {}

func (x *SC_Login) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Login.ProtoReflect.Descriptor instead.
func (*SC_Login) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *SC_Login) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SC_Login) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SC_Login) GetPlayerInfo() *PlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

// *
// 注册请求
type CS_Register struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       string                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	HeadId        int32                  `protobuf:"varint,3,opt,name=headId,proto3" json:"headId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_Register) Reset() {
	*x = CS_Register{}
	mi := &file_msg_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Register) ProtoMessage() {}

func (x *CS_Register) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Register.ProtoReflect.Descriptor instead.
func (*CS_Register) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{6}
}

func (x *CS_Register) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CS_Register) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CS_Register) GetHeadId() int32 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

// *
// 注册响应
type SC_Register struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_Register) Reset() {
	*x = SC_Register{}
	mi := &file_msg_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Register) ProtoMessage() {}

func (x *SC_Register) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Register.ProtoReflect.Descriptor instead.
func (*SC_Register) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{7}
}

func (x *SC_Register) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// *
// 大厅数据请求
type CS_HallInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_HallInfo) Reset() {
	*x = CS_HallInfo{}
	mi := &file_msg_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_HallInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_HallInfo) ProtoMessage() {}

func (x *CS_HallInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_HallInfo.ProtoReflect.Descriptor instead.
func (*CS_HallInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{8}
}

// *
// 大厅数据响应 repeated 数据写法
type SC_HallInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomArr       []*RoomInfo            `protobuf:"bytes,1,rep,name=roomArr,proto3" json:"roomArr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_HallInfo) Reset() {
	*x = SC_HallInfo{}
	mi := &file_msg_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_HallInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_HallInfo) ProtoMessage() {}

func (x *SC_HallInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_HallInfo.ProtoReflect.Descriptor instead.
func (*SC_HallInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{9}
}

func (x *SC_HallInfo) GetRoomArr() []*RoomInfo {
	if x != nil {
		return x.RoomArr
	}
	return nil
}

// *
// 创建房间请求
type CS_CreateRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_CreateRoom) Reset() {
	*x = CS_CreateRoom{}
	mi := &file_msg_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_CreateRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_CreateRoom) ProtoMessage() {}

func (x *CS_CreateRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_CreateRoom.ProtoReflect.Descriptor instead.
func (*CS_CreateRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{10}
}

func (x *CS_CreateRoom) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// 创建房间响应
type SC_CreateRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_CreateRoom) Reset() {
	*x = SC_CreateRoom{}
	mi := &file_msg_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_CreateRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_CreateRoom) ProtoMessage() {}

func (x *SC_CreateRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_CreateRoom.ProtoReflect.Descriptor instead.
func (*SC_CreateRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{11}
}

func (x *SC_CreateRoom) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// *
// 加入房间请求
type CS_JoinRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_JoinRoom) Reset() {
	*x = CS_JoinRoom{}
	mi := &file_msg_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_JoinRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_JoinRoom) ProtoMessage() {}

func (x *CS_JoinRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_JoinRoom.ProtoReflect.Descriptor instead.
func (*CS_JoinRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{12}
}

func (x *CS_JoinRoom) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// 加入房间响应
type SC_JoinRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_JoinRoom) Reset() {
	*x = SC_JoinRoom{}
	mi := &file_msg_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_JoinRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_JoinRoom) ProtoMessage() {}

func (x *SC_JoinRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_JoinRoom.ProtoReflect.Descriptor instead.
func (*SC_JoinRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{13}
}

func (x *SC_JoinRoom) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

// *
// 匹配房间请求
type CS_MatchRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CS_MatchRoom) Reset() {
	*x = CS_MatchRoom{}
	mi := &file_msg_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CS_MatchRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_MatchRoom) ProtoMessage() {}

func (x *CS_MatchRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_MatchRoom.ProtoReflect.Descriptor instead.
func (*CS_MatchRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{14}
}

func (x *CS_MatchRoom) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// 匹配房间响应
type SC_MatchRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SC_MatchRoom) Reset() {
	*x = SC_MatchRoom{}
	mi := &file_msg_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SC_MatchRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_MatchRoom) ProtoMessage() {}

func (x *SC_MatchRoom) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_MatchRoom.ProtoReflect.Descriptor instead.
func (*SC_MatchRoom) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{15}
}

func (x *SC_MatchRoom) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = string([]byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x27, 0x0a, 0x07, 0x43, 0x53, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x27, 0x0a, 0x07, 0x53, 0x43, 0x5f,
	0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x5a, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb7,
	0x03, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x67, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a,
	0x0d, 0x6d, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x51, 0x0a, 0x12, 0x4d, 0x61, 0x70, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x08, 0x43, 0x53, 0x5f, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x0a, 0x0a,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a,
	0x0b, 0x43, 0x53, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x43,
	0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x43, 0x53, 0x5f, 0x48, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x36, 0x0a, 0x0b,
	0x53, 0x43, 0x5f, 0x48, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x41, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x41, 0x72, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x53, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x43, 0x5f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x53,
	0x5f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x43, 0x5f,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1e, 0x0a, 0x0c,
	0x43, 0x53, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0c,
	0x53, 0x43, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x2a, 0xb6, 0x02, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x53,
	0x47, 0x5f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d,
	0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43, 0x5f, 0x50, 0x6f, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x65, 0x12,
	0x10, 0x0a, 0x0c, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10,
	0x66, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x10, 0x67, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43,
	0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10, 0x68, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x48, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x6f,
	0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43, 0x5f, 0x48, 0x61, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x10, 0x70, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x79, 0x12, 0x15, 0x0a, 0x11,
	0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x10, 0x7a, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x7b, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53, 0x47, 0x5f,
	0x53, 0x43, 0x5f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x7c, 0x12, 0x14, 0x0a,
	0x10, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x53, 0x5f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f,
	0x6d, 0x10, 0x7d, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x43, 0x5f, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x7e, 0x2a, 0x40, 0x0a, 0x09, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x72, 0x61, 0x77, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x10,
	0x03, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x76, 0x65, 0x72, 0x10, 0x04, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData []byte
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_msg_proto_rawDesc), len(file_msg_proto_rawDesc)))
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_msg_proto_goTypes = []any{
	(MsgId)(0),            // 0: msg.MsgId
	(RoomState)(0),        // 1: msg.RoomState
	(*CS_Ping)(nil),       // 2: msg.CS_Ping
	(*SC_Pong)(nil),       // 3: msg.SC_Pong
	(*PlayerInfo)(nil),    // 4: msg.PlayerInfo
	(*RoomInfo)(nil),      // 5: msg.RoomInfo
	(*CS_Login)(nil),      // 6: msg.CS_Login
	(*SC_Login)(nil),      // 7: msg.SC_Login
	(*CS_Register)(nil),   // 8: msg.CS_Register
	(*SC_Register)(nil),   // 9: msg.SC_Register
	(*CS_HallInfo)(nil),   // 10: msg.CS_HallInfo
	(*SC_HallInfo)(nil),   // 11: msg.SC_HallInfo
	(*CS_CreateRoom)(nil), // 12: msg.CS_CreateRoom
	(*SC_CreateRoom)(nil), // 13: msg.SC_CreateRoom
	(*CS_JoinRoom)(nil),   // 14: msg.CS_JoinRoom
	(*SC_JoinRoom)(nil),   // 15: msg.SC_JoinRoom
	(*CS_MatchRoom)(nil),  // 16: msg.CS_MatchRoom
	(*SC_MatchRoom)(nil),  // 17: msg.SC_MatchRoom
	nil,                   // 18: msg.RoomInfo.MapPlayerInfoEntry
}
var file_msg_proto_depIdxs = []int32{
	18, // 0: msg.RoomInfo.mapPlayerInfo:type_name -> msg.RoomInfo.MapPlayerInfoEntry
	4,  // 1: msg.SC_Login.PlayerInfo:type_name -> msg.PlayerInfo
	5,  // 2: msg.SC_HallInfo.roomArr:type_name -> msg.RoomInfo
	4,  // 3: msg.RoomInfo.MapPlayerInfoEntry.value:type_name -> msg.PlayerInfo
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_msg_proto_rawDesc), len(file_msg_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
