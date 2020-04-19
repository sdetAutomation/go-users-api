package users

// User ...
type User struct {
	ID 			int64	`json:"id"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Email 		string	`json:"email"`
	DataCreated string	`json:"date_created"`
}
