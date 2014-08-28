// Code generated by protoc-gen-go.
// source: datatransfer.proto
// DO NOT EDIT!

package hadoop_hdfs

import proto "code.google.com/p/goprotobuf/proto"
import math "math"
import hadoop_common "github.com/hortonworks/gohadoop/hadoop_common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Status int32

const (
	Status_SUCCESS            Status = 0
	Status_ERROR              Status = 1
	Status_ERROR_CHECKSUM     Status = 2
	Status_ERROR_INVALID      Status = 3
	Status_ERROR_EXISTS       Status = 4
	Status_ERROR_ACCESS_TOKEN Status = 5
	Status_CHECKSUM_OK        Status = 6
	Status_ERROR_UNSUPPORTED  Status = 7
	Status_OOB_RESTART        Status = 8
	Status_OOB_RESERVED1      Status = 9
	Status_OOB_RESERVED2      Status = 10
	Status_OOB_RESERVED3      Status = 11
)

var Status_name = map[int32]string{
	0:  "SUCCESS",
	1:  "ERROR",
	2:  "ERROR_CHECKSUM",
	3:  "ERROR_INVALID",
	4:  "ERROR_EXISTS",
	5:  "ERROR_ACCESS_TOKEN",
	6:  "CHECKSUM_OK",
	7:  "ERROR_UNSUPPORTED",
	8:  "OOB_RESTART",
	9:  "OOB_RESERVED1",
	10: "OOB_RESERVED2",
	11: "OOB_RESERVED3",
}
var Status_value = map[string]int32{
	"SUCCESS":            0,
	"ERROR":              1,
	"ERROR_CHECKSUM":     2,
	"ERROR_INVALID":      3,
	"ERROR_EXISTS":       4,
	"ERROR_ACCESS_TOKEN": 5,
	"CHECKSUM_OK":        6,
	"ERROR_UNSUPPORTED":  7,
	"OOB_RESTART":        8,
	"OOB_RESERVED1":      9,
	"OOB_RESERVED2":      10,
	"OOB_RESERVED3":      11,
}

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}
func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (x *Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Status_value, data, "Status")
	if err != nil {
		return err
	}
	*x = Status(value)
	return nil
}

type DataTransferEncryptorMessageProto_DataTransferEncryptorStatus int32

const (
	DataTransferEncryptorMessageProto_SUCCESS           DataTransferEncryptorMessageProto_DataTransferEncryptorStatus = 0
	DataTransferEncryptorMessageProto_ERROR_UNKNOWN_KEY DataTransferEncryptorMessageProto_DataTransferEncryptorStatus = 1
	DataTransferEncryptorMessageProto_ERROR             DataTransferEncryptorMessageProto_DataTransferEncryptorStatus = 2
)

var DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR_UNKNOWN_KEY",
	2: "ERROR",
}
var DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_value = map[string]int32{
	"SUCCESS":           0,
	"ERROR_UNKNOWN_KEY": 1,
	"ERROR":             2,
}

func (x DataTransferEncryptorMessageProto_DataTransferEncryptorStatus) Enum() *DataTransferEncryptorMessageProto_DataTransferEncryptorStatus {
	p := new(DataTransferEncryptorMessageProto_DataTransferEncryptorStatus)
	*p = x
	return p
}
func (x DataTransferEncryptorMessageProto_DataTransferEncryptorStatus) String() string {
	return proto.EnumName(DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_name, int32(x))
}
func (x *DataTransferEncryptorMessageProto_DataTransferEncryptorStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_value, data, "DataTransferEncryptorMessageProto_DataTransferEncryptorStatus")
	if err != nil {
		return err
	}
	*x = DataTransferEncryptorMessageProto_DataTransferEncryptorStatus(value)
	return nil
}

type OpWriteBlockProto_BlockConstructionStage int32

