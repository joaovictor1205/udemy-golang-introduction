package main

import "fmt"

type item struct {
	product_id int
	number     int
	price      float64
}

type request struct {
	user_id int
	items   []item
}

func (r request) total() float64 {
	total := 0.0

	for _, item := range r.items {
		total += item.price * float64(item.number)
	}

	return total
}

func main() {

	req := request{
		user_id: 1,

		items: []item{
			item{
				product_id: 1,
				number:     5,
				price:      2.5,
			},
			item{
				product_id: 2,
				number:     6,
				price:      7.5,
			},
		},
	}

	fmt.Println(req)

	fmt.Printf("Total price: %.2f\n", req.total())
}
