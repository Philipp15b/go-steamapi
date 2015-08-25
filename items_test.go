package steamapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMockOkGetPlayerItems(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetMockOKGetPlayerItems())
	}))
	defer ts.Close()

	appID := uint64(2)
	steamID := uint64(1234)
	apiKey := "123"

	expectedPlayerItems := &Inventory{
		Status:        uint(1),
		BackpackSlots: 840,
		Items: []Item{
			Item{
				ID:                uint64(1234567894),
				OriginalID:        uint64(123456),
				Defindex:          5470,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       false,
				Uncraftable:       false,
				InventoryToken:    uint64(19),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes:        []Attribute(nil),
				Equipped: []EquipInfo{
					EquipInfo{
						Class: 23,
						Slot:  6,
					},
				},
			},
			Item{
				ID:                uint64(1234567897),
				OriginalID:        uint64(1234567897),
				Defindex:          5508,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       true,
				Uncraftable:       false,
				InventoryToken:    uint64(25),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes: []Attribute{
					Attribute{
						Defindex:    8,
						Value:       1049511890,
						FloatValue:  0.27789169549942017,
						AccountInfo: (*AccountInfo)(nil),
					},
				},
			},
		},
	}

	playerItems, err := GetPlayerItems(ts.URL, steamID, appID, apiKey)
	if err != nil {
		t.Errorf("GetPlayerItems failure: %v", err)
		return
	}

	// Is result marshallable?
	_, err = json.Marshal(playerItems)
	if err != nil {
		t.Errorf("GetPlayerItems result marshalling failure: %v", err)
		return
	}

	// Need a better test
	if playerItems.Status != expectedPlayerItems.Status ||
		playerItems.BackpackSlots != expectedPlayerItems.BackpackSlots ||
		len(playerItems.Items) != len(expectedPlayerItems.Items) ||
		playerItems.Items[1].Attributes[0].FloatValue != 0.27789169549942017 {

		t.Errorf("GetPlayerItems(%v, %v, %v, %v) == %#v, expected %#v",
			ts.URL, steamID, appID, apiKey, playerItems, expectedPlayerItems)
	}

}