const (
	OpWriteBlockProto_PIPELINE_SETUP_APPEND OpWriteBlockProto_BlockConstructionStage = 0
	// pipeline set up for failed PIPELINE_SETUP_APPEND recovery
	OpWriteBlockProto_PIPELINE_SETUP_APPEND_RECOVERY OpWriteBlockProto_BlockConstructionStage = 1
	// data streaming
	OpWriteBlockProto_DATA_STREAMING OpWriteBlockProto_BlockConstructionStage = 2
	// pipeline setup for failed data streaming recovery
	OpWriteBlockProto_PIPELINE_SETUP_STREAMING_RECOVERY OpWriteBlockProto_BlockConstructionStage = 3
	// close the block and pipeline
	OpWriteBlockProto_PIPELINE_CLOSE OpWriteBlockProto_BlockConstructionStage = 4
	// Recover a failed PIPELINE_CLOSE
	OpWriteBlockProto_PIPELINE_CLOSE_RECOVERY OpWriteBlockProto_BlockConstructionStage = 5
	// pipeline set up for block creation
	OpWriteBlockProto_PIPELINE_SETUP_CREATE OpWriteBlockProto_BlockConstructionStage = 6
	// transfer RBW for adding datanodes
	OpWriteBlockProto_TRANSFER_RBW OpWriteBlockProto_BlockConstructionStage = 7
	// transfer Finalized for adding datanodes
	OpWriteBlockProto_TRANSFER_FINALIZED OpWriteBlockProto_BlockConstructionStage = 8
)

var OpWriteBlockProto_BlockConstructionStage_name = map[int32]string{
	0: "PIPELINE_SETUP_APPEND",
	1: "PIPELINE_SETUP_APPEND_RECOVERY",
	2: "DATA_STREAMING",
	3: "PIPELINE_SETUP_STREAMING_RECOVERY",
	4: "PIPELINE_CLOSE",
	5: "PIPELINE_CLOSE_RECOVERY",
	6: "PIPELINE_SETUP_CREATE",
	7: "TRANSFER_RBW",
	8: "TRANSFER_FINALIZED",
}
var OpWriteBlockProto_BlockConstructionStage_value = map[string]int32{
	"PIPELINE_SETUP_APPEND":             0,
	"PIPELINE_SETUP_APPEND_RECOVERY":    1,
	"DATA_STREAMING":                    2,
	"PIPELINE_SETUP_STREAMING_RECOVERY": 3,
	"PIPELINE_CLOSE":                    4,
	"PIPELINE_CLOSE_RECOVERY":           5,
	"PIPELINE_SETUP_CREATE":             6,
	"TRANSFER_RBW":                      7,
	"TRANSFER_FINALIZED":                8,
}

func (x OpWriteBlockProto_BlockConstructionStage) Enum() *OpWriteBlockProto_BlockConstructionStage {
	p := new(OpWriteBlockProto_BlockConstructionStage)
	*p = x
	return p
}
func (x OpWriteBlockProto_BlockConstructionStage) String() string {
	return proto.EnumName(OpWriteBlockProto_BlockConstructionStage_name, int32(x))
}
func (x *OpWriteBlockProto_BlockConstructionStage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OpWriteBlockProto_BlockConstructionStage_value, data, "OpWriteBlockProto_BlockConstructionStage")
	if err != nil {
		return err
	}
	*x = OpWriteBlockProto_BlockConstructionStage(value)
	return nil
}

type DataTransferEncryptorMessageProto struct {
	Status           *DataTransferEncryptorMessageProto_DataTransferEncryptorStatus `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.DataTransferEncryptorMessageProto_DataTransferEncryptorStatus" json:"status,omitempty"`
	Payload          []byte                                                         `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	Message          *string                                                        `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte                                                         `json:"-"`
}

func (m *DataTransferEncryptorMessageProto) Reset()         { *m = DataTransferEncryptorMessageProto{} }
func (m *DataTransferEncryptorMessageProto) String() string { return proto.CompactTextString(m) }
func (*DataTransferEncryptorMessageProto) ProtoMessage()    {}

func (m *DataTransferEncryptorMessageProto) GetStatus() DataTransferEncryptorMessageProto_DataTransferEncryptorStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return DataTransferEncryptorMessageProto_SUCCESS
}

func (m *DataTransferEncryptorMessageProto) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *DataTransferEncryptorMessageProto) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type BaseHeaderProto struct {
	Block            *ExtendedBlockProto       `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	Token            *hadoop_common.TokenProto `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *BaseHeaderProto) Reset()         { *m = BaseHeaderProto{} }
