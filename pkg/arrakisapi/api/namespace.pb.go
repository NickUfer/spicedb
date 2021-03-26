// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: REDACTEDapi/api/namespace.proto

package REDACTEDapi

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ComputedUserset_Object int32

const (
	ComputedUserset_TUPLE_OBJECT         ComputedUserset_Object = 0
	ComputedUserset_TUPLE_USERSET_OBJECT ComputedUserset_Object = 1
)

// Enum value maps for ComputedUserset_Object.
var (
	ComputedUserset_Object_name = map[int32]string{
		0: "TUPLE_OBJECT",
		1: "TUPLE_USERSET_OBJECT",
	}
	ComputedUserset_Object_value = map[string]int32{
		"TUPLE_OBJECT":         0,
		"TUPLE_USERSET_OBJECT": 1,
	}
)

func (x ComputedUserset_Object) Enum() *ComputedUserset_Object {
	p := new(ComputedUserset_Object)
	*p = x
	return p
}

func (x ComputedUserset_Object) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComputedUserset_Object) Descriptor() protoreflect.EnumDescriptor {
	return file_REDACTEDapi_api_namespace_proto_enumTypes[0].Descriptor()
}

func (ComputedUserset_Object) Type() protoreflect.EnumType {
	return &file_REDACTEDapi_api_namespace_proto_enumTypes[0]
}

func (x ComputedUserset_Object) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComputedUserset_Object.Descriptor instead.
func (ComputedUserset_Object) EnumDescriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{5, 0}
}

type NamespaceDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Relation []*Relation `protobuf:"bytes,2,rep,name=relation,proto3" json:"relation,omitempty"`
}

func (x *NamespaceDefinition) Reset() {
	*x = NamespaceDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceDefinition) ProtoMessage() {}

func (x *NamespaceDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceDefinition.ProtoReflect.Descriptor instead.
func (*NamespaceDefinition) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *NamespaceDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespaceDefinition) GetRelation() []*Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UsersetRewrite *UsersetRewrite `protobuf:"bytes,2,opt,name=userset_rewrite,json=usersetRewrite,proto3" json:"userset_rewrite,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *Relation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Relation) GetUsersetRewrite() *UsersetRewrite {
	if x != nil {
		return x.UsersetRewrite
	}
	return nil
}

type UsersetRewrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RewriteOperation:
	//	*UsersetRewrite_Union
	//	*UsersetRewrite_Intersection
	//	*UsersetRewrite_Exclusion
	RewriteOperation isUsersetRewrite_RewriteOperation `protobuf_oneof:"rewrite_operation"`
}

func (x *UsersetRewrite) Reset() {
	*x = UsersetRewrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersetRewrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersetRewrite) ProtoMessage() {}

func (x *UsersetRewrite) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersetRewrite.ProtoReflect.Descriptor instead.
func (*UsersetRewrite) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{2}
}

func (m *UsersetRewrite) GetRewriteOperation() isUsersetRewrite_RewriteOperation {
	if m != nil {
		return m.RewriteOperation
	}
	return nil
}

func (x *UsersetRewrite) GetUnion() *SetOperation {
	if x, ok := x.GetRewriteOperation().(*UsersetRewrite_Union); ok {
		return x.Union
	}
	return nil
}

func (x *UsersetRewrite) GetIntersection() *SetOperation {
	if x, ok := x.GetRewriteOperation().(*UsersetRewrite_Intersection); ok {
		return x.Intersection
	}
	return nil
}

func (x *UsersetRewrite) GetExclusion() *SetOperation {
	if x, ok := x.GetRewriteOperation().(*UsersetRewrite_Exclusion); ok {
		return x.Exclusion
	}
	return nil
}

type isUsersetRewrite_RewriteOperation interface {
	isUsersetRewrite_RewriteOperation()
}

type UsersetRewrite_Union struct {
	Union *SetOperation `protobuf:"bytes,1,opt,name=union,proto3,oneof"`
}

type UsersetRewrite_Intersection struct {
	Intersection *SetOperation `protobuf:"bytes,2,opt,name=intersection,proto3,oneof"`
}

type UsersetRewrite_Exclusion struct {
	Exclusion *SetOperation `protobuf:"bytes,3,opt,name=exclusion,proto3,oneof"`
}

func (*UsersetRewrite_Union) isUsersetRewrite_RewriteOperation() {}

func (*UsersetRewrite_Intersection) isUsersetRewrite_RewriteOperation() {}

func (*UsersetRewrite_Exclusion) isUsersetRewrite_RewriteOperation() {}

type SetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Child []*SetOperation_Child `protobuf:"bytes,1,rep,name=child,proto3" json:"child,omitempty"`
}

func (x *SetOperation) Reset() {
	*x = SetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperation) ProtoMessage() {}

func (x *SetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperation.ProtoReflect.Descriptor instead.
func (*SetOperation) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *SetOperation) GetChild() []*SetOperation_Child {
	if x != nil {
		return x.Child
	}
	return nil
}

type TupleToUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tupleset        *TupleToUserset_Tupleset `protobuf:"bytes,1,opt,name=tupleset,proto3" json:"tupleset,omitempty"`
	ComputedUserset *ComputedUserset         `protobuf:"bytes,2,opt,name=computed_userset,json=computedUserset,proto3" json:"computed_userset,omitempty"`
}

func (x *TupleToUserset) Reset() {
	*x = TupleToUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleToUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleToUserset) ProtoMessage() {}

func (x *TupleToUserset) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TupleToUserset.ProtoReflect.Descriptor instead.
func (*TupleToUserset) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *TupleToUserset) GetTupleset() *TupleToUserset_Tupleset {
	if x != nil {
		return x.Tupleset
	}
	return nil
}

func (x *TupleToUserset) GetComputedUserset() *ComputedUserset {
	if x != nil {
		return x.ComputedUserset
	}
	return nil
}

type ComputedUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object   ComputedUserset_Object `protobuf:"varint,1,opt,name=object,proto3,enum=ComputedUserset_Object" json:"object,omitempty"`
	Relation string                 `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ComputedUserset) Reset() {
	*x = ComputedUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedUserset) ProtoMessage() {}

