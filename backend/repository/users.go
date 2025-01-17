package repository

import "database/sql"

func CreateStatusPage(db *sql.DB, emailID string, username string, projectlink1 string) (msg string, err error) {

	_, err = db.Exec(`INSERT INTO users (emailID, username,projectlink1)
						VALUES ($1, $2, $3)`, emailID, username, projectlink1)

	if err != nil {
		return "", err
	}
	return "created status page successfully", nil
}
