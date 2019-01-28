package zcl

import "io"

type Function struct {
	general *general
	basic   *basic
}

type basic struct {
	clusterId uint16
}

type general struct{}

type ReadAttributes struct {
	AttributeIDs []uint16
}

type Attribute struct {
	DataType ZclDataType
	Value    interface{}
}

type ReadAttributeStatus struct {
	AttributeID uint16
	Status      ZclStatus
	attribute   *Attribute `cond:"uint:Status==0"`
}

type ReadAttributesResponse struct {
	ReadAttributeStatuses []*ReadAttributeStatus
}

func (a *Attribute) Serialize(w io.Writer) {

}

func (a *Attribute) Deserialize(r io.Reader) {
	var buf [8]byte
	if _, err := r.Read(buf[:1]); err = nil {

	}
	dataType := ZclDataType(buf[0])
	a.DataType = dataType

	switch dataType {
		case ZclDataTypeNoData:
			a.Value = nil
		case ZclDataTypeData8:
			var b [1]byte
			r.Read(b[:])
			a.Value = 
		case ZclDataTypeData16:
		case ZclDataTypeData24:
		case ZclDataTypeData32:
		case ZclDataTypeData40:
		case ZclDataTypeData48:
		case ZclDataTypeData56:
		case ZclDataTypeData64:
		case ZclDataTypeBoolean:
		case ZclDataTypeBitmap8:
		case ZclDataTypeBitmap16:
		case ZclDataTypeBitmap24:
		case ZclDataTypeBitmap32:
		case ZclDataTypeBitmap40:
		case ZclDataTypeBitmap48:
		case ZclDataTypeBitmap56:
		case ZclDataTypeBitmap64:
		case ZclDataTypeUint8:
		case ZclDataTypeUint16:
		case ZclDataTypeUint24:
		case ZclDataTypeUint32:
		case ZclDataTypeUint40:
		case ZclDataTypeUint48:
		case ZclDataTypeUint56:
		case ZclDataTypeUint64:
		case ZclDataTypeInt8:
		case ZclDataTypeInt16:
		case ZclDataTypeInt24:
		case ZclDataTypeInt32:
		case ZclDataTypeInt40:
		case ZclDataTypeInt48:
		case ZclDataTypeInt56:
		case ZclDataTypeInt64:
		case ZclDataTypeEnum8:
		case ZclDataTypeEnum16:
		case ZclDataTypeSemiPrec:
		case ZclDataTypeSinglePrec:
		case ZclDataTypeDoublePrec:
		case ZclDataTypeOctetStr:
		case ZclDataTypeCharStr:
		case ZclDataTypeLongOctetStr:
		case ZclDataTypeLongCharStr:
		case ZclDataTypeArray:
		case ZclDataTypeStruct:
		case ZclDataTypeSet:
		case ZclDataTypeBag:
		case ZclDataTypeTod:
		case ZclDataTypeDate:
		case ZclDataTypeUtc:
		case ZclDataTypeClusterId:
		case ZclDataTypeAttrId:
		case ZclDataTypeBacOid:
		case ZclDataTypeIeeeAddr:
		case ZclDataType_128BitSecKey:
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
