// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: components/visual_details.proto

package components

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

type VisualDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RangeRings    *RangeRings            `protobuf:"bytes,1,opt,name=range_rings,json=rangeRings,proto3,oneof" json:"range_rings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisualDetails) Reset() {
	*x = VisualDetails{}
	mi := &file_components_visual_details_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisualDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualDetails) ProtoMessage() {}

func (x *VisualDetails) ProtoReflect() protoreflect.Message {
	mi := &file_components_visual_details_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualDetails.ProtoReflect.Descriptor instead.
func (*VisualDetails) Descriptor() ([]byte, []int) {
	return file_components_visual_details_proto_rawDescGZIP(), []int{0}
}

func (x *VisualDetails) GetRangeRings() *RangeRings {
	if x != nil {
		return x.RangeRings
	}
	return nil
}

type RangeRings struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinDistanceM  *float64               `protobuf:"fixed64,1,opt,name=min_distance_m,json=minDistanceM,proto3,oneof" json:"min_distance_m,omitempty"`
	MaxDistanceM  *float64               `protobuf:"fixed64,2,opt,name=max_distance_m,json=maxDistanceM,proto3,oneof" json:"max_distance_m,omitempty"`
	RingCount     *uint32                `protobuf:"varint,3,opt,name=ring_count,json=ringCount,proto3,oneof" json:"ring_count,omitempty"`
	RingLineColor *RingLineColor         `protobuf:"bytes,4,opt,name=ring_line_color,json=ringLineColor,proto3,oneof" json:"ring_line_color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RangeRings) Reset() {
	*x = RangeRings{}
	mi := &file_components_visual_details_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RangeRings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeRings) ProtoMessage() {}

