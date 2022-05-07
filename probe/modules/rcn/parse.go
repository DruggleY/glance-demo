package rcn

import (
	"bitbucket.org/struggle888/glance/probe"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
)

var pingReg = regexp.MustCompile(`Success rate is \d+ percent \((\d)/(\d)\)(, round-trip min/avg/max = (\d+)/(\d+)/(\d+) ms)?`)

func parseRcnPingResponse(b []byte) interface{} {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		log.Printf("err in parse rcn response to html: %v", err)
		return nil
	}
	text := doc.Find("pre").Text()
	match := pingReg.FindAllStringSubmatch(text, -1)
	if len(match) > 0 {
		if len(match[0]) == 7 {
			return probe.PingResult{
				SendAmount:    mustInt(match[0][2]),
				ReceiveAmount: mustInt(match[0][1]),
				MaxDelay:      mustInt(match[0][6]),
				MinDelay:      mustInt(match[0][4]),
				AvaDelay:      mustInt(match[0][5]),
			}
		}
	}
	log.Printf("error parsing rcn ping result, not enough match: %s", match)
	return nil
}

func parseRcnBGPResponse(b []byte) interface{} {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		log.Printf("err in parse rcn response to html: %v", err)
		return nil
	}
	return doc.Find("pre").Text()
}

var traceReg = regexp.MustCompile(` (\d{1,2}) +((([\da-zA-Z\-.]+\.[a-zA-z]+ \()?(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\)? (\[.*?] )?(\d{1,4}) msec)|\*)\s+(((([\da-zA-Z\-.]+\.[a-zA-z]+ \()?(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\)? (\[.*?] )?)?(\d{1,4}) msec)|\*)\s+(((([\da-zA-Z\-.]+\.[a-zA-z]+ \()?(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\)? (\[.*?] )?)?(\d{1,4}) msec)|\*)`)

func parseRcnTraceroute(b []byte) interface{} {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		log.Printf("err in parse rcn response to html: %v", err)
		return nil
	}
	text := doc.Find("pre").Text()
	var result probe.TraceRouteResult
	for index, jump := range traceReg.FindAllStringSubmatch(text, -1) {
		if index != mustInt(jump[1])-1 {
			log.Printf("error traceroute index: %s", text)
			return nil
		}
		var j []probe.IpDelay
		lastIp := ""
		if jump[2] == "*" {
			j = append(j, probe.IpDelay{})
		} else {
			j = append(j, probe.IpDelay{
				IP:      jump[5],
				DelayMS: mustInt(jump[7]),
			})
			lastIp = jump[5]
		}
		if jump[8] == "*" {
			j = append(j, probe.IpDelay{})
		} else {
			if jump[12] == "" {
				jump[12] = lastIp
			}
			j = append(j, probe.IpDelay{
				IP:      jump[12],
				DelayMS: mustInt(jump[14]),
			})
			lastIp = jump[12]
		}
		if jump[15] == "*" {
			j = append(j, probe.IpDelay{})
		} else {
			if jump[19] == "" {
				jump[19] = lastIp
			}
			j = append(j, probe.IpDelay{
				IP:      jump[19],
				DelayMS: mustInt(jump[21]),
			})
			lastIp = jump[19]
		}
		result = append(result, j)
	}
	return result
}
