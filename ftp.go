package logs

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/SquadGO/squad-logs-go/internal/parser"
	"github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/iamalone98/eventEmitter"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type FTPReaderConfig struct {
	Host               string
	Username           string
	Password           string
	LogPath            string
	AdminsPath         string
	LogEnabled         bool
	AutoReconnect      bool
	AutoReconnectDelay int
}

type fptReader struct {
	Emitter            eventEmitter.EventEmitter
	sshClient          *ssh.Client
	sftpClient         *sftp.Client
	host               string
	username           string
	password           string
	logPath            string
	adminsPath         string
	autoReconnect      bool
	autoReconnectDelay int
	connected          bool
	reconnecting       bool
}

func (fr *fptReader) Close() {
	if fr.sftpClient != nil {
		fr.sftpClient.Close()
	}

	if fr.sshClient != nil {
		fr.sshClient.Close()
	}
}

func (fr *fptReader) connect() error {
	config := &ssh.ClientConfig{
		User: fr.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(fr.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshClient, err := ssh.Dial("tcp", fr.host, config)
	if err != nil {
		fr.Emitter.Emit(logsEvents.ERROR, true)
		return fmt.Errorf("SSH Connection error: %v", err)
	}
	fr.sshClient = sshClient

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		fr.Emitter.Emit(logsEvents.ERROR, true)
		return fmt.Errorf("SFTP Connection error: %v", err)
	}
	fr.sftpClient = sftpClient
	fr.connected = true
	fr.reconnecting = false

	fr.Emitter.Emit(logsEvents.CONNECTED, true)

	go func() {
		fr.readLogFile()
	}()

	return nil
}

func (fr *fptReader) reconnect() {
	ticker := time.NewTicker(time.Duration(fr.autoReconnectDelay) * time.Second)
	go func() {
	loop:
		for {
			select {
			case <-ticker.C:
				if fr.connected {
					break loop
				}

				fr.Emitter.Emit(logsEvents.RECONNECTING, true)
				fr.reconnecting = true
				fr.connect()
			}
		}
	}()
}

func (fr *fptReader) readLogFile() {
	file, err := fr.sftpClient.Open(fr.logPath)
	if err != nil {
		fr.Emitter.Emit(logsEvents.ERROR, err)
		return
	}

	fileStat, err := file.Stat()
	if err != nil {
		fr.Emitter.Emit(logsEvents.ERROR, err)
		return
	}

	lastSize := fileStat.Size()
	_, err = file.Seek(lastSize, io.SeekStart)
	if err != nil {
		fr.Emitter.Emit(logsEvents.ERROR, err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		newStat, err := file.Stat()
		if err != nil {
			fr.Emitter.Emit(logsEvents.ERROR, err)
			return
		}

		if newStat.Size() < lastSize {
			lastSize = 0
			_, err = file.Seek(lastSize, io.SeekStart)
			if err != nil {
				fr.Emitter.Emit(logsEvents.ERROR, err)
				return
			}
		}

		if newStat.Size() > lastSize {
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					return
				}

				parser.LogParser(line, fr.Emitter)
				lastSize += int64(len(line))
			}
		}
	}
}
