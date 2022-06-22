package domain

//ADAPTER
type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1, "Douglas", "Belo Horizonte", "30494220", "01/07/1992", "1"},
		{2, "Maria", "Belo Horizonte", "30494220", "11/03/1994", "1"},
	}
	return CustomerRepositoryStub{customers}
}
