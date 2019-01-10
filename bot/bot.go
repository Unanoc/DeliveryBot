package bot

import (
	"log"
	"vkbot/database"

	vkapi "github.com/dimonchik0036/vk-api"
)

// Run runs the bot.
func Run(db *database.DB, accessToken string) {
	bot, err := vkapi.NewClientFromToken(accessToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Log(true)

	if err := bot.InitLongPoll(0, 2); err != nil {
		log.Panic(err)
	}

	LPCfg := vkapi.LPConfig{
		Wait: 25,
		Mode: vkapi.LPModeAttachments,
	}
	updates, _, err := bot.GetLPUpdatesChan(100, LPCfg)
	if err != nil {
		log.Panic(err)
	}

	var msg *vkapi.LPMessage

	for update := range updates {
		msg = update.Message

		if msg == nil || !update.IsNewMessage() || msg.Outbox() {
			continue
		}

		go msgHandler(bot, db, msg.FromID, msg.Text)
	}
}
