// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracker.proto

package game_tracker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RegisterResponse_Status int32

const (
	RegisterResponse_OK         RegisterResponse_Status = 0
	RegisterResponse_REGISTERED RegisterResponse_Status = 1
)

var RegisterResponse_Status_name = map[int32]string{
	0: "OK",
	1: "REGISTERED",
}
var RegisterResponse_Status_value = map[string]int32{
	"OK":         0,
	"REGISTERED": 1,
}

func (x RegisterResponse_Status) String() string {
	return proto.EnumName(RegisterResponse_Status_name, int32(x))
}
func (RegisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{2, 0}
}

type MissingResponse_Status int32

const (
	MissingResponse_OK        MissingResponse_Status = 0
	MissingResponse_NOT_EXIST MissingResponse_Status = 1
)

var MissingResponse_Status_name = map[int32]string{
	0: "OK",
	1: "NOT_EXIST",
}
var MissingResponse_Status_value = map[string]int32{
	"OK":        0,
	"NOT_EXIST": 1,
}

func (x MissingResponse_Status) String() string {
	return proto.EnumName(MissingResponse_Status_name, int32(x))
}
func (MissingResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{4, 0}
}

type Player struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	PlayerId             string   `protobuf:"bytes,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{0}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Player) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Player) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type RegisterRequest struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{1}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type RegisterResponse struct {
	Status               RegisterResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_tracker.RegisterResponse_Status" json:"status,omitempty"`
	PlayerList           []*Player               `protobuf:"bytes,2,rep,name=playerList,proto3" json:"playerList,omitempty"`
	Version              int32                   `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	N                    int32                   `protobuf:"varint,5,opt,name=N,proto3" json:"N,omitempty"`
	K                    int32                   `protobuf:"varint,6,opt,name=K,proto3" json:"K,omitempty"`
	StartPort            int32                   `protobuf:"varint,7,opt,name=startPort,proto3" json:"startPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{2}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetStatus() RegisterResponse_Status {
	if m != nil {
		return m.Status
	}
	return RegisterResponse_OK
}

func (m *RegisterResponse) GetPlayerList() []*Player {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

func (m *RegisterResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RegisterResponse) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *RegisterResponse) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *RegisterResponse) GetStartPort() int32 {
	if m != nil {
		return m.StartPort
	}
	return 0
}

type Missing struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Missing) Reset()         { *m = Missing{} }
func (m *Missing) String() string { return proto.CompactTextString(m) }
func (*Missing) ProtoMessage()    {}
func (*Missing) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{3}
}
func (m *Missing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Missing.Unmarshal(m, b)
}
func (m *Missing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Missing.Marshal(b, m, deterministic)
}
func (dst *Missing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Missing.Merge(dst, src)
}
func (m *Missing) XXX_Size() int {
	return xxx_messageInfo_Missing.Size(m)
}
func (m *Missing) XXX_DiscardUnknown() {
	xxx_messageInfo_Missing.DiscardUnknown(m)
}

var xxx_messageInfo_Missing proto.InternalMessageInfo

func (m *Missing) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type MissingResponse struct {
	Status               MissingResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_tracker.MissingResponse_Status" json:"status,omitempty"`
	PlayerList           []*Player              `protobuf:"bytes,2,rep,name=playerList,proto3" json:"playerList,omitempty"`
	Version              int32                  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MissingResponse) Reset()         { *m = MissingResponse{} }
