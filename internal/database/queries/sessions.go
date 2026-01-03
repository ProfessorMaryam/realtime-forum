package queries

import (
	"github.com/google/uuid"
	"real-time-forum/internal/database"
	"time"
	"net/http"
	"errors"

)


func GenerateToken() (string, error) {
	u, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func AddSession(email string) (http.Cookie, error) {

	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE email = ?",
		email).Scan(&userID)
	if err != nil {
		return http.Cookie{}, err
	}

	value, err := GenerateToken()
	expires := time.Now().Add(7 * 24 * time.Hour)
	created := time.Now()

	if err != nil {
		return http.Cookie{}, err
	}

	_, err = database.DB.Exec("INSERT INTO sessions (session_id, user_id, user_agent, real_ip, expires_at, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		value, userID, "agent","0.0.0.0", expires, created)
	if err != nil {
		return http.Cookie{}, errors.New("database insert error")
	}

	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    value,
		Expires:  expires,
		HttpOnly: true,
	}

	return cookie, nil
}

func DeletePastSessions(userID int) error {
	_, err := database.DB.Exec("UPDATE sessions SET valid = false WHERE userID = ?",
		userID)
	if err != nil {
		return err
	}
	return nil
}

func ValidSession(value string) bool {

	var expirey time.Time
	var valid bool
	err := database.DB.QueryRow("SELECT expiresAt, valid FROM sessions WHERE sessionID = ?",
		value).Scan(&expirey, &valid)
	if err != nil {
		return false
	}

	return time.Now().Before(expirey) && valid
}