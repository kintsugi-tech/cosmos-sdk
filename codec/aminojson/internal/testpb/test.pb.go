// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: testpb/test.proto

package testpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnEnum int32

const (
	AnEnum_UNDEFINED AnEnum = 0
	AnEnum_ONE       AnEnum = 1
	AnEnum_TWO       AnEnum = 2
)

// Enum value maps for AnEnum.
var (
	AnEnum_name = map[int32]string{
		0: "UNDEFINED",
		1: "ONE",
		2: "TWO",
	}
	AnEnum_value = map[string]int32{
		"UNDEFINED": 0,
		"ONE":       1,
		"TWO":       2,
	}
)

func (x AnEnum) Enum() *AnEnum {
	p := new(AnEnum)
	*p = x
	return p
}

func (x AnEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testpb_test_proto_enumTypes[0].Descriptor()
}

func (AnEnum) Type() protoreflect.EnumType {
	return &file_testpb_test_proto_enumTypes[0]
}

func (x AnEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnEnum.Descriptor instead.
func (AnEnum) EnumDescriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{0}
}

type ABitOfEverything struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message
	// Generates JSON objects. Message field names are mapped to lowerCamelCase and become JSON object keys. If the json_name field option is specified, the specified value will be used as the key instead. Parsers accept both the lowerCamelCase name (or the one specified by the json_name option) and the original proto field name. null is an accepted value for all field types and treated as the default value of the corresponding field type.
	Message *NestedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// enum
	// The name of the enum value as specified in proto is used. Parsers accept both enum names and integer values.
	Enum AnEnum `protobuf:"varint,2,opt,name=enum,proto3,enum=testpb.AnEnum" json:"enum,omitempty"`
	// map
	// All keys are converted to strings.
	StrMap   map[string]string `protobuf:"bytes,3,rep,name=str_map,json=strMap,proto3" json:"str_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Int32Map map[int32]string  `protobuf:"bytes,4,rep,name=int32_map,json=int32Map,proto3" json:"int32_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BoolMap  map[bool]string   `protobuf:"bytes,5,rep,name=bool_map,json=boolMap,proto3" json:"bool_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// repeated
	Repeated    []int32                 `protobuf:"varint,6,rep,packed,name=repeated,proto3" json:"repeated,omitempty"`
	Str         string                  `protobuf:"bytes,7,opt,name=str,proto3" json:"str,omitempty"`
	Bool        bool                    `protobuf:"varint,8,opt,name=bool,proto3" json:"bool,omitempty"`
	Bytes       []byte                  `protobuf:"bytes,9,opt,name=bytes,proto3" json:"bytes,omitempty"`
	I32         int32                   `protobuf:"varint,10,opt,name=i32,proto3" json:"i32,omitempty"`
	F32         uint32                  `protobuf:"fixed32,11,opt,name=f32,proto3" json:"f32,omitempty"`
	U32         uint32                  `protobuf:"varint,12,opt,name=u32,proto3" json:"u32,omitempty"`
	Si32        int32                   `protobuf:"zigzag32,13,opt,name=si32,proto3" json:"si32,omitempty"`
	Sf32        int32                   `protobuf:"fixed32,14,opt,name=sf32,proto3" json:"sf32,omitempty"`
	I64         int64                   `protobuf:"varint,15,opt,name=i64,proto3" json:"i64,omitempty"`
	F64         uint64                  `protobuf:"fixed64,16,opt,name=f64,proto3" json:"f64,omitempty"`
	U64         uint64                  `protobuf:"varint,17,opt,name=u64,proto3" json:"u64,omitempty"`
	Si64        int64                   `protobuf:"zigzag64,18,opt,name=si64,proto3" json:"si64,omitempty"`
	Sf64        int64                   `protobuf:"fixed64,19,opt,name=sf64,proto3" json:"sf64,omitempty"`
	Float       float32                 `protobuf:"fixed32,20,opt,name=float,proto3" json:"float,omitempty"`
	Double      float64                 `protobuf:"fixed64,21,opt,name=double,proto3" json:"double,omitempty"`
	Any         *anypb.Any              `protobuf:"bytes,22,opt,name=any,proto3" json:"any,omitempty"`
	Timestamp   *timestamppb.Timestamp  `protobuf:"bytes,23,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration    *durationpb.Duration    `protobuf:"bytes,24,opt,name=duration,proto3" json:"duration,omitempty"`
	Struct      *structpb.Struct        `protobuf:"bytes,25,opt,name=struct,proto3" json:"struct,omitempty"`
	BoolValue   *wrapperspb.BoolValue   `protobuf:"bytes,26,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	BytesValue  *wrapperspb.BytesValue  `protobuf:"bytes,27,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	DoubleValue *wrapperspb.DoubleValue `protobuf:"bytes,28,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	FloatValue  *wrapperspb.FloatValue  `protobuf:"bytes,29,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	Int32Value  *wrapperspb.Int32Value  `protobuf:"bytes,30,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value  *wrapperspb.Int64Value  `protobuf:"bytes,31,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,32,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	Uint32Value *wrapperspb.UInt32Value `protobuf:"bytes,33,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value *wrapperspb.UInt64Value `protobuf:"bytes,34,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	FieldMask   *fieldmaskpb.FieldMask  `protobuf:"bytes,35,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	ListValue   *structpb.ListValue     `protobuf:"bytes,36,opt,name=list_value,json=listValue,proto3" json:"list_value,omitempty"`
	Value       *structpb.Value         `protobuf:"bytes,37,opt,name=value,proto3" json:"value,omitempty"`
	NullValue   structpb.NullValue      `protobuf:"varint,38,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue" json:"null_value,omitempty"`
	Empty       *emptypb.Empty          `protobuf:"bytes,39,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *ABitOfEverything) Reset() {
	*x = ABitOfEverything{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ABitOfEverything) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ABitOfEverything) ProtoMessage() {}

