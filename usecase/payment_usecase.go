package usecase

import (
	"teknikal-test/entity"
	"teknikal-test/entity/request"
	"teknikal-test/repository"
	"teknikal-test/service"
	"github.com/sirupsen/logrus"
	"fmt"
)

type PaymentUsecase interface {
	Payment(request request.TransactionRequest) (entity.Transaction, error)
}

type paymentUsecase struct {
	CustomerRepositry     repository.CustomerRepository
	MerchantRepository	  repository.MerchantRepository
	TransactionRepository repository.TransactionRepository
}

func (p *paymentUsecase) Payment(request request.TransactionRequest) (entity.Transaction, error) {
	customer, err := p.CustomerRepositry.FindById(request.CustomerId)
	if err != nil {
		err := fmt.Errorf("customer not found")
		return entity.Transaction{}, err
	}

	merchant, err := p.MerchantRepository.FindById(request.MerchantId)
	if err != nil {
		err := fmt.Errorf("merchant not found")
		return entity.Transaction{}, err
	}

	transaction, err := p.TransactionRepository.Create(request)
	if err != nil {
		err := fmt.Errorf("failed to create transaction")
		return entity.Transaction{}, err
	}

	service.AddLog(logrus.Fields{"customer" : customer.Name, "merchant" : merchant.Name, "amount" : request.Amount},"info","user has been paid")
	service.SaveLog()
	return transaction, nil
}

func NewPaymentUsecase(transactionRepository repository.TransactionRepository, customerRepository repository.CustomerRepository, merchantRepository repository.MerchantRepository) PaymentUsecase {
	return &paymentUsecase{TransactionRepository: transactionRepository,CustomerRepositry : customerRepository, MerchantRepository : merchantRepository}
}
