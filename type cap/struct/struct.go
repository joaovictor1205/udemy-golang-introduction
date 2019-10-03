package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {

	var product1 product

	product1 = product{
		name:     "Test",
		price:    10,
		discount: 0.05,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{
		name:     "Test2",
		price:    5,
		discount: 0.1,
	}

	fmt.Println(product2, product2.priceWithDiscount())

	product3 := product{name: "Test3", price: 15, discount: 0.43}

	fmt.Println(product3, product3.priceWithDiscount())

}
