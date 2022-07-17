package bot

import (
	"BabyGoBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

var Id string

func Start() {

	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	Id = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Checking to make sure that the incoming message is not from the bot
	if m.Author.ID == Id {
		return
	}

	//Checking incoming messages for "I'm __" and most variations for that
	dadJoke, _ := regexp.Compile("[iI]['\"’]?[mM] ([-a-zA-Z0-9’' ]+)")

	var test = dadJoke.FindStringSubmatch(m.Content)

	//Sending a message as response in the form of a dad joke
	if test != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi, "+test[1]+", I'm Baby Goose Bot!")
	}

	//Checking incoming messages for the phrase "ping" with any capitalization and responding with "pong"
	var content = strings.ToLower(m.Content)

	if content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
