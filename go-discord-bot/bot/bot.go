package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vzaldat/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
	}
	u, err1 := goBot.User("@me")

	if err1 != nil {
		fmt.Println(err1.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err2 := goBot.Open()

	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println("Bot is running!!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "puka")
	}
	x := len(m.Content)
	p := "This message has " + strconv.Itoa(x-22) + " Characters"
	if strings.Contains(m.Content, fmt.Sprintf("@%s", s.State.User.ID)) {
		s.ChannelMessageSend(m.ChannelID, p)
		fmt.Println(x, p)
	}
}
