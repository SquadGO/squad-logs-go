package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerQueued(line string) (event string, data interface{}) {
	re := regexp.MustCompile(`^\[([0-9.:\-]+)]\[([ 0-9]*)]LogBeacon: Beacon Join SQJoinBeaconClient RedpointEOS:([0-9a-f]{32})`)
	matches := re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_QUEUED, logsTypes.PlayerQueued{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			EosID:   matches[3],
		}
	}

	return logsEvents.PLAYER_QUEUED, nil
}
