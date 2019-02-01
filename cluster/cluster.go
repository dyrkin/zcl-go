package cluster

import (
	"fmt"

	"github.com/dyrkin/bin"
	zcl "github.com/dyrkin/zcl-go"
	. "github.com/dyrkin/zcl-go/frame"
	"github.com/dyrkin/zcl-go/reflection"
)

type attributeDescriptor struct {
	Name   string
	Type   zcl.ZclDataType
	Access Access
}

type commandDescriptor struct {
	Name    string
	Command interface{}
}

type cluster struct {
	Name                        string
	AttributeDescriptors        map[uint16]*attributeDescriptor
	CommandReceivedDescriptors  map[uint8]*commandDescriptor
	CommandGeneratedDescriptors map[uint8]*commandDescriptor
}

type clusters struct {
	all map[ClusterId]*cluster
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

func New() *clusters {
	return &clusters{map[ClusterId]*cluster{
		Basic: &cluster{
			Name: "Basic",
			AttributeDescriptors: map[uint16]*attributeDescriptor{
				0x0000: &attributeDescriptor{"ZCLVersion", zcl.ZclDataTypeUint8, Read},
				0x0001: &attributeDescriptor{"ApplicationVersion", zcl.ZclDataTypeUint8, Read},
				0x0002: &attributeDescriptor{"StackVersion", zcl.ZclDataTypeUint8, Read},
				0x0003: &attributeDescriptor{"HWVersion", zcl.ZclDataTypeUint8, Read},
				0x0004: &attributeDescriptor{"ManufacturerName", zcl.ZclDataTypeCharStr, Read},
				0x0005: &attributeDescriptor{"ModelIdentifier", zcl.ZclDataTypeCharStr, Read},
				0x0006: &attributeDescriptor{"DateCode", zcl.ZclDataTypeCharStr, Read},
				0x0007: &attributeDescriptor{"PowerSource", zcl.ZclDataTypeEnum8, Read},
				0x0010: &attributeDescriptor{"LocationDescription", zcl.ZclDataTypeCharStr, Read | Write},
				0x0011: &attributeDescriptor{"PhysicalEnvironment", zcl.ZclDataTypeEnum8, Read | Write},
				0x0012: &attributeDescriptor{"DeviceEnabled", zcl.ZclDataTypeBoolean, Read | Write},
				0x0013: &attributeDescriptor{"AlarmMask", zcl.ZclDataTypeBitmap8, Read | Write},
				0x0014: &attributeDescriptor{"DisableLocalConfig", zcl.ZclDataTypeBitmap8, Read | Write},
				0x4000: &attributeDescriptor{"SWBuildID", zcl.ZclDataTypeCharStr, Read},
			},
			CommandReceivedDescriptors: map[uint8]*commandDescriptor{
				0x00: &commandDescriptor{"ResetToFactoryDefaults", &ResetToFactoryDefaultsCommand{}},
			},
		},
		PowerConfiguration: &cluster{
			Name: "PowerConfiguration",
			AttributeDescriptors: map[uint16]*attributeDescriptor{
				0x0000: &attributeDescriptor{"MainsVoltage", zcl.ZclDataTypeUint16, Read},
				0x0001: &attributeDescriptor{"MainsFrequency", zcl.ZclDataTypeUint8, Read},
				0x0010: &attributeDescriptor{"MainsAlarmMask", zcl.ZclDataTypeBitmap8, Read | Write},
				0x0011: &attributeDescriptor{"MainsVoltageMinThreshold", zcl.ZclDataTypeUint16, Read | Write},
				0x0012: &attributeDescriptor{"MainsVoltageMaxThreshold", zcl.ZclDataTypeUint16, Read | Write},
				0x0013: &attributeDescriptor{"MainsVoltageDwellTripPoint", zcl.ZclDataTypeUint16, Read | Write},
				0x0020: &attributeDescriptor{"BatteryVoltage", zcl.ZclDataTypeUint8, Read},
				0x0021: &attributeDescriptor{"BatteryPercentageRemaining", zcl.ZclDataTypeUint8, Read | Reportable},
				0x0030: &attributeDescriptor{"BatteryManufacturer", zcl.ZclDataTypeCharStr, Read | Write},
				0x0031: &attributeDescriptor{"BatterySize", zcl.ZclDataTypeEnum8, Read | Write},
				0x0032: &attributeDescriptor{"BatteryAHrRating", zcl.ZclDataTypeUint16, Read | Write},
				0x0033: &attributeDescriptor{"BatteryQuantity", zcl.ZclDataTypeUint8, Read | Write},
				0x0034: &attributeDescriptor{"BatteryRatedVoltage", zcl.ZclDataTypeUint8, Read | Write},
				0x0035: &attributeDescriptor{"BatteryAlarmMask", zcl.ZclDataTypeBitmap8, Read | Write},
				0x0036: &attributeDescriptor{"BatteryVoltageMinThreshold", zcl.ZclDataTypeUint8, Read | Write},
				0x0037: &attributeDescriptor{"BatteryVoltageThreshold1", zcl.ZclDataTypeUint8, Read | Write},
				0x0038: &attributeDescriptor{"BatteryVoltageThreshold2", zcl.ZclDataTypeUint8, Read | Write},
				0x0039: &attributeDescriptor{"BatteryVoltageThreshold3", zcl.ZclDataTypeUint8, Read | Write},
				0x003a: &attributeDescriptor{"BatteryPercentageMinThreshold", zcl.ZclDataTypeUint8, Read | Write},
				0x003b: &attributeDescriptor{"BatteryPercentageThreshold1", zcl.ZclDataTypeUint8, Read | Write},
				0x003c: &attributeDescriptor{"BatteryPercentageThreshold2", zcl.ZclDataTypeUint8, Read | Write},
				0x003d: &attributeDescriptor{"BatteryPercentageThreshold3", zcl.ZclDataTypeUint8, Read | Write},
				0x003e: &attributeDescriptor{"BatteryAlarmState", zcl.ZclDataTypeBitmap32, Read},
			},
		},
		DeviceTemperatureConfiguration: &cluster{
			Name: "DeviceTemperatureConfiguration",
			AttributeDescriptors: map[uint16]*attributeDescriptor{
				0x0000: &attributeDescriptor{"CurrentTemperature", zcl.ZclDataTypeInt16, Read},
				0x0001: &attributeDescriptor{"MinTempExperienced", zcl.ZclDataTypeInt16, Read},
				0x0002: &attributeDescriptor{"MaxTempExperienced", zcl.ZclDataTypeInt16, Read},
				0x0003: &attributeDescriptor{"OverTempTotalDwell", zcl.ZclDataTypeInt16, Read},
				0x0010: &attributeDescriptor{"DeviceTempAlarmMask", zcl.ZclDataTypeBitmap16, Read | Write},
				0x0011: &attributeDescriptor{"LowTempThreshold", zcl.ZclDataTypeInt16, Read | Write},
				0x0012: &attributeDescriptor{"HighTempThreshold", zcl.ZclDataTypeInt16, Read | Write},
				0x0013: &attributeDescriptor{"LowTempDwellTripPoint", zcl.ZclDataTypeUint24, Read | Write},
				0x0014: &attributeDescriptor{"HighTempDwellTripPoint", zcl.ZclDataTypeUint24, Read | Write},
			},
		},
		Identify: &cluster{
			Name: "Identify",
			AttributeDescriptors: map[uint16]*attributeDescriptor{
				0x0000: &attributeDescriptor{"IdentifyTime", zcl.ZclDataTypeInt16, Read | Write},
			},
			CommandReceivedDescriptors: map[uint8]*commandDescriptor{
				0x00: &commandDescriptor{"Identify", &IdentifyCommand{}},
				0x01: &commandDescriptor{"IdentifyQuery", &IdentifyQuery{}},
				0x40: &commandDescriptor{"TriggerEffect ", &TriggerEffect{}},
			},
			CommandGeneratedDescriptors: map[uint8]*commandDescriptor{
				0x00: &commandDescriptor{"IdentifyQueryResponse ", &IdentifyQueryResponseCommand{}},
			},
		},
	},
	}
}

type CommandExtractor func(commandDescriptors map[uint8]*commandDescriptor) (uint8, *commandDescriptor, error)

type ClusterQuery func(c *clusters) (ClusterId, *cluster, error)

type CommandQuery func(c *cluster) (uint8, *commandDescriptor, error)

//ClusterByID returns cluster by its id
func ClusterByID(clusterId ClusterId) ClusterQuery {
	return func(c *clusters) (ClusterId, *cluster, error) {
		if cluster, ok := c.all[clusterId]; ok {
			return clusterId, cluster, nil
		}
		return 0, nil, fmt.Errorf("Unknown cluster %d", clusterId)
	}
}

func ClusterByName(clusterName string) ClusterQuery {
	return func(c *clusters) (ClusterId, *cluster, error) {
		for k, v := range c.all {
			if v.Name == clusterName {
				return k, v, nil
			}
		}
		return 0, nil, fmt.Errorf("Unknown cluster %q", clusterName)
	}
}

func ReceiveCommand(commandExtractor CommandExtractor) CommandQuery {
	return func(c *cluster) (uint8, *commandDescriptor, error) {
		return commandExtractor(c.CommandReceivedDescriptors)
	}
}

func GeneratedCommand(commandExtractor CommandExtractor) CommandQuery {
	return func(c *cluster) (uint8, *commandDescriptor, error) {
		return commandExtractor(c.CommandGeneratedDescriptors)
	}
}

func CommandByName(commandName string) CommandExtractor {
	return func(commandDescriptors map[uint8]*commandDescriptor) (uint8, *commandDescriptor, error) {
		for commandId, command := range commandDescriptors {
			if command.Name == commandName {
				return commandId, command, nil
			}
		}
		return 0, nil, fmt.Errorf("Unknown command %q", commandName)
	}
}

func CommandById(commandId uint8) CommandExtractor {
	return func(commandDescriptors map[uint8]*commandDescriptor) (uint8, *commandDescriptor, error) {
		if command, ok := commandDescriptors[commandId]; ok {
			return commandId, command, nil
		}
		return 0, nil, fmt.Errorf("Unknown command %d", commandId)
	}
}

func (c *clusters) FunctionalFrame(clusterQuery ClusterQuery, commandQuery CommandQuery, args ...interface{}) (frame *Frame, err error) {
	if _, cluster, err := clusterQuery(c); err == nil {
		if commandId, commandDescriptor, err := commandQuery(cluster); err == nil {
			command := commandDescriptor.Command
			preparedCommand := prepareCommand(command, args...)
			return createFrame(FrameTypeLocal, commandId, preparedCommand), nil
		}
	}
	return
}

func prepareCommand(command interface{}, args ...interface{}) interface{} {
	copy := reflection.Copy(command)
	reflection.ApplyArgs(copy, args...)
	return copy
}

func createFrame(frameType FrameType, commandId uint8, command interface{}) *Frame {
	payload := make([]uint8, 0, 0)
	if command != nil {
		payload = bin.Encode(command)
	}
	return &Frame{
		FrameControl:              &FrameControl{frameType, 0, DirectionClientServer, 0, 0},
		ManufacturerCode:          0,
		TransactionSequenceNumber: 1,
		CommandIdentifier:         commandId,
		Payload:                   payload,
	}
}
