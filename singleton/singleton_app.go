package singleton

func init() {
	orderService := new(OrderService)
	orderService.Save("0001")

	orderDetailService := new(OrderDetailService)
	orderDetailService.Save("0001", "Indomie")
}