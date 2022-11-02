package f5bigip

import "github.com/blushft/go-diagrams/diagram"

type f5Container struct {
	path string
	opts []diagram.NodeOption
}

var F5 = &f5Container{
	opts: diagram.OptionSet{diagram.Provider("f5bigip"), diagram.NodeShape("none")},
	path: "assets/f5bigip/f5bigip",
}

func (c *f5Container) BigIp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/f5bigip/f5bigip/f5bigip.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
