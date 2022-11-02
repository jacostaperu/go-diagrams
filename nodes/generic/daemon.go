package generic

import "github.com/blushft/go-diagrams/diagram"

type daemonContainer struct {
	path string
	opts []diagram.NodeOption
}

var Daemon = &computeContainer{
	opts: diagram.OptionSet{diagram.Provider("generic"), diagram.NodeShape("none")},
	path: "assets/generic/daemon",
}

func (c *computeContainer) FreeRadius(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/generic/daemon/freeradius.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
