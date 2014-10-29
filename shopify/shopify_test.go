package shopify

// Import Testing frameworks needed
import (
	"testing"
	"fmt"
	"github.com/bmizerany/assert"
	simplejson "github.com/bitly/go-simplejson"
	"encoding/json"
)

// Create out store variables for easy access
const (
	store  = "your-store-domain-name-here"
	apiKey = "your-api-key-here"
	pass   = "your-secret-pass-here"
)

// We declare out shop here just to reuse
// it later on.
var shop = New(store, apiKey, pass)
var objIdToDelete int64

// Should create a new Store.
func TestNew(t *testing.T) {

	if shop.store != store || shop.apiKey != apiKey || shop.pass != pass {
		t.Errorf("Error creating client, was suppposed to have store:$v apiKey:$v pass:$v", store, apiKey, pass)
	}
}

// Should make a new Request
func TestRequest(t *testing.T) {

	result, error := shop.Request("GET", "products", nil)

	assert.T(t, error == nil, "should be null")
	assert.T(t, result != nil, "shouldnt be null")
}

// Should make a new Get Request
func TestGet(t *testing.T) {

	products, errors := shop.Get("products/350748043/variants/819701439")

	js, err := simplejson.NewJson(products)

	fmt.Println(js)
	fmt.Println(err)


	assert.T(t, errors == nil)
	assert.T(t, products != nil)


	//	product := products["products"].([]interface{})
	//	fmt.Println(product[0]["vendor"])
	//	fmt.Println(products)



}

// Should make a new Post Request
func TestPost(t *testing.T) {

	str := ` {"product": {"title": "MyProduct","body_html": "<strong>Good snowboard!</strong>","vendor": "Burton","product_type": "Snowboard","variants": [  {	"option1": "First",	"price": "10.00",	"sku": 123  },  {	"option1": "Second",	"price": "20.00",	"sku": "123"  }]}}  `
	var data map[string]interface{}
	json.Unmarshal([]byte(str), &data)

	result, errors := shop.Post("products", data)

	js, err := simplejson.NewJson(result)

	fmt.Println(js)
	fmt.Println(err)

	assert.T(t, errors == nil)
	assert.T(t, result != nil)

	title, _ := js.Get("product").Get("title").String()

	assert.T(t, title == "MyProduct")

	id, _ := js.Get("product").Get("id").Int64()
	objIdToDelete = id
}

// Should make a new Put Request
func TestPut(t *testing.T) {

	str := ` {"product": {"title": "Edited" } } `
	var data map[string]interface{}
	json.Unmarshal([]byte(str), &data)

	endpoint := fmt.Sprintf("products/%v", objIdToDelete)
	result, errors := shop.Put(endpoint, data)

	js, err := simplejson.NewJson(result)

	fmt.Println(js)
	fmt.Println(err)

	title, _ := js.Get("product").Get("title").String()

	assert.T(t, errors == nil)
	assert.T(t, result != nil)

	assert.T(t, title == "Edited")
}

// Should make a new Delete Request
func TestDelete(t *testing.T) {

	endpoint := fmt.Sprintf("products/%v", objIdToDelete)
	result, error := shop.Delete(endpoint)

	assert.T(t, error == nil, "should be null")
	assert.T(t, result != nil, "shouldnt be null")
}

func TestCreateTargetURL(t *testing.T) {

	endpoint := "products"
	result := shop.createTargetURL(endpoint)
	test := fmt.Sprintf("https://%s:%s@%s%s/%s.json", apiKey, pass, store, domain, endpoint)
	assert.Equal(t, result, test)
}



