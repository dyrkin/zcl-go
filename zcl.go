package zcl

import (
	"fmt"
	"github.com/dyrkin/bin"
	"github.com/dyrkin/zcl-go/cluster"
	"github.com/dyrkin/zcl-go/frame"
	"github.com/dyrkin/zcl-go/reflection"
	"github.com/dyrkin/znp-go"
)

type CommandExtractor func(commandDescriptors map[uint8]*cluster.CommandDescriptor) (uint8, *cluster.CommandDescriptor, error)

type ClusterQuery func(c map[cluster.ClusterId]*cluster.Cluster) (cluster.ClusterId, *cluster.Cluster, error)

type CommandQuery func(c *cluster.Cluster) (uint8, *cluster.CommandDescriptor, error)

type ZclFrameControl struct {
	FrameType              frame.FrameType
	ManufacturerSpecific   bool
	Direction              frame.Direction
	DisableDefaultResponse bool
}

type ZclFrame struct {
	FrameControl              *ZclFrameControl
	ManufacturerCode          uint16
	TransactionSequenceNumber uint8
	CommandIdentifier         uint8
	CommandName               string
	Command                   interface{}
}

type ZclIncomingMessage struct {
	GroupID              uint16
	ClusterID            uint16
	SrcAddr              string
	SrcEndpoint          uint8
	DstEndpoint          uint8
	WasBroadcast         bool
	LinkQuality          uint8
	SecurityUse          bool
	Timestamp            uint32
	TransactionSeqNumber uint8
	Data                 *ZclFrame
}

type Zcl struct {
	library *cluster.ClusterLibrary
}

func New() *Zcl {
	return &Zcl{cluster.New()}
}

func (z *Zcl) ToZclIncomingMessage(m *znp.AfIncomingMessage) (*ZclIncomingMessage, error) {
	im := &ZclIncomingMessage{}
	im.GroupID = m.GroupID
	im.ClusterID = m.ClusterID
	im.SrcAddr = m.SrcAddr
	im.SrcEndpoint = m.SrcEndpoint
	im.DstEndpoint = m.DstEndpoint
	im.WasBroadcast = m.WasBroadcast > 0
	im.LinkQuality = m.LinkQuality
	im.SecurityUse = m.SecurityUse > 0
	im.Timestamp = m.Timestamp
	im.TransactionSeqNumber = m.TransSeqNumber
	data, err := z.toZclFrame(m.Data, m.ClusterID)
	im.Data = data
	return im, err
}

func (z *Zcl) toZclFrame(data []uint8, clusterId uint16) (*ZclFrame, error) {
	frame := frame.Decode(data)
	f := &ZclFrame{}
	f.FrameControl = z.toZclFrameControl(frame.FrameControl)
	f.ManufacturerCode = frame.ManufacturerCode
	f.TransactionSequenceNumber = frame.TransactionSequenceNumber
	f.CommandIdentifier = frame.CommandIdentifier
	cmd, name, err := z.toZclCommand(clusterId, frame)
	f.CommandName = name
	f.Command = cmd
	return f, err
}

func (z *Zcl) toZclCommand(clusterId uint16, f *frame.Frame) (interface{}, string, error) {
	var cd *cluster.CommandDescriptor
	var ok bool
	switch f.FrameControl.FrameType {
	case frame.FrameTypeGlobal:
		if cd, ok = z.library.Global()[f.CommandIdentifier]; !ok {
			return nil, "", fmt.Errorf("unsupported global cmd identifier %d", f.CommandIdentifier)
		}
		cmd := cd.Command
		copy := reflection.Copy(cmd)
		bin.Decode(f.Payload, copy)
		z.patchName(copy, clusterId, f.CommandIdentifier)
		return copy, cd.Name, nil
	case frame.FrameTypeLocal:
		var c *cluster.Cluster
		if c, ok = z.library.Clusters()[cluster.ClusterId(clusterId)]; !ok {
			return nil, "", fmt.Errorf("unknown cluster %d", clusterId)
		}
		var commandDescriptors map[uint8]*cluster.CommandDescriptor
		switch f.FrameControl.Direction {
		case frame.DirectionClientServer:
			commandDescriptors = c.CommandDescriptors.Received
		case frame.DirectionServerClient:
			commandDescriptors = c.CommandDescriptors.Generated
		}
		if cd, ok = commandDescriptors[f.CommandIdentifier]; !ok {
			return nil, "", fmt.Errorf("cluster %d doesn't support this cmd %d", clusterId, f.CommandIdentifier)
		}
		cmd := cd.Command
		copy := reflection.Copy(cmd)
		bin.Decode(f.Payload, copy)
		return copy, cd.Name, nil
	}
	return nil, "", fmt.Errorf("unknown frame type")
}

func (z *Zcl) patchName(cmd interface{}, clusterId uint16, commandId uint8) {
	switch cmd := cmd.(type) {
	case *cluster.ReadAttributesResponse:
		for _, v := range cmd.ReadAttributeStatuses {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesCommand:
		for _, v := range cmd.WriteAttributeRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesUndividedCommand:
		for _, v := range cmd.WriteAttributeRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesNoResponseCommand:
		for _, v := range cmd.WriteAttributeRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesResponse:
		for _, v := range cmd.WriteAttributeStatuses {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ConfigureReportingCommand:
		for _, v := range cmd.AttributeReportingConfigurationRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ConfigureReportingResponse:
		for _, v := range cmd.AttributeStatusRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ReadReportingConfigurationCommand:
		for _, v := range cmd.AttributeRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ReadReportingConfigurationResponse:
		for _, v := range cmd.AttributeReportingConfigurationResponseRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ReportAttributesCommand:
		for _, v := range cmd.AttributeReports {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.DiscoverAttributesResponse:
		for _, v := range cmd.AttributeInformations {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.ReadAttributesStructuredCommand:
		for _, v := range cmd.AttributeSelectors {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesStructuredCommand:
		for _, v := range cmd.WriteAttributeStructuredRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.WriteAttributesStructuredResponse:
		for _, v := range cmd.WriteAttributeStatusRecords {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	case *cluster.DiscoverAttributesExtendedResponse:
		for _, v := range cmd.ExtendedAttributeInformations {
			v.AttributeName = z.getAttributeName(clusterId, v.AttributeID)
		}
	}
}

func (z *Zcl) getAttributeName(clusterId uint16, attributeId uint16) string {
	if cluster, ok := z.library.Clusters()[cluster.ClusterId(clusterId)]; ok {
		if attributeDescriptor, ok := cluster.AttributeDescriptors[attributeId]; ok {
			return attributeDescriptor.Name
		}
	}
	return ""
}

func (z *Zcl) toZclFrameControl(frameControl *frame.FrameControl) *ZclFrameControl {
	fc := &ZclFrameControl{}
	fc.FrameType = frameControl.FrameType
	fc.ManufacturerSpecific = frameControl.ManufacturerSpecific > 0
	fc.Direction = frameControl.Direction
	fc.DisableDefaultResponse = frameControl.DisableDefaultResponse > 0
	return fc
}

func (z *Zcl) ClusterLibrary() *cluster.ClusterLibrary {
	return z.library
}