func (m *BaseHeaderProto) String() string { return proto.CompactTextString(m) }
func (*BaseHeaderProto) ProtoMessage()    {}

func (m *BaseHeaderProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BaseHeaderProto) GetToken() *hadoop_common.TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type ClientOperationHeaderProto struct {
	BaseHeader       *BaseHeaderProto `protobuf:"bytes,1,req,name=baseHeader" json:"baseHeader,omitempty"`
	ClientName       *string          `protobuf:"bytes,2,req,name=clientName" json:"clientName,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ClientOperationHeaderProto) Reset()         { *m = ClientOperationHeaderProto{} }
func (m *ClientOperationHeaderProto) String() string { return proto.CompactTextString(m) }
func (*ClientOperationHeaderProto) ProtoMessage()    {}

func (m *ClientOperationHeaderProto) GetBaseHeader() *BaseHeaderProto {
	if m != nil {
		return m.BaseHeader
	}
	return nil
}

func (m *ClientOperationHeaderProto) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

type CachingStrategyProto struct {
	DropBehind       *bool  `protobuf:"varint,1,opt,name=dropBehind" json:"dropBehind,omitempty"`
	Readahead        *int64 `protobuf:"varint,2,opt,name=readahead" json:"readahead,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CachingStrategyProto) Reset()         { *m = CachingStrategyProto{} }
func (m *CachingStrategyProto) String() string { return proto.CompactTextString(m) }
func (*CachingStrategyProto) ProtoMessage()    {}

func (m *CachingStrategyProto) GetDropBehind() bool {
	if m != nil && m.DropBehind != nil {
		return *m.DropBehind
	}
	return false
}

func (m *CachingStrategyProto) GetReadahead() int64 {
	if m != nil && m.Readahead != nil {
		return *m.Readahead
	}
	return 0
}

type OpReadBlockProto struct {
	Header           *ClientOperationHeaderProto `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Offset           *uint64                     `protobuf:"varint,2,req,name=offset" json:"offset,omitempty"`
	Len              *uint64                     `protobuf:"varint,3,req,name=len" json:"len,omitempty"`
	SendChecksums    *bool                       `protobuf:"varint,4,opt,name=sendChecksums,def=1" json:"sendChecksums,omitempty"`
	CachingStrategy  *CachingStrategyProto       `protobuf:"bytes,5,opt,name=cachingStrategy" json:"cachingStrategy,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *OpReadBlockProto) Reset()         { *m = OpReadBlockProto{} }
func (m *OpReadBlockProto) String() string { return proto.CompactTextString(m) }
func (*OpReadBlockProto) ProtoMessage()    {}

const Default_OpReadBlockProto_SendChecksums bool = true

func (m *OpReadBlockProto) GetHeader() *ClientOperationHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OpReadBlockProto) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *OpReadBlockProto) GetLen() uint64 {
	if m != nil && m.Len != nil {
		return *m.Len
	}
	return 0
}

func (m *OpReadBlockProto) GetSendChecksums() bool {
	if m != nil && m.SendChecksums != nil {
		return *m.SendChecksums
	}
	return Default_OpReadBlockProto_SendChecksums
}

func (m *OpReadBlockProto) GetCachingStrategy() *CachingStrategyProto {
	if m != nil {
		return m.CachingStrategy
	}
	return nil
}

