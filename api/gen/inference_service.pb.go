// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: inference_service.proto

package gen

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InferenceServiceIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *InferenceServiceIdentifier) Reset() {
	*x = InferenceServiceIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceServiceIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceServiceIdentifier) ProtoMessage() {}

func (x *InferenceServiceIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceServiceIdentifier.ProtoReflect.Descriptor instead.
func (*InferenceServiceIdentifier) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{0}
}

func (x *InferenceServiceIdentifier) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *InferenceServiceIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Env struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Env) Reset() {
	*x = Env{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Env) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Env) ProtoMessage() {}

func (x *Env) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Env.ProtoReflect.Descriptor instead.
func (*Env) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{1}
}

func (x *Env) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Env) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Env   []*Env `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{2}
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetEnv() []*Env {
	if x != nil {
		return x.Env
	}
	return nil
}

type InferenceServiceTransformer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	MinCpu     string       `protobuf:"bytes,2,opt,name=minCpu,proto3" json:"minCpu,omitempty"`
	MinMemory  string       `protobuf:"bytes,3,opt,name=minMemory,proto3" json:"minMemory,omitempty"`
	MaxCpu     string       `protobuf:"bytes,4,opt,name=maxCpu,proto3" json:"maxCpu,omitempty"`
	MaxMemory  string       `protobuf:"bytes,5,opt,name=maxMemory,proto3" json:"maxMemory,omitempty"`
}

func (x *InferenceServiceTransformer) Reset() {
	*x = InferenceServiceTransformer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceServiceTransformer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceServiceTransformer) ProtoMessage() {}

func (x *InferenceServiceTransformer) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceServiceTransformer.ProtoReflect.Descriptor instead.
func (*InferenceServiceTransformer) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{3}
}

func (x *InferenceServiceTransformer) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *InferenceServiceTransformer) GetMinCpu() string {
	if x != nil {
		return x.MinCpu
	}
	return ""
}

func (x *InferenceServiceTransformer) GetMinMemory() string {
	if x != nil {
		return x.MinMemory
	}
	return ""
}

func (x *InferenceServiceTransformer) GetMaxCpu() string {
	if x != nil {
		return x.MaxCpu
	}
	return ""
}

func (x *InferenceServiceTransformer) GetMaxMemory() string {
	if x != nil {
		return x.MaxMemory
	}
	return ""
}

type InferenceServicePredictor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RuntimeVersion string `protobuf:"bytes,2,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	StorageUri     string `protobuf:"bytes,3,opt,name=storageUri,proto3" json:"storageUri,omitempty"`
	NodeSelector   string `protobuf:"bytes,4,opt,name=nodeSelector,proto3" json:"nodeSelector,omitempty"`
	MinCpu         string `protobuf:"bytes,5,opt,name=minCpu,proto3" json:"minCpu,omitempty"`
	MinMemory      string `protobuf:"bytes,6,opt,name=minMemory,proto3" json:"minMemory,omitempty"`
	MaxCpu         string `protobuf:"bytes,7,opt,name=maxCpu,proto3" json:"maxCpu,omitempty"`
	MaxMemory      string `protobuf:"bytes,8,opt,name=maxMemory,proto3" json:"maxMemory,omitempty"`
}

func (x *InferenceServicePredictor) Reset() {
	*x = InferenceServicePredictor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceServicePredictor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceServicePredictor) ProtoMessage() {}

func (x *InferenceServicePredictor) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceServicePredictor.ProtoReflect.Descriptor instead.
func (*InferenceServicePredictor) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{4}
}

func (x *InferenceServicePredictor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InferenceServicePredictor) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *InferenceServicePredictor) GetStorageUri() string {
	if x != nil {
		return x.StorageUri
	}
	return ""
}

func (x *InferenceServicePredictor) GetNodeSelector() string {
	if x != nil {
		return x.NodeSelector
	}
	return ""
}

func (x *InferenceServicePredictor) GetMinCpu() string {
	if x != nil {
		return x.MinCpu
	}
	return ""
}

func (x *InferenceServicePredictor) GetMinMemory() string {
	if x != nil {
		return x.MinMemory
	}
	return ""
}

