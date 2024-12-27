package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// func AddSubscriber(db *sql.DB, emailID string) error {
// 	_, err := db.Exec(
// 		"INSERT INTO subscribers (email_id) VALUES ($1)", emailID,
// 	)
// 	if err != nil {
// 		log.Printf("Failed to add email id: %v", err)
// 	}
// 	return err
// }

func AddSubscriberToken(db *sql.DB, emailID string, token string) (msg string, err error) {
	var isVerified bool

	log.Print("Executing query to check subscriber")
	err = db.QueryRow(`SELECT is_verified FROM subscribers WHERE email = $1`, emailID).Scan(&isVerified)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("No existing subscriber found; inserting new record")
			_, insertErr := db.Exec(`
                INSERT INTO subscribers (email, token, is_verified, created_at, updated_at)
                VALUES ($1, $2, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
                `, emailID, token)
			if insertErr != nil {
				return "", insertErr
			}
			return "Subscriber added successfully", nil
		}
		log.Printf("Error querying subscriber: %v", err)
		return "", err
	}

	if isVerified {
		log.Printf("Email %s is already verified", emailID)
		return fmt.Sprintf("Email %s is already verified", emailID), nil
	}

	log.Print("Updating subscriber token")
	_, updateErr := db.Exec(`
        UPDATE subscribers
        SET token = $1, updated_at = NOW()
        WHERE email = $2
    `, token, emailID)
	if updateErr != nil {
		log.Printf("Error updating token: %v", updateErr)
		return "", updateErr
	}

	return "Token updated successfully", nil
}

// func AddSubscriberToken(db *sql.DB, emailID string, token string) (msg string, er error) {
// 	var isVerified bool
// 	log.Print("before query 1")
// 	err := db.QueryRow(`SELECT is_verified FROM subscribers WHERE email = $1`, emailID).Scan(&isVerified)
// 	log.Print("after query 1")
// 	if err != nil {
// 		log.Print("before query 2")
// 		if err == sql.ErrNoRows {
// 			db.Exec(`
// 				INSERT INTO subscribers (email, token, is_verified, created_at, updated_at)
// 				VALUES ($1, $2, false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
// 				`, emailID, token)
// 			log.Print("after query 2")
// 			// return "", insertErr
// 		}
// 		log.Print("before ")
// 		// return "", err
// 	}
// 	if isVerified {
// 		log.Printf("email %s is already verified", emailID)
// 		return fmt.Sprintf("email %s is already verified", emailID), nil
// 	}

// 	_, updateErr := db.Exec(`
// 		UPDATE subscribers
// 		SET token = $1, updated_at = NOW()
// 		WHERE email = $2
// 	`, token, emailID)
// 	return "", updateErr

// query := `
// INSERT INTO subscribers (email, token,is_verified,created_at,updated_at)
// VALUES ($1, $2, false, NOW(), NOW())
// ON CONFLICT (email) DO UPDATE SET token=$2, created_at=NOW();
// `
// _, err := db.Exec(query, emailID, token)
// return err
// }

const tokenExpirationDuration = 15 * time.Minute

func VerifySubscriber(db *sql.DB, token string) (string, error) {
	var email string
	var createdAt time.Time

	err := db.QueryRow(`
		SELECT email, created_at
		FROM subscribers
		WHERE token = $1
		AND is_verified = false`,
		token,
	).Scan(&email, &createdAt)
	log.Print(err)
	log.Print(email)
	log.Print(createdAt, email)
	if err == sql.ErrNoRows {
		return "", errors.New("invalid token")
	}
	if err != nil {
		return "", err
	}

	if time.Since(createdAt) > tokenExpirationDuration {
		return "", errors.New("token has expired")
	}

	_, err = db.Exec(`
		UPDATE subscribers
		SET is_verified = true,
			updated_at = CURRENT_TIMESTAMP
		WHERE token = $1`,
		token)

	if err != nil {
		return "", err
	}
	return email, nil
	// query := `UPDATE subscribers SET is_verified = true, updated_at=NOW() WHERE token=$1 RETURNING email;`
	// var email string
	// err := db.QueryRow(query, token).Scan(&email)
	// if err == sql.ErrNoRows {
	// 	return "", errors.New("invalid or expired token")
	// }
	// return email, err
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
