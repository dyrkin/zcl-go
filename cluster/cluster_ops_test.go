package cluster

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestClusterOps(t *testing.T) { TestingT(t) }

type ClusterOpsSuite struct{}

var _ = Suite(&ClusterOpsSuite{})

func (s *ClusterOpsSuite) TestClusterByName(c *C) {
	lib := New()
	clusterId, cluster, err := ClusterByName("Basic")(lib.Clusters)
	c.Assert(err, IsNil)
	c.Assert(clusterId, Equals, Basic)
	c.Assert(cluster, NotNil)
}

func (s *ClusterOpsSuite) TestCommandByName(c *C) {
	lib := New()
	_, cluster, err := ClusterByID(Basic)(lib.Clusters)
	c.Assert(err, IsNil)
	commandId, command, err := CommandByName("ResetToFactoryDefaults")(cluster.CommandDescriptors.Received)
	c.Assert(err, IsNil)
	c.Assert(commandId, Equals, uint8(0x00))
	c.Assert(command, NotNil)
	c.Assert(command.Command, DeepEquals, &ResetToFactoryDefaultsCommand{})
}

func (s *ClusterOpsSuite) TestLocalFrame(c *C) {
	frame, err := Library.LocalFrame(ClusterByID(Identify), ReceiveCommand(CommandByName("Identify")), uint16(16))
	c.Assert(err, IsNil)
	c.Assert(frame.Payload, DeepEquals, []uint8{16, 0})
}
