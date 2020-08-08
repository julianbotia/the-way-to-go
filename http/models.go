package http

// Method Alias
type Method string

// Supported methods
const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Patch  Method = "PATCH"
	Delete Method = "DELETE"
)

// Request Wrapper
type Request struct {
	Method      Method
	BaseURL     string
	Body        []byte
	QueryParams map[string]string
}

// Response Wrapper
type Response struct {
	Status int
	Body   string
}
