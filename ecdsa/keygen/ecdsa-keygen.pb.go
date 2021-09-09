// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ecdsa-keygen.proto

package keygen

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

//
// Represents a BROADCAST message sent during Round 1 of the ECDSA TSS keygen protocol.
type KGRound1Message struct {
	Commitment           []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	PaillierN            []byte   `protobuf:"bytes,2,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	NTilde               []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1                   []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	Dlnproof_1           [][]byte `protobuf:"bytes,6,rep,name=dlnproof_1,json=dlnproof1,proto3" json:"dlnproof_1,omitempty"`
	Dlnproof_2           [][]byte `protobuf:"bytes,7,rep,name=dlnproof_2,json=dlnproof2,proto3" json:"dlnproof_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound1Message) Reset()         { *m = KGRound1Message{} }
func (m *KGRound1Message) String() string { return proto.CompactTextString(m) }
func (*KGRound1Message) ProtoMessage()    {}
func (*KGRound1Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9942204cc0c4f82, []int{0}
}

func (m *KGRound1Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound1Message.Unmarshal(m, b)
}
func (m *KGRound1Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound1Message.Marshal(b, m, deterministic)
}
func (m *KGRound1Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound1Message.Merge(m, src)
}
func (m *KGRound1Message) XXX_Size() int {
	return xxx_messageInfo_KGRound1Message.Size(m)
}
func (m *KGRound1Message) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound1Message.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound1Message proto.InternalMessageInfo

func (m *KGRound1Message) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *KGRound1Message) GetPaillierN() []byte {
	if m != nil {
		return m.PaillierN
	}
	return nil
}

func (m *KGRound1Message) GetNTilde() []byte {
	if m != nil {
		return m.NTilde
	}
	return nil
}

func (m *KGRound1Message) GetH1() []byte {
	if m != nil {
		return m.H1
	}
	return nil
}

func (m *KGRound1Message) GetH2() []byte {
	if m != nil {
		return m.H2
	}
	return nil
}

func (m *KGRound1Message) GetDlnproof_1() [][]byte {
	if m != nil {
		return m.Dlnproof_1
	}
	return nil
}

