package vk

// CallbackEvent is base event
type CallbackEvent struct {
	// ID of group this event occured in
	GroupID int
	// Secret for Callback API
	Secret string
	// VK Event
	Event interface{}
}

// MessageNew -- new message is recieved
//
//easyjson:json
type MessageNew struct {
	Message
}

// Confirmation is used in Callback API.
// It requires listener to reply with Confirmation token instead of normal "ok".
//
//easyjson:json
type Confirmation struct{}