func (x *ComputedUserset) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedUserset.ProtoReflect.Descriptor instead.
func (*ComputedUserset) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *ComputedUserset) GetObject() ComputedUserset_Object {
	if x != nil {
		return x.Object
	}
	return ComputedUserset_TUPLE_OBJECT
}

func (x *ComputedUserset) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type SetOperation_Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ChildType:
	//	*SetOperation_Child_XThis
	//	*SetOperation_Child_ComputedUserset
	//	*SetOperation_Child_TupleToUserset
	//	*SetOperation_Child_UsersetRewrite
	ChildType isSetOperation_Child_ChildType `protobuf_oneof:"child_type"`
}

func (x *SetOperation_Child) Reset() {
	*x = SetOperation_Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOperation_Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperation_Child) ProtoMessage() {}

func (x *SetOperation_Child) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperation_Child.ProtoReflect.Descriptor instead.
func (*SetOperation_Child) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{3, 0}
}

func (m *SetOperation_Child) GetChildType() isSetOperation_Child_ChildType {
	if m != nil {
		return m.ChildType
	}
	return nil
}

func (x *SetOperation_Child) GetXThis() *SetOperation_Child_This {
	if x, ok := x.GetChildType().(*SetOperation_Child_XThis); ok {
		return x.XThis
	}
	return nil
}

func (x *SetOperation_Child) GetComputedUserset() *ComputedUserset {
	if x, ok := x.GetChildType().(*SetOperation_Child_ComputedUserset); ok {
		return x.ComputedUserset
	}
	return nil
}

func (x *SetOperation_Child) GetTupleToUserset() *TupleToUserset {
	if x, ok := x.GetChildType().(*SetOperation_Child_TupleToUserset); ok {
		return x.TupleToUserset
	}
	return nil
}

func (x *SetOperation_Child) GetUsersetRewrite() *UsersetRewrite {
	if x, ok := x.GetChildType().(*SetOperation_Child_UsersetRewrite); ok {
		return x.UsersetRewrite
	}
	return nil
}

type isSetOperation_Child_ChildType interface {
	isSetOperation_Child_ChildType()
}

