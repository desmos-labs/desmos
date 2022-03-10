// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/coingecko/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// QueryCirculatingSupplyRequest is the request type for the Query/CirculatingSupply RPC method
type QueryCirculatingSupplyRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryCirculatingSupplyRequest) Reset()         { *m = QueryCirculatingSupplyRequest{} }
func (m *QueryCirculatingSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingSupplyRequest) ProtoMessage()    {}
func (*QueryCirculatingSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8905c55eddc1e7d, []int{0}
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

// QueryCirculatingSupplyRequest is the response type for the Query/CirculatingSupply RPC method
type QueryCirculatingSupplyResponse struct {
	CirculatingSupply types.Coin `protobuf:"bytes,1,opt,name=circulating_supply,json=circulatingSupply,proto3" json:"circulating_supply"`
}

func (m *QueryCirculatingSupplyResponse) Reset()         { *m = QueryCirculatingSupplyResponse{} }
func (m *QueryCirculatingSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingSupplyResponse) ProtoMessage()    {}
func (*QueryCirculatingSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8905c55eddc1e7d, []int{1}
}
func (m *QueryCirculatingSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingSupplyResponse.Merge(m, src)
}
func (m *QueryCirculatingSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingSupplyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryCirculatingSupplyRequest)(nil), "desmos.coingecko.v1.QueryCirculatingSupplyRequest")
	proto.RegisterType((*QueryCirculatingSupplyResponse)(nil), "desmos.coingecko.v1.QueryCirculatingSupplyResponse")
}

func init() { proto.RegisterFile("desmos/coingecko/v1/query.proto", fileDescriptor_d8905c55eddc1e7d) }

