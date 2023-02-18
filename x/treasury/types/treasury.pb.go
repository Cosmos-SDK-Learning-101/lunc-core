// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/treasury/v1beta1/treasury.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the oracle module.
type Params struct {
	TaxPolicy               PolicyConstraints                      `protobuf:"bytes,1,opt,name=tax_policy,json=taxPolicy,proto3" json:"tax_policy" yaml:"tax_policy"`
	RewardPolicy            PolicyConstraints                      `protobuf:"bytes,2,opt,name=reward_policy,json=rewardPolicy,proto3" json:"reward_policy" yaml:"reward_policy"`
	SeigniorageBurdenTarget github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=seigniorage_burden_target,json=seigniorageBurdenTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"seigniorage_burden_target" yaml:"seigniorage_burden_target"`
	MiningIncrement         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=mining_increment,json=miningIncrement,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mining_increment" yaml:"mining_increment"`
	WindowShort             uint64                                 `protobuf:"varint,5,opt,name=window_short,json=windowShort,proto3" json:"window_short,omitempty" yaml:"window_short"`
	WindowLong              uint64                                 `protobuf:"varint,6,opt,name=window_long,json=windowLong,proto3" json:"window_long,omitempty" yaml:"window_long"`
	WindowProbation         uint64                                 `protobuf:"varint,7,opt,name=window_probation,json=windowProbation,proto3" json:"window_probation,omitempty" yaml:"window_probation"`
	BurnTaxSplit            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=burn_tax_split,json=burnTaxSplit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burn_tax_split" yaml:"burn_tax_split"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{0}
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

func (m *Params) GetTaxPolicy() PolicyConstraints {
	if m != nil {
		return m.TaxPolicy
	}
	return PolicyConstraints{}
}

func (m *Params) GetRewardPolicy() PolicyConstraints {
	if m != nil {
		return m.RewardPolicy
	}
	return PolicyConstraints{}
}

func (m *Params) GetWindowShort() uint64 {
	if m != nil {
		return m.WindowShort
	}
	return 0
}

func (m *Params) GetWindowLong() uint64 {
	if m != nil {
		return m.WindowLong
	}
	return 0
}

func (m *Params) GetWindowProbation() uint64 {
	if m != nil {
		return m.WindowProbation
	}
	return 0
}

// PolicyConstraints - defines policy constraints can be applied in tax & reward policies
type PolicyConstraints struct {
	RateMin       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=rate_min,json=rateMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate_min" yaml:"rate_min"`
	RateMax       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate_max,json=rateMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate_max" yaml:"rate_max"`
	Cap           types.Coin                             `protobuf:"bytes,3,opt,name=cap,proto3" json:"cap" yaml:"cap"`
	ChangeRateMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=change_rate_max,json=changeRateMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"change_rate_max" yaml:"change_rate_max"`
}

func (m *PolicyConstraints) Reset()      { *m = PolicyConstraints{} }
func (*PolicyConstraints) ProtoMessage() {}
func (*PolicyConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{1}
}
func (m *PolicyConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyConstraints.Merge(m, src)
}
func (m *PolicyConstraints) XXX_Size() int {
	return m.Size()
}
func (m *PolicyConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyConstraints proto.InternalMessageInfo

func (m *PolicyConstraints) GetCap() types.Coin {
	if m != nil {
		return m.Cap
	}
	return types.Coin{}
}

// EpochTaxProceeds represents the tax amount
// collected at the current epoch
type EpochTaxProceeds struct {
	TaxProceeds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=tax_proceeds,json=taxProceeds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tax_proceeds" yaml:"tax_proceeds"`
}

func (m *EpochTaxProceeds) Reset()         { *m = EpochTaxProceeds{} }
func (m *EpochTaxProceeds) String() string { return proto.CompactTextString(m) }
func (*EpochTaxProceeds) ProtoMessage()    {}
func (*EpochTaxProceeds) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{2}
}
func (m *EpochTaxProceeds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochTaxProceeds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochTaxProceeds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochTaxProceeds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochTaxProceeds.Merge(m, src)
}
func (m *EpochTaxProceeds) XXX_Size() int {
	return m.Size()
}
func (m *EpochTaxProceeds) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochTaxProceeds.DiscardUnknown(m)
}

