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

func GetProjects(db *sql.DB, username string) (link string, err error) {
	err = db.QueryRow(`SELECT projectlink1 FROM users WHERE  username= $1`, username).Scan(&link)

	if err != nil {
		return "", err
	}
	return link, nil
}
