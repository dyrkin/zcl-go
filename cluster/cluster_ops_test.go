package cluster

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestClusterOps(t *testing.T) { TestingT(t) }

type ClusterOpsSuite struct{}

var _ = Suite(&ClusterOpsSuite{})

func (s *ClusterOpsSuite) TestClusterByName(c *C) {
	zcl := New()
	clusterId, cluster, err := ClusterByName("Basic")(zcl.clusters)
	c.Assert(err, IsNil)
	c.Assert(clusterId, Equals, Basic)
	c.Assert(cluster, NotNil)
}

func (s *ClusterOpsSuite) TestCommandByName(c *C) {
	zcl := New()
	_, cluster, err := ClusterByID(Basic)(zcl.clusters)
	c.Assert(err, IsNil)
	commandId, command, err := CommandByName("ResetToFactoryDefaults")(cluster.CommandDescriptors.Received)
	c.Assert(err, IsNil)
	c.Assert(commandId, Equals, uint8(0x00))
	c.Assert(command, NotNil)
	c.Assert(command.Command, DeepEquals, &ResetToFactoryDefaultsCommand{})
}

func (s *ClusterOpsSuite) TestLocalFrame(c *C) {
	zcl := New()
	frame, err := zcl.LocalFrame(ClusterByID(Identify), ReceiveCommand(CommandByName("Identify")), uint16(16))
	c.Assert(err, IsNil)
	c.Assert(frame.Payload, DeepEquals, []uint8{16, 0})
}
