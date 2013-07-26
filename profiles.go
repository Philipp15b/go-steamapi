package steamapi

import (
	"net/url"
	"strconv"
	"strings"
)

type CommunityVisibilityState int

const (
	Private     CommunityVisibilityState = 1
	FriendsOnly                          = 2
	Public                               = 3
)

type PersonaState int

const (
	// The offline persona state is also
	// used when the steam user has set his profile
	// to private.
	Offline PersonaState = iota

	Online
	Busy
	Away
	Snooze
	LookingToTrade
	LookingToPlay
)

type PlayerSummary struct {
	SteamId                  uint64 `json:",string"`
	CommunityVisibilityState CommunityVisibilityState
	ProfileUrl               string

	ProfileState int // Set to 1 if the player has configured the profile.
	PersonaName  string
	LastLogoff   int64
	PersonaState PersonaState

	SmallAvatarUrl  string `json:"avatar"`       // 32x32
	MediumAvatarUrl string `json:"avatarmedium"` // 64x64
	LargeAvatarUrl  string `json:"avatarfull"`   // 184x184

	TimeCreated   int64  `json:",omitempty"`
	RealName      string `json:",omitempty"`
	PrimaryClanId uint64 `json:",string,omitempty"`
}

type playerSummaryJson struct {
	Response struct {
		Players []PlayerSummary
	}
}

var getPlayerSummaries = NewSteamMethod("ISteamUser", "GetPlayerSummaries", 2)

// Fetches the player summaries for the given Steam Ids.
func GetPlayerSummaries(ids []uint64, apiKey string) ([]PlayerSummary, error) {
	strIds := make([]string, len(ids))
	for _, id := range ids {
		strIds = append(strIds, strconv.FormatUint(id, 10))
	}
	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("steamids", strings.Join(strIds, ","))

	var resp playerSummaryJson
	err := getPlayerSummaries.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Response.Players, nil
}

var resolveVanityUrl = NewSteamMethod("ISteamUser", "ResolveVanityUrl", 1)

type ResolveVanityUrlResponse struct {
	Success int
	SteamId uint64 `json:",omitempty,string"`
	Message string `json:",omitempty"`
}

func ResolveVanityUrl(vanityUrl string, apiKey string) (*ResolveVanityUrlResponse, error) {
	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("vanityurl", vanityUrl)

	var resp ResolveVanityUrlResponse
	err := resolveVanityUrl.Request(data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
