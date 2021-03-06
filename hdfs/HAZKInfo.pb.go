// Code generated by protoc-gen-go.
// source: HAZKInfo.proto
// DO NOT EDIT!

package hadoop_hdfs

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ActiveNodeInfo struct {
	NameserviceId    *string `protobuf:"bytes,1,req,name=nameserviceId" json:"nameserviceId,omitempty"`
	NamenodeId       *string `protobuf:"bytes,2,req,name=namenodeId" json:"namenodeId,omitempty"`
	Hostname         *string `protobuf:"bytes,3,req,name=hostname" json:"hostname,omitempty"`
	Port             *int32  `protobuf:"varint,4,req,name=port" json:"port,omitempty"`
	ZkfcPort         *int32  `protobuf:"varint,5,req,name=zkfcPort" json:"zkfcPort,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ActiveNodeInfo) Reset()         { *m = ActiveNodeInfo{} }
func (m *ActiveNodeInfo) String() string { return proto.CompactTextString(m) }
func (*ActiveNodeInfo) ProtoMessage()    {}

func (m *ActiveNodeInfo) GetNameserviceId() string {
	if m != nil && m.NameserviceId != nil {
		return *m.NameserviceId
	}
	return ""
}

func (m *ActiveNodeInfo) GetNamenodeId() string {
	if m != nil && m.NamenodeId != nil {
		return *m.NamenodeId
	}
	return ""
}

func (m *ActiveNodeInfo) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ActiveNodeInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ActiveNodeInfo) GetZkfcPort() int32 {
	if m != nil && m.ZkfcPort != nil {
		return *m.ZkfcPort
	}
	return 0
}

func init() {
}
