package bot

import (
	"strings"
	"vkbot/database"

	vkapi "github.com/dimonchik0036/vk-api"
)

func msgHandler(bot *vkapi.Client, db *database.DB, userID int64, text string) {
	// if message if "/start", it always resets the state of user and sets stateNull
	if text == "/start" {
		sendMessage(bot, userID, startMsg)
		database.SetUserState(db, userID, stateNull)
		return
	}

	userState := database.GetUserStateByID(db, userID)

	// continue interaction with user if state > stateNull
	if userState > stateNull {
		stateHandler(bot, db, userID, text, userState)
		return
	}

	// the first state change after stateNull
	if strings.ToLower(text) == "анкета" {
		database.SetUserState(db, userID, stateStartingOrder)
		stateHandler(bot, db, userID, text, stateStartingOrder)
		return
	}

	// when stateNull and message != "start"
	sendMessage(bot, userID, errorMsg)
}

func sendMessage(bot *vkapi.Client, userID int64, text string) {
	bot.SendMessage(
		vkapi.NewMessage(
			vkapi.NewDstFromUserID(userID),
			text,
		),
	)
}

func stateHandler(bot *vkapi.Client, db *database.DB, userID int64, text string, state int) {
	switch state {
	case stateStartingOrder:
		sendMessage(bot, userID, askingFirstNameMsg)
		database.SetUserState(db, userID, stateFirstName)
	case stateFirstName:
		sendMessage(bot, userID, askingLastNameMsg)
		database.SetUserState(db, userID, stateLastName)
	case stateLastName:
		sendMessage(bot, userID, askingPhone)
		database.SetUserState(db, userID, statePhone)
	case statePhone:
		sendMessage(bot, userID, askingCompanyMsg)
		database.SetUserState(db, userID, stateCompanyName)
	case stateCompanyName:
		sendMessage(bot, userID, askingAddressMsg)
		database.SetUserState(db, userID, stateAddress)
	case stateAddress:
		sendMessage(bot, userID, askingDateMsg)
		database.SetUserState(db, userID, stateDate)
	case stateDate:
		sendMessage(bot, userID, confirmationMsg)
		profile, err := database.selectProfileByID(db, userID)
		if err != nil {
			sendMessage(bot, userID, confirmationErrorMsg)
			database.SetUserState(db, userID, stateNull)
			return
		}
		sendMessage(bot, userID, profile.String())
		database.SetUserState(db, userID, stateConfirmation)
	case stateConfirmation:
		if strings.ToLower(text) == "да" {
			sendMessage(bot, userID, successMsg)
			database.SetUserState(db, userID, stateNull)
		} else {
			if strings.ToLower(text) == "нет" {
				sendMessage(bot, userID, cancelOrderMsg)
				database.SetUserState(db, userID, stateNull)
			} else {
				sendMessage(bot, userID, confirmationMsg)
			}
		}
	}
}
