package shopify

import (
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	Url "net/url"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

type Shopify struct {
	store    string
	apiKey   string
	pass     string
}

const (
	domain = ".myshopify.com/admin"
)

// Creates a New Shopify Store API object with the
// store, apiKey and pass of your store.
// Usage:
// 	shopify.New("mystore", "XXX","YYY")
func New(store, apiKey, pass string) Shopify {

	shop := Shopify{store: store, apiKey: apiKey, pass: pass}
	//	fmt.Println("[New] Creating Shopify client with: ", store, apiKey, pass)

	return shop
}

// Creates target URL for making a Shopify Request
// to a given endpoint
func (shopify *Shopify) createTargetURL(endpoint string) string {
	result := fmt.Sprintf("https://%s:%s@%s%s/%s.json", shopify.apiKey, shopify.pass, shopify.store, domain, endpoint)
	return result
}

// Extracts Json Bytes from map[string]interface
func getJsonBytesFromMap(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Invalid data object, can't parse to json:")
		fmt.Println("Error:", err)
		fmt.Println("Data:", data)
		return nil, err
	}
	return jsonData, nil
}

// Creates a new Request to Shopify and returns
// the response as a map[string]interface{}.
// method: GET/POST/PUT - string
// url: target endpoint like "products" - string
// data: content to be sent with the request
// Usage: shopify.request("GET","products",nil)
func (shopify *Shopify) Request(method, endpoint string, data map[string]interface{}) ([]byte, []error) {
	//	fmt.Println("[request] Arguments: ", method, endpoint, data)

	jsonData, _ := getJsonBytesFromMap(data)
	//	fmt.Println("[request] data: ", string(jsonData))

	targetURL := shopify.createTargetURL(endpoint)
	//	fmt.Println("[request] targetUrl: ", targetURL)

	request := gorequest.New()
	request.Get(targetURL)

	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}

	_, body, errs := request.End()


	return []byte(body), errs
}

// Makes a GET request to shopify with the
// given endpoint.
// Usage:
// shopify.Get("products/5.json")
// shopify.Get("products/5/variants.json")
func (shopify *Shopify) Get(endpoint string) ([]byte, []error) {

	targetUrl := shopify.createTargetURL(endpoint)

	request := gorequest.New()
	_, body, errs := request.Get(targetUrl).End()

	return []byte(body), errs
}

// Makes a POST request to shopify with the
// given endpoint and data.
// Usage:
// shopify.Post("products", map[string]interface{} = product data map)
func (shopify *Shopify) Post(endpoint string, data map[string]interface{}) ([]byte, []error) {

	targetUrl := shopify.createTargetURL(endpoint)
	jsonData, err := getJsonBytesFromMap(data)
	if err != nil {
		return nil, []error{err}
	}

	request := gorequest.New()
	request.Post(targetUrl)
	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}
	_, body, errs := request.End()

	return []byte(body), errs
}

// Makes a PUT request to shopify with the
// given endpoint and data.
// Usage:
// shopify.Put("products", map[string]interface{} = product data map)
func (shopify *Shopify) Put(endpoint string, data map[string]interface{}) ([]byte, []error) {

	targetUrl := shopify.createTargetURL(endpoint)
	jsonData, err := getJsonBytesFromMap(data)
	if err != nil {
		return nil, []error{err}
	}

	request := gorequest.New()
	request.Put(targetUrl)
	if jsonData != nil && data != nil {
		request.Send(string(jsonData))
	}
	_, body, errs := request.End()

	return []byte(body), errs
}

// Makes a DELETE request to shopify with the
// given endpoint.
// Usage:
// shopify.Delete("products/5.json")
func (shopify *Shopify) Delete(endpoint string) ([]byte, []error) {

	targetUrl := shopify.createTargetURL(endpoint)

	request := gorequest.New()
	_, body, errs := request.Delete(targetUrl).End()

	return []byte(body), errs
}
