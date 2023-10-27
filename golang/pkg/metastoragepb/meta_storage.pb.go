// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: meta_storage.proto

package metastoragepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcedureInfo_State int32

const (
	ProcedureInfo_STATE_INIT     ProcedureInfo_State = 0
	ProcedureInfo_STATE_RUNNING  ProcedureInfo_State = 1
	ProcedureInfo_STATE_FINISHED ProcedureInfo_State = 2
	ProcedureInfo_STATE_FAILED   ProcedureInfo_State = 3
	ProcedureInfo_STATE_CANCELED ProcedureInfo_State = 4
)

// Enum value maps for ProcedureInfo_State.
var (
	ProcedureInfo_State_name = map[int32]string{
		0: "STATE_INIT",
		1: "STATE_RUNNING",
		2: "STATE_FINISHED",
		3: "STATE_FAILED",
		4: "STATE_CANCELED",
	}
	ProcedureInfo_State_value = map[string]int32{
		"STATE_INIT":     0,
		"STATE_RUNNING":  1,
		"STATE_FINISHED": 2,
		"STATE_FAILED":   3,
		"STATE_CANCELED": 4,
	}
)

func (x ProcedureInfo_State) Enum() *ProcedureInfo_State {
	p := new(ProcedureInfo_State)
	*p = x
	return p
}

func (x ProcedureInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_meta_storage_proto_enumTypes[0].Descriptor()
}

func (ProcedureInfo_State) Type() protoreflect.EnumType {
	return &file_meta_storage_proto_enumTypes[0]
}

func (x ProcedureInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcedureInfo_State.Descriptor instead.
func (ProcedureInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_meta_storage_proto_rawDescGZIP(), []int{1, 0}
}

type ProcedureInfo_Typ int32

const (
	ProcedureInfo_CREATE          ProcedureInfo_Typ = 0
	ProcedureInfo_DELETE          ProcedureInfo_Typ = 1
	ProcedureInfo_TRANSFER_LEADER ProcedureInfo_Typ = 2
	ProcedureInfo_MIGRATE         ProcedureInfo_Typ = 3
	ProcedureInfo_SPLIT           ProcedureInfo_Typ = 4
	ProcedureInfo_MERGE           ProcedureInfo_Typ = 5
	ProcedureInfo_SCATTER         ProcedureInfo_Typ = 6
)

// Enum value maps for ProcedureInfo_Typ.
var (
	ProcedureInfo_Typ_name = map[int32]string{
		0: "CREATE",
		1: "DELETE",
		2: "TRANSFER_LEADER",
		3: "MIGRATE",
		4: "SPLIT",
		5: "MERGE",
		6: "SCATTER",
	}
	ProcedureInfo_Typ_value = map[string]int32{
		"CREATE":          0,
		"DELETE":          1,
		"TRANSFER_LEADER": 2,
		"MIGRATE":         3,
		"SPLIT":           4,
		"MERGE":           5,
		"SCATTER":         6,
	}
)

func (x ProcedureInfo_Typ) Enum() *ProcedureInfo_Typ {
	p := new(ProcedureInfo_Typ)
	*p = x
	return p
}

func (x ProcedureInfo_Typ) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureInfo_Typ) Descriptor() protoreflect.EnumDescriptor {
	return file_meta_storage_proto_enumTypes[1].Descriptor()
}

func (ProcedureInfo_Typ) Type() protoreflect.EnumType {
	return &file_meta_storage_proto_enumTypes[1]
}

