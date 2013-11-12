package steamapi

type SteamApp struct {
	AppId uint64
	Name  string
}

type appListJson struct {
	Applist struct {
		Apps []SteamApp
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
