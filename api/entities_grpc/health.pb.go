// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: components/health.proto

package components

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ConnectionStatus int32

const (
	ConnectionStatus_CONNECTION_STATUS_INVALID ConnectionStatus = 0
	ConnectionStatus_CONNECTION_STATUS_ONLINE  ConnectionStatus = 1
	ConnectionStatus_CONNECTION_STATUS_OFFLINE ConnectionStatus = 2
)

// Enum value maps for ConnectionStatus.
var (
	ConnectionStatus_name = map[int32]string{
		0: "CONNECTION_STATUS_INVALID",
		1: "CONNECTION_STATUS_ONLINE",
		2: "CONNECTION_STATUS_OFFLINE",
	}
	ConnectionStatus_value = map[string]int32{
		"CONNECTION_STATUS_INVALID": 0,
		"CONNECTION_STATUS_ONLINE":  1,
		"CONNECTION_STATUS_OFFLINE": 2,
	}
)

func (x ConnectionStatus) Enum() *ConnectionStatus {
	p := new(ConnectionStatus)
	*p = x
	return p
}

func (x ConnectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_components_health_proto_enumTypes[0].Descriptor()
}

func (ConnectionStatus) Type() protoreflect.EnumType {
	return &file_components_health_proto_enumTypes[0]
}

func (x ConnectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionStatus.Descriptor instead.
func (ConnectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{0}
}

type HealthStatus int32

const (
	HealthStatus_HEALTH_STATUS_INVALID   HealthStatus = 0
	HealthStatus_HEALTH_STATUS_HEALTHY   HealthStatus = 1
	HealthStatus_HEALTH_STATUS_WARN      HealthStatus = 2
	HealthStatus_HEALTH_STATUS_FAIL      HealthStatus = 3
	HealthStatus_HEALTH_STATUS_OFFLINE   HealthStatus = 4
	HealthStatus_HEALTH_STATUS_NOT_READY HealthStatus = 5
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_INVALID",
		1: "HEALTH_STATUS_HEALTHY",
		2: "HEALTH_STATUS_WARN",
		3: "HEALTH_STATUS_FAIL",
		4: "HEALTH_STATUS_OFFLINE",
		5: "HEALTH_STATUS_NOT_READY",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_INVALID":   0,
		"HEALTH_STATUS_HEALTHY":   1,
		"HEALTH_STATUS_WARN":      2,
		"HEALTH_STATUS_FAIL":      3,
		"HEALTH_STATUS_OFFLINE":   4,
		"HEALTH_STATUS_NOT_READY": 5,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_components_health_proto_enumTypes[1].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_components_health_proto_enumTypes[1]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{1}
}

type AlertLevel int32

const (
	AlertLevel_ALERT_LEVEL_INVALID  AlertLevel = 0
	AlertLevel_ALERT_LEVEL_ADVISORY AlertLevel = 1
	AlertLevel_ALERT_LEVEL_CAUTION  AlertLevel = 2
	AlertLevel_ALERT_LEVEL_WARNING  AlertLevel = 3
)

// Enum value maps for AlertLevel.
var (
	AlertLevel_name = map[int32]string{
		0: "ALERT_LEVEL_INVALID",
		1: "ALERT_LEVEL_ADVISORY",
		2: "ALERT_LEVEL_CAUTION",
		3: "ALERT_LEVEL_WARNING",
	}
	AlertLevel_value = map[string]int32{
		"ALERT_LEVEL_INVALID":  0,
		"ALERT_LEVEL_ADVISORY": 1,
		"ALERT_LEVEL_CAUTION":  2,
		"ALERT_LEVEL_WARNING":  3,
	}
)

func (x AlertLevel) Enum() *AlertLevel {
	p := new(AlertLevel)
	*p = x
	return p
}

func (x AlertLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_components_health_proto_enumTypes[2].Descriptor()
}

func (AlertLevel) Type() protoreflect.EnumType {
	return &file_components_health_proto_enumTypes[2]
}

func (x AlertLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertLevel.Descriptor instead.
func (AlertLevel) EnumDescriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{2}
}

type Health struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ConnectionStatus *ConnectionStatus      `protobuf:"varint,1,opt,name=connectionStatus,proto3,enum=components.ConnectionStatus,oneof" json:"connectionStatus,omitempty"`
	HealthStatus     *HealthStatus          `protobuf:"varint,2,opt,name=healthStatus,proto3,enum=components.HealthStatus,oneof" json:"healthStatus,omitempty"`
	Components       []*Component           `protobuf:"bytes,3,rep,name=components,proto3" json:"components,omitempty"`
	UpdateTime       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updateTime,proto3,oneof" json:"updateTime,omitempty"`
	ActiveAlerts     []*ActiveAlert         `protobuf:"bytes,5,rep,name=activeAlerts,proto3" json:"activeAlerts,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Health) Reset() {
	*x = Health{}
	mi := &file_components_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_components_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{0}
}

func (x *Health) GetConnectionStatus() ConnectionStatus {
	if x != nil && x.ConnectionStatus != nil {
		return *x.ConnectionStatus
	}
	return ConnectionStatus_CONNECTION_STATUS_INVALID
}

func (x *Health) GetHealthStatus() HealthStatus {
	if x != nil && x.HealthStatus != nil {
		return *x.HealthStatus
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *Health) GetComponents() []*Component {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *Health) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Health) GetActiveAlerts() []*ActiveAlert {
	if x != nil {
		return x.ActiveAlerts
	}
	return nil
}

type Component struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Health        *HealthStatus          `protobuf:"varint,3,opt,name=health,proto3,enum=components.HealthStatus,oneof" json:"health,omitempty"`
	Messages      []*Message             `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updateTime,proto3,oneof" json:"updateTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Component) Reset() {
	*x = Component{}
	mi := &file_components_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_components_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{1}
}

