package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"vkbot/api/bot"
	"vkbot/api/vk"
	"vkbot/database"
)

// Run runs the bot.
func Run(db *database.DB, accessToken string, groupID int) error {
	if accessToken == "" {
		return fmt.Errorf("token is required")
	}

	baseAPI, err := vk.NewBaseAPI(vk.BaseAPIConfig{
		AccessToken: accessToken,
	})
	if err != nil {
		return err
	}

	bot, err := bot.NewBot(baseAPI, bot.BotConfig{
		GroupID: groupID,
		Poller: &bot.LongPoller{
			Wait: 10 * time.Second,
		},
	})
	if err != nil {
		return err
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
		return err
	}

	for e := range events {
		switch ev := e.Event.(type) {
		case vk.MessageNew:
			from := ev.PeerID
			text := ev.Text

			log.Printf("New message from %v: `%v`", from, text)

			if text != "" {
				go msgHandler(bot, db, from, text)
			}
		}
	}

	return err
}
