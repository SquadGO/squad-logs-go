package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func applyExplosiveDamage(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`/^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer](?:ASQProjectile::)?ApplyExplosiveDamage\(\): HitActor=(\S+) DamageCauser=(\S+) DamageInstigator=(\S+) ExplosionLocation=V\((X=[\d\-.]+, Y=[\d\-.]+, Z=[\d\-.]+)\)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.EXPLOSIVE_DAMAGED, logsTypes.ApplyExplosiveDamage{
			Raw:              line,
			Time:             matches[1],
			ChainID:          matches[2],
			Name:             strings.TrimSpace(matches[3]),
			Deployable:       matches[4],
			PlayerController: matches[5],
			Locations:        matches[6],
		}
	}

	return logsEvents.EXPLOSIVE_DAMAGED, nil
}