func (x *Component) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Component) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Component) GetHealth() HealthStatus {
	if x != nil && x.Health != nil {
		return *x.Health
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *Component) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Component) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *HealthStatus          `protobuf:"varint,1,opt,name=status,proto3,enum=components.HealthStatus,oneof" json:"status,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_components_health_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_components_health_proto_msgTypes[2]
	if x != nil {
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
	return file_components_health_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetStatus() HealthStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return HealthStatus_HEALTH_STATUS_INVALID
}

func (x *Message) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ActiveAlert struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AlertCode        *string                `protobuf:"bytes,1,opt,name=alertCode,proto3,oneof" json:"alertCode,omitempty"`
	Description      *string                `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Level            *AlertLevel            `protobuf:"varint,3,opt,name=level,proto3,enum=components.AlertLevel,oneof" json:"level,omitempty"`
	ActivatedTime    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=activatedTime,proto3,oneof" json:"activatedTime,omitempty"`
	ActiveConditions []*ActiveCondition     `protobuf:"bytes,5,rep,name=activeConditions,proto3" json:"activeConditions,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ActiveAlert) Reset() {
	*x = ActiveAlert{}
	mi := &file_components_health_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveAlert) ProtoMessage() {}

func (x *ActiveAlert) ProtoReflect() protoreflect.Message {
	mi := &file_components_health_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveAlert.ProtoReflect.Descriptor instead.
func (*ActiveAlert) Descriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{3}
}

func (x *ActiveAlert) GetAlertCode() string {
	if x != nil && x.AlertCode != nil {
		return *x.AlertCode
	}
	return ""
}

func (x *ActiveAlert) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ActiveAlert) GetLevel() AlertLevel {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return AlertLevel_ALERT_LEVEL_INVALID
}

func (x *ActiveAlert) GetActivatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ActivatedTime
	}
	return nil
}

func (x *ActiveAlert) GetActiveConditions() []*ActiveCondition {
	if x != nil {
		return x.ActiveConditions
	}
	return nil
}

type ActiveCondition struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConditionCode *string                `protobuf:"bytes,1,opt,name=conditionCode,proto3,oneof" json:"conditionCode,omitempty"`
	Description   *string                `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActiveCondition) Reset() {
	*x = ActiveCondition{}
	mi := &file_components_health_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActiveCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveCondition) ProtoMessage() {}

