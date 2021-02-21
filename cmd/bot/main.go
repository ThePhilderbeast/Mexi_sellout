package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"gopkg.in/yaml.v2"
)

// Config the bot config file
type Config struct {
	Username string
	Oauth    string
	LukeBans int
}

var config = Config{}
var client *twitch.Client
var enabledUntil = time.Now()

func main() {

	dat, err := ioutil.ReadFile("./configs/config.yml")
	check(err)
	yaml.Unmarshal(dat, &config)
	fmt.Println(config)

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

	if message.User.DisplayName == "LukeAdrian29" {
		if time.Now().Before(enabledUntil) {
			if rand.Float32() <= 0.50 {
				client.Say(message.Channel, "/timeout LukeAdrian29 1")
				config.LukeBans++

		}
	}

	if strings.HasPrefix(message.Message, "!lulubans") {
		msg := fmt.Sprintf("lulu has been banned %d times", config.LukeBans)
		fmt.Println(message.Channel + ": " + msg)
		client.Say(message.Channel, msg)
	}

	if strings.HasPrefix(message.Message, "!nolulu") {
		if message.User.DisplayName == "Philderbeast" {
			fmt.Println("Enabling bot msg from " + message.User.DisplayName)
			enableBan(message.Message)
			client.Say(message.Channel, "Bot enabled")
		}

		if _, ok := message.User.Badges["moderator"]; ok {
			fmt.Println("Enabling bot msg from " + message.User.DisplayName)
			enableBan(message.Message)
			client.Say(message.Channel, "Bot enabled")
		}

		if _, ok := message.User.Badges["broadcaster"]; ok {
			fmt.Println("Enabling bot msg from " + message.User.DisplayName)
			enableBan(message.Message)
			client.Say(message.Channel, "Bot enabled")
		}
	}
}

func enableBan(message string) {

	arg := strings.Split(message, " ")
	if len(arg) >= 2 {
		minutes, err := strconv.Atoi(arg[1])
		check(err)
		enabledUntil = time.Now()
		enabledUntil = enabledUntil.Add(time.Minute * time.Duration(minutes))
	}
}
