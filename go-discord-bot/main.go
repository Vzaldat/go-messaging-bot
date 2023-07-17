package main

import (
	"fmt"

	"github.com/Vzaldat/discord-bot/bot"
	"github.com/Vzaldat/discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
