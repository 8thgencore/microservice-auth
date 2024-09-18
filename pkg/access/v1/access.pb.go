// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: access.proto

package access_v1

import (
	v1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endpoint where user wants access to.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type AddRoleEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endpoint to which roles will be added.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Roles allowed to access this endpoint.
	AllowedRoles []v1.Role `protobuf:"varint,2,rep,packed,name=allowed_roles,json=allowedRoles,proto3,enum=user_v1.Role" json:"allowed_roles,omitempty"`
}

func (x *AddRoleEndpointRequest) Reset() {
	*x = AddRoleEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleEndpointRequest) ProtoMessage() {}

func (x *AddRoleEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleEndpointRequest.ProtoReflect.Descriptor instead.
func (*AddRoleEndpointRequest) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{1}
}

func (x *AddRoleEndpointRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *AddRoleEndpointRequest) GetAllowedRoles() []v1.Role {
	if x != nil {
		return x.AllowedRoles
	}
	return nil
}

type UpdateRoleEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endpoint to be updated.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Updated roles for this endpoint.
	AllowedRoles []v1.Role `protobuf:"varint,2,rep,packed,name=allowed_roles,json=allowedRoles,proto3,enum=user_v1.Role" json:"allowed_roles,omitempty"`
}

func (x *UpdateRoleEndpointRequest) Reset() {
	*x = UpdateRoleEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleEndpointRequest) ProtoMessage() {}

func (x *UpdateRoleEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleEndpointRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleEndpointRequest) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRoleEndpointRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *UpdateRoleEndpointRequest) GetAllowedRoles() []v1.Role {
	if x != nil {
		return x.AllowedRoles
	}
	return nil
}

type DeleteRoleEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Endpoint to be deleted.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *DeleteRoleEndpointRequest) Reset() {
	*x = DeleteRoleEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleEndpointRequest) ProtoMessage() {}

func (x *DeleteRoleEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleEndpointRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleEndpointRequest) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRoleEndpointRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type GetRoleEndpointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of endpoint permissions.
	EndpointPermissions []*EndpointPermissions `protobuf:"bytes,1,rep,name=endpoint_permissions,json=endpointPermissions,proto3" json:"endpoint_permissions,omitempty"`
}

func (x *GetRoleEndpointsResponse) Reset() {
	*x = GetRoleEndpointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleEndpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleEndpointsResponse) ProtoMessage() {}

func (x *GetRoleEndpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleEndpointsResponse.ProtoReflect.Descriptor instead.
func (*GetRoleEndpointsResponse) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoleEndpointsResponse) GetEndpointPermissions() []*EndpointPermissions {
	if x != nil {
		return x.EndpointPermissions
	}
	return nil
}

