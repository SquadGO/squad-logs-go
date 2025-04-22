package parser

import (
	"regexp"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerDisconnected(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogNet: UChannel::Close: Sending CloseBunch\. ChIndex == [0-9]+\. Name: \[UChannel\] ChIndex: [0-9]+, Closing: [0-9]+ \[UNetConnection\] RemoteAddr: ([\d.]+):[\d]+, Name: EOSIpNetConnection_[0-9]+, Driver: GameNetDriver EOSNetDriver_[0-9]+, IsServer: YES, PC: ([^ ]+PlayerController_C_[0-9]+), Owner: [^ ]+PlayerController_C_[0-9]+, UniqueId: RedpointEOS:([\d\w]+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.PLAYER_DISCONNECTED, logsTypes.PlayerDisconnected{
			Raw:              line,
			Time:             matches[1],
			ChainID:          matches[2],
			Ip:               matches[3],
			PlayerController: matches[4],
			EosID:            matches[5],
		}
	}

	return logsEvents.PLAYER_DISCONNECTED, nil
}
