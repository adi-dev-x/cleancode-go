package user

import "net/url"

type UserRegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"Phone"`
}

func (u *UserRegisterRequest) Valid() url.Values {
	err := url.Values{}

	if len(u.FirstName) < 2 {
		err.Add("first_name", "invalid first name")
	}

	if len(u.Password) < 6 {
		err.Add("password", "password must be greater than 6")
	}

	return err
}
