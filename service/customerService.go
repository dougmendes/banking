package service

import (
	"banking/domain"
	"banking/errs"
)

// CustomerService CONNECT THE REQUEST COME FROM THE PRIMARY PORT(SERVICE) TO THE SECONDARY PORT(REPOSITORY)
type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomersById(string) (*domain.Customer, *errs.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetCustomersById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}

}
