package vk

import (
	"encoding/json"
)

// Group implements VK API namespace `group`
type Group struct {
	ID          int
	Name        string
	ScreenName  string
	Type        string
	Description string
}

// Groups implements VK API namespace `groups`
type Groups struct {
	API API
}

// GroupsGetByIDResponse is response for Groups.GetByID
type GroupsGetByIDResponse []Group

// GetByID Returns information about communities by their IDs.
func (v Groups) GetByID(params GroupsGetByIDParams) (GroupsGetByIDResponse, error) {
	r, err := v.API.Request("groups.getById", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetByIDResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupsGetLongPollServerParams are params for Groups.GetLongPollServer
type GroupsGetLongPollServerParams struct {
	// Community ID
	GroupID int `url:"group_id"`
}

// GroupsGetLongPollServerResponse is response for Groups.GetLongPollServer
type GroupsGetLongPollServerResponse struct {
	// Long Poll key
	Key string `json:"key,omitempty"`
	// Long Poll server address
	Server string `json:"server,omitempty"`
	// Number of the last event
	TS int `json:"ts,omitempty"`
}

// GetLongPollServer Returns the data needed to query a Long Poll server for events
func (v Groups) GetLongPollServer(params GroupsGetLongPollServerParams) (*GroupsGetLongPollServerResponse, error) {
	r, err := v.API.Request("groups.getLongPollServer", params)
	if err != nil {
		return nil, err
	}

	var resp GroupsGetLongPollServerResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GroupsGetByIDParams are params for Groups.GetByID
type GroupsGetByIDParams struct {
	// ID or screen name of the community.
	GroupID string `url:"group_id,omitempty"`
}
