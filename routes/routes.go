package routes

import (
	"byos-go/utils"
	"fmt"
	"strings"
)

// HandlerFunc defines a type for route handler functions
type HandlerFunc func(req *utils.HTTPRequest, params map[string]string) (string, string) // returns (contentType, body)

type DynamicRoute struct {
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

var DynamicRoutes = []DynamicRoute{
	{
		Method:  "GET",
		Pattern: "/api/user/:id",
		HandlerFunc: func(req *utils.HTTPRequest, params map[string]string) (string, string) {
			id := params["id"]
			return "application/json", fmt.Sprintf(`{"user_id": "%s"}`, id)
		},
	},
	{
		Method:  "POST",
		Pattern: "/api/echo",
		HandlerFunc: func(req *utils.HTTPRequest, params map[string]string) (string, string) {
			return "application/json", fmt.Sprintf(`{"echo": %q}`, req.Body)
		},
	},
}

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
func HelloHandler(req *utils.HTTPRequest, _ map[string]string) (string, string) {
	body := `{"message": "Hello from BYOS-GO API"}`
	return "application/json", body
}

// HomeHandler serves plain text for "/"
func HomeHandler(req *utils.HTTPRequest, _ map[string]string) (string, string) {
	return "text/plain", "Welcome to BYOS-GO!"
}

func MatchDynamicRoute(method, path string) (HandlerFunc, map[string]string, bool) {
	for _, route := range DynamicRoutes {
		if route.Method != method {
			continue
		}

		patternParts := strings.Split(route.Pattern, "/")
		pathParts := strings.Split(path, "/")

		if len(patternParts) != len(pathParts) {
			continue
		}

		params := make(map[string]string)
		matched := true
		for i := range patternParts {
			if strings.HasPrefix(patternParts[i], ":") {
				key := strings.TrimPrefix(patternParts[i], ":")
				params[key] = pathParts[i]
			} else if patternParts[i] != pathParts[i] {
				matched = false
				break
			}
		}

		if matched {
			return route.HandlerFunc, params, true
		}
	}

	return nil, nil, false
}
