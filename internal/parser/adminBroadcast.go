package parser

import (
	"regexp"
	"strings"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func adminBroadcast(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquad: ADMIN COMMAND: Message broadcasted <(.+)> from (.+)`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		return logsEvents.ADMIN_BROADCAST, logsTypes.AdminBroadcast{
			Raw:     line,
			Time:    matches[1],
			ChainID: matches[2],
			Message: matches[3],
			From:    strings.TrimSpace(matches[4]),
		}
	}

	return logsEvents.ADMIN_BROADCAST, nil
}
