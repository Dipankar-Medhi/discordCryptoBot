package main

import (
	"fmt"

	"github.com/Dipankar-Medhi/goDiscordBot/bot"
	"github.com/Dipankar-Medhi/goDiscordBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan chan struct{})
	return
}