func (x *ActiveCondition) ProtoReflect() protoreflect.Message {
	mi := &file_components_health_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveCondition.ProtoReflect.Descriptor instead.
func (*ActiveCondition) Descriptor() ([]byte, []int) {
	return file_components_health_proto_rawDescGZIP(), []int{4}
}

func (x *ActiveCondition) GetConditionCode() string {
	if x != nil && x.ConditionCode != nil {
		return *x.ConditionCode
	}
	return ""
}

func (x *ActiveCondition) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

var File_components_health_proto protoreflect.FileDescriptor

const file_components_health_proto_rawDesc = "" +
	"\n" +
	"\x17components/health.proto\x12\n" +
	"components\x1a\x1fgoogle/protobuf/timestamp.proto\"\x84\x03\n" +
	"\x06Health\x12M\n" +
	"\x10connectionStatus\x18\x01 \x01(\x0e2\x1c.components.ConnectionStatusH\x00R\x10connectionStatus\x88\x01\x01\x12A\n" +
	"\fhealthStatus\x18\x02 \x01(\x0e2\x18.components.HealthStatusH\x01R\fhealthStatus\x88\x01\x01\x125\n" +
	"\n" +
	"components\x18\x03 \x03(\v2\x15.components.ComponentR\n" +
	"components\x12?\n" +
	"\n" +
	"updateTime\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampH\x02R\n" +
	"updateTime\x88\x01\x01\x12;\n" +
	"\factiveAlerts\x18\x05 \x03(\v2\x17.components.ActiveAlertR\factiveAlertsB\x13\n" +
	"\x11_connectionStatusB\x0f\n" +
	"\r_healthStatusB\r\n" +
	"\v_updateTime\"\x8c\x02\n" +
	"\tComponent\x12\x13\n" +
	"\x02id\x18\x01 \x01(\tH\x00R\x02id\x88\x01\x01\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x01R\x04name\x88\x01\x01\x125\n" +
	"\x06health\x18\x03 \x01(\x0e2\x18.components.HealthStatusH\x02R\x06health\x88\x01\x01\x12/\n" +
	"\bmessages\x18\x04 \x03(\v2\x13.components.MessageR\bmessages\x12?\n" +
	"\n" +
	"updateTime\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampH\x03R\n" +
	"updateTime\x88\x01\x01B\x05\n" +
	"\x03_idB\a\n" +
	"\x05_nameB\t\n" +
	"\a_healthB\r\n" +
	"\v_updateTime\"v\n" +
	"\aMessage\x125\n" +
	"\x06status\x18\x01 \x01(\x0e2\x18.components.HealthStatusH\x00R\x06status\x88\x01\x01\x12\x1d\n" +
	"\amessage\x18\x02 \x01(\tH\x01R\amessage\x88\x01\x01B\t\n" +
	"\a_statusB\n" +
	"\n" +
	"\b_message\"\xd4\x02\n" +
	"\vActiveAlert\x12!\n" +
	"\talertCode\x18\x01 \x01(\tH\x00R\talertCode\x88\x01\x01\x12%\n" +
	"\vdescription\x18\x02 \x01(\tH\x01R\vdescription\x88\x01\x01\x121\n" +
	"\x05level\x18\x03 \x01(\x0e2\x16.components.AlertLevelH\x02R\x05level\x88\x01\x01\x12E\n" +
	"\ractivatedTime\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampH\x03R\ractivatedTime\x88\x01\x01\x12G\n" +
	"\x10activeConditions\x18\x05 \x03(\v2\x1b.components.ActiveConditionR\x10activeConditionsB\f\n" +
	"\n" +
	"_alertCodeB\x0e\n" +
	"\f_descriptionB\b\n" +
	"\x06_levelB\x10\n" +
	"\x0e_activatedTime\"\x85\x01\n" +
	"\x0fActiveCondition\x12)\n" +
	"\rconditionCode\x18\x01 \x01(\tH\x00R\rconditionCode\x88\x01\x01\x12%\n" +
	"\vdescription\x18\x02 \x01(\tH\x01R\vdescription\x88\x01\x01B\x10\n" +
	"\x0e_conditionCodeB\x0e\n" +
	"\f_description*n\n" +
	"\x10ConnectionStatus\x12\x1d\n" +
	"\x19CONNECTION_STATUS_INVALID\x10\x00\x12\x1c\n" +
	"\x18CONNECTION_STATUS_ONLINE\x10\x01\x12\x1d\n" +
	"\x19CONNECTION_STATUS_OFFLINE\x10\x02*\xac\x01\n" +
	"\fHealthStatus\x12\x19\n" +
	"\x15HEALTH_STATUS_INVALID\x10\x00\x12\x19\n" +
	"\x15HEALTH_STATUS_HEALTHY\x10\x01\x12\x16\n" +
	"\x12HEALTH_STATUS_WARN\x10\x02\x12\x16\n" +
	"\x12HEALTH_STATUS_FAIL\x10\x03\x12\x19\n" +
	"\x15HEALTH_STATUS_OFFLINE\x10\x04\x12\x1b\n" +
	"\x17HEALTH_STATUS_NOT_READY\x10\x05*q\n" +
	"\n" +
	"AlertLevel\x12\x17\n" +
	"\x13ALERT_LEVEL_INVALID\x10\x00\x12\x18\n" +
	"\x14ALERT_LEVEL_ADVISORY\x10\x01\x12\x17\n" +
	"\x13ALERT_LEVEL_CAUTION\x10\x02\x12\x17\n" +
	"\x13ALERT_LEVEL_WARNING\x10\x03B\rZ\v/componentsb\x06proto3"

var (
	file_components_health_proto_rawDescOnce sync.Once
	file_components_health_proto_rawDescData []byte
)

func file_components_health_proto_rawDescGZIP() []byte {
	file_components_health_proto_rawDescOnce.Do(func() {
		file_components_health_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_components_health_proto_rawDesc), len(file_components_health_proto_rawDesc)))
	})
	return file_components_health_proto_rawDescData
}

