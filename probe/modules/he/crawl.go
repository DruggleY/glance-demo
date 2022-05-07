package he

import "bitbucket.org/struggle888/glance/probe"

// Module is reserved, meaningless, but must exist
type Module struct{}

type node struct {
	meta *probe.NodeMeta
}

func (n *node) Meta() *probe.NodeMeta {
	return n.meta
}

func (n *node) Call(function probe.Function, Argument string) interface{} {
	return nil
}

func init() {
	probe.Crawler.AddNode(&node{})
}
