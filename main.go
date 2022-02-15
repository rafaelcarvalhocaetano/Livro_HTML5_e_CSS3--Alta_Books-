package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	kfk "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/rafaelcarvalhocaetano/msgo/email"
	"github.com/rafaelcarvalhocaetano/msgo/kafka"

	gomail "gopkg.in/mail.v2"
)

func main() {
	var emailChannel = make(chan email.Email)
	var kfkMessage = make(chan *kfk.Message)

	d := gomail.NewDialer(
		"smtp.live.com",
		587,
		"rapha.pse@outlook.com",
		"10Sistem@s",
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	es := email.NewMailSender()
	es.From = "rapha.pse@gmail.com"
	es.Dailer = d
	// thread
	go es.Send(emailChannel)

	configMap := &kfk.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"client.id":         "emailapp",
		"group.id":          "emailapp",
	}

	topics := []string{"emails"}
	consumer := kafka.NewConsumer(configMap, topics)

	// thread
	go consumer.Consume(kfkMessage)

	for msg := range kfkMessage {
		var input email.Email
		json.Unmarshal(msg.Value, &input)
		fmt.Println("recebendo mensagem")
		emailChannel <- input
	}

}
