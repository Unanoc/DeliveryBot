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
		database.CreateOrUpdateUserState(db, userID, StateNull)
		return
	}

	userState := database.GetUserStateByID(db, userID)
	stateHandler(bot, db, userID, text, userState)
}

func stateHandler(bot *vkapi.Client, db *database.DB, userID int64, text string, state int) {
	switch state {
	case StateNull:
		if strings.ToLower(text) == "заказ" {
			sendMessage(bot, userID, askingFirstNameMsg)
			database.CreateOrUpdateUserState(db, userID, StateFirstName)
		} else {
			sendMessage(bot, userID, errorMsg)
		}
	case StateFirstName:
		sendMessage(bot, userID, askingLastNameMsg)
		database.CreateOrUpdateUserState(db, userID, StateLastName)
	case StateLastName:
		sendMessage(bot, userID, askingPhone)
		database.CreateOrUpdateUserState(db, userID, StatePhone)
	case StatePhone:
		sendMessage(bot, userID, askingCompanyMsg)
		database.CreateOrUpdateUserState(db, userID, StateCompanyName)
	case StateCompanyName:
		sendMessage(bot, userID, askingAddressMsg)
		database.CreateOrUpdateUserState(db, userID, StateAddress)
	case StateAddress:
		sendMessage(bot, userID, askingDateMsg)
		database.CreateOrUpdateUserState(db, userID, StateDate)
	case StateDate:
		sendMessage(bot, userID, confirmationMsg)
		order, err := database.SelectFinishedOrderByID(db, userID)
		if err != nil {
			sendMessage(bot, userID, confirmationErrorMsg)
			database.CreateOrUpdateUserState(db, userID, StateNull)
			return
		}
		sendMessage(bot, userID, order.String())
		database.CreateOrUpdateUserState(db, userID, StateConfirmation)
	case StateConfirmation:
		if strings.ToLower(text) == "да" {
			sendMessage(bot, userID, successMsg)
			database.SetFinishFlagOrder(db, userID)
			database.CreateOrUpdateUserState(db, userID, StateNull)
		} else {
			if strings.ToLower(text) == "нет" {
				sendMessage(bot, userID, cancelOrderMsg)
				database.CreateOrUpdateUserState(db, userID, StateNull)
			} else {
				sendMessage(bot, userID, confirmationMsg)
			}
		}
	default:
		sendMessage(bot, userID, errorMsg)
	}
}

func sendMessage(bot *vkapi.Client, userID int64, text string) {
	bot.SendMessage(
		vkapi.NewMessage(
			vkapi.NewDstFromUserID(userID),
			text,
		),
	)
}

// Добавить валидацию на дату и возможно не только
// заполнение базы после каждого поля orders
// заполнения таблицы users
