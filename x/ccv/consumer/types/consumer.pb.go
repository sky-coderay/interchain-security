// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/consumer/v1/consumer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CrossChainValidator defines the type used to store validator information
// internal to the consumer CCV module.  Note one cross chain validator entry is
// persisted for each consumer validator, where incoming VSC packets update this
// data, which is eventually forwarded to comet for consumer chain consensus.
//
// Note this type is only used internally to the consumer CCV module.
type CrossChainValidator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Power   int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	// pubkey is the consensus public key of the validator, as a Protobuf Any.
	Pubkey   *types.Any `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty" yaml:"consensus_pubkey"`
	OptedOut bool       `protobuf:"varint,4,opt,name=opted_out,json=optedOut,proto3" json:"opted_out,omitempty"`
}

func (m *CrossChainValidator) Reset()         { *m = CrossChainValidator{} }
func (m *CrossChainValidator) String() string { return proto.CompactTextString(m) }
func (*CrossChainValidator) ProtoMessage()    {}
func (*CrossChainValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b27a82b276e7f93, []int{0}
}
func (m *CrossChainValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossChainValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossChainValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossChainValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossChainValidator.Merge(m, src)
}
func (m *CrossChainValidator) XXX_Size() int {
	return m.Size()
}
func (m *CrossChainValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossChainValidator.DiscardUnknown(m)
}

var xxx_messageInfo_CrossChainValidator proto.InternalMessageInfo

func (m *CrossChainValidator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CrossChainValidator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *CrossChainValidator) GetPubkey() *types.Any {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *CrossChainValidator) GetOptedOut() bool {
	if m != nil {
		return m.OptedOut
	}
	return false
}

// A record storing the state of a slash packet sent to the provider chain
// which may bounce back and forth until handled by the provider.
//
// Note this type is only used internally to the consumer CCV module.
type SlashRecord struct {
	WaitingOnReply bool      `protobuf:"varint,1,opt,name=waiting_on_reply,json=waitingOnReply,proto3" json:"waiting_on_reply,omitempty"`
	SendTime       time.Time `protobuf:"bytes,2,opt,name=send_time,json=sendTime,proto3,stdtime" json:"send_time"`
}

func (m *SlashRecord) Reset()         { *m = SlashRecord{} }
func (m *SlashRecord) String() string { return proto.CompactTextString(m) }
func (*SlashRecord) ProtoMessage()    {}
func (*SlashRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b27a82b276e7f93, []int{1}
}
func (m *SlashRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SlashRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SlashRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SlashRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlashRecord.Merge(m, src)
}
func (m *SlashRecord) XXX_Size() int {
	return m.Size()
}
func (m *SlashRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SlashRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SlashRecord proto.InternalMessageInfo

func (m *SlashRecord) GetWaitingOnReply() bool {
	if m != nil {
		return m.WaitingOnReply
	}
	return false
}

func (m *SlashRecord) GetSendTime() time.Time {
	if m != nil {
		return m.SendTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*CrossChainValidator)(nil), "interchain_security.ccv.consumer.v1.CrossChainValidator")
	proto.RegisterType((*SlashRecord)(nil), "interchain_security.ccv.consumer.v1.SlashRecord")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/consumer/v1/consumer.proto", fileDescriptor_5b27a82b276e7f93)
}

