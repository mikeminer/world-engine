// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NamespacesRequest struct {
}

func (m *NamespacesRequest) Reset()         { *m = NamespacesRequest{} }
func (m *NamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*NamespacesRequest) ProtoMessage()    {}
func (*NamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961e0cabda3d9d6, []int{0}
}
func (m *NamespacesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamespacesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacesRequest.Merge(m, src)
}
func (m *NamespacesRequest) XXX_Size() int {
	return m.Size()
}
func (m *NamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacesRequest proto.InternalMessageInfo

type NamespacesResponse struct {
	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (m *NamespacesResponse) Reset()         { *m = NamespacesResponse{} }
func (m *NamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*NamespacesResponse) ProtoMessage()    {}
func (*NamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961e0cabda3d9d6, []int{1}
}
func (m *NamespacesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamespacesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacesResponse.Merge(m, src)
}
func (m *NamespacesResponse) XXX_Size() int {
	return m.Size()
}
func (m *NamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacesResponse proto.InternalMessageInfo

func (m *NamespacesResponse) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type AddressRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *AddressRequest) Reset()         { *m = AddressRequest{} }
func (m *AddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()    {}
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961e0cabda3d9d6, []int{2}
}
func (m *AddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequest.Merge(m, src)
}
func (m *AddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *AddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequest proto.InternalMessageInfo

func (m *AddressRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type AddressResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AddressResponse) Reset()         { *m = AddressResponse{} }
func (m *AddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()    {}
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8961e0cabda3d9d6, []int{3}
}
func (m *AddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResponse.Merge(m, src)
}
func (m *AddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *AddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResponse proto.InternalMessageInfo

func (m *AddressResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*NamespacesRequest)(nil), "router.v1.NamespacesRequest")
	proto.RegisterType((*NamespacesResponse)(nil), "router.v1.NamespacesResponse")
	proto.RegisterType((*AddressRequest)(nil), "router.v1.AddressRequest")
	proto.RegisterType((*AddressResponse)(nil), "router.v1.AddressResponse")
}

func init() { proto.RegisterFile("router/v1/query.proto", fileDescriptor_8961e0cabda3d9d6) }

var fileDescriptor_8961e0cabda3d9d6 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4b, 0x2b, 0x31,
	0x14, 0xed, 0xbc, 0x07, 0xaf, 0x34, 0xef, 0xf1, 0xc4, 0x58, 0xa1, 0x0e, 0x75, 0x28, 0x83, 0x8b,
	0xa2, 0x34, 0xb1, 0x55, 0x74, 0xad, 0x4b, 0x17, 0x42, 0xeb, 0xce, 0x8d, 0xa4, 0xd3, 0x4b, 0x3a,
	0xd0, 0x26, 0xd3, 0x24, 0x53, 0x5b, 0xc4, 0x8d, 0xbf, 0x40, 0xf0, 0x4f, 0xb9, 0x2c, 0xb8, 0x71,
	0x29, 0xad, 0xff, 0x43, 0x69, 0xe7, 0x53, 0xa8, 0xb8, 0xbb, 0x73, 0xce, 0x99, 0x73, 0xce, 0xbd,
	0x04, 0x6d, 0x2b, 0x19, 0x1a, 0x50, 0x74, 0xdc, 0xa4, 0xa3, 0x10, 0xd4, 0x94, 0x04, 0x4a, 0x1a,
	0x89, 0x4b, 0x11, 0x4c, 0xc6, 0x4d, 0xbb, 0xcc, 0x25, 0x97, 0x2b, 0x94, 0x2e, 0xa7, 0x48, 0x60,
	0x57, 0xb9, 0x94, 0x7c, 0x00, 0x94, 0x05, 0x3e, 0x65, 0x42, 0x48, 0xc3, 0x8c, 0x2f, 0x85, 0x8e,
	0xd9, 0x9c, 0xab, 0x36, 0xcc, 0x40, 0x04, 0xbb, 0x5b, 0x68, 0xf3, 0x92, 0x0d, 0x41, 0x07, 0xcc,
	0x03, 0xdd, 0x81, 0x51, 0x08, 0xda, 0xb8, 0x17, 0x08, 0xe7, 0x41, 0x1d, 0x48, 0xa1, 0x01, 0x1f,
	0x23, 0x24, 0x52, 0xb4, 0x62, 0xd5, 0x7e, 0xd7, 0xff, 0xb6, 0xca, 0x24, 0x6d, 0x45, 0xd2, 0x5f,
	0x3a, 0x39, 0x9d, 0x4b, 0xd0, 0xff, 0xb3, 0x5e, 0x4f, 0x81, 0x4e, 0xdc, 0x71, 0x15, 0x95, 0x52,
	0xbe, 0x62, 0xd5, 0xac, 0x7a, 0xa9, 0x93, 0x01, 0xee, 0x01, 0xda, 0x48, 0xf5, 0x71, 0x70, 0x05,
	0x15, 0x59, 0x04, 0xc5, 0xf2, 0xe4, 0xb3, 0xf5, 0x61, 0xa1, 0x7f, 0xed, 0xe5, 0x8d, 0xae, 0x40,
	0x8d, 0x7d, 0x0f, 0x70, 0x88, 0x50, 0xd6, 0x1c, 0x57, 0xd7, 0xb5, 0x4b, 0x7a, 0xd8, 0xbb, 0xdf,
	0xb0, 0x51, 0xaa, 0xbb, 0xff, 0xf0, 0xf2, 0xfe, 0xf4, 0x6b, 0x0f, 0xbb, 0xf4, 0x56, 0xaa, 0x41,
	0xef, 0x06, 0x04, 0xf7, 0x05, 0xd0, 0xec, 0x8c, 0xd9, 0x92, 0x78, 0x82, 0x8a, 0x71, 0x69, 0xbc,
	0x93, 0x73, 0xfd, 0xba, 0xb8, 0x6d, 0xaf, 0xa3, 0xe2, 0xb4, 0x93, 0x55, 0xda, 0x21, 0x26, 0x3f,
	0xa7, 0xd1, 0xbb, 0x74, 0xbe, 0x3f, 0x6f, 0x3f, 0xcf, 0x1d, 0x6b, 0x36, 0x77, 0xac, 0xb7, 0xb9,
	0x63, 0x3d, 0x2e, 0x9c, 0xc2, 0x6c, 0xe1, 0x14, 0x5e, 0x17, 0x4e, 0xe1, 0xfa, 0x94, 0xfb, 0xa6,
	0x1f, 0x76, 0x89, 0x27, 0x87, 0x94, 0x29, 0x1e, 0xea, 0xc6, 0x80, 0x75, 0x75, 0x64, 0xdf, 0x88,
	0xed, 0xbd, 0x3e, 0xf3, 0x05, 0x9d, 0x24, 0x31, 0x66, 0x1a, 0x80, 0xee, 0xfe, 0x59, 0xbd, 0x8c,
	0xa3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x02, 0xf9, 0x20, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// Namespaces returns all the namespace:address pairs.
	Namespaces(ctx context.Context, in *NamespacesRequest, opts ...grpc.CallOption) (*NamespacesResponse, error)
	// Address returns the address for a given namespace.
	Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Namespaces(ctx context.Context, in *NamespacesRequest, opts ...grpc.CallOption) (*NamespacesResponse, error) {
	out := new(NamespacesResponse)
	err := c.cc.Invoke(ctx, "/router.v1.QueryService/Namespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Address(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/router.v1.QueryService/Address", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// Namespaces returns all the namespace:address pairs.
	Namespaces(context.Context, *NamespacesRequest) (*NamespacesResponse, error)
	// Address returns the address for a given namespace.
	Address(context.Context, *AddressRequest) (*AddressResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Namespaces(ctx context.Context, req *NamespacesRequest) (*NamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Namespaces not implemented")
}
func (*UnimplementedQueryServiceServer) Address(ctx context.Context, req *AddressRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Address not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Namespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Namespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.v1.QueryService/Namespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Namespaces(ctx, req.(*NamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Address_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Address(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.v1.QueryService/Address",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Address(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "router.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Namespaces",
			Handler:    _QueryService_Namespaces_Handler,
		},
		{
			MethodName: "Address",
			Handler:    _QueryService_Address_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router/v1/query.proto",
}

func (m *NamespacesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespacesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamespacesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *NamespacesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespacesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamespacesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for iNdEx := len(m.Namespaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Namespaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NamespacesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *NamespacesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for _, e := range m.Namespaces {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *AddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamespacesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespacesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespacesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespacesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespacesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespacesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespaces = append(m.Namespaces, &Namespace{})
			if err := m.Namespaces[len(m.Namespaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
