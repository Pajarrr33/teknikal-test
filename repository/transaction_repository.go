package repository

import (
	"database/sql"
	"teknikal-test/entity"
	"teknikal-test/entity/request"
	"teknikal-test/service"

	"github.com/sirupsen/logrus"
)

type TransactionRepository interface {
	FindAll() ([]entity.Transaction, error)
	FindById(id string) (entity.Transaction, error)
	Create(request request.TransactionRequest) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	Delete(id string) error
}

type transactionRepository struct {
	db *sql.DB
}

// Create implements TransactionRepository.
func (t *transactionRepository) Create(request request.TransactionRequest) (entity.Transaction, error) {
	query := "INSERT INTO transaction (customer_id, merchant_id, amount) VALUES ($1, $2, $3) RETURNING id,created_at,updated_at"

	var transaction entity.Transaction

	err := t.db.QueryRow(query, request.CustomerId, request.MerchantId, request.Amount).Scan(&transaction.Id,&transaction.CreatedAt,&transaction.UpdatedAt)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when create transaction")
		service.SaveLog()
		return entity.Transaction{}, err
	}

	transaction.CustomerId = request.CustomerId
	transaction.MerchantId = request.MerchantId
	transaction.Amount = request.Amount

	return transaction, nil
}

// Delete implements TransactionRepository.
func (t *transactionRepository) Delete(id string) error {
	query := "DELETE FROM transaction WHERE id = $1"
	_, err := t.db.Exec(query, id)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when delete transaction")
		service.SaveLog()
		return err
	}

	return nil
}

// FindAll implements TransactionRepository.
func (t *transactionRepository) FindAll() ([]entity.Transaction, error) {
	query := "SELECT id, customer_id, merchant_id, amount, created_at, updated_at FROM transaction"

	rows, err := t.db.Query(query)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when find all transaction")
		service.SaveLog()
		return nil, err
	}

	var transactions []entity.Transaction

	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction

		err := rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.MerchantId, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)

		if err != nil {
			service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when find all transaction")
			service.SaveLog()
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when find all transaction")
		service.SaveLog()
		return nil, err
	}

	return transactions, nil
}

// FindById implements TransactionRepository.
func (t *transactionRepository) FindById(id string) (entity.Transaction, error) {
	query := "SELECT id, customer_id, merchant_id, amount, created_at, updated_at FROM transaction WHERE id = $1"

	var transaction entity.Transaction
	err := t.db.QueryRow(query, id).Scan(&transaction.Id, &transaction.CustomerId, &transaction.MerchantId, &transaction.Amount, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when find by id")
		service.SaveLog()
		return entity.Transaction{}, err
	}

	return transaction, nil
}

// Update implements TransactionRepository.
func (t *transactionRepository) Update(transaction entity.Transaction) (entity.Transaction, error) {
	query := "UPDATE transaction SET customer_id = $1, merchant_id = $2, amount = $3, updated_at = $4 WHERE id = $5 RETURNING id, customer_id, merchant_id, amount, created_at, updated_at"

	err := t.db.QueryRow(query,transaction.CustomerId,transaction.MerchantId,transaction.Amount,transaction.UpdatedAt,transaction.Id).Scan(&transaction.Id,&transaction.CustomerId,&transaction.MerchantId,&transaction.Amount,&transaction.CreatedAt,&transaction.UpdatedAt)
	if err != nil {
		service.AddLog(logrus.Fields{"location" : "transaction repository", "error" : err}, "error", "error occured when update transaction")
		service.SaveLog()
		return entity.Transaction{}, err
	}

	return transaction,nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}
