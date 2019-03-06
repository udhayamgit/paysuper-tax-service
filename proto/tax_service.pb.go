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
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{0}
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

type TaxEntry struct {
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

func (m *TaxEntry) Reset()         { *m = TaxEntry{} }
func (m *TaxEntry) String() string { return proto.CompactTextString(m) }
func (*TaxEntry) ProtoMessage()    {}
func (*TaxEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{1}
}
func (m *TaxEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaxEntry.Unmarshal(m, b)
}
func (m *TaxEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaxEntry.Marshal(b, m, deterministic)
}
func (dst *TaxEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaxEntry.Merge(dst, src)
}
func (m *TaxEntry) XXX_Size() int {
	return xxx_messageInfo_TaxEntry.Size(m)
}
func (m *TaxEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TaxEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TaxEntry proto.InternalMessageInfo

func (m *TaxEntry) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *TaxEntry) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *TaxEntry) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *TaxEntry) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TaxEntry) GetRate() float32 {
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
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{2}
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
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{3}
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
	Rate                 *TaxEntry `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	NotAccurate          bool      `protobuf:"varint,2,opt,name=not_accurate,json=notAccurate,proto3" json:"not_accurate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetSingleRateResponse) Reset()         { *m = GetSingleRateResponse{} }
func (m *GetSingleRateResponse) String() string { return proto.CompactTextString(m) }
func (*GetSingleRateResponse) ProtoMessage()    {}
func (*GetSingleRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{4}
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

func (m *GetSingleRateResponse) GetRate() *TaxEntry {
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
	Rates                []*TaxEntry `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRatesResponse) Reset()         { *m = GetRatesResponse{} }
func (m *GetRatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetRatesResponse) ProtoMessage()    {}
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tax_service_b919e59ef8b42c44, []int{5}
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

func (m *GetRatesResponse) GetRates() []*TaxEntry {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyResponse)(nil), "tax_service.EmptyResponse")
	proto.RegisterType((*TaxEntry)(nil), "tax_service.TaxEntry")
	proto.RegisterType((*GetRatesQuery)(nil), "tax_service.GetRatesQuery")
	proto.RegisterType((*RateLookupQuery)(nil), "tax_service.RateLookupQuery")
	proto.RegisterType((*GetSingleRateResponse)(nil), "tax_service.GetSingleRateResponse")
	proto.RegisterType((*GetRatesResponse)(nil), "tax_service.GetRatesResponse")
}

func init() { proto.RegisterFile("tax_service.proto", fileDescriptor_tax_service_b919e59ef8b42c44) }

var fileDescriptor_tax_service_b919e59ef8b42c44 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x4f, 0xea, 0x50,
	0x10, 0xa5, 0x40, 0xa1, 0x6f, 0x78, 0x84, 0xf7, 0x6e, 0xc4, 0xd4, 0x46, 0x13, 0xbc, 0x2b, 0x8c,
	0x09, 0x0b, 0xdc, 0x6b, 0x8c, 0x12, 0x5c, 0xb8, 0xb1, 0x65, 0x4f, 0x6a, 0x19, 0x4c, 0x23, 0xf4,
	0x36, 0xed, 0xd4, 0x00, 0x89, 0x5b, 0x7f, 0x91, 0x3f, 0xd0, 0xdc, 0xdb, 0x56, 0x28, 0x62, 0xd8,
	0xb0, 0xbb, 0x73, 0xce, 0x7c, 0x9c, 0x9e, 0x99, 0xc2, 0x7f, 0x72, 0x17, 0xe3, 0x18, 0xa3, 0x37,
	0xdf, 0xc3, 0x5e, 0x18, 0x09, 0x12, 0xac, 0xb1, 0x01, 0xf1, 0x16, 0x34, 0x07, 0xf3, 0x90, 0x96,
	0x36, 0xc6, 0xa1, 0x08, 0x62, 0xe4, 0xef, 0x60, 0x8c, 0xdc, 0xc5, 0x20, 0xa0, 0x68, 0xc9, 0x4e,
	0xc0, 0x58, 0xf9, 0xe1, 0xd8, 0x13, 0x13, 0x34, 0xb5, 0x8e, 0xd6, 0xfd, 0x63, 0xd7, 0x57, 0x7e,
	0x78, 0x27, 0x26, 0xc8, 0x4c, 0xa8, 0x7b, 0x22, 0x91, 0x59, 0x66, 0x39, 0x65, 0xb2, 0x90, 0x31,
	0xa8, 0x7a, 0x3e, 0x2d, 0xcd, 0x8a, 0x82, 0xd5, 0x9b, 0x1d, 0x81, 0x1e, 0x93, 0x4b, 0x68, 0x56,
	0x15, 0x98, 0x06, 0x32, 0x33, 0x92, 0x60, 0xad, 0xa3, 0x75, 0xcb, 0xb6, 0x7a, 0xf3, 0x0f, 0x0d,
	0x9a, 0x43, 0x24, 0xdb, 0x25, 0x8c, 0x9f, 0x12, 0x3c, 0xb0, 0x88, 0x99, 0x3f, 0xf7, 0x49, 0x89,
	0xd0, 0xed, 0x34, 0x60, 0xc7, 0x50, 0x13, 0xd3, 0x69, 0x8c, 0x64, 0xea, 0x0a, 0xce, 0x22, 0x1e,
	0x42, 0x4b, 0x8a, 0x78, 0x14, 0xe2, 0x35, 0x09, 0x0f, 0xaf, 0xe4, 0xa7, 0x1d, 0x1c, 0xa1, 0x3d,
	0x44, 0x72, 0xfc, 0xe0, 0x65, 0x86, 0x72, 0x74, 0xbe, 0x12, 0x76, 0x91, 0xf9, 0x24, 0x67, 0x36,
	0xfa, 0xed, 0xde, 0xe6, 0x4a, 0xf3, 0x5d, 0xa5, 0xf6, 0xb1, 0x73, 0xf8, 0x1b, 0x08, 0x1a, 0xbb,
	0x9e, 0x97, 0xa8, 0x12, 0x29, 0xc6, 0xb0, 0x1b, 0x81, 0xa0, 0xdb, 0x0c, 0xe2, 0x37, 0xf0, 0x2f,
	0x37, 0xf8, 0x7b, 0xc2, 0x25, 0xe8, 0x92, 0x8b, 0x4d, 0xad, 0x53, 0xf9, 0x7d, 0x44, 0x9a, 0xd3,
	0xff, 0x2c, 0x03, 0x8c, 0xdc, 0x85, 0x93, 0xd2, 0xec, 0x1a, 0xea, 0x4e, 0xda, 0x8f, 0xed, 0xae,
	0xb3, 0xac, 0x02, 0x5c, 0x3c, 0xb7, 0x12, 0x7b, 0x00, 0xb8, 0xc7, 0x19, 0x92, 0xfa, 0x66, 0x76,
	0x5a, 0xc8, 0xdd, 0xda, 0xc0, 0x9e, 0x4e, 0x8e, 0x3a, 0x9d, 0xb5, 0x81, 0x7b, 0x9a, 0xf1, 0x02,
	0xbb, 0xd3, 0x7a, 0x5e, 0x62, 0x43, 0x30, 0x72, 0xbb, 0x98, 0xb5, 0x5d, 0xb1, 0x3e, 0x53, 0xeb,
	0x6c, 0x27, 0xb7, 0x6e, 0xf4, 0x5c, 0x53, 0x7f, 0xdf, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8d, 0x46, 0x48, 0x7d, 0x92, 0x03, 0x00, 0x00,
}