type ChecksumProto struct {
	Type             *ChecksumTypeProto `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.ChecksumTypeProto" json:"type,omitempty"`
	BytesPerChecksum *uint32            `protobuf:"varint,2,req,name=bytesPerChecksum" json:"bytesPerChecksum,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ChecksumProto) Reset()         { *m = ChecksumProto{} }
func (m *ChecksumProto) String() string { return proto.CompactTextString(m) }
func (*ChecksumProto) ProtoMessage()    {}

func (m *ChecksumProto) GetType() ChecksumTypeProto {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ChecksumTypeProto_CHECKSUM_NULL
}

func (m *ChecksumProto) GetBytesPerChecksum() uint32 {
	if m != nil && m.BytesPerChecksum != nil {
		return *m.BytesPerChecksum
	}
	return 0
}

type OpWriteBlockProto struct {
	Header                *ClientOperationHeaderProto               `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Targets               []*DatanodeInfoProto                      `protobuf:"bytes,2,rep,name=targets" json:"targets,omitempty"`
	Source                *DatanodeInfoProto                        `protobuf:"bytes,3,opt,name=source" json:"source,omitempty"`
	Stage                 *OpWriteBlockProto_BlockConstructionStage `protobuf:"varint,4,req,name=stage,enum=hadoop.hdfs.OpWriteBlockProto_BlockConstructionStage" json:"stage,omitempty"`
	PipelineSize          *uint32                                   `protobuf:"varint,5,req,name=pipelineSize" json:"pipelineSize,omitempty"`
	MinBytesRcvd          *uint64                                   `protobuf:"varint,6,req,name=minBytesRcvd" json:"minBytesRcvd,omitempty"`
	MaxBytesRcvd          *uint64                                   `protobuf:"varint,7,req,name=maxBytesRcvd" json:"maxBytesRcvd,omitempty"`
	LatestGenerationStamp *uint64                                   `protobuf:"varint,8,req,name=latestGenerationStamp" json:"latestGenerationStamp,omitempty"`
	// *
	// The requested checksum mechanism for this block write.
	RequestedChecksum *ChecksumProto        `protobuf:"bytes,9,req,name=requestedChecksum" json:"requestedChecksum,omitempty"`
	CachingStrategy   *CachingStrategyProto `protobuf:"bytes,10,opt,name=cachingStrategy" json:"cachingStrategy,omitempty"`
	XXX_unrecognized  []byte                `json:"-"`
}

func (m *OpWriteBlockProto) Reset()         { *m = OpWriteBlockProto{} }
func (m *OpWriteBlockProto) String() string { return proto.CompactTextString(m) }
func (*OpWriteBlockProto) ProtoMessage()    {}

func (m *OpWriteBlockProto) GetHeader() *ClientOperationHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OpWriteBlockProto) GetTargets() []*DatanodeInfoProto {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *OpWriteBlockProto) GetSource() *DatanodeInfoProto {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *OpWriteBlockProto) GetStage() OpWriteBlockProto_BlockConstructionStage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return OpWriteBlockProto_PIPELINE_SETUP_APPEND
}

func (m *OpWriteBlockProto) GetPipelineSize() uint32 {
	if m != nil && m.PipelineSize != nil {
		return *m.PipelineSize
	}
	return 0
}

func (m *OpWriteBlockProto) GetMinBytesRcvd() uint64 {
	if m != nil && m.MinBytesRcvd != nil {
		return *m.MinBytesRcvd
	}
	return 0
}

func (m *OpWriteBlockProto) GetMaxBytesRcvd() uint64 {
	if m != nil && m.MaxBytesRcvd != nil {
		return *m.MaxBytesRcvd
	}
	return 0
}

func (m *OpWriteBlockProto) GetLatestGenerationStamp() uint64 {
	if m != nil && m.LatestGenerationStamp != nil {
		return *m.LatestGenerationStamp
	}
	return 0
}

func (m *OpWriteBlockProto) GetRequestedChecksum() *ChecksumProto {
	if m != nil {
		return m.RequestedChecksum
	}
	return nil
}

func (m *OpWriteBlockProto) GetCachingStrategy() *CachingStrategyProto {
	if m != nil {
		return m.CachingStrategy
	}
	return nil
}

type OpTransferBlockProto struct {
	Header           *ClientOperationHeaderProto `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Targets          []*DatanodeInfoProto        `protobuf:"bytes,2,rep,name=targets" json:"targets,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *OpTransferBlockProto) Reset()         { *m = OpTransferBlockProto{} }
func (m *OpTransferBlockProto) String() string { return proto.CompactTextString(m) }
func (*OpTransferBlockProto) ProtoMessage()    {}

func (m *OpTransferBlockProto) GetHeader() *ClientOperationHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OpTransferBlockProto) GetTargets() []*DatanodeInfoProto {
	if m != nil {
		return m.Targets
	}
	return nil
}

