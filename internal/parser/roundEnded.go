package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func roundEnded(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogGameState: Match State Changed from InProgress to WaitingPostMatch`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.ROUND_ENDED, logsTypes.RoundEnded{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
		}
	}

	return logsEvents.ROUND_ENDED, nil
}
