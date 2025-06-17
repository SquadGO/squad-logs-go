package parser

import (
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
	"regexp"
	"strings"
)

func beaconHandshake(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.\-:]+)\]\[([0-9]+)\]LogBeacon: Handshake complete for (.+)!$`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.BEACON_HANDSHAKE, logsTypes.BeaconHandshake{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			Client:  strings.TrimSpace(matches[3]),
		}
	}

	return logsEvents.BEACON_HANDSHAKE, nil
}
