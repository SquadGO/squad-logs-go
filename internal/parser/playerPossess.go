package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerPossess(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQPlayerController::)?OnPossess\(\): PC=(.+) \(Online IDs: EOS: ([\w\d]{32}) steam: (\d{17})\) Pawn=([A-z0-9_]+)_C`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_POSSESS, logsTypes.PlayerPossess{
			Raw:              line,
			Time:             matches[1],
			ChainID:          matches[2],
			Name:             strings.TrimSpace(matches[3]),
			EosID:            matches[4],
			SteamID:          matches[5],
			PossessClassname: matches[6],
		}
	}

	return logsEvents.PLAYER_POSSESS, nil
}
