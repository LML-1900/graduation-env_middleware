// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: env_service.proto

package grpc_env_service

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

type DataType int32

const (
	DataType_ZERO DataType = 0
	DataType_DEM  DataType = 1
	DataType_PBF  DataType = 2
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "ZERO",
		1: "DEM",
		2: "PBF",
	}
	DataType_value = map[string]int32{
		"ZERO": 0,
		"DEM":  1,
		"PBF":  2,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_env_service_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_env_service_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{0}
}

type UpdateCraterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateCraterResponse) Reset() {
	*x = UpdateCraterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCraterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCraterResponse) ProtoMessage() {}

func (x *UpdateCraterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCraterResponse.ProtoReflect.Descriptor instead.
func (*UpdateCraterResponse) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCraterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Position) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type Area struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bottomleft *Position `protobuf:"bytes,1,opt,name=bottomleft,proto3" json:"bottomleft,omitempty"`
	Topright   *Position `protobuf:"bytes,2,opt,name=topright,proto3" json:"topright,omitempty"`
}

func (x *Area) Reset() {
	*x = Area{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Area) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Area) ProtoMessage() {}

func (x *Area) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Area.ProtoReflect.Descriptor instead.
func (*Area) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{2}
}

func (x *Area) GetBottomleft() *Position {
	if x != nil {
		return x.Bottomleft
	}
	return nil
}

func (x *Area) GetTopright() *Position {
	if x != nil {
		return x.Topright
	}
	return nil
}

type Crater struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width float64   `protobuf:"fixed64,1,opt,name=width,proto3" json:"width,omitempty"`
	Depth float64   `protobuf:"fixed64,2,opt,name=depth,proto3" json:"depth,omitempty"`
	Pos   *Position `protobuf:"bytes,3,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *Crater) Reset() {
	*x = Crater{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crater) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crater) ProtoMessage() {}

func (x *Crater) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crater.ProtoReflect.Descriptor instead.
func (*Crater) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{3}
}

func (x *Crater) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Crater) GetDepth() float64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Crater) GetPos() *Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

type GetStaticDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Area     *Area    `protobuf:"bytes,1,opt,name=area,proto3" json:"area,omitempty"`
	Level    int32    `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	DataType DataType `protobuf:"varint,4,opt,name=dataType,proto3,enum=env_data_service.DataType" json:"dataType,omitempty"`
}

func (x *GetStaticDataRequest) Reset() {
	*x = GetStaticDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticDataRequest) ProtoMessage() {}

func (x *GetStaticDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticDataRequest.ProtoReflect.Descriptor instead.
func (*GetStaticDataRequest) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetStaticDataRequest) GetArea() *Area {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *GetStaticDataRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetStaticDataRequest) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_ZERO
}

type GetStaticDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TileID  string `protobuf:"bytes,1,opt,name=tileID,proto3" json:"tileID,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetStaticDataResponse) Reset() {
	*x = GetStaticDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaticDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticDataResponse) ProtoMessage() {}

func (x *GetStaticDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticDataResponse.ProtoReflect.Descriptor instead.
func (*GetStaticDataResponse) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetStaticDataResponse) GetTileID() string {
	if x != nil {
		return x.TileID
	}
	return ""
}

func (x *GetStaticDataResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type StartStopPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *Position `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *Position `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *StartStopPoints) Reset() {
	*x = StartStopPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStopPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStopPoints) ProtoMessage() {}

func (x *StartStopPoints) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStopPoints.ProtoReflect.Descriptor instead.
func (*StartStopPoints) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{6}
}

func (x *StartStopPoints) GetStart() *Position {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *StartStopPoints) GetEnd() *Position {
	if x != nil {
		return x.End
	}
	return nil
}

type RoutePoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos []*Position `protobuf:"bytes,1,rep,name=pos,proto3" json:"pos,omitempty"`
}

func (x *RoutePoints) Reset() {
	*x = RoutePoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutePoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutePoints) ProtoMessage() {}

func (x *RoutePoints) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutePoints.ProtoReflect.Descriptor instead.
func (*RoutePoints) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{7}
}

func (x *RoutePoints) GetPos() []*Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

type Obstacle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos   *Position `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Cause string    `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
}

func (x *Obstacle) Reset() {
	*x = Obstacle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Obstacle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Obstacle) ProtoMessage() {}

func (x *Obstacle) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Obstacle.ProtoReflect.Descriptor instead.
func (*Obstacle) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{8}
}

func (x *Obstacle) GetPos() *Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Obstacle) GetCause() string {
	if x != nil {
		return x.Cause
	}
	return ""
}

type UpdateObstacleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateObstacleResponse) Reset() {
	*x = UpdateObstacleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_env_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateObstacleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateObstacleResponse) ProtoMessage() {}

func (x *UpdateObstacleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_env_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateObstacleResponse.ProtoReflect.Descriptor instead.
func (*UpdateObstacleResponse) Descriptor() ([]byte, []int) {
	return file_env_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateObstacleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_env_service_proto protoreflect.FileDescriptor

var file_env_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x61, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x7a, 0x0a,
	0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x6c,
	0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x6c, 0x65, 0x66,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x74, 0x6f, 0x70, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x62, 0x0a, 0x06, 0x43, 0x72, 0x61,
	0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12,
	0x2c, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x90, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6c, 0x65, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x0f, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x30,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x3b,
	0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a,
	0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x4e, 0x0a, 0x08, 0x4f,
	0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0x26, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x5a,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x50, 0x42, 0x46, 0x10, 0x02, 0x32, 0xfb, 0x02, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x64, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x2e, 0x65,
	0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x52, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x26, 0x2e, 0x65, 0x6e,
	0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x73, 0x74, 0x61,
	0x63, 0x6c, 0x65, 0x1a, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x73,
	0x74, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x6e, 0x76, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x6e, 0x76, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_env_service_proto_rawDescOnce sync.Once
	file_env_service_proto_rawDescData = file_env_service_proto_rawDesc
)

func file_env_service_proto_rawDescGZIP() []byte {
	file_env_service_proto_rawDescOnce.Do(func() {
		file_env_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_env_service_proto_rawDescData)
	})
	return file_env_service_proto_rawDescData
}

var file_env_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_env_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_env_service_proto_goTypes = []interface{}{
	(DataType)(0),                  // 0: env_data_service.DataType
	(*UpdateCraterResponse)(nil),   // 1: env_data_service.UpdateCraterResponse
	(*Position)(nil),               // 2: env_data_service.Position
	(*Area)(nil),                   // 3: env_data_service.Area
	(*Crater)(nil),                 // 4: env_data_service.Crater
	(*GetStaticDataRequest)(nil),   // 5: env_data_service.GetStaticDataRequest
	(*GetStaticDataResponse)(nil),  // 6: env_data_service.GetStaticDataResponse
	(*StartStopPoints)(nil),        // 7: env_data_service.StartStopPoints
	(*RoutePoints)(nil),            // 8: env_data_service.RoutePoints
	(*Obstacle)(nil),               // 9: env_data_service.Obstacle
	(*UpdateObstacleResponse)(nil), // 10: env_data_service.UpdateObstacleResponse
}
var file_env_service_proto_depIdxs = []int32{
	2,  // 0: env_data_service.Area.bottomleft:type_name -> env_data_service.Position
	2,  // 1: env_data_service.Area.topright:type_name -> env_data_service.Position
	2,  // 2: env_data_service.Crater.pos:type_name -> env_data_service.Position
	3,  // 3: env_data_service.GetStaticDataRequest.area:type_name -> env_data_service.Area
	0,  // 4: env_data_service.GetStaticDataRequest.dataType:type_name -> env_data_service.DataType
	2,  // 5: env_data_service.StartStopPoints.start:type_name -> env_data_service.Position
	2,  // 6: env_data_service.StartStopPoints.end:type_name -> env_data_service.Position
	2,  // 7: env_data_service.RoutePoints.pos:type_name -> env_data_service.Position
	2,  // 8: env_data_service.Obstacle.pos:type_name -> env_data_service.Position
	5,  // 9: env_data_service.EnvironmentData.GetStaticData:input_type -> env_data_service.GetStaticDataRequest
	4,  // 10: env_data_service.EnvironmentData.UpdateCrater:input_type -> env_data_service.Crater
	9,  // 11: env_data_service.EnvironmentData.UpdateObstacle:input_type -> env_data_service.Obstacle
	7,  // 12: env_data_service.EnvironmentData.GetRoutePoints:input_type -> env_data_service.StartStopPoints
	6,  // 13: env_data_service.EnvironmentData.GetStaticData:output_type -> env_data_service.GetStaticDataResponse
	1,  // 14: env_data_service.EnvironmentData.UpdateCrater:output_type -> env_data_service.UpdateCraterResponse
	10, // 15: env_data_service.EnvironmentData.UpdateObstacle:output_type -> env_data_service.UpdateObstacleResponse
	8,  // 16: env_data_service.EnvironmentData.GetRoutePoints:output_type -> env_data_service.RoutePoints
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_env_service_proto_init() }
func file_env_service_proto_init() {
	if File_env_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_env_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCraterResponse); i {
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
		file_env_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_env_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Area); i {
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
		file_env_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crater); i {
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
		file_env_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticDataRequest); i {
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
		file_env_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaticDataResponse); i {
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
		file_env_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStopPoints); i {
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
		file_env_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutePoints); i {
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
		file_env_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Obstacle); i {
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
		file_env_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateObstacleResponse); i {
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
			RawDescriptor: file_env_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_env_service_proto_goTypes,
		DependencyIndexes: file_env_service_proto_depIdxs,
		EnumInfos:         file_env_service_proto_enumTypes,
		MessageInfos:      file_env_service_proto_msgTypes,
	}.Build()
	File_env_service_proto = out.File
	file_env_service_proto_rawDesc = nil
	file_env_service_proto_goTypes = nil
	file_env_service_proto_depIdxs = nil
}
