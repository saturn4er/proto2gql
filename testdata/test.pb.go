// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package testdata is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	RootMessage
	Empty
	MessageWithEmpty
	RootMessage2
*/
package testdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/saturn4er/proto2gql/testdata/common"

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

type RootEnum int32

const (
	RootEnum_RootEnumVal0 RootEnum = 0
	RootEnum_RootEnumVal1 RootEnum = 1
	// It's a RootEnumVal2
	RootEnum_RootEnumVal2 RootEnum = 2
)

var RootEnum_name = map[int32]string{
	0: "RootEnumVal0",
	1: "RootEnumVal1",
	2: "RootEnumVal2",
}
var RootEnum_value = map[string]int32{
	"RootEnumVal0": 0,
	"RootEnumVal1": 1,
	"RootEnumVal2": 2,
}

func (x RootEnum) String() string {
	return proto.EnumName(RootEnum_name, int32(x))
}
func (RootEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RootMessage_NestedEnum int32

const (
	RootMessage_NestedEnumVal0 RootMessage_NestedEnum = 0
	RootMessage_NestedEnumVal1 RootMessage_NestedEnum = 1
)

var RootMessage_NestedEnum_name = map[int32]string{
	0: "NestedEnumVal0",
	1: "NestedEnumVal1",
}
var RootMessage_NestedEnum_value = map[string]int32{
	"NestedEnumVal0": 0,
	"NestedEnumVal1": 1,
}

func (x RootMessage_NestedEnum) String() string {
	return proto.EnumName(RootMessage_NestedEnum_name, int32(x))
}
func (RootMessage_NestedEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type RootMessage_NestedMessage_NestedNestedEnum int32

const (
	RootMessage_NestedMessage_NestedNestedEnumVal0 RootMessage_NestedMessage_NestedNestedEnum = 0
	RootMessage_NestedMessage_NestedNestedEnumVal1 RootMessage_NestedMessage_NestedNestedEnum = 1
	RootMessage_NestedMessage_NestedNestedEnumVal2 RootMessage_NestedMessage_NestedNestedEnum = 2
	RootMessage_NestedMessage_NestedNestedEnumVal3 RootMessage_NestedMessage_NestedNestedEnum = 3
)

var RootMessage_NestedMessage_NestedNestedEnum_name = map[int32]string{
	0: "NestedNestedEnumVal0",
	1: "NestedNestedEnumVal1",
	2: "NestedNestedEnumVal2",
	3: "NestedNestedEnumVal3",
}
var RootMessage_NestedMessage_NestedNestedEnum_value = map[string]int32{
	"NestedNestedEnumVal0": 0,
	"NestedNestedEnumVal1": 1,
	"NestedNestedEnumVal2": 2,
	"NestedNestedEnumVal3": 3,
}

func (x RootMessage_NestedMessage_NestedNestedEnum) String() string {
	return proto.EnumName(RootMessage_NestedMessage_NestedNestedEnum_name, int32(x))
}
func (RootMessage_NestedMessage_NestedNestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type RootMessage struct {
	// enum_map
	MapEnum map[int32]RootMessage_NestedEnum `protobuf:"bytes,1,rep,name=map_enum,json=mapEnum" json:"map_enum,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=example.RootMessage_NestedEnum"`
	// scalar map
	MapScalar map[int32]int32                      `protobuf:"bytes,28,rep,name=map_scalar,json=mapScalar" json:"map_scalar,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	MapMsg    map[int32]*RootMessage_NestedMessage `protobuf:"bytes,2,rep,name=map_msg,json=mapMsg" json:"map_msg,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// repeated Message
	RMsg []*RootMessage_NestedMessage `protobuf:"bytes,3,rep,name=r_msg,json=rMsg" json:"r_msg,omitempty"`
	// repeated Scalar
	RScalar []int32 `protobuf:"varint,4,rep,packed,name=r_scalar,json=rScalar" json:"r_scalar,omitempty"`
	// repeated Enum
	REnum []RootEnum `protobuf:"varint,5,rep,packed,name=r_enum,json=rEnum,enum=example.RootEnum" json:"r_enum,omitempty"`
	// repeated empty message
	REmptyMsg []*Empty `protobuf:"bytes,6,rep,name=r_empty_msg,json=rEmptyMsg" json:"r_empty_msg,omitempty"`
	// non-repeated Enum
	NREnum common.CommonEnum `protobuf:"varint,7,opt,name=n_r_enum,json=nREnum,enum=common.CommonEnum" json:"n_r_enum,omitempty"`
	// non-repeated Scalar
	NRScalar int32 `protobuf:"varint,8,opt,name=n_r_scalar,json=nRScalar" json:"n_r_scalar,omitempty"`
	// non-repeated Message
	NRMsg *common.CommonMessage `protobuf:"bytes,9,opt,name=n_r_msg,json=nRMsg" json:"n_r_msg,omitempty"`
	// field from context
	ScalarFromContext int32 `protobuf:"varint,10,opt,name=scalar_from_context,json=scalarFromContext" json:"scalar_from_context,omitempty"`
	// non-repeated empty message field
	NREmptyMsg *Empty `protobuf:"bytes,11,opt,name=n_r_empty_msg,json=nREmptyMsg" json:"n_r_empty_msg,omitempty"`
	// Types that are valid to be assigned to EnumFirstOneoff:
	//	*RootMessage_EFOE
	//	*RootMessage_EFOS
	//	*RootMessage_EFOM
	//	*RootMessage_EFOEm
	EnumFirstOneoff isRootMessage_EnumFirstOneoff `protobuf_oneof:"enum_first_oneoff"`
	// Types that are valid to be assigned to ScalarFirstOneoff:
	//	*RootMessage_SFOS
	//	*RootMessage_SFOE
	//	*RootMessage_SFOMes
	//	*RootMessage_SFOM
	ScalarFirstOneoff isRootMessage_ScalarFirstOneoff `protobuf_oneof:"scalar_first_oneoff"`
	// Types that are valid to be assigned to MessageFirstOneoff:
	//	*RootMessage_MFOM
	//	*RootMessage_MFOS
	//	*RootMessage_MFOE
	//	*RootMessage_MFOEm
	MessageFirstOneoff isRootMessage_MessageFirstOneoff `protobuf_oneof:"message_first_oneoff"`
	// Types that are valid to be assigned to EmptyFirstOneoff:
	//	*RootMessage_EmFOEm
	//	*RootMessage_EmFOS
	//	*RootMessage_EmFOEn
	//	*RootMessage_EmFOM
	EmptyFirstOneoff isRootMessage_EmptyFirstOneoff `protobuf_oneof:"empty_first_oneoff"`
}

func (m *RootMessage) Reset()                    { *m = RootMessage{} }
func (m *RootMessage) String() string            { return proto.CompactTextString(m) }
func (*RootMessage) ProtoMessage()               {}
func (*RootMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isRootMessage_EnumFirstOneoff interface {
	isRootMessage_EnumFirstOneoff()
}
type isRootMessage_ScalarFirstOneoff interface {
	isRootMessage_ScalarFirstOneoff()
}
type isRootMessage_MessageFirstOneoff interface {
	isRootMessage_MessageFirstOneoff()
}
type isRootMessage_EmptyFirstOneoff interface {
	isRootMessage_EmptyFirstOneoff()
}

type RootMessage_EFOE struct {
	EFOE common.CommonEnum `protobuf:"varint,12,opt,name=e_f_o_e,json=eFOE,enum=common.CommonEnum,oneof"`
}
type RootMessage_EFOS struct {
	EFOS int32 `protobuf:"varint,13,opt,name=e_f_o_s,json=eFOS,oneof"`
}
type RootMessage_EFOM struct {
	EFOM *common.CommonMessage `protobuf:"bytes,14,opt,name=e_f_o_m,json=eFOM,oneof"`
}
type RootMessage_EFOEm struct {
	EFOEm *Empty `protobuf:"bytes,15,opt,name=e_f_o_em,json=eFOEm,oneof"`
}
type RootMessage_SFOS struct {
	SFOS int32 `protobuf:"varint,16,opt,name=s_f_o_s,json=sFOS,oneof"`
}
type RootMessage_SFOE struct {
	SFOE RootEnum `protobuf:"varint,17,opt,name=s_f_o_e,json=sFOE,enum=example.RootEnum,oneof"`
}
type RootMessage_SFOMes struct {
	SFOMes *RootMessage2 `protobuf:"bytes,18,opt,name=s_f_o_mes,json=sFOMes,oneof"`
}
type RootMessage_SFOM struct {
	SFOM *Empty `protobuf:"bytes,19,opt,name=s_f_o_m,json=sFOM,oneof"`
}
type RootMessage_MFOM struct {
	MFOM *RootMessage2 `protobuf:"bytes,20,opt,name=m_f_o_m,json=mFOM,oneof"`
}
type RootMessage_MFOS struct {
	MFOS int32 `protobuf:"varint,21,opt,name=m_f_o_s,json=mFOS,oneof"`
}
type RootMessage_MFOE struct {
	MFOE RootEnum `protobuf:"varint,22,opt,name=m_f_o_e,json=mFOE,enum=example.RootEnum,oneof"`
}
type RootMessage_MFOEm struct {
	MFOEm *Empty `protobuf:"bytes,23,opt,name=m_f_o_em,json=mFOEm,oneof"`
}
type RootMessage_EmFOEm struct {
	EmFOEm *Empty `protobuf:"bytes,24,opt,name=em_f_o_em,json=emFOEm,oneof"`
}
type RootMessage_EmFOS struct {
	EmFOS int32 `protobuf:"varint,25,opt,name=em_f_o_s,json=emFOS,oneof"`
}
type RootMessage_EmFOEn struct {
	EmFOEn RootEnum `protobuf:"varint,26,opt,name=em_f_o_en,json=emFOEn,enum=example.RootEnum,oneof"`
}
type RootMessage_EmFOM struct {
	EmFOM *RootMessage2 `protobuf:"bytes,27,opt,name=em_f_o_m,json=emFOM,oneof"`
}

func (*RootMessage_EFOE) isRootMessage_EnumFirstOneoff()     {}
func (*RootMessage_EFOS) isRootMessage_EnumFirstOneoff()     {}
func (*RootMessage_EFOM) isRootMessage_EnumFirstOneoff()     {}
func (*RootMessage_EFOEm) isRootMessage_EnumFirstOneoff()    {}
func (*RootMessage_SFOS) isRootMessage_ScalarFirstOneoff()   {}
func (*RootMessage_SFOE) isRootMessage_ScalarFirstOneoff()   {}
func (*RootMessage_SFOMes) isRootMessage_ScalarFirstOneoff() {}
func (*RootMessage_SFOM) isRootMessage_ScalarFirstOneoff()   {}
func (*RootMessage_MFOM) isRootMessage_MessageFirstOneoff()  {}
func (*RootMessage_MFOS) isRootMessage_MessageFirstOneoff()  {}
func (*RootMessage_MFOE) isRootMessage_MessageFirstOneoff()  {}
func (*RootMessage_MFOEm) isRootMessage_MessageFirstOneoff() {}
func (*RootMessage_EmFOEm) isRootMessage_EmptyFirstOneoff()  {}
func (*RootMessage_EmFOS) isRootMessage_EmptyFirstOneoff()   {}
func (*RootMessage_EmFOEn) isRootMessage_EmptyFirstOneoff()  {}
func (*RootMessage_EmFOM) isRootMessage_EmptyFirstOneoff()   {}

func (m *RootMessage) GetEnumFirstOneoff() isRootMessage_EnumFirstOneoff {
	if m != nil {
		return m.EnumFirstOneoff
	}
	return nil
}
func (m *RootMessage) GetScalarFirstOneoff() isRootMessage_ScalarFirstOneoff {
	if m != nil {
		return m.ScalarFirstOneoff
	}
	return nil
}
func (m *RootMessage) GetMessageFirstOneoff() isRootMessage_MessageFirstOneoff {
	if m != nil {
		return m.MessageFirstOneoff
	}
	return nil
}
func (m *RootMessage) GetEmptyFirstOneoff() isRootMessage_EmptyFirstOneoff {
	if m != nil {
		return m.EmptyFirstOneoff
	}
	return nil
}

func (m *RootMessage) GetMapEnum() map[int32]RootMessage_NestedEnum {
	if m != nil {
		return m.MapEnum
	}
	return nil
}

func (m *RootMessage) GetMapScalar() map[int32]int32 {
	if m != nil {
		return m.MapScalar
	}
	return nil
}

func (m *RootMessage) GetMapMsg() map[int32]*RootMessage_NestedMessage {
	if m != nil {
		return m.MapMsg
	}
	return nil
}

func (m *RootMessage) GetRMsg() []*RootMessage_NestedMessage {
	if m != nil {
		return m.RMsg
	}
	return nil
}

func (m *RootMessage) GetRScalar() []int32 {
	if m != nil {
		return m.RScalar
	}
	return nil
}

func (m *RootMessage) GetREnum() []RootEnum {
	if m != nil {
		return m.REnum
	}
	return nil
}

func (m *RootMessage) GetREmptyMsg() []*Empty {
	if m != nil {
		return m.REmptyMsg
	}
	return nil
}

func (m *RootMessage) GetNREnum() common.CommonEnum {
	if m != nil {
		return m.NREnum
	}
	return common.CommonEnum_CommonEnumVal0
}

func (m *RootMessage) GetNRScalar() int32 {
	if m != nil {
		return m.NRScalar
	}
	return 0
}

func (m *RootMessage) GetNRMsg() *common.CommonMessage {
	if m != nil {
		return m.NRMsg
	}
	return nil
}

func (m *RootMessage) GetScalarFromContext() int32 {
	if m != nil {
		return m.ScalarFromContext
	}
	return 0
}

func (m *RootMessage) GetNREmptyMsg() *Empty {
	if m != nil {
		return m.NREmptyMsg
	}
	return nil
}

func (m *RootMessage) GetEFOE() common.CommonEnum {
	if x, ok := m.GetEnumFirstOneoff().(*RootMessage_EFOE); ok {
		return x.EFOE
	}
	return common.CommonEnum_CommonEnumVal0
}

func (m *RootMessage) GetEFOS() int32 {
	if x, ok := m.GetEnumFirstOneoff().(*RootMessage_EFOS); ok {
		return x.EFOS
	}
	return 0
}

func (m *RootMessage) GetEFOM() *common.CommonMessage {
	if x, ok := m.GetEnumFirstOneoff().(*RootMessage_EFOM); ok {
		return x.EFOM
	}
	return nil
}

func (m *RootMessage) GetEFOEm() *Empty {
	if x, ok := m.GetEnumFirstOneoff().(*RootMessage_EFOEm); ok {
		return x.EFOEm
	}
	return nil
}

func (m *RootMessage) GetSFOS() int32 {
	if x, ok := m.GetScalarFirstOneoff().(*RootMessage_SFOS); ok {
		return x.SFOS
	}
	return 0
}

func (m *RootMessage) GetSFOE() RootEnum {
	if x, ok := m.GetScalarFirstOneoff().(*RootMessage_SFOE); ok {
		return x.SFOE
	}
	return RootEnum_RootEnumVal0
}

func (m *RootMessage) GetSFOMes() *RootMessage2 {
	if x, ok := m.GetScalarFirstOneoff().(*RootMessage_SFOMes); ok {
		return x.SFOMes
	}
	return nil
}

func (m *RootMessage) GetSFOM() *Empty {
	if x, ok := m.GetScalarFirstOneoff().(*RootMessage_SFOM); ok {
		return x.SFOM
	}
	return nil
}

func (m *RootMessage) GetMFOM() *RootMessage2 {
	if x, ok := m.GetMessageFirstOneoff().(*RootMessage_MFOM); ok {
		return x.MFOM
	}
	return nil
}

func (m *RootMessage) GetMFOS() int32 {
	if x, ok := m.GetMessageFirstOneoff().(*RootMessage_MFOS); ok {
		return x.MFOS
	}
	return 0
}

func (m *RootMessage) GetMFOE() RootEnum {
	if x, ok := m.GetMessageFirstOneoff().(*RootMessage_MFOE); ok {
		return x.MFOE
	}
	return RootEnum_RootEnumVal0
}

func (m *RootMessage) GetMFOEm() *Empty {
	if x, ok := m.GetMessageFirstOneoff().(*RootMessage_MFOEm); ok {
		return x.MFOEm
	}
	return nil
}

func (m *RootMessage) GetEmFOEm() *Empty {
	if x, ok := m.GetEmptyFirstOneoff().(*RootMessage_EmFOEm); ok {
		return x.EmFOEm
	}
	return nil
}

func (m *RootMessage) GetEmFOS() int32 {
	if x, ok := m.GetEmptyFirstOneoff().(*RootMessage_EmFOS); ok {
		return x.EmFOS
	}
	return 0
}

func (m *RootMessage) GetEmFOEn() RootEnum {
	if x, ok := m.GetEmptyFirstOneoff().(*RootMessage_EmFOEn); ok {
		return x.EmFOEn
	}
	return RootEnum_RootEnumVal0
}

func (m *RootMessage) GetEmFOM() *RootMessage2 {
	if x, ok := m.GetEmptyFirstOneoff().(*RootMessage_EmFOM); ok {
		return x.EmFOM
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RootMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RootMessage_OneofMarshaler, _RootMessage_OneofUnmarshaler, _RootMessage_OneofSizer, []interface{}{
		(*RootMessage_EFOE)(nil),
		(*RootMessage_EFOS)(nil),
		(*RootMessage_EFOM)(nil),
		(*RootMessage_EFOEm)(nil),
		(*RootMessage_SFOS)(nil),
		(*RootMessage_SFOE)(nil),
		(*RootMessage_SFOMes)(nil),
		(*RootMessage_SFOM)(nil),
		(*RootMessage_MFOM)(nil),
		(*RootMessage_MFOS)(nil),
		(*RootMessage_MFOE)(nil),
		(*RootMessage_MFOEm)(nil),
		(*RootMessage_EmFOEm)(nil),
		(*RootMessage_EmFOS)(nil),
		(*RootMessage_EmFOEn)(nil),
		(*RootMessage_EmFOM)(nil),
	}
}

func _RootMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RootMessage)
	// enum_first_oneoff
	switch x := m.EnumFirstOneoff.(type) {
	case *RootMessage_EFOE:
		b.EncodeVarint(12<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.EFOE))
	case *RootMessage_EFOS:
		b.EncodeVarint(13<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.EFOS))
	case *RootMessage_EFOM:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EFOM); err != nil {
			return err
		}
	case *RootMessage_EFOEm:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EFOEm); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RootMessage.EnumFirstOneoff has unexpected type %T", x)
	}
	// scalar_first_oneoff
	switch x := m.ScalarFirstOneoff.(type) {
	case *RootMessage_SFOS:
		b.EncodeVarint(16<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SFOS))
	case *RootMessage_SFOE:
		b.EncodeVarint(17<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SFOE))
	case *RootMessage_SFOMes:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SFOMes); err != nil {
			return err
		}
	case *RootMessage_SFOM:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SFOM); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RootMessage.ScalarFirstOneoff has unexpected type %T", x)
	}
	// message_first_oneoff
	switch x := m.MessageFirstOneoff.(type) {
	case *RootMessage_MFOM:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MFOM); err != nil {
			return err
		}
	case *RootMessage_MFOS:
		b.EncodeVarint(21<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MFOS))
	case *RootMessage_MFOE:
		b.EncodeVarint(22<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MFOE))
	case *RootMessage_MFOEm:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MFOEm); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RootMessage.MessageFirstOneoff has unexpected type %T", x)
	}
	// empty_first_oneoff
	switch x := m.EmptyFirstOneoff.(type) {
	case *RootMessage_EmFOEm:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EmFOEm); err != nil {
			return err
		}
	case *RootMessage_EmFOS:
		b.EncodeVarint(25<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.EmFOS))
	case *RootMessage_EmFOEn:
		b.EncodeVarint(26<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.EmFOEn))
	case *RootMessage_EmFOM:
		b.EncodeVarint(27<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EmFOM); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RootMessage.EmptyFirstOneoff has unexpected type %T", x)
	}
	return nil
}

