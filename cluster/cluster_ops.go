package cluster

import (
	"fmt"

	"github.com/dyrkin/bin"
	"github.com/dyrkin/zcl-go/reflection"

	"github.com/dyrkin/zcl-go/frame"
)

type CommandExtractor func(commandDescriptors map[uint8]*commandDescriptor) (uint8, *commandDescriptor, error)

type ClusterQuery func(c *Clusters) (ClusterId, *cluster, error)

type CommandQuery func(c *cluster) (uint8, *commandDescriptor, error)

func (c *Zcl) LocalFrame(clusterQuery ClusterQuery, commandQuery CommandQuery, args ...interface{}) (f *frame.Frame, err error) {
	if _, cluster, err := clusterQuery(c.clusters); err == nil {
		if commandId, commandDescriptor, err := commandQuery(cluster); err == nil {
			command := commandDescriptor.Command
			preparedCommand := prepareCommand(command, args...)
			return createFrame(frame.FrameTypeLocal, commandId, preparedCommand), nil
		}
	}
	return
}

func ClusterByID(clusterId ClusterId) ClusterQuery {
	return func(c *Clusters) (ClusterId, *cluster, error) {
		if cluster, ok := c.all[clusterId]; ok {
			return clusterId, cluster, nil
		}
		return 0, nil, fmt.Errorf("Unknown cluster %d", clusterId)
	}
}

func ClusterByName(clusterName string) ClusterQuery {
	return func(c *Clusters) (ClusterId, *cluster, error) {
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
		return commandExtractor(c.CommandDescriptors.Received)
	}
}

func GeneratedCommand(commandExtractor CommandExtractor) CommandQuery {
	return func(c *cluster) (uint8, *commandDescriptor, error) {
		return commandExtractor(c.CommandDescriptors.Generated)
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

func prepareCommand(command interface{}, args ...interface{}) interface{} {
	copy := reflection.Copy(command)
	reflection.ApplyArgs(copy, args...)
	return copy
}

func createFrame(frameType frame.FrameType, commandId uint8, command interface{}) *frame.Frame {
	payload := make([]uint8, 0, 0)
	if command != nil {
		payload = bin.Encode(command)
	}
	return &frame.Frame{
		FrameControl:              &frame.FrameControl{frameType, 0, frame.DirectionClientServer, 0, 0},
		ManufacturerCode:          0,
		TransactionSequenceNumber: 1,
		CommandIdentifier:         commandId,
		Payload:                   payload,
	}
}
