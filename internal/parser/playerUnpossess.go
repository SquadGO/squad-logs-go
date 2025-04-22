package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerUnpossess(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQPlayerController::)?OnUnPossess\(\): PC=(.+) \(Online IDs: EOS: ([\w\d]{32}) steam: (\d{17})\)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_UNPOSSESS, logsTypes.PlayerUnpossess{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			Name:    strings.TrimSpace(matches[3]),
			EosID:   matches[4],
			SteamID: matches[5],
		}
	}

	return logsEvents.PLAYER_UNPOSSESS, nil
}
