// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/oracle/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params                        Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeederDelegations             []FeederDelegation             `protobuf:"bytes,2,rep,name=feeder_delegations,json=feederDelegations,proto3" json:"feeder_delegations"`
	ExchangeRates                 ExchangeRateTuples             `protobuf:"bytes,3,rep,name=exchange_rates,json=exchangeRates,proto3,castrepeated=ExchangeRateTuples" json:"exchange_rates"`
	MissCounters                  []MissCounter                  `protobuf:"bytes,4,rep,name=miss_counters,json=missCounters,proto3" json:"miss_counters"`
	AggregateExchangeRatePrevotes []AggregateExchangeRatePrevote `protobuf:"bytes,5,rep,name=aggregate_exchange_rate_prevotes,json=aggregateExchangeRatePrevotes,proto3" json:"aggregate_exchange_rate_prevotes"`
	AggregateExchangeRateVotes    []AggregateExchangeRateVote    `protobuf:"bytes,6,rep,name=aggregate_exchange_rate_votes,json=aggregateExchangeRateVotes,proto3" json:"aggregate_exchange_rate_votes"`
	TobinTaxes                    []TobinTax                     `protobuf:"bytes,7,rep,name=tobin_taxes,json=tobinTaxes,proto3" json:"tobin_taxes"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff46fd82c752f1f, []int{0}
}

func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}

func (m *GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeederDelegations() []FeederDelegation {
	if m != nil {
		return m.FeederDelegations
	}
	return nil
}

func (m *GenesisState) GetExchangeRates() ExchangeRateTuples {
	if m != nil {
		return m.ExchangeRates
	}
	return nil
}

func (m *GenesisState) GetMissCounters() []MissCounter {
	if m != nil {
		return m.MissCounters
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRatePrevotes() []AggregateExchangeRatePrevote {
	if m != nil {
		return m.AggregateExchangeRatePrevotes
	}
	return nil
}

func (m *GenesisState) GetAggregateExchangeRateVotes() []AggregateExchangeRateVote {
	if m != nil {
		return m.AggregateExchangeRateVotes
	}
	return nil
}

func (m *GenesisState) GetTobinTaxes() []TobinTax {
	if m != nil {
		return m.TobinTaxes
	}
	return nil
}

// FeederDelegation is the address for where oracle feeder authority are
// delegated to. By default this struct is only used at genesis to feed in
// default feeder addresses.
type FeederDelegation struct {
	FeederAddress    string `protobuf:"bytes,1,opt,name=feeder_address,json=feederAddress,proto3" json:"feeder_address,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (m *FeederDelegation) Reset()         { *m = FeederDelegation{} }
func (m *FeederDelegation) String() string { return proto.CompactTextString(m) }
func (*FeederDelegation) ProtoMessage()    {}
func (*FeederDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff46fd82c752f1f, []int{1}
}

func (m *FeederDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *FeederDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeederDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *FeederDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeederDelegation.Merge(m, src)
}

func (m *FeederDelegation) XXX_Size() int {
	return m.Size()
}

func (m *FeederDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeederDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_FeederDelegation proto.InternalMessageInfo

func (m *FeederDelegation) GetFeederAddress() string {
	if m != nil {
		return m.FeederAddress
	}
	return ""
}

func (m *FeederDelegation) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

// MissCounter defines an miss counter and validator address pair used in
// oracle module's genesis state
type MissCounter struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	MissCounter      uint64 `protobuf:"varint,2,opt,name=miss_counter,json=missCounter,proto3" json:"miss_counter,omitempty"`
}

func (m *MissCounter) Reset()         { *m = MissCounter{} }
func (m *MissCounter) String() string { return proto.CompactTextString(m) }
func (*MissCounter) ProtoMessage()    {}
func (*MissCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff46fd82c752f1f, []int{2}
}

func (m *MissCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MissCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MissCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissCounter.Merge(m, src)
}

func (m *MissCounter) XXX_Size() int {
	return m.Size()
}

