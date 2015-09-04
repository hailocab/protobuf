// Code generated by protoc-gen-go.
// source: test_objects.proto
// DO NOT EDIT!

package jsonpb

import proto "github.com/hailocab/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Widget_Color int32

const (
	Widget_RED   Widget_Color = 0
	Widget_GREEN Widget_Color = 1
	Widget_BLUE  Widget_Color = 2
)

var Widget_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Widget_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Widget_Color) Enum() *Widget_Color {
	p := new(Widget_Color)
	*p = x
	return p
}
func (x Widget_Color) String() string {
	return proto.EnumName(Widget_Color_name, int32(x))
}
func (x *Widget_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Widget_Color_value, data, "Widget_Color")
	if err != nil {
		return err
	}
	*x = Widget_Color(value)
	return nil
}

// Test message for holding primitive types.
type Simple struct {
	OBool            *bool    `protobuf:"varint,1,opt,name=o_bool" json:"o_bool,omitempty"`
	OInt32           *int32   `protobuf:"varint,2,opt,name=o_int32" json:"o_int32,omitempty"`
	OInt64           *int64   `protobuf:"varint,3,opt,name=o_int64" json:"o_int64,omitempty"`
	OUint32          *uint32  `protobuf:"varint,4,opt,name=o_uint32" json:"o_uint32,omitempty"`
	OUint64          *uint64  `protobuf:"varint,5,opt,name=o_uint64" json:"o_uint64,omitempty"`
	OSint32          *int32   `protobuf:"zigzag32,6,opt,name=o_sint32" json:"o_sint32,omitempty"`
	OSint64          *int64   `protobuf:"zigzag64,7,opt,name=o_sint64" json:"o_sint64,omitempty"`
	OFloat           *float32 `protobuf:"fixed32,8,opt,name=o_float" json:"o_float,omitempty"`
	ODouble          *float64 `protobuf:"fixed64,9,opt,name=o_double" json:"o_double,omitempty"`
	OString          *string  `protobuf:"bytes,10,opt,name=o_string" json:"o_string,omitempty"`
	OBytes           []byte   `protobuf:"bytes,11,opt,name=o_bytes" json:"o_bytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}

func (m *Simple) GetOBool() bool {
	if m != nil && m.OBool != nil {
		return *m.OBool
	}
	return false
}

func (m *Simple) GetOInt32() int32 {
	if m != nil && m.OInt32 != nil {
		return *m.OInt32
	}
	return 0
}

func (m *Simple) GetOInt64() int64 {
	if m != nil && m.OInt64 != nil {
		return *m.OInt64
	}
	return 0
}

func (m *Simple) GetOUint32() uint32 {
	if m != nil && m.OUint32 != nil {
		return *m.OUint32
	}
	return 0
}

func (m *Simple) GetOUint64() uint64 {
	if m != nil && m.OUint64 != nil {
		return *m.OUint64
	}
	return 0
}

func (m *Simple) GetOSint32() int32 {
	if m != nil && m.OSint32 != nil {
		return *m.OSint32
	}
	return 0
}

func (m *Simple) GetOSint64() int64 {
	if m != nil && m.OSint64 != nil {
		return *m.OSint64
	}
	return 0
}

func (m *Simple) GetOFloat() float32 {
	if m != nil && m.OFloat != nil {
		return *m.OFloat
	}
	return 0
}

func (m *Simple) GetODouble() float64 {
	if m != nil && m.ODouble != nil {
		return *m.ODouble
	}
	return 0
}

func (m *Simple) GetOString() string {
	if m != nil && m.OString != nil {
		return *m.OString
	}
	return ""
}

func (m *Simple) GetOBytes() []byte {
	if m != nil {
		return m.OBytes
	}
	return nil
}

// Test message for holding repeated primitives.
type Repeats struct {
	RBool            []bool    `protobuf:"varint,1,rep,name=r_bool" json:"r_bool,omitempty"`
	RInt32           []int32   `protobuf:"varint,2,rep,name=r_int32" json:"r_int32,omitempty"`
	RInt64           []int64   `protobuf:"varint,3,rep,name=r_int64" json:"r_int64,omitempty"`
	RUint32          []uint32  `protobuf:"varint,4,rep,name=r_uint32" json:"r_uint32,omitempty"`
	RUint64          []uint64  `protobuf:"varint,5,rep,name=r_uint64" json:"r_uint64,omitempty"`
	RSint32          []int32   `protobuf:"zigzag32,6,rep,name=r_sint32" json:"r_sint32,omitempty"`
	RSint64          []int64   `protobuf:"zigzag64,7,rep,name=r_sint64" json:"r_sint64,omitempty"`
	RFloat           []float32 `protobuf:"fixed32,8,rep,name=r_float" json:"r_float,omitempty"`
	RDouble          []float64 `protobuf:"fixed64,9,rep,name=r_double" json:"r_double,omitempty"`
	RString          []string  `protobuf:"bytes,10,rep,name=r_string" json:"r_string,omitempty"`
	RBytes           [][]byte  `protobuf:"bytes,11,rep,name=r_bytes" json:"r_bytes,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Repeats) Reset()         { *m = Repeats{} }
func (m *Repeats) String() string { return proto.CompactTextString(m) }
func (*Repeats) ProtoMessage()    {}

func (m *Repeats) GetRBool() []bool {
	if m != nil {
		return m.RBool
	}
	return nil
}

func (m *Repeats) GetRInt32() []int32 {
	if m != nil {
		return m.RInt32
	}
	return nil
}

func (m *Repeats) GetRInt64() []int64 {
	if m != nil {
		return m.RInt64
	}
	return nil
}

func (m *Repeats) GetRUint32() []uint32 {
	if m != nil {
		return m.RUint32
	}
	return nil
}

func (m *Repeats) GetRUint64() []uint64 {
	if m != nil {
		return m.RUint64
	}
	return nil
}

func (m *Repeats) GetRSint32() []int32 {
	if m != nil {
		return m.RSint32
	}
	return nil
}

func (m *Repeats) GetRSint64() []int64 {
	if m != nil {
		return m.RSint64
	}
	return nil
}

func (m *Repeats) GetRFloat() []float32 {
	if m != nil {
		return m.RFloat
	}
	return nil
}

func (m *Repeats) GetRDouble() []float64 {
	if m != nil {
		return m.RDouble
	}
	return nil
}

func (m *Repeats) GetRString() []string {
	if m != nil {
		return m.RString
	}
	return nil
}

func (m *Repeats) GetRBytes() [][]byte {
	if m != nil {
		return m.RBytes
	}
	return nil
}

// Test message for holding enums and nested messages.
type Widget struct {
	Color            *Widget_Color  `protobuf:"varint,1,opt,name=color,enum=jsonpb.Widget_Color" json:"color,omitempty"`
	RColor           []Widget_Color `protobuf:"varint,2,rep,name=r_color,enum=jsonpb.Widget_Color" json:"r_color,omitempty"`
	Simple           *Simple        `protobuf:"bytes,10,opt,name=simple" json:"simple,omitempty"`
	RSimple          []*Simple      `protobuf:"bytes,11,rep,name=r_simple" json:"r_simple,omitempty"`
	Repeats          *Repeats       `protobuf:"bytes,20,opt,name=repeats" json:"repeats,omitempty"`
	RRepeats         []*Repeats     `protobuf:"bytes,21,rep,name=r_repeats" json:"r_repeats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Widget) Reset()         { *m = Widget{} }
func (m *Widget) String() string { return proto.CompactTextString(m) }
func (*Widget) ProtoMessage()    {}

func (m *Widget) GetColor() Widget_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Widget_RED
}

func (m *Widget) GetRColor() []Widget_Color {
	if m != nil {
		return m.RColor
	}
	return nil
}

func (m *Widget) GetSimple() *Simple {
	if m != nil {
		return m.Simple
	}
	return nil
}

func (m *Widget) GetRSimple() []*Simple {
	if m != nil {
		return m.RSimple
	}
	return nil
}

func (m *Widget) GetRepeats() *Repeats {
	if m != nil {
		return m.Repeats
	}
	return nil
}

func (m *Widget) GetRRepeats() []*Repeats {
	if m != nil {
		return m.RRepeats
	}
	return nil
}

type Maps struct {
	MInt64Str        map[int64]string `protobuf:"bytes,1,rep,name=m_int64_str" json:"m_int64_str,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MBoolSimple      map[bool]*Simple `protobuf:"bytes,2,rep,name=m_bool_simple" json:"m_bool_simple,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Maps) Reset()         { *m = Maps{} }
func (m *Maps) String() string { return proto.CompactTextString(m) }
func (*Maps) ProtoMessage()    {}

func (m *Maps) GetMInt64Str() map[int64]string {
	if m != nil {
		return m.MInt64Str
	}
	return nil
}

func (m *Maps) GetMBoolSimple() map[bool]*Simple {
	if m != nil {
		return m.MBoolSimple
	}
	return nil
}

type MsgWithOneof struct {
	// Types that are valid to be assigned to Union:
	//	*MsgWithOneof_Title
	//	*MsgWithOneof_Salary
	Union            isMsgWithOneof_Union `protobuf_oneof:"union"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MsgWithOneof) Reset()         { *m = MsgWithOneof{} }
func (m *MsgWithOneof) String() string { return proto.CompactTextString(m) }
func (*MsgWithOneof) ProtoMessage()    {}

type isMsgWithOneof_Union interface {
	isMsgWithOneof_Union()
}

type MsgWithOneof_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title"`
}
type MsgWithOneof_Salary struct {
	Salary int64 `protobuf:"varint,2,opt,name=salary"`
}

func (*MsgWithOneof_Title) isMsgWithOneof_Union()  {}
func (*MsgWithOneof_Salary) isMsgWithOneof_Union() {}

func (m *MsgWithOneof) GetUnion() isMsgWithOneof_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *MsgWithOneof) GetTitle() string {
	if x, ok := m.GetUnion().(*MsgWithOneof_Title); ok {
		return x.Title
	}
	return ""
}

func (m *MsgWithOneof) GetSalary() int64 {
	if x, ok := m.GetUnion().(*MsgWithOneof_Salary); ok {
		return x.Salary
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MsgWithOneof) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _MsgWithOneof_OneofMarshaler, _MsgWithOneof_OneofUnmarshaler, []interface{}{
		(*MsgWithOneof_Title)(nil),
		(*MsgWithOneof_Salary)(nil),
	}
}

func _MsgWithOneof_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MsgWithOneof)
	// union
	switch x := m.Union.(type) {
	case *MsgWithOneof_Title:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Title)
	case *MsgWithOneof_Salary:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Salary))
	case nil:
	default:
		return fmt.Errorf("MsgWithOneof.Union has unexpected type %T", x)
	}
	return nil
}

func _MsgWithOneof_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MsgWithOneof)
	switch tag {
	case 1: // union.title
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Union = &MsgWithOneof_Title{x}
		return true, err
	case 2: // union.salary
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &MsgWithOneof_Salary{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func init() {
	proto.RegisterEnum("jsonpb.Widget_Color", Widget_Color_name, Widget_Color_value)
}
