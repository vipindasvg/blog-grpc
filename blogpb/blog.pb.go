// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRequest) Reset()         { *m = ReadBlogRequest{} }
func (m *ReadBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRequest) ProtoMessage()    {}
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{3}
}

func (m *ReadBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRequest.Unmarshal(m, b)
}
func (m *ReadBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRequest.Marshal(b, m, deterministic)
}
func (m *ReadBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRequest.Merge(m, src)
}
func (m *ReadBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRequest.Size(m)
}
func (m *ReadBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRequest proto.InternalMessageInfo

func (m *ReadBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogResponse) Reset()         { *m = ReadBlogResponse{} }
func (m *ReadBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBlogResponse) ProtoMessage()    {}
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{4}
}

func (m *ReadBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogResponse.Unmarshal(m, b)
}
func (m *ReadBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogResponse.Marshal(b, m, deterministic)
}
func (m *ReadBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogResponse.Merge(m, src)
}
func (m *ReadBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBlogResponse.Size(m)
}
func (m *ReadBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogResponse proto.InternalMessageInfo

func (m *ReadBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRequest) Reset()         { *m = UpdateBlogRequest{} }
func (m *UpdateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRequest) ProtoMessage()    {}
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{5}
}

func (m *UpdateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRequest.Unmarshal(m, b)
}
func (m *UpdateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRequest.Merge(m, src)
}
func (m *UpdateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRequest.Size(m)
}
func (m *UpdateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRequest proto.InternalMessageInfo

func (m *UpdateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogResponse) Reset()         { *m = UpdateBlogResponse{} }
func (m *UpdateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogResponse) ProtoMessage()    {}
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{6}
}

func (m *UpdateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogResponse.Unmarshal(m, b)
}
func (m *UpdateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogResponse.Marshal(b, m, deterministic)
}
func (m *UpdateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogResponse.Merge(m, src)
}
func (m *UpdateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogResponse.Size(m)
}
func (m *UpdateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogResponse proto.InternalMessageInfo

func (m *UpdateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRequest) Reset()         { *m = DeleteBlogRequest{} }
func (m *DeleteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRequest) ProtoMessage()    {}
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{7}
}

func (m *DeleteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRequest.Unmarshal(m, b)
}
func (m *DeleteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRequest.Merge(m, src)
}
func (m *DeleteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRequest.Size(m)
}
func (m *DeleteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRequest proto.InternalMessageInfo

func (m *DeleteBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type DeleteBlogResponse struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogResponse) Reset()         { *m = DeleteBlogResponse{} }
func (m *DeleteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogResponse) ProtoMessage()    {}
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{8}
}

func (m *DeleteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogResponse.Unmarshal(m, b)
}
func (m *DeleteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogResponse.Merge(m, src)
}
func (m *DeleteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogResponse.Size(m)
}
func (m *DeleteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogResponse proto.InternalMessageInfo

func (m *DeleteBlogResponse) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ListBlogRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRequest) Reset()         { *m = ListBlogRequest{} }
func (m *ListBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ListBlogRequest) ProtoMessage()    {}
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{9}
}

func (m *ListBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRequest.Unmarshal(m, b)
}
func (m *ListBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRequest.Marshal(b, m, deterministic)
}
func (m *ListBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRequest.Merge(m, src)
}
func (m *ListBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ListBlogRequest.Size(m)
}
func (m *ListBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRequest proto.InternalMessageInfo

type ListBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogResponse) Reset()         { *m = ListBlogResponse{} }
func (m *ListBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ListBlogResponse) ProtoMessage()    {}
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{10}
}

