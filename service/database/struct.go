package database

import "time"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Id      int       `json:"id"`
	ChatId  int       `json:"chat_id"`
	Sender  User      `json:"sender"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
	Comment string    `json:"comment"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Chat struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Members  []User    `json:"members"`
	Messages []Message `json:"messages"`
}
