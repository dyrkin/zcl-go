package zcl

import (
	"encoding/binary"
	"io"
	"strconv"

	"github.com/dyrkin/bin"
	"github.com/dyrkin/bin/util"
	"github.com/dyrkin/composer"
)

type ReadAttributesCommand struct {
	AttributeIDs []uint16
}

type TimeOfDay struct {
	Hours      uint8
	Minutes    uint8
	Seconds    uint8
	Hundredths uint8
}

type Date struct {
	Year       uint8
	Month      uint8
	DayOfMonth uint8
	DayOfWeek  uint8
}

type Attribute struct {
	DataType ZclDataType
	Value    interface{}
}

type ReadAttributeStatus struct {
	AttributeID uint16
	Status      ZclStatus
	Attribute   *Attribute `cond:"uint:Status==0"`
}

type ReadAttributesResponse struct {
	ReadAttributeStatuses []*ReadAttributeStatus
}

type WriteAttributeRecord struct {
	AttributeID uint16
	Attribute   *Attribute
}

type WriteAttributesCommand struct {
	WriteAttributeRecords []*WriteAttributeRecord
}

type WriteAttributeStatus struct {
	Status      ZclStatus
	AttributeID uint16
}

type WriteAttributesResponse struct {
	WriteAttributeStatuses []*WriteAttributeStatus
}

type AttributeReportingConfigurationRecord struct {
	Direction                ReportDirection
	AttributeID              uint16
	AttributeDataType        ZclDataType `cond:"uint:Direction==0"`
	MinimumReportingInterval uint16      `cond:"uint:Direction==0"`
	MaximumReportingInterval uint16      `cond:"uint:Direction==0"`
	ReportableChange         *Attribute  `cond:"uint:Direction==0"`
	TimeoutPeriod            uint16      `cond:"uint:Direction==1"`
}

type ConfigureReportingCommand struct {
	AttributeReportingConfigurationRecords []*AttributeReportingConfigurationRecord
}

type AttributeStatusRecord struct {
	Status      ZclStatus
	Direction   ReportDirection
	AttributeID uint16
}

type ConfigureReportingResponse struct {
	AttributeStatusRecords []*AttributeStatusRecord
}

type AttributeRecord struct {
	Direction   ReportDirection
	AttributeID uint16
}

type ReadReportingConfigurationCommand struct {
	AttributeRecords []*AttributeRecord
}

type AttributeReportingConfigurationResponseRecord struct {
	Status                   ZclStatus
	Direction                ReportDirection
	AttributeID              uint16
	AttributeDataType        ZclDataType `cond:"uint:Direction==0;uint:Status==0"`
	MinimumReportingInterval uint16      `cond:"uint:Direction==0;uint:Status==0"`
	MaximumReportingInterval uint16      `cond:"uint:Direction==0;uint:Status==0"`
	ReportableChange         *Attribute  `cond:"uint:Direction==0;uint:Status==0"`
	TimeoutPeriod            uint16      `cond:"uint:Direction==1;uint:Status==0"`
}

type ReadReportingConfigurationResponse struct {
	AttributeReportingConfigurationResponseRecords []*AttributeReportingConfigurationResponseRecord
}
type AttributeReport struct {
	AttributeID uint16
	Attribute   *Attribute
}

type ReportAttributesCommand struct {
	AttributeReports []*AttributeReport
}

type DefaultResponseCommand struct {
	CommandID uint8
	Status    ZclStatus
}

type DiscoverAttributesCommand struct {
	StartAttributeID            uint16
	MaximumAttributeIdentifiers uint8
}

type AttributeInformation struct {
	AttributeID       uint16
	AttributeDataType ZclDataType
}

type DiscoverAttributesResponse struct {
	DiscoveryComplete     uint8
	AttributeInformations []*AttributeInformation
}

type AttributeSelector struct {
	AttributeID uint16
	Selector    []uint16 `size:"1"`
}

type ReadAttributesStructuredCommand struct {
	AttributeSelectors []*AttributeSelector
}

type WriteAttributeStructuredRecord struct {
	AttributeID uint16
	Selector    []uint16 `size:"1"`
	Attribute   *Attribute
}

