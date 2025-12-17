package infra

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client

func InitMQTT(broker string) mqtt.Client {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID("greenhouse-backend")

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		return nil
	}
	return Client
}

func Publish(topic string, payload string) error {
	token := Client.Publish(topic, 1, false, payload)
	token.Wait()
	return token.Error()
}

func IsConnected() bool {
	return Client != nil && Client.IsConnected()
}
