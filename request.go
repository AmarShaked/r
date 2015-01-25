package requests

// A Request represents an HTTP request received by a server.
type Request struct {
	Headers map[string]string
}