type WriteAttributesStructuredCommand struct {
	WriteAttributeStructuredRecords []*WriteAttributeStructuredRecord
}

type WriteAttributeStatusRecord struct {
	Status      ZclStatus
	AttributeID uint16
	Selector    []uint16 `size:"1"`
}

type WriteAttributesStructuredResponse struct {
	WriteAttributeStatusRecords []*WriteAttributeStatusRecord
}

type DiscoverCommandsReceivedCommand struct {
	StartCommandID            uint8
	MaximumCommandIdentifiers uint8
}

type DiscoverCommandsReceivedResponse struct {
	DiscoveryComplete  uint8
	CommandIdentifiers []uint8
}

type DiscoverCommandsGeneratedCommand struct {
	StartCommandID            uint8
	MaximumCommandIdentifiers uint8
}

type DiscoverCommandsGeneratedResponse struct {
	DiscoveryComplete  uint8
	CommandIdentifiers []uint8
}

type DiscoverAttributesExtendedCommand struct {
	StartAttributeID            uint16
	MaximumAttributeIdentifiers uint8
}

type AttributeAccessControl struct {
	Readable   uint8 `bits:"0b00000001" bitmask:"start"`
	Writeable  uint8 `bits:"0b00000010"`
	Reportable uint8 `bits:"0b00000100" bitmask:"end"`
}

type ExtendedAttributeInformation struct {
	AttributeID            uint16
	AttributeDataType      ZclDataType
	AttributeAccessControl *AttributeAccessControl
}

type DiscoverAttributesExtendedResponse struct {
	DiscoveryComplete             uint8
	ExtendedAttributeInformations []*ExtendedAttributeInformation
}

func ReadAttributesCommandFrame(command *ReadAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReadAttributes, command, direction, disableDefaultResponse)
}

func ReadAttributesResponseFrame(command *ReadAttributesResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReadAttributesResponse, command, direction, disableDefaultResponse)
}

func WriteAttributesCommandFrame(command *WriteAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributes, command, direction, disableDefaultResponse)
}

func WriteAttributesUndividedCommandFrame(command *WriteAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributesUndivided, command, direction, disableDefaultResponse)
}

func WriteAttributesResponseFrame(command *WriteAttributesResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributesResponse, command, direction, disableDefaultResponse)
}

func WriteAttributesNoResponseCommandFrame(command *WriteAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributesNoResponse, command, direction, disableDefaultResponse)
}

func ConfigureReportingCommandFrame(command *ConfigureReportingCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandConfigureReporting, command, direction, disableDefaultResponse)
}

func ConfigureReportingResponseFrame(command *ConfigureReportingResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandConfigureReportingResponse, command, direction, disableDefaultResponse)
}

func ReadReportingConfigurationCommandFrame(command *ReadReportingConfigurationCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReadReportingConfiguration, command, direction, disableDefaultResponse)
}

func ReadReportingConfigurationResponseFrame(command *ReadReportingConfigurationResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReadReportingConfigurationResponse, command, direction, disableDefaultResponse)
}

func ReportAttributesCommandFrame(command *ReportAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReportAttributes, command, direction, disableDefaultResponse)
}

func DefaultResponseCommandFrame(command *DefaultResponseCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDefaultResponse, command, direction, disableDefaultResponse)
}

func DiscoverAttributesCommandFrame(command *DiscoverAttributesCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverAttributes, command, direction, disableDefaultResponse)
}

func DiscoverAttributesResponseFrame(command *DiscoverAttributesResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverAttributesResponse, command, direction, disableDefaultResponse)
}

func ReadAttributesStructuredCommandFrame(command *ReadAttributesStructuredCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandReadAttributesStructured, command, direction, disableDefaultResponse)
}

func WriteAttributesStructuredCommandFrame(command *WriteAttributesStructuredCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributesStructured, command, direction, disableDefaultResponse)
}

func WriteAttributesStructuredResponseFrame(command *WriteAttributesStructuredResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandWriteAttributesStructuredResponse, command, direction, disableDefaultResponse)
}

func DiscoverCommandsReceivedCommandFrame(command *DiscoverCommandsReceivedCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverCommandsReceived, command, direction, disableDefaultResponse)
}

