package steamapi

// GetMockOKGetPlayerItems steam API mock response gives you item_id, def_index...
// https://api.steampowered.com/IEconItems_XXXX/GetPlayerItems/v1/?key=XXXXXX&format=json&steamid=XXXXX
func GetMockOKGetPlayerItems() string {
	return `
		{
			"result": {
				"status": 1,
				"num_backpack_slots": 840,
				"items": [
					{
						"id": 1234567894,
						"original_id": 1234567894,
						"defindex": 5470,
						"level": 1,
						"quality": 4,
						"inventory": 19,
						"quantity": 1,
						"equipped": [
							{
								"class": 23,
								"slot": 6
							}
						]
						
					},
					{
						"id": 1234567895,
						"original_id": 1234567895,
						"defindex": 15001,
						"level": 1,
						"quality": 4,
						"inventory": 30,
						"quantity": 1,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1065353216,
								"float_value": 1
							},
							{
								"defindex": 16,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							}
						]
						
					},
					{
						"id": 1234567896,
						"original_id": 1234567896,
						"defindex": 10068,
						"level": 1,
						"quality": 4,
						"inventory": 5,
						"quantity": 1,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1065353216,
								"float_value": 1
							}
						]
						
					},
					{
						"id": 1234567897,
						"original_id": 1234567897,
						"defindex": 5508,
						"level": 1,
						"quality": 4,
						"inventory": 25,
						"quantity": 1,
						"equipped": [
							{
								"class": 4,
								"slot": 4
							}
						]
						
					},
					{
						"id": 1234567898,
						"original_id": 1234567898,
						"defindex": 7480,
						"level": 1,
						"quality": 4,
						"inventory": 142,
						"quantity": 1,
						"equipped": [
							{
								"class": 15,
								"slot": 0
							}
						]
						,
						"flag_cannot_trade": true,
						"attributes": [
							{
								"defindex": 153,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							},
							{
								"defindex": 213,
								"value": 1,
								"float_value": 1.4012984643248171e-045
							}
						]
						
					}
				]
				
			}
		}
	`
}
