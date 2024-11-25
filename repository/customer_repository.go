package repository

import (
	"database/sql"
	"teknikal-test/entity"
	"teknikal-test/entity/request"
	"teknikal-test/service"

	"github.com/sirupsen/logrus"
)

type CustomerRepository interface {
	FindAll() ([]entity.Customer, error)
	FindById(id string) (entity.Customer, error)
	FindByEmail(email string) (entity.Customer, error)
	Create(request request.RegisterRequest) (entity.Customer, error)
	Update(customer entity.Customer) (entity.Customer, error)
	Delete(id string) error
}

type customerRepository struct {
	db *sql.DB
}

// FindByEmail implements CustomerRepository.
func (c *customerRepository) FindByEmail(email string) (entity.Customer, error) {
	query := "SELECT id, name, email, password FROM customers WHERE email = $1"

	var customer entity.Customer
	err := c.db.QueryRow(query, email).Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Password)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when find by email")
		service.SaveLog()
		return entity.Customer{}, err
	}

	return customer, nil
}

// Create implements CustomerRepository.
func (c *customerRepository) Create(request request.RegisterRequest) (entity.Customer, error) {
	query := "INSERT INTO customers (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email,  balance, created_at, updated_at"

	var customer entity.Customer
	err := c.db.QueryRow(query, request.Name, request.Email, request.Password).Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when create customer")
		service.SaveLog()
		return entity.Customer{}, err
	}

	return customer, nil
}

// Delete implements CustomerRepository.
func (c *customerRepository) Delete(id string) error {
	query := "DELETE FROM customers WHERE id = $1"
	_, err := c.db.Exec(query, id)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when delete customer")
		service.SaveLog()
		return err
	}

	return nil
}

// FindAll implements CustomerRepository.
func (c *customerRepository) FindAll() ([]entity.Customer, error) {
	query := "SELECT id, name, email, balance, created_at, updated_at FROM customers"
	rows, err := c.db.Query(query)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when find all customer")
		service.SaveLog()
		return nil, err
	}

	var customers []entity.Customer

	defer rows.Close()
	for rows.Next() {
		var customer entity.Customer

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)

		if err != nil {
			service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when find all customer")
			service.SaveLog()
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when find all customer")
		service.SaveLog()
		return nil, err
	}

	return customers, nil
}

// FindById implements CustomerRepository.
func (c *customerRepository) FindById(id string) (entity.Customer, error) {
	query := "SELECT id, name, email, balance, created_at, updated_at FROM customers WHERE id = $1"

	var customer entity.Customer
	err := c.db.QueryRow(query, id).Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when find by id")
		service.SaveLog()
		return entity.Customer{}, err
	}

	return customer, nil
}

// Update implements CustomerRepository.
func (c *customerRepository) Update(customer entity.Customer) (entity.Customer, error) {
	query := "UPDATE customers SET name = $1, email = $2,balance = $3, updated_at = $4 WHERE id = $5 RETURNING id, name, email, balance, created_at, updated_at"

	err := c.db.QueryRow(query, customer.Name, customer.Email, customer.Balance, customer.UpdatedAt, customer.Id).Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt)

	if err != nil {
		service.AddLog(logrus.Fields{"location" : "customer repository", "error" : err}, "error", "error occured when update customer")
		service.SaveLog()
		return entity.Customer{}, err
	}

	return customer, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db}
}
