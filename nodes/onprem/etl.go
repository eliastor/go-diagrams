package onprem

import "github.com/blushft/go-diagrams/diagram"

type etlContainer struct {
	path string
	opts []diagram.NodeOption
}

var Etl = &etlContainer{
	opts: diagram.OptionSet{diagram.Provider("onprem"), diagram.NodeShape("none")},
	path: "assets/onprem/etl",
}

func (c *etlContainer) Embulk(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/onprem/etl/embulk.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
