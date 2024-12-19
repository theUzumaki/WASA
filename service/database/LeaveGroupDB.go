package database

func (db *appdbimpl) LeaveGroup(groupId string, userId string) error {

	_, err := db.c.Exec("DELETE FROM chat_user WHERE chatId = ? AND userId = ?", groupId, userId)

	return err
}
