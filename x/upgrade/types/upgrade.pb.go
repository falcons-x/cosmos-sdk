// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/upgrade/v1beta1/upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

// Plan specifies information about a planned upgrade and when it should occur.
type Plan struct {
	// Sets the name for the upgrade. This name will be used by the upgraded
	// version of the software to apply any special "on-upgrade" commands during
	// the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no
	// upgrade handler with this name has been set in the software, it will be
	// assumed that the software is out-of-date when the upgrade Time or Height is
	// reached and the software will exit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *Plan) Reset()      { *m = Plan{} }
func (*Plan) ProtoMessage() {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

// SoftwareUpgradeProposal is a gov Content type for initiating a software
// upgrade.
type SoftwareUpgradeProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Plan        Plan   `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan"`
}

func (m *SoftwareUpgradeProposal) Reset()      { *m = SoftwareUpgradeProposal{} }
func (*SoftwareUpgradeProposal) ProtoMessage() {}
func (*SoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{1}
}
func (m *SoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SoftwareUpgradeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoftwareUpgradeProposal.Merge(m, src)
}
func (m *SoftwareUpgradeProposal) XXX_Size() int {
	return m.Size()
}
func (m *SoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SoftwareUpgradeProposal proto.InternalMessageInfo

// CancelSoftwareUpgradeProposal is a gov Content type for cancelling a software
// upgrade.
type CancelSoftwareUpgradeProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *CancelSoftwareUpgradeProposal) Reset()      { *m = CancelSoftwareUpgradeProposal{} }
func (*CancelSoftwareUpgradeProposal) ProtoMessage() {}
func (*CancelSoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{2}
}
func (m *CancelSoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelSoftwareUpgradeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelSoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.Merge(m, src)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Size() int {
	return m.Size()
}
func (m *CancelSoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CancelSoftwareUpgradeProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Plan)(nil), "cosmos.upgrade.v1beta1.Plan")
	proto.RegisterType((*SoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal")
	proto.RegisterType((*CancelSoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal")
}

func init() {
	proto.RegisterFile("cosmos/upgrade/v1beta1/upgrade.proto", fileDescriptor_ccf2a7d4d7b48dca)
}

var fileDescriptor_ccf2a7d4d7b48dca = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0xc7, 0x3b, 0x30, 0x34, 0xdc, 0x61, 0x43, 0x1a, 0xc2, 0xed, 0x25, 0xd7, 0x81, 0x10, 0x17,
	0x2c, 0xb4, 0x0d, 0x9a, 0xb8, 0x70, 0x89, 0x3b, 0x16, 0x86, 0xd4, 0xb8, 0x31, 0x71, 0x31, 0x2d,
	0xd3, 0xd2, 0xd8, 0xf6, 0x34, 0xed, 0xa0, 0xf2, 0x14, 0xfa, 0x08, 0x3e, 0x0e, 0x4b, 0x96, 0xac,
	0x8c, 0xc0, 0xc6, 0xc7, 0x30, 0x9d, 0x29, 0xc6, 0x18, 0x97, 0xae, 0xe6, 0x7f, 0xce, 0xfc, 0xce,
	0xd7, 0xcc, 0x21, 0x87, 0x1e, 0xe4, 0x31, 0xe4, 0xf6, 0x3c, 0x0d, 0x32, 0x36, 0xe5, 0xf6, 0xfd,
	0xd0, 0xe5, 0x82, 0x0d, 0xf7, 0xb6, 0x95, 0x66, 0x20, 0xc0, 0x68, 0x2b, 0xca, 0xda, 0x7b, 0x4b,
	0xaa, 0xf3, 0x2f, 0x00, 0x08, 0x22, 0x6e, 0x4b, 0xca, 0x9d, 0xfb, 0x36, 0x4b, 0x16, 0x2a, 0xa4,
	0xd3, 0x0a, 0x20, 0x00, 0x29, 0xed, 0x42, 0x95, 0xde, 0xee, 0xf7, 0x00, 0x11, 0xc6, 0x3c, 0x17,
	0x2c, 0x4e, 0x15, 0xd0, 0xf7, 0x09, 0x9e, 0x44, 0x2c, 0x31, 0x0c, 0x82, 0x13, 0x16, 0x73, 0x13,
	0xf5, 0xd0, 0xe0, 0x8f, 0x23, 0xb5, 0xd1, 0x26, 0xfa, 0x8c, 0x87, 0xc1, 0x4c, 0x98, 0xd5, 0x1e,
	0x1a, 0x54, 0x9d, 0xd2, 0x2a, 0xd8, 0x30, 0xf1, 0xc1, 0xc4, 0x8a, 0x2d, 0xf4, 0x39, 0x7e, 0x7f,
	0xe9, 0xa2, 0x31, 0xae, 0x57, 0x9a, 0xd5, 0x31, 0xae, 0xd7, 0x9a, 0xba, 0x83, 0x8b, 0x52, 0x8e,
	0x0e, 0xa9, 0x08, 0x21, 0xe9, 0x3f, 0x21, 0xf2, 0xf7, 0x0a, 0x7c, 0xf1, 0xc0, 0x32, 0x7e, 0xad,
	0xa6, 0x9a, 0x64, 0x90, 0x42, 0xce, 0x22, 0xa3, 0x45, 0x6a, 0x22, 0x14, 0xd1, 0xbe, 0xb8, 0x32,
	0x8c, 0x1e, 0x69, 0x4c, 0x79, 0xee, 0x65, 0xa1, 0x4c, 0x60, 0x56, 0xe4, 0xdd, 0x57, 0x97, 0x71,
	0x46, 0x70, 0x1a, 0xb1, 0x44, 0x76, 0xd7, 0x38, 0xf9, 0x6f, 0xfd, 0xfc, 0x68, 0x56, 0x31, 0xdf,
	0x08, 0x2f, 0x5f, 0xbb, 0x9a, 0x23, 0x79, 0xd5, 0x6b, 0xff, 0x96, 0x1c, 0x5c, 0xb0, 0xc4, 0xe3,
	0xd1, 0x2f, 0xb7, 0xa5, 0xd2, 0x8f, 0x2e, 0x97, 0x1b, 0xaa, 0xad, 0x37, 0x54, 0x5b, 0x6e, 0x29,
	0x5a, 0x6d, 0x29, 0x7a, 0xdb, 0x52, 0xf4, 0xbc, 0xa3, 0xda, 0x6a, 0x47, 0xb5, 0xf5, 0x8e, 0x6a,
	0x37, 0x47, 0x41, 0x28, 0x66, 0x73, 0xd7, 0xf2, 0x20, 0xb6, 0xcb, 0xad, 0x50, 0xc7, 0x71, 0x3e,
	0xbd, 0xb3, 0x1f, 0x3f, 0x57, 0x44, 0x2c, 0x52, 0x9e, 0xbb, 0xba, 0xfc, 0xaf, 0xd3, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0x05, 0xd5, 0x12, 0x41, 0x02, 0x00, 0x00,
}

func (this *Plan) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Plan)
	if !ok {
		that2, ok := that.(Plan)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	return true
}
func (this *SoftwareUpgradeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SoftwareUpgradeProposal)
	if !ok {
		that2, ok := that.(SoftwareUpgradeProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Plan.Equal(&that1.Plan) {
		return false
	}
	return true
}
func (this *CancelSoftwareUpgradeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CancelSoftwareUpgradeProposal)
	if !ok {
		that2, ok := that.(CancelSoftwareUpgradeProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	return true
}
func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintUpgrade(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SoftwareUpgradeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SoftwareUpgradeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SoftwareUpgradeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CancelSoftwareUpgradeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelSoftwareUpgradeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelSoftwareUpgradeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintUpgrade(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovUpgrade(uint64(m.Height))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func (m *SoftwareUpgradeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = m.Plan.Size()
	n += 1 + l + sovUpgrade(uint64(l))
	return n
}

func (m *CancelSoftwareUpgradeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovUpgrade(uint64(l))
	}
	return n
}

func sovUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpgrade(x uint64) (n int) {
	return sovUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func (m *SoftwareUpgradeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: SoftwareUpgradeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func (m *CancelSoftwareUpgradeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpgrade
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
			return fmt.Errorf("proto: CancelSoftwareUpgradeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelSoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpgrade
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
				return ErrInvalidLengthUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUpgrade
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
func skipUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
					return 0, ErrIntOverflowUpgrade
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
				return 0, ErrInvalidLengthUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUpgrade = fmt.Errorf("proto: unexpected end of group")
)