func _RootMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RootMessage)
	switch tag {
	case 12: // enum_first_oneoff.e_f_o_e
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.EnumFirstOneoff = &RootMessage_EFOE{common.CommonEnum(x)}
		return true, err
	case 13: // enum_first_oneoff.e_f_o_s
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.EnumFirstOneoff = &RootMessage_EFOS{int32(x)}
		return true, err
	case 14: // enum_first_oneoff.e_f_o_m
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CommonMessage)
		err := b.DecodeMessage(msg)
		m.EnumFirstOneoff = &RootMessage_EFOM{msg}
		return true, err
	case 15: // enum_first_oneoff.e_f_o_em
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.EnumFirstOneoff = &RootMessage_EFOEm{msg}
		return true, err
	case 16: // scalar_first_oneoff.s_f_o_s
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ScalarFirstOneoff = &RootMessage_SFOS{int32(x)}
		return true, err
	case 17: // scalar_first_oneoff.s_f_o_e
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ScalarFirstOneoff = &RootMessage_SFOE{RootEnum(x)}
		return true, err
	case 18: // scalar_first_oneoff.s_f_o_mes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RootMessage2)
		err := b.DecodeMessage(msg)
		m.ScalarFirstOneoff = &RootMessage_SFOMes{msg}
		return true, err
	case 19: // scalar_first_oneoff.s_f_o_m
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.ScalarFirstOneoff = &RootMessage_SFOM{msg}
		return true, err
	case 20: // message_first_oneoff.m_f_o_m
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RootMessage2)
		err := b.DecodeMessage(msg)
		m.MessageFirstOneoff = &RootMessage_MFOM{msg}
		return true, err
	case 21: // message_first_oneoff.m_f_o_s
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MessageFirstOneoff = &RootMessage_MFOS{int32(x)}
		return true, err
	case 22: // message_first_oneoff.m_f_o_e
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MessageFirstOneoff = &RootMessage_MFOE{RootEnum(x)}
		return true, err
	case 23: // message_first_oneoff.m_f_o_em
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.MessageFirstOneoff = &RootMessage_MFOEm{msg}
		return true, err
	case 24: // empty_first_oneoff.em_f_o_em
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.EmptyFirstOneoff = &RootMessage_EmFOEm{msg}
		return true, err
	case 25: // empty_first_oneoff.em_f_o_s
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.EmptyFirstOneoff = &RootMessage_EmFOS{int32(x)}
		return true, err
	case 26: // empty_first_oneoff.em_f_o_en
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.EmptyFirstOneoff = &RootMessage_EmFOEn{RootEnum(x)}
		return true, err
	case 27: // empty_first_oneoff.em_f_o_m
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RootMessage2)
		err := b.DecodeMessage(msg)
		m.EmptyFirstOneoff = &RootMessage_EmFOM{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RootMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RootMessage)
	// enum_first_oneoff
	switch x := m.EnumFirstOneoff.(type) {
	case *RootMessage_EFOE:
		n += proto.SizeVarint(12<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EFOE))
	case *RootMessage_EFOS:
		n += proto.SizeVarint(13<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EFOS))
	case *RootMessage_EFOM:
		s := proto.Size(x.EFOM)
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RootMessage_EFOEm:
		s := proto.Size(x.EFOEm)
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// scalar_first_oneoff
	switch x := m.ScalarFirstOneoff.(type) {
	case *RootMessage_SFOS:
		n += proto.SizeVarint(16<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SFOS))
	case *RootMessage_SFOE:
		n += proto.SizeVarint(17<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SFOE))
	case *RootMessage_SFOMes:
		s := proto.Size(x.SFOMes)
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RootMessage_SFOM:
		s := proto.Size(x.SFOM)
		n += proto.SizeVarint(19<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// message_first_oneoff
	switch x := m.MessageFirstOneoff.(type) {
	case *RootMessage_MFOM:
		s := proto.Size(x.MFOM)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RootMessage_MFOS:
		n += proto.SizeVarint(21<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MFOS))
	case *RootMessage_MFOE:
		n += proto.SizeVarint(22<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MFOE))
	case *RootMessage_MFOEm:
		s := proto.Size(x.MFOEm)
		n += proto.SizeVarint(23<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// empty_first_oneoff
	switch x := m.EmptyFirstOneoff.(type) {
	case *RootMessage_EmFOEm:
		s := proto.Size(x.EmFOEm)
		n += proto.SizeVarint(24<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RootMessage_EmFOS:
		n += proto.SizeVarint(25<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EmFOS))
	case *RootMessage_EmFOEn:
		n += proto.SizeVarint(26<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.EmFOEn))
	case *RootMessage_EmFOM:
		s := proto.Size(x.EmFOM)
		n += proto.SizeVarint(27<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RootMessage_NestedMessage struct {
	SubREnum    []RootMessage_NestedEnum                     `protobuf:"varint,1,rep,packed,name=sub_r_enum,json=subREnum,enum=example.RootMessage_NestedEnum" json:"sub_r_enum,omitempty"`
	SubSubREnum []RootMessage_NestedMessage_NestedNestedEnum `protobuf:"varint,2,rep,packed,name=sub_sub_r_enum,json=subSubREnum,enum=example.RootMessage_NestedMessage_NestedNestedEnum" json:"sub_sub_r_enum,omitempty"`
}

func (m *RootMessage_NestedMessage) Reset()                    { *m = RootMessage_NestedMessage{} }
func (m *RootMessage_NestedMessage) String() string            { return proto.CompactTextString(m) }
func (*RootMessage_NestedMessage) ProtoMessage()               {}
func (*RootMessage_NestedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *RootMessage_NestedMessage) GetSubREnum() []RootMessage_NestedEnum {
	if m != nil {
		return m.SubREnum
	}
	return nil
}

func (m *RootMessage_NestedMessage) GetSubSubREnum() []RootMessage_NestedMessage_NestedNestedEnum {
	if m != nil {
		return m.SubSubREnum
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MessageWithEmpty struct {
	Empt *Empty `protobuf:"bytes,1,opt,name=empt" json:"empt,omitempty"`
}

func (m *MessageWithEmpty) Reset()                    { *m = MessageWithEmpty{} }
func (m *MessageWithEmpty) String() string            { return proto.CompactTextString(m) }
func (*MessageWithEmpty) ProtoMessage()               {}
func (*MessageWithEmpty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MessageWithEmpty) GetEmpt() *Empty {
	if m != nil {
		return m.Empt
	}
	return nil
}

type RootMessage2 struct {
	SomeField int32 `protobuf:"varint,1,opt,name=some_field,json=someField" json:"some_field,omitempty"`
}

func (m *RootMessage2) Reset()                    { *m = RootMessage2{} }
func (m *RootMessage2) String() string            { return proto.CompactTextString(m) }
func (*RootMessage2) ProtoMessage()               {}
func (*RootMessage2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RootMessage2) GetSomeField() int32 {
	if m != nil {
		return m.SomeField
	}
	return 0
}

func init() {
	proto.RegisterType((*RootMessage)(nil), "example.RootMessage")
	proto.RegisterType((*RootMessage_NestedMessage)(nil), "example.RootMessage.NestedMessage")
	proto.RegisterType((*Empty)(nil), "example.Empty")
	proto.RegisterType((*MessageWithEmpty)(nil), "example.MessageWithEmpty")
	proto.RegisterType((*RootMessage2)(nil), "example.RootMessage2")
	proto.RegisterEnum("example.RootEnum", RootEnum_name, RootEnum_value)
	proto.RegisterEnum("example.RootMessage_NestedEnum", RootMessage_NestedEnum_name, RootMessage_NestedEnum_value)
	proto.RegisterEnum("example.RootMessage_NestedMessage_NestedNestedEnum", RootMessage_NestedMessage_NestedNestedEnum_name, RootMessage_NestedMessage_NestedNestedEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServiceExample service

type ServiceExampleClient interface {
	GetQueryMethod(ctx context.Context, in *RootMessage, opts ...grpc.CallOption) (*RootMessage2, error)
	// rpc comment
	MutationMethod(ctx context.Context, in *RootMessage2, opts ...grpc.CallOption) (*RootMessage_NestedMessage, error)
	EmptyMsgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	MsgsWithEpmty(ctx context.Context, in *MessageWithEmpty, opts ...grpc.CallOption) (*MessageWithEmpty, error)
}

type serviceExampleClient struct {
	cc *grpc.ClientConn
}

func NewServiceExampleClient(cc *grpc.ClientConn) ServiceExampleClient {
	return &serviceExampleClient{cc}
}

func (c *serviceExampleClient) GetQueryMethod(ctx context.Context, in *RootMessage, opts ...grpc.CallOption) (*RootMessage2, error) {
	out := new(RootMessage2)
	err := grpc.Invoke(ctx, "/example.ServiceExample/getQueryMethod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceExampleClient) MutationMethod(ctx context.Context, in *RootMessage2, opts ...grpc.CallOption) (*RootMessage_NestedMessage, error) {
	out := new(RootMessage_NestedMessage)
	err := grpc.Invoke(ctx, "/example.ServiceExample/mutationMethod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceExampleClient) EmptyMsgs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/example.ServiceExample/EmptyMsgs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceExampleClient) MsgsWithEpmty(ctx context.Context, in *MessageWithEmpty, opts ...grpc.CallOption) (*MessageWithEmpty, error) {
	out := new(MessageWithEmpty)
	err := grpc.Invoke(ctx, "/example.ServiceExample/MsgsWithEpmty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceExample service

type ServiceExampleServer interface {
	GetQueryMethod(context.Context, *RootMessage) (*RootMessage2, error)
	// rpc comment
	MutationMethod(context.Context, *RootMessage2) (*RootMessage_NestedMessage, error)
	EmptyMsgs(context.Context, *Empty) (*Empty, error)
	MsgsWithEpmty(context.Context, *MessageWithEmpty) (*MessageWithEmpty, error)
}

func RegisterServiceExampleServer(s *grpc.Server, srv ServiceExampleServer) {
	s.RegisterService(&_ServiceExample_serviceDesc, srv)
}

func _ServiceExample_GetQueryMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceExampleServer).GetQueryMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ServiceExample/GetQueryMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceExampleServer).GetQueryMethod(ctx, req.(*RootMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceExample_MutationMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootMessage2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceExampleServer).MutationMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ServiceExample/MutationMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceExampleServer).MutationMethod(ctx, req.(*RootMessage2))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceExample_EmptyMsgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceExampleServer).EmptyMsgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ServiceExample/EmptyMsgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceExampleServer).EmptyMsgs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceExample_MsgsWithEpmty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageWithEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceExampleServer).MsgsWithEpmty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ServiceExample/MsgsWithEpmty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceExampleServer).MsgsWithEpmty(ctx, req.(*MessageWithEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceExample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.ServiceExample",
	HandlerType: (*ServiceExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getQueryMethod",
			Handler:    _ServiceExample_GetQueryMethod_Handler,
		},
		{
			MethodName: "mutationMethod",
			Handler:    _ServiceExample_MutationMethod_Handler,
		},
		{
			MethodName: "EmptyMsgs",
			Handler:    _ServiceExample_EmptyMsgs_Handler,
		},
		{
			MethodName: "MsgsWithEpmty",
			Handler:    _ServiceExample_MsgsWithEpmty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 962 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xce, 0x9d, 0x7d, 0x7e, 0x19, 0x27, 0xd7, 0xcb, 0xc6, 0x69, 0x36, 0xa6, 0x08, 0x63, 0x3e,
	0xd4, 0x6d, 0xa9, 0xd3, 0x5c, 0x0a, 0x14, 0xd4, 0x4a, 0x28, 0x95, 0xa3, 0x48, 0xc8, 0x54, 0x6c,
	0x24, 0x40, 0x20, 0x74, 0xba, 0x24, 0x6b, 0xc7, 0xc2, 0x7b, 0x67, 0x76, 0xf7, 0xaa, 0x9a, 0x5f,
	0xc1, 0x07, 0x7e, 0x02, 0x3f, 0x14, 0xed, 0xcb, 0xb9, 0x67, 0xe7, 0x9c, 0xf6, 0xd3, 0x79, 0x67,
	0x9e, 0x79, 0xe6, 0x99, 0x59, 0xcf, 0xdc, 0x01, 0x48, 0x2a, 0xe4, 0x60, 0xce, 0x53, 0x99, 0xa2,
	0x3a, 0x7d, 0x17, 0xb3, 0xf9, 0x8c, 0x76, 0xf6, 0xae, 0x52, 0xc6, 0xd2, 0xe4, 0xc8, 0x3c, 0x8c,
	0xb7, 0xf7, 0xef, 0x3d, 0x68, 0x91, 0x34, 0x95, 0x23, 0x2a, 0x44, 0x3c, 0xa1, 0xe8, 0x25, 0x34,
	0x58, 0x3c, 0x8f, 0x68, 0x92, 0x31, 0xec, 0x74, 0x2b, 0xfd, 0x56, 0xf8, 0xf9, 0xc0, 0x12, 0x0c,
	0x0a, 0xb8, 0xc1, 0x28, 0x9e, 0x0f, 0x93, 0x8c, 0x0d, 0x13, 0xc9, 0x17, 0xa4, 0xce, 0xcc, 0x09,
	0x9d, 0x02, 0xa8, 0x68, 0x71, 0x15, 0xcf, 0x62, 0x8e, 0x1f, 0xe8, 0xf8, 0x2f, 0x36, 0xc5, 0x5f,
	0x68, 0x94, 0x61, 0x68, 0xb2, 0xfc, 0x8c, 0xbe, 0x05, 0x45, 0x17, 0x31, 0x31, 0xc1, 0xae, 0x26,
	0xe8, 0x6e, 0x22, 0x18, 0x89, 0x89, 0x89, 0xae, 0x31, 0x7d, 0x40, 0xdf, 0x80, 0xc7, 0x75, 0x60,
	0x45, 0x07, 0xf6, 0x4a, 0x03, 0x7f, 0xa4, 0x42, 0xd2, 0x6b, 0x7b, 0x22, 0x55, 0xae, 0x02, 0x0f,
	0xa1, 0xc1, 0x73, 0xd5, 0xd5, 0x6e, 0xa5, 0xef, 0x91, 0x3a, 0xb7, 0x72, 0xfa, 0x50, 0xe3, 0xa6,
	0x1d, 0x5e, 0xb7, 0xd2, 0xf7, 0xc3, 0xdd, 0x15, 0x52, 0x55, 0x35, 0xf1, 0xb8, 0x2e, 0x7e, 0x00,
	0x2d, 0x1e, 0x51, 0x36, 0x97, 0x0b, 0xad, 0xa1, 0xa6, 0x35, 0xf8, 0x4b, 0xf8, 0x50, 0x79, 0x48,
	0x93, 0xeb, 0xa7, 0x4a, 0xfa, 0x25, 0x34, 0x92, 0xc8, 0x72, 0xd7, 0xbb, 0x4e, 0xdf, 0x0f, 0xd1,
	0xc0, 0xde, 0xcd, 0x6b, 0xfd, 0xd0, 0xe4, 0xb5, 0x84, 0x68, 0xf6, 0x07, 0x00, 0x0a, 0x6d, 0x45,
	0x36, 0xba, 0x4e, 0xdf, 0x23, 0x8d, 0x84, 0x58, 0x95, 0x4f, 0xa1, 0xae, 0xbc, 0x2a, 0x6f, 0xb3,
	0xeb, 0xf4, 0x5b, 0xe1, 0xfe, 0x2a, 0x55, 0x5e, 0xae, 0x97, 0x10, 0x95, 0x7a, 0x00, 0x7b, 0x86,
	0x28, 0x1a, 0xf3, 0x94, 0x45, 0x57, 0x69, 0x22, 0xe9, 0x3b, 0x89, 0x41, 0xb3, 0xee, 0x1a, 0xd7,
	0x19, 0x4f, 0xd9, 0x6b, 0xe3, 0x40, 0xc7, 0xb0, 0xa3, 0xa5, 0x2e, 0x8b, 0x6b, 0xe9, 0x24, 0xeb,
	0xc5, 0x41, 0x42, 0x96, 0xd5, 0x3d, 0x81, 0x3a, 0x8d, 0xc6, 0x51, 0x1a, 0x51, 0xbc, 0xbd, 0xa9,
	0xb8, 0xf3, 0x2d, 0x52, 0xa5, 0x67, 0x6f, 0x86, 0xe8, 0x20, 0x07, 0x0b, 0xbc, 0xa3, 0x34, 0x58,
	0xc7, 0x05, 0x1a, 0xe4, 0x0e, 0x86, 0xfd, 0x3b, 0xea, 0xb2, 0xf8, 0x11, 0x7a, 0x04, 0x0d, 0x9b,
	0x95, 0xe1, 0x7b, 0x65, 0x1a, 0xcf, 0xb7, 0x88, 0xa7, 0x52, 0x32, 0x95, 0x53, 0xd8, 0x9c, 0x81,
	0xce, 0xe9, 0x90, 0xaa, 0x50, 0x39, 0x1f, 0xe7, 0x0e, 0x8a, 0x77, 0xb5, 0xf2, 0xdb, 0x57, 0x6e,
	0xb1, 0x43, 0x14, 0x42, 0xd3, 0x60, 0x19, 0x15, 0x18, 0x59, 0x85, 0x25, 0xff, 0xba, 0xf0, 0xdc,
	0x21, 0x35, 0x71, 0xf6, 0x66, 0x44, 0x05, 0x7a, 0x98, 0xf3, 0x33, 0xbc, 0x57, 0x2a, 0xd1, 0x90,
	0x8f, 0x54, 0xf1, 0xcc, 0x02, 0xdb, 0x77, 0x51, 0xbb, 0xa4, 0xca, 0x14, 0xfe, 0x20, 0xc7, 0x0b,
	0xbc, 0xaf, 0x2b, 0x32, 0x0e, 0x5d, 0x11, 0xb3, 0x15, 0xdd, 0xdf, 0x54, 0x91, 0xc1, 0x0e, 0x55,
	0x07, 0x59, 0xde, 0xc1, 0x83, 0x52, 0x79, 0x2e, 0xf1, 0x98, 0xee, 0xe0, 0x13, 0x68, 0xd2, 0x25,
	0x16, 0x97, 0x62, 0x2b, 0xa4, 0x46, 0x0d, 0xf8, 0x10, 0x1a, 0x34, 0x57, 0x77, 0xa8, 0xd5, 0x55,
	0x88, 0x47, 0x99, 0xb9, 0xe4, 0x25, 0x4f, 0x82, 0x3b, 0x9b, 0x04, 0xe6, 0x54, 0x09, 0x7a, 0xb6,
	0xa4, 0x62, 0xf8, 0x93, 0xbb, 0x1a, 0x63, 0x33, 0x8c, 0x3a, 0xff, 0xb9, 0xb0, 0xb3, 0x32, 0xf7,
	0xe8, 0x15, 0x80, 0xc8, 0x2e, 0xf3, 0xf1, 0x73, 0xf4, 0x68, 0x7f, 0x76, 0xc7, 0xbe, 0xd0, 0xb3,
	0xd8, 0x10, 0xd9, 0xa5, 0x99, 0xc6, 0x5f, 0xc1, 0x57, 0xe1, 0x05, 0x0a, 0x57, 0x53, 0x9c, 0x7c,
	0x78, 0xe5, 0xd8, 0x53, 0x81, 0xb6, 0x25, 0xb2, 0xcb, 0x0b, 0xcb, 0xdc, 0xfb, 0x1b, 0x82, 0x75,
	0x00, 0xc2, 0xd0, 0x5e, 0xb7, 0xfd, 0x1c, 0xcf, 0x9e, 0x05, 0x5b, 0x1b, 0x3c, 0xc7, 0x81, 0xb3,
	0xc1, 0x13, 0x06, 0xee, 0x06, 0xcf, 0x49, 0x50, 0xe9, 0xfc, 0x0e, 0xdb, 0xc5, 0xbd, 0x8e, 0x02,
	0xa8, 0xfc, 0x49, 0x17, 0xd8, 0xd1, 0x6b, 0x41, 0xfd, 0x44, 0x5f, 0x81, 0xf7, 0x36, 0x9e, 0x65,
	0x14, 0xbb, 0xfa, 0x9a, 0x3e, 0xd8, 0x31, 0x83, 0xfe, 0xce, 0x7d, 0xe1, 0x74, 0x5e, 0x82, 0xbf,
	0xba, 0xf4, 0x4b, 0xe8, 0xdb, 0x45, 0x7a, 0xaf, 0x18, 0xfd, 0x07, 0xb4, 0x0a, 0x1b, 0xbf, 0x24,
	0xf4, 0x45, 0x31, 0xf4, 0xe3, 0x76, 0xff, 0x7b, 0xfa, 0xde, 0x73, 0x80, 0x42, 0xbf, 0x11, 0xf8,
	0xb7, 0x3a, 0xbd, 0x6e, 0x3b, 0x0e, 0x9c, 0xd3, 0x3d, 0xd8, 0x55, 0x77, 0x1f, 0x8d, 0xa7, 0x5c,
	0xc8, 0x28, 0x4d, 0x68, 0x3a, 0x1e, 0x9f, 0xee, 0xbf, 0xdf, 0xad, 0x45, 0xf3, 0x7d, 0x68, 0x33,
	0x93, 0x77, 0xd5, 0xde, 0x06, 0x64, 0xd6, 0x6a, 0xd1, 0xda, 0xab, 0x83, 0xa7, 0x07, 0xa8, 0xf7,
	0x35, 0x04, 0x56, 0xee, 0x2f, 0x53, 0x79, 0xa3, 0x6d, 0xa8, 0x07, 0x55, 0x15, 0xa2, 0xab, 0xbf,
	0xbd, 0x84, 0xb5, 0xaf, 0xf7, 0x14, 0xb6, 0x8b, 0xa3, 0x80, 0x3e, 0x05, 0x10, 0x29, 0x53, 0xb9,
	0xe9, 0xec, 0xda, 0xf6, 0xad, 0xa9, 0x2c, 0x67, 0xca, 0xf0, 0xf8, 0x7b, 0x68, 0xe4, 0x83, 0x86,
	0x02, 0x13, 0x5a, 0xa8, 0x7d, 0xd5, 0xa2, 0xfe, 0x5d, 0xab, 0x96, 0x30, 0x70, 0xc3, 0x7f, 0x5c,
	0xf0, 0x2f, 0x28, 0x7f, 0x3b, 0xbd, 0xa2, 0x43, 0xa3, 0x07, 0xbd, 0x02, 0x7f, 0x42, 0xe5, 0x4f,
	0x19, 0xe5, 0x8b, 0x11, 0x95, 0x37, 0xe9, 0x35, 0x6a, 0x97, 0xdd, 0x4a, 0xa7, 0x7c, 0x7a, 0xd1,
	0x0f, 0xe0, 0xb3, 0x4c, 0xc6, 0x72, 0xaa, 0xd6, 0xbc, 0x0e, 0x2f, 0x07, 0x76, 0x3e, 0xe2, 0xae,
	0xd5, 0xae, 0xca, 0x5f, 0x4d, 0x02, 0xad, 0xb5, 0xac, 0xb3, 0x76, 0x46, 0x43, 0xd8, 0x51, 0x38,
	0xdd, 0xf1, 0x39, 0x93, 0x0b, 0x74, 0xb8, 0x04, 0xac, 0x5f, 0x46, 0x67, 0xb3, 0xeb, 0xf4, 0xd1,
	0x6f, 0x0f, 0x27, 0x53, 0x79, 0x93, 0x5d, 0xaa, 0x77, 0xd6, 0x91, 0x88, 0x65, 0xc6, 0x93, 0xe7,
	0x94, 0x1f, 0xe9, 0x2f, 0xaf, 0x70, 0xf2, 0xd7, 0xec, 0x48, 0x7d, 0xa6, 0x5d, 0xc7, 0x32, 0xbe,
	0xac, 0x69, 0xdb, 0xc9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0x92, 0x3b, 0x01, 0xb9, 0x09,
	0x00, 0x00,
}
