package zcl

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEncode(c *C) {
	frame := &Frame{
		&FrameControl{1, 1, DirectionClientServer, 1, 0},
		123,
		1,
		5,
		[]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	res := Encode(frame)
	c.Assert(res, DeepEquals, []uint8{0x15, 0x7b, 0x0, 0x1, 0x5, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})

	frame = &Frame{
		&FrameControl{1, 0, DirectionClientServer, 1, 0},
		0,
		1,
		5,
		[]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	res = Encode(frame)
	c.Assert(res, DeepEquals, []uint8{0x11, 0x1, 0x5, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
}

func (s *MySuite) TestDecode(c *C) {
	frame := &Frame{
		&FrameControl{1, 1, DirectionClientServer, 1, 0},
		123,
		1,
		5,
		[]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	res := Decode([]uint8{0x15, 0x7b, 0x0, 0x1, 0x5, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	c.Assert(res, DeepEquals, frame)

	frame = &Frame{
		&FrameControl{1, 0, DirectionClientServer, 1, 0},
		0,
		1,
		5,
		[]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	res = Decode([]uint8{0x11, 0x1, 0x5, 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	c.Assert(res, DeepEquals, frame)
}
