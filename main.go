package main

import (
	"context"
	"log"
	"time"

	"github.com/mixinmessenger/bot-api-go-client"

	"github.com/ibigbug/caoliu/config"
	"github.com/ibigbug/caoliu/mixin"
)

func main() {
	var handler = mixin.Messenger{}

	for {
		if err := bot.Loop(context.Background(), handler, config.ClientID, config.SessionID, config.PrivateKey); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