var fileDescriptor_5b27a82b276e7f93 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x52, 0x28, 0xee, 0x06, 0x21, 0x64, 0x22, 0xe1, 0x06, 0xc9, 0x89, 0xcc, 0xc5, 0x97,
	0xda, 0x6a, 0xca, 0x09, 0x89, 0x43, 0xd3, 0x23, 0x87, 0x22, 0x83, 0x40, 0xe2, 0x62, 0xad, 0xd7,
	0x8b, 0x63, 0x61, 0xef, 0xac, 0xf6, 0xc7, 0x65, 0x79, 0x8a, 0x3e, 0x0c, 0xaf, 0x80, 0x54, 0x71,
	0xea, 0x91, 0x53, 0x41, 0xc9, 0x1b, 0xf0, 0x04, 0xc8, 0x3f, 0x09, 0xe2, 0xe7, 0x36, 0xf3, 0xcd,
	0x7e, 0x33, 0xdf, 0xcc, 0x7e, 0x78, 0x51, 0x72, 0xcd, 0x24, 0x5d, 0x91, 0x92, 0xa7, 0x8a, 0x51,
	0x23, 0x4b, 0x6d, 0x63, 0x4a, 0x9b, 0x98, 0x02, 0x57, 0xa6, 0x66, 0x32, 0x6e, 0x8e, 0x77, 0x71,
	0x24, 0x24, 0x68, 0x70, 0x9f, 0xfc, 0x87, 0x13, 0x51, 0xda, 0x44, 0xbb, 0x77, 0xcd, 0xf1, 0xf4,
	0xb0, 0x00, 0x28, 0x2a, 0x16, 0x77, 0x94, 0xcc, 0xbc, 0x8f, 0x09, 0xb7, 0x3d, 0x7f, 0x3a, 0x29,
	0xa0, 0x80, 0x2e, 0x8c, 0xdb, 0x68, 0x40, 0x0f, 0x29, 0xa8, 0x1a, 0x54, 0xda, 0x17, 0xfa, 0x64,
	0x28, 0xcd, 0xfe, 0xee, 0xa5, 0xcb, 0x9a, 0x29, 0x4d, 0x6a, 0xd1, 0x3f, 0x08, 0xbe, 0x20, 0xfc,
	0xf0, 0x4c, 0x82, 0x52, 0x67, 0xad, 0xa8, 0x37, 0xa4, 0x2a, 0x73, 0xa2, 0x41, 0xba, 0x1e, 0xbe,
	0x4b, 0xf2, 0x5c, 0x32, 0xa5, 0x3c, 0x34, 0x47, 0xe1, 0xbd, 0x64, 0x9b, 0xba, 0x13, 0x7c, 0x47,
	0xc0, 0x05, 0x93, 0xde, 0xad, 0x39, 0x0a, 0xf7, 0x92, 0x3e, 0x71, 0x09, 0xde, 0x17, 0x26, 0xfb,
	0xc0, 0xac, 0xb7, 0x37, 0x47, 0xe1, 0x78, 0x31, 0x89, 0xfa, 0xc9, 0xd1, 0x76, 0x72, 0x74, 0xca,
	0xed, 0xf2, 0xe4, 0xe7, 0xcd, 0xec, 0x91, 0x25, 0x75, 0xf5, 0x2c, 0x68, 0x37, 0x66, 0x5c, 0x19,
	0x95, 0xf6, 0xbc, 0xe0, 0xeb, 0xe7, 0xa3, 0xc9, 0xa0, 0x9d, 0x4a, 0x2b, 0x34, 0x44, 0x2f, 0x4d,
	0xf6, 0x82, 0xd9, 0x64, 0x68, 0xec, 0x3e, 0xc6, 0x07, 0x20, 0x34, 0xcb, 0x53, 0x30, 0xda, 0xbb,
	0x3d, 0x47, 0xa1, 0x93, 0x38, 0x1d, 0x70, 0x6e, 0x74, 0xf0, 0x09, 0x8f, 0x5f, 0x55, 0x44, 0xad,
	0x12, 0x46, 0x41, 0xe6, 0x6e, 0x88, 0x1f, 0x5c, 0x90, 0x52, 0x97, 0xbc, 0x48, 0x81, 0xa7, 0x92,
	0x89, 0xca, 0x76, 0x7b, 0x38, 0xc9, 0xfd, 0x01, 0x3f, 0xe7, 0x49, 0x8b, 0xba, 0xa7, 0xf8, 0x40,
	0x31, 0x9e, 0xa7, 0xed, 0x61, 0xba, 0x95, 0xc6, 0x8b, 0xe9, 0x3f, 0xda, 0x5f, 0x6f, 0xaf, 0xb6,
	0x74, 0xae, 0x6e, 0x66, 0xa3, 0xcb, 0xef, 0x33, 0x94, 0x38, 0x2d, 0xad, 0x2d, 0x2c, 0xdf, 0x5e,
	0xad, 0x7d, 0x74, 0xbd, 0xf6, 0xd1, 0x8f, 0xb5, 0x8f, 0x2e, 0x37, 0xfe, 0xe8, 0x7a, 0xe3, 0x8f,
	0xbe, 0x6d, 0xfc, 0xd1, 0xbb, 0xe7, 0x45, 0xa9, 0x57, 0x26, 0x8b, 0x28, 0xd4, 0xc3, 0xbf, 0xc4,
	0xbf, 0x1d, 0x70, 0xb4, 0x73, 0x4d, 0xf3, 0x34, 0xfe, 0xf8, 0xa7, 0x75, 0xb4, 0x15, 0x4c, 0x65,
	0xfb, 0x9d, 0x80, 0x93, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0xf1, 0xa5, 0x1e, 0x6b, 0x02,
	0x00, 0x00,
}

func (m *CrossChainValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossChainValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossChainValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OptedOut {
		i--
		if m.OptedOut {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Pubkey != nil {
		{
			size, err := m.Pubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConsumer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Power != 0 {
		i = encodeVarintConsumer(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintConsumer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SlashRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlashRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SlashRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.SendTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.SendTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintConsumer(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if m.WaitingOnReply {
		i--
		if m.WaitingOnReply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsumer(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsumer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrossChainValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovConsumer(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovConsumer(uint64(m.Power))
	}
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovConsumer(uint64(l))
	}
	if m.OptedOut {
		n += 2
	}
	return n
}

func (m *SlashRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WaitingOnReply {
		n += 2
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.SendTime)
	n += 1 + l + sovConsumer(uint64(l))
	return n
}

func sovConsumer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsumer(x uint64) (n int) {
	return sovConsumer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrossChainValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsumer
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
			return fmt.Errorf("proto: CrossChainValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossChainValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsumer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsumer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pubkey == nil {
				m.Pubkey = &types.Any{}
			}
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptedOut", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
			m.OptedOut = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConsumer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsumer
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
func (m *SlashRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsumer
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
			return fmt.Errorf("proto: SlashRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlashRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WaitingOnReply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
			m.WaitingOnReply = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumer
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
				return ErrInvalidLengthConsumer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsumer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.SendTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsumer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsumer
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
func skipConsumer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsumer
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
					return 0, ErrIntOverflowConsumer
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
					return 0, ErrIntOverflowConsumer
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
				return 0, ErrInvalidLengthConsumer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsumer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsumer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsumer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsumer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsumer = fmt.Errorf("proto: unexpected end of group")
)
