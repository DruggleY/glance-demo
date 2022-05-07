package probe

import "net"

//LGCrawler is Core crawler
type LGCrawler struct {
	nodes []Node
}

func (c *LGCrawler) AddNode(n Node) {
	c.nodes = append(c.nodes, n)
}

// Node define a probe node
type Node interface {
	Meta() *NodeMeta
	Call(function Function, Argument string) interface{}
}

type NodeMeta struct {
	Location           NodeLocation
	Url                string
	SupportedFunctions Function
	AS                 int
	Domain             string
	IP                 net.IP
}

type NodeLocation struct {
	Continent Continent
	Country   string
	City      string
}

/* Result */
type PingResult struct {
	SendAmount    int `json:"send_amount"`
	ReceiveAmount int `json:"receive_amount"`
	MaxDelay      int `json:"max_delay"`
	MinDelay      int `json:"min_delay"`
	AvaDelay      int `json:"ava_delay"`
}

type IpDelay struct {
	IP      string `json:"ip"`
	DelayMS int    `json:"delay_ms"`
}

type TraceRouteResult [][]IpDelay

type ShowBGPResult string