func (x *ABitOfEverything) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ABitOfEverything.ProtoReflect.Descriptor instead.
func (*ABitOfEverything) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{0}
}

func (x *ABitOfEverything) GetMessage() *NestedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ABitOfEverything) GetEnum() AnEnum {
	if x != nil {
		return x.Enum
	}
	return AnEnum_UNDEFINED
}

func (x *ABitOfEverything) GetStrMap() map[string]string {
	if x != nil {
		return x.StrMap
	}
	return nil
}

func (x *ABitOfEverything) GetInt32Map() map[int32]string {
	if x != nil {
		return x.Int32Map
	}
	return nil
}

func (x *ABitOfEverything) GetBoolMap() map[bool]string {
	if x != nil {
		return x.BoolMap
	}
	return nil
}

func (x *ABitOfEverything) GetRepeated() []int32 {
	if x != nil {
		return x.Repeated
	}
	return nil
}

func (x *ABitOfEverything) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *ABitOfEverything) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *ABitOfEverything) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *ABitOfEverything) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *ABitOfEverything) GetF32() uint32 {
	if x != nil {
		return x.F32
	}
	return 0
}

func (x *ABitOfEverything) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

func (x *ABitOfEverything) GetSi32() int32 {
	if x != nil {
		return x.Si32
	}
	return 0
}

func (x *ABitOfEverything) GetSf32() int32 {
	if x != nil {
		return x.Sf32
	}
	return 0
}

func (x *ABitOfEverything) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *ABitOfEverything) GetF64() uint64 {
	if x != nil {
		return x.F64
	}
	return 0
}

func (x *ABitOfEverything) GetU64() uint64 {
	if x != nil {
		return x.U64
	}
	return 0
}

func (x *ABitOfEverything) GetSi64() int64 {
	if x != nil {
		return x.Si64
	}
	return 0
}

func (x *ABitOfEverything) GetSf64() int64 {
	if x != nil {
		return x.Sf64
	}
	return 0
}

func (x *ABitOfEverything) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *ABitOfEverything) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *ABitOfEverything) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *ABitOfEverything) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ABitOfEverything) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *ABitOfEverything) GetStruct() *structpb.Struct {
	if x != nil {
		return x.Struct
	}
	return nil
}

func (x *ABitOfEverything) GetBoolValue() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValue
	}
	return nil
}

func (x *ABitOfEverything) GetBytesValue() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *ABitOfEverything) GetDoubleValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValue
	}
	return nil
}

func (x *ABitOfEverything) GetFloatValue() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValue
	}
	return nil
}

