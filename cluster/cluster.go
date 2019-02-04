package cluster

type AttributeDescriptor struct {
	Name   string
	Type   ZclDataType
	Access Access
}

type CommandDescriptor struct {
	Name    string
	Command interface{}
}

type CommandDescriptors struct {
	Received  map[uint8]*CommandDescriptor
	Generated map[uint8]*CommandDescriptor
}

type Cluster struct {
	Name                 string
	AttributeDescriptors map[uint16]*AttributeDescriptor
	CommandDescriptors   *CommandDescriptors
}

type ClusterLibrary struct {
	global   map[uint8]*CommandDescriptor
	clusters map[ClusterId]*Cluster
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
	LevelControl                   ClusterId = 0x0008
	MultistateInput                ClusterId = 0x0012
)

func New() *ClusterLibrary {
	return &ClusterLibrary{
		global: map[uint8]*CommandDescriptor{
			0x00: &CommandDescriptor{"ReadAttributes", &ReadAttributesCommand{}},
			0x01: &CommandDescriptor{"ReadAttributesResponse", &ReadAttributesResponse{}},
			0x02: &CommandDescriptor{"WriteAttributes", &WriteAttributesCommand{}},
			0x03: &CommandDescriptor{"WriteAttributesUndivided", &WriteAttributesUndividedCommand{}},
			0x04: &CommandDescriptor{"WriteAttributesResponse", &WriteAttributesResponse{}},
			0x05: &CommandDescriptor{"WriteAttributesNoResponse", &WriteAttributesNoResponseCommand{}},
			0x06: &CommandDescriptor{"ConfigureReporting", &ConfigureReportingCommand{}},
			0x07: &CommandDescriptor{"ConfigureReportingResponse", &ConfigureReportingResponse{}},
			0x08: &CommandDescriptor{"ReadReportingConfiguration", &ReadReportingConfigurationCommand{}},
			0x09: &CommandDescriptor{"ReadReportingConfigurationResponse", &ReadReportingConfigurationResponse{}},
			0x0a: &CommandDescriptor{"ReportAttributes", &ReportAttributesCommand{}},
			0x0b: &CommandDescriptor{"DefaultResponse", &DefaultResponseCommand{}},
			0x0c: &CommandDescriptor{"DiscoverAttributes", &DiscoverAttributesCommand{}},
			0x0d: &CommandDescriptor{"DiscoverAttributesResponse", &DiscoverAttributesResponse{}},
			0x0e: &CommandDescriptor{"ReadAttributesStructured", &ReadAttributesStructuredCommand{}},
			0x0f: &CommandDescriptor{"WriteAttributesStructured", &WriteAttributesStructuredCommand{}},
			0x10: &CommandDescriptor{"WriteAttributesStructuredResponse", &WriteAttributesStructuredResponse{}},
			0x11: &CommandDescriptor{"DiscoverCommandsReceived", &DiscoverCommandsReceivedCommand{}},
			0x12: &CommandDescriptor{"DiscoverCommandsReceivedResponse", &DiscoverCommandsReceivedResponse{}},
			0x13: &CommandDescriptor{"DiscoverCommandsGenerated", &DiscoverCommandsGeneratedCommand{}},
			0x14: &CommandDescriptor{"DiscoverCommandsGeneratedResponse", &DiscoverCommandsGeneratedResponse{}},
			0x15: &CommandDescriptor{"DiscoverAttributesExtended", &DiscoverAttributesExtendedCommand{}},
			0x16: &CommandDescriptor{"DiscoverAttributesExtendedResponse", &DiscoverAttributesExtendedResponse{}},
		},
		clusters: map[ClusterId]*Cluster{
			Basic: &Cluster{
				Name: "Basic",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: &AttributeDescriptor{"ZLibraryVersion", ZclDataTypeUint8, Read},
					0x0001: &AttributeDescriptor{"ApplicationVersion", ZclDataTypeUint8, Read},
					0x0002: &AttributeDescriptor{"StackVersion", ZclDataTypeUint8, Read},
					0x0003: &AttributeDescriptor{"HWVersion", ZclDataTypeUint8, Read},
					0x0004: &AttributeDescriptor{"ManufacturerName", ZclDataTypeCharStr, Read},
					0x0005: &AttributeDescriptor{"ModelIdentifier", ZclDataTypeCharStr, Read},
					0x0006: &AttributeDescriptor{"DateCode", ZclDataTypeCharStr, Read},
					0x0007: &AttributeDescriptor{"PowerSource", ZclDataTypeEnum8, Read},
					0x0010: &AttributeDescriptor{"LocationDescription", ZclDataTypeCharStr, Read | Write},
					0x0011: &AttributeDescriptor{"PhysicalEnvironment", ZclDataTypeEnum8, Read | Write},
					0x0012: &AttributeDescriptor{"DeviceEnabled", ZclDataTypeBoolean, Read | Write},
					0x0013: &AttributeDescriptor{"AlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0014: &AttributeDescriptor{"DisableLocalConfig", ZclDataTypeBitmap8, Read | Write},
					0x4000: &AttributeDescriptor{"SWBuildID", ZclDataTypeCharStr, Read},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: &CommandDescriptor{"ResetToFactoryDefaults", &ResetToFactoryDefaultsCommand{}},
					},
				},
			},
			PowerConfiguration: &Cluster{
				Name: "PowerConfiguration",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: &AttributeDescriptor{"MainsVoltage", ZclDataTypeUint16, Read},
					0x0001: &AttributeDescriptor{"MainsFrequency", ZclDataTypeUint8, Read},
					0x0010: &AttributeDescriptor{"MainsAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0011: &AttributeDescriptor{"MainsVoltageMinThreshold", ZclDataTypeUint16, Read | Write},
					0x0012: &AttributeDescriptor{"MainsVoltageMaxThreshold", ZclDataTypeUint16, Read | Write},
					0x0013: &AttributeDescriptor{"MainsVoltageDwellTripPoint", ZclDataTypeUint16, Read | Write},
					0x0020: &AttributeDescriptor{"BatteryVoltage", ZclDataTypeUint8, Read},
					0x0021: &AttributeDescriptor{"BatteryPercentageRemaining", ZclDataTypeUint8, Read | Reportable},
					0x0030: &AttributeDescriptor{"BatteryManufacturer", ZclDataTypeCharStr, Read | Write},
					0x0031: &AttributeDescriptor{"BatterySize", ZclDataTypeEnum8, Read | Write},
					0x0032: &AttributeDescriptor{"BatteryAHrRating", ZclDataTypeUint16, Read | Write},
					0x0033: &AttributeDescriptor{"BatteryQuantity", ZclDataTypeUint8, Read | Write},
					0x0034: &AttributeDescriptor{"BatteryRatedVoltage", ZclDataTypeUint8, Read | Write},
					0x0035: &AttributeDescriptor{"BatteryAlarmMask", ZclDataTypeBitmap8, Read | Write},
					0x0036: &AttributeDescriptor{"BatteryVoltageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x0037: &AttributeDescriptor{"BatteryVoltageThreshold1", ZclDataTypeUint8, Read | Write},
					0x0038: &AttributeDescriptor{"BatteryVoltageThreshold2", ZclDataTypeUint8, Read | Write},
					0x0039: &AttributeDescriptor{"BatteryVoltageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003a: &AttributeDescriptor{"BatteryPercentageMinThreshold", ZclDataTypeUint8, Read | Write},
					0x003b: &AttributeDescriptor{"BatteryPercentageThreshold1", ZclDataTypeUint8, Read | Write},
					0x003c: &AttributeDescriptor{"BatteryPercentageThreshold2", ZclDataTypeUint8, Read | Write},
					0x003d: &AttributeDescriptor{"BatteryPercentageThreshold3", ZclDataTypeUint8, Read | Write},
					0x003e: &AttributeDescriptor{"BatteryAlarmState", ZclDataTypeBitmap32, Read},
				},
			},
			DeviceTemperatureConfiguration: &Cluster{
				Name: "DeviceTemperatureConfiguration",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: &AttributeDescriptor{"CurrentTemperature", ZclDataTypeInt16, Read},
					0x0001: &AttributeDescriptor{"MinTempExperienced", ZclDataTypeInt16, Read},
					0x0002: &AttributeDescriptor{"MaxTempExperienced", ZclDataTypeInt16, Read},
					0x0003: &AttributeDescriptor{"OverTempTotalDwell", ZclDataTypeInt16, Read},
					0x0010: &AttributeDescriptor{"DeviceTempAlarmMask", ZclDataTypeBitmap16, Read | Write},
					0x0011: &AttributeDescriptor{"LowTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0012: &AttributeDescriptor{"HighTempThreshold", ZclDataTypeInt16, Read | Write},
					0x0013: &AttributeDescriptor{"LowTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
					0x0014: &AttributeDescriptor{"HighTempDwellTripPoint", ZclDataTypeUint24, Read | Write},
				},
			},
			Identify: &Cluster{
				Name: "Identify",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: &AttributeDescriptor{"IdentifyTime", ZclDataTypeInt16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: &CommandDescriptor{"Identify", &IdentifyCommand{}},
						0x01: &CommandDescriptor{"IdentifyQuery", &IdentifyQueryCommand{}},
						0x40: &CommandDescriptor{"TriggerEffect ", &TriggerEffectCommand{}},
					},
					Generated: map[uint8]*CommandDescriptor{
						0x00: &CommandDescriptor{"IdentifyQueryResponse ", &IdentifyQueryResponse{}},
					},
				},
			},
			LevelControl: &Cluster{
				Name: "LevelControl",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x0000: &AttributeDescriptor{"CurrentLevel", ZclDataTypeUint8, Read | Reportable},
					0x0001: &AttributeDescriptor{"RemainingTime", ZclDataTypeUint16, Read},
					0x0010: &AttributeDescriptor{"OnOffTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0011: &AttributeDescriptor{"OnLevel", ZclDataTypeUint8, Read | Write},
					0x0012: &AttributeDescriptor{"OnTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0013: &AttributeDescriptor{"OffTransitionTime", ZclDataTypeUint16, Read | Write},
					0x0014: &AttributeDescriptor{"DefaultMoveRate", ZclDataTypeUint16, Read | Write},
				},
				CommandDescriptors: &CommandDescriptors{
					Received: map[uint8]*CommandDescriptor{
						0x00: &CommandDescriptor{"MoveToLevel ", &MoveToLevelCommand{}},
						0x01: &CommandDescriptor{"Move", &MoveCommand{}},
						0x02: &CommandDescriptor{"Step ", &StepCommand{}},
						0x03: &CommandDescriptor{"Stop ", &StopCommand{}},
						0x04: &CommandDescriptor{"MoveToLevel/OnOff", &MoveToLevelOnOffCommand{}},
						0x05: &CommandDescriptor{"Move/OnOff", &MoveOnOffCommand{}},
						0x06: &CommandDescriptor{"Step/OnOff", &StepOnOffCommand{}},
						0x07: &CommandDescriptor{"Stop/OnOff", &StopOnOffCommand{}},
					},
				},
			},
			MultistateInput: &Cluster{
				Name: "MultistateInput",
				AttributeDescriptors: map[uint16]*AttributeDescriptor{
					0x000E: &AttributeDescriptor{"StateText", ZclDataTypeArray, Read | Write},
					0x001C: &AttributeDescriptor{"Description", ZclDataTypeCharStr, Read | Write},
					0x004A: &AttributeDescriptor{"NumberOfStates", ZclDataTypeUint16, Read | Write},
					0x0051: &AttributeDescriptor{"OutOfService", ZclDataTypeBoolean, Read | Write},
					0x0055: &AttributeDescriptor{"PresentValue", ZclDataTypeUint16, Read | Write},
					0x0067: &AttributeDescriptor{"Reliability", ZclDataTypeEnum8, Read | Write},
					0x006F: &AttributeDescriptor{"StatusFlags", ZclDataTypeBitmap8, Read},
					0x0100: &AttributeDescriptor{"ApplicationType", ZclDataTypeUint32, Read},
				},
			},
		},
	}
}

func (cl *ClusterLibrary) Clusters() map[ClusterId]*Cluster {
	return cl.clusters
}

func (cl *ClusterLibrary) Global() map[uint8]*CommandDescriptor {
	return cl.global
}
