// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package game_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TakeSlotResponse_Status int32

const (
	TakeSlotResponse_OK                     TakeSlotResponse_Status = 0
	TakeSlotResponse_INVALID_INPUT          TakeSlotResponse_Status = 1
	TakeSlotResponse_SLOT_TAKEN             TakeSlotResponse_Status = 2
	TakeSlotResponse_I_AM_NOT_A_SERVER      TakeSlotResponse_Status = 3
	TakeSlotResponse_I_AM_ONLY_BACKUP       TakeSlotResponse_Status = 4
	TakeSlotResponse_SLAVE_INIT_IN_PROGRESS TakeSlotResponse_Status = 5
)

var TakeSlotResponse_Status_name = map[int32]string{
	0: "OK",
	1: "INVALID_INPUT",
	2: "SLOT_TAKEN",
	3: "I_AM_NOT_A_SERVER",
	4: "I_AM_ONLY_BACKUP",
	5: "SLAVE_INIT_IN_PROGRESS",
}

var TakeSlotResponse_Status_value = map[string]int32{
	"OK":                     0,
	"INVALID_INPUT":          1,
	"SLOT_TAKEN":             2,
	"I_AM_NOT_A_SERVER":      3,
	"I_AM_ONLY_BACKUP":       4,
	"SLAVE_INIT_IN_PROGRESS": 5,
}

func (x TakeSlotResponse_Status) String() string {
	return proto.EnumName(TakeSlotResponse_Status_name, int32(x))
}

func (TakeSlotResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3, 0}
}

