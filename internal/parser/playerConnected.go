package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerConnected(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: PostLogin: NewPlayer: [^\s]+ .+PersistentLevel\.([^\s]+) \(IP: ([\d.]+) \| Online IDs: EOS: ([0-9a-f]{32}) steam: (\d+)\)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_CONNECTED, logsTypes.PlayerConnected{
			Raw:              line,
			Time:             matches[1],
			ChainID:          matches[2],
			PlayerController: matches[3],
			Ip:               matches[4],
			EosID:            matches[5],
			SteamID:          matches[6],
		}
	}

	return logsEvents.PLAYER_CONNECTED, nil
}