type OpReplaceBlockProto struct {
	Header           *BaseHeaderProto   `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	DelHint          *string            `protobuf:"bytes,2,req,name=delHint" json:"delHint,omitempty"`
	Source           *DatanodeInfoProto `protobuf:"bytes,3,req,name=source" json:"source,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *OpReplaceBlockProto) Reset()         { *m = OpReplaceBlockProto{} }
func (m *OpReplaceBlockProto) String() string { return proto.CompactTextString(m) }
func (*OpReplaceBlockProto) ProtoMessage()    {}

func (m *OpReplaceBlockProto) GetHeader() *BaseHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OpReplaceBlockProto) GetDelHint() string {
	if m != nil && m.DelHint != nil {
		return *m.DelHint
	}
	return ""
}

func (m *OpReplaceBlockProto) GetSource() *DatanodeInfoProto {
	if m != nil {
		return m.Source
	}
	return nil
}

type OpCopyBlockProto struct {
	Header           *BaseHeaderProto `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *OpCopyBlockProto) Reset()         { *m = OpCopyBlockProto{} }
func (m *OpCopyBlockProto) String() string { return proto.CompactTextString(m) }
func (*OpCopyBlockProto) ProtoMessage()    {}

func (m *OpCopyBlockProto) GetHeader() *BaseHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

type OpBlockChecksumProto struct {
	Header           *BaseHeaderProto `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *OpBlockChecksumProto) Reset()         { *m = OpBlockChecksumProto{} }
func (m *OpBlockChecksumProto) String() string { return proto.CompactTextString(m) }
func (*OpBlockChecksumProto) ProtoMessage()    {}

func (m *OpBlockChecksumProto) GetHeader() *BaseHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

