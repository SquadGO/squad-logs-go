package parser

import (
	"github.com/iamalone98/eventEmitter"
)

type Parser func(string) (event string, data interface{})

var parsers = []Parser{
	adminBroadcast,
	applyExplosiveDamage,
	deployableDamaged,
	newGame,
	playerConnected,
	playerDamaged,
	playerDied,
	playerDisconnected,
	playerPossess,
	playerRevived,
	playerSuicide,
	playerUnpossess,
	playerWounded,
	roundEnded,
	roundTickets,
	roundWinner,
	serverTickrate,
	squadCreated,
	vehicleDamaged,
	uchannelClose,
	beaconHandshake,
}

func LogParser(line string, emitter eventEmitter.EventEmitter) {
	for _, fn := range parsers {
		event, data := fn(line)

		if data != nil {
			emitter.Emit(event, data)
			break
		}
	}
}
