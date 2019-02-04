package cluster

import (
	"fmt"

	"github.com/dyrkin/bin"
	"github.com/dyrkin/zcl-go/cluster"
	"github.com/dyrkin/zcl-go/frame"
	"github.com/dyrkin/zcl-go/reflection"
	znp "github.com/dyrkin/znp-go"
)

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

type ZclOutgoingOptions struct {
	WildcardProfileID uint8
	APSAck            bool
	DiscoverRoute     bool
	APSSecurity       bool
	SkipRouting       bool
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

type ZclOutgoingMessage struct {
	DstAddr              string
	DstEndpoint          uint8
	SrcEndpoint          uint8
	ClusterID            uint16
	TransactionSeqNumber uint8
	Options              *ZclOutgoingOptions
	Radius               uint8
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
	command, name, err := z.toZclCommand(clusterId, frame)
	f.CommandName = name
	f.Command = command
	return f, err
}

func (z *Zcl) toZclCommand(clusterId uint16, f *frame.Frame) (interface{}, string, error) {
	var cd *cluster.CommandDescriptor
	var ok bool
	switch f.FrameControl.FrameType {
	case frame.FrameTypeGlobal:
		if cd, ok = z.library.Global()[f.CommandIdentifier]; !ok {
			return nil, "", fmt.Errorf("Unsupported global command identifier %d", f.CommandIdentifier)
		}
	case frame.FrameTypeLocal:
		var c *cluster.Cluster
		if c, ok = z.library.Clusters()[cluster.ClusterId(clusterId)]; !ok {
			return nil, "", fmt.Errorf("Unknown cluster %d", clusterId)
		}
		var commandDescriptors map[uint8]*cluster.CommandDescriptor
		switch f.FrameControl.Direction {
		case frame.DirectionClientServer:
			commandDescriptors = c.CommandDescriptors.Received
		case frame.DirectionServerClient:
			commandDescriptors = c.CommandDescriptors.Generated
		}
		if cd, ok = commandDescriptors[f.CommandIdentifier]; !ok {
			return nil, "", fmt.Errorf("Cluster %d doesn't support this command %d", clusterId, f.CommandIdentifier)
		}
	}
	command := cd.Command
	copy := reflection.Copy(command)
	bin.Decode(f.Payload, copy)
	return copy, cd.Name, nil
}

func (z *Zcl) toZclFrameControl(frameControl *frame.FrameControl) *ZclFrameControl {
	fc := &ZclFrameControl{}
	fc.FrameType = frameControl.FrameType
	fc.ManufacturerSpecific = frameControl.ManufacturerSpecific > 0
	fc.Direction = frameControl.Direction
	fc.DisableDefaultResponse = frameControl.DisableDefaultResponse > 0
	return fc
}

func (z *Zcl) FromZclIncomingMessage(m *ZclIncomingMessage) *znp.AfIncomingMessage {
	im := &znp.AfIncomingMessage{}
	im.GroupID = m.GroupID
	im.ClusterID = m.ClusterID
	im.SrcAddr = m.SrcAddr
	im.SrcEndpoint = m.SrcEndpoint
	im.DstEndpoint = m.DstEndpoint
	im.WasBroadcast = flag(m.WasBroadcast)
	im.LinkQuality = m.LinkQuality
	im.SecurityUse = flag(m.SecurityUse)
	im.Timestamp = m.Timestamp
	im.TransSeqNumber = m.TransactionSeqNumber
	im.Data = z.fromZclFrame(m.Data, m.ClusterID)
	return im
}

func (z *Zcl) fromZclFrame(zf *ZclFrame, clusterId uint16) []uint8 {
	f := &frame.Frame{}
	f.FrameControl = z.fromZclFrameControl(zf.FrameControl)
	f.ManufacturerCode = zf.ManufacturerCode
	f.TransactionSequenceNumber = zf.TransactionSequenceNumber
	f.CommandIdentifier = zf.CommandIdentifier
	f.Payload = bin.Encode(zf.Command)
	return frame.Encode(f)
}

func (z *Zcl) fromZclFrameControl(frameControl *ZclFrameControl) *frame.FrameControl {
	fc := &frame.FrameControl{}
	fc.FrameType = frameControl.FrameType
	fc.ManufacturerSpecific = flag(frameControl.ManufacturerSpecific)
	fc.Direction = frameControl.Direction
	fc.DisableDefaultResponse = flag(frameControl.DisableDefaultResponse)
	return fc
}

func (z *Zcl) ClusterLibrary() *cluster.ClusterLibrary {
	return z.library
}

func flag(boolean bool) uint8 {
	if boolean {
		return 1
	}
	return 0
}
