package steamapi

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// CommunityVisibilityState contains the visibility of the user
type CommunityVisibilityState int

const (
	// Private community visibility state
	Private CommunityVisibilityState = 1
	// FriendsOnly community visibility state
	FriendsOnly CommunityVisibilityState = 2
	// Public community visibility state
	Public CommunityVisibilityState = 3
)

// PersonaState is the visibility state
type PersonaState int

const (
	// Offline persona state is also
	// used when the steam user has set his profile
	// to private.
	Offline PersonaState = iota

	// Online is online
	Online
	// Busy is busy
	Busy
	// Away is away
	Away
	// Snooze is sniooze
	Snooze
	// LookingToTrade is looking to trade
	LookingToTrade
	// LookingToPlay is looking ot play
	LookingToPlay
)

// PlayerSummary gives an overall state of the user in steam community
type PlayerSummary struct {
	SteamID                  uint64 `json:",string"`
	CommunityVisibilityState CommunityVisibilityState
	ProfileURL               string

	ProfileState int // Set to 1 if the player has configured the profile.
	PersonaName  string
	LastLogoff   int64
	PersonaState PersonaState

	SmallAvatarURL  string `json:"avatar"`       // 32x32
	MediumAvatarURL string `json:"avatarmedium"` // 64x64
	LargeAvatarURL  string `json:"avatarfull"`   // 184x184

	TimeCreated   int64  `json:",omitempty"`
	RealName      string `json:",omitempty"`
	GameExtraInfo string `json:",omitempty"`

	PrimaryClanID uint64 `json:",string,omitempty"`
	GameID uint64 `json:",string,omitempty"`
	GameServerIp string `json:",omitempty"`
}

type playerSummaryJSON struct {
	Response struct {
		Players []PlayerSummary
	}
}

// GetPlayerSummaries Fetches the player summaries for the given Steam Ids.
func GetPlayerSummaries(ids []uint64, apiKey string) ([]PlayerSummary, error) {
	var getPlayerSummaries = NewSteamMethod("ISteamUser", "GetPlayerSummaries", 2)
	strIds := make([]string, len(ids))
	for _, id := range ids {
		strIds = append(strIds, strconv.FormatUint(id, 10))
	}
	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("steamids", strings.Join(strIds, ","))

	var resp playerSummaryJSON
	err := getPlayerSummaries.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Response.Players, nil
}

// ResolveVanityURLResponse resolves the response from steam
type ResolveVanityURLResponse struct {
	Success int
	SteamID uint64 `json:",omitempty,string"`
	Message string `json:",omitempty"`
}

// ResolveVanityURL should return a response
func ResolveVanityURL(vanityURL string, apiKey string) (*ResolveVanityURLResponse, error) {
	var resolveVanityURL = NewSteamMethod("ISteamUser", "ResolveVanityURL", 1)
	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("vanityURL", vanityURL)

	var resp struct {
		Response ResolveVanityURLResponse
	}
	err := resolveVanityURL.Request(data, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Response.Success != 1 {
		err = errors.New(resp.Response.Message)
	}
	return &resp.Response, err
}
