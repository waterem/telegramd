// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.tl.handshake.proto

package mtproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// /////////////////////////////////////////////////////////////////////////////
// Server_DH_inner_data <--
//  + TL_server_DH_inner_data
//
type Server_DHInnerData_Data struct {
	Nonce       []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	G           int32  `protobuf:"varint,3,opt,name=g" json:"g,omitempty"`
	DhPrime     string `protobuf:"bytes,4,opt,name=dh_prime,json=dhPrime" json:"dh_prime,omitempty"`
	GA          string `protobuf:"bytes,5,opt,name=g_a,json=gA" json:"g_a,omitempty"`
	ServerTime  int32  `protobuf:"varint,6,opt,name=server_time,json=serverTime" json:"server_time,omitempty"`
}

func (m *Server_DHInnerData_Data) Reset()                    { *m = Server_DHInnerData_Data{} }
func (m *Server_DHInnerData_Data) String() string            { return proto.CompactTextString(m) }
func (*Server_DHInnerData_Data) ProtoMessage()               {}
func (*Server_DHInnerData_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Server_DHInnerData_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Server_DHInnerData_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *Server_DHInnerData_Data) GetG() int32 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *Server_DHInnerData_Data) GetDhPrime() string {
	if m != nil {
		return m.DhPrime
	}
	return ""
}

func (m *Server_DHInnerData_Data) GetGA() string {
	if m != nil {
		return m.GA
	}
	return ""
}

func (m *Server_DHInnerData_Data) GetServerTime() int32 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

type Server_DHInnerData struct {
	Constructor TLConstructor            `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *Server_DHInnerData_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *Server_DHInnerData) Reset()                    { *m = Server_DHInnerData{} }
func (m *Server_DHInnerData) String() string            { return proto.CompactTextString(m) }
func (*Server_DHInnerData) ProtoMessage()               {}
func (*Server_DHInnerData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Server_DHInnerData) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *Server_DHInnerData) GetData2() *Server_DHInnerData_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// server_DH_inner_data#b5890dba nonce:int128 server_nonce:int128 g:int dh_prime:string g_a:string server_time:int = Server_DH_inner_data;
type TLServer_DHInnerData struct {
	Data2 *Server_DHInnerData_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLServer_DHInnerData) Reset()                    { *m = TLServer_DHInnerData{} }
func (m *TLServer_DHInnerData) String() string            { return proto.CompactTextString(m) }
func (*TLServer_DHInnerData) ProtoMessage()               {}
func (*TLServer_DHInnerData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *TLServer_DHInnerData) GetData2() *Server_DHInnerData_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Client_DH_Inner_Data <--
//  + TL_client_DH_inner_data
//
type Client_DH_Inner_Data_Data struct {
	Nonce       []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	RetryId     int64  `protobuf:"varint,3,opt,name=retry_id,json=retryId" json:"retry_id,omitempty"`
	GB          string `protobuf:"bytes,4,opt,name=g_b,json=gB" json:"g_b,omitempty"`
}

func (m *Client_DH_Inner_Data_Data) Reset()                    { *m = Client_DH_Inner_Data_Data{} }
func (m *Client_DH_Inner_Data_Data) String() string            { return proto.CompactTextString(m) }
func (*Client_DH_Inner_Data_Data) ProtoMessage()               {}
func (*Client_DH_Inner_Data_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Client_DH_Inner_Data_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Client_DH_Inner_Data_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *Client_DH_Inner_Data_Data) GetRetryId() int64 {
	if m != nil {
		return m.RetryId
	}
	return 0
}

func (m *Client_DH_Inner_Data_Data) GetGB() string {
	if m != nil {
		return m.GB
	}
	return ""
}

type Client_DH_Inner_Data struct {
	Constructor TLConstructor              `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *Client_DH_Inner_Data_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *Client_DH_Inner_Data) Reset()                    { *m = Client_DH_Inner_Data{} }
