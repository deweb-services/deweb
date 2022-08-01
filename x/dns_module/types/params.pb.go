// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dns_module/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	DomainPrice                  uint64   `protobuf:"varint,1,opt,name=domainPrice,proto3" json:"domainPrice,omitempty"`
	DomainExpirationHours        int64    `protobuf:"varint,2,opt,name=domainExpirationHours,proto3" json:"domainExpirationHours,omitempty"`
	DomainOwnerProlongationHours int64    `protobuf:"varint,3,opt,name=domainOwnerProlongationHours,proto3" json:"domainOwnerProlongationHours,omitempty"`
	SubDomainPrice               uint64   `protobuf:"varint,4,opt,name=subDomainPrice,proto3" json:"subDomainPrice,omitempty"`
	BlockTLDs                    []string `protobuf:"bytes,5,rep,name=blockTLDs,proto3" json:"blockTLDs,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bdbc71c42463c07, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDomainPrice() uint64 {
	if m != nil {
		return m.DomainPrice
	}
	return 0
}

func (m *Params) GetDomainExpirationHours() int64 {
	if m != nil {
		return m.DomainExpirationHours
	}
	return 0
}

func (m *Params) GetDomainOwnerProlongationHours() int64 {
	if m != nil {
		return m.DomainOwnerProlongationHours
	}
	return 0
}

func (m *Params) GetSubDomainPrice() uint64 {
	if m != nil {
		return m.SubDomainPrice
	}
	return 0
}

func (m *Params) GetBlockTLDs() []string {
	if m != nil {
		return m.BlockTLDs
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "dewebservices.domain.Params")
}

func init() { proto.RegisterFile("dns_module/params.proto", fileDescriptor_6bdbc71c42463c07) }

var fileDescriptor_6bdbc71c42463c07 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xc9, 0x2b, 0x8e,
	0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0x49, 0x2d, 0x4f, 0x4d, 0x2a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x2d, 0xd6, 0x4b, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x2b, 0xd0, 0x07, 0xb1, 0x20, 0x6a, 0x95, 0xbe, 0x31, 0x72, 0xb1, 0x05, 0x80, 0x35, 0x0b, 0x29,
	0x70, 0x71, 0x43, 0x94, 0x06, 0x14, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04,
	0x21, 0x0b, 0x09, 0x99, 0x70, 0x89, 0x42, 0xb8, 0xae, 0x15, 0x05, 0x99, 0x45, 0x89, 0x25, 0x99,
	0xf9, 0x79, 0x1e, 0xf9, 0xa5, 0x45, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0xd8, 0x25,
	0x85, 0x9c, 0xb8, 0x64, 0x20, 0x12, 0xfe, 0xe5, 0x79, 0xa9, 0x45, 0x01, 0x45, 0xf9, 0x39, 0xf9,
	0x79, 0xe9, 0x48, 0x9a, 0x99, 0xc1, 0x9a, 0xf1, 0xaa, 0x11, 0x52, 0xe3, 0xe2, 0x2b, 0x2e, 0x4d,
	0x72, 0x41, 0x72, 0x1e, 0x0b, 0xd8, 0x79, 0x68, 0xa2, 0x42, 0x4a, 0x5c, 0x9c, 0x49, 0x39, 0xf9,
	0xc9, 0xd9, 0x21, 0x3e, 0x2e, 0xc5, 0x12, 0xac, 0x0a, 0xcc, 0x1a, 0x9c, 0x4e, 0x2c, 0x27, 0xee,
	0xc9, 0x33, 0x04, 0x21, 0x84, 0xad, 0x58, 0x66, 0x2c, 0x90, 0x67, 0x70, 0xf2, 0x39, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xa3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0x70, 0x48, 0xea, 0xc2, 0x82, 0x12, 0xc2, 0xd5, 0xaf, 0xd0, 0x47, 0x0a,
	0xf9, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x68, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0xdb, 0x08, 0xec, 0x94, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockTLDs) > 0 {
		for iNdEx := len(m.BlockTLDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlockTLDs[iNdEx])
			copy(dAtA[i:], m.BlockTLDs[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.BlockTLDs[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.SubDomainPrice != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SubDomainPrice))
		i--
		dAtA[i] = 0x20
	}
	if m.DomainOwnerProlongationHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DomainOwnerProlongationHours))
		i--
		dAtA[i] = 0x18
	}
	if m.DomainExpirationHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DomainExpirationHours))
		i--
		dAtA[i] = 0x10
	}
	if m.DomainPrice != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DomainPrice))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DomainPrice != 0 {
		n += 1 + sovParams(uint64(m.DomainPrice))
	}
	if m.DomainExpirationHours != 0 {
		n += 1 + sovParams(uint64(m.DomainExpirationHours))
	}
	if m.DomainOwnerProlongationHours != 0 {
		n += 1 + sovParams(uint64(m.DomainOwnerProlongationHours))
	}
	if m.SubDomainPrice != 0 {
		n += 1 + sovParams(uint64(m.SubDomainPrice))
	}
	if len(m.BlockTLDs) > 0 {
		for _, s := range m.BlockTLDs {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainPrice", wireType)
			}
			m.DomainPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DomainPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainExpirationHours", wireType)
			}
			m.DomainExpirationHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DomainExpirationHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainOwnerProlongationHours", wireType)
			}
			m.DomainOwnerProlongationHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DomainOwnerProlongationHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubDomainPrice", wireType)
			}
			m.SubDomainPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubDomainPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTLDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockTLDs = append(m.BlockTLDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
