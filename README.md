# requests
HTTP library for Go language
## Getting Started

### Make a Request
Making a request with Requests is very simple.
Begin by installing the Requests package.
~~~ go
go get github.com/AmarShaked/requests
import "github.com/AmarShaked/requests"
~~~
Now, let’s try to get a webpage. For this example, let’s get GitHub’s public timeline
~~~ go
r, _ := requests.Get('https://api.github.com/events')
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