// *
// An ID uniquely identifying a shared memory segment.
type ShortCircuitShmIdProto struct {
	Hi               *int64 `protobuf:"varint,1,req,name=hi" json:"hi,omitempty"`
	Lo               *int64 `protobuf:"varint,2,req,name=lo" json:"lo,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShortCircuitShmIdProto) Reset()         { *m = ShortCircuitShmIdProto{} }
func (m *ShortCircuitShmIdProto) String() string { return proto.CompactTextString(m) }
func (*ShortCircuitShmIdProto) ProtoMessage()    {}

func (m *ShortCircuitShmIdProto) GetHi() int64 {
	if m != nil && m.Hi != nil {
		return *m.Hi
	}
	return 0
}

func (m *ShortCircuitShmIdProto) GetLo() int64 {
	if m != nil && m.Lo != nil {
		return *m.Lo
	}
	return 0
}

// *
// An ID uniquely identifying a slot within a shared memory segment.
type ShortCircuitShmSlotProto struct {
	ShmId            *ShortCircuitShmIdProto `protobuf:"bytes,1,req,name=shmId" json:"shmId,omitempty"`
	SlotIdx          *int32                  `protobuf:"varint,2,req,name=slotIdx" json:"slotIdx,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ShortCircuitShmSlotProto) Reset()         { *m = ShortCircuitShmSlotProto{} }
func (m *ShortCircuitShmSlotProto) String() string { return proto.CompactTextString(m) }
func (*ShortCircuitShmSlotProto) ProtoMessage()    {}

func (m *ShortCircuitShmSlotProto) GetShmId() *ShortCircuitShmIdProto {
	if m != nil {
		return m.ShmId
	}
	return nil
}

func (m *ShortCircuitShmSlotProto) GetSlotIdx() int32 {
	if m != nil && m.SlotIdx != nil {
		return *m.SlotIdx
	}
	return 0
}

type OpRequestShortCircuitAccessProto struct {
	Header *BaseHeaderProto `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	// * In order to get short-circuit access to block data, clients must set this
	// to the highest version of the block data that they can understand.
	// Currently 1 is the only version, but more versions may exist in the future
	// if the on-disk format changes.
	MaxVersion *uint32 `protobuf:"varint,2,req,name=maxVersion" json:"maxVersion,omitempty"`
	// *
	// The shared memory slot to use, if we are using one.
	SlotId           *ShortCircuitShmSlotProto `protobuf:"bytes,3,opt,name=slotId" json:"slotId,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *OpRequestShortCircuitAccessProto) Reset()         { *m = OpRequestShortCircuitAccessProto{} }
func (m *OpRequestShortCircuitAccessProto) String() string { return proto.CompactTextString(m) }
func (*OpRequestShortCircuitAccessProto) ProtoMessage()    {}

func (m *OpRequestShortCircuitAccessProto) GetHeader() *BaseHeaderProto {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *OpRequestShortCircuitAccessProto) GetMaxVersion() uint32 {
	if m != nil && m.MaxVersion != nil {
		return *m.MaxVersion
	}
	return 0
}

func (m *OpRequestShortCircuitAccessProto) GetSlotId() *ShortCircuitShmSlotProto {
	if m != nil {
		return m.SlotId
	}
	return nil
}

type ReleaseShortCircuitAccessRequestProto struct {
	SlotId           *ShortCircuitShmSlotProto `protobuf:"bytes,1,req,name=slotId" json:"slotId,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ReleaseShortCircuitAccessRequestProto) Reset()         { *m = ReleaseShortCircuitAccessRequestProto{} }
func (m *ReleaseShortCircuitAccessRequestProto) String() string { return proto.CompactTextString(m) }
func (*ReleaseShortCircuitAccessRequestProto) ProtoMessage()    {}

func (m *ReleaseShortCircuitAccessRequestProto) GetSlotId() *ShortCircuitShmSlotProto {
	if m != nil {
		return m.SlotId
	}
	return nil
}

type ReleaseShortCircuitAccessResponseProto struct {
	Status           *Status `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	Error            *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReleaseShortCircuitAccessResponseProto) Reset() {
	*m = ReleaseShortCircuitAccessResponseProto{}
}
func (m *ReleaseShortCircuitAccessResponseProto) String() string { return proto.CompactTextString(m) }
func (*ReleaseShortCircuitAccessResponseProto) ProtoMessage()    {}

func (m *ReleaseShortCircuitAccessResponseProto) GetStatus() Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Status_SUCCESS
}

func (m *ReleaseShortCircuitAccessResponseProto) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type ShortCircuitShmRequestProto struct {
	// The name of the client requesting the shared memory segment.  This is
	// purely for logging / debugging purposes.
	ClientName       *string `protobuf:"bytes,1,req,name=clientName" json:"clientName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ShortCircuitShmRequestProto) Reset()         { *m = ShortCircuitShmRequestProto{} }
func (m *ShortCircuitShmRequestProto) String() string { return proto.CompactTextString(m) }
func (*ShortCircuitShmRequestProto) ProtoMessage()    {}

func (m *ShortCircuitShmRequestProto) GetClientName() string {
	if m != nil && m.ClientName != nil {
		return *m.ClientName
	}
	return ""
}

type ShortCircuitShmResponseProto struct {
	Status           *Status                 `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	Error            *string                 `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Id               *ShortCircuitShmIdProto `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ShortCircuitShmResponseProto) Reset()         { *m = ShortCircuitShmResponseProto{} }
func (m *ShortCircuitShmResponseProto) String() string { return proto.CompactTextString(m) }
func (*ShortCircuitShmResponseProto) ProtoMessage()    {}

func (m *ShortCircuitShmResponseProto) GetStatus() Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Status_SUCCESS
}

func (m *ShortCircuitShmResponseProto) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *ShortCircuitShmResponseProto) GetId() *ShortCircuitShmIdProto {
	if m != nil {
		return m.Id
	}
	return nil
}

