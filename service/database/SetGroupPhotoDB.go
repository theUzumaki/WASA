package database

func (db *appdbimpl) SetGroupPhoto(id string, newpicture string) error {

	_, err := db.c.Exec("UPDATE chats SET picture = ? WHERE chatId = ?", newpicture, id)
	if err != nil {
		return err
	}

	return err
}
