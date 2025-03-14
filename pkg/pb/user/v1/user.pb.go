// user.proto
// This file defines the User API v1 for managing user operations
// such as create, get, update, and delete.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.21.12
// source: user.proto

package user_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Role defines the various roles a user can have in the system.
type Role int32

const (
	// Unknown or unspecified role.
	Role_UNKNOWN_UNSPECIFIED Role = 0
	// Regular user role.
	Role_USER Role = 1
	// Admin user role.
	Role_ADMIN Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "UNKNOWN_UNSPECIFIED",
		1: "USER",
		2: "ADMIN",
	}
	Role_value = map[string]int32{
		"UNKNOWN_UNSPECIFIED": 0,
		"USER":                1,
		"ADMIN":               2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

// User represents a user in the system.
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Email of the user.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// Role of the user.
	Role Role `protobuf:"varint,4,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
	// Timestamp when the user was created.
	Created *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	// Timestamp when the user info was last updated.
	Updated       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UNKNOWN_UNSPECIFIED
}

func (x *User) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *User) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

// UserCreate represents the data required to create a new user.
type UserCreate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the user to create.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email of the user to create.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Password of the user to create.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Password confirmation of the user to create.
	PasswordConfirm string `protobuf:"bytes,4,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	// Role of the user to create.
	Role          Role `protobuf:"varint,5,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreate) Reset() {
	*x = UserCreate{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreate) ProtoMessage() {}

func (x *UserCreate) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreate.ProtoReflect.Descriptor instead.
func (*UserCreate) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserCreate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserCreate) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreate) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

func (x *UserCreate) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UNKNOWN_UNSPECIFIED
}

// UserUpdate represents the data required to update an existing user.
type UserUpdate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to update.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// [optional] Name of the user to update.
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// [optional] Email of the user to update.
	Email *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// Role of the user to update.
	Role          Role `protobuf:"varint,4,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserUpdate) Reset() {
	*x = UserUpdate{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdate) ProtoMessage() {}

func (x *UserUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdate.ProtoReflect.Descriptor instead.
func (*UserUpdate) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserUpdate) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UserUpdate) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *UserUpdate) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_UNKNOWN_UNSPECIFIED
}

// CreateRequest represents the request to create a user.
type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *UserCreate            `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetUser() *UserCreate {
	if x != nil {
		return x.User
	}
	return nil
}

// CreateResponse represents the response after a user is created.
type CreateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the created user.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetRequest represents the request to get a user by ID.
type GetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to retrieve.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetResponse represents the response containing the user info.
type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// UpdateRequest represents the request to update a user.
type UpdateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *UserUpdate            `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetUser() *UserUpdate {
	if x != nil {
		return x.User
	}
	return nil
}

