package steamapi

// GetMockOKGetAssetClassInfo steam API gives class_id, market_hash_name...
// http://api.steampowered.com/ISteamEconomy/GetAssetClassInfo/v0001?key=XXX&format=json&language=en&appid=XXX&class_count=2&classid0=123456789
func GetMockOKGetAssetClassInfo() string {
	return `
		{
			"result": {
				"123456789": {
					"icon_url": "W_I_5GLm4wPcv9jJQ7z7tz_l_0sEIYUhRfbF4arNQkgGQGKd3kMuVpMgCwRZrhSfeEqb1qNMeO7lDgsvJYj2VkHyNb-A-UWkTe9Xc8Rgd2sbj9_ugkgSUXffBrFHXNQrvM7K0Ay7XgXDLWdun9gFgPqagJWGCPPO6UywK3ID03w",
					"icon_url_large": "W_I_5GLm4wPcv9jJQ7z7tz_l_0sEIYUhRfbF4arNQkgGQGKd3kMuVpMgCwRZrhSfeEqb1qNMeO7lDgsvJYj2VkHyNb-A-UWkTe9Xc8RgBmMYzo69mB0TByTSDb8RDYMpupzD1APoW1HCcWFun4wGivufgpfQUqHSrESyJVJuk7o-hPMuyZ4",
					"icon_drag_url": "",
					"name": "Ye Olde Pipe",
					"market_hash_name": "Ye Olde Pipe",
					"market_name": "Ye Olde Pipe",
					"name_color": "D2D2D2",
					"background_color": "",
					"type": "Common Pipe",
					"tradable": "1",
					"marketable": "1",
					"commodity": "0",
					"market_tradable_restriction": "7",
					"market_marketable_restriction": "7",
					"fraudwarnings": "",
					"descriptions": {
						"0": {
							"type": "html",
							"value": "Used By: Kunkka",
							"app_data": ""
						},
						"1": {
							"type": "html",
							"value": " ",
							"app_data": ""
						},
						"2": {
							"type": "html",
							"value": "Armaments of Leviathan",
							"color": "9da1a9",
							"app_data": {
								"def_index": "20267",
								"is_itemset_name": "1"
							}
						},
						"3": {
							"type": "html",
							"value": "Admiral's Foraged Cap",
							"color": "6c7075",
							"app_data": {
								"def_index": "5463"
							}
						},
						"4": {
							"type": "html",
							"value": "Admiral's Stash",
							"color": "6c7075",
							"app_data": {
								"def_index": "5464"
							}
						},
						"5": {
							"type": "html",
							"value": "Claddish Gauntlets",
							"color": "6c7075",
							"app_data": {
								"def_index": "5465"
							}
						},
						"6": {
							"type": "html",
							"value": "Claddish Guard",
							"color": "6c7075",
							"app_data": {
								"def_index": "5466"
							}
						},
						"7": {
							"type": "html",
							"value": "Claddish Hightops",
							"color": "6c7075",
							"app_data": {
								"def_index": "5467"
							}
						},
						"8": {
							"type": "html",
							"value": "Neptunian Sabre",
							"color": "6c7075",
							"app_data": {
								"def_index": "5468"
							}
						},
						"9": {
							"type": "html",
							"value": "Admiral's Salty Shawl",
							"color": "6c7075",
							"app_data": {
								"def_index": "5469"
							}
						},
						"10": {
							"type": "html",
							"value": "Ye Olde Pipe",
							"color": "6c7075",
							"app_data": {
								"def_index": "5470"
							}
						},
						"11": {
							"type": "html",
							"value": "An old pipe Kunkka plucked from a tidepool after being shipwrecked, the old seadog claims it brings him fortune and good luck, and thus never ever takes it from his scurvy-ridden mouth. \r\n\t",
							"app_data": ""
						}
					},
					"tags": {
						"0": {
							"internal_name": "unique",
							"name": "Standard",
							"category": "Quality",
							"color": "D2D2D2",
							"category_name": "Quality"
						},
						"1": {
							"internal_name": "Rarity_Common",
							"name": "Common",
							"category": "Rarity",
							"color": "b0c3d9",
							"category_name": "Rarity"
						},
						"2": {
							"internal_name": "wearable",
							"name": "Wearable",
							"category": "Type",
							"category_name": "Type"
						},
						"3": {
							"internal_name": "neck",
							"name": "Neck",
							"category": "Slot",
							"category_name": "Slot"
						},
						"4": {
							"internal_name": "npc_dota_hero_kunkka",
							"name": "Kunkka",
							"category": "Hero",
							"category_name": "Hero"
						}
					},
					"classid": "123456789"
				},
				"success": true
			}
		}
	`
}
