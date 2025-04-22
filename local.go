package logs

import (
	"bufio"
	"io"
	"os"

	"github.com/SquadGO/squad-logs-go/internal/parser"
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/iamalone98/eventEmitter"
)

type LocalReaderConfig struct {
	LogPath            string
	AdminsPath         string
	AutoReconnect      bool
	AutoReconnectDelay int
}

type localReader struct {
	Emitter    eventEmitter.EventEmitter
	logPath    string
	adminsPath string
	logFile    *os.File
}

func (lr *localReader) Close() {
	if lr.logFile != nil {
		lr.logFile.Close()
	}
}

func (lr *localReader) connect() error {
	file, err := os.Open(lr.logPath)
	if err != nil {
		lr.Emitter.Emit(logsEvents.ERROR, err)
		return err
	}

	lr.logFile = file

	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		lr.Emitter.Emit(logsEvents.ERROR, err)
		return err
	}

	lr.Emitter.Emit(logsEvents.CONNECTED, true)

	reader := bufio.NewReader(file)

	go func() {
		for {
			line, _ := reader.ReadString('\n')

			parser.LogParser(line, lr.Emitter)
		}
	}()

	return nil
}
