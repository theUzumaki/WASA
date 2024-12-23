/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	LoginManager(user User) (int, string, error)

	SetName(oldname string, newname string) error

	NewChat(id string, chat Chat) error

	LeaveGroup(idgroup string, iduser string) error

	AddToGroup(iduser string, idchat string, newiduser string) error

	GetMyConversations(id string) ([]Chat, error)

	GetConversation(id string) (Chat, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableNum int
	err := db.QueryRow(`SELECT COUNT(name) FROM sqlite_master WHERE type='table';`).Scan(&tableNum)
	if err != nil {
		return nil, err
	}
	if tableNum != 6 {
		_, err = db.Exec(users)
		if err != nil {
			return nil, fmt.Errorf("users table creation error %w", err)
		}
		_, err = db.Exec(messages)
		if err != nil {
			return nil, fmt.Errorf("messages table creation error %w", err)
		}
		_, err = db.Exec(chats)
		if err != nil {
			return nil, fmt.Errorf("chats table creation error %w", err)
		}
		_, err = db.Exec(chat_message)
		if err != nil {
			return nil, fmt.Errorf("chat_message table creation error %w", err)
		}
		_, err = db.Exec(chat_user)
		if err != nil {
			return nil, fmt.Errorf("chat_user table creation error %w", err)
		}
		_, err = db.Exec(message_user)
		if err != nil {
			return nil, fmt.Errorf("message_user table creation error %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
