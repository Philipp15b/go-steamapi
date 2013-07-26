package steamapi

import (
	"net/url"
	"strconv"
)

type Relationship string

const (
	All    Relationship = "all"
	Friend              = "friend"
)

type SteamFriend struct {
	SteamId      uint64 `json:",string"`
	Relationship Relationship
	FriendSince  int64 `json:"friend_since"`
}

type playerFriendsListJson struct {
	Friendslist *struct {
		Friends []SteamFriend
	}
}

var getFriendsList = NewSteamMethod("ISteamUser", "GetFriendList", 1)

// Fetches the friends of the given steam id and returns the result.
//
// It returns nil if the profile is private or if there were no friends
// found for the given relationship. In either one of both cases, no error
// is returned.
func GetFriendsList(id uint64, filter Relationship, apiKey string) ([]SteamFriend, error) {
	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamid", strconv.FormatUint(id, 10))
	data.Add("relationship", string(filter))

	var resp playerFriendsListJson
	err := getFriendsList.Request(data, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Friendslist == nil {
		return nil, nil
	}
	return resp.Friendslist.Friends, nil
}
