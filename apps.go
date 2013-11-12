package steamapi

import (
	"errors"
	"net/url"
	"strconv"
)

type SteamApp struct {
	AppId uint64
	Name  string
}

type appListJson struct {
	Applist struct {
		Apps []SteamApp
	}
}

type upToDateCheckJson struct {
	Response struct {
		Success        bool
		UpToDate       bool   `json:"up_to_date"`
		Listable       bool   `json:"version_is_listable"`
		CurrentVersion uint   `json:"required_version,omitempty"`
		Message        string `json:"omitempty"`
		Error          string `json:"omitempty"`
	}
}

func GetAppList() ([]SteamApp, error) {
	getAppList := NewSteamMethod("ISteamApps", "GetAppList", 2)

	var resp appListJson
	err := getAppList.Request(nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Applist.Apps, nil
}

func IsAppUpToDate(app int, version uint) (bool, error) {
	upToDateCheck := NewSteamMethod("ISteamApps", "UpToDateCheck", 1)

	vals := url.Values{}
	vals.Add("appid", strconv.Itoa(app))
	vals.Add("version", strconv.FormatUint(uint64(version), 10))

	var resp upToDateCheckJson
	err := upToDateCheck.Request(vals, &resp)
	if err != nil {
		return false, err
	}
	if !resp.Response.Success {
		return false, errors.New(resp.Response.Error)
	}
	return resp.Response.UpToDate, nil
}

func GetCurrentAppVersion(app int) (uint, error) {
	upToDateCheck := NewSteamMethod("ISteamApps", "UpToDateCheck", 1)

	vals := url.Values{}
	vals.Add("appid", strconv.Itoa(app))
	vals.Add("version", "1")

	var resp upToDateCheckJson
	err := upToDateCheck.Request(vals, &resp)
	if err != nil {
		return 0, err
	}
	if !resp.Response.Success {
		return 0, errors.New(resp.Response.Error)
	}
	return resp.Response.CurrentVersion, nil
}
