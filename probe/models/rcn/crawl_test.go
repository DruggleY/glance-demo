package rcn

import (
	"bitbucket.org/struggle888/glance/probe"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	node := &rcnRouter{
		name: "Washington DC:cisco-xr",
		meta: &probe.NodeMeta{
			Location:           probe.NodeLocation{},
			SupportedFunctions: probe.FunctionTraceroute | probe.FunctionPing | probe.FunctionShowBGP,
			AS:                 6079,
			Domain:             "",
			IP:                 nil,
		},
	}
	jsonPrint(node.Call(probe.FunctionPing, "139.162.205.205"))
	jsonPrint(node.Call(probe.FunctionPing, "222.194.15.1"))
	jsonPrint(node.Call(probe.FunctionPing, "39.106.66.97"))
	jsonPrint(node.Call(probe.FunctionTraceroute, "139.162.205.205"))
	jsonPrint(node.Call(probe.FunctionTraceroute, "222.194.15.1"))
	jsonPrint(node.Call(probe.FunctionTraceroute, "39.106.66.97"))
	fmt.Println(node.Call(probe.FunctionShowBGP, "139.162.205.205"))
	fmt.Println(node.Call(probe.FunctionShowBGP, "222.194.15.1"))
	fmt.Println(node.Call(probe.FunctionShowBGP, "39.106.66.97"))
}

func jsonPrint(item interface{}) {
	b, _ := json.Marshal(item)
	fmt.Println(string(b))
}