func (m *Client_DH_Inner_Data) String() string            { return proto.CompactTextString(m) }
func (*Client_DH_Inner_Data) ProtoMessage()               {}
func (*Client_DH_Inner_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Client_DH_Inner_Data) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *Client_DH_Inner_Data) GetData2() *Client_DH_Inner_Data_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// client_DH_inner_data#6643b654 nonce:int128 server_nonce:int128 retry_id:long g_b:string = Client_DH_Inner_Data;
type TLClient_DHInnerData struct {
	Data2 *Client_DH_Inner_Data_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLClient_DHInnerData) Reset()                    { *m = TLClient_DHInnerData{} }
func (m *TLClient_DHInnerData) String() string            { return proto.CompactTextString(m) }
func (*TLClient_DHInnerData) ProtoMessage()               {}
func (*TLClient_DHInnerData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *TLClient_DHInnerData) GetData2() *Client_DH_Inner_Data_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Set_client_DH_params_answer <--
//  + TL_dh_gen_ok
//  + TL_dh_gen_retry
//  + TL_dh_gen_fail
//
type SetClient_DHParamsAnswer_Data struct {
	Nonce         []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce   []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	NewNonceHash1 []byte `protobuf:"bytes,3,opt,name=new_nonce_hash1,json=newNonceHash1,proto3" json:"new_nonce_hash1,omitempty"`
	NewNonceHash2 []byte `protobuf:"bytes,4,opt,name=new_nonce_hash2,json=newNonceHash2,proto3" json:"new_nonce_hash2,omitempty"`
	NewNonceHash3 []byte `protobuf:"bytes,5,opt,name=new_nonce_hash3,json=newNonceHash3,proto3" json:"new_nonce_hash3,omitempty"`
}

func (m *SetClient_DHParamsAnswer_Data) Reset()                    { *m = SetClient_DHParamsAnswer_Data{} }
func (m *SetClient_DHParamsAnswer_Data) String() string            { return proto.CompactTextString(m) }
func (*SetClient_DHParamsAnswer_Data) ProtoMessage()               {}
func (*SetClient_DHParamsAnswer_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *SetClient_DHParamsAnswer_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *SetClient_DHParamsAnswer_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *SetClient_DHParamsAnswer_Data) GetNewNonceHash1() []byte {
	if m != nil {
		return m.NewNonceHash1
	}
	return nil
}

func (m *SetClient_DHParamsAnswer_Data) GetNewNonceHash2() []byte {
	if m != nil {
		return m.NewNonceHash2
	}
	return nil
}

func (m *SetClient_DHParamsAnswer_Data) GetNewNonceHash3() []byte {
	if m != nil {
		return m.NewNonceHash3
	}
	return nil
}

type SetClient_DHParamsAnswer struct {
	Constructor TLConstructor                  `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *SetClient_DHParamsAnswer_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *SetClient_DHParamsAnswer) Reset()                    { *m = SetClient_DHParamsAnswer{} }
func (m *SetClient_DHParamsAnswer) String() string            { return proto.CompactTextString(m) }
func (*SetClient_DHParamsAnswer) ProtoMessage()               {}
func (*SetClient_DHParamsAnswer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *SetClient_DHParamsAnswer) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *SetClient_DHParamsAnswer) GetData2() *SetClient_DHParamsAnswer_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// dh_gen_ok#3bcbf734 nonce:int128 server_nonce:int128 new_nonce_hash1:int128 = Set_client_DH_params_answer;
type TLDhGenOk struct {
	Data2 *SetClient_DHParamsAnswer_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDhGenOk) Reset()                    { *m = TLDhGenOk{} }
func (m *TLDhGenOk) String() string            { return proto.CompactTextString(m) }
func (*TLDhGenOk) ProtoMessage()               {}
func (*TLDhGenOk) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *TLDhGenOk) GetData2() *SetClient_DHParamsAnswer_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// dh_gen_retry#46dc1fb9 nonce:int128 server_nonce:int128 new_nonce_hash2:int128 = Set_client_DH_params_answer;
type TLDhGenRetry struct {
	Data2 *SetClient_DHParamsAnswer_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDhGenRetry) Reset()                    { *m = TLDhGenRetry{} }
func (m *TLDhGenRetry) String() string            { return proto.CompactTextString(m) }
func (*TLDhGenRetry) ProtoMessage()               {}
func (*TLDhGenRetry) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *TLDhGenRetry) GetData2() *SetClient_DHParamsAnswer_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// dh_gen_fail#a69dae02 nonce:int128 server_nonce:int128 new_nonce_hash3:int128 = Set_client_DH_params_answer;
type TLDhGenFail struct {
	Data2 *SetClient_DHParamsAnswer_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDhGenFail) Reset()                    { *m = TLDhGenFail{} }
func (m *TLDhGenFail) String() string            { return proto.CompactTextString(m) }
func (*TLDhGenFail) ProtoMessage()               {}
func (*TLDhGenFail) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *TLDhGenFail) GetData2() *SetClient_DHParamsAnswer_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// DestroyAuthKeyRes <--
//  + TL_destroy_auth_key_ok
//  + TL_destroy_auth_key_none
//  + TL_destroy_auth_key_fail
//
type DestroyAuthKeyRes_Data struct {
}

func (m *DestroyAuthKeyRes_Data) Reset()                    { *m = DestroyAuthKeyRes_Data{} }
func (m *DestroyAuthKeyRes_Data) String() string            { return proto.CompactTextString(m) }
func (*DestroyAuthKeyRes_Data) ProtoMessage()               {}
func (*DestroyAuthKeyRes_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

type DestroyAuthKeyRes struct {
	Constructor TLConstructor           `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *DestroyAuthKeyRes_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *DestroyAuthKeyRes) Reset()                    { *m = DestroyAuthKeyRes{} }
