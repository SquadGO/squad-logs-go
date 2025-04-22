# SquadRcon

This library is designed for the game Squad, it will give you the ability to easily parse game logs. It is possible to read the file locally or protocol SFTP. I hope this will be useful to you!

## Install

```text
go get -u github.com/SquadGO/squad-logs-go
```

## Quick start example

```golang
import (
  "fmt"
  logs "github.com/SquadGO/squad-logs-go"
  "github.com/SquadGO/squad-logs-go/logsEvents"
	"github.com/SquadGO/squad-logs-go/logsTypes"
)

func main() {
  fr, err := logs.NewFTPReader(
		logs.FTPReaderConfig{
			Host:               "127.0.0.1:22",
			Username:           "root",
			Password:           "password",
			LogPath:            "/server/SquadGame.log",
			AdminsPath:         "/server/Admins.cfg", // TODO
			AutoReconnect:      true,
			AutoReconnectDelay: 5,
		},
	)

  if err != nil {
    fmt.Println("[LOGS] Error: ", err)
    return
  }

  defer fr.Close()

  fmt.Println("[LOGS] Connection successful")

  /* Listeners works after first initialization */

  fr.Emitter.On(logsEvents.CONNECTED, func(_ interface{}) {
    fmt.Println("[LOGS] Connection successful")
  })

  fr.Emitter.On(logsEvents.RECONNECTING, func(i interface{}) {
    fmt.Println("[LOGS] Reconnecting")
  })

  fr.Emitter.On(logsEvents.CLOSE, func(_ interface{}) {
    fmt.Println("[LOGS] Connection closed")
  })

  fr.Emitter.On(logsEvents.ERROR, func(err interface{}) {
    fmt.Println("[LOGS] Error: ", err)
  })

  fr.Emitter.On(logsEvents.PLAYER_DAMAGED, func(i interface{}) {
    if data, ok := i.(logsTypes.PlayerDamaged); ok {
      fmt.Println(data)
    }
  })

  // Use to prevent the program from ending
  select {}
}
```

## Listeners

| Listeners               | Returns              |
| ----------------------- | -------------------- |
| **CONNECTED**           | `nil`                |
| **RECONNECTING**        | `nil`                |
| **CLOSE**               | `nil`                |
| **ERROR**               | `Error`              |
| **ADMIN_BROADCAST**     | `AdminBroadcast`     |
| **DEPLOYABLE_DAMAGED**  | `DeployableDamaged`  |
| **NEW_GAME**            | `NewGame`            |
| **PLAYER_CONNECTED**    | `PlayerConnected`    |
| **PLAYER_DISCONNECTED** | `PlayerDisconnected` |
| **PLAYER_DAMAGED**      | `PlayerDamaged`      |
| **PLAYER_DIED**         | `PlayerDied`         |
| **PLAYER_POSSESS**      | `PlayerPossess`      |
| **PLAYER_UNPOSSESS**    | `PlayerUnpossess`    |
| **PLAYER_REVIVED**      | `PlayerRevived`      |
| **PLAYER_SUICIDE**      | `PlayerSuicide`      |
| **PLAYER_WOUNDED**      | `PlayerWounded`      |
| **ROUND_WINNER**        | `RoundWinner`        |
| **ROUND_ENDED**         | `RoundEnded`         |
| **ROUND_TICKETS**       | `RoundTickets`       |
| **SQUAD_CREATED**       | `SquadCreated`       |
| **VEHICLE_DAMAGED**     | `VehicleDamaged`     |
| **SERVER_TICKRATE**     | `ServerTickrate`     |
