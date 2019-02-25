package snitch

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the MessageBroker now ")

	return client
}

func Init(brokerURL, serviceName, topic string) error {
	uri, err := url.Parse(brokerURL)
	if err != nil {
		log.Fatal(err)
	}
	ServiceName = serviceName
	Topic = topic
	snitchClient = connect("pub", uri)
	return err
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ws://%s", uri.Host))
	// opts.SetUsername(uri.User.Username())
	// password, _ := uri.User.Password()
	// opts.SetPassword(password)
	// opts.SetClientID(clientId)
	return opts
}

func SendLog(payload interface{}) {
	log.Println("Now Publishing to MQTT ")
	snitchClient.Publish(Topic, 2, false, payload)
}
