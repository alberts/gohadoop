// Code generated by protoc-gen-go.
// source: xattr.proto
// DO NOT EDIT!

package hadoop_hdfs

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type XAttrSetFlagProto int32

const (
	XAttrSetFlagProto_XATTR_CREATE  XAttrSetFlagProto = 1
	XAttrSetFlagProto_XATTR_REPLACE XAttrSetFlagProto = 2
)

var XAttrSetFlagProto_name = map[int32]string{
	1: "XATTR_CREATE",
	2: "XATTR_REPLACE",
}
var XAttrSetFlagProto_value = map[string]int32{
	"XATTR_CREATE":  1,
	"XATTR_REPLACE": 2,
}

func (x XAttrSetFlagProto) Enum() *XAttrSetFlagProto {
	p := new(XAttrSetFlagProto)
	*p = x
	return p
}
func (x XAttrSetFlagProto) String() string {
	return proto.EnumName(XAttrSetFlagProto_name, int32(x))
}
func (x *XAttrSetFlagProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XAttrSetFlagProto_value, data, "XAttrSetFlagProto")
	if err != nil {
		return err
	}
	*x = XAttrSetFlagProto(value)
	return nil
}

type XAttrProto_XAttrNamespaceProto int32

const (
	XAttrProto_USER     XAttrProto_XAttrNamespaceProto = 0
	XAttrProto_TRUSTED  XAttrProto_XAttrNamespaceProto = 1
	XAttrProto_SECURITY XAttrProto_XAttrNamespaceProto = 2
	XAttrProto_SYSTEM   XAttrProto_XAttrNamespaceProto = 3
)

var XAttrProto_XAttrNamespaceProto_name = map[int32]string{
	0: "USER",
	1: "TRUSTED",
	2: "SECURITY",
	3: "SYSTEM",
}
var XAttrProto_XAttrNamespaceProto_value = map[string]int32{
	"USER":     0,
	"TRUSTED":  1,
	"SECURITY": 2,
	"SYSTEM":   3,
}

func (x XAttrProto_XAttrNamespaceProto) Enum() *XAttrProto_XAttrNamespaceProto {
	p := new(XAttrProto_XAttrNamespaceProto)
	*p = x
	return p
}
func (x XAttrProto_XAttrNamespaceProto) String() string {
	return proto.EnumName(XAttrProto_XAttrNamespaceProto_name, int32(x))
}
func (x *XAttrProto_XAttrNamespaceProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(XAttrProto_XAttrNamespaceProto_value, data, "XAttrProto_XAttrNamespaceProto")
	if err != nil {
		return err
	}
	*x = XAttrProto_XAttrNamespaceProto(value)
	return nil
}

type XAttrProto struct {
	Namespace        *XAttrProto_XAttrNamespaceProto `protobuf:"varint,1,req,name=namespace,enum=hadoop.hdfs.XAttrProto_XAttrNamespaceProto" json:"namespace,omitempty"`
	Name             *string                         `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Value            []byte                          `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *XAttrProto) Reset()         { *m = XAttrProto{} }
func (m *XAttrProto) String() string { return proto.CompactTextString(m) }
func (*XAttrProto) ProtoMessage()    {}

func (m *XAttrProto) GetNamespace() XAttrProto_XAttrNamespaceProto {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return XAttrProto_USER
}

func (m *XAttrProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *XAttrProto) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type XAttrEditLogProto struct {
	Src              *string       `protobuf:"bytes,1,opt,name=src" json:"src,omitempty"`
	XAttrs           []*XAttrProto `protobuf:"bytes,2,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *XAttrEditLogProto) Reset()         { *m = XAttrEditLogProto{} }
func (m *XAttrEditLogProto) String() string { return proto.CompactTextString(m) }
func (*XAttrEditLogProto) ProtoMessage()    {}

func (m *XAttrEditLogProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *XAttrEditLogProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

type SetXAttrRequestProto struct {
	Src              *string     `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttr            *XAttrProto `protobuf:"bytes,2,opt,name=xAttr" json:"xAttr,omitempty"`
	Flag             *uint32     `protobuf:"varint,3,opt,name=flag" json:"flag,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SetXAttrRequestProto) Reset()         { *m = SetXAttrRequestProto{} }
func (m *SetXAttrRequestProto) String() string { return proto.CompactTextString(m) }
func (*SetXAttrRequestProto) ProtoMessage()    {}

func (m *SetXAttrRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *SetXAttrRequestProto) GetXAttr() *XAttrProto {
	if m != nil {
		return m.XAttr
	}
	return nil
}

func (m *SetXAttrRequestProto) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

type SetXAttrResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetXAttrResponseProto) Reset()         { *m = SetXAttrResponseProto{} }
func (m *SetXAttrResponseProto) String() string { return proto.CompactTextString(m) }
func (*SetXAttrResponseProto) ProtoMessage()    {}

type GetXAttrsRequestProto struct {
	Src              *string       `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttrs           []*XAttrProto `protobuf:"bytes,2,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetXAttrsRequestProto) Reset()         { *m = GetXAttrsRequestProto{} }
func (m *GetXAttrsRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetXAttrsRequestProto) ProtoMessage()    {}

func (m *GetXAttrsRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *GetXAttrsRequestProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

type GetXAttrsResponseProto struct {
	XAttrs           []*XAttrProto `protobuf:"bytes,1,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetXAttrsResponseProto) Reset()         { *m = GetXAttrsResponseProto{} }
func (m *GetXAttrsResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetXAttrsResponseProto) ProtoMessage()    {}

func (m *GetXAttrsResponseProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

type ListXAttrsRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListXAttrsRequestProto) Reset()         { *m = ListXAttrsRequestProto{} }
func (m *ListXAttrsRequestProto) String() string { return proto.CompactTextString(m) }
func (*ListXAttrsRequestProto) ProtoMessage()    {}

func (m *ListXAttrsRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type ListXAttrsResponseProto struct {
	XAttrs           []*XAttrProto `protobuf:"bytes,1,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ListXAttrsResponseProto) Reset()         { *m = ListXAttrsResponseProto{} }
func (m *ListXAttrsResponseProto) String() string { return proto.CompactTextString(m) }
func (*ListXAttrsResponseProto) ProtoMessage()    {}

func (m *ListXAttrsResponseProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

type RemoveXAttrRequestProto struct {
	Src              *string     `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttr            *XAttrProto `protobuf:"bytes,2,opt,name=xAttr" json:"xAttr,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RemoveXAttrRequestProto) Reset()         { *m = RemoveXAttrRequestProto{} }
func (m *RemoveXAttrRequestProto) String() string { return proto.CompactTextString(m) }
func (*RemoveXAttrRequestProto) ProtoMessage()    {}

func (m *RemoveXAttrRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *RemoveXAttrRequestProto) GetXAttr() *XAttrProto {
	if m != nil {
		return m.XAttr
	}
	return nil
}

type RemoveXAttrResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveXAttrResponseProto) Reset()         { *m = RemoveXAttrResponseProto{} }
func (m *RemoveXAttrResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveXAttrResponseProto) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("hadoop.hdfs.XAttrSetFlagProto", XAttrSetFlagProto_name, XAttrSetFlagProto_value)
	proto.RegisterEnum("hadoop.hdfs.XAttrProto_XAttrNamespaceProto", XAttrProto_XAttrNamespaceProto_name, XAttrProto_XAttrNamespaceProto_value)
}