var xxx_messageInfo_EpochTaxProceeds proto.InternalMessageInfo

func (m *EpochTaxProceeds) GetTaxProceeds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TaxProceeds
	}
	return nil
}

// EpochInitialIssuance represents initial issuance
// of the currrent epoch
type EpochInitialIssuance struct {
	Issuance github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=issuance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"issuance" yaml:"issuance"`
}

func (m *EpochInitialIssuance) Reset()         { *m = EpochInitialIssuance{} }
func (m *EpochInitialIssuance) String() string { return proto.CompactTextString(m) }
func (*EpochInitialIssuance) ProtoMessage()    {}
func (*EpochInitialIssuance) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{3}
}
func (m *EpochInitialIssuance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInitialIssuance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInitialIssuance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInitialIssuance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInitialIssuance.Merge(m, src)
}
func (m *EpochInitialIssuance) XXX_Size() int {
	return m.Size()
}
func (m *EpochInitialIssuance) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInitialIssuance.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInitialIssuance proto.InternalMessageInfo

func (m *EpochInitialIssuance) GetIssuance() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Issuance
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "terra.treasury.v1beta1.Params")
	proto.RegisterType((*PolicyConstraints)(nil), "terra.treasury.v1beta1.PolicyConstraints")
	proto.RegisterType((*EpochTaxProceeds)(nil), "terra.treasury.v1beta1.EpochTaxProceeds")
	proto.RegisterType((*EpochInitialIssuance)(nil), "terra.treasury.v1beta1.EpochInitialIssuance")
}

func init() {
	proto.RegisterFile("terra/treasury/v1beta1/treasury.proto", fileDescriptor_353bb3a9c554268e)
}

var fileDescriptor_353bb3a9c554268e = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xbb, 0x6f, 0xe3, 0x36,
	0x1c, 0xc7, 0xad, 0x3a, 0x75, 0x1c, 0xda, 0xa9, 0x13, 0x25, 0x4d, 0x94, 0xb4, 0xb0, 0x0c, 0x02,
	0x2d, 0xdc, 0x21, 0x32, 0x92, 0x0e, 0x05, 0xb2, 0x14, 0x75, 0xda, 0x26, 0x06, 0x5a, 0xc0, 0x50,
	0x32, 0x15, 0x05, 0x04, 0x5a, 0x26, 0x64, 0xa2, 0x16, 0x29, 0x90, 0x74, 0x23, 0x77, 0xef, 0x56,
	0x14, 0x87, 0x9b, 0x82, 0x9b, 0x32, 0xdf, 0x5f, 0x92, 0x31, 0xe3, 0xe1, 0x06, 0xdf, 0x21, 0x59,
	0x6e, 0xf6, 0x5f, 0x70, 0x10, 0x49, 0xbf, 0x72, 0x4f, 0xe3, 0x26, 0xf1, 0xf7, 0xfa, 0xfc, 0xbe,
	0xa4, 0xf8, 0x00, 0xdf, 0x48, 0xcc, 0x39, 0x6a, 0x48, 0x8e, 0x91, 0x18, 0xf0, 0x61, 0xe3, 0xef,
	0xc3, 0x0e, 0x96, 0xe8, 0x70, 0xea, 0xf0, 0x12, 0xce, 0x24, 0xb3, 0x77, 0x54, 0x9a, 0x37, 0xf5,
	0x9a, 0xb4, 0xfd, 0xed, 0x88, 0x45, 0x4c, 0xa5, 0x34, 0xb2, 0x91, 0xce, 0xde, 0xaf, 0x86, 0x4c,
	0xc4, 0x4c, 0x34, 0x3a, 0x48, 0xe0, 0x29, 0x31, 0x64, 0x84, 0xea, 0x38, 0xbc, 0x2a, 0x80, 0x42,
	0x1b, 0x71, 0x14, 0x0b, 0x3b, 0x04, 0x40, 0xa2, 0x34, 0x48, 0x58, 0x9f, 0x84, 0x43, 0xc7, 0xaa,
	0x59, 0xf5, 0xd2, 0xd1, 0x77, 0xde, 0xdb, 0xbb, 0x79, 0x6d, 0x95, 0x75, 0xc2, 0xa8, 0x90, 0x1c,
	0x11, 0x2a, 0x45, 0x73, 0xef, 0x66, 0xe4, 0xe6, 0xc6, 0x23, 0x77, 0x73, 0x88, 0xe2, 0xfe, 0x31,
	0x9c, 0xa1, 0xa0, 0xbf, 0x26, 0x51, 0xaa, 0x0b, 0xec, 0x3e, 0x58, 0xe7, 0xf8, 0x12, 0xf1, 0xee,
	0xa4, 0xcf, 0x67, 0xcb, 0xf6, 0xf9, 0xda, 0xf4, 0xd9, 0xd6, 0x7d, 0x16, 0x68, 0xd0, 0x2f, 0x6b,
	0xdb, 0x74, 0xfb, 0xdf, 0x02, 0x7b, 0x02, 0x93, 0x88, 0x12, 0xc6, 0x51, 0x84, 0x83, 0xce, 0x80,
	0x77, 0x31, 0x0d, 0x24, 0xe2, 0x11, 0x96, 0x4e, 0xbe, 0x66, 0xd5, 0xd7, 0x9a, 0x7e, 0xc6, 0x7b,
	0x3e, 0x72, 0xbf, 0x8d, 0x88, 0xec, 0x0d, 0x3a, 0x5e, 0xc8, 0xe2, 0x86, 0x59, 0x34, 0xfd, 0x39,
	0x10, 0xdd, 0xbf, 0x1a, 0x72, 0x98, 0x60, 0xe1, 0xfd, 0x8c, 0xc3, 0xf1, 0xc8, 0xad, 0xe9, 0xce,
	0xef, 0x04, 0x43, 0x7f, 0x77, 0x2e, 0xd6, 0x54, 0xa1, 0x0b, 0x15, 0xb1, 0x25, 0xd8, 0x88, 0x09,
	0x25, 0x34, 0x0a, 0x08, 0x0d, 0x39, 0x8e, 0x31, 0x95, 0xce, 0x8a, 0x92, 0xd1, 0x5a, 0x5a, 0xc6,
	0xae, 0x96, 0xf1, 0x90, 0x07, 0xfd, 0x8a, 0x76, 0xb5, 0x26, 0x1e, 0xfb, 0x18, 0x94, 0x2f, 0x09,
	0xed, 0xb2, 0xcb, 0x40, 0xf4, 0x18, 0x97, 0xce, 0xe7, 0x35, 0xab, 0xbe, 0xd2, 0xdc, 0x1d, 0x8f,
	0xdc, 0x2d, 0xcd, 0x98, 0x8f, 0x42, 0xbf, 0xa4, 0xcd, 0xf3, 0xcc, 0xb2, 0x7f, 0x00, 0xc6, 0x0c,
	0xfa, 0x8c, 0x46, 0x4e, 0x41, 0x95, 0xee, 0x8c, 0x47, 0xae, 0xbd, 0x50, 0x9a, 0x05, 0xa1, 0x0f,
	0xb4, 0xf5, 0x1b, 0xa3, 0x91, 0xfd, 0x2b, 0xd8, 0x30, 0xb1, 0x84, 0xb3, 0x0e, 0x92, 0x84, 0x51,
	0x67, 0x55, 0x55, 0x7f, 0x35, 0x13, 0xff, 0x30, 0x03, 0xfa, 0x15, 0xed, 0x6a, 0x4f, 0x3c, 0x76,
	0x0c, 0xbe, 0xe8, 0x0c, 0x78, 0xb6, 0xb6, 0x69, 0x20, 0x92, 0x3e, 0x91, 0x4e, 0x51, 0x2d, 0xd8,
	0xe9, 0xd2, 0x0b, 0xf6, 0xa5, 0xee, 0xb9, 0x48, 0x83, 0x7e, 0x39, 0x73, 0x5c, 0xa0, 0xf4, 0x3c,
	0x33, 0x8f, 0x8b, 0x57, 0xd7, 0x6e, 0xee, 0xd5, 0xb5, 0x6b, 0xc1, 0xff, 0xf2, 0x60, 0xf3, 0x8d,
	0xed, 0x67, 0xff, 0x09, 0x8a, 0x1c, 0x49, 0x1c, 0xc4, 0x84, 0xaa, 0x33, 0xb2, 0xd6, 0xfc, 0x69,
	0x69, 0x21, 0x15, 0xb3, 0x75, 0x0d, 0x07, 0xfa, 0xab, 0xd9, 0xf0, 0x77, 0x42, 0x67, 0x74, 0x94,
	0xaa, 0x93, 0xf1, 0xc9, 0x74, 0x94, 0x4e, 0xe8, 0x28, 0xb5, 0x7f, 0x04, 0xf9, 0x10, 0x25, 0x6a,
	0xdf, 0x97, 0x8e, 0xf6, 0x3c, 0x5d, 0xef, 0x65, 0x57, 0xc3, 0xf4, 0xbc, 0x9d, 0x30, 0x42, 0x9b,
	0xb6, 0x39, 0x62, 0x40, 0x93, 0x42, 0x94, 0x40, 0x3f, 0xab, 0xb4, 0x13, 0x50, 0x09, 0x7b, 0x88,
	0x46, 0x38, 0x98, 0xaa, 0xd4, 0xbb, 0xf7, 0x6c, 0x69, 0x95, 0x3b, 0x86, 0xbd, 0x88, 0x83, 0xfe,
	0xba, 0xf6, 0xf8, 0x5a, 0xf2, 0xdc, 0xef, 0x78, 0x62, 0x81, 0x8d, 0x5f, 0x12, 0x16, 0xf6, 0x2e,
	0x50, 0xda, 0xe6, 0x2c, 0xc4, 0xb8, 0x2b, 0xec, 0x7f, 0x2d, 0x50, 0x56, 0x37, 0x8d, 0x71, 0x38,
	0x56, 0x2d, 0xff, 0xfe, 0xb9, 0x9d, 0x9a, 0xb9, 0x6d, 0xcd, 0x5d, 0x53, 0xa6, 0x18, 0x3e, 0x7d,
	0xe1, 0xd6, 0x3f, 0x62, 0x02, 0x19, 0x47, 0xf8, 0x25, 0x39, 0xd3, 0x01, 0x1f, 0x5b, 0x60, 0x5b,
	0x89, 0x6b, 0x51, 0x22, 0x09, 0xea, 0xb7, 0x84, 0x18, 0x20, 0x1a, 0x62, 0xfb, 0x1f, 0x50, 0x24,
	0x66, 0xfc, 0x61, 0x6d, 0x27, 0x46, 0x9b, 0xf9, 0x83, 0x93, 0xc2, 0xe5, 0x74, 0x4d, 0xfb, 0x35,
	0xcf, 0x6e, 0xee, 0xaa, 0xd6, 0xed, 0x5d, 0xd5, 0x7a, 0x79, 0x57, 0xb5, 0x1e, 0xdd, 0x57, 0x73,
	0xb7, 0xf7, 0xd5, 0xdc, 0xb3, 0xfb, 0x6a, 0xee, 0x0f, 0x6f, 0x9e, 0xd6, 0x47, 0x42, 0x90, 0xf0,
	0x40, 0xbf, 0x3e, 0x21, 0xe3, 0xb8, 0x91, 0xce, 0x1e, 0x21, 0x45, 0xee, 0x14, 0xd4, 0x63, 0xf1,
	0xfd, 0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xb0, 0x90, 0xa5, 0xa3, 0x06, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.TaxPolicy.Equal(&that1.TaxPolicy) {
		return false
	}
	if !this.RewardPolicy.Equal(&that1.RewardPolicy) {
		return false
	}
	if !this.SeigniorageBurdenTarget.Equal(that1.SeigniorageBurdenTarget) {
		return false
	}
	if !this.MiningIncrement.Equal(that1.MiningIncrement) {
		return false
	}
	if this.WindowShort != that1.WindowShort {
		return false
	}
	if this.WindowLong != that1.WindowLong {
		return false
	}
	if this.WindowProbation != that1.WindowProbation {
		return false
	}
	if !this.BurnTaxSplit.Equal(that1.BurnTaxSplit) {
		return false
	}
	return true
}
func (this *PolicyConstraints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PolicyConstraints)
	if !ok {
		that2, ok := that.(PolicyConstraints)
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
	if !this.RateMin.Equal(that1.RateMin) {
		return false
	}
	if !this.RateMax.Equal(that1.RateMax) {
		return false
	}
	if !this.Cap.Equal(&that1.Cap) {
		return false
	}
	if !this.ChangeRateMax.Equal(that1.ChangeRateMax) {
		return false
	}
	return true
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
	{
		size := m.BurnTaxSplit.Size()
		i -= size
		if _, err := m.BurnTaxSplit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.WindowProbation != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowProbation))
		i--
		dAtA[i] = 0x38
	}
	if m.WindowLong != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowLong))
		i--
		dAtA[i] = 0x30
	}
	if m.WindowShort != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowShort))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MiningIncrement.Size()
		i -= size
		if _, err := m.MiningIncrement.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SeigniorageBurdenTarget.Size()
		i -= size
		if _, err := m.SeigniorageBurdenTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.RewardPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.TaxPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PolicyConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ChangeRateMax.Size()
		i -= size
		if _, err := m.ChangeRateMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Cap.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RateMax.Size()
		i -= size
		if _, err := m.RateMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.RateMin.Size()
		i -= size
		if _, err := m.RateMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EpochTaxProceeds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochTaxProceeds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochTaxProceeds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaxProceeds) > 0 {
		for iNdEx := len(m.TaxProceeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxProceeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTreasury(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EpochInitialIssuance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInitialIssuance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInitialIssuance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Issuance) > 0 {
		for iNdEx := len(m.Issuance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issuance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTreasury(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreasury(v)
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
	l = m.TaxPolicy.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.RewardPolicy.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.SeigniorageBurdenTarget.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.MiningIncrement.Size()
	n += 1 + l + sovTreasury(uint64(l))
	if m.WindowShort != 0 {
		n += 1 + sovTreasury(uint64(m.WindowShort))
	}
	if m.WindowLong != 0 {
		n += 1 + sovTreasury(uint64(m.WindowLong))
	}
	if m.WindowProbation != 0 {
		n += 1 + sovTreasury(uint64(m.WindowProbation))
	}
	l = m.BurnTaxSplit.Size()
	n += 1 + l + sovTreasury(uint64(l))
	return n
}

func (m *PolicyConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RateMin.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.RateMax.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.Cap.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.ChangeRateMax.Size()
	n += 1 + l + sovTreasury(uint64(l))
	return n
}

func (m *EpochTaxProceeds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TaxProceeds) > 0 {
		for _, e := range m.TaxProceeds {
			l = e.Size()
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func (m *EpochInitialIssuance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Issuance) > 0 {
		for _, e := range m.Issuance {
			l = e.Size()
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func sovTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreasury(x uint64) (n int) {
	return sovTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeigniorageBurdenTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SeigniorageBurdenTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiningIncrement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MiningIncrement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowShort", wireType)
			}
			m.WindowShort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowShort |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowLong", wireType)
			}
			m.WindowLong = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowLong |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowProbation", wireType)
			}
			m.WindowProbation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowProbation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnTaxSplit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnTaxSplit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *PolicyConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: PolicyConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeRateMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChangeRateMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *EpochTaxProceeds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: EpochTaxProceeds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochTaxProceeds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxProceeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxProceeds = append(m.TaxProceeds, types.Coin{})
			if err := m.TaxProceeds[len(m.TaxProceeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *EpochInitialIssuance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: EpochInitialIssuance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInitialIssuance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuance = append(m.Issuance, types.Coin{})
			if err := m.Issuance[len(m.Issuance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func skipTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
				return 0, ErrInvalidLengthTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreasury = fmt.Errorf("proto: unexpected end of group")
)
