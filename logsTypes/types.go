package logsTypes

type AdminBroadcast struct {
	Raw     string
	Time    string
	ChainID string
	Message string
	From    string
}

type ApplyExplosiveDamage struct {
	Raw              string
	Time             string
	ChainID          string
	Name             string
	Deployable       string
	PlayerController string
	Locations        string
}

type DeployableDamaged struct {
	Raw             string
	Time            string
	ChainID         string
	Deployable      string
	Damage          float64
	Weapon          string
	Name            string
	DamageType      string
	HealthRemaining float64
}

type NewGame struct {
	Raw            string
	Time           string
	ChainID        string
	Dlc            string
	MapClassname   string
	LayerClassname string
}

type PlayerConnected struct {
	Raw              string
	Time             string
	ChainID          string
	PlayerController string
	Ip               string
	EosID            string
	SteamID          string
}

type PlayerDamaged struct {
	Raw                string
	Time               string
	ChainID            string
	VictimName         string
	Damage             float64
	AttackerName       string
	AttackerEOSID      string
	AttackerSteamID    string
	AttackerController string
	Weapon             string
}

type PlayerDied struct {
	Raw                string
	Time               string
	ChainID            string
	VictimName         string
	Damage             float64
	AttackerController string
	AttackerEOSID      string
	AttackerSteamID    string
	Weapon             string // nullptr or controller
}

type PlayerDisconnected struct {
	Raw              string
	Time             string
	ChainID          string
	Ip               string
	PlayerController string
	EosID            string
}

type PlayerPossess struct {
	Raw              string
	Time             string
	ChainID          string
	Name             string
	EosID            string
	SteamID          string
	PossessClassname string
}

type PlayerRevived struct {
	Raw            string
	Time           string
	ChainID        string
	ReviverName    string
	ReviverEOSID   string
	ReviverSteamID string
	VictimName     string
	VictimEOSID    string
	VictimSteamID  string
}

type PlayerSuicide struct {
	Raw     string
	Time    string
	ChainID string
	Name    string
}

type PlayerUnpossess struct {
	Raw     string
	Time    string
	ChainID string
	Name    string
	EosID   string
	SteamID string
}

type PlayerWounded struct {
	Raw                      string
	Time                     string
	ChainID                  string
	VictimName               string
	Damage                   float64
	AttackerPlayerController string
	AttackerEOSID            string
	AttackerSteamID          string
	Weapon                   string
}

type PlayerQueued struct {
	Raw     string
	Time    string
	ChainID string
	EosID   string
}

type RoundEnded struct {
	Raw     string
	Time    string
	ChainID string
}

type RoundTickets struct {
	Raw        string
	Time       string
	ChainID    string
	Team       string
	Subfaction string
	Faction    string
	Action     string
	Tickets    int
	Layer      string
	Level      string
}

type RoundWinner struct {
	Raw     string
	Time    string
	ChainID string
	Winner  string
	Layer   string
}

type ServerTickrate struct {
	Raw      string
	Time     string
	ChainID  string
	TickRate float64
}

type SquadCreated struct {
	Raw       string
	Time      string
	ChainID   string
	Name      string
	EosID     string
	SteamID   string
	SquadID   string
	SquadName string
	TeamName  string
}

type VehicleDamaged struct {
	Raw             string
	Time            string
	ChainID         string
	VictimVehicle   string
	Damage          float64
	AttackerVehicle string
	AttackerName    string
	AttackerEOSID   string
	AttackerSteamID string
	HealthRemaining string
}

type BeaconHandshake struct {
	Raw     string
	Time    string
	ChainID string
	Client  string
}