var file_components_health_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_components_health_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_components_health_proto_goTypes = []any{
	(ConnectionStatus)(0),         // 0: components.ConnectionStatus
	(HealthStatus)(0),             // 1: components.HealthStatus
	(AlertLevel)(0),               // 2: components.AlertLevel
	(*Health)(nil),                // 3: components.Health
	(*Component)(nil),             // 4: components.Component
	(*Message)(nil),               // 5: components.Message
	(*ActiveAlert)(nil),           // 6: components.ActiveAlert
	(*ActiveCondition)(nil),       // 7: components.ActiveCondition
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_components_health_proto_depIdxs = []int32{
	0,  // 0: components.Health.connectionStatus:type_name -> components.ConnectionStatus
	1,  // 1: components.Health.healthStatus:type_name -> components.HealthStatus
	4,  // 2: components.Health.components:type_name -> components.Component
	8,  // 3: components.Health.updateTime:type_name -> google.protobuf.Timestamp
	6,  // 4: components.Health.activeAlerts:type_name -> components.ActiveAlert
	1,  // 5: components.Component.health:type_name -> components.HealthStatus
	5,  // 6: components.Component.messages:type_name -> components.Message
	8,  // 7: components.Component.updateTime:type_name -> google.protobuf.Timestamp
	1,  // 8: components.Message.status:type_name -> components.HealthStatus
	2,  // 9: components.ActiveAlert.level:type_name -> components.AlertLevel
	8,  // 10: components.ActiveAlert.activatedTime:type_name -> google.protobuf.Timestamp
	7,  // 11: components.ActiveAlert.activeConditions:type_name -> components.ActiveCondition
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_components_health_proto_init() }
func file_components_health_proto_init() {
	if File_components_health_proto != nil {
		return
	}
	file_components_health_proto_msgTypes[0].OneofWrappers = []any{}
	file_components_health_proto_msgTypes[1].OneofWrappers = []any{}
	file_components_health_proto_msgTypes[2].OneofWrappers = []any{}
	file_components_health_proto_msgTypes[3].OneofWrappers = []any{}
	file_components_health_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_components_health_proto_rawDesc), len(file_components_health_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_components_health_proto_goTypes,
		DependencyIndexes: file_components_health_proto_depIdxs,
		EnumInfos:         file_components_health_proto_enumTypes,
		MessageInfos:      file_components_health_proto_msgTypes,
	}.Build()
	File_components_health_proto = out.File
	file_components_health_proto_goTypes = nil
	file_components_health_proto_depIdxs = nil
}
