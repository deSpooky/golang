package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	db, err := InitDB("dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cache := NewCache()
	cards, _ := LoadCards(db)
	cache.LoadFromDB(cards)

	nc, _ := nats.Connect(nats.DefaultURL)
	SubscribeNATS(nc, cache, db)

	http.HandleFunc("/card", GetCardHandler(cache))

	go func() {
		fmt.Println("http server running on :8080")
		http.ListenAndServe(":8080", nil)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}