func (m *MissingResponse) String() string { return proto.CompactTextString(m) }
func (*MissingResponse) ProtoMessage()    {}
func (*MissingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_dc7c183669eed543, []int{4}
}
func (m *MissingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissingResponse.Unmarshal(m, b)
}
func (m *MissingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissingResponse.Marshal(b, m, deterministic)
}
func (dst *MissingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissingResponse.Merge(dst, src)
}
func (m *MissingResponse) XXX_Size() int {
	return xxx_messageInfo_MissingResponse.Size(m)
}
func (m *MissingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MissingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MissingResponse proto.InternalMessageInfo

func (m *MissingResponse) GetStatus() MissingResponse_Status {
	if m != nil {
		return m.Status
	}
	return MissingResponse_OK
}

func (m *MissingResponse) GetPlayerList() []*Player {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

func (m *MissingResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "game_tracker.Player")
	proto.RegisterType((*RegisterRequest)(nil), "game_tracker.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "game_tracker.RegisterResponse")
	proto.RegisterType((*Missing)(nil), "game_tracker.Missing")
	proto.RegisterType((*MissingResponse)(nil), "game_tracker.MissingResponse")
	proto.RegisterEnum("game_tracker.RegisterResponse_Status", RegisterResponse_Status_name, RegisterResponse_Status_value)
	proto.RegisterEnum("game_tracker.MissingResponse_Status", MissingResponse_Status_name, MissingResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrackerServiceClient is the client API for TrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrackerServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	ReportMissing(ctx context.Context, in *Missing, opts ...grpc.CallOption) (*MissingResponse, error)
}

type trackerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTrackerServiceClient(cc *grpc.ClientConn) TrackerServiceClient {
	return &trackerServiceClient{cc}
}

func (c *trackerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/game_tracker.TrackerService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerServiceClient) ReportMissing(ctx context.Context, in *Missing, opts ...grpc.CallOption) (*MissingResponse, error) {
	out := new(MissingResponse)
	err := c.cc.Invoke(ctx, "/game_tracker.TrackerService/ReportMissing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServiceServer is the server API for TrackerService service.
type TrackerServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	ReportMissing(context.Context, *Missing) (*MissingResponse, error)
}

func RegisterTrackerServiceServer(s *grpc.Server, srv TrackerServiceServer) {
	s.RegisterService(&_TrackerService_serviceDesc, srv)
}

func _TrackerService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_tracker.TrackerService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackerService_ReportMissing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Missing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServiceServer).ReportMissing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_tracker.TrackerService/ReportMissing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServiceServer).ReportMissing(ctx, req.(*Missing))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrackerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game_tracker.TrackerService",
	HandlerType: (*TrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _TrackerService_Register_Handler,
		},
		{
			MethodName: "ReportMissing",
			Handler:    _TrackerService_ReportMissing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracker.proto",
}

func init() { proto.RegisterFile("tracker.proto", fileDescriptor_tracker_dc7c183669eed543) }

var fileDescriptor_tracker_dc7c183669eed543 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4f, 0xea, 0x40,
	0x10, 0xc0, 0xd9, 0x02, 0x05, 0xe6, 0x41, 0x69, 0x36, 0xef, 0x25, 0x0d, 0x79, 0x28, 0x69, 0x24,
	0xe1, 0x62, 0x0f, 0xe8, 0x51, 0x6f, 0x36, 0x48, 0xaa, 0x40, 0xb6, 0x3d, 0x78, 0x23, 0x15, 0x37,
	0x64, 0xa3, 0xd2, 0xba, 0xbb, 0x90, 0xf8, 0x85, 0xfc, 0x1c, 0x7e, 0x31, 0x13, 0xd3, 0x2d, 0xff,
	0x5a, 0x95, 0x93, 0xb7, 0xce, 0xe4, 0xd7, 0xc9, 0xcc, 0x6f, 0x66, 0xa1, 0x21, 0x79, 0x38, 0x7b,
	0xa4, 0xdc, 0x89, 0x79, 0x24, 0x23, 0x5c, 0x9f, 0x87, 0xcf, 0x74, 0xba, 0xce, 0xd9, 0xd7, 0xa0,
	0x4f, 0x9e, 0xc2, 0x57, 0xca, 0xb1, 0x01, 0x1a, 0x8b, 0x2d, 0xd4, 0x41, 0xbd, 0x1a, 0xd1, 0x58,
	0x8c, 0x31, 0x94, 0xe2, 0x88, 0x4b, 0x4b, 0xeb, 0xa0, 0x5e, 0x99, 0xa8, 0x6f, 0xdc, 0x82, 0x6a,
	0xac, 0xe8, 0xe1, 0x83, 0x55, 0x54, 0xe4, 0x36, 0xb6, 0x4f, 0xa1, 0x49, 0xe8, 0x9c, 0x09, 0x49,
	0x39, 0xa1, 0x2f, 0x4b, 0x2a, 0xb2, 0x38, 0xca, 0xe1, 0x1f, 0x08, 0xcc, 0x1d, 0x2f, 0xe2, 0x68,
	0x21, 0x28, 0xbe, 0x04, 0x5d, 0xc8, 0x50, 0x2e, 0x85, 0xc2, 0x8d, 0x7e, 0xd7, 0xd9, 0x6f, 0xd6,
	0xc9, 0xf3, 0x8e, 0xaf, 0x60, 0xb2, 0xfe, 0x09, 0x9f, 0x03, 0xa4, 0xf5, 0x6f, 0x98, 0x48, 0x1a,
	0x2f, 0xf6, 0xfe, 0xf4, 0xff, 0x66, 0x4b, 0xa4, 0xc3, 0x92, 0x3d, 0x0e, 0x5b, 0x50, 0x59, 0x51,
	0x2e, 0x58, 0xb4, 0xb0, 0x4a, 0x6a, 0xd6, 0x4d, 0x88, 0xeb, 0x80, 0x46, 0x56, 0x59, 0xe5, 0xd0,
	0x28, 0x89, 0x3c, 0x4b, 0x4f, 0x23, 0x0f, 0xff, 0x87, 0x9a, 0x90, 0x21, 0x97, 0x93, 0xc4, 0x51,
	0x45, 0x65, 0x77, 0x09, 0xbb, 0x03, 0x7a, 0xda, 0x1b, 0xd6, 0x41, 0x1b, 0x7b, 0x66, 0x01, 0x1b,
	0x00, 0xc4, 0x1d, 0x0c, 0xfd, 0xc0, 0x25, 0xee, 0x95, 0x89, 0xec, 0x2e, 0x54, 0x6e, 0x99, 0x10,
	0x6c, 0x31, 0x3f, 0xa8, 0xe9, 0x1d, 0x41, 0x73, 0xcd, 0x6d, 0x2d, 0x5d, 0xe4, 0x2c, 0x9d, 0x64,
	0x47, 0xcc, 0xe1, 0xbf, 0x2e, 0xa9, 0x98, 0x91, 0x64, 0x1f, 0x7f, 0x19, 0xb5, 0x01, 0xb5, 0xd1,
	0x38, 0x98, 0xba, 0x77, 0x43, 0x3f, 0x30, 0x51, 0xff, 0x0d, 0x81, 0x11, 0xa4, 0x95, 0x7d, 0xca,
	0x57, 0x6c, 0x46, 0xb1, 0x07, 0xd5, 0xcd, 0x2e, 0x71, 0xfb, 0xa7, 0x1d, 0xab, 0x1b, 0x6a, 0x1d,
	0x1d, 0x3e, 0x01, 0xbb, 0x80, 0x07, 0xd0, 0x20, 0x34, 0x39, 0xcf, 0x8d, 0xcf, 0x7f, 0xdf, 0xfa,
	0x68, 0xb5, 0x0f, 0x6a, 0xb2, 0x0b, 0xf7, 0xba, 0x7a, 0x20, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x62, 0xbb, 0xaa, 0x07, 0x31, 0x03, 0x00, 0x00,
}