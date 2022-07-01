package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

// CustomerService CONNECT THE REQUEST COME FROM THE PRIMARY PORT(SERVICE) TO THE SECONDARY PORT(REPOSITORY)
type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomersById(string) (*dto.CustomerResponse, *errs.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, cus := range c {
		response = append(response, cus.ToDto())
	}
	return response, nil
}
func (s DefaultCustomerService) GetCustomersById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}

}