func (m *DestroyAuthKeyRes) String() string            { return proto.CompactTextString(m) }
func (*DestroyAuthKeyRes) ProtoMessage()               {}
func (*DestroyAuthKeyRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *DestroyAuthKeyRes) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *DestroyAuthKeyRes) GetData2() *DestroyAuthKeyRes_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
type TLDestroyAuthKeyOk struct {
	Data2 *DestroyAuthKeyRes_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDestroyAuthKeyOk) Reset()                    { *m = TLDestroyAuthKeyOk{} }
func (m *TLDestroyAuthKeyOk) String() string            { return proto.CompactTextString(m) }
func (*TLDestroyAuthKeyOk) ProtoMessage()               {}
func (*TLDestroyAuthKeyOk) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *TLDestroyAuthKeyOk) GetData2() *DestroyAuthKeyRes_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
type TLDestroyAuthKeyNone struct {
	Data2 *DestroyAuthKeyRes_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDestroyAuthKeyNone) Reset()                    { *m = TLDestroyAuthKeyNone{} }
func (m *TLDestroyAuthKeyNone) String() string            { return proto.CompactTextString(m) }
func (*TLDestroyAuthKeyNone) ProtoMessage()               {}
func (*TLDestroyAuthKeyNone) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{14} }

func (m *TLDestroyAuthKeyNone) GetData2() *DestroyAuthKeyRes_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
type TLDestroyAuthKeyFail struct {
	Data2 *DestroyAuthKeyRes_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLDestroyAuthKeyFail) Reset()                    { *m = TLDestroyAuthKeyFail{} }
func (m *TLDestroyAuthKeyFail) String() string            { return proto.CompactTextString(m) }
func (*TLDestroyAuthKeyFail) ProtoMessage()               {}
func (*TLDestroyAuthKeyFail) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{15} }

func (m *TLDestroyAuthKeyFail) GetData2() *DestroyAuthKeyRes_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// ResPQ <--
//  + TL_resPQ
//
type ResPQ_Data struct {
	Nonce                       []byte  `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce                 []byte  `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	Pq                          string  `protobuf:"bytes,3,opt,name=pq" json:"pq,omitempty"`
	ServerPublicKeyFingerprints []int64 `protobuf:"varint,4,rep,packed,name=server_public_key_fingerprints,json=serverPublicKeyFingerprints" json:"server_public_key_fingerprints,omitempty"`
}

func (m *ResPQ_Data) Reset()                    { *m = ResPQ_Data{} }
func (m *ResPQ_Data) String() string            { return proto.CompactTextString(m) }
func (*ResPQ_Data) ProtoMessage()               {}
func (*ResPQ_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{16} }

func (m *ResPQ_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *ResPQ_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *ResPQ_Data) GetPq() string {
	if m != nil {
		return m.Pq
	}
	return ""
}

