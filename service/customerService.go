package service

import (
	"banking/domain"
	"banking/errs"
)

// CustomerService CONNECT THE REQUEST COME FROM THE PRIMARY PORT(SERVICE) TO THE SECONDARY PORT(REPOSITORY)
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomersById(string) (*domain.Customer, *errs.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) GetCustomersById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}

}