func (x *InferenceServicePredictor) GetMaxCpu() string {
	if x != nil {
		return x.MaxCpu
	}
	return ""
}

func (x *InferenceServicePredictor) GetMaxMemory() string {
	if x != nil {
		return x.MaxMemory
	}
	return ""
}

type CreateInferenceServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace               string                       `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                    string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DefaultTransformerImage string                       `protobuf:"bytes,3,opt,name=defaultTransformerImage,proto3" json:"defaultTransformerImage,omitempty"`
	Predictor               *InferenceServicePredictor   `protobuf:"bytes,4,opt,name=predictor,proto3" json:"predictor,omitempty"`
	Transformer             *InferenceServiceTransformer `protobuf:"bytes,5,opt,name=transformer,proto3" json:"transformer,omitempty"`
}

func (x *CreateInferenceServiceRequest) Reset() {
	*x = CreateInferenceServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInferenceServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInferenceServiceRequest) ProtoMessage() {}

func (x *CreateInferenceServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInferenceServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateInferenceServiceRequest) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateInferenceServiceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateInferenceServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateInferenceServiceRequest) GetDefaultTransformerImage() string {
	if x != nil {
		return x.DefaultTransformerImage
	}
	return ""
}

func (x *CreateInferenceServiceRequest) GetPredictor() *InferenceServicePredictor {
	if x != nil {
		return x.Predictor
	}
	return nil
}

func (x *CreateInferenceServiceRequest) GetTransformer() *InferenceServiceTransformer {
	if x != nil {
		return x.Transformer
	}
	return nil
}

type DeployModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeployModelResponse) Reset() {
	*x = DeployModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployModelResponse) ProtoMessage() {}

func (x *DeployModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployModelResponse.ProtoReflect.Descriptor instead.
func (*DeployModelResponse) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeployModelResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InferenceServiceCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastTransitionTime string `protobuf:"bytes,1,opt,name=lastTransitionTime,proto3" json:"lastTransitionTime,omitempty"`
	Status             string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *InferenceServiceCondition) Reset() {
	*x = InferenceServiceCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceServiceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceServiceCondition) ProtoMessage() {}

func (x *InferenceServiceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceServiceCondition.ProtoReflect.Descriptor instead.
func (*InferenceServiceCondition) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{7}
}

func (x *InferenceServiceCondition) GetLastTransitionTime() string {
	if x != nil {
		return x.LastTransitionTime
	}
	return ""
}

func (x *InferenceServiceCondition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InferenceServiceCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetInferenceServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready      bool                         `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	Conditions []*InferenceServiceCondition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	PredictUrl string                       `protobuf:"bytes,3,opt,name=predictUrl,proto3" json:"predictUrl,omitempty"`
}

func (x *GetInferenceServiceResponse) Reset() {
	*x = GetInferenceServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInferenceServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInferenceServiceResponse) ProtoMessage() {}

func (x *GetInferenceServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInferenceServiceResponse.ProtoReflect.Descriptor instead.
func (*GetInferenceServiceResponse) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetInferenceServiceResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *GetInferenceServiceResponse) GetConditions() []*InferenceServiceCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *GetInferenceServiceResponse) GetPredictUrl() string {
	if x != nil {
		return x.PredictUrl
	}
	return ""
}

type InferenceServiceEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predict string `protobuf:"bytes,1,opt,name=predict,proto3" json:"predict,omitempty"`
}

func (x *InferenceServiceEndpoints) Reset() {
	*x = InferenceServiceEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inference_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceServiceEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceServiceEndpoints) ProtoMessage() {}

