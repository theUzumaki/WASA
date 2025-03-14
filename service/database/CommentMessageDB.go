package database

import (
	"database/sql"
	"errors"
	"strconv"
	"wasatext/service/structs"
)

func (db *appdbimpl) CommentMessage(messageid string, comment string, senderid string) error {

	var comm structs.Comment
	row := db.c.QueryRow("SELECT comment_message.commentId FROM comment_message JOIN comment_user ON comment_message.commentId = comment_user.commentId WHERE comment_message.messageId = ? AND comment_user.userId = ?", messageid, senderid)
	if err := row.Scan(&comm.Id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		} else {
			return err
		}
	} else {
		_, err = db.c.Exec("UPDATE comments SET content = ? WHERE commentId = ?", comment, comm.Id)
		if err != nil {
			return err
		}
		return nil
	}

	row = db.c.QueryRow("SELECT MAX(commentId) FROM comments")
	var stringId string
	err := row.Scan(&stringId)
	if stringId == "" {
		comm.Id = 0
	} else if err != nil {
		return err
	} else {
		comm.Id, err = strconv.Atoi(stringId)
		if err != nil {
			return err
		}
	}
	comm.Id = comm.Id + 1

	_, err = db.c.Exec("INSERT INTO comments  VALUES (?, ?);", comm.Id, comment)
	if err != nil {
		return err
	}

	mess_id, err := strconv.Atoi(messageid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO comment_message VALUES (?, ?);", comm.Id, mess_id)
	if err != nil {
		return err
	}

	send_id, err := strconv.Atoi(senderid)
	if err != nil {
		return err
	}

	_, err = db.c.Exec("INSERT INTO comment_user VALUES (?, ?);", comm.Id, send_id)
	if err != nil {
		return err
	}

	return err
}
