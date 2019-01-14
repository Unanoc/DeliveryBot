package vk

// UsersGetParams are params for Users.Get
type UsersGetParams struct {
	UserIDs string `url:"user_ids,omitempty"`
	Fields string `url:"fields,omitempty"`
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
