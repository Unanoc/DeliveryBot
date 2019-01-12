package vk

// UsersGetParams are params for Users.Get
type UsersGetParams struct {
	// User IDs or screen names ('screen_name'). By default, current user ID.
	UserIDs CSVStringSlice `url:"user_ids,omitempty"`
	// Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities',
	Fields CSVStringSlice `url:"fields,omitempty"`
	// Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
	NameCase string `url:"name_case,omitempty"`
}

// UsersGetResponse is response for Users.Get
type UsersGetResponse []User

// User is a structure of VK User
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       int    `json:"sex,omitempty"`
	BirthDate string `json:"bdate,omitempty"`
}

// CSVStringSlice is a string slice which gets encoded
// as comma-separated string
// API methods often accept arrays of strings, which
// should be encoded as one comma-separated parameter
// This is helper type which implements query.Encoder
type CSVStringSlice []string