func (x *RangeRings) ProtoReflect() protoreflect.Message {
	mi := &file_components_visual_details_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeRings.ProtoReflect.Descriptor instead.
func (*RangeRings) Descriptor() ([]byte, []int) {
	return file_components_visual_details_proto_rawDescGZIP(), []int{1}
}

func (x *RangeRings) GetMinDistanceM() float64 {
	if x != nil && x.MinDistanceM != nil {
		return *x.MinDistanceM
	}
	return 0
}

func (x *RangeRings) GetMaxDistanceM() float64 {
	if x != nil && x.MaxDistanceM != nil {
		return *x.MaxDistanceM
	}
	return 0
}

func (x *RangeRings) GetRingCount() uint32 {
	if x != nil && x.RingCount != nil {
		return *x.RingCount
	}
	return 0
}

func (x *RangeRings) GetRingLineColor() *RingLineColor {
	if x != nil {
		return x.RingLineColor
	}
	return nil
}

type RingLineColor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Red           *float32               `protobuf:"fixed32,1,opt,name=red,proto3,oneof" json:"red,omitempty"`
	Green         *float32               `protobuf:"fixed32,2,opt,name=green,proto3,oneof" json:"green,omitempty"`
	Blue          *float32               `protobuf:"fixed32,3,opt,name=blue,proto3,oneof" json:"blue,omitempty"`
	Alpha         *float32               `protobuf:"fixed32,4,opt,name=alpha,proto3,oneof" json:"alpha,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RingLineColor) Reset() {
	*x = RingLineColor{}
	mi := &file_components_visual_details_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RingLineColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RingLineColor) ProtoMessage() {}

func (x *RingLineColor) ProtoReflect() protoreflect.Message {
	mi := &file_components_visual_details_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RingLineColor.ProtoReflect.Descriptor instead.
func (*RingLineColor) Descriptor() ([]byte, []int) {
	return file_components_visual_details_proto_rawDescGZIP(), []int{2}
}

func (x *RingLineColor) GetRed() float32 {
	if x != nil && x.Red != nil {
		return *x.Red
	}
	return 0
}

func (x *RingLineColor) GetGreen() float32 {
	if x != nil && x.Green != nil {
		return *x.Green
	}
	return 0
}

func (x *RingLineColor) GetBlue() float32 {
	if x != nil && x.Blue != nil {
		return *x.Blue
	}
	return 0
}

func (x *RingLineColor) GetAlpha() float32 {
	if x != nil && x.Alpha != nil {
		return *x.Alpha
	}
	return 0
}

var File_components_visual_details_proto protoreflect.FileDescriptor

const file_components_visual_details_proto_rawDesc = "" +
	"\n" +
	"\x1fcomponents/visual_details.proto\x12\n" +
	"components\"]\n" +
	"\rVisualDetails\x12<\n" +
	"\vrange_rings\x18\x01 \x01(\v2\x16.components.RangeRingsH\x00R\n" +
	"rangeRings\x88\x01\x01B\x0e\n" +
	"\f_range_rings\"\x97\x02\n" +
	"\n" +
	"RangeRings\x12)\n" +
	"\x0emin_distance_m\x18\x01 \x01(\x01H\x00R\fminDistanceM\x88\x01\x01\x12)\n" +
	"\x0emax_distance_m\x18\x02 \x01(\x01H\x01R\fmaxDistanceM\x88\x01\x01\x12\"\n" +
	"\n" +
	"ring_count\x18\x03 \x01(\rH\x02R\tringCount\x88\x01\x01\x12F\n" +
	"\x0fring_line_color\x18\x04 \x01(\v2\x19.components.RingLineColorH\x03R\rringLineColor\x88\x01\x01B\x11\n" +
	"\x0f_min_distance_mB\x11\n" +
	"\x0f_max_distance_mB\r\n" +
	"\v_ring_countB\x12\n" +
	"\x10_ring_line_color\"\x9a\x01\n" +
	"\rRingLineColor\x12\x15\n" +
	"\x03red\x18\x01 \x01(\x02H\x00R\x03red\x88\x01\x01\x12\x19\n" +
	"\x05green\x18\x02 \x01(\x02H\x01R\x05green\x88\x01\x01\x12\x17\n" +
	"\x04blue\x18\x03 \x01(\x02H\x02R\x04blue\x88\x01\x01\x12\x19\n" +
	"\x05alpha\x18\x04 \x01(\x02H\x03R\x05alpha\x88\x01\x01B\x06\n" +
	"\x04_redB\b\n" +
	"\x06_greenB\a\n" +
	"\x05_blueB\b\n" +
	"\x06_alphaB\rZ\v/componentsb\x06proto3"

var (
	file_components_visual_details_proto_rawDescOnce sync.Once
	file_components_visual_details_proto_rawDescData []byte
)

func file_components_visual_details_proto_rawDescGZIP() []byte {
	file_components_visual_details_proto_rawDescOnce.Do(func() {
		file_components_visual_details_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_components_visual_details_proto_rawDesc), len(file_components_visual_details_proto_rawDesc)))
	})
	return file_components_visual_details_proto_rawDescData
}

var file_components_visual_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_components_visual_details_proto_goTypes = []any{
	(*VisualDetails)(nil), // 0: components.VisualDetails
	(*RangeRings)(nil),    // 1: components.RangeRings
	(*RingLineColor)(nil), // 2: components.RingLineColor
}
var file_components_visual_details_proto_depIdxs = []int32{
	1, // 0: components.VisualDetails.range_rings:type_name -> components.RangeRings
	2, // 1: components.RangeRings.ring_line_color:type_name -> components.RingLineColor
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_components_visual_details_proto_init() }
func file_components_visual_details_proto_init() {
	if File_components_visual_details_proto != nil {
		return
	}
	file_components_visual_details_proto_msgTypes[0].OneofWrappers = []any{}
	file_components_visual_details_proto_msgTypes[1].OneofWrappers = []any{}
	file_components_visual_details_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_components_visual_details_proto_rawDesc), len(file_components_visual_details_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_components_visual_details_proto_goTypes,
		DependencyIndexes: file_components_visual_details_proto_depIdxs,
		MessageInfos:      file_components_visual_details_proto_msgTypes,
	}.Build()
	File_components_visual_details_proto = out.File
	file_components_visual_details_proto_goTypes = nil
	file_components_visual_details_proto_depIdxs = nil
}
