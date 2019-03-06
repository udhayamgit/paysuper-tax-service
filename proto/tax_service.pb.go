// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tax_service.proto

package tax_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{0}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (dst *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(dst, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type VATRate struct {
	// @inject_tag: json:"zip"
	ZipCode string `protobuf:"bytes,1,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	// @inject_tag: json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: gorm:"type:varchar(2)"
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// @inject_tag: json:"rate"
	Rate                 float32  `protobuf:"fixed32,6,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VATRate) Reset()         { *m = VATRate{} }
func (m *VATRate) String() string { return proto.CompactTextString(m) }
func (*VATRate) ProtoMessage()    {}
func (*VATRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{1}
}
func (m *VATRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VATRate.Unmarshal(m, b)
}
func (m *VATRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VATRate.Marshal(b, m, deterministic)
}
func (dst *VATRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VATRate.Merge(dst, src)
}
func (m *VATRate) XXX_Size() int {
	return xxx_messageInfo_VATRate.Size(m)
}
func (m *VATRate) XXX_DiscardUnknown() {
	xxx_messageInfo_VATRate.DiscardUnknown(m)
}

var xxx_messageInfo_VATRate proto.InternalMessageInfo

func (m *VATRate) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *VATRate) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *VATRate) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *VATRate) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *VATRate) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

type GetRatesQuery struct {
	// @inject_tag: json:"zip"
	ZipCode string `protobuf:"bytes,1,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	// @inject_tag: json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: json:"limit,omitempty"
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// @inject_tag: json:"offset,omitempty"
	Offset               int32    `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRatesQuery) Reset()         { *m = GetRatesQuery{} }
func (m *GetRatesQuery) String() string { return proto.CompactTextString(m) }
func (*GetRatesQuery) ProtoMessage()    {}
func (*GetRatesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{2}
}
func (m *GetRatesQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatesQuery.Unmarshal(m, b)
}
func (m *GetRatesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatesQuery.Marshal(b, m, deterministic)
}
func (dst *GetRatesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatesQuery.Merge(dst, src)
}
func (m *GetRatesQuery) XXX_Size() int {
	return xxx_messageInfo_GetRatesQuery.Size(m)
}
func (m *GetRatesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatesQuery proto.InternalMessageInfo

func (m *GetRatesQuery) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *GetRatesQuery) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *GetRatesQuery) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *GetRatesQuery) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetRatesQuery) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type RateLookupQuery struct {
	// @inject_tag: json:"zip"
	ZipCode string `protobuf:"bytes,1,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	// @inject_tag: json:"country"
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// @inject_tag: json:"city,omitempty"
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	// @inject_tag: json:"state,omitempty"
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLookupQuery) Reset()         { *m = RateLookupQuery{} }
func (m *RateLookupQuery) String() string { return proto.CompactTextString(m) }
func (*RateLookupQuery) ProtoMessage()    {}
func (*RateLookupQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{3}
}
func (m *RateLookupQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLookupQuery.Unmarshal(m, b)
}
func (m *RateLookupQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLookupQuery.Marshal(b, m, deterministic)
}
func (dst *RateLookupQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLookupQuery.Merge(dst, src)
}
func (m *RateLookupQuery) XXX_Size() int {
	return xxx_messageInfo_RateLookupQuery.Size(m)
}
func (m *RateLookupQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLookupQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RateLookupQuery proto.InternalMessageInfo

func (m *RateLookupQuery) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *RateLookupQuery) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *RateLookupQuery) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RateLookupQuery) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type GetSingleRateResponse struct {
	Rate                 *VATRate `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	NotAccurate          bool     `protobuf:"varint,2,opt,name=not_accurate,json=notAccurate,proto3" json:"not_accurate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSingleRateResponse) Reset()         { *m = GetSingleRateResponse{} }
func (m *GetSingleRateResponse) String() string { return proto.CompactTextString(m) }
func (*GetSingleRateResponse) ProtoMessage()    {}
func (*GetSingleRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{4}
}
func (m *GetSingleRateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSingleRateResponse.Unmarshal(m, b)
}
func (m *GetSingleRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSingleRateResponse.Marshal(b, m, deterministic)
}
func (dst *GetSingleRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleRateResponse.Merge(dst, src)
}
func (m *GetSingleRateResponse) XXX_Size() int {
	return xxx_messageInfo_GetSingleRateResponse.Size(m)
}
func (m *GetSingleRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleRateResponse proto.InternalMessageInfo

func (m *GetSingleRateResponse) GetRate() *VATRate {
	if m != nil {
		return m.Rate
	}
	return nil
}

func (m *GetSingleRateResponse) GetNotAccurate() bool {
	if m != nil {
		return m.NotAccurate
	}
	return false
}

type GetRatesResponse struct {
	// @inject_tag: json:"rates"
	Rates                []*VATRate `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetRatesResponse) Reset()         { *m = GetRatesResponse{} }
func (m *GetRatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetRatesResponse) ProtoMessage()    {}
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_d074d3a9a64ce81a, []int{5}
}
func (m *GetRatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatesResponse.Unmarshal(m, b)
}
func (m *GetRatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatesResponse.Marshal(b, m, deterministic)
}
func (dst *GetRatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatesResponse.Merge(dst, src)
}
func (m *GetRatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetRatesResponse.Size(m)
}
func (m *GetRatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatesResponse proto.InternalMessageInfo

func (m *GetRatesResponse) GetRates() []*VATRate {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyResponse)(nil), "tax_service.EmptyResponse")
	proto.RegisterType((*VATRate)(nil), "tax_service.VATRate")
	proto.RegisterType((*GetRatesQuery)(nil), "tax_service.GetRatesQuery")
	proto.RegisterType((*RateLookupQuery)(nil), "tax_service.RateLookupQuery")
	proto.RegisterType((*GetSingleRateResponse)(nil), "tax_service.GetSingleRateResponse")
	proto.RegisterType((*GetRatesResponse)(nil), "tax_service.GetRatesResponse")
}