var fileDescriptor_d8905c55eddc1e7d = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4f, 0xdb, 0x40,
	0x14, 0xc7, 0xed, 0xaa, 0xa9, 0x5a, 0x77, 0x8a, 0x9b, 0xa1, 0x89, 0xda, 0x73, 0x95, 0xa9, 0x4b,
	0xee, 0x64, 0x67, 0xa8, 0xd4, 0xa5, 0x52, 0xb2, 0x56, 0x48, 0x84, 0x8d, 0x25, 0x3a, 0x3b, 0x87,
	0xb1, 0x70, 0xee, 0x39, 0xb9, 0xb3, 0x85, 0x41, 0x2c, 0x4c, 0x8c, 0x48, 0xfc, 0x03, 0xf9, 0x27,
	0x98, 0x59, 0x33, 0x46, 0x62, 0x61, 0x42, 0x28, 0x61, 0xe0, 0xcf, 0x40, 0xbe, 0x73, 0x48, 0x04,
	0x01, 0x89, 0xcd, 0x4f, 0x9f, 0xe7, 0xef, 0x8f, 0x67, 0x5b, 0xce, 0x80, 0x89, 0x21, 0x08, 0x12,
	0x40, 0xc4, 0x43, 0x16, 0x1c, 0x00, 0xc9, 0x5c, 0x32, 0x4a, 0xd9, 0x38, 0xc7, 0xc9, 0x18, 0x24,
	0xd8, 0xdf, 0xf4, 0x02, 0x7e, 0x5a, 0xc0, 0x99, 0xdb, 0xa8, 0x85, 0x10, 0x82, 0xe2, 0xa4, 0x78,
	0xd2, 0xab, 0x8d, 0x1f, 0x21, 0x40, 0x18, 0x33, 0x42, 0x93, 0x88, 0x50, 0xce, 0x41, 0x52, 0x19,
	0x01, 0x17, 0x25, 0xad, 0x97, 0x54, 0x4d, 0x7e, 0xba, 0x47, 0x28, 0xcf, 0x97, 0x28, 0x80, 0xc2,
	0xa3, 0xaf, 0x15, 0xf5, 0x50, 0x22, 0xa4, 0x27, 0xe2, 0x53, 0xc1, 0x48, 0xe6, 0xfa, 0x4c, 0x52,
	0x57, 0x85, 0xd5, 0xbc, 0xf9, 0xcf, 0xfa, 0xb9, 0x5d, 0xa4, 0xed, 0x46, 0xe3, 0x20, 0x8d, 0xa9,
	0x8c, 0x78, 0xb8, 0x93, 0x26, 0x49, 0x9c, 0xf7, 0xd8, 0x28, 0x65, 0x42, 0xda, 0x35, 0xab, 0x32,
	0x60, 0x1c, 0x86, 0xdf, 0xcd, 0x5f, 0xe6, 0xef, 0x2f, 0x3d, 0x3d, 0xfc, 0xfd, 0x7c, 0x36, 0x71,
	0x8c, 0x87, 0x89, 0x63, 0x34, 0x8f, 0x2c, 0xf4, 0x9a, 0x80, 0x48, 0x80, 0x0b, 0x66, 0x6f, 0x59,
	0x76, 0xb0, 0x82, 0x7d, 0xa1, 0xa8, 0x92, 0xfb, 0xea, 0xd5, 0x71, 0x99, 0xb6, 0xc8, 0x87, 0xcb,
	0x7c, 0xb8, 0x0b, 0x11, 0xef, 0x7c, 0x9c, 0xde, 0x3a, 0x46, 0xaf, 0x1a, 0x3c, 0xd7, 0x5d, 0x79,
	0x7b, 0x57, 0xa6, 0x55, 0x51, 0xe6, 0xf6, 0xa5, 0x69, 0x55, 0x5f, 0x24, 0xb0, 0x3d, 0xbc, 0xe1,
	0xf8, 0xf8, 0xcd, 0xbe, 0x8d, 0xf6, 0xbb, 0xde, 0xd1, 0x15, 0x9b, 0x7f, 0x4e, 0xaf, 0xef, 0x2f,
	0x3e, 0xb8, 0x36, 0x21, 0x9b, 0x7e, 0x87, 0xb5, 0x0a, 0x2d, 0xdd, 0x9e, 0x1c, 0xab, 0x33, 0x9e,
	0x74, 0xfe, 0x4f, 0xe7, 0xc8, 0x9c, 0xcd, 0x91, 0x79, 0x37, 0x47, 0xe6, 0xf9, 0x02, 0x19, 0xb3,
	0x05, 0x32, 0x6e, 0x16, 0xc8, 0xd8, 0xf5, 0xc2, 0x48, 0xee, 0xa7, 0x3e, 0x0e, 0x60, 0x58, 0x8a,
	0xb6, 0x62, 0xea, 0x8b, 0xa5, 0x41, 0xe6, 0x91, 0xc3, 0x35, 0x17, 0x99, 0x27, 0x4c, 0xf8, 0x9f,
	0xd4, 0x37, 0x6d, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xec, 0x42, 0xf3, 0x95, 0x02, 0x00,
	0x00,
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
	// CirculatingSupply queries the circulating supply (Total supply - community pool - vesting tokens) of the given
	// denom
	CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*QueryCirculatingSupplyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*QueryCirculatingSupplyResponse, error) {
	out := new(QueryCirculatingSupplyResponse)
	err := c.cc.Invoke(ctx, "/desmos.coingecko.v1.Query/CirculatingSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CirculatingSupply queries the circulating supply (Total supply - community pool - vesting tokens) of the given
	// denom
	CirculatingSupply(context.Context, *QueryCirculatingSupplyRequest) (*QueryCirculatingSupplyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CirculatingSupply(ctx context.Context, req *QueryCirculatingSupplyRequest) (*QueryCirculatingSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CirculatingSupply not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/desmos.coingecko.v1.Query/CirculatingSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CirculatingSupply(ctx, req.(*QueryCirculatingSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.coingecko.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CirculatingSupply",
			Handler:    _Query_CirculatingSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/coingecko/v1/query.proto",
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
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CirculatingSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	return n
}

func (m *QueryCirculatingSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CirculatingSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
func (m *QueryCirculatingSupplyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCirculatingSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CirculatingSupply", wireType)
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
			if err := m.CirculatingSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
