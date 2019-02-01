package cluster

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestCluster(t *testing.T) { TestingT(t) }

type ClusterSuite struct{}

var _ = Suite(&ClusterSuite{})

func (s *ClusterSuite) TestClusterByName(c *C) {
	clusters := New()
	clusterId, cluster, err := ClusterByName("Basic")(clusters)
	c.Assert(err, IsNil)
	c.Assert(clusterId, Equals, Basic)
	c.Assert(cluster, NotNil)
}

func (s *ClusterSuite) TestCommandByName(c *C) {
	clusters := New()
	_, cluster, err := ClusterByID(Basic)(clusters)
	c.Assert(err, IsNil)
	commandId, command, err := CommandByName("ResetToFactoryDefaults")(cluster.CommandReceivedDescriptors)
	c.Assert(err, IsNil)
	c.Assert(commandId, Equals, uint8(0x00))
	c.Assert(command, NotNil)
	c.Assert(command.Command, DeepEquals, &ResetToFactoryDefaultsCommand{})
}

func (s *ClusterSuite) TestFunctionalFrame(c *C) {
	clusters := New()
	frame, err := clusters.FunctionalFrame(ClusterByID(Identify), ReceiveCommand(CommandByName("Identify")), uint16(16))
	c.Assert(err, IsNil)
	c.Assert(frame.Payload, DeepEquals, []uint8{16, 0})
}
