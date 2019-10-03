package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	product1 := product{
		ID:    1,
		Nome:  "Product 1",
		Preco: 1.99,
		Tags:  []string{"promotion", "eletronic"},
	}
	product1Json, _ := json.Marshal(product1)
	fmt.Println(string(product1Json))

	var product2 product
	jsonString := `{"id":"2", "nome":"Product 2", "preco":"2.99", "tags":["Test","TestTest"]}`
	json.Unmarshal([]byte(jsonString), &product2)
	fmt.Println(product2.Tags[1])

}
