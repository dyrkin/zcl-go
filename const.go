package zcl

type ZclClusterId uint16

const (
	//General Clusters
	ZclClusterIdGenBasic                 ZclClusterId = 0x0000
	ZclClusterIdGenPowerCfg              ZclClusterId = 0x0001
	ZclClusterIdGenDeviceTempConfig      ZclClusterId = 0x0002
	ZclClusterIdGenIdentify              ZclClusterId = 0x0003
	ZclClusterIdGenGroups                ZclClusterId = 0x0004
	ZclClusterIdGenScenes                ZclClusterId = 0x0005
	ZclClusterIdGenOnOff                 ZclClusterId = 0x0006
	ZclClusterIdGenOnOffSwitchConfig     ZclClusterId = 0x0007
	ZclClusterIdGenLevelControl          ZclClusterId = 0x0008
	ZclClusterIdGenAlarms                ZclClusterId = 0x0009
	ZclClusterIdGenTime                  ZclClusterId = 0x000A
	ZclClusterIdGenLocation              ZclClusterId = 0x000B
	ZclClusterIdGenAnalogInputBasic      ZclClusterId = 0x000C
	ZclClusterIdGenAnalogOutputBasic     ZclClusterId = 0x000D
	ZclClusterIdGenAnalogValueBasic      ZclClusterId = 0x000E
	ZclClusterIdGenBinaryInputBasic      ZclClusterId = 0x000F
	ZclClusterIdGenBinaryOutputBasic     ZclClusterId = 0x0010
	ZclClusterIdGenBinaryValueBasic      ZclClusterId = 0x0011
	ZclClusterIdGenMultistateInputBasic  ZclClusterId = 0x0012
	ZclClusterIdGenMultistateOutputBasic ZclClusterId = 0x0013
	ZclClusterIdGenMultistateValueBasic  ZclClusterId = 0x0014
	ZclClusterIdGenCommissioning         ZclClusterId = 0x0015
	ZclClusterIdGenPartition             ZclClusterId = 0x0016

	ZclClusterIdOta ZclClusterId = 0x0019

	ZclClusterIdGenPowerProfile     ZclClusterId = 0x001A
	ZclClusterIdGenApplianceControl ZclClusterId = 0x001B

	ZclClusterIdGenPollControl ZclClusterId = 0x0020

	ZclClusterIdGreenPower ZclClusterId = 0x0021

	// Retail Clusters
	ZclClusterIdMobileDeviceConfiguration ZclClusterId = 0x0022
	ZclClusterIdNeighborCleaning          ZclClusterId = 0x0023
	ZclClusterIdNearestGateway            ZclClusterId = 0x0024

	// Closures Clusters
	ZclClusterIdClosuresShadeConfig    ZclClusterId = 0x0100
	ZclClusterIdClosuresDoorLock       ZclClusterId = 0x0101
	ZclClusterIdClosuresWindowCovering ZclClusterId = 0x0102

	// HVAC Clusters
	ZclClusterIdHvacPumpConfigControl       ZclClusterId = 0x0200
	ZclClusterIdHvacThermostat              ZclClusterId = 0x0201
	ZclClusterIdHvacFanControl              ZclClusterId = 0x0202
	ZclClusterIdHvacDihumidificationControl ZclClusterId = 0x0203
	ZclClusterIdHvacUserInterfaceConfig     ZclClusterId = 0x0204

	// Lighting Clusters
	ZclClusterIdLightingColorControl  ZclClusterId = 0x0300
	ZclClusterIdLightingBallastConfig ZclClusterId = 0x0301

	// Measurement and Sensing Clusters
	ZclClusterIdMsIlluminanceMeasurement        ZclClusterId = 0x0400
	ZclClusterIdMsIlluminanceLevelSensingConfig ZclClusterId = 0x0401
	ZclClusterIdMsTemperatureMeasurement        ZclClusterId = 0x0402
	ZclClusterIdMsPressureMeasurement           ZclClusterId = 0x0403
	ZclClusterIdMsFlowMeasurement               ZclClusterId = 0x0404
	ZclClusterIdMsRelativeHumidity              ZclClusterId = 0x0405
	ZclClusterIdMsOccupancySensing              ZclClusterId = 0x0406

	// Security and Safety (SS) Clusters
	ZclClusterIdSsIasZone ZclClusterId = 0x0500
	ZclClusterIdSsIasAce  ZclClusterId = 0x0501
	ZclClusterIdSsIasWd   ZclClusterId = 0x0502

	// Protocol Interfaces
	ZclClusterIdPiGenericTunnel             ZclClusterId = 0x0600
	ZclClusterIdPiBacnetProtocolTunnel      ZclClusterId = 0x0601
	ZclClusterIdPiAnalogInputBacnetReg      ZclClusterId = 0x0602
	ZclClusterIdPiAnalogInputBacnetExt      ZclClusterId = 0x0603
	ZclClusterIdPiAnalogOutputBacnetReg     ZclClusterId = 0x0604
	ZclClusterIdPiAnalogOutputBacnetExt     ZclClusterId = 0x0605
	ZclClusterIdPiAnalogValueBacnetReg      ZclClusterId = 0x0606
	ZclClusterIdPiAnalogValueBacnetExt      ZclClusterId = 0x0607
	ZclClusterIdPiBinaryInputBacnetReg      ZclClusterId = 0x0608
	ZclClusterIdPiBinaryInputBacnetExt      ZclClusterId = 0x0609
	ZclClusterIdPiBinaryOutputBacnetReg     ZclClusterId = 0x060A
	ZclClusterIdPiBinaryOutputBacnetExt     ZclClusterId = 0x060B
	ZclClusterIdPiBinaryValueBacnetReg      ZclClusterId = 0x060C
	ZclClusterIdPiBinaryValueBacnetExt      ZclClusterId = 0x060D
	ZclClusterIdPiMultistateInputBacnetReg  ZclClusterId = 0x060E
	ZclClusterIdPiMultistateInputBacnetExt  ZclClusterId = 0x060F
	ZclClusterIdPiMultistateOutputBacnetReg ZclClusterId = 0x0610
	ZclClusterIdPiMultistateOutputBacnetExt ZclClusterId = 0x0611
	ZclClusterIdPiMultistateValueBacnetReg  ZclClusterId = 0x0612
	ZclClusterIdPiMultistateValueBacnetExt  ZclClusterId = 0x0613
	ZclClusterIdPi_11073ProtocolTunnel      ZclClusterId = 0x0614
	ZclClusterIdPiIso7818ProtocolTunnel     ZclClusterId = 0x0615
	ZclClusterIdPiRetailTunnel              ZclClusterId = 0x0617

	// Smart Energy (SE) Clusters
	ZclClusterIdSePrice      ZclClusterId = 0x0700
	ZclClusterIdSeDrlc       ZclClusterId = 0x0701
	ZclClusterIdSeMetering   ZclClusterId = 0x0702
	ZclClusterIdSeMessaging  ZclClusterId = 0x0703
	ZclClusterIdSeTunneling  ZclClusterId = 0x0704
	ZclClusterIdSePrepayment ZclClusterId = 0x0705
	ZclClusterIdSeEnergyMgmt ZclClusterId = 0x0706
	ZclClusterIdSeCalendar   ZclClusterId = 0x0707
	ZclClusterIdSeDeviceMgmt ZclClusterId = 0x0708
	ZclClusterIdSeEvents     ZclClusterId = 0x0709
	ZclClusterIdSeMduPairing ZclClusterId = 0x070A

	ZclClusterIdSeKeyEstablishment ZclClusterId = 0x0800

	ZclClusterIdTelecommunicationsInformation     ZclClusterId = 0x0900
	ZclClusterIdTelecommunicationsVoiceOverZigbee ZclClusterId = 0x0904
	ZclClusterIdTelecommunicationsChatting        ZclClusterId = 0x0905

	ZclClusterIdHaApplianceIdentification ZclClusterId = 0x0B00
	ZclClusterIdHaMeterIdentification     ZclClusterId = 0x0B01
	ZclClusterIdHaApplianceEventsAlerts   ZclClusterId = 0x0B02
	ZclClusterIdHaApplianceStatistics     ZclClusterId = 0x0B03
	ZclClusterIdHaElectricalMeasurement   ZclClusterId = 0x0B04
	ZclClusterIdHaDiagnostic              ZclClusterId = 0x0B05

	// Light Link cluster
	ZclClusterIdTouchlink ZclClusterId = 0x1000
)

