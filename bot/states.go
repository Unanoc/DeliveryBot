package bot

import (
	"time"
)

// Order is a struct for order confirmation.
type Order struct {
	OrderID    int
	UserID     int64
	FirstName  string
	LastName   string
	Phone      string
	Company    string
	Address    string
	Date       time.Time
	IsFinished bool
}

// user's states
const (
	stateNull = iota
	stateStartingOrder
	stateFirstName
	stateLastName
	statePhone
	stateCompanyName
	stateAddress
	stateDate
	stateConfirmation
)

// bot's replies
const (
	// greetings
	startMsg = `
	Привет! Я умею легко организовывать процесс доставки.
	Для этого необходимо заполнить анкету! 
	Отправь мне сообщение "Анкета" и начинай заполнять.
	`

	// error
	errorMsg = `
	Что-то я ничего не понял... отправь сообщение "/start", чтобы понять
	какие сообщения я понимаю.
	`

	// profile
	askingFirstNameMsg = `
	Введи имя (для отмены процесса, напиши "/start")
	`
	askingLastNameMsg = `
	Введи фамилию
	`
	askingPhone = `
	Введи номер телефона 
	`
	askingCompanyMsg = `
	Введи название компании, от которой будет осуществляться доставка
	`
	askingAddressMsg = `
	Введи адрес, куда нужно доставить
	`
	askingDateMsg = `
	Введи дату доставки
	`
	confirmationMsg = `
	Правильно ли составлен заказ? (да/нет либо для отмены напишите "/start")
	`
	confirmationErrorMsg = `
	Не удалось составить анкету, попробуйте снова
	`
	successMsg = `
	Заказ успешно оформлен!
	`
	cancelOrderMsg = `
	Заказ отменен
	`
)
