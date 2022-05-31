package factory

import "fmt"

func init() {
	manager1 := createManager("Manager1")
	staff1 := createStaff("Staff1")

	fmt.Println(manager1, staff1)
}