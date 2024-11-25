package repository

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"teknikal-test/service"
)

type ExpiredRepository interface {
	Insert(token string) error
	GetExpiredByToken(token string) (bool, error)
}

type expiredRepository struct {
	db *sql.DB
}

func (e *expiredRepository) Insert(token string) error {
	query := "INSERT INTO expired_token (token) VALUES ($1);"
	_, err := e.db.Exec(query, token)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "expired repository", "error" : err}, "error", "error occured when insert expired token")
		service.SaveLog()
		return err
	}

	return nil
}

func (e *expiredRepository) GetExpiredByToken(token string) (bool, error) {
    query := "SELECT token FROM expired_token WHERE token = $1;"
    var result string
    err := e.db.QueryRow(query, token).Scan(&result)
    if err == sql.ErrNoRows {
        return false, nil // Token does not exist
    } else if err != nil {
        service.AddLog(logrus.Fields{"location": "expired repository", "error": err}, "error", "error occurred when getting expired token")
        service.SaveLog()
        return false, err // Other errors
    }
    return true, nil // Token exists
}


func NewExpiredRepository(db *sql.DB) ExpiredRepository {
	return &expiredRepository{db}
}