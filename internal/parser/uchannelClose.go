package parser

import (
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
	"regexp"
	"strconv"
	"strings"
)

func uchannelClose(line string) (event string, data interface{}) {
	var re *regexp.Regexp
	var matches []string

	re = regexp.MustCompile(`^\[([0-9.\-:]+)\]\[([0-9]+)\]LogNet: UChannel::Close: Sending CloseBunch\. ChIndex == [0-9]+\. Name: \[UChannel\] ChIndex: ([0-9]+), Closing: ([0-9]+) \[UNetConnection\] RemoteAddr: ([^,]+), Name: ([^,]+), Driver: ([^,]+), IsServer: (YES|NO), PC: ([^,]+), Owner: ([^,]+), UniqueId: (.+)$`)
	matches = re.FindStringSubmatch(line)

	if matches != nil {
		chIndex, _ := strconv.Atoi(matches[3])
		closing, _ := strconv.Atoi(matches[4])
		return logsEvents.UCHANNEL_CLOSE, logsTypes.UChannelClose{
			Raw:        line,
			Time:       matches[1],
			ChainID:    matches[2],
			ChIndex:    chIndex,
			Closing:    closing,
			RemoteAddr: matches[5],
			Name:       strings.TrimSpace(matches[6]),
			Driver:     strings.TrimSpace(matches[7]),
			IsServer:   matches[8] == "YES",
			PC:         strings.TrimSpace(matches[9]),
			Owner:      strings.TrimSpace(matches[10]),
			UniqueId:   matches[11],
		}
	}

	return logsEvents.UCHANNEL_CLOSE, nil
}
