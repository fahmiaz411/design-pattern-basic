package factory

func createStaff(name string) *Employee {
	return NewEmployee(name, "Staff", 5000000)
}

func createManager(name string) *Employee {
	return NewEmployee(name, "Manager", 10000000)
}