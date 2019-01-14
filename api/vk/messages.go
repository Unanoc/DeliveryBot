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
	UserID int `url:"user_id,omitempty"`
	RandomID int `url:"random_id,omitempty"`
	PeerID int `url:"peer_id,omitempty"`
	Domain string `url:"domain,omitempty"`
	ChatID int `url:"chat_id,omitempty"`
	Message      string `url:"message,omitempty"`
	Notification bool   `url:"notification,omitempty"`
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
