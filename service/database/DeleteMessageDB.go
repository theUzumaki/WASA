package database

func (db *appdbimpl) DeleteMessage(messageid string) error {

	_, err := db.c.Exec("DELETE FROM messages WHERE messageId = ?", messageid)
	if err != nil {
		return err
	}

	return err
}
