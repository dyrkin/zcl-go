package cluster

type attributeDescriptor struct {
	Name   string
	Type   ZclDataType
	Access Access
}

type commandDescriptor struct {
	Name    string
	Command interface{}
}

type CommandDescriptors struct {
	Received  map[uint8]*commandDescriptor
	Generated map[uint8]*commandDescriptor
}

type cluster struct {
	Name                 string
	AttributeDescriptors map[uint16]*attributeDescriptor
	CommandDescriptors   *CommandDescriptors
}

type Clusters struct {
	all map[ClusterId]*cluster
}

type Zcl struct {
	global   map[uint8]*commandDescriptor
	clusters *Clusters
}

type Access uint8

const (
	Read       Access = 0x01
	Write      Access = 0x02
	Reportable Access = 0x04
)

type ClusterId uint16

const (
	Basic                          ClusterId = 0x0000
	PowerConfiguration             ClusterId = 0x0001
	DeviceTemperatureConfiguration ClusterId = 0x0002
	Identify                       ClusterId = 0x0003
)

func New() *Zcl {
	return &Zcl{
		global: map[uint8]*commandDescriptor{
			0x00: &commandDescriptor{"ReadAttributes", &ReadAttributesCommand{}},
			0x01: &commandDescriptor{"ReadAttributesResponse", &ReadAttributesResponse{}},
			0x02: &commandDescriptor{"WriteAttributes", &WriteAttributesCommand{}},
			0x03: &commandDescriptor{"WriteAttributesUndivided", &WriteAttributesUndividedCommand{}},
			0x04: &commandDescriptor{"WriteAttributesResponse", &WriteAttributesResponse{}},
			0x05: &commandDescriptor{"WriteAttributesNoResponse", &WriteAttributesNoResponseCommand{}},
			0x06: &commandDescriptor{"ConfigureReporting", &ConfigureReportingCommand{}},
			0x07: &commandDescriptor{"ConfigureReportingResponse", &ConfigureReportingResponse{}},
			0x08: &commandDescriptor{"ReadReportingConfiguration", &ReadReportingConfigurationCommand{}},
			0x09: &commandDescriptor{"ReadReportingConfigurationResponse", &ReadReportingConfigurationResponse{}},
			0x0a: &commandDescriptor{"ReportAttributes", &ReportAttributesCommand{}},
			0x0b: &commandDescriptor{"DefaultResponse", &DefaultResponseCommand{}},
			0x0c: &commandDescriptor{"DiscoverAttributes", &DiscoverAttributesCommand{}},
			0x0d: &commandDescriptor{"DiscoverAttributesResponse", &DiscoverAttributesResponse{}},
			0x0e: &commandDescriptor{"ReadAttributesStructured", &ReadAttributesStructuredCommand{}},
			0x0f: &commandDescriptor{"WriteAttributesStructured", &WriteAttributesStructuredCommand{}},
			0x10: &commandDescriptor{"WriteAttributesStructuredResponse", &WriteAttributesStructuredResponse{}},
			0x11: &commandDescriptor{"DiscoverCommandsReceived", &DiscoverCommandsReceivedCommand{}},
			0x12: &commandDescriptor{"DiscoverCommandsReceivedResponse", &DiscoverCommandsReceivedResponse{}},
			0x13: &commandDescriptor{"DiscoverCommandsGenerated", &DiscoverCommandsGeneratedCommand{}},
			0x14: &commandDescriptor{"DiscoverCommandsGeneratedResponse", &DiscoverCommandsGeneratedResponse{}},
			0x15: &commandDescriptor{"DiscoverAttributesExtended", &DiscoverAttributesExtendedCommand{}},
			0x16: &commandDescriptor{"DiscoverAttributesExtendedResponse", &DiscoverAttributesExtendedResponse{}},
		},
		clusters: &Clusters{map[ClusterId]*cluster{
			Basic: &cluster{
				Name: "Basic",
				AttributeDescriptors: map[uint16]*attributeDescriptor{
					0x0000: &attributeDescriptor{"ZCLVersion", ZclDataTypeUint8, Read},
					0x0001: &attributeDescriptor{"ApplicationVersion", ZclDataTypeUint8, Read},
					0x0002: &attributeDescriptor{"StackVersion", ZclDataTypeUint8, Read},
					0x0003: &attributeDescriptor{"HWVersion", ZclDataTypeUint8, Read},
					0x0004: &attributeDescriptor{"ManufacturerName", ZclDataTypeCharStr, Read},
					0x0005: &attributeDescriptor{"ModelIdentifier", ZclDataTypeCharStr, Read},
					0x0006: &attributeDescriptor{"DateCode", ZclDataTypeCharStr, Read},
					0x0007: &attributeDescriptor{"PowerSource", ZclDataTypeEnum8, Read},
					0x0010: &attributeDescriptor{"LocationDescription", ZclDataTypeCharStr, Read | Write},
					0x0011: &attributeDescriptor{"PhysicalEnvironment", ZclDataTypeEnum8, Read | Write},
					0x0012: &attributeDescriptor{"DeviceEnabled", ZclDataTypeBoolean, Read | Write},
					0x0013: &attributeDescriptor{"AlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0014: &attributeDescriptor{"DisableLocalConfig", ZclDataTypeBitmap8, Read | Write},
					0x4000: &attributeDescriptor{"SWBuildID", ZclDataTypeCharStr, Read},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*commandDescriptor{
						0x00: &commandDescriptor{"ResetToFactoryDefaults", &ResetToFactoryDefaultsCommand{}},
					},
				},
			},
			PowerConfiguration: &cluster{
				Name: "PowerConfiguration",
				AttributeDescriptors: map[uint16]*attributeDescriptor{
					0x0000: &attributeDescriptor{"MainsVoltage", ZclDataTypeUint16, Read},
					0x0001: &attributeDescriptor{"MainsFrequency", ZclDataTypeUint8, Read},
					0x0010: &attributeDescriptor{"MainsAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0011: &attributeDescriptor{"MainsVoltageMinThreshold", ZclDataTypeUint16, Read | Write},
					0x0012: &attributeDescriptor{"MainsVoltageMaxThreshold", ZclDataTypeUint16, Read | Write},
					0x0013: &attributeDescriptor{"MainsVoltageDwellTripPoint", ZclDataTypeUint16, Read | Write},
					0x0020: &attributeDescriptor{"BatteryVoltage", ZclDataTypeUint8, Read},
					0x0021: &attributeDescriptor{"BatteryPercentageRemaining", ZclDataTypeUint8, Read | Reportable},
					0x0030: &attributeDescriptor{"BatteryManufacturer", ZclDataTypeCharStr, Read | Write},
					0x0031: &attributeDescriptor{"BatterySize", ZclDataTypeEnum8, Read | Write},
					0x0032: &attributeDescriptor{"BatteryAHrRating", ZclDataTypeUint16, Read | Write},
					0x0033: &attributeDescriptor{"BatteryQuantity", ZclDataTypeUint8, Read | Write},
					0x0034: &attributeDescriptor{"BatteryRatedVoltage", ZclDataTypeUint8, Read | Write},
					0x0035: &attributeDescriptor{"BatteryAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0036: &attributeDescriptor{"BatteryVoltageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x0037: &attributeDescriptor{"BatteryVoltageThreshold1", ZclDataTypeUint8, Read | Write},
					0x0038: &attributeDescriptor{"BatteryVoltageThreshold2", ZclDataTypeUint8, Read | Write},
					0x0039: &attributeDescriptor{"BatteryVoltageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003a: &attributeDescriptor{"BatteryPercentageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x003b: &attributeDescriptor{"BatteryPercentageThreshold1", ZclDataTypeUint8, Read | Write},
					0x003c: &attributeDescriptor{"BatteryPercentageThreshold2", ZclDataTypeUint8, Read | Write},
					0x003d: &attributeDescriptor{"BatteryPercentageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003e: &attributeDescriptor{"BatteryAlarmState", ZclDataTypeBitmap32, Read},
				},
			},
			DeviceTemperatureConfiguration: &cluster{
				Name: "DeviceTemperatureConfiguration",
				AttributeDescriptors: map[uint16]*attributeDescriptor{
					0x0000: &attributeDescriptor{"CurrentTemperature", ZclDataTypeInt16, Read},
					0x0001: &attributeDescriptor{"MinTempExperienced", ZclDataTypeInt16, Read},
					0x0002: &attributeDescriptor{"MaxTempExperienced", ZclDataTypeInt16, Read},
					0x0003: &attributeDescriptor{"OverTempTotalDwell", ZclDataTypeInt16, Read},
					0x0010: &attributeDescriptor{"DeviceTempAlarmMask", ZclDataTypeBitmap16, Read | Write},
					0x0011: &attributeDescriptor{"LowTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0012: &attributeDescriptor{"HighTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0013: &attributeDescriptor{"LowTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
					0x0014: &attributeDescriptor{"HighTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
				},
			},
			Identify: &cluster{
				Name: "Identify",
				AttributeDescriptors: map[uint16]*attributeDescriptor{
					0x0000: &attributeDescriptor{"IdentifyTime", ZclDataTypeInt16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*commandDescriptor{
						0x00: &commandDescriptor{"Identify", &IdentifyCommand{}},
						0x01: &commandDescriptor{"IdentifyQuery", &IdentifyQuery{}},
						0x40: &commandDescriptor{"TriggerEffect ", &TriggerEffect{}},
					},
					Generated: map[uint8]*commandDescriptor{
						0x00: &commandDescriptor{"IdentifyQueryResponse ", &IdentifyQueryResponseCommand{}},
					},
				},
			},
		},
		},
	}
}
