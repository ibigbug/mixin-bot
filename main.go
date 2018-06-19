package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/mixinmessenger/bot-api-go-client"

	"github.com/ibigbug/caoliu/config"
	"github.com/ibigbug/caoliu/mixin"
)

func main() {
	godotenv.Load("./scripts/.env")
	var handler = mixin.Messenger{}

	for {
		if err := bot.Loop(context.Background(), handler, config.ClientID, config.SessionID, config.PrivateKey); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