func (m *MissCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MissCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MissCounter proto.InternalMessageInfo

func (m *MissCounter) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MissCounter) GetMissCounter() uint64 {
	if m != nil {
		return m.MissCounter
	}
	return 0
}

// TobinTax defines an denom and tobin_tax pair used in
// oracle module's genesis state
type TobinTax struct {
	Denom    string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	TobinTax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=tobin_tax,json=tobinTax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tobin_tax"`
}

func (m *TobinTax) Reset()         { *m = TobinTax{} }
func (m *TobinTax) String() string { return proto.CompactTextString(m) }
func (*TobinTax) ProtoMessage()    {}
func (*TobinTax) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ff46fd82c752f1f, []int{3}
}

func (m *TobinTax) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *TobinTax) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TobinTax.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *TobinTax) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TobinTax.Merge(m, src)
}

func (m *TobinTax) XXX_Size() int {
	return m.Size()
}

func (m *TobinTax) XXX_DiscardUnknown() {
	xxx_messageInfo_TobinTax.DiscardUnknown(m)
}

var xxx_messageInfo_TobinTax proto.InternalMessageInfo

func (m *TobinTax) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "terra.oracle.v1beta1.GenesisState")
	proto.RegisterType((*FeederDelegation)(nil), "terra.oracle.v1beta1.FeederDelegation")
	proto.RegisterType((*MissCounter)(nil), "terra.oracle.v1beta1.MissCounter")
	proto.RegisterType((*TobinTax)(nil), "terra.oracle.v1beta1.TobinTax")
}

func init() {
	proto.RegisterFile("terra/oracle/v1beta1/genesis.proto", fileDescriptor_7ff46fd82c752f1f)
}

