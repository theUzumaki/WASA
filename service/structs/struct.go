package structs

import (
	"time"
)

type Text struct {
	Text string `json:"content"`
}

type Id struct {
	Id string `json:"id"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Comment struct {
	Content string `json:"content"`
	Id      int    `json:"id"`
}

type CommentSender struct {
	Sender  User    `json:"sender"`
	Comment Comment `json:"comment"`
}

type Message struct {
	Id          int             `json:"id"`
	ChatId      int             `json:"chat_id"`
	Sender      User            `json:"sender"`
	CommSenders []CommentSender `json:"comm_senders"`
	Date        time.Time       `json:"date"`
	Content     string          `json:"content"`
	Checkmark   bool            `json:"checkmark"`
	ReplyId     int             `json:"reply_id" default:"-1"`
}

type Group struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Chat struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Members  []User    `json:"members"`
	Messages []Message `json:"messages"`
	Picture  string    `json:"picture"`
}
