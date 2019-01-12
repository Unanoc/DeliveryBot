// package main

// import (
// 	"encoding/json"
// 	"flag"
// 	"io/ioutil"
// 	"log"
// 	"vkbot/bot"
// 	"vkbot/config"
// 	"vkbot/database"
// )

// var pathToConfig = flag.String("config", "", "Path to configuration JSON file")

// func main() {
// 	flag.Parse()

// 	// Setting of configuration
// 	configJSON, err := ioutil.ReadFile(*pathToConfig)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	config := config.Config{}
// 	if err := json.Unmarshal(configJSON, &config); err != nil {
// 		log.Panic(err)
// 	}

// 	// Connecting to database
// 	db := database.DB{}
// 	if err := db.Connect(config.Connection); err != nil {
// 		log.Panic(err)
// 	}
// 	defer db.Disconnect()

// 	// Launching bot
// 	bot.Run(&db, config.AccessToken)
// }

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"vkbot/api/bot"
	"vkbot/api/vk"
)

var token = "7e420872611f4251cb09eb44333647906d9d125d58be469bafc26d57e79369d69350f17b3b2a39c9a068a"
var groupID = 176353269

const (
	apiBaseURL = "https://api.vk.com/method/"
)

func main() {
	if token == "" {
		log.Fatal("token is required")
	}

	baseAPI, err := vk.NewBaseAPI(vk.BaseAPIConfig{
		AccessToken: token,
	})
	if err != nil {
		log.Fatal("Cant create baseAPI:", err)
	}

	bot, err := bot.NewBot(baseAPI, bot.BotConfig{
		GroupID: groupID,
		Poller: &bot.LongPoller{
			Wait: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatal("Cant create bot:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(cancel context.CancelFunc) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		cancel()
	}(cancel)

	events, err := bot.StartPolling(ctx, 0)
	if err != nil {
		log.Fatal("Cat start polling:", err)
	}

	// докопипастить остальную часть с events
	for e := range events {
		// log.Printf("Got event: %+v", e)

		switch ev := e.Event.(type) {
		case vk.MessageNew:
			from := ev.PeerID
			text := ev.Text
			msgID := ev.ID

			log.Printf("New message(%v) from %v: `%v`", msgID, from, text)

			if text != "" {
				resp, err := vk.Messages{bot}.Send(vk.MessagesSendParams{
					PeerID:  from,
					Message: text,
				})

				if err != nil {
					log.Printf("Cant send reply to (%v): %v", msgID, err)
				} else {
					log.Printf("Sent reply to (%v): reply id %v", msgID, resp)
				}
			}
		}
	}

	log.Printf("Bye!")
}
