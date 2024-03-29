package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	// FirstName   string `json:"first_name"`
	// LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"-"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"-"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		// way 1:
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	// way 2:
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}

func (users Users) Marshall(isPublic bool) interface{} {
	// first approche
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}