func DiscoverCommandsReceivedResponseFrame(command *DiscoverCommandsReceivedResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverCommandsReceivedResponse, command, direction, disableDefaultResponse)
}

func DiscoverCommandsGeneratedCommandFrame(command *DiscoverCommandsGeneratedCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverCommandsGenerated, command, direction, disableDefaultResponse)
}

func DiscoverCommandsGeneratedResponseFrame(command *DiscoverCommandsGeneratedResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverCommandsGenerated, command, direction, disableDefaultResponse)
}

func DiscoverAttributesExtendedCommandFrame(command *DiscoverAttributesExtendedCommand, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverAttributesExtended, command, direction, disableDefaultResponse)
}

func DiscoverAttributesExtendedResponseFrame(command *DiscoverAttributesExtendedResponse, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationNonManufacturerSpecificFrame(ZclCommandDiscoverAttributesExtendedResponse, command, direction, disableDefaultResponse)
}

func (a *Attribute) Serialize(w io.Writer) {
	c := composer.NewWithW(w)
	c.Uint8(uint8(a.DataType))

	switch a.DataType {
	case ZclDataTypeNoData:
	case ZclDataTypeData8:
		b := a.Value.([1]byte)
		c.Bytes(b[:])
	case ZclDataTypeData16:
		b := a.Value.([2]byte)
		c.Bytes(b[:])
	case ZclDataTypeData24:
		b := a.Value.([3]byte)
		c.Bytes(b[:])
	case ZclDataTypeData32:
		b := a.Value.([4]byte)
		c.Bytes(b[:])
	case ZclDataTypeData40:
		b := a.Value.([5]byte)
		c.Bytes(b[:])
	case ZclDataTypeData48:
		b := a.Value.([6]byte)
		c.Bytes(b[:])
	case ZclDataTypeData56:
		b := a.Value.([7]byte)
		c.Bytes(b[:])
	case ZclDataTypeData64:
		b := a.Value.([8]byte)
		c.Bytes(b[:])
	case ZclDataTypeBoolean:
		b := a.Value.(bool)
		if b {
			c.Byte(1)
		} else {
			c.Byte(0)
		}
	case ZclDataTypeBitmap8:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 1)
	case ZclDataTypeBitmap16:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 2)
	case ZclDataTypeBitmap24:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 3)
	case ZclDataTypeBitmap32:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 4)
	case ZclDataTypeBitmap40:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 5)
	case ZclDataTypeBitmap48:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 6)
	case ZclDataTypeBitmap56:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 7)
	case ZclDataTypeBitmap64:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 8)
	case ZclDataTypeUint8:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 1)
	case ZclDataTypeUint16:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 2)
	case ZclDataTypeUint24:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 3)
	case ZclDataTypeUint32:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 4)
	case ZclDataTypeUint40:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 5)
	case ZclDataTypeUint48:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 6)
	case ZclDataTypeUint56:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 7)
	case ZclDataTypeUint64:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 8)
	case ZclDataTypeInt8:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 1)
	case ZclDataTypeInt16:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 2)
	case ZclDataTypeInt24:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 3)
	case ZclDataTypeInt32:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 4)
	case ZclDataTypeInt40:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 5)
	case ZclDataTypeInt48:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 6)
	case ZclDataTypeInt56:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 7)
	case ZclDataTypeInt64:
		b := a.Value.(int64)
		c.Int(binary.LittleEndian, b, 8)
	case ZclDataTypeEnum8:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 1)
	case ZclDataTypeEnum16:
		b := a.Value.(uint64)
		c.Uint(binary.LittleEndian, b, 2)
	case ZclDataTypeSemiPrec:
	case ZclDataTypeSinglePrec:
	case ZclDataTypeDoublePrec:
	case ZclDataTypeOctetStr:
		b := a.Value.(string)
		c.Uint8(uint8(len(b)))
		c.String(b)
	case ZclDataTypeCharStr:
		b := a.Value.(string)
		c.Uint8(uint8(len(b)))
		c.String(b)
	case ZclDataTypeLongOctetStr:
		b := a.Value.(string)
		c.Uint16le(uint16(len(b)))
		c.String(b)
	case ZclDataTypeLongCharStr:
		b := a.Value.(string)
		c.Uint16le(uint16(len(b)))
		c.String(b)
	case ZclDataTypeArray:
	case ZclDataTypeStruct:
	case ZclDataTypeSet:
	case ZclDataTypeBag:
	case ZclDataTypeTod:
		b := a.Value.(*TimeOfDay)
		c.Uint8(b.Hours)
		c.Uint8(b.Minutes)
		c.Uint8(b.Seconds)
		c.Uint8(b.Hundredths)
	case ZclDataTypeDate:
		b := a.Value.(*Date)
		c.Uint8(b.Year)
		c.Uint8(b.Month)
		c.Uint8(b.DayOfMonth)
		c.Uint8(b.DayOfWeek)
	case ZclDataTypeUtc:
		b := a.Value.(uint32)
		c.Uint32le(b)
	case ZclDataTypeClusterId:
		b := a.Value.(uint16)
		c.Uint16le(b)
	case ZclDataTypeAttrId:
		b := a.Value.(uint16)
		c.Uint16le(b)
	case ZclDataTypeBacOid:
		b := a.Value.(uint32)
		c.Uint32le(b)
	case ZclDataTypeIeeeAddr:
		b := a.Value.(string)
		v, _ := strconv.ParseUint(b[2:], 16, 64)
		c.Uint64le(v)
	case ZclDataType_128BitSecKey:
		b := a.Value.([16]byte)
		c.Bytes(b[:])
	case ZclDataTypeUnknown:

	}
	c.Flush()
}

