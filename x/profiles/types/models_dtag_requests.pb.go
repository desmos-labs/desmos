// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models_dtag_requests.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// DTagTransferRequest represent a DTag transfer request between two users
type DTagTransferRequest struct {
	DTagToTrade string `protobuf:"bytes,1,opt,name=dtag_to_trade,json=dtagToTrade,proto3" json:"dtag_to_trade,omitempty" yaml:"dtag_to_trade"`
	Sender      string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Receiver    string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty" yaml:"receiver"`
}

func (m *DTagTransferRequest) Reset()         { *m = DTagTransferRequest{} }
func (m *DTagTransferRequest) String() string { return proto.CompactTextString(m) }
func (*DTagTransferRequest) ProtoMessage()    {}
func (*DTagTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f2e5360e821c5e, []int{0}
}
func (m *DTagTransferRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagTransferRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagTransferRequest.Merge(m, src)
}
func (m *DTagTransferRequest) XXX_Size() int {
	return m.Size()
}
func (m *DTagTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DTagTransferRequest proto.InternalMessageInfo

func (m *DTagTransferRequest) GetDTagToTrade() string {
	if m != nil {
		return m.DTagToTrade
	}
	return ""
}

func (m *DTagTransferRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *DTagTransferRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

// DTagTransferRequests contains a list of DTagTransferRequest
type DTagTransferRequests struct {
	Requests []DTagTransferRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *DTagTransferRequests) Reset()         { *m = DTagTransferRequests{} }
func (m *DTagTransferRequests) String() string { return proto.CompactTextString(m) }
func (*DTagTransferRequests) ProtoMessage()    {}
func (*DTagTransferRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f2e5360e821c5e, []int{1}
}
func (m *DTagTransferRequests) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagTransferRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagTransferRequests.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagTransferRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagTransferRequests.Merge(m, src)
}
func (m *DTagTransferRequests) XXX_Size() int {
	return m.Size()
}
func (m *DTagTransferRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagTransferRequests.DiscardUnknown(m)
}

var xxx_messageInfo_DTagTransferRequests proto.InternalMessageInfo

func (m *DTagTransferRequests) GetRequests() []DTagTransferRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*DTagTransferRequest)(nil), "desmos.profiles.v1beta1.DTagTransferRequest")
	proto.RegisterType((*DTagTransferRequests)(nil), "desmos.profiles.v1beta1.DTagTransferRequests")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models_dtag_requests.proto", fileDescriptor_08f2e5360e821c5e)
}

