package steamapi

import (
	"net"
	"net/url"
)

type RegionId int

const (
	World RegionId = 255

	USEast RegionId = iota
	USWest
	SouthAmerica
	Europe
	Asia
	Australia
	MiddleEast
	Africa
)

type serverInfoJson struct {
	Response struct {
		Success bool
		Servers []ServerInfo
	}
}

type ServerInfo struct {
	Addr string

	// Seems to always be 65534
	// TODO: find out meaning of value
	Gmsindex   uint32
	Message    string
	AppId      uint64
	GameDir    string
	Region     RegionId
	VACSecured bool `json:"secure"`
	LANOnly    bool `json:"lan"`
	GamePort   uint16
	SpecPort   uint16
}

func GetServerInfo(ip net.IP) ([]ServerInfo, error) {
	getServerInfo := NewSteamMethod("ISteamApps", "GetServersAtAddress", 1)

	vals := url.Values{}
	vals.Add("addr", ip.String())

	var resp serverInfoJson
	err := getServerInfo.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Response.Servers, nil
}
