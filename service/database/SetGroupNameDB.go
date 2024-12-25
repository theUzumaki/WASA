package database

func (db *appdbimpl) SetGroupName(id string, newname string) error {

	_, err := db.c.Exec("UPDATE chats SET chatName = ? WHERE chatId = ?", newname, id)

	return err
}
