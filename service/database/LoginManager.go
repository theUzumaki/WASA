package database

import (
	"fmt"
	"log"
	"strconv"
)

func (db *appdbimpl) LoginManager(user User) (string, error) {

	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", user.Name)
	if row.Scan(nil) == nil {
		return "user exist", nil
	}

	row = db.c.QueryRow("SELECT MAX(userId) FROM users")
	var id int
	row.Scan(&id)
	user.Id = id + 1

	fmt.Println("id: " + strconv.Itoa(user.Id) + " name: " + user.Name)
	_, err := db.c.Query("INSERT INTO users (userId, userName) VALUES (?, ?)", user.Id, user.Name)
	if err != nil {
		fmt.Println("FIRST ERROR")
		log.Fatal(err)
	}

	rowcheck, errcheck := db.c.Query("SELECT * FROM users")
	if errcheck != nil {
		fmt.Println("SECOND ERROR")
		log.Fatal(errcheck)
	}
	var name string
	rowcheck.Scan(&name)
	result := "ROW: " + name
	fmt.Println(result)
	return "", nil
}
