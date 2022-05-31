package builder

type Customer struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	// New Property
	Adress string
	Age    int
}

func NewCustomer(customer *Customer) *Customer {
	return customer
}