package database

func (db *appdbimpl) LeaveGroup(groupId string) error {

	_, err := db.c.Exec("DELETE FROM chatsCollection WHERE chatsId = ?", groupId)

	return err
}
