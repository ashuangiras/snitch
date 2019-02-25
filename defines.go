package snitch

import mqtt "github.com/eclipse/paho.mqtt.golang"

// // const (
// // 	TypeNetworkLog         = 1
// // 	TypeApplicationLog     = 2
// // 	TypeInternalHandlerLog = 3
// // 	TypeInternalMethodLog  = 4
// // )

// var TypeText = map[string]string{
// 	"TypeNetworkLog":         "Network traffic logs",
// 	"TypeApplicationLog":     "Overall Application Specific logs",
// 	"TypeInternalMethodLog":  "Application Internal methods logs",
// 	"TypeInternalHandlerLog": "Internal handlers logs",
// }

// func SnitchTypeText(code string) string {
// 	return TypeText[code]
// }
var snitchClient mqtt.Client
var ServiceName string
var Topic string
