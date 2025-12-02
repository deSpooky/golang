package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func SubscribeNATS(nc *nats.Conn, cache *Cache, dbConn *sql.DB) error {
	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{Name: "foo"})

	cfg := jetstream.ConsumerConfig{Name: "cons"}
	cons, err := js.CreateOrUpdateConsumer(ctx, "foo", cfg)
	if err != nil {
		return err
	}

	cc, _ := cons.Consume(func(msg jetstream.Msg) {
		defer msg.Ack()
		card := PostCard{}
		if err := json.Unmarshal(msg.Data(), &card); err != nil {
			fmt.Println("invalid data from nats:", err)
			return
		}
		cache.Set(card)
		if err := InsertCard(dbConn, card); err != nil {
			fmt.Println("db insert error:", err)
		}
	})
	defer cc.Stop()
	return nil
}