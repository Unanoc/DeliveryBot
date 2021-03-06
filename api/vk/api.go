package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	apiBaseURL = "https://api.vk.com/method/"
	apiVersion = "5.80"
)

// API is entity which can perform API Requests
// using HTTPClient client
type API interface {
	// Request performs an API request
	//
	// - method is method name
	// - params: See BuildRequestParams
	Request(method string, params interface{}) (json.RawMessage, error)

	HTTPClient() *http.Client
}

// BaseAPI is a helper type used for making requests
type BaseAPI struct {
	BaseURL     string `url:"-"`
	AccessToken string `url:"access_token,omitempty"`
	Version     string `url:"v,omitempty"`
	Language    string `url:"lang,omitempty"`

	client *http.Client
}

// BaseAPIConfig represents configuration used for BaseAPI creation
type BaseAPIConfig struct {
	// Required
	AccessToken string
	// Optional: if nil, lang is not passed in requests
	Language string
	// Optional: if nil, http.DefaultClient is used
	Client *http.Client
}

// NewBaseAPI creates and initializes a new BaseAPI instance
func NewBaseAPI(cfg BaseAPIConfig) (*BaseAPI, error) {
	if cfg.AccessToken == "" {
		return nil, errors.New("AccessToken is required")
	}

	client := cfg.Client
	if client == nil {
		client = http.DefaultClient
	}

	return &BaseAPI{
		AccessToken: cfg.AccessToken,
		Version:     apiVersion,
		BaseURL:     apiBaseURL,
		Language:    cfg.Language,

		client: client,
	}, nil
}

// APIError is a type representing errors returned by VK API
//
//easyjson:json
type APIError struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string
		Value string
	}
}

// Error implements error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("vk.APIError %d: %s", e.Code, e.Message)
}

// APIResponse is a type representing general response returned by VK API
//
//easyjson:json
type APIResponse struct {
	Error    *APIError
	Response json.RawMessage
}

// HTTPClient conforms to API interface
func (vk *BaseAPI) HTTPClient() *http.Client {
	return vk.client
}

// Request conforms to API interface
func (vk *BaseAPI) Request(method string, params interface{}) (json.RawMessage, error) {
	u, err := url.Parse(vk.BaseURL + method)
	if err != nil {
		return nil, err
	}

	q, err := BuildRequestParams(params)
	if err != nil {
		return nil, err
	}

	if baseParams, err := query.Values(vk); err == nil {
		MergeURLValues(q, baseParams)
	} else {
		return nil, err
	}

	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r, err := vk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	resp := APIResponse{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Response, nil
}
