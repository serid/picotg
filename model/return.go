package model

type Message struct {
	MessageId int  `json:"message_id"`
	From      User `json:"from"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
