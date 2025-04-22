package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerRevived(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: (.+) \(Online IDs: EOS: ([0-9a-f]{32}) steam: (\d{17})\) has revived (.+) \(Online IDs: EOS: ([0-9a-f]{32}) steam: (\d{17})\)\.`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_REVIVED, logsTypes.PlayerRevived{
			Raw:            line,
			Time:           matches[1],
			ChainID:        matches[2],
			ReviverName:    strings.TrimSpace(matches[3]),
			ReviverEOSID:   matches[4],
			ReviverSteamID: matches[5],
			VictimName:     strings.TrimSpace(matches[6]),
			VictimEOSID:    matches[7],
			VictimSteamID:  matches[8],
		}
	}

	return logsEvents.PLAYER_REVIVED, nil
}