func (m *ResPQ_Data) GetServerPublicKeyFingerprints() []int64 {
	if m != nil {
		return m.ServerPublicKeyFingerprints
	}
	return nil
}

type ResPQ struct {
	Constructor TLConstructor `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *ResPQ_Data   `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *ResPQ) Reset()                    { *m = ResPQ{} }
func (m *ResPQ) String() string            { return proto.CompactTextString(m) }
func (*ResPQ) ProtoMessage()               {}
func (*ResPQ) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{17} }

func (m *ResPQ) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *ResPQ) GetData2() *ResPQ_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
type TLResPQ struct {
	Data2 *ResPQ_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLResPQ) Reset()                    { *m = TLResPQ{} }
func (m *TLResPQ) String() string            { return proto.CompactTextString(m) }
func (*TLResPQ) ProtoMessage()               {}
func (*TLResPQ) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{18} }

func (m *TLResPQ) GetData2() *ResPQ_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// P_Q_inner_data <--
//  + TL_p_q_inner_data
//
type P_QInnerData_Data struct {
	Pq          string `protobuf:"bytes,1,opt,name=pq" json:"pq,omitempty"`
	P           string `protobuf:"bytes,2,opt,name=p" json:"p,omitempty"`
	Q           string `protobuf:"bytes,3,opt,name=q" json:"q,omitempty"`
	Nonce       []byte `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce []byte `protobuf:"bytes,5,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	NewNonce    []byte `protobuf:"bytes,6,opt,name=new_nonce,json=newNonce,proto3" json:"new_nonce,omitempty"`
}

func (m *P_QInnerData_Data) Reset()                    { *m = P_QInnerData_Data{} }
func (m *P_QInnerData_Data) String() string            { return proto.CompactTextString(m) }
func (*P_QInnerData_Data) ProtoMessage()               {}
func (*P_QInnerData_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{19} }

func (m *P_QInnerData_Data) GetPq() string {
	if m != nil {
		return m.Pq
	}
	return ""
}

func (m *P_QInnerData_Data) GetP() string {
	if m != nil {
		return m.P
	}
	return ""
}

func (m *P_QInnerData_Data) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *P_QInnerData_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *P_QInnerData_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *P_QInnerData_Data) GetNewNonce() []byte {
	if m != nil {
		return m.NewNonce
	}
	return nil
}

type P_QInnerData struct {
	Constructor TLConstructor      `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *P_QInnerData_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *P_QInnerData) Reset()                    { *m = P_QInnerData{} }
func (m *P_QInnerData) String() string            { return proto.CompactTextString(m) }
func (*P_QInnerData) ProtoMessage()               {}
func (*P_QInnerData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{20} }

func (m *P_QInnerData) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *P_QInnerData) GetData2() *P_QInnerData_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// p_q_inner_data#83c95aec pq:string p:string q:string nonce:int128 server_nonce:int128 new_nonce:int256 = P_Q_inner_data;
type TLPQInnerData struct {
	Data2 *P_QInnerData_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLPQInnerData) Reset()                    { *m = TLPQInnerData{} }
func (m *TLPQInnerData) String() string            { return proto.CompactTextString(m) }
func (*TLPQInnerData) ProtoMessage()               {}
func (*TLPQInnerData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{21} }

func (m *TLPQInnerData) GetData2() *P_QInnerData_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Server_DH_Params <--
//  + TL_server_DH_params_fail
//  + TL_server_DH_params_ok
//
type Server_DH_Params_Data struct {
	Nonce           []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce     []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	NewNonceHash    []byte `protobuf:"bytes,3,opt,name=new_nonce_hash,json=newNonceHash,proto3" json:"new_nonce_hash,omitempty"`
	EncryptedAnswer string `protobuf:"bytes,4,opt,name=encrypted_answer,json=encryptedAnswer" json:"encrypted_answer,omitempty"`
}

func (m *Server_DH_Params_Data) Reset()                    { *m = Server_DH_Params_Data{} }
func (m *Server_DH_Params_Data) String() string            { return proto.CompactTextString(m) }
func (*Server_DH_Params_Data) ProtoMessage()               {}
func (*Server_DH_Params_Data) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{22} }

func (m *Server_DH_Params_Data) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Server_DH_Params_Data) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *Server_DH_Params_Data) GetNewNonceHash() []byte {
	if m != nil {
		return m.NewNonceHash
	}
	return nil
}

func (m *Server_DH_Params_Data) GetEncryptedAnswer() string {
	if m != nil {
		return m.EncryptedAnswer
	}
	return ""
}

type Server_DH_Params struct {
	Constructor TLConstructor          `protobuf:"varint,1,opt,name=constructor,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data2       *Server_DH_Params_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *Server_DH_Params) Reset()                    { *m = Server_DH_Params{} }
func (m *Server_DH_Params) String() string            { return proto.CompactTextString(m) }
func (*Server_DH_Params) ProtoMessage()               {}
func (*Server_DH_Params) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{23} }

