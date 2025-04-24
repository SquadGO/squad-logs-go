package parser

import (
	"regexp"
	"strconv"

	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func roundTickets(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.:-]+)]\[([ 0-9]*)]LogSquadGameEvents: Display: Team ([0-9]), (.*) \( ?(.*?) ?\) has (won|lost) the match with ([0-9]+) Tickets on layer (.*) \(level (.*)\)!`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		tickets, err := strconv.Atoi(matches[7])
		if err != nil {
			return logsEvents.ROUND_TICKETS, nil
		}

		return logsEvents.ROUND_TICKETS, logsTypes.RoundTickets{
			Raw:        line,
			Time:       matches[1],
			ChainID:    matches[2],
			Team:       matches[3],
			Subfaction: matches[4],
			Faction:    matches[5],
			Action:     matches[6],
			Tickets:    tickets,
			Layer:      matches[8],
			Level:      matches[9],
		}
	}

	return logsEvents.ROUND_TICKETS, nil
}
