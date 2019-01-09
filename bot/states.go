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
	Приветствую! Я умею легко организовывать процесс доставки.
	Для этого необходимо оформить заказ! 
	Отправьте мне сообщение "Заказ" и начинай заполнять.
	`

	// error
	errorMsg = `
	Что-то я ничего не понял... отправьте сообщение "/start", чтобы понять
	какие сообщения я понимаю.
	`

	// order
	askingFirstNameMsg = `
	Введите имя
	(для отмены процесса, напишите "/start")
	`
	askingLastNameMsg = `
	Введите фамилию
	`
	askingPhone = `
	Введите номер телефона 
	`
	askingCompanyMsg = `
	Введите название компании, от которой будет осуществляться доставка
	`
	askingAddressMsg = `
	Введите адрес доставки
	`
	askingDateMsg = `
	Введите дату доставки
	`
	confirmationMsg = `
	Правильно ли составлен заказ? 
	(да/нет либо для отмены напишите "/start")
	`
	confirmationErrorMsg = `
	Не удалось составить заказ, попробуйте снова...
	`
	successMsg = `
	Заказ успешно оформлен! С вами свяжутся в ближайшее время!
	`
	cancelOrderMsg = `
	Заказ отменен
	`
)
