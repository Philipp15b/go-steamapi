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

type steamid struct {
	X int
	Y int
	Z int
}

func NewIDFrom32bit(i uint32) (id steamid) {
	id.Y = int(i % 2)
	id.Z = int(i / 2)
	return
}

func NewIDFrom64bit(i uint64) (id steamid) {
	i -= 0x0110000100000000
	id = NewIDFrom32bit(uint32(i))
	return
}

func NewIDFromString(s string) (id steamid, err error) {
	validid := regexp.MustCompile("STEAM_\\d:\\d:\\d{1,}")
	
	if !validid.MatchString(s) {
		err = ErrInvalidId
	}
	
	tmp := strings.Split(s, ":")
	id.X, _ = strconv.Atoi(strings.Split(tmp[0], "_")[1])
	id.Y, _ = strconv.Atoi(tmp[1])
	id.Z, _ = strconv.Atoi(tmp[2])
	return
}

func (id steamid) String() (string) {
	s := "STEAM_"
	s += strconv.Itoa(id.X) + ":"
	s += strconv.Itoa(id.Y) + ":"
	s += strconv.Itoa(id.Z)
	return s
}

func (id *steamid) As32Bit() (i uint32) {
	i = uint32(id.Z*2 + id.Y)
	return
}

func (id *steamid) As64Bit() (i uint64) {
	i = uint64(id.As32Bit()) + 0x0110000100000000
	return
}
