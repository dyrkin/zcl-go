package zcl

import (
	"testing"

	"github.com/dyrkin/bin"
	. "gopkg.in/check.v1"
)

func TestZcl(t *testing.T) { TestingT(t) }

type ZclSuite struct{}

var _ = Suite(&ZclSuite{})

func (s *ZclSuite) TestDecodeReadAttributesResponse(c *C) {
	res := &ReadAttributesResponse{}
	bin.Decode([]byte{
		127, 0, 0, byte(ZclDataTypeNoData), //NoData
		128, 0, 1, //Status Failed
		129, 0, 0, byte(ZclDataTypeData24), 0x12, 0x13, 0x14, //Data24
		130, 0, 0, byte(ZclDataTypeBitmap24), 0x12, 0x00, 0x00, //Bitmap24
		131, 0, 0, byte(ZclDataTypeBitmap32), 0x12, 0x00, 0x00, 0x00, //Bitmap32
		132, 0, 0, byte(ZclDataTypeInt24), 0xf7, 0xff, 0xff, //Int24
		133, 0, 0, byte(ZclDataTypeArray), 0x02, 0x00, byte(ZclDataTypeInt24), 0xf8, 0xff, 0xff, byte(ZclDataTypeInt24), 0xf7, 0xff, 0xff, //Array of Int24
	}, res)
	expected := &ReadAttributesResponse{
		[]*ReadAttributeStatus{
			&ReadAttributeStatus{127, ZclStatusSuccess, &Attribute{ZclDataTypeNoData, nil}},
			&ReadAttributeStatus{128, ZclStatusFailure, nil},
			&ReadAttributeStatus{129, ZclStatusSuccess, &Attribute{ZclDataTypeData24, [3]byte{0x12, 0x13, 0x14}}},
			&ReadAttributeStatus{130, ZclStatusSuccess, &Attribute{ZclDataTypeBitmap24, uint64(0x12)}},
			&ReadAttributeStatus{131, ZclStatusSuccess, &Attribute{ZclDataTypeBitmap32, uint64(0x12)}},
			&ReadAttributeStatus{132, ZclStatusSuccess, &Attribute{ZclDataTypeInt24, int64(-9)}},
			&ReadAttributeStatus{133, ZclStatusSuccess, &Attribute{ZclDataTypeArray,
				[]*Attribute{&Attribute{ZclDataTypeInt24, int64(-8)}, &Attribute{ZclDataTypeInt24, int64(-9)}}},
			},
		},
	}
	c.Assert(res, DeepEquals, expected)
}

func (s *ZclSuite) TestEncodeReadAttributesResponse(c *C) {
	a := &ReadAttributesResponse{
		[]*ReadAttributeStatus{
			&ReadAttributeStatus{127, ZclStatusSuccess, &Attribute{ZclDataTypeNoData, nil}},
			&ReadAttributeStatus{128, ZclStatusFailure, nil},
			&ReadAttributeStatus{129, ZclStatusSuccess, &Attribute{ZclDataTypeData24, [3]byte{0x12, 0x13, 0x14}}},
			&ReadAttributeStatus{130, ZclStatusSuccess, &Attribute{ZclDataTypeBitmap24, uint64(0x12)}},
			&ReadAttributeStatus{131, ZclStatusSuccess, &Attribute{ZclDataTypeBitmap32, uint64(0x12)}},
			&ReadAttributeStatus{132, ZclStatusSuccess, &Attribute{ZclDataTypeInt24, int64(-9)}},
			&ReadAttributeStatus{133, ZclStatusSuccess, &Attribute{ZclDataTypeArray,
				[]*Attribute{&Attribute{ZclDataTypeInt24, int64(-8)}, &Attribute{ZclDataTypeInt24, int64(-9)}}},
			},
		},
	}
	res := bin.Encode(a)
	expected := []byte{
		127, 0, 0, byte(ZclDataTypeNoData), //NoData
		128, 0, 1, //Status Failed
		129, 0, 0, byte(ZclDataTypeData24), 0x12, 0x13, 0x14, //Data24
		130, 0, 0, byte(ZclDataTypeBitmap24), 0x12, 0x00, 0x00, //Bitmap24
		131, 0, 0, byte(ZclDataTypeBitmap32), 0x12, 0x00, 0x00, 0x00, //Bitmap32
		132, 0, 0, byte(ZclDataTypeInt24), 0xf7, 0xff, 0xff, //Int24
		133, 0, 0, byte(ZclDataTypeArray), 0x02, 0x00, byte(ZclDataTypeInt24), 0xf8, 0xff, 0xff, byte(ZclDataTypeInt24), 0xf7, 0xff, 0xff, //Array of Int24
	}
	c.Assert(res, DeepEquals, expected)
}