func (x ProcedureInfo_Typ) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcedureInfo_Typ.Descriptor instead.
func (ProcedureInfo_Typ) EnumDescriptor() ([]byte, []int) {
	return file_meta_storage_proto_rawDescGZIP(), []int{1, 1}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id             uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint       string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	LeaderPriority int32  `protobuf:"varint,4,opt,name=leader_priority,json=leaderPriority,proto3" json:"leader_priority,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_meta_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_meta_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Member) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Member) GetLeaderPriority() int32 {
	if x != nil {
		return x.LeaderPriority
	}
	return 0
}

type ProcedureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Typ          ProcedureInfo_Typ   `protobuf:"varint,2,opt,name=typ,proto3,enum=meta_storage.ProcedureInfo_Typ" json:"typ,omitempty"`
	State        ProcedureInfo_State `protobuf:"varint,3,opt,name=state,proto3,enum=meta_storage.ProcedureInfo_State" json:"state,omitempty"`
	ExtendedInfo []byte              `protobuf:"bytes,4,opt,name=extended_info,json=extendedInfo,proto3" json:"extended_info,omitempty"`
}

func (x *ProcedureInfo) Reset() {
	*x = ProcedureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureInfo) ProtoMessage() {}

func (x *ProcedureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_meta_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureInfo.ProtoReflect.Descriptor instead.
func (*ProcedureInfo) Descriptor() ([]byte, []int) {
	return file_meta_storage_proto_rawDescGZIP(), []int{1}
}

func (x *ProcedureInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProcedureInfo) GetTyp() ProcedureInfo_Typ {
	if x != nil {
		return x.Typ
	}
	return ProcedureInfo_CREATE
}

func (x *ProcedureInfo) GetState() ProcedureInfo_State {
	if x != nil {
		return x.State
	}
	return ProcedureInfo_STATE_INIT
}

func (x *ProcedureInfo) GetExtendedInfo() []byte {
	if x != nil {
		return x.ExtendedInfo
	}
	return nil
}

var File_meta_storage_proto protoreflect.FileDescriptor

var file_meta_storage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x22, 0x71, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xfa, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x54, 0x79, 0x70, 0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x64, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x62,
	0x0a, 0x03, 0x54, 0x79, 0x70, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45,
	0x52, 0x47, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x43, 0x41, 0x54, 0x54, 0x45, 0x52,
	0x10, 0x06, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x65, 0x72, 0x65, 0x73, 0x44, 0x42, 0x2f, 0x63, 0x65, 0x72, 0x65, 0x73, 0x64, 0x62,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_storage_proto_rawDescOnce sync.Once
	file_meta_storage_proto_rawDescData = file_meta_storage_proto_rawDesc
)

func file_meta_storage_proto_rawDescGZIP() []byte {
	file_meta_storage_proto_rawDescOnce.Do(func() {
		file_meta_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_storage_proto_rawDescData)
	})
	return file_meta_storage_proto_rawDescData
}

var file_meta_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_meta_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_storage_proto_goTypes = []interface{}{
	(ProcedureInfo_State)(0), // 0: meta_storage.ProcedureInfo.State
	(ProcedureInfo_Typ)(0),   // 1: meta_storage.ProcedureInfo.Typ
	(*Member)(nil),           // 2: meta_storage.Member
	(*ProcedureInfo)(nil),    // 3: meta_storage.ProcedureInfo
}
var file_meta_storage_proto_depIdxs = []int32{
	1, // 0: meta_storage.ProcedureInfo.typ:type_name -> meta_storage.ProcedureInfo.Typ
	0, // 1: meta_storage.ProcedureInfo.state:type_name -> meta_storage.ProcedureInfo.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_meta_storage_proto_init() }
func file_meta_storage_proto_init() {
	if File_meta_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_meta_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureInfo); i {
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
			RawDescriptor: file_meta_storage_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_storage_proto_goTypes,
		DependencyIndexes: file_meta_storage_proto_depIdxs,
		EnumInfos:         file_meta_storage_proto_enumTypes,
		MessageInfos:      file_meta_storage_proto_msgTypes,
	}.Build()
	File_meta_storage_proto = out.File
	file_meta_storage_proto_rawDesc = nil
	file_meta_storage_proto_goTypes = nil
	file_meta_storage_proto_depIdxs = nil
}
