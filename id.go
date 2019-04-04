package steamapi

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidId = errors.New("Invalid Steam ID")
)

type steamId struct {
	X uint32
	Y uint32
	Z uint32
}

func NewIdFrom32bit(i uint32) (id steamId) {
	id.Y = i % 2
	id.Z = i / 2
	return
}

func NewIdFrom64bit(i uint64) (id steamId) {
	i -= 0x0110000100000000
	id = NewIdFrom32bit(uint32(i))
	return
}

func NewIdFromVanityUrl(vanityUrl, apiKey string) (id steamId, err error) {
	resp, err := ResolveVanityURL(vanityUrl, apiKey)
	if err != nil {
		return
	}

	id = NewIdFrom64bit(resp.SteamID)
	return
}

func NewIdFromString(s string) (id steamId, err error) {
	validid := regexp.MustCompile("STEAM_\\d:\\d:\\d{1,}")

	if !validid.MatchString(s) {
		err = ErrInvalidId
	}

	tmp := strings.Split(s, ":")
	tmpX, _ := strconv.ParseUint(strings.Split(tmp[0], "_")[1], 10, 32)
	tmpY, _ := strconv.ParseUint(tmp[1], 10, 32)
	tmpZ, _ := strconv.ParseUint(tmp[2], 10, 32)

	id.X = uint32(tmpX)
	id.Y = uint32(tmpY)
	id.Z = uint32(tmpZ)
	return
}

func (id steamId) String() (s string) {
	s = "STEAM_"
	s += strconv.FormatUint(uint64(id.X), 10) + ":"
	s += strconv.FormatUint(uint64(id.Y), 10) + ":"
	s += strconv.FormatUint(uint64(id.Z), 10)
	return
}

func (id steamId) As32Bit() (i uint32) {
	i = id.Z*2 + id.Y
	return
}

func (id steamId) As64Bit() (i uint64) {
	i = uint64(id.As32Bit()) + 0x0110000100000000
	return
}