func (a *Attribute) Deserialize(r io.Reader) {
	c := composer.NewWithR(r)
	dataType, _ := c.ReadByte()
	a.DataType = ZclDataType(dataType)

	switch a.DataType {
	case ZclDataTypeNoData:
		a.Value = nil
	case ZclDataTypeData8:
		var buf [1]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData16:
		var buf [2]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData24:
		var buf [3]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData32:
		var buf [4]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData40:
		var buf [5]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData48:
		var buf [6]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData56:
		var buf [7]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeData64:
		var buf [8]byte
		c.ReadBuf(buf[:])
		a.Value = buf
	case ZclDataTypeBoolean:
		b, _ := c.ReadByte()
		a.Value = b > 0
	case ZclDataTypeBitmap8:
		a.Value = c.ReadUint(binary.LittleEndian, 1)
	case ZclDataTypeBitmap16:
		a.Value = c.ReadUint(binary.LittleEndian, 2)
	case ZclDataTypeBitmap24:
		a.Value = c.ReadUint(binary.LittleEndian, 3)
	case ZclDataTypeBitmap32:
		a.Value = c.ReadUint(binary.LittleEndian, 4)
	case ZclDataTypeBitmap40:
		a.Value = c.ReadUint(binary.LittleEndian, 5)
	case ZclDataTypeBitmap48:
		a.Value = c.ReadUint(binary.LittleEndian, 6)
	case ZclDataTypeBitmap56:
		a.Value = c.ReadUint(binary.LittleEndian, 7)
	case ZclDataTypeBitmap64:
		a.Value = c.ReadUint(binary.LittleEndian, 8)
	case ZclDataTypeUint8:
		a.Value = c.ReadUint(binary.LittleEndian, 1)
	case ZclDataTypeUint16:
		a.Value = c.ReadUint(binary.LittleEndian, 2)
	case ZclDataTypeUint24:
		a.Value = c.ReadUint(binary.LittleEndian, 3)
	case ZclDataTypeUint32:
		a.Value = c.ReadUint(binary.LittleEndian, 4)
	case ZclDataTypeUint40:
		a.Value = c.ReadUint(binary.LittleEndian, 5)
	case ZclDataTypeUint48:
		a.Value = c.ReadUint(binary.LittleEndian, 6)
	case ZclDataTypeUint56:
		a.Value = c.ReadUint(binary.LittleEndian, 7)
	case ZclDataTypeUint64:
		a.Value = c.ReadUint(binary.LittleEndian, 8)
	case ZclDataTypeInt8:
		a.Value = c.ReadInt(binary.LittleEndian, 1)
	case ZclDataTypeInt16:
		a.Value = c.ReadInt(binary.LittleEndian, 2)
	case ZclDataTypeInt24:
		a.Value = c.ReadInt(binary.LittleEndian, 3)
	case ZclDataTypeInt32:
		a.Value = c.ReadInt(binary.LittleEndian, 4)
	case ZclDataTypeInt40:
		a.Value = c.ReadInt(binary.LittleEndian, 5)
	case ZclDataTypeInt48:
		a.Value = c.ReadInt(binary.LittleEndian, 6)
	case ZclDataTypeInt56:
		a.Value = c.ReadInt(binary.LittleEndian, 7)
	case ZclDataTypeInt64:
		a.Value = c.ReadInt(binary.LittleEndian, 8)
	case ZclDataTypeEnum8:
		a.Value = c.ReadUint(binary.LittleEndian, 1)
	case ZclDataTypeEnum16:
		a.Value = c.ReadUint(binary.LittleEndian, 2)
	case ZclDataTypeSemiPrec:
	case ZclDataTypeSinglePrec:
	case ZclDataTypeDoublePrec:
	case ZclDataTypeOctetStr:
		len, _ := c.ReadByte()
		a.Value, _ = c.ReadString(int(len))
	case ZclDataTypeCharStr:
		len, _ := c.ReadByte()
		a.Value, _ = c.ReadString(int(len))
	case ZclDataTypeLongOctetStr:
		len, _ := c.ReadUint16le()
		a.Value, _ = c.ReadString(int(len))
	case ZclDataTypeLongCharStr:
		len, _ := c.ReadUint16le()
		a.Value, _ = c.ReadString(int(len))
	case ZclDataTypeArray:
	case ZclDataTypeStruct:
	case ZclDataTypeSet:
	case ZclDataTypeBag:
	case ZclDataTypeTod:
		hours, _ := c.ReadUint8()
		minutes, _ := c.ReadUint8()
		seconds, _ := c.ReadUint8()
		hundredths, _ := c.ReadUint8()
		a.Value = &TimeOfDay{hours, minutes, seconds, hundredths}
	case ZclDataTypeDate:
		year, _ := c.ReadUint8()
		month, _ := c.ReadUint8()
		dayOfMonth, _ := c.ReadUint8()
		dayOfWeek, _ := c.ReadUint8()
		a.Value = &Date{year, month, dayOfMonth, dayOfWeek}
	case ZclDataTypeUtc:
		a.Value, _ = c.ReadUint32le()
	case ZclDataTypeClusterId:
		a.Value, _ = c.ReadUint16le()
	case ZclDataTypeAttrId:
		a.Value, _ = c.ReadUint16le()
	case ZclDataTypeBacOid:
		a.Value, _ = c.ReadUint32le()
	case ZclDataTypeIeeeAddr:
		v, _ := c.ReadUint64le()
		a.Value, _ = util.UintToHexString(v, 8)
	case ZclDataType_128BitSecKey:
		var key [16]byte
		_ = c.ReadBuf(key[:])
		a.Value = key
	case ZclDataTypeUnknown:

	}
}

func ToFoundationNonManufacturerSpecificFrame(commandIdentifier ZclCommand, command interface{}, direction Direction, disableDefaultResponse bool) *Frame {
	return ToFoundationManufacturerSpecificFrame(commandIdentifier, command, false, 0, direction,
		disableDefaultResponse)
}

func ToFoundationManufacturerSpecificFrame(commandIdentifier ZclCommand, command interface{}, manufacturerSpecific bool, manufacturerCode uint16,
	direction Direction, disableDefaultResponse bool) *Frame {
	return ToFrame(commandIdentifier, command, FrameTypeGlobal, manufacturerSpecific, manufacturerCode, direction,
		disableDefaultResponse)
}

func ToFrame(commandIdentifier ZclCommand, command interface{}, frameType FrameType, manufacturerSpecific bool, manufacturerCode uint16,
	direction Direction, disableDefaultResponse bool) *Frame {
	return &Frame{
		&FrameControl{frameType, flag(manufacturerSpecific), direction, flag(disableDefaultResponse), 0},
		manufacturerCode, 1,
		commandIdentifier,
		bin.Encode(command),
	}
}

func flag(boolean bool) uint8 {
	if boolean {
		return 1
	}
	return 0
}