func init() { proto.RegisterFile("tax_service.proto", fileDescriptor_tax_service_d074d3a9a64ce81a) }

var fileDescriptor_tax_service_d074d3a9a64ce81a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x4f, 0xea, 0x50,
	0x10, 0xa5, 0x40, 0x29, 0x6f, 0x78, 0x84, 0xf7, 0x6e, 0xd0, 0xd4, 0x46, 0x13, 0xbc, 0x2b, 0xe2,
	0x82, 0x05, 0xae, 0x35, 0x21, 0x6a, 0x70, 0xe1, 0xc6, 0x96, 0xb8, 0x25, 0xb5, 0x0c, 0xa6, 0x11,
	0x7a, 0x9b, 0x76, 0x6a, 0x28, 0x71, 0xed, 0x1f, 0xf2, 0x0f, 0x9a, 0x7b, 0x5b, 0x84, 0x22, 0x86,
	0x0d, 0xbb, 0x3b, 0x67, 0x3e, 0xce, 0xe9, 0x99, 0x29, 0xfc, 0x27, 0x77, 0x31, 0x8e, 0x31, 0x7a,
	0xf3, 0x3d, 0xec, 0x85, 0x91, 0x20, 0xc1, 0x1a, 0x1b, 0x10, 0x6f, 0x41, 0xf3, 0x6e, 0x1e, 0x52,
	0x6a, 0x63, 0x1c, 0x8a, 0x20, 0x46, 0xfe, 0x0e, 0xc6, 0xd3, 0x60, 0x64, 0xbb, 0x84, 0xec, 0x04,
	0xea, 0x4b, 0x3f, 0x1c, 0x7b, 0x62, 0x82, 0xa6, 0xd6, 0xd1, 0xba, 0x7f, 0x6c, 0x63, 0xe9, 0x87,
	0x37, 0x62, 0x82, 0xcc, 0x04, 0xc3, 0x13, 0x49, 0x40, 0x51, 0x6a, 0x96, 0xb3, 0x4c, 0x1e, 0x32,
	0x06, 0x55, 0xcf, 0xa7, 0xd4, 0xac, 0x28, 0x58, 0xbd, 0x59, 0x1b, 0xf4, 0x98, 0x5c, 0x42, 0xb3,
	0xaa, 0xc0, 0x2c, 0x90, 0x95, 0x91, 0x04, 0x6b, 0x1d, 0xad, 0x5b, 0xb6, 0xd5, 0x9b, 0x7f, 0x68,
	0xd0, 0x1c, 0x22, 0x49, 0xfa, 0xf8, 0x31, 0xc1, 0x28, 0x3d, 0xa8, 0x88, 0x99, 0x3f, 0xf7, 0x49,
	0x89, 0xd0, 0xed, 0x2c, 0x60, 0xc7, 0x50, 0x13, 0xd3, 0x69, 0x8c, 0x64, 0xea, 0x0a, 0xce, 0x23,
	0x1e, 0x42, 0x4b, 0x8a, 0x78, 0x10, 0xe2, 0x35, 0x09, 0x0f, 0xaf, 0xe4, 0xa7, 0x1d, 0x7c, 0x02,
	0x47, 0x43, 0x24, 0xc7, 0x0f, 0x5e, 0x66, 0x28, 0xa9, 0x57, 0x1b, 0x61, 0xdd, 0xdc, 0x27, 0xc9,
	0xd9, 0xe8, 0xb7, 0x7b, 0x9b, 0x1b, 0xcd, 0x57, 0x95, 0xb9, 0xc7, 0xce, 0xe1, 0x6f, 0x20, 0x68,
	0xec, 0x7a, 0x5e, 0xa2, 0x3a, 0xa4, 0x96, 0xba, 0xdd, 0x08, 0x04, 0x0d, 0x72, 0x88, 0x5f, 0xc3,
	0xbf, 0x95, 0xbf, 0xdf, 0x04, 0x17, 0xa0, 0xcb, 0x5c, 0x6c, 0x6a, 0x9d, 0xca, 0xaf, 0x0c, 0x59,
	0x49, 0xff, 0xb3, 0x0c, 0x30, 0x72, 0x17, 0x4e, 0x96, 0x65, 0x57, 0x60, 0x38, 0xd9, 0x38, 0xb6,
	0xb3, 0xcd, 0xb2, 0x0a, 0x68, 0xf1, 0xd4, 0x4a, 0xec, 0x1e, 0xe0, 0x16, 0x67, 0x48, 0xea, 0x83,
	0xd9, 0x69, 0xa1, 0x76, 0xcb, 0xfe, 0x3d, 0x93, 0x1c, 0x75, 0x37, 0x6b, 0xf7, 0xf6, 0x0c, 0xe3,
	0x85, 0xec, 0x4e, 0xdf, 0x79, 0x89, 0x0d, 0xa1, 0xbe, 0x32, 0x8b, 0x59, 0xdb, 0x1d, 0xeb, 0x1b,
	0xb5, 0xce, 0x76, 0xe6, 0xd6, 0x83, 0x9e, 0x6b, 0xea, 0xcf, 0xbb, 0xfc, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0x33, 0xfe, 0x3d, 0x8e, 0x03, 0x00, 0x00,
}
