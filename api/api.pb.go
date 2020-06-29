// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: api.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetPumpByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPumpByIdRequest) Reset() {
	*x = GetPumpByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPumpByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPumpByIdRequest) ProtoMessage() {}

func (x *GetPumpByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPumpByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPumpByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPumpByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPumpByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pump *Pump `protobuf:"bytes,1,opt,name=pump,proto3" json:"pump,omitempty"`
}

func (x *GetPumpByIdResponse) Reset() {
	*x = GetPumpByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPumpByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPumpByIdResponse) ProtoMessage() {}

func (x *GetPumpByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPumpByIdResponse.ProtoReflect.Descriptor instead.
func (*GetPumpByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetPumpByIdResponse) GetPump() *Pump {
	if x != nil {
		return x.Pump
	}
	return nil
}

type ListPumpsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPumpsRequest) Reset() {
	*x = ListPumpsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPumpsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPumpsRequest) ProtoMessage() {}

func (x *ListPumpsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPumpsRequest.ProtoReflect.Descriptor instead.
func (*ListPumpsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type ListPumpsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pumps []*Pump `protobuf:"bytes,1,rep,name=pumps,proto3" json:"pumps,omitempty"`
}

func (x *ListPumpsResponse) Reset() {
	*x = ListPumpsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPumpsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPumpsResponse) ProtoMessage() {}

func (x *ListPumpsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPumpsResponse.ProtoReflect.Descriptor instead.
func (*ListPumpsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListPumpsResponse) GetPumps() []*Pump {
	if x != nil {
		return x.Pumps
	}
	return nil
}

type Pump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique device identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// device display name
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	// device manufacturers
	Manufacturers []string `protobuf:"bytes,3,rep,name=manufacturers,proto3" json:"manufacturers,omitempty"`
	// device model
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Pump) Reset() {
	*x = Pump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pump) ProtoMessage() {}

func (x *Pump) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pump.ProtoReflect.Descriptor instead.
func (*Pump) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Pump) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pump) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Pump) GetManufacturers() []string {
	if x != nil {
		return x.Manufacturers
	}
	return nil
}

func (x *Pump) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type GetCgmByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCgmByIdRequest) Reset() {
	*x = GetCgmByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCgmByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCgmByIdRequest) ProtoMessage() {}

func (x *GetCgmByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCgmByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCgmByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetCgmByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCgmByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cgm *Cgm `protobuf:"bytes,1,opt,name=cgm,proto3" json:"cgm,omitempty"`
}

func (x *GetCgmByIdResponse) Reset() {
	*x = GetCgmByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCgmByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCgmByIdResponse) ProtoMessage() {}

func (x *GetCgmByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCgmByIdResponse.ProtoReflect.Descriptor instead.
func (*GetCgmByIdResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetCgmByIdResponse) GetCgm() *Cgm {
	if x != nil {
		return x.Cgm
	}
	return nil
}

type ListCgmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCgmsRequest) Reset() {
	*x = ListCgmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCgmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCgmsRequest) ProtoMessage() {}

func (x *ListCgmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCgmsRequest.ProtoReflect.Descriptor instead.
func (*ListCgmsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

type ListCgmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cgms []*Cgm `protobuf:"bytes,1,rep,name=cgms,proto3" json:"cgms,omitempty"`
}

func (x *ListCgmsResponse) Reset() {
	*x = ListCgmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCgmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCgmsResponse) ProtoMessage() {}

func (x *ListCgmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCgmsResponse.ProtoReflect.Descriptor instead.
func (*ListCgmsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListCgmsResponse) GetCgms() []*Cgm {
	if x != nil {
		return x.Cgms
	}
	return nil
}

type Cgm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique device identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// device display name
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	// device manufacturers
	Manufacturers []string `protobuf:"bytes,3,rep,name=manufacturers,proto3" json:"manufacturers,omitempty"`
	// device model
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Cgm) Reset() {
	*x = Cgm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cgm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cgm) ProtoMessage() {}

func (x *Cgm) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cgm.ProtoReflect.Descriptor instead.
func (*Cgm) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *Cgm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cgm) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Cgm) GetManufacturers() []string {
	if x != nil {
		return x.Manufacturers
	}
	return nil
}