var fileDescriptor_08f2e5360e821c5e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x6f, 0xda, 0x40,
	0x14, 0xc7, 0x7d, 0xa5, 0x42, 0xd4, 0x08, 0x55, 0x35, 0x48, 0x75, 0x19, 0x7c, 0xc8, 0x4b, 0xa9,
	0xd4, 0xfa, 0x04, 0xdd, 0x18, 0x51, 0x87, 0x4a, 0x95, 0x3a, 0x58, 0x4c, 0x59, 0xac, 0xb3, 0xfd,
	0x70, 0x2c, 0xd9, 0x9c, 0x73, 0x77, 0xa0, 0xf0, 0x2d, 0x32, 0x66, 0xe4, 0xe3, 0x30, 0x32, 0x66,
	0xb2, 0x22, 0xb3, 0x64, 0xe6, 0x13, 0x44, 0xf6, 0x19, 0x50, 0x22, 0xa2, 0x6c, 0xef, 0xdd, 0xff,
	0xf7, 0xff, 0xbf, 0x7b, 0x7a, 0xfa, 0x38, 0x04, 0x91, 0x32, 0x41, 0x32, 0xce, 0xe6, 0x71, 0x02,
	0x82, 0xac, 0x46, 0x3e, 0x48, 0x3a, 0x22, 0x29, 0x0b, 0x21, 0x11, 0x5e, 0x28, 0x69, 0xe4, 0x71,
	0xb8, 0x59, 0x82, 0x90, 0xc2, 0xc9, 0x38, 0x93, 0xcc, 0xf8, 0xaa, 0x3c, 0xce, 0xd1, 0xe3, 0xd4,
	0x9e, 0x7e, 0x2f, 0x62, 0x11, 0xab, 0x18, 0x52, 0x56, 0x0a, 0xef, 0x7f, 0x8b, 0x18, 0x8b, 0x12,
	0x20, 0x55, 0xe7, 0x2f, 0xe7, 0x84, 0x2e, 0xd6, 0xb5, 0x84, 0x5f, 0x4b, 0x32, 0x4e, 0x41, 0x48,
	0x9a, 0x66, 0x47, 0x6f, 0xc0, 0xca, 0x51, 0x9e, 0x0a, 0x55, 0x4d, 0x2d, 0x0d, 0xdf, 0xf9, 0x79,
	0xec, 0x07, 0x8a, 0xb4, 0xb7, 0x48, 0xef, 0xfe, 0x99, 0xd1, 0x68, 0xc6, 0xe9, 0x42, 0xcc, 0x81,
	0xbb, 0x6a, 0x1d, 0xe3, 0x9f, 0xde, 0xa9, 0xd6, 0x93, 0xcc, 0x93, 0x9c, 0x86, 0x60, 0xa2, 0x01,
	0x1a, 0x7e, 0x9a, 0x7e, 0x2f, 0x72, 0xdc, 0xae, 0x78, 0x36, 0x2b, 0x9f, 0x0f, 0x39, 0xee, 0xad,
	0x69, 0x9a, 0x4c, 0xec, 0x17, 0xb4, 0xed, 0xb6, 0xcb, 0xbe, 0x86, 0x8c, 0x1f, 0x7a, 0x53, 0xc0,
	0x22, 0x04, 0x6e, 0x7e, 0xa8, 0x52, 0xbe, 0x1c, 0x72, 0xdc, 0x51, 0x36, 0xf5, 0x6e, 0xbb, 0x35,
	0x60, 0x10, 0xbd, 0xc5, 0x21, 0x80, 0x78, 0x05, 0xdc, 0x6c, 0x54, 0x70, 0xf7, 0x90, 0xe3, 0xcf,
	0x0a, 0x3e, 0x2a, 0xb6, 0x7b, 0x82, 0x26, 0xad, 0xfb, 0x0d, 0x46, 0x4f, 0x1b, 0x8c, 0xec, 0x4c,
	0xef, 0x5d, 0xd8, 0x44, 0x18, 0xff, 0xcb, 0x48, 0x55, 0x9b, 0x68, 0xd0, 0x18, 0xb6, 0xc7, 0x3f,
	0x9d, 0x37, 0xae, 0xe4, 0x5c, 0x08, 0x98, 0x7e, 0xdc, 0xe6, 0x58, 0x73, 0x4f, 0x19, 0xe7, 0x89,
	0xd3, 0xbf, 0xdb, 0xc2, 0x42, 0xbb, 0xc2, 0x42, 0x8f, 0x85, 0x85, 0xee, 0xf6, 0x96, 0xb6, 0xdb,
	0x5b, 0xda, 0xc3, 0xde, 0xd2, 0xae, 0x9c, 0x28, 0x96, 0xd7, 0x4b, 0xdf, 0x09, 0x58, 0x4a, 0xd4,
	0xac, 0x5f, 0x09, 0xf5, 0x45, 0x5d, 0x93, 0xdb, 0xf3, 0x65, 0xe4, 0x3a, 0x03, 0xe1, 0x37, 0xab,
	0x6b, 0xfc, 0x7e, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xbb, 0xd9, 0x36, 0x73, 0x02, 0x00, 0x00,
}

func (this *DTagTransferRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DTagTransferRequest)
	if !ok {
		that2, ok := that.(DTagTransferRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DTagToTrade != that1.DTagToTrade {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	if this.Receiver != that1.Receiver {
		return false
	}
	return true
}
func (this *DTagTransferRequests) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DTagTransferRequests)
	if !ok {
		that2, ok := that.(DTagTransferRequests)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Requests) != len(that1.Requests) {
		return false
	}
	for i := range this.Requests {
		if !this.Requests[i].Equal(&that1.Requests[i]) {
			return false
		}
	}
	return true
}
func (m *DTagTransferRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagTransferRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagTransferRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DTagToTrade) > 0 {
		i -= len(m.DTagToTrade)
		copy(dAtA[i:], m.DTagToTrade)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.DTagToTrade)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DTagTransferRequests) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagTransferRequests) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagTransferRequests) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintModelsDtagRequests(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsDtagRequests(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsDtagRequests(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DTagTransferRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DTagToTrade)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	return n
}

func (m *DTagTransferRequests) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovModelsDtagRequests(uint64(l))
		}
	}
	return n
}

func sovModelsDtagRequests(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsDtagRequests(x uint64) (n int) {
	return sovModelsDtagRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DTagTransferRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsDtagRequests
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
			return fmt.Errorf("proto: DTagTransferRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagTransferRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTagToTrade", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTagToTrade = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsDtagRequests
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
func (m *DTagTransferRequests) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsDtagRequests
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
			return fmt.Errorf("proto: DTagTransferRequests: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagTransferRequests: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, DTagTransferRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsDtagRequests
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
func skipModelsDtagRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsDtagRequests
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
					return 0, ErrIntOverflowModelsDtagRequests
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
					return 0, ErrIntOverflowModelsDtagRequests
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
				return 0, ErrInvalidLengthModelsDtagRequests
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsDtagRequests
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsDtagRequests
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsDtagRequests        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsDtagRequests          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsDtagRequests = fmt.Errorf("proto: unexpected end of group")
)
