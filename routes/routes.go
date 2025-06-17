package routes

// HandlerFunc defines a type for route handler functions
type HandlerFunc func() (string, string) // returns (contentType, body)

// RouteTable maps method + path to a handler
var RouteTable = map[string]HandlerFunc{
	"GET /api/hello": HelloHandler,
	"GET /":          HomeHandler,
}

// MatchRoute finds a matching route handler
func MatchRoute(method string, path string) (HandlerFunc, bool) {
	key := method + " " + path
	handler, exists := RouteTable[key]
	return handler, exists
}

// HelloHandler is a sample API handler
func HelloHandler() (string, string) {
	body := `{"message": "Hello from BYOS-GO API"}`
	return "application/json", body
}

// HomeHandler serves plain text for "/"
func HomeHandler() (string, string) {
	return "text/plain", "Welcome to BYOS-GO!"
}
