package api

import (
	"time"
	"wasatext/service/database"
)

type id struct {
	Id string `json:"id"`
}

type text struct {
	Text string `json:"content"`
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (user User) ApiUserToDB() database.User {
	newUser := database.User{
		Id:   user.Id,
		Name: user.Name,
	}
	return newUser
}

type Message struct {
	Id      int       `json:"id"`
	ChatId  int       `json:"chat_id"`
	Sender  User      `json:"sender"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
	Comment string    `json:"comment"`
}

func (message Message) ApiMessageToDB() database.Message {
	newMessage := database.Message{
		Id:      message.Id,
		ChatId:  message.ChatId,
		Sender:  message.Sender.ApiUserToDB(),
		Date:    message.Date,
		Content: message.Content,
		Comment: message.Comment,
	}
	return newMessage
}

/*
type Group struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (group Group) ApiGroupToDB() database.Group {
	newGroup := database.Group{
		Id:   group.Id,
		Name: group.Name,
	}
	return newGroup
}
*/

type Chat struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Members  []User    `json:"members"`
	Messages []Message `json:"messages"`
	Picture  string    `json:"picture"`
}

func (chat Chat) ApiChatToDB() database.Chat {

	var s_newMembers []database.User = make([]database.User, len(chat.Members))
	for i := 0; i < len(chat.Members); i++ {
		s_newMembers[i] = chat.Members[i].ApiUserToDB()
	}

	var s_newMessages []database.Message = make([]database.Message, len(chat.Messages))
	for i := 0; i < len(chat.Messages); i++ {
		s_newMessages = append(s_newMessages, chat.Messages[i].ApiMessageToDB())
	}

	newChat := database.Chat{
		Id:       chat.Id,
		Name:     chat.Name,
		Members:  s_newMembers,
		Messages: s_newMessages,
	}
	return newChat
}
