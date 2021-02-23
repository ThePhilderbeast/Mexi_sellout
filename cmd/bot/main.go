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
}

type victim struct {
	Username     string
	EnabledUntil time.Time
	EnabledCount int
	TimeoutCount int
}

var config = Config{}
var client *twitch.Client
var victims []victim

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

	for _, v := range victims {
		if v.Username == strings.ToLower(message.User.DisplayName) {
			if time.Now().Before(v.EnabledUntil) {
				if rand.Float32() <= 0.50 {
					client.Say(message.Channel, "/timeout message.User.DisplayName 1")
					v.TimeoutCount++
				}
			}
		}
	}

	if strings.HasPrefix(message.Message, "!nomore") {
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
	victimName := "lukeadrian29"
	if len(arg) >= 2 {
		if len(arg) == 3 {
			victimName = strings.ToLower(arg[2])
		}

		vic := victim{
			Username:     strings.ToLower(victimName),
			EnabledUntil: time.Now(),
			EnabledCount: 0,
			TimeoutCount: 0,
		}

		for _, v := range victims {
			if v.Username == victimName {
				vic = v
				break
			}
		}

		minutes, err := strconv.Atoi(arg[1])
		check(err)
		vic.EnabledUntil = time.Now()
		vic.EnabledUntil = vic.EnabledUntil.Add(time.Minute * time.Duration(minutes))
		vic.EnabledCount++
	}
}