type ZclDataType uint8

const (
	ZclDataTypeNoData        ZclDataType = 0x00
	ZclDataTypeData8         ZclDataType = 0x08
	ZclDataTypeData16        ZclDataType = 0x09
	ZclDataTypeData24        ZclDataType = 0x0a
	ZclDataTypeData32        ZclDataType = 0x0b
	ZclDataTypeData40        ZclDataType = 0x0c
	ZclDataTypeData48        ZclDataType = 0x0d
	ZclDataTypeData56        ZclDataType = 0x0e
	ZclDataTypeData64        ZclDataType = 0x0f
	ZclDataTypeBoolean       ZclDataType = 0x10
	ZclDataTypeBitmap8       ZclDataType = 0x18
	ZclDataTypeBitmap16      ZclDataType = 0x19
	ZclDataTypeBitmap24      ZclDataType = 0x1a
	ZclDataTypeBitmap32      ZclDataType = 0x1b
	ZclDataTypeBitmap40      ZclDataType = 0x1c
	ZclDataTypeBitmap48      ZclDataType = 0x1d
	ZclDataTypeBitmap56      ZclDataType = 0x1e
	ZclDataTypeBitmap64      ZclDataType = 0x1f
	ZclDataTypeUint8         ZclDataType = 0x20
	ZclDataTypeUint16        ZclDataType = 0x21
	ZclDataTypeUint24        ZclDataType = 0x22
	ZclDataTypeUint32        ZclDataType = 0x23
	ZclDataTypeUint40        ZclDataType = 0x24
	ZclDataTypeUint48        ZclDataType = 0x25
	ZclDataTypeUint56        ZclDataType = 0x26
	ZclDataTypeUint64        ZclDataType = 0x27
	ZclDataTypeInt8          ZclDataType = 0x28
	ZclDataTypeInt16         ZclDataType = 0x29
	ZclDataTypeInt24         ZclDataType = 0x2a
	ZclDataTypeInt32         ZclDataType = 0x2b
	ZclDataTypeInt40         ZclDataType = 0x2c
	ZclDataTypeInt48         ZclDataType = 0x2d
	ZclDataTypeInt56         ZclDataType = 0x2e
	ZclDataTypeInt64         ZclDataType = 0x2f
	ZclDataTypeEnum8         ZclDataType = 0x30
	ZclDataTypeEnum16        ZclDataType = 0x31
	ZclDataTypeSemiPrec      ZclDataType = 0x38
	ZclDataTypeSinglePrec    ZclDataType = 0x39
	ZclDataTypeDoublePrec    ZclDataType = 0x3a
	ZclDataTypeOctetStr      ZclDataType = 0x41
	ZclDataTypeCharStr       ZclDataType = 0x42
	ZclDataTypeLongOctetStr  ZclDataType = 0x43
	ZclDataTypeLongCharStr   ZclDataType = 0x44
	ZclDataTypeArray         ZclDataType = 0x48
	ZclDataTypeStruct        ZclDataType = 0x4c
	ZclDataTypeSet           ZclDataType = 0x50
	ZclDataTypeBag           ZclDataType = 0x51
	ZclDataTypeTod           ZclDataType = 0xe0
	ZclDataTypeDate          ZclDataType = 0xe1
	ZclDataTypeUtc           ZclDataType = 0xe2
	ZclDataTypeClusterId     ZclDataType = 0xe8
	ZclDataTypeAttrId        ZclDataType = 0xe9
	ZclDataTypeBacOid        ZclDataType = 0xea
	ZclDataTypeIeeeAddr      ZclDataType = 0xf0
	ZclDataType_128BitSecKey ZclDataType = 0xf1
	ZclDataTypeUnknown       ZclDataType = 0xff
)

