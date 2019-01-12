package vk

import (
	"strconv"
)

//easyjson:json
type Message struct {
	ID       int    `json:"id"`
	Date     int    `json:"date"`
	PeerID   int    `json:"peer_id"`
	FromID   int    `json:"from_id"`
	Text     string `json:"text"`
	RandomID int    `json:"random_id"`
}

// Messages implements VK API namespace `messages`
type Messages struct {
	API API
}

// MessagesSendParams are params for Messages.Send
type MessagesSendParams struct {
	// User ID (by default — current user).
	UserID int `url:"user_id,omitempty"`
	// Unique identifier to avoid resending the message.
	RandomID int `url:"random_id,omitempty"`
	// Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
	PeerID int `url:"peer_id,omitempty"`
	// User's short address (for example, 'illarionov').
	Domain string `url:"domain,omitempty"`
	// ID of conversation the message will relate to.
	ChatID int `url:"chat_id,omitempty"`
	// (Required if 'attachments' is not set.) Text of the message.
	Message      string `url:"message,omitempty"`
	Notification bool   `url:"notification,omitempty"`
	// Group ID (for group messages with group access token)
	GroupID int `url:"group_id,omitempty"`
}

// MessagesSendResponse is response for Messages.Send
// Message ID
type MessagesSendResponse int

// Send Sends a message.
func (v Messages) Send(params MessagesSendParams) (MessagesSendResponse, error) {
	r, err := v.API.Request("messages.send", params)
	if err != nil {
		return 0, err
	}

	var resp MessagesSendResponse

	var cnv int
	cnv, err = strconv.Atoi(string(r))
	resp = MessagesSendResponse(cnv)

	if err != nil {
		return 0, err
	}
	return resp, nil
}
