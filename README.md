# r [![GoDoc](http://godoc.org/github.com/AmarShaked/r?status.png)](http://godoc.org/github.com/AmarShaked/r)

Powerful package for quick and simple HTTP requests in Go language.

## Table of Contents
* [Why use r](#why-use-r)
* [How to install](#how-to-install)
* [Features](#features)
* [API](#api)
 * [Getting Started](#getting-started) 
 * [Quick requests](#quick-requests)
 * [Advance requests](#advance-requests)
    * [Basic advance request](#basic-advance-request)
    * [Builtin functions](#builtin-functions)
    * [Basic authentication](#basic-authentication)
    * [Request body](#body)
    * [Headers](#headers)
* [Contribute and things ToDo](#contribute-and-things-todo)


### Why use r
r takes all the power that language Go implements about http requests and makes all the work simpler and cleaner.
There is no need to use special and complex types as http.Header and url.Value.
Things like reading the response body, cookies, Headers and conversion from Json to custom type are now a simple tasks.

### Features
* Extremely simple to use.
* Make your code cleaner.
* Used simple types like string, map and bytes.
* Wrapping the common error handling for you.


### How to install
For install use the go get tool, just type this in your terminal:
~~~ bash
go get github.com/AmarShaked/r
~~~

###API

#### Getting Started
Let’s try to get a webpage. For this example, let’s get GitHub’s public timeline
~~~ go
res, err := r.Get('https://api.github.com/events')
// Some error handling
~~~
Now, we have a Response (not a regular http.Response) type called res. We can get all the information we need from this type.
For example, to print the response body we will have to just do this:
~~~ go
res.Text() // Get the response body as string
~~~

If we try to do this without r package, we will have to write something like this:
~~~ go
resp, err := http.Get("https://api.github.com/events")
// Some error handling
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// Error handling again
string(body)
~~~

#### Quick requests
We can send a quick requests of all HTTP requests types.
Usually we use quick requests in case we need to send a request with no special settings.

GET, HEAD, OPTIONS are the easiest:
~~~ go
r, _ := r.Get('http://httpbin.org/get')
r, _ := r.Options('http://httpbin.org/get')
r, _ := r.Head('http://httpbin.org/get')
~~~

In POST method, we get map of strings, and we use it like a PostForm.
~~~ go
data := map[string]string{"test": "shaked", "test2": "shaked2"}
r, _ := r.Post('http://httpbin.org/post', data)

// "form": {
//      "test": "shaked",
//      "test2": "shaked2"
//  }
~~~
PUT and DELETE receive a string for the body:
~~~ go
r, _ := r.Put('http://httpbin.org/put', "Test Body")
r, _ := r.Delete('http://httpbin.org/delete', "Test Body")
~~~

#### Advance requests
Advanced requests are designed for situations in which we want to add more things on the request.
For example: basic authentication, Headers, cookies, etc.

##### Basic advance request

~~~ go
req := &r.Request{Method: "GET",
                  Url: "http://httpbin.org/get"}
                  
res, err := req.Do()
~~~

##### Builtin functions
We do not have to provide a method to the request, but to use special functions that the type offer.
The Request type provides us functions for all types of methods.

~~~go
req := &r.Request{Url: "http://httpbin.org/get"}
                  
res, err := req.Get()
res, err := req.Post()
res, err := req.Put()
res, err := req.Delete()
res, err := req.Options()
res, err := req.Head()
~~~

##### Basic authentication
To add a Authentication header, just pass map of strings with your credentials.

~~~ go
req := &r.Request{Url: "http://httpbin.org/get",
                  Auth: []string{"username","password"}
                  
res, err := req.Post()
~~~

##### Body
In the body value we can put, string, bytes, io.Reader, and custom types that will parse as Json.
~~~ go
req := &r.Request{Url: "http://httpbin.org/get", Body: "string Body"}
req := &r.Request{Url: "http://httpbin.org/get", Body: []byte("bytes Body")}
req := &r.Request{Url: "http://httpbin.org/get", Body: strings.NewReader("string Body")}

type Animal struct {
    Name 
}

req := &r.Request{Url: "http://httpbin.org/get", Body: &Animal{Name: "dog"}}

res, err := req.Post()
~~~

#### Headers
The headers value gets a map of strings with all the header the you want to put in your request.
~~~ go
hdrs := map[string]string{"Header1": "Header 1 Value", "Header2": "Header 2 value"}
req := &r.Request{Url: "http://httpbin.org/get",
                  Headers: hdrs}
                  
res, err := req.Post()
~~~

### Contribute and things ToDo
We have many more things to add to the package. We work on it all the time and we need your help.
* Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug. 
* Fork the repository on GitHub to start making your changes to the master branch (or branch off of it).
* Send email to shakedamar@gmail.com for any questions you may have.
* Give as a star ;)