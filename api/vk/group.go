package vk

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
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
//easyjson:json
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
	// IDs or screen names of communities.
	GroupIDs CSVStringSlice `url:"group_ids,omitempty"`
	// ID or screen name of the community.
	GroupID string `url:"group_id,omitempty"`
	// Group fields to return.
	Fields CSVStringSlice `url:"fields,omitempty"`
}

// CSVStringSlice is a string slice which gets encoded
// as comma-separated string
// API methods often accept arrays of strings, which
// should be encoded as one comma-separated parameter
// This is helper type which implements query.Encoder
type CSVStringSlice []string

// EncodeValues conforms to query.Encoder inteface
func (csv CSVStringSlice) EncodeValues(key string, v *url.Values) error {
	// Doesnt handle strings with , in them, but there's no use for that
	if val := []string(csv); len(val) != 0 {
		encoded := strings.Join(val, ",")
		v.Set(key, encoded)
	}

	return nil
}

// CSVIntSlice is an int slice which gets encoded
// as comma-separated string
// API methods sometimes accept arrays of ints, which
// should be encoded as one comma-separated parameter
// This is helper type which implements query.Encoder
type CSVIntSlice []int

// EncodeValues conforms to query.Encoder inteface
func (csv CSVIntSlice) EncodeValues(key string, v *url.Values) error {
	strCSV := make(CSVStringSlice, len(csv))

	for i, v := range csv {
		strCSV[i] = strconv.Itoa(v)
	}

	return strCSV.EncodeValues(key, v)
}

// BoolInt is bool type which conforms to easyjson.Unmarshaler interface
// and unmarshals from VK's favorite 1/0 int bools
type BoolInt bool
