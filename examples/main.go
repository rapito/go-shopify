package main

import (
	"github.com/rapito/go-shopify/shopify"
	simplejson "github.com/bitly/go-simplejson"
	"fmt"
)

const (
	store  = "your-store-domain-name-here"
	apiKey = "your-api-key-here"
	pass   = "your-secret-pass-here"
)

// Create a new shopify object with your store
// domain, api key and password
var shop = shopify.New(store, apiKey, pass)

func main() {

	fetchAllProducts()
	fetchOneProduct(395386591)
}

func fetchOneProduct(id int64) {
	fmt.Println("[fetchOneProduct]")

	// Call any of the api CRUD methods
	endpoint := fmt.Sprintf("products/%v", id)
	result, _ := shop.Get(endpoint)

	fmt.Println("Result")
	fmt.Println(string(result))

	fmt.Println("===============")
}


func fetchAllProducts() {
	fmt.Println("[fetchAllProducts]")

	// Call any of the api CRUD methods
	result, _ := shop.Get("products")

	// Do what you want with the []byte response.
	// In this case we are using simplejson to handle it.
	jsonData, _ := simplejson.NewJson(result)

	products := jsonData.Get("products")
	product := products.GetIndex(0)

	title, _ := product.Get("title").String()

	fmt.Println("Full result: ");
	fmt.Println(string(result));

	fmt.Println("First Product title: ");
	fmt.Println(title);

	fmt.Println("===============")
}
