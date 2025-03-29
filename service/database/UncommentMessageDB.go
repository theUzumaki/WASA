package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UncommentMessage(messageid string, senderid string) error {

	row := db.c.QueryRow("SELECT comments.commentId FROM comments JOIN comment_message ON comment_message.commentId = comments.commentId JOIN comment_user ON comment_user.commentId = comments.commentId WHERE comment_message.messageId = ? AND comment_user.userId = ?", messageid, senderid)

	var id string
	err := row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("no comments")
	} else if err != nil {
		return err
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE comments.commentId = ?", id)
	if err != nil {
		return err
	}

	return err
}
