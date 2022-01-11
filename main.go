package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ProductList struct {
	ID        string
	Name      string
	CreatedAt int
}

func main() {
	// Print time
	a := time.Now()
	fmt.Println(a)

	// Read txt file
	data := mustRead("data/conversation.txt")
	fmt.Println(string(data))

	// Read JSON file
	data = mustRead("data/productlist.json")

	var products []ProductList
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
