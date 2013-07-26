package steamapi

import (
	"net/url"
	"strconv"
)

type schemaJson struct {
	Result Schema
}

// A game schema.
type Schema struct {
	Status                               int
	FullSchemaUrl                        string `json:"items_game_url"`
	Qualities                            map[string]int
	QualityNames                         map[string]string
	OriginNames                          []Origin
	Items                                []SchemaItem
	Attributes                           []SchemaAttribute
	ItemSets                             []ItemSet
	AttributeControlledAttachedParticles []ParticleEffect     `json:"attribute_controlled_attached_particles"`
	ItemLevels                           []ItemRankSet        `json:"item_levels"`
	KillEaterScoreTypes                  []KillEaterScoreType `json:"kill_eater_score_types"`
}

// Finds an item by its defindex in a Schema.
func (s *Schema) Item(defindex int) *SchemaItem {
	for _, item := range s.Items {
		if item.Defindex == defindex {
			return &item
		}
	}
	return nil
}

// A Schema description of an Item.
type SchemaItem struct {
	Name               string
	Defindex           int
	ItemClass          string                `json:"item_class"`
	ItemTypeName       string                `json:"item_type_name"`
	ItemName           string                `json:"item_name"`
	Description        string                `json:"item_description,omitempty"`
	ProperName         bool                  `json:"proper_name"`
	Slot               string                `json:"item_slot"`
	DefaultQuality     int                   `json:"item_quality"`
	InventoryImage     *string               `json:"image_inventory"` // this is null for the "Free Trial Premium Upgrade"
	ImageUrl           string                `json:"image_url"`
	ImageUrlLarge      string                `json:"image_url_large"`
	DropType           string                `json:"drop_type,omitempty"`
	ItemSet            string                `json:"item_set,omitempty"`
	HolidayRestriction string                `json:"holiday_restriction"`
	MinLevel           int                   `json:"min_ilevel"`
	MaxLevel           int                   `json:"max_ilevel"`
	CraftClass         string                `json:"craft_class,omitempty"`
	Capabilities       map[string]bool       `json:",omitempty"`
	UsedByClasses      []string              `json:"used_by_classes,omitempty"`
	ClassLoadoutSlots  map[string]string     `json:"per_class_loadout_slots,omitempty"`
	Styles             []Style               `json:",omitempty"`
	Attributes         []SchemaItemAttribute `json:",omitempty"`
}

// An item origin as defined in the Schema.
type Origin struct {
	Id   int `json:"origin"`
	Name string
}

// An item style
type Style struct {
	Name string
}

type SchemaItemAttribute struct {
	Name  string
	Class string
	Value float64
}

type SchemaAttribute struct {
	Name              string
	Defindex          int
	Class             string `json:"attribute_class"`
	MinValue          float64
	MaxValue          float64
	Description       string `json:"description_string,omitempty"`
	DescriptionFormat string `json:"description_format,omitempty"`
	EffectType        string `json:"effect_type"`
	Hidden            bool
	StoredAsInteger   bool `json:"stored_as_integer"`
}

type ItemSet struct {
	InternalName string `json:"item_set"`
	Name         string
	StoreBundle  string `json:"store_bundle,omitempty"`
	Items        []string
	Attributes   []SchemaItemAttribute `json:",omitempty"`
}

type ParticleEffect struct {
	System           string
	Id               int
	AttachToRootbone bool   `json:"attach_to_rootbone"`
	Attachment       string `json:",omitempty"`
	Name             string
}

type ItemRankSet struct {
	Name   string
	Levels []ItemRank
}

type ItemRank struct {
	Level         int
	RequiredScore int `json:"required_score"`
	Name          string
}

type KillEaterScoreType struct {
	Id   int    `json:"type"`
	Name string `json:"type_name"`
}

// Fetches the Schema for the given game.
func GetSchema(app int, language, apiKey string) (*Schema, error) {
	getSchema := NewSteamMethod("IEconItems_"+strconv.Itoa(app), "GetSchema", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("language", language)

	var resp schemaJson
	err := getSchema.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