func (m *KGRound1Message) GetDlnproof_2() [][]byte {
	if m != nil {
		return m.Dlnproof_2
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message1 struct {
	Share                []byte   `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound2Message1) Reset()         { *m = KGRound2Message1{} }
func (m *KGRound2Message1) String() string { return proto.CompactTextString(m) }
func (*KGRound2Message1) ProtoMessage()    {}
func (*KGRound2Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9942204cc0c4f82, []int{1}
}

func (m *KGRound2Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound2Message1.Unmarshal(m, b)
}
func (m *KGRound2Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound2Message1.Marshal(b, m, deterministic)
}
func (m *KGRound2Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound2Message1.Merge(m, src)
}
func (m *KGRound2Message1) XXX_Size() int {
	return xxx_messageInfo_KGRound2Message1.Size(m)
}
func (m *KGRound2Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound2Message1.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound2Message1 proto.InternalMessageInfo

func (m *KGRound2Message1) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message2 struct {
	DeCommitment         [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound2Message2) Reset()         { *m = KGRound2Message2{} }
func (m *KGRound2Message2) String() string { return proto.CompactTextString(m) }
func (*KGRound2Message2) ProtoMessage()    {}
func (*KGRound2Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9942204cc0c4f82, []int{2}
}

func (m *KGRound2Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound2Message2.Unmarshal(m, b)
}
func (m *KGRound2Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound2Message2.Marshal(b, m, deterministic)
}
func (m *KGRound2Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound2Message2.Merge(m, src)
}
func (m *KGRound2Message2) XXX_Size() int {
	return xxx_messageInfo_KGRound2Message2.Size(m)
}
func (m *KGRound2Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound2Message2.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound2Message2 proto.InternalMessageInfo

func (m *KGRound2Message2) GetDeCommitment() [][]byte {
	if m != nil {
		return m.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 3 of the ECDSA TSS keygen protocol.
type KGRound3Message struct {
	PaillierProof        [][]byte `protobuf:"bytes,1,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KGRound3Message) Reset()         { *m = KGRound3Message{} }
func (m *KGRound3Message) String() string { return proto.CompactTextString(m) }
func (*KGRound3Message) ProtoMessage()    {}
func (*KGRound3Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9942204cc0c4f82, []int{3}
}

func (m *KGRound3Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KGRound3Message.Unmarshal(m, b)
}
func (m *KGRound3Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KGRound3Message.Marshal(b, m, deterministic)
}
func (m *KGRound3Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KGRound3Message.Merge(m, src)
}
func (m *KGRound3Message) XXX_Size() int {
	return xxx_messageInfo_KGRound3Message.Size(m)
}
func (m *KGRound3Message) XXX_DiscardUnknown() {
	xxx_messageInfo_KGRound3Message.DiscardUnknown(m)
}

var xxx_messageInfo_KGRound3Message proto.InternalMessageInfo

func (m *KGRound3Message) GetPaillierProof() [][]byte {
	if m != nil {
		return m.PaillierProof
	}
	return nil
}

func init() {
	proto.RegisterType((*KGRound1Message)(nil), "binance.tsslib.ecdsa.keygen.KGRound1Message")
	proto.RegisterType((*KGRound2Message1)(nil), "binance.tsslib.ecdsa.keygen.KGRound2Message1")
	proto.RegisterType((*KGRound2Message2)(nil), "binance.tsslib.ecdsa.keygen.KGRound2Message2")
	proto.RegisterType((*KGRound3Message)(nil), "binance.tsslib.ecdsa.keygen.KGRound3Message")
}

func init() { proto.RegisterFile("ecdsa-keygen.proto", fileDescriptor_b9942204cc0c4f82) }

var fileDescriptor_b9942204cc0c4f82 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4b, 0x84, 0x40,
	0x18, 0xc5, 0x59, 0xb7, 0x75, 0xe9, 0xc3, 0xb5, 0x18, 0x82, 0x06, 0xa2, 0x58, 0x8c, 0x60, 0x2f,
	0x19, 0x33, 0x7b, 0xa8, 0x73, 0x1d, 0x3a, 0x44, 0x11, 0x4b, 0xa7, 0x2e, 0xa2, 0xce, 0xd7, 0x3a,
	0xa4, 0x33, 0xe2, 0xd8, 0xa1, 0xbf, 0xb0, 0x7f, 0x2b, 0x1c, 0x47, 0x49, 0x3a, 0xbe, 0xdf, 0xfb,
	0x7c, 0xf8, 0xde, 0x00, 0xc1, 0x5c, 0x98, 0xf4, 0xfa, 0x13, 0xbf, 0xf7, 0xa8, 0xe2, 0xba, 0xd1,
	0xad, 0x26, 0x67, 0x99, 0x54, 0xa9, 0xca, 0x31, 0x6e, 0x8d, 0x29, 0x65, 0x16, 0xdb, 0x93, 0xb8,
	0x3f, 0x89, 0x7e, 0x66, 0x70, 0xf4, 0xf4, 0xb8, 0xd3, 0x5f, 0x4a, 0xb0, 0x67, 0x34, 0x26, 0xdd,
	0x23, 0xb9, 0x00, 0xc8, 0x75, 0x55, 0xc9, 0xb6, 0x42, 0xd5, 0xd2, 0xd9, 0x7a, 0xb6, 0x09, 0x76,
	0x7f, 0x08, 0x39, 0x07, 0xa8, 0x53, 0x59, 0x96, 0x12, 0x9b, 0x44, 0x51, 0xcf, 0xfa, 0x87, 0x03,
	0x79, 0x21, 0xa7, 0xb0, 0x54, 0x49, 0x2b, 0x4b, 0x81, 0x74, 0x6e, 0x3d, 0x5f, 0xbd, 0x75, 0x8a,
	0x84, 0xe0, 0x15, 0x8c, 0x1e, 0x58, 0xe6, 0x15, 0xcc, 0x6a, 0x4e, 0x17, 0x4e, 0xf3, 0x2e, 0x57,
	0x94, 0xaa, 0x6e, 0xb4, 0xfe, 0x48, 0x18, 0xf5, 0xd7, 0xf3, 0x2e, 0x77, 0x20, 0x6c, 0x62, 0x73,
	0xba, 0x9c, 0xda, 0x3c, 0xda, 0xc0, 0xb1, 0x2b, 0xc2, 0x5d, 0x11, 0x46, 0x4e, 0x60, 0x61, 0x8a,
	0xb4, 0x41, 0x57, 0xa2, 0x17, 0xd1, 0xed, 0xbf, 0x4b, 0x4e, 0x2e, 0x61, 0x25, 0x30, 0x99, 0xd4,
	0xee, 0xf2, 0x03, 0x81, 0x0f, 0x23, 0x8b, 0xee, 0xc6, 0xad, 0xb6, 0xc3, 0x56, 0x57, 0x10, 0x8e,
	0x5b, 0xd8, 0x1f, 0x71, 0x1f, 0xae, 0x06, 0xfa, 0xda, 0xc1, 0xfb, 0xf0, 0x3d, 0xb0, 0xb3, 0xdf,
	0xf4, 0xb3, 0x67, 0xbe, 0x7d, 0x9a, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xf7, 0xd8,
	0xfe, 0xb0, 0x01, 0x00, 0x00,
}
