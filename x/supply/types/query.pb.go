// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/supply/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// QueryTotalSupplyRequest is the request type for Query/TotalSupply RPC method
type QueryTotalSupplyRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryTotalSupplyRequest) Reset()         { *m = QueryTotalSupplyRequest{} }
func (m *QueryTotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyRequest) ProtoMessage()    {}
func (*QueryTotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{0}
}
func (m *QueryTotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyRequest.Merge(m, src)
}
func (m *QueryTotalSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyRequest proto.InternalMessageInfo

// QueryCirculatingSupplyRequest is the request type for the Query/CirculatingSupply RPC method
type QueryCirculatingSupplyRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryCirculatingSupplyRequest) Reset()         { *m = QueryCirculatingSupplyRequest{} }
func (m *QueryCirculatingSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingSupplyRequest) ProtoMessage()    {}
func (*QueryCirculatingSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{1}
}
func (m *QueryCirculatingSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingSupplyRequest.Merge(m, src)
}
func (m *QueryCirculatingSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingSupplyRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryTotalSupplyRequest)(nil), "desmos.supply.v1.QueryTotalSupplyRequest")
	proto.RegisterType((*QueryCirculatingSupplyRequest)(nil), "desmos.supply.v1.QueryCirculatingSupplyRequest")
}

func init() { proto.RegisterFile("desmos/supply/v1/query.proto", fileDescriptor_821941a4adac8710) }

