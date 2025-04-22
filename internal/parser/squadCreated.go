package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func squadCreated(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: (.+) \(Online IDs: EOS: ([0-9a-f]{32}) steam: (\d{17})\) has created Squad (\d+) \(Squad Name: (.+)\) on (.+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.SQUAD_CREATED, logsTypes.SquadCreated{
			Raw:       line,
			Time:      matches[1],
			ChainID:   matches[2],
			Name:      strings.TrimSpace(matches[3]),
			EosID:     matches[4],
			SteamID:   matches[5],
			SquadID:   matches[6],
			SquadName: matches[7],
			TeamName:  matches[8],
		}
	}

	return logsEvents.SQUAD_CREATED, nil
}
