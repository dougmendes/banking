package domain

import "banking/errs"

//DOMAIN OBJECT
type Customer struct {
	ID          int `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	//status == 1 status==0 status == ''
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
