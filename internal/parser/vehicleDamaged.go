package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func vehicleDamaged(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadTrace: \[DedicatedServer]ASQVehicleSeat::TraceAndMessageClient\(\): (.+): (.+) damage taken by causer (.+) instigator \(Online Ids: (.+?)\) EOS: ([0-9a-f]{32}) steam: (\d{17}) health remaining (.+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		damage, err := strconv.ParseFloat(strings.TrimSpace(matches[4]), 64)
		if err != nil {
			return logsEvents.VEHICLE_DAMAGED, nil
		}

		return logsEvents.VEHICLE_DAMAGED, logsTypes.VehicleDamaged{
			Raw:             line,
			Time:            matches[1],
			ChainID:         matches[2],
			VictimVehicle:   matches[3],
			Damage:          damage,
			AttackerVehicle: matches[5],
			AttackerName:    strings.TrimSpace(matches[6]),
			AttackerEOSID:   matches[7],
			AttackerSteamID: matches[8],
			HealthRemaining: matches[9],
		}
	}

	return logsEvents.VEHICLE_DAMAGED, nil
}