type SetOperation_Child_XThis struct {
	XThis *SetOperation_Child_This `protobuf:"bytes,1,opt,name=_this,json=This,proto3,oneof"`
}

type SetOperation_Child_ComputedUserset struct {
	ComputedUserset *ComputedUserset `protobuf:"bytes,2,opt,name=computed_userset,json=computedUserset,proto3,oneof"`
}

type SetOperation_Child_TupleToUserset struct {
	TupleToUserset *TupleToUserset `protobuf:"bytes,3,opt,name=tuple_to_userset,json=tupleToUserset,proto3,oneof"`
}

type SetOperation_Child_UsersetRewrite struct {
	UsersetRewrite *UsersetRewrite `protobuf:"bytes,4,opt,name=userset_rewrite,json=usersetRewrite,proto3,oneof"`
}

func (*SetOperation_Child_XThis) isSetOperation_Child_ChildType() {}

func (*SetOperation_Child_ComputedUserset) isSetOperation_Child_ChildType() {}

func (*SetOperation_Child_TupleToUserset) isSetOperation_Child_ChildType() {}

func (*SetOperation_Child_UsersetRewrite) isSetOperation_Child_ChildType() {}

type SetOperation_Child_This struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetOperation_Child_This) Reset() {
	*x = SetOperation_Child_This{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOperation_Child_This) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOperation_Child_This) ProtoMessage() {}

func (x *SetOperation_Child_This) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOperation_Child_This.ProtoReflect.Descriptor instead.
func (*SetOperation_Child_This) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{3, 0, 0}
}