type PacketHeaderProto struct {
	// All fields must be fixed-length!
	OffsetInBlock     *int64 `protobuf:"fixed64,1,req,name=offsetInBlock" json:"offsetInBlock,omitempty"`
	Seqno             *int64 `protobuf:"fixed64,2,req,name=seqno" json:"seqno,omitempty"`
	LastPacketInBlock *bool  `protobuf:"varint,3,req,name=lastPacketInBlock" json:"lastPacketInBlock,omitempty"`
	DataLen           *int32 `protobuf:"fixed32,4,req,name=dataLen" json:"dataLen,omitempty"`
	SyncBlock         *bool  `protobuf:"varint,5,opt,name=syncBlock,def=0" json:"syncBlock,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *PacketHeaderProto) Reset()         { *m = PacketHeaderProto{} }
func (m *PacketHeaderProto) String() string { return proto.CompactTextString(m) }
func (*PacketHeaderProto) ProtoMessage()    {}

const Default_PacketHeaderProto_SyncBlock bool = false

func (m *PacketHeaderProto) GetOffsetInBlock() int64 {
	if m != nil && m.OffsetInBlock != nil {
		return *m.OffsetInBlock
	}
	return 0
}

func (m *PacketHeaderProto) GetSeqno() int64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *PacketHeaderProto) GetLastPacketInBlock() bool {
	if m != nil && m.LastPacketInBlock != nil {
		return *m.LastPacketInBlock
	}
	return false
}

func (m *PacketHeaderProto) GetDataLen() int32 {
	if m != nil && m.DataLen != nil {
		return *m.DataLen
	}
	return 0
}

func (m *PacketHeaderProto) GetSyncBlock() bool {
	if m != nil && m.SyncBlock != nil {
		return *m.SyncBlock
	}
	return Default_PacketHeaderProto_SyncBlock
}

type PipelineAckProto struct {
	Seqno                  *int64   `protobuf:"zigzag64,1,req,name=seqno" json:"seqno,omitempty"`
	Status                 []Status `protobuf:"varint,2,rep,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	DownstreamAckTimeNanos *uint64  `protobuf:"varint,3,opt,name=downstreamAckTimeNanos,def=0" json:"downstreamAckTimeNanos,omitempty"`
	XXX_unrecognized       []byte   `json:"-"`
}

func (m *PipelineAckProto) Reset()         { *m = PipelineAckProto{} }
func (m *PipelineAckProto) String() string { return proto.CompactTextString(m) }
func (*PipelineAckProto) ProtoMessage()    {}

const Default_PipelineAckProto_DownstreamAckTimeNanos uint64 = 0

func (m *PipelineAckProto) GetSeqno() int64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *PipelineAckProto) GetStatus() []Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PipelineAckProto) GetDownstreamAckTimeNanos() uint64 {
	if m != nil && m.DownstreamAckTimeNanos != nil {
		return *m.DownstreamAckTimeNanos
	}
	return Default_PipelineAckProto_DownstreamAckTimeNanos
}

// *
// Sent as part of the BlockOpResponseProto
// for READ_BLOCK and COPY_BLOCK operations.
type ReadOpChecksumInfoProto struct {
	Checksum *ChecksumProto `protobuf:"bytes,1,req,name=checksum" json:"checksum,omitempty"`
	// *
	// The offset into the block at which the first packet
	// will start. This is necessary since reads will align
	// backwards to a checksum chunk boundary.
	ChunkOffset      *uint64 `protobuf:"varint,2,req,name=chunkOffset" json:"chunkOffset,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReadOpChecksumInfoProto) Reset()         { *m = ReadOpChecksumInfoProto{} }
func (m *ReadOpChecksumInfoProto) String() string { return proto.CompactTextString(m) }
func (*ReadOpChecksumInfoProto) ProtoMessage()    {}

func (m *ReadOpChecksumInfoProto) GetChecksum() *ChecksumProto {
	if m != nil {
		return m.Checksum
	}
	return nil
}

func (m *ReadOpChecksumInfoProto) GetChunkOffset() uint64 {
	if m != nil && m.ChunkOffset != nil {
		return *m.ChunkOffset
	}
	return 0
}

type BlockOpResponseProto struct {
	Status             *Status                       `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	FirstBadLink       *string                       `protobuf:"bytes,2,opt,name=firstBadLink" json:"firstBadLink,omitempty"`
	ChecksumResponse   *OpBlockChecksumResponseProto `protobuf:"bytes,3,opt,name=checksumResponse" json:"checksumResponse,omitempty"`
	ReadOpChecksumInfo *ReadOpChecksumInfoProto      `protobuf:"bytes,4,opt,name=readOpChecksumInfo" json:"readOpChecksumInfo,omitempty"`
	// * explanatory text which may be useful to log on the client side
	Message *string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	// * If the server chooses to agree to the request of a client for
	// short-circuit access, it will send a response message with the relevant
	// file descriptors attached.
	//
	// In the body of the message, this version number will be set to the
	// specific version number of the block data that the client is about to
	// read.
	ShortCircuitAccessVersion *uint32 `protobuf:"varint,6,opt,name=shortCircuitAccessVersion" json:"shortCircuitAccessVersion,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *BlockOpResponseProto) Reset()         { *m = BlockOpResponseProto{} }
func (m *BlockOpResponseProto) String() string { return proto.CompactTextString(m) }
func (*BlockOpResponseProto) ProtoMessage()    {}

func (m *BlockOpResponseProto) GetStatus() Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Status_SUCCESS
}

func (m *BlockOpResponseProto) GetFirstBadLink() string {
	if m != nil && m.FirstBadLink != nil {
		return *m.FirstBadLink
	}
	return ""
}

func (m *BlockOpResponseProto) GetChecksumResponse() *OpBlockChecksumResponseProto {
	if m != nil {
		return m.ChecksumResponse
	}
	return nil
}

func (m *BlockOpResponseProto) GetReadOpChecksumInfo() *ReadOpChecksumInfoProto {
	if m != nil {
		return m.ReadOpChecksumInfo
	}
	return nil
}

func (m *BlockOpResponseProto) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *BlockOpResponseProto) GetShortCircuitAccessVersion() uint32 {
	if m != nil && m.ShortCircuitAccessVersion != nil {
		return *m.ShortCircuitAccessVersion
	}
	return 0
}

// *
// Message sent from the client to the DN after reading the entire
// read request.
type ClientReadStatusProto struct {
	Status           *Status `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientReadStatusProto) Reset()         { *m = ClientReadStatusProto{} }
