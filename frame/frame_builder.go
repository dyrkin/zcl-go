package frame

import (
	"errors"
	"github.com/dyrkin/bin"
)

type frameConfiguration struct {
	transactionIdProvider            func() uint8
	frameType                        FrameType
	frameTypeConfigured              bool
	manufacturerCode                 uint16
	manufacturerCodeConfigured       bool
	direction                        Direction
	directionConfigured              bool
	disableDefaultResponse           bool
	disableDefaultResponseConfigured bool
	commandId                        uint8
	commandIdConfigured              bool
	command                          interface{}
	commandConfigured                bool
}

type Builder interface {
	FrameType(frameType FrameType) Builder
	ManufacturerCode(manufacturerCode uint16) Builder
	Direction(direction Direction) Builder
	DisableDefaultResponse(disableDefaultResponse bool) Builder
	CommandId(commandId uint8) Builder
	Command(command interface{}) Builder
	Build() (*Frame, error)
}

var defaultTransactionIdProvider func() uint8

func New() Builder {
	return &frameConfiguration{transactionIdProvider: defaultTransactionIdProvider}
}

func (f *frameConfiguration) IdGenerator(transactionIdProvider func() uint8) Builder {
	f.transactionIdProvider = transactionIdProvider
	return f
}

func (f *frameConfiguration) FrameType(frameType FrameType) Builder {
	f.frameType = frameType
	f.frameTypeConfigured = true
	return f
}

func (f *frameConfiguration) ManufacturerCode(manufacturerCode uint16) Builder {
	f.manufacturerCode = manufacturerCode
	f.manufacturerCodeConfigured = true
	return f
}

func (f *frameConfiguration) Direction(direction Direction) Builder {
	f.direction = direction
	f.directionConfigured = true
	return f
}

func (f *frameConfiguration) DisableDefaultResponse(disableDefaultResponse bool) Builder {
	f.disableDefaultResponse = disableDefaultResponse
	f.disableDefaultResponseConfigured = true
	return f
}

func (f *frameConfiguration) CommandId(commandId uint8) Builder {
	f.commandId = commandId
	f.commandIdConfigured = true
	return f
}

func (f *frameConfiguration) Command(command interface{}) Builder {
	f.command = command
	f.commandConfigured = true
	return f
}

func (f *frameConfiguration) Build() (*Frame, error) {
	if err := f.validateConfiguration(); err != nil {
		return nil, err
	}
	frame := &Frame{}
	frame.FrameControl = &FrameControl{}
	frame.FrameControl.FrameType = f.frameType
	frame.FrameControl.ManufacturerSpecific = flag(f.manufacturerCodeConfigured)
	frame.FrameControl.Direction = f.direction
	frame.FrameControl.DisableDefaultResponse = flag(f.disableDefaultResponse)
	frame.ManufacturerCode = f.manufacturerCode
	frame.TransactionSequenceNumber = f.transactionIdProvider()
	frame.CommandIdentifier = f.commandId
	if f.commandConfigured {
		frame.Payload = bin.Encode(f.command)
	} else {
		frame.Payload = make([]uint8, 0, 0)
	}
	return frame, nil
}

func (f *frameConfiguration) validateConfiguration() error {
	if !f.frameTypeConfigured {
		return errors.New("frame type must be set")
	}
	if !f.commandIdConfigured {
		return errors.New("command id must be set")
	}
	if !f.directionConfigured {
		return errors.New("direction must be set")
	}
	return nil
}

func flag(flag bool) uint8 {
	if flag {
		return 1
	} else {
		return 0
	}
}

func MakeDefaultTransactionIdProvider() func() uint8 {
	transactionId := uint8(1)
	return func() uint8 {
		transactionId = transactionId + 1
		if transactionId > 255 {
			transactionId = 1
		}
		return transactionId
	}
}

func init() {
	defaultTransactionIdProvider = MakeDefaultTransactionIdProvider()
}
