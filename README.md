go-shopify
==========

[![GoDoc](https://godoc.org/github.com/rapito/go-shopify/shopify?status.svg)](https://godoc.org/github.com/rapito/go-shopify/shopify)  [![baby-gopher](https://raw.github.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Simple API made with **go** to make **CRUD** request to your **Shopify Store**.

Installation
------------
```
go get github.com/rapito/go-shopify
```

How-to-use
----------


- Get Requests

```
    import "fmt"
    import "github.com/rapito/go-shopify/shopify"
    ...
    
    shop := shopify.New(storeDomain,apiKey,pass)
    result, _ := shop.Get("products")
    
    fmt.Println(string(result))
```

- Check out the *examples* folder for simple usage.
- Read some of the tests at *shopify_test.go* for complete CRUD examples.

Contribution
------------
 
 - You may fork this library and modify it as you please.
 - You can make a pull request and I will be happy to check it out and merge it.
 - If you find a bug, create an issue and I will do my best to fix it (someday). 

Original Work
-------------

While I was looking for something cool to do with this new language im learning 
(Go, obviously), I ran into [hammond-bones'](https://github.com/hammond-bones/) **go-shopify** 
library. Which inspired me to start creating this one. 

- Fork it at: [go-shopify](https://github.com/hammond-bones/go-shopify)

Buy me a Drink
-------------
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate?hosted_button_id=FFC6KNAX9SKZU)

Links
-----

While I was on my *go-trip* to create this api, I found some awesome libs which made 
my life easier.
Check them out, hopefully they'll do the same for you:
 
 - http://github.com/parnurzeal/gorequest
 - http://github.com/bmizerany/assert
 - http://github.com/avelino/awesome-go
 
 Other APIs
 ----------
 
 - http://github.com/rapito/go-spotify
 
