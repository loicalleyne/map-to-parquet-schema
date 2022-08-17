package main

var exampleMap = map[string]interface{}{
	"request": map[string]interface{}{
		"datetime":    "2021-07-27 02:59:59",
		"ip":          "172.222.233.111",
		"host":        "www.domain.com",
		"uri":         "/api/v1/",
		"request_uri": "/api/v1/",
		"referer":     "",
		"useragent":   "",
	},
	"metadata": map[string]interface{}{
		"compression_ratio": 1,
		"coroutine_uuid":    "46c45675-5c10-4094-a124-f8615f2b10db",
		"hostname":          "server.domain.com",
	},
	"incoming_request": map[string]interface{}{
		"id":                "1d25bb6a-3e09-4a8e-8b0f-93a0e9383904",
		"external_id":       "afwef",
		"direct_connection": true,
		"remote_ip":         "67.49.160.53",
	},
	"entity": map[string]interface{}{
		"id":                      39,
		"api_key":                 "1d6492bb-490e-4cb9-1d24-2a24db8d07e7",
		"tracking_protocol":       "emulator",
		"event_proxy":             nil,
		"feature_support_enabled": true,
		"primary_currency":        "EUR",
		"rate_percent":            10,
	},
	"screen": map[string]interface{}{
		"id": 86233,
		"ids": []interface{}{
			132, 453535, 13412341,
		},
		"external_id": "com.domain:215426709",
		"width":       1080,
		"height":      1920,
	},
	"geo": map[string]interface{}{
		"ip":        "0.0.0.0",
		"country":   "NL",
		"region":    "06",
		"city":      "Noordhoek",
		"dma":       "1234",
		"zip":       "1345-A",
		"latitude":  51.558441162109375,
		"longitude": 5.078000068664551,
	},
	"schedule": map[string]interface{}{
		"id": 18604307,
	},
	"audience": map[string]interface{}{
		"origin":    "file",
		"SOMEFLOAT": 5.255000114440918,
	},
	"provider": map[string]interface{}{
		"id":               42,
		"protocol":         "openrtb",
		"feature1_enabled": false,
		"feature2_enabled": true,
		"primary_currency": "USD",
	},
	"outgoing_request": map[string]interface{}{
		"string_array": []interface{}{
			"P0OW42XDHA",
			"KLTFJYG9FX",
		},
		"id": "9abbb401-d29a-4025-bd45-416b2ebf13e3",
	},
}
