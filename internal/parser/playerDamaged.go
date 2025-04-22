package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerDamaged(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: Player:(.+) ActualDamage=([0-9.]+) from (.+) \(Online IDs: EOS: ([0-9a-f]{32}) steam: (\d{17}) \| Player Controller ID: ([^ ]+)\)caused by ([A-z_0-9-]+)_C`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		damage, err := strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
		if err != nil {
			return logsEvents.PLAYER_DAMAGED, nil
		}

		return logsEvents.PLAYER_DAMAGED, logsTypes.PlayerDamaged{
			Raw:                line,
			Time:               matches[1],
			ChainID:            matches[2],
			VictimName:         strings.TrimSpace(matches[3]),
			Damage:             damage,
			AttackerName:       strings.TrimSpace(matches[5]),
			AttackerEOSID:      matches[6],
			AttackerSteamID:    matches[7],
			AttackerController: matches[8],
			Weapon:             matches[9],
		}
	}

	return logsEvents.PLAYER_DAMAGED, nil
}
