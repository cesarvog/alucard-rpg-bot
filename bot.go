package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var dice *Dice
//var goBot *discordgo.Session

const (
	help = "Use me like this\nTo do a test/contest use '!d n h' where 'n' is normal dice pool and 'h' hunger dice pool\ne.g: !d 4 2\nTo do a rouse check use '!hunger'"
)

func Start() {
	//var goBot *discordgo.Session
	goBot, err := discordgo.New("Bot " + token)
	dice = &Dice{}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Content
	if ! isTalkingToMe(m) { 
		return
	}

	params := strings.Split(msg, " ")

	if params[0] == "!d" { //normal roll
		responseDice(params, s, m)
	} else if params[0] == "!hunger" {
		s.ChannelMessageSend(m.ChannelID, dice.Hunger())
		return 
	}
}


func responseDice(params []string, s *discordgo.Session, m *discordgo.MessageCreate) { 
	if len(params) != 3 {
		s.ChannelMessageSend(m.ChannelID, help)
		return
	}

	n, err := strconv.Atoi(params[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, help)
		return
	}

	h, err := strconv.Atoi(params[2])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, help)
		return
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Master <@%s>\n", m.Author.ID)
	fmt.Fprintf(&b, "```css\n%s\n```", dice.Roll(n, h))

	fmt.Println(b.String())
	s.ChannelMessageSend(m.ChannelID, b.String())
}

func isTalkingToMe(m *discordgo.MessageCreate) bool {
	if m.Author.ID == BotId {
		return false
	}

	msg := m.Content
	if len(msg) <= 0 || string(msg[0]) != "!" {
		return false
	}

	return true
}

