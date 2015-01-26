# requests
HTTP library for Go language

## Getting Started

Install the requests package 
~~~
go get github.com/AmarShaked/requests
~~~

~~~ go
package main

import "github.com/AmarShaked/requests"

func main() {

	url := "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=israel"

	res, _ := requests.Get(url)

	// Print the body as string
	fmt.Println(res.Text())

	// Get the header
	fmt.Println(res.Headers("content-type"))

	// Parse as json
	var json JSONTEST
	res.Json(&json)
}

~~~
