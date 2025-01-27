package database

func (db *appdbimpl) AddToGroup(userid string, chatid string, newuserid string) error {

	_, err := db.c.Exec("INSERT INTO chat_user VALUES (?,?)", chatid, newuserid)
	if err != nil {
		return err
	}

	return err
}