type Coordinate struct {
	Row                  int32    `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Col                  int32    `protobuf:"varint,2,opt,name=col,proto3" json:"col,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *Coordinate) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

type TakeSlotRequest struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MoveToCoordinate     *Coordinate `protobuf:"bytes,2,opt,name=move_to_coordinate,json=moveToCoordinate,proto3" json:"move_to_coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TakeSlotRequest) Reset()         { *m = TakeSlotRequest{} }
func (m *TakeSlotRequest) String() string { return proto.CompactTextString(m) }
func (*TakeSlotRequest) ProtoMessage()    {}
func (*TakeSlotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *TakeSlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSlotRequest.Unmarshal(m, b)
}
func (m *TakeSlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSlotRequest.Marshal(b, m, deterministic)
}
func (m *TakeSlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSlotRequest.Merge(m, src)
}
func (m *TakeSlotRequest) XXX_Size() int {
	return xxx_messageInfo_TakeSlotRequest.Size(m)
}
func (m *TakeSlotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeSlotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeSlotRequest proto.InternalMessageInfo

func (m *TakeSlotRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TakeSlotRequest) GetMoveToCoordinate() *Coordinate {
	if m != nil {
		return m.MoveToCoordinate
	}
	return nil
}

type PlayerState struct {
	PlayerId             string      `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	CurrentPosition      *Coordinate `protobuf:"bytes,2,opt,name=current_position,json=currentPosition,proto3" json:"current_position,omitempty"`
	Score                int32       `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PlayerState) Reset()         { *m = PlayerState{} }
func (m *PlayerState) String() string { return proto.CompactTextString(m) }
func (*PlayerState) ProtoMessage()    {}
func (*PlayerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *PlayerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerState.Unmarshal(m, b)
}
func (m *PlayerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerState.Marshal(b, m, deterministic)
}
func (m *PlayerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerState.Merge(m, src)
}
func (m *PlayerState) XXX_Size() int {
	return xxx_messageInfo_PlayerState.Size(m)
}
func (m *PlayerState) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerState.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerState proto.InternalMessageInfo

func (m *PlayerState) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerState) GetCurrentPosition() *Coordinate {
	if m != nil {
		return m.CurrentPosition
	}
	return nil
}

func (m *PlayerState) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type TakeSlotResponse struct {
	Status               TakeSlotResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_pb.TakeSlotResponse_Status" json:"status,omitempty"`
	PlayerStates         []*PlayerState          `protobuf:"bytes,2,rep,name=player_states,json=playerStates,proto3" json:"player_states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TakeSlotResponse) Reset()         { *m = TakeSlotResponse{} }
func (m *TakeSlotResponse) String() string { return proto.CompactTextString(m) }
func (*TakeSlotResponse) ProtoMessage()    {}
func (*TakeSlotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *TakeSlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSlotResponse.Unmarshal(m, b)
}
func (m *TakeSlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSlotResponse.Marshal(b, m, deterministic)
}
func (m *TakeSlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSlotResponse.Merge(m, src)
}
func (m *TakeSlotResponse) XXX_Size() int {
	return xxx_messageInfo_TakeSlotResponse.Size(m)
}
func (m *TakeSlotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeSlotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TakeSlotResponse proto.InternalMessageInfo

func (m *TakeSlotResponse) GetStatus() TakeSlotResponse_Status {
	if m != nil {
		return m.Status
	}
	return TakeSlotResponse_OK
}

func (m *TakeSlotResponse) GetPlayerStates() []*PlayerState {
	if m != nil {
		return m.PlayerStates
	}
	return nil
}

type HeartbeatRequest struct {
	PlayerId             string    `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Registry             *Registry `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *HeartbeatRequest) GetRegistry() *Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type HeartbeatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("game_pb.TakeSlotResponse_Status", TakeSlotResponse_Status_name, TakeSlotResponse_Status_value)
	proto.RegisterType((*Coordinate)(nil), "game_pb.Coordinate")
	proto.RegisterType((*TakeSlotRequest)(nil), "game_pb.TakeSlotRequest")
	proto.RegisterType((*PlayerState)(nil), "game_pb.PlayerState")
	proto.RegisterType((*TakeSlotResponse)(nil), "game_pb.TakeSlotResponse")
	proto.RegisterType((*HeartbeatRequest)(nil), "game_pb.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "game_pb.HeartbeatResponse")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x1d, 0x12, 0x92, 0x09, 0x49, 0x37, 0xdb, 0x80, 0x8c, 0xb9, 0x44, 0x3e, 0xf5, 0x42,
	0x84, 0xc2, 0x05, 0x2e, 0x48, 0xa6, 0xb5, 0x8a, 0x95, 0x60, 0x5b, 0x6b, 0x37, 0x12, 0x17, 0x56,
	0x8e, 0xbd, 0xaa, 0xac, 0x26, 0x59, 0xb3, 0xbb, 0x29, 0xaa, 0xc4, 0x81, 0x0b, 0x7f, 0x81, 0xdf,
	0x8b, 0xfc, 0x51, 0xbb, 0xaa, 0x4a, 0x6f, 0x3b, 0x6f, 0x9e, 0xdf, 0x9b, 0x37, 0x23, 0x03, 0x5c,
	0xc5, 0x3b, 0x36, 0xcf, 0x05, 0x57, 0x1c, 0x3f, 0x2f, 0xde, 0x34, 0xdf, 0x98, 0x23, 0x25, 0xe2,
	0xe4, 0x9a, 0x89, 0x0a, 0xb7, 0xde, 0x01, 0x9c, 0x71, 0x2e, 0xd2, 0x6c, 0x1f, 0x2b, 0x86, 0x11,
	0x74, 0x04, 0xff, 0x69, 0x68, 0x33, 0xed, 0xb4, 0x4b, 0x8a, 0x67, 0x81, 0x24, 0x7c, 0x6b, 0xe8,
	0x15, 0x92, 0xf0, 0xad, 0x95, 0xc2, 0x71, 0x14, 0x5f, 0xb3, 0x70, 0xcb, 0x15, 0x61, 0x3f, 0x0e,
	0x4c, 0x2a, 0x3c, 0x06, 0x3d, 0x4b, 0xcb, 0xaf, 0x06, 0x44, 0xcf, 0x52, 0x6c, 0x03, 0xde, 0xf1,
	0x1b, 0x46, 0x15, 0xa7, 0x49, 0x23, 0x5e, 0x6a, 0x0c, 0x17, 0x27, 0xf3, 0x7a, 0x92, 0x79, 0xeb,
	0x4b, 0x50, 0x41, 0x8f, 0x78, 0x8b, 0x58, 0xbf, 0x35, 0x18, 0x06, 0xdb, 0xf8, 0x96, 0x89, 0x50,
	0x15, 0x93, 0xbd, 0x81, 0x41, 0x5e, 0x96, 0xb4, 0x71, 0xea, 0x57, 0x80, 0x9b, 0xe2, 0x4f, 0x80,
	0x92, 0x83, 0x10, 0x6c, 0xaf, 0x68, 0xce, 0x65, 0xa6, 0x32, 0xbe, 0x7f, 0xca, 0xed, 0xb8, 0x26,
	0x07, 0x35, 0x17, 0x4f, 0xa1, 0x2b, 0x13, 0x2e, 0x98, 0xd1, 0x29, 0x63, 0x56, 0x85, 0xf5, 0x47,
	0x07, 0xd4, 0x26, 0x95, 0x39, 0xdf, 0x4b, 0x86, 0x3f, 0x40, 0x4f, 0xaa, 0x58, 0x1d, 0x64, 0x39,
	0xc4, 0x78, 0x31, 0x6b, 0x0c, 0x1e, 0x52, 0xe7, 0x61, 0xc9, 0x23, 0x35, 0x1f, 0x7f, 0x84, 0x51,
	0x9d, 0xa0, 0x00, 0x98, 0x34, 0xf4, 0x59, 0xe7, 0x74, 0xb8, 0x98, 0x36, 0x02, 0xf7, 0xe2, 0x92,
	0x17, 0x79, 0x5b, 0x48, 0xeb, 0x17, 0xf4, 0x2a, 0x31, 0xdc, 0x03, 0xdd, 0x5f, 0xa2, 0x23, 0x3c,
	0x81, 0x91, 0xeb, 0xad, 0xed, 0x95, 0x7b, 0x4e, 0x5d, 0x2f, 0xb8, 0x8c, 0x90, 0x86, 0xc7, 0x00,
	0xe1, 0xca, 0x8f, 0x68, 0x64, 0x2f, 0x1d, 0x0f, 0xe9, 0xf8, 0x25, 0x4c, 0x5c, 0x6a, 0x7f, 0xa5,
	0x9e, 0x1f, 0x51, 0x9b, 0x86, 0x0e, 0x59, 0x3b, 0x04, 0x75, 0xf0, 0x14, 0x50, 0x09, 0xfb, 0xde,
	0xea, 0x1b, 0xfd, 0x6c, 0x9f, 0x2d, 0x2f, 0x03, 0xf4, 0x0c, 0x9b, 0xf0, 0x2a, 0x5c, 0xd9, 0x6b,
	0x87, 0xba, 0x9e, 0x1b, 0x51, 0xd7, 0xa3, 0x01, 0xf1, 0x2f, 0x88, 0x13, 0x86, 0xa8, 0x6b, 0x7d,
	0x07, 0xf4, 0x85, 0xc5, 0x42, 0x6d, 0x58, 0xdc, 0x5c, 0xfc, 0xc9, 0x73, 0xbc, 0x85, 0xbe, 0x60,
	0x57, 0x99, 0x54, 0xe2, 0xb6, 0x3e, 0xc3, 0xa4, 0x09, 0x49, 0xea, 0x06, 0x69, 0x28, 0xd6, 0x09,
	0x4c, 0xee, 0xe9, 0x57, 0xcb, 0x5b, 0xfc, 0xd5, 0x60, 0x78, 0x11, 0xef, 0x58, 0xc8, 0xc4, 0x4d,
	0x96, 0x30, 0x6c, 0x43, 0xff, 0x6e, 0xc1, 0xd8, 0x78, 0x64, 0xe7, 0xe5, 0x58, 0xe6, 0xeb, 0xff,
	0x5e, 0xc3, 0x3a, 0xc2, 0xe7, 0x30, 0x68, 0x7c, 0x70, 0xcb, 0x7c, 0x98, 0xcd, 0x34, 0x1f, 0x6b,
	0xdd, 0xa9, 0x6c, 0x7a, 0xe5, 0x7f, 0xf3, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xf7,
	0x27, 0x96, 0x5d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceClient interface {
	TakeSlot(ctx context.Context, in *TakeSlotRequest, opts ...grpc.CallOption) (*TakeSlotResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) TakeSlot(ctx context.Context, in *TakeSlotRequest, opts ...grpc.CallOption) (*TakeSlotResponse, error) {
	out := new(TakeSlotResponse)
	err := c.cc.Invoke(ctx, "/game_pb.GameService/TakeSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/game_pb.GameService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
type GameServiceServer interface {
	TakeSlot(context.Context, *TakeSlotRequest) (*TakeSlotResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

// UnimplementedGameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (*UnimplementedGameServiceServer) TakeSlot(ctx context.Context, req *TakeSlotRequest) (*TakeSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeSlot not implemented")
}
func (*UnimplementedGameServiceServer) Heartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_TakeSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).TakeSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_pb.GameService/TakeSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).TakeSlot(ctx, req.(*TakeSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_pb.GameService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game_pb.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TakeSlot",
			Handler:    _GameService_TakeSlot_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _GameService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}
