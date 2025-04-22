package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func playerWounded(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQSoldier::)?Wound\(\): Player:(.+) KillingDamage=(?:-)*([0-9.]+) from ([A-z_0-9]+) \(Online IDs: EOS: ([\w\d]{32}) steam: (\d{17}) \| Controller ID: ([\w\d]+)\) caused by ([A-z_0-9-]+)_C`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		damage, err := strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
		if err != nil {
			return logsEvents.PLAYER_WOUNDED, nil
		}

		return logsEvents.PLAYER_WOUNDED, logsTypes.PlayerWounded{
			Raw:                      line,
			Time:                     matches[1],
			ChainID:                  matches[2],
			VictimName:               strings.TrimSpace(matches[3]),
			Damage:                   damage,
			AttackerPlayerController: matches[5],
			AttackerEOSID:            matches[6],
			AttackerSteamID:          matches[7],
			// matches[8] === matches[5]
			Weapon: matches[9],
		}
	}

	return logsEvents.PLAYER_WOUNDED, nil
}
