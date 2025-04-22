package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerSuicide(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: Warning: Suicide (.+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_SUICIDE, logsTypes.PlayerSuicide{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			Name:    strings.TrimSpace(matches[3]),
		}
	}

	return logsEvents.PLAYER_SUICIDE, nil
}
