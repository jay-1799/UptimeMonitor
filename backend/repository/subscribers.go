package repository

import (
	"database/sql"
	"errors"
	"log"
)

func AddSubscriber(db *sql.DB, emailID string) error {
	_, err := db.Exec(
		"INSERT INTO subscribers (email_id) VALUES ($1)", emailID,
	)
	if err != nil {
		log.Printf("Failed to add email id: %v", err)
	}
	return err
}

func AddSubscriberToken(db *sql.DB, emailID string, token string) error {
	query := `
	INSERT INTO subscribers (email, token,is_verified,created_at,updated_at) 
	VALUES ($1, $2, false, NOW(), NOW())
	ON CONFLICT (email) DO UPDATE SET token=$2, created_at=NOW();
	`
	_, err := db.Exec(query, emailID, token)
	return err
}

func VerifySubscriber(db *sql.DB, token string) (string, error) {
	query := `UPDATE subscribers SET is_verified = true, updated_at=NOW() WHERE token=$1 RETURNING email;`
	var email string
	err := db.QueryRow(query, token).Scan(&email)
	if err == sql.ErrNoRows {
		return "", errors.New("invalid or expired token")
	}
	return email, err
}

func FetchAllSubscribers(db *sql.DB) ([]string, error) {
	query := `SELECT email FROM subscribers WHERE is_verified=true;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []string
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return emails, nil
}
