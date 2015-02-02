# r
HTTP library for Go language

## Getting Started

### Make a Request
Making a request with Requests is very simple.
Begin by installing the Requests package.
~~~ go
go get github.com/AmarShaked/r
import "github.com/AmarShaked/r"
~~~
Now, let’s try to get a webpage. For this example, let’s get GitHub’s public timeline
~~~ go
r, _ := r.Get('https://api.github.com/events')
~~~
Now, we have a Response (not a regular http.Response) object called r. We can get all the information we need from this object.
For example:
~~~ go
r.StatusCode // 200
r.Text() // Get the response body as string
r.Json(&jsonObject) // Parse the json data and stores the result in the value pointed to by jsonObject.
r.Headers('content-type') // 'application/json'
r.Cookies('key') // Return a cookie value as string
~~~
### Quick requests
You can send a quick requests of all HTTP requests types.

GET, HEAD, OPTIONS are the easiest:
~~~ go
r, _ := r.Get('http://httpbin.org/get')
r, _ := r.Options('http://httpbin.org/get')
r, _ := r.Head('http://httpbin.org/get')
~~~

POST, PUT, DELETE Usually contain form data:
~~~ go
data := map[string]string{"test": "shaked", "test2": "shaked2"} // "form": {"test": "shaked","test2": "shaked2"}
r, _ := r.Post('http://httpbin.org/post', data)
r, _ := r.Put('http://httpbin.org/put', data)
r, _ := r.Delete('http://httpbin.org/delete', data)
~~~