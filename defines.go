package snitch

import mqtt "github.com/eclipse/paho.mqtt.golang"

var snitchClient mqtt.Client
var Topic, ServiceName string
