package steamapi

import (
	"net/url"
	"strconv"
)

// Relationship is a type of relationship
type Relationship string

const (
	// All is a type of relationship
	All Relationship = "all"
	// Friend is a type of relationship
	Friend Relationship = "friend"
)

// SteamFriend is a relationship between two steam users
type SteamFriend struct {
	SteamID      uint64 `json:",string"`
	Relationship Relationship
	FriendSince  int64 `json:"friend_since"`
}

type playerFriendsListJSON struct {
	Friendslist *struct {
		Friends []SteamFriend
	}
}

// GetFriendsList Fetches the friends of the given steam id and returns the result.
//
// It returns nil if the profile is private or if there were no friends
// found for the given relationship. In either one of both cases, no error
// is returned.
func GetFriendsList(steamID uint64, filter Relationship, apiKey string) ([]SteamFriend, error) {

	var getFriendsList = NewSteamMethod("ISteamUser", "GetFriendList", 1)

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamid", strconv.FormatUint(steamID, 10))
	data.Add("relationship", string(filter))

	var resp playerFriendsListJSON
	err := getFriendsList.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Friendslist == nil {
		return nil, nil
	}
	return resp.Friendslist.Friends, nil
}
