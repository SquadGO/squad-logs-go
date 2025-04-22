package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func newGame(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogWorld: Bringing World \/([A-z]+)\/(?:Maps\/)?([A-z0-9-]+)\/(?:.+\/)?([A-z0-9-]+)(?:\.[A-z0-9-]+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.NEW_GAME, logsTypes.NewGame{
			Raw:            line,
			Time:           matches[1],
			ChainID:        matches[2],
			Dlc:            matches[3],
			MapClassname:   matches[4],
			LayerClassname: matches[5],
		}
	}

	return logsEvents.NEW_GAME, nil
}
