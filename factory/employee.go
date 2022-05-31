package factory

type Employee struct {
	Name   string
	Title  string
	Salary int
}

func NewEmployee(name, title string, salary int) *Employee {
	return &Employee{
		Name:   name,
		Title:  title,
		Salary: salary,
	}
}