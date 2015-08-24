// Package steamapi provides an interface to the
// Steam Web API methods.
package steamapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// BaseSteamWebURL is the steam url used to do requests
const BaseSteamAPIURL = "https://api.steampowered.com"

// A SteamMethod represents a Steam Web API method.
type SteamMethod string

// NewSteamMethod creates a new SteamMethod.
func NewSteamMethod(baseSteamAPIURL, interf, method string, version int) SteamMethod {
	m := fmt.Sprintf("%v/%v/%v/v%v/", baseSteamAPIURL, interf, method, strconv.Itoa(version))
	return SteamMethod(m)
}

// Request makes a request to the Steam Web API with the given
// url values and stores the result in v.
//
// Returns an error if the return status code was not 200.
func (s SteamMethod) Request(data url.Values, v interface{}) error {
	url := string(s)
	if data != nil {
		url += "?" + data.Encode()
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("steamapi %s Status code %d", s, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(&v)
}
