package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Products struct {
	ID        string
	Name      string
	CreatedAt int
}

func main() {
	// Read txt file
	data := mustRead("data/conversation.txt")
	fmt.Println("Txt file: ", string(data))

	// Read JSON file
	data = mustRead("data/products.json")

	var products []Products
	err := json.Unmarshal(data, &products)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Print products with field name: %+v\n", products)
	fmt.Println("First product:", products[0])
}

func mustRead(filename string) []byte {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}
