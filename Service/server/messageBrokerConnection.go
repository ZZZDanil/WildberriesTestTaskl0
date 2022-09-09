package service

import (
	"log"

	"github.com/nats-io/stan.go"
)

type MessageBroker struct {
	nats stan.Conn
}

func (mb *MessageBroker) ConnectToMessageBroker(mbo *MessageBrokerOptions) {
	var err error
	mb.nats, err = stan.Connect(mbo.ClusterID, mbo.ClientID)
	if err != nil {
		log.Print(err)
	}
}

func (mb *MessageBroker) SubscribeMessageBroker(cache *DataBaseCache, db *DataBase) {
	if mb.nats != nil {
		_, err := mb.nats.Subscribe("foo", func(m *stan.Msg) {
			s := string(m.Data[:])
			trueData, err := cache.AddToLocalCache(s)
			if err == nil {
				_, err := db.connection.Exec("insert into order_table_cache (data) values ($1)", trueData)
				if err != nil {
					log.Print(err)
				}
			} else {
				log.Print(err)
			}
		}, stan.DurableName("my-durable"))
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Print("MessageBroker is nil")
	}
}

func (mb *MessageBroker) PublishMessageBroker(message []byte) {
	if mb.nats != nil {
		err := mb.nats.Publish("foo", message)
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Print("MessageBroker is nil")
	}
}