func (m *ClientReadStatusProto) String() string { return proto.CompactTextString(m) }
func (*ClientReadStatusProto) ProtoMessage()    {}

func (m *ClientReadStatusProto) GetStatus() Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Status_SUCCESS
}

type DNTransferAckProto struct {
	Status           *Status `protobuf:"varint,1,req,name=status,enum=hadoop.hdfs.Status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DNTransferAckProto) Reset()         { *m = DNTransferAckProto{} }
func (m *DNTransferAckProto) String() string { return proto.CompactTextString(m) }
func (*DNTransferAckProto) ProtoMessage()    {}

func (m *DNTransferAckProto) GetStatus() Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Status_SUCCESS
}

type OpBlockChecksumResponseProto struct {
	BytesPerCrc      *uint32            `protobuf:"varint,1,req,name=bytesPerCrc" json:"bytesPerCrc,omitempty"`
	CrcPerBlock      *uint64            `protobuf:"varint,2,req,name=crcPerBlock" json:"crcPerBlock,omitempty"`
	Md5              []byte             `protobuf:"bytes,3,req,name=md5" json:"md5,omitempty"`
	CrcType          *ChecksumTypeProto `protobuf:"varint,4,opt,name=crcType,enum=hadoop.hdfs.ChecksumTypeProto" json:"crcType,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *OpBlockChecksumResponseProto) Reset()         { *m = OpBlockChecksumResponseProto{} }
func (m *OpBlockChecksumResponseProto) String() string { return proto.CompactTextString(m) }
func (*OpBlockChecksumResponseProto) ProtoMessage()    {}

func (m *OpBlockChecksumResponseProto) GetBytesPerCrc() uint32 {
	if m != nil && m.BytesPerCrc != nil {
		return *m.BytesPerCrc
	}
	return 0
}

func (m *OpBlockChecksumResponseProto) GetCrcPerBlock() uint64 {
	if m != nil && m.CrcPerBlock != nil {
		return *m.CrcPerBlock
	}
	return 0
}

func (m *OpBlockChecksumResponseProto) GetMd5() []byte {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *OpBlockChecksumResponseProto) GetCrcType() ChecksumTypeProto {
	if m != nil && m.CrcType != nil {
		return *m.CrcType
	}
	return ChecksumTypeProto_CHECKSUM_NULL
}

func init() {
	proto.RegisterEnum("hadoop.hdfs.Status", Status_name, Status_value)
	proto.RegisterEnum("hadoop.hdfs.DataTransferEncryptorMessageProto_DataTransferEncryptorStatus", DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_name, DataTransferEncryptorMessageProto_DataTransferEncryptorStatus_value)
	proto.RegisterEnum("hadoop.hdfs.OpWriteBlockProto_BlockConstructionStage", OpWriteBlockProto_BlockConstructionStage_name, OpWriteBlockProto_BlockConstructionStage_value)
}
