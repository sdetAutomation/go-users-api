package users

import "encoding/json"

// PublicUser ...
type PublicUser struct {
	ID 			int64	`json:"id"`
	DataCreated string	`json:"date_created"`
	Status		string	`json:"status"`
}

// PrivateUser ...
type PrivateUser struct {
	ID 			int64	`json:"id"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email 		string	`json:"email"`
	DataCreated string	`json:"date_created"`
	Status		string	`json:"status"`
}

// Marshal ...
func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

// Marshal ....
func (user *User) Marshal(isPublic bool) interface {} {
	if isPublic {
		return PublicUser {
			ID: 			user.ID,
			DataCreated: 	user.DataCreated,
			Status: 		user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}