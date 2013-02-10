package steamapi

import (
	"net/url"
	"strconv"
)

type playerItemsJson struct {
	Result Inventory
}

type Inventory struct {
	Status        uint
	BackpackSlots int `json:"num_backpack_slots"`
	Items         []Item
}

type Item struct {
	Id                uint32
	OriginalId        uint32 `json:"original_id"`
	Defindex          int
	Level             int
	Quanitity         int
	Origin            int
	Untradeable       bool   `json:"flag_cannot_trade,omitempty"`
	Uncraftable       bool   `json:"flag_cannot_craft,omitempty"`
	InventoryToken    uint32 `json:",inventory"`
	Quality           int
	CustomName        string      `json:"custom_name,omitempty"`
	CustomDescription string      `json:"custom_description,omitempty"`
	Attributes        []Attribute `json:",omitempty"`
	Equipped          []EquipInfo `json:",omitempty"`
}

func (i *Item) Position() uint16 {
	return uint16(i.InventoryToken & 0xFFFF)
}

type Attribute struct {
	Defindex    int
	Value       int
	FloatValue  float64      `json:",omitempty"`
	AccountInfo *AccountInfo `json:",omitempty"`
}

type AccountInfo struct {
	SteamId     uint64 `json:",string"`
	PersonaName string
}

type EquipInfo struct {
	Class int
	Slot  int
}

// Fetches the player summaries for the given Steam Ids.
func GetPlayerItems(id uint64, app int, apiKey string) (*Inventory, error) {
	getPlayerItems := NewSteamMethod("IEconItems_"+strconv.Itoa(app), "GetPlayerItems", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("SteamId", strconv.FormatUint(id, 10))

	var resp playerItemsJson
	err := getPlayerItems.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