var fileDescriptor_7ff46fd82c752f1f = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0xfe, 0xfb, 0xb5, 0x9b, 0xb6, 0x6a, 0x57, 0x39, 0x44, 0xd1, 0xaf, 0xce, 0x1f,
	0x89, 0x52, 0x09, 0xd5, 0x56, 0xc3, 0x8d, 0x5b, 0x43, 0x0b, 0x42, 0x80, 0x54, 0x99, 0x88, 0x03,
	0x08, 0x59, 0x1b, 0x7b, 0xe2, 0x1a, 0x62, 0x6f, 0xb4, 0xb3, 0x89, 0xc2, 0x91, 0x37, 0xe0, 0x29,
	0x38, 0xf0, 0x24, 0x3d, 0xf6, 0x88, 0x38, 0x14, 0x94, 0xbc, 0x08, 0xf2, 0xee, 0x26, 0x31, 0xc5,
	0x45, 0xe2, 0x94, 0xec, 0xec, 0x67, 0xbe, 0xdf, 0x99, 0xf5, 0x68, 0x48, 0x4b, 0x82, 0x10, 0xcc,
	0xe5, 0x82, 0x05, 0x03, 0x70, 0xc7, 0x27, 0x3d, 0x90, 0xec, 0xc4, 0x8d, 0x20, 0x05, 0x8c, 0xd1,
	0x19, 0x0a, 0x2e, 0x39, 0xad, 0x28, 0xc6, 0xd1, 0x8c, 0x63, 0x98, 0x5a, 0x25, 0xe2, 0x11, 0x57,
	0x80, 0x9b, 0xfd, 0xd3, 0x6c, 0xad, 0x59, 0xa8, 0x67, 0x52, 0x15, 0xd2, 0xfa, 0xb2, 0x4e, 0xb6,
	0x9f, 0x6a, 0x83, 0x57, 0x92, 0x49, 0xa0, 0x8f, 0xc8, 0xc6, 0x90, 0x09, 0x96, 0x60, 0xd5, 0x6a,
	0x58, 0x47, 0xe5, 0xf6, 0xff, 0x4e, 0x91, 0xa1, 0x73, 0xa1, 0x98, 0xce, 0xda, 0xd5, 0x4d, 0xbd,
	0xe4, 0x99, 0x0c, 0xfa, 0x96, 0xd0, 0x3e, 0x40, 0x08, 0xc2, 0x0f, 0x61, 0x00, 0x11, 0x93, 0x31,
	0x4f, 0xb1, 0xba, 0xd2, 0x58, 0x3d, 0x2a, 0xb7, 0x0f, 0x8b, 0x75, 0x9e, 0x28, 0xfe, 0x6c, 0x81,
	0x1b, 0xc5, 0xfd, 0xfe, 0xad, 0x38, 0xd2, 0xf7, 0x64, 0x17, 0x26, 0xc1, 0x25, 0x4b, 0x23, 0xf0,
	0x05, 0x93, 0x80, 0xd5, 0x55, 0x25, 0x7c, 0xbf, 0x58, 0xf8, 0xdc, 0xb0, 0x1e, 0x93, 0xd0, 0x1d,
	0x0d, 0x07, 0xd0, 0xa9, 0x65, 0xca, 0x5f, 0x7f, 0xd4, 0xe9, 0x1f, 0x57, 0xe8, 0xed, 0x40, 0x2e,
	0x86, 0xf4, 0x05, 0xd9, 0x49, 0x62, 0x44, 0x3f, 0xe0, 0xa3, 0x54, 0x82, 0xc0, 0xea, 0x9a, 0xb2,
	0x6a, 0x16, 0x5b, 0xbd, 0x8c, 0x11, 0x1f, 0x6b, 0xd2, 0x94, 0xbf, 0x9d, 0x2c, 0x43, 0x48, 0x3f,
	0x59, 0xa4, 0xc1, 0xa2, 0x48, 0x64, 0xad, 0x80, 0xff, 0x5b, 0x13, 0xfe, 0x50, 0xc0, 0x98, 0x67,
	0xcd, 0xac, 0x2b, 0x87, 0x76, 0xb1, 0xc3, 0xe9, 0x3c, 0x3b, 0x5f, 0xfa, 0x85, 0x4e, 0x35, 0x96,
	0x07, 0xec, 0x2f, 0x0c, 0xd2, 0x09, 0x39, 0xb8, 0xab, 0x04, 0xed, 0xbf, 0xa1, 0xfc, 0xdd, 0x7f,
	0xf0, 0x7f, 0xbd, 0x34, 0xaf, 0xb1, 0xbb, 0x00, 0xa4, 0xe7, 0xa4, 0x2c, 0x79, 0x2f, 0x4e, 0x7d,
	0xc9, 0x26, 0x80, 0xd5, 0xff, 0x94, 0x8f, 0x5d, 0xec, 0xd3, 0xcd, 0xc0, 0x2e, 0x9b, 0x18, 0x59,
	0x22, 0xcd, 0x19, 0xb0, 0xd5, 0x27, 0x7b, 0xb7, 0x67, 0x85, 0xde, 0x23, 0xbb, 0x66, 0xde, 0x58,
	0x18, 0x0a, 0x40, 0x3d, 0xb3, 0x5b, 0xde, 0x8e, 0x8e, 0x9e, 0xea, 0x20, 0x7d, 0x40, 0xf6, 0xc7,
	0x6c, 0x10, 0x87, 0x4c, 0xf2, 0x25, 0xb9, 0xa2, 0xc8, 0xbd, 0xc5, 0x85, 0x81, 0x5b, 0xef, 0x48,
	0x39, 0xf7, 0x3d, 0x8b, 0x73, 0xad, 0xe2, 0x5c, 0xda, 0x24, 0xdb, 0xf9, 0xb1, 0x51, 0x1e, 0x6b,
	0x5e, 0x39, 0x37, 0x0c, 0xad, 0x84, 0x6c, 0xce, 0x9b, 0xa4, 0x15, 0xb2, 0x1e, 0x42, 0xca, 0x13,
	0xa3, 0xa7, 0x0f, 0xf4, 0x39, 0xd9, 0x5a, 0xbc, 0x97, 0xae, 0xb2, 0xe3, 0x64, 0xaf, 0xf1, 0xfd,
	0xa6, 0x7e, 0x18, 0xc5, 0xf2, 0x72, 0xd4, 0x73, 0x02, 0x9e, 0xb8, 0x01, 0xc7, 0x84, 0xa3, 0xf9,
	0x39, 0xc6, 0xf0, 0x83, 0x2b, 0x3f, 0x0e, 0x01, 0x9d, 0x33, 0x08, 0xbc, 0xcd, 0xf9, 0xbb, 0x75,
	0x9e, 0x5d, 0x4d, 0x6d, 0xeb, 0x7a, 0x6a, 0x5b, 0x3f, 0xa7, 0xb6, 0xf5, 0x79, 0x66, 0x97, 0xae,
	0x67, 0x76, 0xe9, 0xdb, 0xcc, 0x2e, 0xbd, 0x71, 0xf3, 0x5a, 0x03, 0x86, 0x18, 0x07, 0xc7, 0x7a,
	0x5d, 0x04, 0x5c, 0x80, 0x3b, 0x6e, 0xbb, 0x93, 0xf9, 0xe2, 0x50, 0xc2, 0xbd, 0x0d, 0xb5, 0x30,
	0x1e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x41, 0x32, 0xa6, 0xa5, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TobinTaxes) > 0 {
		for iNdEx := len(m.TobinTaxes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TobinTaxes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRateVotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRateVotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for iNdEx := len(m.AggregateExchangeRatePrevotes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AggregateExchangeRatePrevotes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MissCounters) > 0 {
		for iNdEx := len(m.MissCounters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissCounters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ExchangeRates) > 0 {
		for iNdEx := len(m.ExchangeRates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeederDelegations) > 0 {
		for iNdEx := len(m.FeederDelegations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeederDelegations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeederDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeederDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeederDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeederAddress) > 0 {
		i -= len(m.FeederAddress)
		copy(dAtA[i:], m.FeederAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.FeederAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MissCounter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TobinTax) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TobinTax) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TobinTax) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TobinTax.Size()
		i -= size
		if _, err := m.TobinTax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeederDelegations) > 0 {
		for _, e := range m.FeederDelegations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ExchangeRates) > 0 {
		for _, e := range m.ExchangeRates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissCounters) > 0 {
		for _, e := range m.MissCounters {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRatePrevotes) > 0 {
		for _, e := range m.AggregateExchangeRatePrevotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AggregateExchangeRateVotes) > 0 {
		for _, e := range m.AggregateExchangeRateVotes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TobinTaxes) > 0 {
		for _, e := range m.TobinTaxes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *FeederDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeederAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MissCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MissCounter != 0 {
		n += 1 + sovGenesis(uint64(m.MissCounter))
	}
	return n
}

func (m *TobinTax) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.TobinTax.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederDelegations = append(m.FeederDelegations, FeederDelegation{})
			if err := m.FeederDelegations[len(m.FeederDelegations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeRates = append(m.ExchangeRates, ExchangeRateTuple{})
			if err := m.ExchangeRates[len(m.ExchangeRates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissCounters = append(m.MissCounters, MissCounter{})
			if err := m.MissCounters[len(m.MissCounters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRatePrevotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRatePrevotes = append(m.AggregateExchangeRatePrevotes, AggregateExchangeRatePrevote{})
			if err := m.AggregateExchangeRatePrevotes[len(m.AggregateExchangeRatePrevotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateExchangeRateVotes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateExchangeRateVotes = append(m.AggregateExchangeRateVotes, AggregateExchangeRateVote{})
			if err := m.AggregateExchangeRateVotes[len(m.AggregateExchangeRateVotes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TobinTaxes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TobinTaxes = append(m.TobinTaxes, TobinTax{})
			if err := m.TobinTaxes[len(m.TobinTaxes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *FeederDelegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: FeederDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeederDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeederAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeederAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *MissCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MissCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissCounter", wireType)
			}
			m.MissCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *TobinTax) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: TobinTax: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TobinTax: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TobinTax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TobinTax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
