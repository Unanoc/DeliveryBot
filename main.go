package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"vkbot/bot"
	"vkbot/config"
	"vkbot/database"
)

var pathToConfig = flag.String("config", "", "Path to configuration JSON file")

func main() {
	flag.Parse()

	// Setting of configuration
	configJSON, err := ioutil.ReadFile(*pathToConfig)
	if err != nil {
		log.Panic(err)
	}

	config := config.Config{}
	if err := json.Unmarshal(configJSON, &config); err != nil {
		log.Panic(err)
	}

	// Connecting to database
	db := database.DB{}
	if err := db.Connect(config.Connection); err != nil {
		log.Panic(err)
	}
	defer db.Disconnect()

	// Launching bot
	if err := bot.Run(&db, config.AccessToken, config.GroupID); err != nil {
		log.Panic(err)
	}
}
