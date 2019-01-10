package bot

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"vkbot/database"

	vkapi "github.com/dimonchik0036/vk-api"
)

func msgHandler(bot *vkapi.Client, db *database.DB, userID int64, text string) {
	userState := database.GetUserStateByID(db, userID)

	if text == "/start" {
		sendMessage(bot, userID, startMsg)
		saveUserInfo(bot, db, userID)
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

func getAgeByBirth(bday string) int {
	dayMonthYear := strings.Split(bday, ".")
	if len(dayMonthYear) < 3 {
		return -1
	}
	userYear, _ := strconv.Atoi(dayMonthYear[2])
	userMonth, _ := strconv.Atoi(dayMonthYear[1])
	userDay, _ := strconv.Atoi(dayMonthYear[0])

	yearNow, monthNow, dayNow := time.Now().Date()
	age := yearNow - userYear

	if int(monthNow) < userMonth {
		age--
	}

	if int(monthNow) == userMonth {
		if dayNow < userDay {
			age--
		}
	}
	return age
}

func saveUserInfo(bot *vkapi.Client, db *database.DB, userID int64) {
	info, err := bot.UsersInfo(vkapi.NewDstFromUserID(userID), "bdate", "first_name", "last_name", "sex")
	if err != nil {
		fmt.Println(err)
		return
	}

	database.CreateOrUpdateUser(
		db,
		userID,
		info[0].FirstName,
		info[0].LastName,
		getAgeByBirth(info[0].Bdate),
		info[0].Sex,
	)
}
