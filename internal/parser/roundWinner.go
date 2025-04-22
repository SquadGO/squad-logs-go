package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func roundWinner(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQGameMode::)?DetermineMatchWinner\(\): (.+) won on (.+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.ROUND_WINNER, logsTypes.RoundWinner{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			Winner:  matches[3],
			Layer:   matches[4],
		}
	}

	return logsEvents.ROUND_WINNER, nil
}
