// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: image.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Manifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag      string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	MimeType string                 `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Digest   string                 `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Created  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	Labels   map[string]string      `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Platform string                 `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	Size     int64                  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *Manifest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Manifest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Manifest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Manifest) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Manifest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Manifest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Manifest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ImageListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImageListRequest) Reset() {
	*x = ImageListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageListRequest) ProtoMessage() {}

func (x *ImageListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageListRequest.ProtoReflect.Descriptor instead.
func (*ImageListRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

type ImageListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*ImageListResponse_Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ImageListResponse) Reset() {
	*x = ImageListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageListResponse) ProtoMessage() {}

func (x *ImageListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageListResponse.ProtoReflect.Descriptor instead.
func (*ImageListResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2}
}

func (x *ImageListResponse) GetImages() []*ImageListResponse_Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type ImageInspectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ImageInspectRequest) Reset() {
	*x = ImageInspectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInspectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInspectRequest) ProtoMessage() {}

func (x *ImageInspectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInspectRequest.ProtoReflect.Descriptor instead.
func (*ImageInspectRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{3}
}

func (x *ImageInspectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ImageInspectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *ImageInspectResponse_Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ImageInspectResponse) Reset() {
	*x = ImageInspectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInspectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInspectResponse) ProtoMessage() {}

func (x *ImageInspectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInspectResponse.ProtoReflect.Descriptor instead.
func (*ImageInspectResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{4}
}

func (x *ImageInspectResponse) GetImage() *ImageInspectResponse_Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type ImageRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ImageRemoveRequest) Reset() {
	*x = ImageRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRemoveRequest) ProtoMessage() {}

func (x *ImageRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRemoveRequest.ProtoReflect.Descriptor instead.
func (*ImageRemoveRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{5}
}

func (x *ImageRemoveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ImageRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manifests []*Manifest `protobuf:"bytes,1,rep,name=manifests,proto3" json:"manifests,omitempty"`
}

func (x *ImageRemoveResponse) Reset() {
	*x = ImageRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRemoveResponse) ProtoMessage() {}

func (x *ImageRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRemoveResponse.ProtoReflect.Descriptor instead.
func (*ImageRemoveResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{6}
}

func (x *ImageRemoveResponse) GetManifests() []*Manifest {
	if x != nil {
		return x.Manifests
	}
	return nil
}

type ImageListResponse_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ManifestsCount int64     `protobuf:"varint,2,opt,name=manifestsCount,proto3" json:"manifestsCount,omitempty"`
	Latest         *Manifest `protobuf:"bytes,3,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (x *ImageListResponse_Image) Reset() {
	*x = ImageListResponse_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageListResponse_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageListResponse_Image) ProtoMessage() {}

func (x *ImageListResponse_Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageListResponse_Image.ProtoReflect.Descriptor instead.
func (*ImageListResponse_Image) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ImageListResponse_Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageListResponse_Image) GetManifestsCount() int64 {
	if x != nil {
		return x.ManifestsCount
	}
	return 0
}

func (x *ImageListResponse_Image) GetLatest() *Manifest {
	if x != nil {
		return x.Latest
	}
	return nil
}

type ImageInspectResponse_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manifests []*Manifest `protobuf:"bytes,2,rep,name=manifests,proto3" json:"manifests,omitempty"`
}

func (x *ImageInspectResponse_Image) Reset() {
	*x = ImageInspectResponse_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInspectResponse_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInspectResponse_Image) ProtoMessage() {}

func (x *ImageInspectResponse_Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInspectResponse_Image.ProtoReflect.Descriptor instead.
func (*ImageInspectResponse_Image) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ImageInspectResponse_Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageInspectResponse_Image) GetManifests() []*Manifest {
	if x != nil {
		return x.Manifests
	}
	return nil
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x12, 0x0a, 0x10, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb3, 0x01,
	0x0a, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x69, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x95,
	0x01, 0x0a, 0x14, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x47, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x09, 0x6d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x41, 0x0a, 0x13, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x73, 0x32, 0xd1, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x7a, 0x79, 0x2d, 0x6d, 0x61, 0x78, 0x2f,
	0x64, 0x69, 0x75, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_image_proto_goTypes = []interface{}{
	(*Manifest)(nil),                   // 0: pb.Manifest
	(*ImageListRequest)(nil),           // 1: pb.ImageListRequest
	(*ImageListResponse)(nil),          // 2: pb.ImageListResponse
	(*ImageInspectRequest)(nil),        // 3: pb.ImageInspectRequest
	(*ImageInspectResponse)(nil),       // 4: pb.ImageInspectResponse
	(*ImageRemoveRequest)(nil),         // 5: pb.ImageRemoveRequest
	(*ImageRemoveResponse)(nil),        // 6: pb.ImageRemoveResponse
	nil,                                // 7: pb.Manifest.LabelsEntry
	(*ImageListResponse_Image)(nil),    // 8: pb.ImageListResponse.Image
	(*ImageInspectResponse_Image)(nil), // 9: pb.ImageInspectResponse.Image
	(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
}
var file_image_proto_depIdxs = []int32{
	10, // 0: pb.Manifest.created:type_name -> google.protobuf.Timestamp
	7,  // 1: pb.Manifest.labels:type_name -> pb.Manifest.LabelsEntry
	8,  // 2: pb.ImageListResponse.images:type_name -> pb.ImageListResponse.Image
	9,  // 3: pb.ImageInspectResponse.image:type_name -> pb.ImageInspectResponse.Image
	0,  // 4: pb.ImageRemoveResponse.manifests:type_name -> pb.Manifest
	0,  // 5: pb.ImageListResponse.Image.latest:type_name -> pb.Manifest
	0,  // 6: pb.ImageInspectResponse.Image.manifests:type_name -> pb.Manifest
	1,  // 7: pb.ImageService.ImageList:input_type -> pb.ImageListRequest
	3,  // 8: pb.ImageService.ImageInspect:input_type -> pb.ImageInspectRequest
	5,  // 9: pb.ImageService.ImageRemove:input_type -> pb.ImageRemoveRequest
	2,  // 10: pb.ImageService.ImageList:output_type -> pb.ImageListResponse
	4,  // 11: pb.ImageService.ImageInspect:output_type -> pb.ImageInspectResponse
	6,  // 12: pb.ImageService.ImageRemove:output_type -> pb.ImageRemoveResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manifest); i {
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
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageListRequest); i {
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
		file_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageListResponse); i {
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
		file_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInspectRequest); i {
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
		file_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInspectResponse); i {
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
		file_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRemoveRequest); i {
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
		file_image_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRemoveResponse); i {
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
		file_image_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageListResponse_Image); i {
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
		file_image_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInspectResponse_Image); i {
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
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
