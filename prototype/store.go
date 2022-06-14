package prototype

import "fmt"

type Store struct {
	Name     string
	City     string
	Country  string
	Category string
}

func init() {
	store1 := Store{
		Name:     "Toko X",
		City:     "Jakarta",
		Country:  "Indonesia",
		Category: "Gadget",
	}
	// Clone store2 and change several specific field
	store2 := store1
	store2.Name = "Toko Y"

	fmt.Print(store1, store2)
}