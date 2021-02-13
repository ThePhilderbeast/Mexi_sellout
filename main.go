package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

type Config struct {
	Username string
	Oauth    string
	LukeBans int
}

var bansEnabled = false
var config = Config{}
var client *twitch.Client

func main() {

	dat, err := ioutil.ReadFile("./config.yml")
	check(err)
	fmt.Print(string(dat))

	client = twitch.NewClient(config.Username, "oauth:"+config.Oauth)

	client.OnPrivateMessage(commandsHandler)
	client.OnNoticeMessage(func(message twitch.NoticeMessage) {
		fmt.Println(message.Message)
	})

	client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		fmt.Println(message.Message)
	})
	client.Join("mexi")

	err = client.Connect()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func commandsHandler(message twitch.PrivateMessage) {
	if message.User.DisplayName == "LukeAdrian29" && bansEnabled {
		if rand.Float32() <= 0.75 {
			client.Say(message.Channel, "/timeout LukeAdrian29 1")
			config.LukeBans++
		} else {
			fmt.Println("he lives")
		}
	}

	if strings.HasPrefix(message.Message, "!nolulu") && bansEnabled {
		client.Say(message.Channel, "No Lulu!")
		client.Say(message.Channel, "/timeout LukeAdrian29 1")
		config.LukeBans++
	}

	if strings.HasPrefix(message.Message, "!lulubans") {
		msg := fmt.Sprintf("lulu has been banned %d times", config.LukeBans)
		fmt.Println(message.Channel + ": " + msg)
		client.Say(message.Channel, msg)
	}
}
