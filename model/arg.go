package model

type SendMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