func (m *Server_DH_Params) GetConstructor() TLConstructor {
	if m != nil {
		return m.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (m *Server_DH_Params) GetData2() *Server_DH_Params_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// server_DH_params_fail#79cb045d nonce:int128 server_nonce:int128 new_nonce_hash:int128 = Server_DH_Params;
type TLServer_DHParamsFail struct {
	Data2 *Server_DH_Params_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLServer_DHParamsFail) Reset()                    { *m = TLServer_DHParamsFail{} }
func (m *TLServer_DHParamsFail) String() string            { return proto.CompactTextString(m) }
func (*TLServer_DHParamsFail) ProtoMessage()               {}
func (*TLServer_DHParamsFail) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{24} }

func (m *TLServer_DHParamsFail) GetData2() *Server_DH_Params_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

// server_DH_params_ok#d0e8075c nonce:int128 server_nonce:int128 encrypted_answer:string = Server_DH_Params;
type TLServer_DHParamsOk struct {
	Data2 *Server_DH_Params_Data `protobuf:"bytes,2,opt,name=data2" json:"data2,omitempty"`
}

func (m *TLServer_DHParamsOk) Reset()                    { *m = TLServer_DHParamsOk{} }
func (m *TLServer_DHParamsOk) String() string            { return proto.CompactTextString(m) }
func (*TLServer_DHParamsOk) ProtoMessage()               {}
func (*TLServer_DHParamsOk) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{25} }

func (m *TLServer_DHParamsOk) GetData2() *Server_DH_Params_Data {
	if m != nil {
		return m.Data2
	}
	return nil
}

func init() {
	proto.RegisterType((*Server_DHInnerData_Data)(nil), "mtproto.Server_DH_inner_data_Data")
	proto.RegisterType((*Server_DHInnerData)(nil), "mtproto.Server_DH_inner_data")
	proto.RegisterType((*TLServer_DHInnerData)(nil), "mtproto.TL_server_DH_inner_data")
	proto.RegisterType((*Client_DH_Inner_Data_Data)(nil), "mtproto.Client_DH_Inner_Data_Data")
	proto.RegisterType((*Client_DH_Inner_Data)(nil), "mtproto.Client_DH_Inner_Data")
	proto.RegisterType((*TLClient_DHInnerData)(nil), "mtproto.TL_client_DH_inner_data")
	proto.RegisterType((*SetClient_DHParamsAnswer_Data)(nil), "mtproto.Set_client_DH_params_answer_Data")
	proto.RegisterType((*SetClient_DHParamsAnswer)(nil), "mtproto.Set_client_DH_params_answer")
	proto.RegisterType((*TLDhGenOk)(nil), "mtproto.TL_dh_gen_ok")
	proto.RegisterType((*TLDhGenRetry)(nil), "mtproto.TL_dh_gen_retry")
	proto.RegisterType((*TLDhGenFail)(nil), "mtproto.TL_dh_gen_fail")
	proto.RegisterType((*DestroyAuthKeyRes_Data)(nil), "mtproto.DestroyAuthKeyRes_Data")
	proto.RegisterType((*DestroyAuthKeyRes)(nil), "mtproto.DestroyAuthKeyRes")
	proto.RegisterType((*TLDestroyAuthKeyOk)(nil), "mtproto.TL_destroy_auth_key_ok")
	proto.RegisterType((*TLDestroyAuthKeyNone)(nil), "mtproto.TL_destroy_auth_key_none")
	proto.RegisterType((*TLDestroyAuthKeyFail)(nil), "mtproto.TL_destroy_auth_key_fail")
	proto.RegisterType((*ResPQ_Data)(nil), "mtproto.ResPQ_Data")
	proto.RegisterType((*ResPQ)(nil), "mtproto.ResPQ")
	proto.RegisterType((*TLResPQ)(nil), "mtproto.TL_resPQ")
	proto.RegisterType((*P_QInnerData_Data)(nil), "mtproto.P_Q_inner_data_Data")
	proto.RegisterType((*P_QInnerData)(nil), "mtproto.P_Q_inner_data")
	proto.RegisterType((*TLPQInnerData)(nil), "mtproto.TL_p_q_inner_data")
	proto.RegisterType((*Server_DH_Params_Data)(nil), "mtproto.Server_DH_Params_Data")
	proto.RegisterType((*Server_DH_Params)(nil), "mtproto.Server_DH_Params")
	proto.RegisterType((*TLServer_DHParamsFail)(nil), "mtproto.TL_server_DH_params_fail")
	proto.RegisterType((*TLServer_DHParamsOk)(nil), "mtproto.TL_server_DH_params_ok")
}

