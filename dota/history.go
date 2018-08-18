package dota

import (
	"github.com/fasmat/go-steamapi"
	"net/url"
	"strconv"
	"time"
)

type MatchFilter struct {
	PlayerName       string
	HeroId           DotaHero
	Skill            DotaSkill
	DateMin          time.Time
	DateMax          time.Time
	MinPlayers       uint
	AccountId        uint32
	LeagueId         uint32
	StartAtMatchId   uint64
	MatchesRequested uint
}

type matchHistoryJson struct {
	Result HistoryResult
}

type HistoryResult struct {
	Status           uint
	StatusDetail     string `json:",omitempty"`
	NumResults       uint   `json:"num_results"`
	TotalResults     uint   `json:"total_results"`
	ResultsRemaining uint   `json:"results_remaining"`
	Matches          []Match
}

type Match struct {
	MatchId         uint64        `json:"match_id"`
	MatchSequenceNo uint          `json:"match_seq_num"`
	MatchStart      uint          `json:"start_time"`
	LobbyType       DotaLobbyType `json:"lobby_type"`
	Players         []PlayerSummary
}

type PlayerSummary struct {
	AccountId  uint32         `json:"account_id"`
	PlayerSlot DotaPlayerSlot `json:"player_slot"`
	HeroId     uint           `json:"hero_id"`
}

func GetMatchHistory(filter MatchFilter, gameMode DotaGameMode, app int, apiKey string) ([]Match, error) {
	getMatchHistory := steamapi.NewSteamMethod("IDOTA2Match_"+strconv.Itoa(app), "GetMatchHistory", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)

	// Add vals to url if any filter value was set 
	if gameMode != AnyMode {
		vals.Add("game_mode", strconv.FormatUint(uint64(gameMode), 10))
	}
	if len(filter.PlayerName) > 0 {
		vals.Add("player_name", filter.PlayerName)
	}
	if filter.HeroId != 0 {
		vals.Add("hero_id", strconv.FormatUint(uint64(filter.HeroId), 10))
	}
	if filter.Skill != 0 {
		vals.Add("skill", strconv.FormatUint(uint64(filter.Skill), 10))
	}
	if !filter.DateMin.IsZero() {
		vals.Add("date_min", strconv.FormatInt(filter.DateMin.Unix(), 10))
	}
	if !filter.DateMax.IsZero() {
		vals.Add("date_min", strconv.FormatInt(filter.DateMax.Unix(), 10))
	}
	if filter.MinPlayers != 0 {
		vals.Add("min_players", strconv.FormatUint(uint64(filter.MinPlayers), 10))
	}
	if filter.AccountId != 0 {
		vals.Add("account_id", strconv.FormatUint(uint64(filter.AccountId), 10))
	}
	if filter.LeagueId != 0 {
		vals.Add("league_id", strconv.FormatUint(uint64(filter.LeagueId), 10))
	}
	if filter.StartAtMatchId != 0 {
		vals.Add("start_at_match_id", strconv.FormatUint(filter.StartAtMatchId, 10))
	}
	if filter.MatchesRequested > 0 {
		vals.Add("matches_requested", strconv.FormatUint(uint64(filter.MatchesRequested), 10))
	}

	var resp matchHistoryJson
	err := getMatchHistory.Request(vals, &resp)
	if err != nil {
		return nil, err
	}

	m := make([]Match, 0)
	m = append(m, resp.Result.Matches...)
	if filter.MatchesRequested > 0 {
		filter.MatchesRequested -= resp.Result.NumResults
		if filter.MatchesRequested == 0 {
			// Fetched as many as requested but less then available
			return m, nil
		}
	}

	if resp.Result.ResultsRemaining == 0 {
		// Fetched less as requested but all available
		return m, nil
	}

	// Fetched less as requested and there are still available
	// Start one match earlier than the earliest already available
	filter.StartAtMatchId = m[len(m)-1].MatchId - 1
	m2, err := GetMatchHistory(filter, gameMode, app, apiKey)
	if err != nil {
		return nil, err
	}
	m = append(m, m2...)
	return m, nil
}