func (x *ABitOfEverything) GetInt32Value() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Value
	}
	return nil
}

func (x *ABitOfEverything) GetInt64Value() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Value
	}
	return nil
}

func (x *ABitOfEverything) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

func (x *ABitOfEverything) GetUint32Value() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Value
	}
	return nil
}

func (x *ABitOfEverything) GetUint64Value() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Value
	}
	return nil
}

func (x *ABitOfEverything) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *ABitOfEverything) GetListValue() *structpb.ListValue {
	if x != nil {
		return x.ListValue
	}
	return nil
}

func (x *ABitOfEverything) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ABitOfEverything) GetNullValue() structpb.NullValue {
	if x != nil {
		return x.NullValue
	}
	return structpb.NullValue(0)
}

func (x *ABitOfEverything) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type NestedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar int32  `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *NestedMessage) Reset() {
	*x = NestedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testpb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedMessage) ProtoMessage() {}

func (x *NestedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testpb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedMessage.ProtoReflect.Descriptor instead.
func (*NestedMessage) Descriptor() ([]byte, []int) {
	return file_testpb_test_proto_rawDescGZIP(), []int{1}
}

func (x *NestedMessage) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *NestedMessage) GetBar() int32 {
	if x != nil {
		return x.Bar
	}
	return 0
}

var File_testpb_test_proto protoreflect.FileDescriptor

var file_testpb_test_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x0e, 0x0a, 0x10, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45,
	0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x62, 0x2e, 0x41, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x3d,
	0x0a, 0x07, 0x73, 0x74, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45,
	0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x12, 0x43, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66,
	0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x4d,
	0x61, 0x70, 0x12, 0x40, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x42,
	0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x4d, 0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x66, 0x33, 0x32,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75,
	0x33, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x04, 0x73, 0x69, 0x33, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x66, 0x33, 0x32, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x04, 0x73, 0x66, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x36,
	0x34, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x36, 0x34, 0x18, 0x10, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x66, 0x36, 0x34, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x36, 0x34,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x36, 0x34, 0x18, 0x12, 0x20, 0x01, 0x28, 0x12, 0x52, 0x04,
	0x73, 0x69, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x66, 0x36, 0x34, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x10, 0x52, 0x04, 0x73, 0x66, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x24, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6e,
	0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6e, 0x75, 0x6c,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18,
	0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3b, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c,
	0x42, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33, 0x0a, 0x0d, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x61, 0x72, 0x2a, 0x29, 0x0a,
	0x06, 0x41, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x7f, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x42, 0x09, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x54, 0x65, 0x73, 0x74,
	0x70, 0x62, 0xca, 0x02, 0x06, 0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0xe2, 0x02, 0x12, 0x54, 0x65,
	0x73, 0x74, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x06, 0x54, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_testpb_test_proto_rawDescOnce sync.Once
	file_testpb_test_proto_rawDescData = file_testpb_test_proto_rawDesc
)

func file_testpb_test_proto_rawDescGZIP() []byte {
	file_testpb_test_proto_rawDescOnce.Do(func() {
		file_testpb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testpb_test_proto_rawDescData)
	})
	return file_testpb_test_proto_rawDescData
}

var file_testpb_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_testpb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_testpb_test_proto_goTypes = []interface{}{
	(AnEnum)(0),                    // 0: testpb.AnEnum
	(*ABitOfEverything)(nil),       // 1: testpb.ABitOfEverything
	(*NestedMessage)(nil),          // 2: testpb.NestedMessage
	nil,                            // 3: testpb.ABitOfEverything.StrMapEntry
	nil,                            // 4: testpb.ABitOfEverything.Int32MapEntry
	nil,                            // 5: testpb.ABitOfEverything.BoolMapEntry
	(*anypb.Any)(nil),              // 6: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),    // 8: google.protobuf.Duration
	(*structpb.Struct)(nil),        // 9: google.protobuf.Struct
	(*wrapperspb.BoolValue)(nil),   // 10: google.protobuf.BoolValue
	(*wrapperspb.BytesValue)(nil),  // 11: google.protobuf.BytesValue
	(*wrapperspb.DoubleValue)(nil), // 12: google.protobuf.DoubleValue
	(*wrapperspb.FloatValue)(nil),  // 13: google.protobuf.FloatValue
	(*wrapperspb.Int32Value)(nil),  // 14: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),  // 15: google.protobuf.Int64Value
	(*wrapperspb.StringValue)(nil), // 16: google.protobuf.StringValue
	(*wrapperspb.UInt32Value)(nil), // 17: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 18: google.protobuf.UInt64Value
	(*fieldmaskpb.FieldMask)(nil),  // 19: google.protobuf.FieldMask
	(*structpb.ListValue)(nil),     // 20: google.protobuf.ListValue
	(*structpb.Value)(nil),         // 21: google.protobuf.Value
	(structpb.NullValue)(0),        // 22: google.protobuf.NullValue
	(*emptypb.Empty)(nil),          // 23: google.protobuf.Empty
}
var file_testpb_test_proto_depIdxs = []int32{
	2,  // 0: testpb.ABitOfEverything.message:type_name -> testpb.NestedMessage
	0,  // 1: testpb.ABitOfEverything.enum:type_name -> testpb.AnEnum
	3,  // 2: testpb.ABitOfEverything.str_map:type_name -> testpb.ABitOfEverything.StrMapEntry
	4,  // 3: testpb.ABitOfEverything.int32_map:type_name -> testpb.ABitOfEverything.Int32MapEntry
	5,  // 4: testpb.ABitOfEverything.bool_map:type_name -> testpb.ABitOfEverything.BoolMapEntry
	6,  // 5: testpb.ABitOfEverything.any:type_name -> google.protobuf.Any
	7,  // 6: testpb.ABitOfEverything.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 7: testpb.ABitOfEverything.duration:type_name -> google.protobuf.Duration
	9,  // 8: testpb.ABitOfEverything.struct:type_name -> google.protobuf.Struct
	10, // 9: testpb.ABitOfEverything.bool_value:type_name -> google.protobuf.BoolValue
	11, // 10: testpb.ABitOfEverything.bytes_value:type_name -> google.protobuf.BytesValue
	12, // 11: testpb.ABitOfEverything.double_value:type_name -> google.protobuf.DoubleValue
	13, // 12: testpb.ABitOfEverything.float_value:type_name -> google.protobuf.FloatValue
	14, // 13: testpb.ABitOfEverything.int32_value:type_name -> google.protobuf.Int32Value
	15, // 14: testpb.ABitOfEverything.int64_value:type_name -> google.protobuf.Int64Value
	16, // 15: testpb.ABitOfEverything.string_value:type_name -> google.protobuf.StringValue
	17, // 16: testpb.ABitOfEverything.uint32_value:type_name -> google.protobuf.UInt32Value
	18, // 17: testpb.ABitOfEverything.uint64_value:type_name -> google.protobuf.UInt64Value
	19, // 18: testpb.ABitOfEverything.field_mask:type_name -> google.protobuf.FieldMask
	20, // 19: testpb.ABitOfEverything.list_value:type_name -> google.protobuf.ListValue
	21, // 20: testpb.ABitOfEverything.value:type_name -> google.protobuf.Value
	22, // 21: testpb.ABitOfEverything.null_value:type_name -> google.protobuf.NullValue
	23, // 22: testpb.ABitOfEverything.empty:type_name -> google.protobuf.Empty
	23, // [23:23] is the sub-list for method output_type
	23, // [23:23] is the sub-list for method input_type
	23, // [23:23] is the sub-list for extension type_name
	23, // [23:23] is the sub-list for extension extendee
	0,  // [0:23] is the sub-list for field type_name
}

func init() { file_testpb_test_proto_init() }
func file_testpb_test_proto_init() {
	if File_testpb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testpb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ABitOfEverything); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testpb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testpb_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testpb_test_proto_goTypes,
		DependencyIndexes: file_testpb_test_proto_depIdxs,
		EnumInfos:         file_testpb_test_proto_enumTypes,
		MessageInfos:      file_testpb_test_proto_msgTypes,
	}.Build()
	File_testpb_test_proto = out.File
	file_testpb_test_proto_rawDesc = nil
	file_testpb_test_proto_goTypes = nil
	file_testpb_test_proto_depIdxs = nil
}