type TupleToUserset_Tupleset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation string `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *TupleToUserset_Tupleset) Reset() {
	*x = TupleToUserset_Tupleset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleToUserset_Tupleset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleToUserset_Tupleset) ProtoMessage() {}

func (x *TupleToUserset_Tupleset) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_api_namespace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TupleToUserset_Tupleset.ProtoReflect.Descriptor instead.
func (*TupleToUserset_Tupleset) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_api_namespace_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TupleToUserset_Tupleset) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

var File_REDACTEDapi_api_namespace_proto protoreflect.FileDescriptor

var file_REDACTEDapi_api_namespace_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x72, 0x72, 0x61, 0x6b, 0x69, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x48, 0xfa, 0x42, 0x45, 0x72, 0x43, 0x28, 0x80, 0x01, 0x32, 0x3e, 0x5e, 0x28, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2f, 0x29, 0x3f, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x28, 0x40, 0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d,
	0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32,
	0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x2f, 0x0a,
	0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x3d,
	0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52,
	0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a,
	0x09, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x18, 0x0a, 0x11, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x22, 0xfc, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x40, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x42, 0x15, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0xfa,
	0x42, 0x0a, 0x92, 0x01, 0x07, 0x22, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x1a, 0xa9, 0x02, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x2f, 0x0a,
	0x05, 0x5f, 0x74, 0x68, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x53,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x2e, 0x54, 0x68, 0x69, 0x73, 0x48, 0x00, 0x52, 0x04, 0x54, 0x68, 0x69, 0x73, 0x12, 0x47,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0e,
	0x74, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x44,
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x52, 0x65, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x1a, 0x06, 0x0a, 0x04, 0x54, 0x68, 0x69, 0x73, 0x42, 0x11, 0x0a, 0x0a,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22,
	0xe8, 0x01, 0x0a, 0x0e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x74, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x1a, 0x4f, 0x0a, 0x08, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x28, 0x40,
	0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f,
	0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x39,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x42, 0x24,
	0x72, 0x22, 0x28, 0x40, 0x32, 0x1e, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a,
	0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b, 0x32, 0x2c, 0x36, 0x32, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x24, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34,
	0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x55, 0x50, 0x4c,
	0x45, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x55,
	0x50, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x45, 0x54, 0x5f, 0x4f, 0x42, 0x4a, 0x45,
	0x43, 0x54, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x72, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x61, 0x72, 0x72, 0x61, 0x6b, 0x69, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_REDACTEDapi_api_namespace_proto_rawDescOnce sync.Once
	file_REDACTEDapi_api_namespace_proto_rawDescData = file_REDACTEDapi_api_namespace_proto_rawDesc
)

func file_REDACTEDapi_api_namespace_proto_rawDescGZIP() []byte {
	file_REDACTEDapi_api_namespace_proto_rawDescOnce.Do(func() {
		file_REDACTEDapi_api_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_REDACTEDapi_api_namespace_proto_rawDescData)
	})
	return file_REDACTEDapi_api_namespace_proto_rawDescData
}

var file_REDACTEDapi_api_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_REDACTEDapi_api_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_REDACTEDapi_api_namespace_proto_goTypes = []interface{}{
	(ComputedUserset_Object)(0),     // 0: ComputedUserset.Object
	(*NamespaceDefinition)(nil),     // 1: NamespaceDefinition
	(*Relation)(nil),                // 2: Relation
	(*UsersetRewrite)(nil),          // 3: UsersetRewrite
	(*SetOperation)(nil),            // 4: SetOperation
	(*TupleToUserset)(nil),          // 5: TupleToUserset
	(*ComputedUserset)(nil),         // 6: ComputedUserset
	(*SetOperation_Child)(nil),      // 7: SetOperation.Child
	(*SetOperation_Child_This)(nil), // 8: SetOperation.Child.This
	(*TupleToUserset_Tupleset)(nil), // 9: TupleToUserset.Tupleset
}
var file_REDACTEDapi_api_namespace_proto_depIdxs = []int32{
	2,  // 0: NamespaceDefinition.relation:type_name -> Relation
	3,  // 1: Relation.userset_rewrite:type_name -> UsersetRewrite
	4,  // 2: UsersetRewrite.union:type_name -> SetOperation
	4,  // 3: UsersetRewrite.intersection:type_name -> SetOperation
	4,  // 4: UsersetRewrite.exclusion:type_name -> SetOperation
	7,  // 5: SetOperation.child:type_name -> SetOperation.Child
	9,  // 6: TupleToUserset.tupleset:type_name -> TupleToUserset.Tupleset
	6,  // 7: TupleToUserset.computed_userset:type_name -> ComputedUserset
	0,  // 8: ComputedUserset.object:type_name -> ComputedUserset.Object
	8,  // 9: SetOperation.Child._this:type_name -> SetOperation.Child.This
	6,  // 10: SetOperation.Child.computed_userset:type_name -> ComputedUserset
	5,  // 11: SetOperation.Child.tuple_to_userset:type_name -> TupleToUserset
	3,  // 12: SetOperation.Child.userset_rewrite:type_name -> UsersetRewrite
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_REDACTEDapi_api_namespace_proto_init() }
func file_REDACTEDapi_api_namespace_proto_init() {
	if File_REDACTEDapi_api_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_REDACTEDapi_api_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceDefinition); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersetRewrite); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOperation); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TupleToUserset); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedUserset); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOperation_Child); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOperation_Child_This); i {
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
		file_REDACTEDapi_api_namespace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TupleToUserset_Tupleset); i {
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
	file_REDACTEDapi_api_namespace_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UsersetRewrite_Union)(nil),
		(*UsersetRewrite_Intersection)(nil),
		(*UsersetRewrite_Exclusion)(nil),
	}
	file_REDACTEDapi_api_namespace_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*SetOperation_Child_XThis)(nil),
		(*SetOperation_Child_ComputedUserset)(nil),
		(*SetOperation_Child_TupleToUserset)(nil),
		(*SetOperation_Child_UsersetRewrite)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_REDACTEDapi_api_namespace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_REDACTEDapi_api_namespace_proto_goTypes,
		DependencyIndexes: file_REDACTEDapi_api_namespace_proto_depIdxs,
		EnumInfos:         file_REDACTEDapi_api_namespace_proto_enumTypes,
		MessageInfos:      file_REDACTEDapi_api_namespace_proto_msgTypes,
	}.Build()
	File_REDACTEDapi_api_namespace_proto = out.File
	file_REDACTEDapi_api_namespace_proto_rawDesc = nil
	file_REDACTEDapi_api_namespace_proto_goTypes = nil
	file_REDACTEDapi_api_namespace_proto_depIdxs = nil
}