func (x *InferenceServiceEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_inference_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceServiceEndpoints.ProtoReflect.Descriptor instead.
func (*InferenceServiceEndpoints) Descriptor() ([]byte, []int) {
	return file_inference_service_proto_rawDescGZIP(), []int{9}
}

func (x *InferenceServiceEndpoints) GetPredict() string {
	if x != nil {
		return x.Predict
	}
	return ""
}

var File_inference_service_proto protoreflect.FileDescriptor

var file_inference_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x1a, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x03, 0x45, 0x6e, 0x76,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0xb9, 0x01,
	0x0a, 0x1b, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x69, 0x6e, 0x43, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x43, 0x70, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x87, 0x02, 0x0a, 0x19, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x55, 0x72, 0x69, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x43, 0x70,
	0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x43, 0x70, 0x75, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x61, 0x78, 0x43, 0x70, 0x75, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x22, 0x8d, 0x02, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x3c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x42, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x77, 0x0a, 0x19, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x22, 0x35, 0x0a, 0x19, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x32, 0xcf, 0x03, 0x0a, 0x10, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x22, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x7d, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x93, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x8c, 0x01, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x2a, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x65, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inference_service_proto_rawDescOnce sync.Once
	file_inference_service_proto_rawDescData = file_inference_service_proto_rawDesc
)

func file_inference_service_proto_rawDescGZIP() []byte {
	file_inference_service_proto_rawDescOnce.Do(func() {
		file_inference_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_inference_service_proto_rawDescData)
	})
	return file_inference_service_proto_rawDescData
}

var file_inference_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_inference_service_proto_goTypes = []interface{}{
	(*InferenceServiceIdentifier)(nil),    // 0: api.InferenceServiceIdentifier
	(*Env)(nil),                           // 1: api.Env
	(*Container)(nil),                     // 2: api.Container
	(*InferenceServiceTransformer)(nil),   // 3: api.InferenceServiceTransformer
	(*InferenceServicePredictor)(nil),     // 4: api.InferenceServicePredictor
	(*CreateInferenceServiceRequest)(nil), // 5: api.CreateInferenceServiceRequest
	(*DeployModelResponse)(nil),           // 6: api.DeployModelResponse
	(*InferenceServiceCondition)(nil),     // 7: api.InferenceServiceCondition
	(*GetInferenceServiceResponse)(nil),   // 8: api.GetInferenceServiceResponse
	(*InferenceServiceEndpoints)(nil),     // 9: api.InferenceServiceEndpoints
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_inference_service_proto_depIdxs = []int32{
	1,  // 0: api.Container.env:type_name -> api.Env
	2,  // 1: api.InferenceServiceTransformer.containers:type_name -> api.Container
	4,  // 2: api.CreateInferenceServiceRequest.predictor:type_name -> api.InferenceServicePredictor
	3,  // 3: api.CreateInferenceServiceRequest.transformer:type_name -> api.InferenceServiceTransformer
	7,  // 4: api.GetInferenceServiceResponse.conditions:type_name -> api.InferenceServiceCondition
	5,  // 5: api.InferenceService.CreateInferenceService:input_type -> api.CreateInferenceServiceRequest
	0,  // 6: api.InferenceService.GetInferenceService:input_type -> api.InferenceServiceIdentifier
	0,  // 7: api.InferenceService.DeleteInferenceService:input_type -> api.InferenceServiceIdentifier
	8,  // 8: api.InferenceService.CreateInferenceService:output_type -> api.GetInferenceServiceResponse
	8,  // 9: api.InferenceService.GetInferenceService:output_type -> api.GetInferenceServiceResponse
	10, // 10: api.InferenceService.DeleteInferenceService:output_type -> google.protobuf.Empty
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_inference_service_proto_init() }
func file_inference_service_proto_init() {
	if File_inference_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inference_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceServiceIdentifier); i {
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
		file_inference_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Env); i {
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
		file_inference_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_inference_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceServiceTransformer); i {
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
		file_inference_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceServicePredictor); i {
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
		file_inference_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInferenceServiceRequest); i {
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
		file_inference_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployModelResponse); i {
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
		file_inference_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceServiceCondition); i {
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
		file_inference_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInferenceServiceResponse); i {
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
		file_inference_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceServiceEndpoints); i {
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
			RawDescriptor: file_inference_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inference_service_proto_goTypes,
		DependencyIndexes: file_inference_service_proto_depIdxs,
		MessageInfos:      file_inference_service_proto_msgTypes,
	}.Build()
	File_inference_service_proto = out.File
	file_inference_service_proto_rawDesc = nil
	file_inference_service_proto_goTypes = nil
	file_inference_service_proto_depIdxs = nil
}
