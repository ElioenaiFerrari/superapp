// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: catalog.proto

package services

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

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *GetCatalogRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId int64  `protobuf:"varint,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *Catalog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Catalog) GetRestaurantId() int64 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

func (x *Catalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCatalogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId int64 `protobuf:"varint,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
}

func (x *ListCatalogsRequest) Reset() {
	*x = ListCatalogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsRequest) ProtoMessage() {}

func (x *ListCatalogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsRequest.ProtoReflect.Descriptor instead.
func (*ListCatalogsRequest) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *ListCatalogsRequest) GetRestaurantId() int64 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

type ListCatalogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalogs []*Catalog `protobuf:"bytes,1,rep,name=catalogs,proto3" json:"catalogs,omitempty"`
}

func (x *ListCatalogsResponse) Reset() {
	*x = ListCatalogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsResponse) ProtoMessage() {}

func (x *ListCatalogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsResponse.ProtoReflect.Descriptor instead.
func (*ListCatalogsResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *ListCatalogsResponse) GetCatalogs() []*Catalog {
	if x != nil {
		return x.Catalogs
	}
	return nil
}

type GetCatalogQRCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrCode []byte `protobuf:"bytes,1,opt,name=qr_code,json=qrCode,proto3" json:"qr_code,omitempty"`
}

func (x *GetCatalogQRCodeResponse) Reset() {
	*x = GetCatalogQRCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogQRCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogQRCodeResponse) ProtoMessage() {}

func (x *GetCatalogQRCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogQRCodeResponse.ProtoReflect.Descriptor instead.
func (*GetCatalogQRCodeResponse) Descriptor() ([]byte, []int) {
	return file_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *GetCatalogQRCodeResponse) GetQrCode() []byte {
	if x != nil {
		return x.QrCode
	}
	return nil
}

var File_catalog_proto protoreflect.FileDescriptor

var file_catalog_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52,
	0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xf8, 0x01, 0x0a, 0x0e, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x51, 0x52, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x69, 0x6f, 0x65, 0x6e, 0x61, 0x69, 0x46, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_proto_rawDescOnce sync.Once
	file_catalog_proto_rawDescData = file_catalog_proto_rawDesc
)

func file_catalog_proto_rawDescGZIP() []byte {
	file_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_proto_rawDescData)
	})
	return file_catalog_proto_rawDescData
}

var file_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_catalog_proto_goTypes = []interface{}{
	(*GetCatalogRequest)(nil),        // 0: services.GetCatalogRequest
	(*Catalog)(nil),                  // 1: services.Catalog
	(*ListCatalogsRequest)(nil),      // 2: services.ListCatalogsRequest
	(*ListCatalogsResponse)(nil),     // 3: services.ListCatalogsResponse
	(*GetCatalogQRCodeResponse)(nil), // 4: services.GetCatalogQRCodeResponse
}
var file_catalog_proto_depIdxs = []int32{
	1, // 0: services.ListCatalogsResponse.catalogs:type_name -> services.Catalog
	0, // 1: services.CatalogService.GetCatalog:input_type -> services.GetCatalogRequest
	2, // 2: services.CatalogService.ListCatalogs:input_type -> services.ListCatalogsRequest
	0, // 3: services.CatalogService.GetCatalogQRCode:input_type -> services.GetCatalogRequest
	1, // 4: services.CatalogService.GetCatalog:output_type -> services.Catalog
	3, // 5: services.CatalogService.ListCatalogs:output_type -> services.ListCatalogsResponse
	4, // 6: services.CatalogService.GetCatalogQRCode:output_type -> services.GetCatalogQRCodeResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_catalog_proto_init() }
func file_catalog_proto_init() {
	if File_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
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
		file_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
		file_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsRequest); i {
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
		file_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsResponse); i {
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
		file_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogQRCodeResponse); i {
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
			RawDescriptor: file_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_proto_depIdxs,
		MessageInfos:      file_catalog_proto_msgTypes,
	}.Build()
	File_catalog_proto = out.File
	file_catalog_proto_rawDesc = nil
	file_catalog_proto_goTypes = nil
	file_catalog_proto_depIdxs = nil
}
