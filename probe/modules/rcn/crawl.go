// Package rcn
// Update At: 2022-05-07
// By Druggle
package rcn

import (
	"bitbucket.org/struggle888/glance/probe"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Module is reserved, meaningless, but must exist
type Module struct{}

// rcnRouter is node of this lg
type rcnRouter struct {
	meta *probe.NodeMeta
	name string
}

var functionMap = map[probe.Function]string{
	probe.FunctionPing:       "ping",
	probe.FunctionTraceroute: "trace",
	probe.FunctionShowBGP:    "prefix",
}

var parseFunctionMap = map[probe.Function]func([]byte) interface{}{
	probe.FunctionPing:       parseRcnPingResponse,
	probe.FunctionTraceroute: parseRcnTraceroute,
	probe.FunctionShowBGP:    parseRcnBGPResponse,
}

func (n *rcnRouter) Meta() *probe.NodeMeta {
	return n.meta
}

// Call is node core function
func (n *rcnRouter) Call(function probe.Function, Argument string) interface{} {
	param, ok := functionMap[function]
	if !ok {
		log.Printf("unsupported function: %d\n", function)
		return nil
	}
	payload := url.Values{
		"args":   {Argument},
		"router": {n.name},
		"submit": {"Submit"},
		"query":  {param},
	}
	resp, err := http.PostForm("http://lg.rcn.net/lg.cgi", payload)
	if err != nil {
		log.Printf("error in query rcn: %v", err)
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error in parse rcn response: %v", err)
		return nil
	}
	_ = resp.Body.Close()
	return parseFunctionMap[function](b)
}

func init() {
	var nameCityMap = map[string]string{
		"Allentown PA:cisco-xr":    "Allentown",
		"Chicago IL:cisco-xr":      "Chicago",
		"Washington DC:cisco-xr":   "Washington",
		"New York NY:cisco-xr":     "New York",
		"Philadelphia PA:cisco-xr": "Philadelphia",
		"Boston MA:cisco-xr":       "Boston",
	}
	for name, city := range nameCityMap {
		probe.Crawler.AddNode(&rcnRouter{
			name: name,
			meta: &probe.NodeMeta{
				Location: probe.NodeLocation{
					Continent: probe.ContinentNorthAmerica,
					Country:   "United States",
					City:      city,
				},
				Url:                "http://lg.rcn.net",
				SupportedFunctions: probe.FunctionTraceroute | probe.FunctionPing | probe.FunctionShowBGP,
				AS:                 6079,
				Domain:             "",
				IP:                 nil,
			},
		})
	}

}

func mustInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