type ZclStatus uint8

const (
	ZclStatusSuccess ZclStatus = 0x00
	ZclStatusFailure ZclStatus = 0x01
	// 0x02-0x7D are reserved.
	ZclStatusNotAuthorized            ZclStatus = 0x7E
	ZclStatusMalformedCommand         ZclStatus = 0x80
	ZclStatusUnsupClusterCommand      ZclStatus = 0x81
	ZclStatusUnsupGeneralCommand      ZclStatus = 0x82
	ZclStatusUnsupManuClusterCommand  ZclStatus = 0x83
	ZclStatusUnsupManuGeneralCommand  ZclStatus = 0x84
	ZclStatusInvalidField             ZclStatus = 0x85
	ZclStatusUnsupportedAttribute     ZclStatus = 0x86
	ZclStatusInvalidValue             ZclStatus = 0x87
	ZclStatusReadOnly                 ZclStatus = 0x88
	ZclStatusInsufficientSpace        ZclStatus = 0x89
	ZclStatusDuplicateExists          ZclStatus = 0x8a
	ZclStatusNotFound                 ZclStatus = 0x8b
	ZclStatusUnreportableAttribute    ZclStatus = 0x8c
	ZclStatusInvalidDataType          ZclStatus = 0x8d
	ZclStatusInvalidSelector          ZclStatus = 0x8e
	ZclStatusWriteOnly                ZclStatus = 0x8f
	ZclStatusInconsistentStartupState ZclStatus = 0x90
	ZclStatusDefinedOutOfBand         ZclStatus = 0x91
	ZclStatusInconsistent             ZclStatus = 0x92
	ZclStatusActionDenied             ZclStatus = 0x93
	ZclStatusTimeout                  ZclStatus = 0x94
	ZclStatusAbort                    ZclStatus = 0x95
	ZclStatusInvalidImage             ZclStatus = 0x96
	ZclStatusWaitForData              ZclStatus = 0x97
	ZclStatusNoImageAvailable         ZclStatus = 0x98
	ZclStatusRequireMoreImage         ZclStatus = 0x99

	// 0xbd-bf are reserved.
	ZclStatusHardwareFailure  ZclStatus = 0xc0
	ZclStatusSoftwareFailure  ZclStatus = 0xc1
	ZclStatusCalibrationError ZclStatus = 0xc2
	// 0xc3-0xff are reserved.
	ZclStatusCmdHasRsp ZclStatus = 0xFF // Non-standard status (used for Default Rsp)
)

