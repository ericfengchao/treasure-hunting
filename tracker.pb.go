// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracker.proto

package treasure_hunting

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
	return fileDescriptor_a0ba8625d8751af3, []int{3, 0}
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
	return fileDescriptor_a0ba8625d8751af3, []int{5, 0}
}

type Player struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	PlayerId             string   `protobuf:"bytes,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
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

type Registry struct {
	PlayerList           []*Player `protobuf:"bytes,1,rep,name=player_list,json=playerList,proto3" json:"player_list,omitempty"`
	Version              int32     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{1}
}

func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetPlayerList() []*Player {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

func (m *Registry) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type RegisterRequest struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
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
	Status               RegisterResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_pb.RegisterResponse_Status" json:"status,omitempty"`
	Registry             *Registry               `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	N                    int32                   `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	K                    int32                   `protobuf:"varint,4,opt,name=K,proto3" json:"K,omitempty"`
	AssignedPort         int32                   `protobuf:"varint,5,opt,name=assigned_port,json=assignedPort,proto3" json:"assigned_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
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

func (m *RegisterResponse) GetRegistry() *Registry {
	if m != nil {
		return m.Registry
	}
	return nil
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

func (m *RegisterResponse) GetAssignedPort() int32 {
	if m != nil {
		return m.AssignedPort
	}
	return 0
}

type Missing struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Missing) Reset()         { *m = Missing{} }
func (m *Missing) String() string { return proto.CompactTextString(m) }
func (*Missing) ProtoMessage()    {}
func (*Missing) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{4}
}

func (m *Missing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Missing.Unmarshal(m, b)
}
func (m *Missing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Missing.Marshal(b, m, deterministic)
}
func (m *Missing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Missing.Merge(m, src)
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
	Status               MissingResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_pb.MissingResponse_Status" json:"status,omitempty"`
	Registry             *Registry              `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MissingResponse) Reset()         { *m = MissingResponse{} }
func (m *MissingResponse) String() string { return proto.CompactTextString(m) }
func (*MissingResponse) ProtoMessage()    {}
func (*MissingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0ba8625d8751af3, []int{5}
}

func (m *MissingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MissingResponse.Unmarshal(m, b)
}
func (m *MissingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MissingResponse.Marshal(b, m, deterministic)
}
func (m *MissingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissingResponse.Merge(m, src)
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

func (m *MissingResponse) GetRegistry() *Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

func init() {
	proto.RegisterEnum("game_pb.RegisterResponse_Status", RegisterResponse_Status_name, RegisterResponse_Status_value)
	proto.RegisterEnum("game_pb.MissingResponse_Status", MissingResponse_Status_name, MissingResponse_Status_value)
	proto.RegisterType((*Player)(nil), "game_pb.Player")
	proto.RegisterType((*Registry)(nil), "game_pb.Registry")
	proto.RegisterType((*RegisterRequest)(nil), "game_pb.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "game_pb.RegisterResponse")
	proto.RegisterType((*Missing)(nil), "game_pb.Missing")
	proto.RegisterType((*MissingResponse)(nil), "game_pb.MissingResponse")
}

func init() { proto.RegisterFile("tracker.proto", fileDescriptor_a0ba8625d8751af3) }

var fileDescriptor_a0ba8625d8751af3 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x69, 0xe3, 0x24, 0x93, 0xda, 0x31, 0x73, 0x5a, 0xca, 0xa1, 0xd6, 0x22, 0xa1,
	0x5c, 0xb0, 0x90, 0x39, 0xc0, 0x85, 0x03, 0x12, 0x16, 0xb2, 0x02, 0x69, 0xb5, 0xb6, 0x10, 0x37,
	0xcb, 0x6d, 0x56, 0xd6, 0x8a, 0x62, 0x9b, 0xdd, 0x6d, 0xa5, 0xbc, 0x06, 0x47, 0x1e, 0x8e, 0x67,
	0x41, 0x5d, 0xff, 0x41, 0x32, 0x84, 0x0b, 0x37, 0xcf, 0xfa, 0xb7, 0x9e, 0xef, 0xfb, 0x66, 0x0c,
	0xae, 0x51, 0xc5, 0xcd, 0x17, 0xa1, 0xc2, 0x46, 0xd5, 0xa6, 0xc6, 0x79, 0x59, 0x7c, 0x15, 0x79,
	0x73, 0xcd, 0x12, 0x70, 0xae, 0x6e, 0x8b, 0x83, 0x50, 0xe8, 0xc1, 0x54, 0x36, 0x94, 0x04, 0x64,
	0xb3, 0xe4, 0x53, 0xd9, 0x20, 0xc2, 0x69, 0x53, 0x2b, 0x43, 0xa7, 0x01, 0xd9, 0xcc, 0xb8, 0x7d,
	0xc6, 0x27, 0xb0, 0x6c, 0x2c, 0x9d, 0xcb, 0x3d, 0x3d, 0xb1, 0xe8, 0xa2, 0x3d, 0x48, 0xf6, 0xec,
	0x13, 0x2c, 0xb8, 0x28, 0xa5, 0x36, 0xea, 0x80, 0x2f, 0x60, 0xd5, 0x81, 0xb7, 0x52, 0x1b, 0x4a,
	0x82, 0x93, 0xcd, 0x2a, 0x5a, 0x87, 0x5d, 0xd7, 0xb0, 0x6d, 0xc9, 0xa1, 0x65, 0x3e, 0x48, 0x6d,
	0x90, 0xc2, 0xfc, 0x5e, 0x28, 0x2d, 0xeb, 0xaa, 0xeb, 0xd8, 0x97, 0x2c, 0x84, 0x75, 0xfb, 0x5d,
	0xa1, 0xb8, 0xf8, 0x76, 0x27, 0xf4, 0x48, 0x07, 0x19, 0xe9, 0xf8, 0x49, 0xc0, 0xff, 0x7d, 0x41,
	0x37, 0x75, 0xa5, 0x05, 0xbe, 0x06, 0x47, 0x9b, 0xc2, 0xdc, 0x69, 0x8b, 0x7b, 0x51, 0x30, 0x68,
	0x19, 0xa3, 0x61, 0x6a, 0x39, 0xde, 0xf1, 0xf8, 0x1c, 0x16, 0xaa, 0xb3, 0x65, 0x95, 0xad, 0xa2,
	0x47, 0xa3, 0xbb, 0xea, 0xc0, 0x07, 0x04, 0xcf, 0x80, 0xec, 0x6c, 0x34, 0x33, 0x4e, 0x76, 0x0f,
	0xd5, 0x96, 0x9e, 0xb6, 0xd5, 0x16, 0x9f, 0x82, 0x5b, 0x68, 0x2d, 0xcb, 0x4a, 0xec, 0x73, 0x9b,
	0xed, 0xcc, 0xbe, 0x39, 0xeb, 0x0f, 0xaf, 0x6a, 0x65, 0x58, 0x00, 0x4e, 0xab, 0x00, 0x1d, 0x98,
	0x5e, 0x6e, 0xfd, 0x09, 0x7a, 0x00, 0x3c, 0x7e, 0x9f, 0xa4, 0x59, 0xcc, 0xe3, 0x77, 0x3e, 0x61,
	0xcf, 0x60, 0xfe, 0x51, 0x6a, 0x2d, 0xab, 0xf2, 0xdf, 0x41, 0xfc, 0x20, 0xb0, 0xee, 0xc0, 0x21,
	0x87, 0x57, 0xa3, 0x1c, 0x2e, 0x06, 0x2f, 0x23, 0xf2, 0xff, 0x62, 0x60, 0x17, 0x7f, 0xb8, 0x70,
	0x61, 0xb9, 0xbb, 0xcc, 0xf2, 0xf8, 0x73, 0x92, 0x66, 0x3e, 0x89, 0xbe, 0x13, 0xf0, 0xb2, 0x76,
	0x27, 0x53, 0xa1, 0xee, 0xe5, 0x8d, 0xc0, 0xb7, 0xfd, 0x02, 0x09, 0x85, 0xf4, 0x2f, 0xf3, 0xb1,
	0xb3, 0x3f, 0x7f, 0x7c, 0x74, 0x72, 0x6c, 0x82, 0x6f, 0xc0, 0xe5, 0xe2, 0x21, 0xda, 0x3e, 0x20,
	0x7f, 0xec, 0xef, 0x9c, 0x1e, 0x73, 0xcc, 0x26, 0xd7, 0x8e, 0xfd, 0x3b, 0x5e, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xce, 0x31, 0x89, 0x18, 0x2e, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/game_pb.TrackerService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerServiceClient) ReportMissing(ctx context.Context, in *Missing, opts ...grpc.CallOption) (*MissingResponse, error) {
	out := new(MissingResponse)
	err := c.cc.Invoke(ctx, "/game_pb.TrackerService/ReportMissing", in, out, opts...)
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

// UnimplementedTrackerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTrackerServiceServer struct {
}

func (*UnimplementedTrackerServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedTrackerServiceServer) ReportMissing(ctx context.Context, req *Missing) (*MissingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMissing not implemented")
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
		FullMethod: "/game_pb.TrackerService/Register",
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
		FullMethod: "/game_pb.TrackerService/ReportMissing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServiceServer).ReportMissing(ctx, req.(*Missing))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrackerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game_pb.TrackerService",
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
