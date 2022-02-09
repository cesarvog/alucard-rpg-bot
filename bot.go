package bot

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

var help = "Use me like this\nTo do a test/contest use '!d n h' where 'n' is normal dice pool and 'h' hunger dice pool\ne.g: !d 4 2\nTo do a rouse check use '!hunger'"
func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

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
	if m.Author.ID == BotId {
		return
	}

	//!d 10 2     -> roll 10 normal dies and two of hunger
	msg := m.Content
	if len(msg) <= 0 || msg[0:1] != "!" {
		return
	}

	params := strings.Split(msg, " ")

	if params[0] == "!d" { //normal roll
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
		fmt.Fprintf(&b, "\n```css\n%s\n```", dies.Roll(n, h))

		fmt.Println(b.String())
		s.ChannelMessageSend(m.ChannelID, b.String())
		return
	} else if params[0] == "!hunger" && msg[1:5] == "roll" {
		s.ChannelMessageSend(m.ChannelID, dies.Hunger())
		return 
	} else {
		return
	}
}
