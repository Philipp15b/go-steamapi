package dota

import (
	"github.com/fasmat/go-steamapi"
	"net/url"
	"strconv"
)

type matchResultJson struct {
	Result MatchResult
}

type MatchResult struct {
	Players               []Player
	Season                uint `json:"item_0,omitempty"`
	RadiantWin            bool `json:"radiant_win"`
	Duration              uint
	MatchStart            uint               `json:"start_time"`
	MatchId               uint64             `json:"match_id"`
	MatchSequenceNo       uint               `json:"match_seq_num"`
	TowerStatusRadiant    DotaTowerStatus    `json:"tower_status_radiant"`
	TowerStatusDire       DotaTowerStatus    `json:"tower_status_dire"`
	BarracksStatusRadiant DotaBarracksStatus `json:"barracks_status_radiant"`
	BarracksStatusDire    DotaBarracksStatus `json:"barracks_status_dire"`
	Cluster               uint
	FirstBloodTime        int           `json:"first_blood_time"`
	LobbyType             DotaLobbyType `json:"lobby_type"`
	HumanPlayers          uint          `json:"human_players"`
	LeagueId              uint
	PositiveVotes         uint         `json:"positive_votes"`
	NegativeVotes         uint         `json:"negative_votes"`
	GameMode              DotaGameMode `json:"game_mode"`
	PicksBans             []PickBan    `json:"picks_bans,omitempty"`
}

type Player struct {
	AccountId     uint32         `json:"account_id"`
	PlayerSlot    DotaPlayerSlot `json:"player_slot"`
	HeroId        uint           `json:"hero_id"`
	Item0         uint           `json:"item_0"`
	Item1         uint           `json:"item_1"`
	Item2         uint           `json:"item_2"`
	Item3         uint           `json:"item_3"`
	Item4         uint           `json:"item_4"`
	Item5         uint           `json:"item_5"`
	Kills         uint
	Deaths        uint
	Assists       uint
	LeaverStatus  DotaLeaverStatus `json:"leaver_status"`
	GoldRemaining uint             `json:"gold"`
	LastHits      uint             `json:"last_hits"`
	Denies        uint             `json:"denies"`
	GPM           uint             `json:"gold_per_min"`
	XPM           uint             `json:"xp_per_min"`
	GoldSpent     uint             `json:"gold_spent"`
	HeroDamage    uint             `json:"hero_damage"`
	TowerDamage   uint             `json:"tower_damage"`
	HeroHealing   uint             `json:"hero_healing"`
	Level         uint
	Abilities     []Ability `json:"ability_upgrades"`
	Units         []Unit    `json:"additional_units,omitempty"`
}

type Ability struct {
	Id           uint `json:"ability"`
	TimeUpgraded int  `json:"time"`
	Level        uint
}

type Unit struct {
	Name  string `json:"unitname"`
	Item0 uint   `json:"item_0"`
	Item1 uint   `json:"item_1"`
	Item2 uint   `json:"item_2"`
	Item3 uint   `json:"item_3"`
	Item4 uint   `json:"item_4"`
	Item5 uint   `json:"item_5"`
}

type PickBan struct {
	IsPick    bool // picked or banned
	HeroId    uint
	ByRadiant bool // if radiant or dire chose to pick/ban
	Order     uint // the sequence in which pick/ban was done (0 - 19)
}

// Fetches statistics of a specific Match.
func GetMatchDetails(matchid uint64, app int, apiKey string) (*MatchResult, error) {
	getMatchDetails := steamapi.NewSteamMethod("IDOTA2Match_"+strconv.Itoa(app), "GetMatchDetails", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("match_id", strconv.FormatUint(matchid, 10))

	var resp matchResultJson
	err := getMatchDetails.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
