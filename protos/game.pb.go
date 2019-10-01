// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package game_pb

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

type CopyResponse_Status int32

const (
	CopyResponse_OK           CopyResponse_Status = 0
	CopyResponse_UPDATE_ERROR CopyResponse_Status = 1
	CopyResponse_NULL_ERROR   CopyResponse_Status = 2
)

var CopyResponse_Status_name = map[int32]string{
	0: "OK",
	1: "UPDATE_ERROR",
	2: "NULL_ERROR",
}
var CopyResponse_Status_value = map[string]int32{
	"OK":           0,
	"UPDATE_ERROR": 1,
	"NULL_ERROR":   2,
}

func (x CopyResponse_Status) String() string {
	return proto.EnumName(CopyResponse_Status_name, int32(x))
}
func (CopyResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{5, 0}
}

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
	return fileDescriptor_game_bb62401282007a66, []int{8, 0}
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
	return fileDescriptor_game_bb62401282007a66, []int{0}
}
func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (dst *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(dst, src)
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

type Slot struct {
	Treasure             bool     `protobuf:"varint,1,opt,name=treasure,proto3" json:"treasure,omitempty"`
	PlayerId             string   `protobuf:"bytes,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slot) Reset()         { *m = Slot{} }
func (m *Slot) String() string { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()    {}
func (*Slot) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{1}
}
func (m *Slot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slot.Unmarshal(m, b)
}
func (m *Slot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slot.Marshal(b, m, deterministic)
}
func (dst *Slot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slot.Merge(dst, src)
}
func (m *Slot) XXX_Size() int {
	return xxx_messageInfo_Slot.Size(m)
}
func (m *Slot) XXX_DiscardUnknown() {
	xxx_messageInfo_Slot.DiscardUnknown(m)
}

var xxx_messageInfo_Slot proto.InternalMessageInfo

func (m *Slot) GetTreasure() bool {
	if m != nil {
		return m.Treasure
	}
	return false
}

func (m *Slot) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type SlotRow struct {
	Slot                 []*Slot  `protobuf:"bytes,1,rep,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlotRow) Reset()         { *m = SlotRow{} }
func (m *SlotRow) String() string { return proto.CompactTextString(m) }
func (*SlotRow) ProtoMessage()    {}
func (*SlotRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{2}
}
func (m *SlotRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotRow.Unmarshal(m, b)
}
func (m *SlotRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotRow.Marshal(b, m, deterministic)
}
func (dst *SlotRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotRow.Merge(dst, src)
}
func (m *SlotRow) XXX_Size() int {
	return xxx_messageInfo_SlotRow.Size(m)
}
func (m *SlotRow) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotRow.DiscardUnknown(m)
}

var xxx_messageInfo_SlotRow proto.InternalMessageInfo

func (m *SlotRow) GetSlot() []*Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

type Grid struct {
	SlotRows []*SlotRow `protobuf:"bytes,1,rep,name=slotRows,proto3" json:"slotRows,omitempty"`
	// indices for fast retrieval
	TreasureSlots        []int32          `protobuf:"varint,2,rep,packed,name=treasureSlots,proto3" json:"treasureSlots,omitempty"`
	PlayerSlots          map[string]int32 `protobuf:"bytes,3,rep,name=playerSlots,proto3" json:"playerSlots,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	EmptySlots           []int32          `protobuf:"varint,4,rep,packed,name=emptySlots,proto3" json:"emptySlots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Grid) Reset()         { *m = Grid{} }
func (m *Grid) String() string { return proto.CompactTextString(m) }
func (*Grid) ProtoMessage()    {}
func (*Grid) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{3}
}
func (m *Grid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grid.Unmarshal(m, b)
}
func (m *Grid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grid.Marshal(b, m, deterministic)
}
func (dst *Grid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grid.Merge(dst, src)
}
func (m *Grid) XXX_Size() int {
	return xxx_messageInfo_Grid.Size(m)
}
func (m *Grid) XXX_DiscardUnknown() {
	xxx_messageInfo_Grid.DiscardUnknown(m)
}

var xxx_messageInfo_Grid proto.InternalMessageInfo

func (m *Grid) GetSlotRows() []*SlotRow {
	if m != nil {
		return m.SlotRows
	}
	return nil
}

func (m *Grid) GetTreasureSlots() []int32 {
	if m != nil {
		return m.TreasureSlots
	}
	return nil
}

func (m *Grid) GetPlayerSlots() map[string]int32 {
	if m != nil {
		return m.PlayerSlots
	}
	return nil
}

func (m *Grid) GetEmptySlots() []int32 {
	if m != nil {
		return m.EmptySlots
	}
	return nil
}

type CopyRequest struct {
	Grid                 *Grid    `protobuf:"bytes,1,opt,name=grid,proto3" json:"grid,omitempty"`
	StateVersion         int32    `protobuf:"varint,2,opt,name=stateVersion,proto3" json:"stateVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CopyRequest) Reset()         { *m = CopyRequest{} }
func (m *CopyRequest) String() string { return proto.CompactTextString(m) }
func (*CopyRequest) ProtoMessage()    {}
func (*CopyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{4}
}
func (m *CopyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyRequest.Unmarshal(m, b)
}
func (m *CopyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyRequest.Marshal(b, m, deterministic)
}
func (dst *CopyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyRequest.Merge(dst, src)
}
func (m *CopyRequest) XXX_Size() int {
	return xxx_messageInfo_CopyRequest.Size(m)
}
func (m *CopyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CopyRequest proto.InternalMessageInfo

func (m *CopyRequest) GetGrid() *Grid {
	if m != nil {
		return m.Grid
	}
	return nil
}

func (m *CopyRequest) GetStateVersion() int32 {
	if m != nil {
		return m.StateVersion
	}
	return 0
}

type CopyResponse struct {
	Status               CopyResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=game_pb.CopyResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CopyResponse) Reset()         { *m = CopyResponse{} }
func (m *CopyResponse) String() string { return proto.CompactTextString(m) }
func (*CopyResponse) ProtoMessage()    {}
func (*CopyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{5}
}
func (m *CopyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyResponse.Unmarshal(m, b)
}
func (m *CopyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyResponse.Marshal(b, m, deterministic)
}
func (dst *CopyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyResponse.Merge(dst, src)
}
func (m *CopyResponse) XXX_Size() int {
	return xxx_messageInfo_CopyResponse.Size(m)
}
func (m *CopyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CopyResponse proto.InternalMessageInfo

func (m *CopyResponse) GetStatus() CopyResponse_Status {
	if m != nil {
		return m.Status
	}
	return CopyResponse_OK
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
	return fileDescriptor_game_bb62401282007a66, []int{6}
}
func (m *TakeSlotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSlotRequest.Unmarshal(m, b)
}
func (m *TakeSlotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSlotRequest.Marshal(b, m, deterministic)
}
func (dst *TakeSlotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSlotRequest.Merge(dst, src)
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
	return fileDescriptor_game_bb62401282007a66, []int{7}
}
func (m *PlayerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerState.Unmarshal(m, b)
}
func (m *PlayerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerState.Marshal(b, m, deterministic)
}
func (dst *PlayerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerState.Merge(dst, src)
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
	return fileDescriptor_game_bb62401282007a66, []int{8}
}
func (m *TakeSlotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeSlotResponse.Unmarshal(m, b)
}
func (m *TakeSlotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeSlotResponse.Marshal(b, m, deterministic)
}
func (dst *TakeSlotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeSlotResponse.Merge(dst, src)
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

type Registry struct {
	Players              []*Registry_Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Version              int32              `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{9}
}
func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (dst *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(dst, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetPlayers() []*Registry_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *Registry) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Registry_Player struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JoinOrder            int32    `protobuf:"varint,2,opt,name=join_order,json=joinOrder,proto3" json:"join_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Registry_Player) Reset()         { *m = Registry_Player{} }
func (m *Registry_Player) String() string { return proto.CompactTextString(m) }
func (*Registry_Player) ProtoMessage()    {}
func (*Registry_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_bb62401282007a66, []int{9, 0}
}
func (m *Registry_Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry_Player.Unmarshal(m, b)
}
func (m *Registry_Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry_Player.Marshal(b, m, deterministic)
}
func (dst *Registry_Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry_Player.Merge(dst, src)
}
func (m *Registry_Player) XXX_Size() int {
	return xxx_messageInfo_Registry_Player.Size(m)
}
func (m *Registry_Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Registry_Player proto.InternalMessageInfo

func (m *Registry_Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Registry_Player) GetJoinOrder() int32 {
	if m != nil {
		return m.JoinOrder
	}
	return 0
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
	return fileDescriptor_game_bb62401282007a66, []int{10}
}
func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (dst *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(dst, src)
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
	return fileDescriptor_game_bb62401282007a66, []int{11}
}
func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (dst *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(dst, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Coordinate)(nil), "game_pb.Coordinate")
	proto.RegisterType((*Slot)(nil), "game_pb.Slot")
	proto.RegisterType((*SlotRow)(nil), "game_pb.Slot_row")
	proto.RegisterType((*Grid)(nil), "game_pb.Grid")
	proto.RegisterMapType((map[string]int32)(nil), "game_pb.Grid.PlayerSlotsEntry")
	proto.RegisterType((*CopyRequest)(nil), "game_pb.CopyRequest")
	proto.RegisterType((*CopyResponse)(nil), "game_pb.CopyResponse")
	proto.RegisterType((*TakeSlotRequest)(nil), "game_pb.TakeSlotRequest")
	proto.RegisterType((*PlayerState)(nil), "game_pb.PlayerState")
	proto.RegisterType((*TakeSlotResponse)(nil), "game_pb.TakeSlotResponse")
	proto.RegisterType((*Registry)(nil), "game_pb.Registry")
	proto.RegisterType((*Registry_Player)(nil), "game_pb.Registry.Player")
	proto.RegisterType((*HeartbeatRequest)(nil), "game_pb.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "game_pb.HeartbeatResponse")
	proto.RegisterEnum("game_pb.CopyResponse_Status", CopyResponse_Status_name, CopyResponse_Status_value)
	proto.RegisterEnum("game_pb.TakeSlotResponse_Status", TakeSlotResponse_Status_name, TakeSlotResponse_Status_value)
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
	StatusCopy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyResponse, error)
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

func (c *gameServiceClient) StatusCopy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyResponse, error) {
	out := new(CopyResponse)
	err := c.cc.Invoke(ctx, "/game_pb.GameService/StatusCopy", in, out, opts...)
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
	StatusCopy(context.Context, *CopyRequest) (*CopyResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
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

func _GameService_StatusCopy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).StatusCopy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_pb.GameService/StatusCopy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).StatusCopy(ctx, req.(*CopyRequest))
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
			MethodName: "StatusCopy",
			Handler:    _GameService_StatusCopy_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _GameService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_game_bb62401282007a66) }

var fileDescriptor_game_bb62401282007a66 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0xaf, 0x9d, 0x34, 0x4d, 0x26, 0x6d, 0xcf, 0xd9, 0xeb, 0x21, 0x13, 0xe0, 0x54, 0x56, 0x3c,
	0xf4, 0xa5, 0x11, 0x0a, 0x48, 0x1c, 0x20, 0x55, 0x98, 0xd6, 0x2a, 0x56, 0x83, 0x13, 0xad, 0xdd,
	0x48, 0xbc, 0xb0, 0x72, 0x93, 0x55, 0x65, 0x2e, 0xc9, 0x9a, 0xb5, 0x93, 0x23, 0x12, 0x0f, 0xbc,
	0xf0, 0x01, 0xf8, 0x6c, 0x7c, 0x1e, 0x24, 0xb4, 0x7f, 0xec, 0xfc, 0x69, 0xef, 0xde, 0x3c, 0x33,
	0xbf, 0xf9, 0xcd, 0xcc, 0x6f, 0x66, 0x65, 0x80, 0xc7, 0x64, 0xce, 0x7a, 0x99, 0xe0, 0x05, 0x47,
	0x47, 0xf2, 0x9b, 0x66, 0x0f, 0xf8, 0x4b, 0x80, 0x6b, 0xce, 0xc5, 0x34, 0x5d, 0x24, 0x05, 0x43,
	0x0e, 0xd4, 0x04, 0x7f, 0xe7, 0x5a, 0xe7, 0xd6, 0xc5, 0x21, 0x91, 0x9f, 0xd2, 0x33, 0xe1, 0x33,
	0xd7, 0xd6, 0x9e, 0x09, 0x9f, 0xe1, 0x2b, 0xa8, 0x47, 0x33, 0x5e, 0xa0, 0x2e, 0x34, 0x0b, 0xc1,
	0x92, 0x7c, 0x29, 0x98, 0x4a, 0x68, 0x92, 0xca, 0x96, 0xb1, 0x6c, 0x96, 0xac, 0x99, 0x08, 0xa6,
	0x2a, 0xb5, 0x45, 0x2a, 0x1b, 0x5f, 0x42, 0x53, 0xe6, 0x53, 0xc9, 0xfe, 0x39, 0xd4, 0xf3, 0x19,
	0x2f, 0x5c, 0xeb, 0xbc, 0x76, 0xd1, 0xee, 0x9f, 0xf4, 0x4c, 0x57, 0x3d, 0x09, 0x20, 0x2a, 0x84,
	0xff, 0xb3, 0xa0, 0x7e, 0x2b, 0xd2, 0x29, 0xba, 0x84, 0xa6, 0x74, 0x10, 0xfe, 0x2e, 0x37, 0xf8,
	0xce, 0x0e, 0x5e, 0x12, 0x92, 0x0a, 0x82, 0xbe, 0x80, 0x93, 0xb2, 0x1d, 0x19, 0xcd, 0x5d, 0xfb,
	0xbc, 0x76, 0x71, 0x48, 0x76, 0x9d, 0xe8, 0x07, 0x68, 0xeb, 0xc6, 0x34, 0xa6, 0xa6, 0x78, 0x5f,
	0x57, 0xbc, 0xb2, 0x70, 0x6f, 0xb4, 0x01, 0xf8, 0x8b, 0x42, 0xac, 0xc9, 0x76, 0x0a, 0x7a, 0x0d,
	0xc0, 0xe6, 0x59, 0xb1, 0xd6, 0x04, 0x75, 0x55, 0x64, 0xcb, 0xd3, 0xbd, 0x02, 0x67, 0x9f, 0x40,
	0x8a, 0xfa, 0x96, 0xad, 0x95, 0x6a, 0x2d, 0x22, 0x3f, 0xd1, 0x19, 0x1c, 0xae, 0x92, 0xd9, 0x92,
	0x19, 0xa1, 0xb5, 0xf1, 0x9d, 0xfd, 0xc6, 0xc2, 0x31, 0xb4, 0xaf, 0x79, 0xb6, 0x26, 0xec, 0xf7,
	0x25, 0xcb, 0x0b, 0xa9, 0xd8, 0xa3, 0x48, 0xa7, 0x2a, 0x77, 0x5b, 0x31, 0xd9, 0x29, 0x51, 0x21,
	0x84, 0xe1, 0x38, 0x2f, 0x92, 0x82, 0x8d, 0x99, 0xc8, 0x53, 0xbe, 0x30, 0x94, 0x3b, 0x3e, 0xfc,
	0x07, 0x1c, 0x6b, 0xd6, 0x3c, 0xe3, 0x8b, 0x9c, 0xa1, 0xaf, 0xa1, 0x21, 0xe3, 0xcb, 0x5c, 0x11,
	0x9f, 0xf6, 0x3f, 0xad, 0x88, 0xb7, 0x61, 0xbd, 0x48, 0x61, 0x88, 0xc1, 0xe2, 0x3e, 0x34, 0xb4,
	0x07, 0x35, 0xc0, 0x1e, 0xde, 0x39, 0x07, 0xc8, 0x81, 0xe3, 0xfb, 0xd1, 0x8d, 0x17, 0xfb, 0xd4,
	0x27, 0x64, 0x48, 0x1c, 0x0b, 0x9d, 0x02, 0x84, 0xf7, 0x83, 0x81, 0xb1, 0x6d, 0x3c, 0x85, 0x17,
	0x71, 0xf2, 0x56, 0xc9, 0x5f, 0xce, 0x74, 0x0a, 0xb6, 0x99, 0xa8, 0x45, 0xec, 0x74, 0x8a, 0x3c,
	0x40, 0x73, 0xbe, 0x62, 0xb4, 0xe0, 0x74, 0x52, 0xdd, 0xa6, 0x1a, 0xa3, 0xdd, 0x7f, 0xb9, 0xd5,
	0x58, 0x19, 0x22, 0x8e, 0x84, 0xc7, 0x7c, 0xe3, 0xc1, 0x7f, 0x59, 0xd0, 0x36, 0xb2, 0xcb, 0xb1,
	0xd1, 0x27, 0xd0, 0xd2, 0x4b, 0xa3, 0x55, 0xa5, 0xea, 0x22, 0xd1, 0x15, 0x38, 0x93, 0xa5, 0x10,
	0x6c, 0x51, 0xd0, 0x8c, 0xe7, 0x69, 0x51, 0x8a, 0xf6, 0x9e, 0x6a, 0x2f, 0x0c, 0x78, 0x64, 0xb0,
	0x72, 0x79, 0xf9, 0x84, 0x0b, 0xe6, 0xd6, 0xf4, 0xf2, 0x94, 0x81, 0xff, 0xb6, 0xc1, 0xd9, 0x4c,
	0x6a, 0x74, 0x7e, 0xb3, 0xa7, 0xf3, 0x79, 0x55, 0x60, 0x1f, 0xba, 0xa7, 0x35, 0xfa, 0x16, 0x4e,
	0xcc, 0x04, 0x6a, 0x91, 0xfa, 0x9e, 0xdb, 0xfd, 0xb3, 0x8a, 0x60, 0x6b, 0x5c, 0x72, 0x9c, 0x6d,
	0x8c, 0x1c, 0xff, 0xf9, 0x64, 0x4d, 0x1d, 0x38, 0x09, 0xc2, 0xb1, 0x37, 0x08, 0x6e, 0x68, 0x10,
	0x8e, 0xee, 0x63, 0xbd, 0xa7, 0x68, 0x30, 0x8c, 0x69, 0xec, 0xdd, 0xf9, 0xa1, 0x63, 0xa3, 0x57,
	0xd0, 0x09, 0xa8, 0xf7, 0x33, 0x0d, 0x87, 0x31, 0xf5, 0x68, 0xe4, 0x93, 0xb1, 0x4f, 0x9c, 0x1a,
	0x3a, 0x03, 0x47, 0xb9, 0x87, 0xe1, 0xe0, 0x17, 0xfa, 0xa3, 0x77, 0x7d, 0x77, 0x3f, 0x72, 0xea,
	0xa8, 0x0b, 0x1f, 0x45, 0x03, 0x6f, 0xec, 0xd3, 0x20, 0x0c, 0x62, 0x1a, 0x84, 0x74, 0x44, 0x86,
	0xb7, 0xc4, 0x8f, 0x22, 0xe7, 0x10, 0xff, 0x63, 0x41, 0x93, 0xb0, 0xc7, 0x34, 0x97, 0x97, 0xdf,
	0x87, 0x23, 0xdd, 0x5a, 0xf9, 0x86, 0xdd, 0xaa, 0xff, 0x12, 0x63, 0x06, 0x21, 0x25, 0x10, 0xb9,
	0x70, 0xb4, 0xda, 0x39, 0xe5, 0xd2, 0xec, 0x7e, 0x03, 0x0d, 0x0d, 0x7e, 0x72, 0x42, 0x9f, 0x01,
	0xfc, 0xc6, 0xd3, 0x05, 0xe5, 0x62, 0xca, 0x84, 0x49, 0x6b, 0x49, 0xcf, 0x50, 0x3a, 0xf0, 0xaf,
	0xe0, 0xfc, 0xc4, 0x12, 0x51, 0x3c, 0xb0, 0xa4, 0xba, 0xc2, 0x0f, 0x9e, 0xc8, 0x25, 0x34, 0x85,
	0xe9, 0xcf, 0x9c, 0x46, 0xe7, 0x49, 0xe3, 0xa4, 0x82, 0xe0, 0x97, 0xd0, 0xd9, 0xe2, 0xd7, 0x0b,
	0xed, 0xff, 0x6b, 0x41, 0xfb, 0x36, 0x99, 0xb3, 0x88, 0x89, 0x55, 0x3a, 0x61, 0xc8, 0x83, 0x66,
	0xb9, 0x74, 0xe4, 0x3e, 0x73, 0x07, 0xaa, 0xad, 0xee, 0xc7, 0xef, 0xbd, 0x10, 0x7c, 0x80, 0xbe,
	0x07, 0xd0, 0x9b, 0x95, 0xaf, 0x14, 0x9d, 0xed, 0x3d, 0x5a, 0x4d, 0xf0, 0xea, 0xd9, 0xa7, 0x8c,
	0x0f, 0xd0, 0x0d, 0xb4, 0xaa, 0x26, 0xd1, 0xa6, 0xcc, 0xbe, 0x30, 0xdd, 0xee, 0x73, 0xa1, 0x92,
	0xe5, 0xa1, 0xa1, 0x7e, 0x28, 0x5f, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xa5, 0xbe, 0x20,
	0x5e, 0x06, 0x00, 0x00,
}
