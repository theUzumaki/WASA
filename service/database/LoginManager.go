package database

import (
	"fmt"
	"log"
	"strconv"
)

func (db *appdbimpl) LoginManager(user User) (string, error) {

	fmt.Println("NAME SEARCHING FOR: <" + user.Name + ">")
	row := db.c.QueryRow("SELECT * FROM users WHERE userName = ?", user.Name)
	if row.Err() != nil {
		log.Fatal(row.Err())
	}

	void1 := ""
	void2 := ""
	err := row.Scan(&void1, &void2)
	if err == row.Err() {
		return "user exist", nil
	}

	row = db.c.QueryRow("SELECT MAX(userId) FROM users")
	var id int
	row.Scan(&id)
	user.Id = id + 1

	fmt.Println("id: " + strconv.Itoa(user.Id) + " name: " + user.Name)
	_, err = db.c.Exec("INSERT INTO users (userId, userName) VALUES (?, ?);", user.Id, user.Name)
	if err != nil {
		log.Fatal(err)
	}

	row = db.c.QueryRow("SELECT userName FROM users WHERE userName = ?", user.Name)

	var name string
	row.Scan(&name)
	result := "ROW: <" + name + ">"
	fmt.Println(result)
	return "", nil
}
