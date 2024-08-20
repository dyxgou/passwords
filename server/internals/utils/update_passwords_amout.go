package utils

import (
	"database/sql"
	"errors"
)

func UpdatePasswordsAmount(db *sql.DB, userId int) error {
	updateUserPasswords := `UPDATE users SET passwords = passwords + 1 WHERE id = ?`

	stmt, err := db.Prepare(updateUserPasswords)
	if err != nil {
		return errors.New("Couldn't update user column")
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId)
	if err != nil {
		return errors.New("Couldn't exec update query")
	}

	return nil
}
