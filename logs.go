package logs

import (
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/iamalone98/eventEmitter"
)

func NewFTPReader(config FTPReaderConfig) (*fptReader, error) {
	fr := &fptReader{
		Emitter:            eventEmitter.NewEventEmitter(),
		host:               config.Host,
		username:           config.Username,
		password:           config.Password,
		logPath:            config.LogPath,
		adminsPath:         config.AdminsPath,
		autoReconnect:      config.AutoReconnect,
		autoReconnectDelay: config.AutoReconnectDelay,
	}

	if err := fr.connect(); err != nil {
		return nil, err
	}

	fr.Emitter.On(logsEvents.ERROR, func(i interface{}) {
		fr.connected = false

		if fr.autoReconnect && fr.autoReconnectDelay > 0 && !fr.reconnecting {
			fr.reconnect()
		}
	})

	return fr, nil
}

func NewLocalReader(config LocalReaderConfig) (*localReader, error) {
	lr := &localReader{
		Emitter:    eventEmitter.NewEventEmitter(),
		logPath:    config.LogPath,
		adminsPath: config.AdminsPath,
	}

	if err := lr.connect(); err != nil {
		return nil, err
	}

	return lr, nil
}
