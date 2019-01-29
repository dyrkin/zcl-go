package zcl

import "github.com/dyrkin/bin"

type FrameControl struct {
	FrameType              FrameType `bits:"0b00000011" bitmask:"start"`
	ManufacturerSpecific   uint8     `bits:"0b00000100"`
	Direction              Direction `bits:"0b00001000"`
	DisableDefaultResponse uint8     `bits:"0b00010000"`
	Reserved               uint8     `bits:"0b11100000" bitmask:"end"`
}

type Frame struct {
	FrameControl              *FrameControl
	ManufacturerCode          uint16 `cond:"uint:FrameControl.ManufacturerSpecific==1"`
	TransactionSequenceNumber uint8
	CommandIdentifier         ZclCommand
	Payload                   []uint8
}

func Decode(buf []uint8) *Frame {
	frame := &Frame{}
	bin.Decode(buf, frame)
	return frame
}

func Encode(frame *Frame) []uint8 {
	return bin.Encode(frame)
}
