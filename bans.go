package steamapi

import (
	"net/url"
	"strconv"
	"strings"
)

type playerBansJSON struct {
	Players []PlayerBan
}

// PlayerBan contains all ban status for community, VAC and economy
type PlayerBan struct {
	SteamID          uint64 `json:"SteamId,string"`
	CommunityBanned  bool
	VACBanned        bool
	EconomyBan       string
	NumberOfVACBans  uint
	DaysSinceLastBan uint
	NumberOfGameBans uint
}

// GetPlayerBans takes a list of steamIDs and returns PlayerBan slice
func GetPlayerBans(steamIDs []uint64, apiKey string) ([]PlayerBan, error) {
	var getPlayerBans = NewSteamMethod("ISteamUser", "GetPlayerBans", 1)
	strSteamIDs := make([]string, len(steamIDs))
	for _, id := range steamIDs {
		strSteamIDs = append(strSteamIDs, strconv.FormatUint(id, 10))
	}

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamids", strings.Join(strSteamIDs, ","))

	var resp playerBansJSON
	err := getPlayerBans.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Players, nil
}
