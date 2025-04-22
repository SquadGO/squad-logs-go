package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func serverTickrate(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: USQGameState: Server Tick Rate: ([0-9.]+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		tickrate, err := strconv.ParseFloat(strings.TrimSpace(matches[3]), 64)
		if err != nil {
			return logsEvents.SERVER_TICKRATE, nil
		}

		return logsEvents.SERVER_TICKRATE, logsTypes.ServerTickrate{
			Raw:      line,
			Time:     matches[1],
			ChainID:  matches[2],
			TickRate: tickrate,
		}
	}

	return logsEvents.SERVER_TICKRATE, nil
}