func init() { proto.RegisterFile("schema.tl.handshake.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xcd, 0x6e, 0xea, 0x46,
	0x14, 0xc7, 0x35, 0x10, 0x02, 0x9c, 0xb8, 0x90, 0x38, 0x1f, 0x85, 0x52, 0x25, 0xd4, 0xaa, 0x2a,
	0xb2, 0x41, 0x2a, 0x34, 0x52, 0x76, 0x55, 0x12, 0xd4, 0x26, 0x0a, 0x4d, 0xc1, 0x61, 0x3f, 0x32,
	0xf6, 0xd4, 0xb6, 0x80, 0xb1, 0x19, 0x0f, 0x8d, 0xd8, 0x54, 0x55, 0xd5, 0x55, 0x5f, 0xa0, 0x55,
	0x77, 0x7d, 0x82, 0x3e, 0xcc, 0x7d, 0xa1, 0x2b, 0x8f, 0xbf, 0x70, 0x20, 0xb9, 0xf7, 0xe2, 0x7b,
	0x77, 0x9e, 0x99, 0xff, 0xf9, 0xfb, 0xcc, 0xef, 0x8c, 0x8f, 0x07, 0xea, 0x9e, 0x6e, 0x91, 0x99,
	0xd6, 0xe6, 0xd3, 0xb6, 0xa5, 0x51, 0xc3, 0xb3, 0xb4, 0x09, 0x69, 0xbb, 0xcc, 0xe1, 0x8e, 0x5c,
	0x9c, 0x71, 0xf1, 0xf0, 0xc5, 0x71, 0xa2, 0xd1, 0x99, 0xde, 0xed, 0x04, 0xeb, 0xca, 0xff, 0x08,
	0xea, 0x8f, 0x84, 0xfd, 0x4a, 0x18, 0xee, 0xdd, 0x62, 0x9b, 0x52, 0xc2, 0xb0, 0xa1, 0x71, 0x0d,
	0xf7, 0x34, 0xae, 0xc9, 0x47, 0x50, 0xa0, 0x0e, 0xd5, 0x49, 0x0d, 0x35, 0x51, 0x4b, 0x52, 0x83,
	0x81, 0xfc, 0x15, 0x48, 0x5e, 0x10, 0x12, 0x2c, 0xe6, 0xc4, 0xe2, 0x5e, 0x30, 0xf7, 0x20, 0x24,
	0x12, 0x20, 0xb3, 0x96, 0x6f, 0xa2, 0x56, 0x41, 0x45, 0xa6, 0x5c, 0x87, 0x92, 0x61, 0x61, 0x97,
	0xd9, 0x33, 0x52, 0xdb, 0x69, 0xa2, 0x56, 0x59, 0x2d, 0x1a, 0xd6, 0xc0, 0x1f, 0xca, 0x55, 0xc8,
	0x9b, 0x58, 0xab, 0x15, 0xc4, 0x6c, 0xce, 0xbc, 0x92, 0xcf, 0x20, 0x34, 0xc2, 0xdc, 0x97, 0xef,
	0x0a, 0x0f, 0x08, 0xa6, 0x46, 0xf6, 0x8c, 0x28, 0x7f, 0x21, 0x38, 0xda, 0x94, 0xb1, 0x7c, 0x09,
	0x7b, 0xba, 0x43, 0x3d, 0xce, 0x16, 0x3a, 0x77, 0x98, 0x48, 0xb9, 0xd2, 0x39, 0x69, 0x87, 0x00,
	0xda, 0xa3, 0xfe, 0x4d, 0xb2, 0xaa, 0xae, 0x4a, 0xe5, 0x4b, 0x28, 0xf8, 0x0e, 0x1d, 0xb1, 0x93,
	0xbd, 0x8e, 0x12, 0xc7, 0xbc, 0x48, 0x46, 0x0d, 0x02, 0x94, 0x47, 0xf8, 0x7c, 0xd4, 0xc7, 0xde,
	0xe6, 0x74, 0xb6, 0x35, 0xfd, 0x1d, 0x41, 0xfd, 0x66, 0x6a, 0x13, 0xca, 0x7d, 0xd1, 0x9d, 0x10,
	0xf5, 0xb2, 0xd7, 0xa4, 0x0e, 0x25, 0x46, 0x38, 0x5b, 0x62, 0xdb, 0x10, 0xa5, 0xc9, 0xab, 0x45,
	0x31, 0xbe, 0x33, 0x82, 0x2a, 0x8c, 0xc3, 0xda, 0xe4, 0xcc, 0x6b, 0x01, 0x79, 0x53, 0x0a, 0x9f,
	0x02, 0xf2, 0x8b, 0x5b, 0x4d, 0x43, 0xd6, 0x63, 0xd9, 0xfb, 0x40, 0x7e, 0xa7, 0xe9, 0x1b, 0x04,
	0xcd, 0x47, 0xc2, 0x57, 0x6c, 0x5d, 0x8d, 0x69, 0x33, 0x0f, 0x6b, 0xd4, 0x7b, 0x8a, 0x76, 0xbb,
	0x35, 0xeb, 0x6f, 0xa0, 0x4a, 0xc9, 0x53, 0xb0, 0x8e, 0x2d, 0xcd, 0xb3, 0xbe, 0x15, 0xc8, 0x25,
	0xf5, 0x33, 0x4a, 0x9e, 0x84, 0xe4, 0xd6, 0x9f, 0x5c, 0xd7, 0x75, 0x44, 0x11, 0x9e, 0xe9, 0x3a,
	0xeb, 0xba, 0xae, 0xf8, 0x64, 0x9e, 0xe9, 0xba, 0xca, 0x3f, 0x08, 0x1a, 0xaf, 0xec, 0x2a, 0x43,
	0xf9, 0xbe, 0x4f, 0x93, 0x3e, 0x5f, 0x39, 0xce, 0xaf, 0x43, 0x8c, 0x80, 0xff, 0x0c, 0xd2, 0xa8,
	0x8f, 0x0d, 0x0b, 0x9b, 0x84, 0x62, 0x67, 0x92, 0xdd, 0x50, 0x85, 0x6a, 0x62, 0x28, 0x4e, 0x72,
	0x76, 0xcf, 0x21, 0x54, 0x12, 0xcf, 0x5f, 0x34, 0x7b, 0x9a, 0xdd, 0xb2, 0x06, 0x27, 0x3d, 0xe2,
	0x71, 0xe6, 0x2c, 0xaf, 0x16, 0xdc, 0xba, 0x27, 0x4b, 0x95, 0x78, 0x42, 0xa0, 0xfc, 0x89, 0xe0,
	0x60, 0x6d, 0x29, 0x43, 0x89, 0x2e, 0xd2, 0xa9, 0x9e, 0xc5, 0x31, 0x9b, 0xdf, 0x9f, 0x14, 0xe6,
	0xc4, 0xdf, 0x73, 0xa0, 0xc1, 0xda, 0x82, 0x5b, 0x78, 0x42, 0x96, 0x7e, 0x89, 0xb6, 0x34, 0x1c,
	0x42, 0x6d, 0x93, 0x21, 0x75, 0x28, 0xf9, 0xc8, 0x96, 0xa2, 0x42, 0x5b, 0x5a, 0xfe, 0x8b, 0x00,
	0x54, 0xe2, 0x0d, 0x86, 0x19, 0x3f, 0xf5, 0x0a, 0xe4, 0xdc, 0xb9, 0xf8, 0xba, 0xcb, 0x6a, 0xce,
	0x9d, 0xcb, 0x37, 0x70, 0x1a, 0x86, 0xb8, 0x8b, 0xf1, 0xd4, 0xd6, 0x83, 0x44, 0x6d, 0x6a, 0x12,
	0xe6, 0x32, 0x9b, 0x72, 0xaf, 0xb6, 0xd3, 0xcc, 0xb7, 0xf2, 0x6a, 0x23, 0x50, 0x0d, 0x84, 0xe8,
	0x9e, 0x2c, 0x7f, 0x58, 0x91, 0x28, 0x53, 0x28, 0x88, 0xdc, 0x32, 0x9c, 0x86, 0xf3, 0x34, 0x96,
	0xc3, 0x38, 0x26, 0xd9, 0x74, 0x84, 0xe2, 0x02, 0x4a, 0xa3, 0x3e, 0x66, 0xe2, 0x85, 0x1f, 0x10,
	0xf6, 0x37, 0x82, 0xc3, 0x01, 0x1e, 0xae, 0xdd, 0x1a, 0x02, 0x22, 0x28, 0x26, 0x22, 0x01, 0x72,
	0x85, 0x5d, 0x59, 0x45, 0xae, 0x3f, 0x8a, 0x70, 0xa1, 0x79, 0x82, 0x7d, 0xe7, 0x35, 0xec, 0x85,
	0x75, 0xec, 0x0d, 0x28, 0xc7, 0x1d, 0x51, 0xdc, 0x12, 0x24, 0xb5, 0x14, 0xf5, 0x42, 0xe5, 0x37,
	0xa8, 0xa4, 0x13, 0xcb, 0xc0, 0xb1, 0x93, 0x06, 0xf2, 0x65, 0x1c, 0xb3, 0x61, 0xeb, 0x11, 0x99,
	0x1f, 0xe1, 0x60, 0xd4, 0xc7, 0x2e, 0x9e, 0xaf, 0xa6, 0xb0, 0x8d, 0xd1, 0x7f, 0x08, 0x8e, 0x93,
	0xfb, 0xc2, 0x20, 0x68, 0x32, 0xd9, 0xce, 0xeb, 0xd7, 0x50, 0x49, 0xff, 0x4a, 0xc2, 0x3f, 0x93,
	0xb4, 0xfa, 0x27, 0x91, 0xcf, 0x61, 0x9f, 0x50, 0x9d, 0x2d, 0x5d, 0x4e, 0x8c, 0xb0, 0xab, 0x85,
	0xd7, 0x83, 0x6a, 0x3c, 0x7f, 0x25, 0xa6, 0x95, 0x3f, 0x10, 0xec, 0x3f, 0xcf, 0x31, 0x03, 0xef,
	0xef, 0xd2, 0x98, 0x4e, 0x37, 0xdc, 0x9b, 0x56, 0x38, 0x44, 0xa0, 0x06, 0xa2, 0x41, 0x24, 0x17,
	0xb1, 0xb0, 0x1f, 0x8b, 0x06, 0xb1, 0x9d, 0xe3, 0x83, 0x68, 0x8b, 0x6b, 0x8e, 0xce, 0x64, 0x3b,
	0xbf, 0xeb, 0x16, 0x34, 0x74, 0x67, 0xd6, 0xa6, 0x64, 0xbc, 0x98, 0x6a, 0xf6, 0xac, 0x4d, 0xa8,
	0x69, 0x53, 0x12, 0xc5, 0x5e, 0x17, 0x7f, 0x1a, 0x0d, 0xfc, 0x87, 0xdb, 0xdc, 0x78, 0x57, 0xcc,
	0x74, 0xdf, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x79, 0x7a, 0x63, 0xd7, 0x0b, 0x00, 0x00,
}