type ZclCommand uint8

const (
	ZclCommandReadAttributes                     ZclCommand = 0x00
	ZclCommandReadAttributesResponse             ZclCommand = 0x01
	ZclCommandWriteAttributes                    ZclCommand = 0x02
	ZclCommandWriteAttributesUndivided           ZclCommand = 0x03
	ZclCommandWriteAttributesResponse            ZclCommand = 0x04
	ZclCommandWriteAttributesNoResponse          ZclCommand = 0x05
	ZclCommandConfigureReporting                 ZclCommand = 0x06
	ZclCommandConfigureReportingResponse         ZclCommand = 0x07
	ZclCommandReadReportingConfiguration         ZclCommand = 0x08
	ZclCommandReadReportingConfigurationResponse ZclCommand = 0x09
	ZclCommandReportAttributes                   ZclCommand = 0x0a
	ZclCommandDefaultResponse                    ZclCommand = 0x0b
	ZclCommandDiscoverAttributes                 ZclCommand = 0x0c
	ZclCommandDiscoverAttributesResponse         ZclCommand = 0x0d
	ZclCommandReadAttributesStructured           ZclCommand = 0x0e
	ZclCommandWriteAttributesStructured          ZclCommand = 0x0f
	ZclCommandWriteAttributesStructuredResponse  ZclCommand = 0x10
	ZclCommandDiscoverCommandsReceived           ZclCommand = 0x11
	ZclCommandDiscoverCommandsReceivedResponse   ZclCommand = 0x12
	ZclCommandDiscoverCommandsGenerated          ZclCommand = 0x13
	ZclCommandDiscoverCommandsGeneratedResponse  ZclCommand = 0x14
	ZclCommandDiscoverAttributesExtended         ZclCommand = 0x15
	ZclCommandDiscoverAttributesExtendedResponse ZclCommand = 0x16
)
