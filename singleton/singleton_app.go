package singleton

// Singleton will use the same object, in sample
// This singleton use 1 connection to handle all
// of its query
func init() {
	db := Connect("localhost", "root", "1234")

	orderService := &OrderService{DB: db}
	orderService.Save("0001")

	orderDetailService := &OrderDetailService{DB: db}
	orderDetailService.Save("0001", "Indomie")
}