package dota

type DotaGameMode int

const (
	AnyMode = -1

	AllPick DotaGameMode = iota
	SingleDraft
	AllRandom
	RandomDraft
	CaptainsDraft
	CaptainsMode
	DeathMode
	Diretide
	ReverseCaptainsMode
	TheGreeviling
	TutorialGame
	MidOnly
	LeastPlayed
	NewPlayerPool
	CompendiumMatchmaking
)

type DotaSkill uint

const (
	AnySkill DotaSkill = iota
	Normal
	High
	VeryHigh
)

type DotaLeaverStatus uint

const (
	None DotaLeaverStatus = iota
	Disconnected
	DisconnectedTooLong
	Abandoned
	AFK
	NeverConnected
	NeverConnectedTooLong
)

type DotaLobbyType int

const (
	Invalid DotaLobbyType = -1

	PublicMatchMaking DotaLobbyType = iota
	Practice
	Tournament
	Tutorial
	Coop
	TeamMatch
	SoloQueue
)

type DotaPlayerSlot uint8

func (d DotaPlayerSlot) IsDire() bool {
	if d&(1<<7) > 0 {
		return true
	}
	return false
}

func (d DotaPlayerSlot) GetPosition() (p uint) {
	p = uint(d & ((1 << 7) - 1))
	return
}

type DotaTeam uint

const (
	Radiant DotaTeam = iota
	Dire
)

// TODO: add methods that read information from bits
type DotaTowerStatus uint16

// TODO: add methods that read information from bits
type DotaBarracksStatus uint16