func (x *Cgm) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x70,
	0x75, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x75, 0x6d, 0x70, 0x52, 0x04, 0x70, 0x75, 0x6d, 0x70, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x75, 0x6d, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x6d, 0x70, 0x52, 0x05, 0x70,
	0x75, 0x6d, 0x70, 0x73, 0x22, 0x74, 0x0a, 0x04, 0x50, 0x75, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x67, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x67, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x67, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x67, 0x6d, 0x52, 0x03, 0x63, 0x67,
	0x6d, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x67, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x67, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x67, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x67, 0x6d,
	0x52, 0x04, 0x63, 0x67, 0x6d, 0x73, 0x22, 0x73, 0x0a, 0x03, 0x43, 0x67, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x32, 0xf3, 0x02, 0x0a, 0x07,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x6d, 0x70, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x6d, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x75, 0x6d, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x6d, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x6d, 0x70, 0x73,
	0x12, 0x5c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x67, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x67, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x67, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x67, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x67, 0x6d, 0x73, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x67, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x67, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x67, 0x6d,
	0x73, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_goTypes = []interface{}{
	(*GetPumpByIdRequest)(nil),  // 0: api.GetPumpByIdRequest
	(*GetPumpByIdResponse)(nil), // 1: api.GetPumpByIdResponse
	(*ListPumpsRequest)(nil),    // 2: api.ListPumpsRequest
	(*ListPumpsResponse)(nil),   // 3: api.ListPumpsResponse
	(*Pump)(nil),                // 4: api.Pump
	(*GetCgmByIdRequest)(nil),   // 5: api.GetCgmByIdRequest
	(*GetCgmByIdResponse)(nil),  // 6: api.GetCgmByIdResponse
	(*ListCgmsRequest)(nil),     // 7: api.ListCgmsRequest
	(*ListCgmsResponse)(nil),    // 8: api.ListCgmsResponse
	(*Cgm)(nil),                 // 9: api.Cgm
}
var file_api_proto_depIdxs = []int32{
	4, // 0: api.GetPumpByIdResponse.pump:type_name -> api.Pump
	4, // 1: api.ListPumpsResponse.pumps:type_name -> api.Pump
	9, // 2: api.GetCgmByIdResponse.cgm:type_name -> api.Cgm
	9, // 3: api.ListCgmsResponse.cgms:type_name -> api.Cgm
	0, // 4: api.Devices.GetPumpById:input_type -> api.GetPumpByIdRequest
	2, // 5: api.Devices.ListPumps:input_type -> api.ListPumpsRequest
	5, // 6: api.Devices.GetCgmById:input_type -> api.GetCgmByIdRequest
	7, // 7: api.Devices.ListCgms:input_type -> api.ListCgmsRequest
	1, // 8: api.Devices.GetPumpById:output_type -> api.GetPumpByIdResponse
	3, // 9: api.Devices.ListPumps:output_type -> api.ListPumpsResponse
	6, // 10: api.Devices.GetCgmById:output_type -> api.GetCgmByIdResponse
	8, // 11: api.Devices.ListCgms:output_type -> api.ListCgmsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPumpByIdRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPumpByIdResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPumpsRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPumpsResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pump); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCgmByIdRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCgmByIdResponse); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCgmsRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCgmsResponse); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cgm); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicesClient interface {
	// Get pump by id
	GetPumpById(ctx context.Context, in *GetPumpByIdRequest, opts ...grpc.CallOption) (*GetPumpByIdResponse, error)
	// List pumps
	ListPumps(ctx context.Context, in *ListPumpsRequest, opts ...grpc.CallOption) (*ListPumpsResponse, error)
	// Get cgm by id
	GetCgmById(ctx context.Context, in *GetCgmByIdRequest, opts ...grpc.CallOption) (*GetCgmByIdResponse, error)
	// List cgms
	ListCgms(ctx context.Context, in *ListCgmsRequest, opts ...grpc.CallOption) (*ListCgmsResponse, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) GetPumpById(ctx context.Context, in *GetPumpByIdRequest, opts ...grpc.CallOption) (*GetPumpByIdResponse, error) {
	out := new(GetPumpByIdResponse)
	err := c.cc.Invoke(ctx, "/api.Devices/GetPumpById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListPumps(ctx context.Context, in *ListPumpsRequest, opts ...grpc.CallOption) (*ListPumpsResponse, error) {
	out := new(ListPumpsResponse)
	err := c.cc.Invoke(ctx, "/api.Devices/ListPumps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetCgmById(ctx context.Context, in *GetCgmByIdRequest, opts ...grpc.CallOption) (*GetCgmByIdResponse, error) {
	out := new(GetCgmByIdResponse)
	err := c.cc.Invoke(ctx, "/api.Devices/GetCgmById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListCgms(ctx context.Context, in *ListCgmsRequest, opts ...grpc.CallOption) (*ListCgmsResponse, error) {
	out := new(ListCgmsResponse)
	err := c.cc.Invoke(ctx, "/api.Devices/ListCgms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
type DevicesServer interface {
	// Get pump by id
	GetPumpById(context.Context, *GetPumpByIdRequest) (*GetPumpByIdResponse, error)
	// List pumps
	ListPumps(context.Context, *ListPumpsRequest) (*ListPumpsResponse, error)
	// Get cgm by id
	GetCgmById(context.Context, *GetCgmByIdRequest) (*GetCgmByIdResponse, error)
	// List cgms
	ListCgms(context.Context, *ListCgmsRequest) (*ListCgmsResponse, error)
}

// UnimplementedDevicesServer can be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (*UnimplementedDevicesServer) GetPumpById(context.Context, *GetPumpByIdRequest) (*GetPumpByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPumpById not implemented")
}
func (*UnimplementedDevicesServer) ListPumps(context.Context, *ListPumpsRequest) (*ListPumpsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPumps not implemented")
}
func (*UnimplementedDevicesServer) GetCgmById(context.Context, *GetCgmByIdRequest) (*GetCgmByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCgmById not implemented")
}
func (*UnimplementedDevicesServer) ListCgms(context.Context, *ListCgmsRequest) (*ListCgmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCgms not implemented")
}

func RegisterDevicesServer(s *grpc.Server, srv DevicesServer) {
	s.RegisterService(&_Devices_serviceDesc, srv)
}

func _Devices_GetPumpById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPumpByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetPumpById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Devices/GetPumpById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetPumpById(ctx, req.(*GetPumpByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListPumps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPumpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListPumps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Devices/ListPumps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListPumps(ctx, req.(*ListPumpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetCgmById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCgmByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetCgmById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Devices/GetCgmById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetCgmById(ctx, req.(*GetCgmByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListCgms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCgmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListCgms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Devices/ListCgms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListCgms(ctx, req.(*ListCgmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPumpById",
			Handler:    _Devices_GetPumpById_Handler,
		},
		{
			MethodName: "ListPumps",
			Handler:    _Devices_ListPumps_Handler,
		},
		{
			MethodName: "GetCgmById",
			Handler:    _Devices_GetCgmById_Handler,
		},
		{
			MethodName: "ListCgms",
			Handler:    _Devices_ListCgms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
