package main

import (
	"fmt"
	"alucard-rpg-bot/bot"
	"alucard-rpg-bot/config"
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


