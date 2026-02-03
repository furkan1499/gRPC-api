// Simple proto implementation for demo

package proto

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset()         { *x = GetUserRequest{} }
func (x *GetUserRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetUserRequest) ProtoMessage()    {}
func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	return nil
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetUserResponse) Reset()         { *x = GetUserResponse{} }
func (x *GetUserResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetUserResponse) ProtoMessage()    {}
func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserRequest) Reset()         { *x = CreateUserRequest{} }
func (x *CreateUserRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateUserRequest) ProtoMessage()    {}
func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateUserResponse) Reset()         { *x = CreateUserResponse{} }
func (x *CreateUserResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*CreateUserResponse) ProtoMessage()    {}
func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	return nil
}