// DeleteRequest represents the request to delete a user by ID.
type DeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to delete.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetMeResponse represents the response containing the current user's information.
type GetMeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Full user information
	User          *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMeResponse) Reset() {
	*x = GetMeResponse{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeResponse) ProtoMessage() {}

func (x *GetMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeResponse.ProtoReflect.Descriptor instead.
func (*GetMeResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetMeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// UpdateMeRequest represents the data that authenticated users can update about themselves.
type UpdateMeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// [optional] Name of the user to update
	Name *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// [optional] Email of the user to update
	Email         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateMeRequest) Reset() {
	*x = UpdateMeRequest{}
	mi := &file_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMeRequest) ProtoMessage() {}

func (x *UpdateMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMeRequest.ProtoReflect.Descriptor instead.
func (*UpdateMeRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateMeRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateMeRequest) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

// ChangePasswordRequest represents the request to change a user's password.
type ChangePasswordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current password for verification
	CurrentPassword string `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	// New password to set
	NewPassword   string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	mi := &file_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *ChangePasswordRequest) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x08, 0x18, 0x80,
	0x02, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x08, 0x18, 0x80,
	0x02, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x12, 0x2b, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0xd3, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x72, 0x07, 0x10, 0x01, 0x18, 0x32, 0xd0,
	0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0xd0, 0x01, 0x01, 0x60,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x91, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x72, 0x07, 0x10, 0x01, 0x18, 0x32, 0xd0,
	0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0xd0, 0x01, 0x01, 0x60,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x7d, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x08, 0x18, 0x80, 0x02, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x08, 0x18, 0x80, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2a, 0x34, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x17, 0x0a, 0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x32, 0x9b, 0x05,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01,
	0x2a, 0x32, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x32, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a,
	0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x0e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01,
	0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0xa7, 0x01, 0x92, 0x41,
	0x64, 0x12, 0x21, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x22, 0x0e, 0x0a,
	0x0c, 0x57, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x57, 0x68, 0x69, 0x74, 0x65, 0x32, 0x05, 0x31,
	0x2e, 0x30, 0x2e, 0x30, 0x1a, 0x17, 0x7b, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x48, 0x4f, 0x53, 0x54,
	0x7d, 0x3a, 0x7b, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x7d, 0x2a, 0x02, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x38, 0x74, 0x68, 0x67, 0x65, 0x6e, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_user_proto_goTypes = []any{
	(Role)(0),                      // 0: user_v1.Role
	(*User)(nil),                   // 1: user_v1.User
	(*UserCreate)(nil),             // 2: user_v1.UserCreate
	(*UserUpdate)(nil),             // 3: user_v1.UserUpdate
	(*CreateRequest)(nil),          // 4: user_v1.CreateRequest
	(*CreateResponse)(nil),         // 5: user_v1.CreateResponse
	(*GetRequest)(nil),             // 6: user_v1.GetRequest
	(*GetResponse)(nil),            // 7: user_v1.GetResponse
	(*UpdateRequest)(nil),          // 8: user_v1.UpdateRequest
	(*DeleteRequest)(nil),          // 9: user_v1.DeleteRequest
	(*GetMeResponse)(nil),          // 10: user_v1.GetMeResponse
	(*UpdateMeRequest)(nil),        // 11: user_v1.UpdateMeRequest
	(*ChangePasswordRequest)(nil),  // 12: user_v1.ChangePasswordRequest
	(*timestamppb.Timestamp)(nil),  // 13: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 14: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 15: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: user_v1.User.role:type_name -> user_v1.Role
	13, // 1: user_v1.User.created:type_name -> google.protobuf.Timestamp
	13, // 2: user_v1.User.updated:type_name -> google.protobuf.Timestamp
	0,  // 3: user_v1.UserCreate.role:type_name -> user_v1.Role
	14, // 4: user_v1.UserUpdate.name:type_name -> google.protobuf.StringValue
	14, // 5: user_v1.UserUpdate.email:type_name -> google.protobuf.StringValue
	0,  // 6: user_v1.UserUpdate.role:type_name -> user_v1.Role
	2,  // 7: user_v1.CreateRequest.user:type_name -> user_v1.UserCreate
	1,  // 8: user_v1.GetResponse.user:type_name -> user_v1.User
	3,  // 9: user_v1.UpdateRequest.user:type_name -> user_v1.UserUpdate
	1,  // 10: user_v1.GetMeResponse.user:type_name -> user_v1.User
	14, // 11: user_v1.UpdateMeRequest.name:type_name -> google.protobuf.StringValue
	14, // 12: user_v1.UpdateMeRequest.email:type_name -> google.protobuf.StringValue
	4,  // 13: user_v1.UserV1.Create:input_type -> user_v1.CreateRequest
	6,  // 14: user_v1.UserV1.Get:input_type -> user_v1.GetRequest
	8,  // 15: user_v1.UserV1.Update:input_type -> user_v1.UpdateRequest
	9,  // 16: user_v1.UserV1.Delete:input_type -> user_v1.DeleteRequest
	15, // 17: user_v1.UserV1.GetMe:input_type -> google.protobuf.Empty
	11, // 18: user_v1.UserV1.UpdateMe:input_type -> user_v1.UpdateMeRequest
	15, // 19: user_v1.UserV1.DeleteMe:input_type -> google.protobuf.Empty
	12, // 20: user_v1.UserV1.ChangePassword:input_type -> user_v1.ChangePasswordRequest
	5,  // 21: user_v1.UserV1.Create:output_type -> user_v1.CreateResponse
	7,  // 22: user_v1.UserV1.Get:output_type -> user_v1.GetResponse
	15, // 23: user_v1.UserV1.Update:output_type -> google.protobuf.Empty
	15, // 24: user_v1.UserV1.Delete:output_type -> google.protobuf.Empty
	10, // 25: user_v1.UserV1.GetMe:output_type -> user_v1.GetMeResponse
	15, // 26: user_v1.UserV1.UpdateMe:output_type -> google.protobuf.Empty
	15, // 27: user_v1.UserV1.DeleteMe:output_type -> google.protobuf.Empty
	15, // 28: user_v1.UserV1.ChangePassword:output_type -> google.protobuf.Empty
	21, // [21:29] is the sub-list for method output_type
	13, // [13:21] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