func (m *ListBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogResponse.Unmarshal(m, b)
}
func (m *ListBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogResponse.Marshal(b, m, deterministic)
}
func (m *ListBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogResponse.Merge(m, src)
}
func (m *ListBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ListBlogResponse.Size(m)
}
func (m *ListBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogResponse proto.InternalMessageInfo

func (m *ListBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
	proto.RegisterType((*ReadBlogRequest)(nil), "blog.ReadBlogRequest")
	proto.RegisterType((*ReadBlogResponse)(nil), "blog.ReadBlogResponse")
	proto.RegisterType((*UpdateBlogRequest)(nil), "blog.UpdateBlogRequest")
	proto.RegisterType((*UpdateBlogResponse)(nil), "blog.UpdateBlogResponse")
	proto.RegisterType((*DeleteBlogRequest)(nil), "blog.DeleteBlogRequest")
	proto.RegisterType((*DeleteBlogResponse)(nil), "blog.DeleteBlogResponse")
	proto.RegisterType((*ListBlogRequest)(nil), "blog.ListBlogRequest")
	proto.RegisterType((*ListBlogResponse)(nil), "blog.ListBlogResponse")
}

func init() { proto.RegisterFile("blogpb/blog.proto", fileDescriptor_1cd072c3eda6f7ba) }

var fileDescriptor_1cd072c3eda6f7ba = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0x0c, 0x15, 0xa1, 0x0c, 0x89, 0xd8, 0x8d, 0x4a, 0x83, 0x89, 0x31, 0x9c, 0x8c, 0x51, 0x34,
	0xe0, 0xc5, 0x78, 0x20, 0x41, 0x2f, 0x24, 0x9e, 0x30, 0x5e, 0xbc, 0x10, 0x60, 0x5f, 0xb0, 0x49,
	0xc3, 0xd6, 0x76, 0xf1, 0x1b, 0xfd, 0x2c, 0xb3, 0xbb, 0xac, 0xad, 0x2d, 0xc4, 0x7a, 0x6a, 0xdf,
	0xdb, 0x79, 0x6f, 0x66, 0x67, 0xb2, 0xf0, 0xe6, 0xa1, 0x58, 0x46, 0xf3, 0x1b, 0xf5, 0xe9, 0x45,
	0xb1, 0x90, 0x82, 0x55, 0xd5, 0x7f, 0x77, 0x81, 0xea, 0x28, 0x14, 0x4b, 0x76, 0x00, 0x27, 0xe0,
	0x7e, 0xe5, 0xbc, 0x72, 0xd1, 0x98, 0x38, 0x01, 0x67, 0xa7, 0x68, 0xcc, 0xd6, 0xf2, 0x5d, 0xc4,
	0xd3, 0x80, 0xfb, 0x8e, 0x6e, 0xbb, 0xa6, 0x31, 0xe6, 0xec, 0x08, 0xfb, 0x32, 0x90, 0x21, 0xf9,
	0x7b, 0xfa, 0xc0, 0x14, 0xcc, 0x47, 0x7d, 0x21, 0x56, 0x92, 0x56, 0xd2, 0xaf, 0xea, 0xbe, 0x2d,
	0xbb, 0x03, 0x78, 0x8f, 0x31, 0xcd, 0x24, 0x29, 0xaa, 0x09, 0x7d, 0xac, 0x29, 0x91, 0xec, 0x0c,
	0x5a, 0x81, 0xe6, 0x6c, 0xf6, 0xd1, 0xd3, 0xd2, 0x34, 0xc0, 0x28, 0xbb, 0x03, 0xcb, 0x0e, 0x25,
	0x91, 0x58, 0x25, 0xf4, 0xe7, 0xd4, 0x25, 0x5a, 0x13, 0x9a, 0xf1, 0x2c, 0x51, 0x1b, 0x75, 0x75,
	0x34, 0xfd, 0xb9, 0x5f, 0x4d, 0x95, 0x63, 0xde, 0xed, 0xe3, 0x30, 0xc5, 0x96, 0xdc, 0x3f, 0x80,
	0xf7, 0x1a, 0xf1, 0xff, 0x5f, 0x25, 0x3b, 0x54, 0x92, 0xea, 0x0a, 0xde, 0x13, 0x85, 0xf4, 0x9b,
	0x6a, 0xe7, 0x65, 0xae, 0xc1, 0xb2, 0xe8, 0x0d, 0xc7, 0x4e, 0xb8, 0x87, 0xd6, 0x73, 0x90, 0xc8,
	0xcc, 0x6a, 0x65, 0x47, 0xda, 0x2a, 0xa7, 0xb1, 0xff, 0xe5, 0xa0, 0xa9, 0xca, 0x17, 0x8a, 0x3f,
	0x83, 0x05, 0xb1, 0x21, 0x90, 0x86, 0xc6, 0xda, 0x06, 0x5f, 0xc8, 0xbe, 0xe3, 0x17, 0x0f, 0x36,
	0x84, 0xf7, 0x70, 0x6d, 0x26, 0xec, 0xd8, 0xa0, 0x72, 0x79, 0x76, 0x4e, 0xf2, 0xed, 0xcd, 0xe8,
	0x10, 0x48, 0x5d, 0xb6, 0xdc, 0x85, 0xb0, 0x2c, 0xf7, 0x96, 0x40, 0x86, 0x40, 0x6a, 0xa1, 0x5d,
	0x50, 0x88, 0xc0, 0x2e, 0xd8, 0xe2, 0xf6, 0x03, 0x5c, 0xeb, 0xa0, 0x15, 0x9f, 0x33, 0xd9, 0x8a,
	0xcf, 0x1b, 0x7d, 0x5b, 0x19, 0xb9, 0x6f, 0x35, 0xf3, 0x48, 0xe7, 0x35, 0xfd, 0x40, 0x07, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x8d, 0x96, 0xcd, 0xb5, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error) {
	out := new(UpdateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error) {
	out := new(DeleteBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogClient interface {
	Recv() (*ListBlogResponse, error)
	grpc.ClientStream
}

type blogServiceListBlogClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogClient) Recv() (*ListBlogResponse, error) {
	m := new(ListBlogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogResponse, error)
	ListBlog(*ListBlogRequest, BlogService_ListBlogServer) error
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlogRequest) (*ReadBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogRequest) (*UpdateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) DeleteBlog(ctx context.Context, req *DeleteBlogRequest) (*DeleteBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlog(req *ListBlogRequest, srv BlogService_ListBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlog(m, &blogServiceListBlogServer{stream})
}

type BlogService_ListBlogServer interface {
	Send(*ListBlogResponse) error
	grpc.ServerStream
}

type blogServiceListBlogServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogServer) Send(m *ListBlogResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _BlogService_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blogpb/blog.proto",
}
