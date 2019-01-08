package bot

// user's states
const (
	StateNull = iota
	StateFirstName
	StateLastName
	StatePhone
	StateCompanyName
	StateAddress
	StateDate
	StateConfirmation
)

// bot's replies
const (
	// greetings
	startMsg = `
	Привет! Я умею легко организовывать процесс доставки.
	Для этого необходимо заполнить анкету! 
	Отправь мне сообщение "Заказ" и начинай заполнять.
	`

	// error
	errorMsg = `
	Что-то я ничего не понял... отправь сообщение "/start", чтобы понять
	какие сообщения я понимаю.
	`

	// order
	askingFirstNameMsg = `
	Введи имя
	(для отмены процесса, напиши "/start")
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
	Правильно ли составлен заказ? 
	(да/нет либо для отмены напишите "/start")
	`
	confirmationErrorMsg = `
	Не удалось составить заказ, попробуйте снова...
	`
	successMsg = `
	Заказ успешно оформлен!
	`
	cancelOrderMsg = `
	Заказ отменен
	`
)
