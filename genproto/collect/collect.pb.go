// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: collect/collect.proto

package collect

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CollectPostsRequest struct {
	Start                bool     `protobuf:"varint,1,opt,name=start,proto3" json:"start"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectPostsRequest) Reset()         { *m = CollectPostsRequest{} }
func (m *CollectPostsRequest) String() string { return proto.CompactTextString(m) }
func (*CollectPostsRequest) ProtoMessage()    {}
func (*CollectPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a99e93162ef549a, []int{0}
}
func (m *CollectPostsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectPostsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectPostsRequest.Merge(m, src)
}
func (m *CollectPostsRequest) XXX_Size() int {
	return m.Size()
}
func (m *CollectPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectPostsRequest proto.InternalMessageInfo

func (m *CollectPostsRequest) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

type CollectPostsResponse struct {
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors"`
	Code                 int64    `protobuf:"varint,3,opt,name=code,proto3" json:"code"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectPostsResponse) Reset()         { *m = CollectPostsResponse{} }
func (m *CollectPostsResponse) String() string { return proto.CompactTextString(m) }
func (*CollectPostsResponse) ProtoMessage()    {}
func (*CollectPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a99e93162ef549a, []int{1}
}
func (m *CollectPostsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectPostsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectPostsResponse.Merge(m, src)
}
func (m *CollectPostsResponse) XXX_Size() int {
	return m.Size()
}
func (m *CollectPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectPostsResponse proto.InternalMessageInfo

func (m *CollectPostsResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *CollectPostsResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type CheckStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusRequest) Reset()         { *m = CheckStatusRequest{} }
func (m *CheckStatusRequest) String() string { return proto.CompactTextString(m) }
func (*CheckStatusRequest) ProtoMessage()    {}
func (*CheckStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a99e93162ef549a, []int{2}
}
func (m *CheckStatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckStatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusRequest.Merge(m, src)
}
func (m *CheckStatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *CheckStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusRequest proto.InternalMessageInfo

type CheckStatusResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors"`
	Code                 int64    `protobuf:"varint,3,opt,name=code,proto3" json:"code"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatusResponse) Reset()         { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()    {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a99e93162ef549a, []int{3}
}
func (m *CheckStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusResponse.Merge(m, src)
}
func (m *CheckStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *CheckStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusResponse proto.InternalMessageInfo

func (m *CheckStatusResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *CheckStatusResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *CheckStatusResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*CollectPostsRequest)(nil), "collect.CollectPostsRequest")
	proto.RegisterType((*CollectPostsResponse)(nil), "collect.CollectPostsResponse")
	proto.RegisterType((*CheckStatusRequest)(nil), "collect.CheckStatusRequest")
	proto.RegisterType((*CheckStatusResponse)(nil), "collect.CheckStatusResponse")
}

func init() { proto.RegisterFile("collect/collect.proto", fileDescriptor_3a99e93162ef549a) }

var fileDescriptor_3a99e93162ef549a = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xc9,
	0x49, 0x4d, 0x2e, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x08, 0x42, 0x3e, 0x37, 0x37, 0x3f, 0x0f, 0x22, 0xad, 0xa4, 0xcd, 0x25, 0xec, 0x0c, 0x11,
	0x0f, 0xc8, 0x2f, 0x2e, 0x29, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62,
	0x2d, 0x2e, 0x49, 0x2c, 0x2a, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x70, 0x94, 0x82,
	0xb8, 0x44, 0x50, 0x15, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xa9, 0x71, 0xb1, 0xa5, 0x16,
	0x15, 0xe5, 0x17, 0x15, 0x4b, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0xf1, 0xe9, 0xc1, 0xdc, 0xe0,
	0x0a, 0x12, 0x0e, 0x82, 0xca, 0x0a, 0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2b,
	0x30, 0x6a, 0x30, 0x07, 0x81, 0xd9, 0x4a, 0x22, 0x5c, 0x42, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0xc1,
	0x25, 0x89, 0x25, 0xa5, 0x30, 0xfb, 0x95, 0x32, 0xb9, 0x84, 0x51, 0x44, 0xa1, 0x16, 0x89, 0x71,
	0xb1, 0x15, 0x83, 0x45, 0xa0, 0xee, 0x82, 0xf2, 0x28, 0x71, 0x80, 0x93, 0xc0, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b,
	0x38, 0x68, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xc8, 0xbb, 0x94, 0x52, 0x01, 0x00,
	0x00,
}

func (m *CollectPostsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectPostsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectPostsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Start {
		i--
		if m.Start {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CollectPostsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectPostsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectPostsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintCollect(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Errors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollect(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *CheckStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckStatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *CheckStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintCollect(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Errors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollect(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollect(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollect(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CollectPostsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CollectPostsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovCollect(uint64(l))
		}
	}
	if m.Code != 0 {
		n += 1 + sovCollect(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckStatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckStatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	if len(m.Errors) > 0 {
		for _, e := range m.Errors {
			l = e.Size()
			n += 1 + l + sovCollect(uint64(l))
		}
	}
	if m.Code != 0 {
		n += 1 + sovCollect(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCollect(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollect(x uint64) (n int) {
	return sovCollect(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CollectPostsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollect
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
			return fmt.Errorf("proto: CollectPostsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectPostsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Start = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCollect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CollectPostsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollect
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
			return fmt.Errorf("proto: CollectPostsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectPostsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
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
				return ErrInvalidLengthCollect
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, &Error{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollect
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
			return fmt.Errorf("proto: CheckStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCollect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollect
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
			return fmt.Errorf("proto: CheckStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
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
				return ErrInvalidLengthCollect
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, &Error{})
			if err := m.Errors[len(m.Errors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCollect(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollect
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
					return 0, ErrIntOverflowCollect
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
					return 0, ErrIntOverflowCollect
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
				return 0, ErrInvalidLengthCollect
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollect
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollect
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollect        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollect          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollect = fmt.Errorf("proto: unexpected end of group")
)