var fileDescriptor_821941a4adac8710 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0xce, 0x15, 0x7f, 0xa2, 0xe7, 0xe0, 0xcf, 0xa3, 0x60, 0x0d, 0x35, 0x2d, 0x1d, 0xa4, 0x05,
	0x7b, 0x47, 0xac, 0x93, 0xa3, 0xe2, 0xe0, 0x68, 0x2b, 0x0e, 0x2e, 0xe5, 0x92, 0x9c, 0xf1, 0x24,
	0xbd, 0xbb, 0xe6, 0x2e, 0xb1, 0x45, 0x5c, 0x1c, 0xc4, 0x51, 0x70, 0x71, 0xec, 0xee, 0x17, 0x71,
	0x2c, 0xb8, 0x38, 0x4a, 0xeb, 0xe0, 0xc7, 0x90, 0xdc, 0x25, 0x2a, 0x2d, 0xc5, 0xc9, 0x2d, 0xcf,
	0xfb, 0xe4, 0x7d, 0xfe, 0xbc, 0x09, 0xec, 0x26, 0x4c, 0x2f, 0xa4, 0x26, 0xba, 0x50, 0x2a, 0x5b,
	0x93, 0x32, 0x24, 0xcb, 0x82, 0xe5, 0x6b, 0xac, 0x72, 0x69, 0x24, 0x3a, 0x77, 0x2c, 0x76, 0x2c,
	0x2e, 0x43, 0xbf, 0x9d, 0xca, 0x54, 0x5a, 0x92, 0x54, 0x4f, 0xee, 0x3d, 0xbf, 0x9b, 0x4a, 0x99,
	0x66, 0x8c, 0x50, 0xc5, 0x09, 0x15, 0x42, 0x1a, 0x6a, 0xb8, 0x14, 0xba, 0x66, 0x6f, 0xd4, 0xac,
	0x45, 0x51, 0xf1, 0x9c, 0x50, 0x51, 0x1b, 0xf8, 0xc1, 0x21, 0xf5, 0x2a, 0xa7, 0x4a, 0xb1, 0xfc,
	0xf7, 0x6a, 0x2c, 0xab, 0x00, 0x73, 0xe7, 0xe8, 0x40, 0xb3, 0xea, 0x10, 0x89, 0xa8, 0x66, 0xa4,
	0x0c, 0x23, 0x66, 0x68, 0x48, 0x62, 0xc9, 0x85, 0xe3, 0x07, 0x09, 0xbc, 0xfe, 0xb8, 0xaa, 0xf2,
	0x44, 0x1a, 0x9a, 0xcd, 0x6c, 0x81, 0x29, 0x5b, 0x16, 0x4c, 0x1b, 0xd4, 0x86, 0x67, 0x09, 0x13,
	0x72, 0xd1, 0x01, 0x7d, 0x30, 0xbc, 0x3c, 0x75, 0x00, 0x8d, 0xe0, 0x79, 0xc2, 0x4b, 0x9e, 0xb0,
	0x7c, 0xce, 0x56, 0x4a, 0x0a, 0x26, 0x4c, 0xa7, 0xd5, 0x07, 0xc3, 0x0b, 0xd3, 0xab, 0xf5, 0xfc,
	0x61, 0x3d, 0xbe, 0x77, 0xe9, 0xfd, 0xa6, 0xe7, 0xfd, 0xdc, 0xf4, 0xbc, 0xc1, 0x4b, 0x78, 0xd3,
	0xba, 0x3c, 0xe0, 0x79, 0x5c, 0x64, 0xd4, 0x70, 0x91, 0xfe, 0x2f, 0xaf, 0x3b, 0x9f, 0x5b, 0xf0,
	0xcc, 0x9a, 0xa1, 0x77, 0x00, 0x5e, 0xf9, 0xab, 0x17, 0x1a, 0xe1, 0xc3, 0x0f, 0x85, 0x4f, 0x74,
	0xf7, 0xbb, 0xd8, 0x9d, 0x1c, 0x37, 0x27, 0xc7, 0x33, 0x93, 0x73, 0x91, 0x3e, 0xa5, 0x59, 0xc1,
	0x06, 0xf8, 0xed, 0xd7, 0x1f, 0x1f, 0x5b, 0x43, 0x74, 0x8b, 0x1c, 0xfd, 0x17, 0xa6, 0xd2, 0x1a,
	0xd7, 0xf8, 0xb5, 0xad, 0xf1, 0x06, 0x7d, 0x02, 0xf0, 0xda, 0x51, 0x75, 0x44, 0x4e, 0xc4, 0x39,
	0x75, 0xa4, 0x7f, 0x84, 0xba, 0x6b, 0x43, 0x61, 0x74, 0xfb, 0x38, 0x54, 0xfc, 0x47, 0xf1, 0x20,
	0xda, 0xfd, 0x47, 0x5f, 0x76, 0x01, 0xd8, 0xee, 0x02, 0xf0, 0x7d, 0x17, 0x80, 0x0f, 0xfb, 0xc0,
	0xdb, 0xee, 0x03, 0xef, 0xdb, 0x3e, 0xf0, 0x9e, 0x91, 0x94, 0x9b, 0x17, 0x45, 0x84, 0x63, 0xb9,
	0xa8, 0x15, 0xc7, 0x19, 0x8d, 0x74, 0xa3, 0x5e, 0x4e, 0xc8, 0xaa, 0xb1, 0x30, 0x6b, 0xc5, 0x74,
	0x74, 0xd1, 0xc6, 0x9a, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xba, 0xd5, 0x88, 0x5a, 0x2d, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// TotalSupply queries the total supply of the given denom
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	// CirculatingSupply queries the amount of tokens circulating in the market of the given denom
	CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/CirculatingSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// TotalSupply queries the total supply of the given denom
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*wrapperspb.StringValue, error)
	// CirculatingSupply queries the amount of tokens circulating in the market of the given denom
	CirculatingSupply(context.Context, *QueryCirculatingSupplyRequest) (*wrapperspb.StringValue, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *QueryTotalSupplyRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) CirculatingSupply(ctx context.Context, req *QueryCirculatingSupplyRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CirculatingSupply not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CirculatingSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCirculatingSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CirculatingSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/CirculatingSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CirculatingSupply(ctx, req.(*QueryCirculatingSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.supply.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "CirculatingSupply",
			Handler:    _Query_CirculatingSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/supply/v1/query.proto",
}

func (m *QueryTotalSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
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
func (m *QueryTotalSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func (m *QueryCirculatingSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTotalSupplyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryCirculatingSupplyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCirculatingSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
