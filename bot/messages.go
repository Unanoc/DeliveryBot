package bot

import (
	"strings"
	"vkbot/database"

	vkapi "github.com/dimonchik0036/vk-api"
)

func msgHandler(bot *vkapi.Client, db *database.DB, userID int64, text string) {
	userState := database.GetUserStateByID(db, userID)

	if text == "/start" {
		sendMessage(bot, userID, startMsg)
		//заполнение таблицы Users по входящему сообщению

		// removing order if it is not completed
		if userState != 0 {
			database.DeleteOrder(db, userID)
		}
		database.CreateOrUpdateUserState(db, userID, StateNull)
		return
	}

	stateHandler(bot, db, userID, text, userState)
}

func stateHandler(bot *vkapi.Client, db *database.DB, userID int64, text string, state int) {
	switch state {
	case StateNull:
		if strings.ToLower(text) == "заказ" {
			sendMessage(bot, userID, askingFirstNameMsg)
			database.CreateOrder(db, userID)
			database.CreateOrUpdateUserState(db, userID, StateFirstName)
		} else {
			sendMessage(bot, userID, errorMsg)
		}
	case StateFirstName:
		sendMessage(bot, userID, askingLastNameMsg)
		database.UpdateOrder(db, userID, "firstname", text)
		database.CreateOrUpdateUserState(db, userID, StateLastName)
	case StateLastName:
		sendMessage(bot, userID, askingPhone)
		database.UpdateOrder(db, userID, "lastname", text)
		database.CreateOrUpdateUserState(db, userID, StatePhone)
	case StatePhone:
		sendMessage(bot, userID, askingCompanyMsg)
		database.UpdateOrder(db, userID, "phone", text)
		database.CreateOrUpdateUserState(db, userID, StateCompanyName)
	case StateCompanyName:
		sendMessage(bot, userID, askingAddressMsg)
		database.UpdateOrder(db, userID, "company", text)
		database.CreateOrUpdateUserState(db, userID, StateAddress)
	case StateAddress:
		sendMessage(bot, userID, askingDateMsg)
		database.UpdateOrder(db, userID, "address", text)
		database.CreateOrUpdateUserState(db, userID, StateDate)
	case StateDate:
		sendMessage(bot, userID, confirmationMsg)
		database.UpdateOrder(db, userID, "delivery_date", text)
		order, err := database.SelectOrderByID(db, userID)
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
			database.CreateOrUpdateUserState(db, userID, StateNull)
		} else {
			if strings.ToLower(text) == "нет" {
				sendMessage(bot, userID, cancelOrderMsg)
				database.CreateOrUpdateUserState(db, userID, StateNull)
				database.DeleteOrder(db, userID)
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

// заполнения таблицы users
