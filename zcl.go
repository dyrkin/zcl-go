package zcl

import (
	"encoding/binary"
	"io"
	"strconv"

	"github.com/dyrkin/bin/util"
	"github.com/dyrkin/composer"
)

type ReadAttributes struct {
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

// func (v *ReadAttributes) ToFrame() *Frame {
// 	c := composer.New()
// 	frameControl := &FrameControl{FrameTypeGlobal, 0x0, v.Direction, v.DisableDefaultResponse, 0}
// 	payload :=
// 	return &Frame{frameControl, 0, 1, ZclCommandReadAttributes, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
// }

// func (v *ReadAttributes) FromFrame(frame *Frame) *ReadAttributes {
// 	frameControl := &FrameControl{FrameTypeGlobal, 0x0, v.Direction, v.DisableDefaultResponse, 0}
// 	return &Frame{frameControl, 0, 1, ZclCommandReadAttributes, []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
// }
