package repository

import (
	"database/sql"
	"teknikal-test/entity"
)

type MerchantRepository interface {
	FindById(id string) (entity.Merchant, error)
}

type merchantRepository struct {
	db *sql.DB
}

func (r *merchantRepository) FindById(id string) (entity.Merchant, error) {
	var merchant entity.Merchant

	query := "SELECT id, name, category, contact FROM merchant WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&merchant.Id, &merchant.Name, &merchant.Category, &merchant.Contact)

	if err != nil {
		return entity.Merchant{}, err
	}

	return merchant, nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}