// Represents endpoint permissions including the endpoint and the roles allowed to access it.
type EndpointPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint     string    `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AllowedRoles []v1.Role `protobuf:"varint,2,rep,packed,name=allowed_roles,json=allowedRoles,proto3,enum=user_v1.Role" json:"allowed_roles,omitempty"`
}

func (x *EndpointPermissions) Reset() {
	*x = EndpointPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointPermissions) ProtoMessage() {}

func (x *EndpointPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointPermissions.ProtoReflect.Descriptor instead.
func (*EndpointPermissions) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{5}
}

func (x *EndpointPermissions) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *EndpointPermissions) GetAllowedRoles() []v1.Role {
	if x != nil {
		return x.AllowedRoles
	}
	return nil
}

var File_access_proto protoreflect.FileDescriptor

var file_access_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b,
	0x72, 0x19, 0x10, 0x01, 0x18, 0xff, 0x01, 0x32, 0x12, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2f, 0x2e, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x72, 0x19, 0x10, 0x01, 0x18, 0xff, 0x01, 0x32, 0x12,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2f, 0x2e, 0x2d, 0x5d,
	0x2b, 0x24, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x72,
	0x19, 0x10, 0x01, 0x18, 0xff, 0x01, 0x32, 0x12, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x5f, 0x2f, 0x2e, 0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x22, 0x57, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x72, 0x19, 0x10, 0x01, 0x18, 0xff, 0x01, 0x32, 0x12, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2f, 0x2e, 0x2d, 0x5d, 0x2b,
	0x24, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x6d, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x14, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x72, 0x19, 0x10, 0x01, 0x18, 0xff, 0x01,
	0x32, 0x12, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2f, 0x2e,
	0x2d, 0x5d, 0x2b, 0x24, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3c,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0xc2, 0x04, 0x0a,
	0x08, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x56, 0x31, 0x12, 0x55, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x71, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x77, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a,
	0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x7f, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x2a, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d, 0x12, 0x72, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x38, 0x74, 0x68, 0x67, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_access_proto_rawDescOnce sync.Once
	file_access_proto_rawDescData = file_access_proto_rawDesc
)

func file_access_proto_rawDescGZIP() []byte {
	file_access_proto_rawDescOnce.Do(func() {
		file_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_access_proto_rawDescData)
	})
	return file_access_proto_rawDescData
}

var file_access_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_access_proto_goTypes = []any{
	(*CheckRequest)(nil),              // 0: access_v1.CheckRequest
	(*AddRoleEndpointRequest)(nil),    // 1: access_v1.AddRoleEndpointRequest
	(*UpdateRoleEndpointRequest)(nil), // 2: access_v1.UpdateRoleEndpointRequest
	(*DeleteRoleEndpointRequest)(nil), // 3: access_v1.DeleteRoleEndpointRequest
	(*GetRoleEndpointsResponse)(nil),  // 4: access_v1.GetRoleEndpointsResponse
	(*EndpointPermissions)(nil),       // 5: access_v1.EndpointPermissions
	(v1.Role)(0),                      // 6: user_v1.Role
	(*emptypb.Empty)(nil),             // 7: google.protobuf.Empty
}
var file_access_proto_depIdxs = []int32{
	6, // 0: access_v1.AddRoleEndpointRequest.allowed_roles:type_name -> user_v1.Role
	6, // 1: access_v1.UpdateRoleEndpointRequest.allowed_roles:type_name -> user_v1.Role
	5, // 2: access_v1.GetRoleEndpointsResponse.endpoint_permissions:type_name -> access_v1.EndpointPermissions
	6, // 3: access_v1.EndpointPermissions.allowed_roles:type_name -> user_v1.Role
	0, // 4: access_v1.AccessV1.Check:input_type -> access_v1.CheckRequest
	1, // 5: access_v1.AccessV1.AddRoleEndpoint:input_type -> access_v1.AddRoleEndpointRequest
	2, // 6: access_v1.AccessV1.UpdateRoleEndpoint:input_type -> access_v1.UpdateRoleEndpointRequest
	3, // 7: access_v1.AccessV1.DeleteRoleEndpoint:input_type -> access_v1.DeleteRoleEndpointRequest
	7, // 8: access_v1.AccessV1.GetRoleEndpoints:input_type -> google.protobuf.Empty
	7, // 9: access_v1.AccessV1.Check:output_type -> google.protobuf.Empty
	7, // 10: access_v1.AccessV1.AddRoleEndpoint:output_type -> google.protobuf.Empty
	7, // 11: access_v1.AccessV1.UpdateRoleEndpoint:output_type -> google.protobuf.Empty
	7, // 12: access_v1.AccessV1.DeleteRoleEndpoint:output_type -> google.protobuf.Empty
	4, // 13: access_v1.AccessV1.GetRoleEndpoints:output_type -> access_v1.GetRoleEndpointsResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_access_proto_init() }
func file_access_proto_init() {
	if File_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_access_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckRequest); i {
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
		file_access_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddRoleEndpointRequest); i {
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
		file_access_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRoleEndpointRequest); i {
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
		file_access_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRoleEndpointRequest); i {
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
		file_access_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoleEndpointsResponse); i {
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
		file_access_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EndpointPermissions); i {
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
			RawDescriptor: file_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_access_proto_goTypes,
		DependencyIndexes: file_access_proto_depIdxs,
		MessageInfos:      file_access_proto_msgTypes,
	}.Build()
	File_access_proto = out.File
	file_access_proto_rawDesc = nil
	file_access_proto_goTypes = nil
	file_access_proto_depIdxs = nil
}
