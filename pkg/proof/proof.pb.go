// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/proof/proof.proto

package proof

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import merkle "github.com/tendermint/tendermint/crypto/merkle"

import bytes "bytes"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KVProof struct {
	Height               int64         `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Proof                *merkle.Proof `protobuf:"bytes,2,opt,name=proof" json:"proof,omitempty"`
	Contract             []byte        `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	Key                  []byte        `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte        `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Version              []byte        `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *KVProof) Reset()         { *m = KVProof{} }
func (m *KVProof) String() string { return proto.CompactTextString(m) }
func (*KVProof) ProtoMessage()    {}
func (*KVProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_proof_999dd9998818ea93, []int{0}
}
func (m *KVProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *KVProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVProof.Merge(dst, src)
}
func (m *KVProof) XXX_Size() int {
	return m.Size()
}
func (m *KVProof) XXX_DiscardUnknown() {
	xxx_messageInfo_KVProof.DiscardUnknown(m)
}

var xxx_messageInfo_KVProof proto.InternalMessageInfo

func (m *KVProof) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *KVProof) GetProof() *merkle.Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *KVProof) GetContract() []byte {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *KVProof) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KVProof) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KVProof) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*KVProof)(nil), "proof.KVProof")
}
func (this *KVProof) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KVProof)
	if !ok {
		that2, ok := that.(KVProof)
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
	if this.Height != that1.Height {
		return false
	}
	if !this.Proof.Equal(that1.Proof) {
		return false
	}
	if !bytes.Equal(this.Contract, that1.Contract) {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	if !bytes.Equal(this.Version, that1.Version) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *KVProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVProof) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProof(dAtA, i, uint64(m.Height))
	}
	if m.Proof != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProof(dAtA, i, uint64(m.Proof.Size()))
		n1, err := m.Proof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Contract) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProof(dAtA, i, uint64(len(m.Contract)))
		i += copy(dAtA[i:], m.Contract)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProof(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProof(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProof(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKVProof(r randyProof, easy bool) *KVProof {
	this := &KVProof{}
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	if r.Intn(10) != 0 {
		this.Proof = merkle.NewPopulatedProof(r, easy)
	}
	v1 := r.Intn(100)
	this.Contract = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Contract[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(100)
	this.Key = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Key[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.Value = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.Value[i] = byte(r.Intn(256))
	}
	v4 := r.Intn(100)
	this.Version = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.Version[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProof(r, 7)
	}
	return this
}

type randyProof interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProof(r randyProof) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProof(r randyProof) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneProof(r)
	}
	return string(tmps)
}
func randUnrecognizedProof(r randyProof, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProof(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProof(dAtA []byte, r randyProof, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProof(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateProof(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateProof(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProof(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProof(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProof(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProof(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *KVProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovProof(uint64(m.Height))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovProof(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProof(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KVProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KVProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &merkle.Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = append(m.Contract[:0], dAtA[iNdEx:postIndex]...)
			if m.Contract == nil {
				m.Contract = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProof
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProof
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProof
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProof(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProof = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pkg/proof/proof.proto", fileDescriptor_proof_999dd9998818ea93) }

var fileDescriptor_proof_999dd9998818ea93 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xf5, 0x08, 0x49, 0x91, 0x01, 0x09, 0x59, 0x05, 0x59, 0x19, 0xac, 0x08, 0x96, 0x2c,
	0x24, 0x12, 0x8c, 0x6c, 0xac, 0x2c, 0x28, 0x03, 0x7b, 0x13, 0x5c, 0x27, 0x4a, 0x93, 0x17, 0x19,
	0xa7, 0x52, 0xff, 0x87, 0x81, 0x4f, 0x61, 0xe4, 0x13, 0xc0, 0x5f, 0xc1, 0x88, 0xfa, 0x1c, 0xaa,
	0x2c, 0xf6, 0x3d, 0xf7, 0xfa, 0xda, 0x7e, 0xec, 0x72, 0x68, 0x75, 0x3e, 0x18, 0xc4, 0xb5, 0x5f,
	0xb3, 0xc1, 0xa0, 0x45, 0x1e, 0x12, 0xc4, 0xb7, 0xba, 0xb1, 0xf5, 0x58, 0x66, 0x15, 0x76, 0xb9,
	0x46, 0x8d, 0x39, 0xa5, 0xe5, 0xb8, 0x26, 0x22, 0x20, 0xe5, 0x5b, 0xf1, 0xc3, 0xec, 0xb8, 0x55,
	0xfd, 0xab, 0x32, 0x5d, 0xd3, 0xdb, 0xb9, 0xac, 0xcc, 0x6e, 0xb0, 0x98, 0x77, 0xca, 0xb4, 0x1b,
	0x35, 0x6d, 0xbe, 0x7c, 0xfd, 0x0e, 0x6c, 0xf1, 0xf4, 0xf2, 0xbc, 0x7f, 0x97, 0x5f, 0xb1, 0xa8,
	0x56, 0x8d, 0xae, 0xad, 0x80, 0x04, 0xd2, 0xa0, 0x98, 0x88, 0xdf, 0x30, 0xff, 0x31, 0x71, 0x94,
	0x40, 0x7a, 0x7a, 0x77, 0x9e, 0x4d, 0x37, 0x50, 0xab, 0xf0, 0x19, 0x8f, 0xd9, 0x49, 0x85, 0xbd,
	0x35, 0xab, 0xca, 0x8a, 0x20, 0x81, 0xf4, 0xac, 0x38, 0x30, 0xbf, 0x60, 0x41, 0xab, 0x76, 0xe2,
	0x98, 0xec, 0xbd, 0xe4, 0x4b, 0x16, 0x6e, 0x57, 0x9b, 0x51, 0x89, 0x90, 0x3c, 0x0f, 0x5c, 0xb0,
	0xc5, 0x56, 0x99, 0xb7, 0x06, 0x7b, 0x11, 0x91, 0xff, 0x8f, 0x8f, 0xcb, 0xdf, 0x1f, 0x09, 0x1f,
	0x4e, 0xc2, 0xa7, 0x93, 0xf0, 0xe5, 0x24, 0x7c, 0x3b, 0x09, 0x65, 0x44, 0x33, 0xdc, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0xe9, 0xc5, 0x3e, 0x4f, 0x01, 0x00, 0x00,
}