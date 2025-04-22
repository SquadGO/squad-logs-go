package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func deployableDamaged(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQDeployable::)?TakeDamage\(\): ([A-z0-9_]+)_C_[0-9]+: ([0-9.]+) damage attempt by causer ([A-z0-9_]+)_C_[0-9]+ instigator (.+) with damage type ([A-z0-9_]+)_C health remaining ([0-9.]+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		damage, err := strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)

		if err != nil {
			return logsEvents.DEPLOYABLE_DAMAGED, nil
		}

		healthRemaining, err := strconv.ParseFloat(strings.TrimSpace(matches[8]), 64)

		if err != nil {
			return logsEvents.DEPLOYABLE_DAMAGED, nil
		}

		return logsEvents.DEPLOYABLE_DAMAGED, logsTypes.DeployableDamaged{
			Raw:             line,
			Time:            matches[1],
			ChainID:         matches[2],
			Deployable:      matches[3],
			Damage:          damage,
			Weapon:          matches[5],
			Name:            strings.TrimSpace(matches[6]),
			DamageType:      matches[7],
			HealthRemaining: healthRemaining,
		}
	}

	return logsEvents.DEPLOYABLE_DAMAGED, nil
}
