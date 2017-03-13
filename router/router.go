// Package router calls certain handlers depending on the HTTP request
//
// I realize that things like gorilla/mux exist. I'm doing this to learn Go, and
// writing my own router means I have to munge around a lot of different data
// types and generally solve simple problems. This is a learning exercise, and
// not a perfect router.
package router

import (
	"net/http"
	"strings"
)

// RouteDefinition a URL, Method, and corresponding handler function key
type RouteDefinition struct {
	URL     string
	Method  string
	Handler string
}

// RouteResponse status code, headers, and a body. Body can be freeform, will be
// marshalled to JSON and sent to browser
type RouteResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       interface{}
}

// NewResponse is a factory for routerResponse's above
func NewResponse() RouteResponse {
	response := RouteResponse{
		StatusCode: http.StatusOK,
		Headers:    make(map[string]string),
		Body:       struct{}{}}

	return response
}

// RouteHandler is a functionn that returns a RouteResponse
type RouteHandler func(req *http.Request) (RouteResponse, error)

var routes []RouteDefinition

var handlers map[string]RouteHandler

func init() {
	routes = make([]RouteDefinition, 0)
	handlers = make(map[string]RouteHandler)

	routes = append(routes, RouteDefinition{URL: "/workouts", Method: "GET", Handler: "getWorkouts"})
	routes = append(routes, RouteDefinition{URL: "/workouts", Method: "POST", Handler: "createWorkout"})
	routes = append(routes, RouteDefinition{URL: "/workouts/:workoutID", Method: "GET", Handler: "getWorkout"})

	handlers["404"] = func(req *http.Request) (RouteResponse, error) {
		response := NewResponse()
		response.StatusCode = 404

		return response, nil
	}

	handlers["getWorkouts"] = getWorkouts
	handlers["getWorkout"] = getWorkout
	handlers["createWorkout"] = createWorkout
}

func findRoute(request *http.Request) RouteDefinition {
	urlParts := tokenize(request.URL.Path)

	// If no other route is matched below, this 404 handler will be returned
	matchingRoute := RouteDefinition{Handler: "404"}

	// Loop over routes. Only one can match
	for i := 0; i < len(routes); i++ {
		routeParts := tokenize(routes[i].URL)
		matchingParts := 0

		// mismatched methods is a guaranteed miss
		if request.Method != routes[i].Method {
			continue
		}

		// different number of URL segments is a guaranteed miss
		if len(routeParts) != len(urlParts) {
			continue
		}

		// Loop over each segment and make sure each matches
		for j := 0; j < len(routeParts); j++ {
			if strings.HasPrefix(routeParts[j], ":") {
				matchingParts++
			}

			if routeParts[j] == urlParts[j] {
				matchingParts++
			}
		}

		if matchingParts == len(routeParts) {
			matchingRoute = routes[i]
			break
		}
	}

	return matchingRoute
}

func tokenize(u string) []string {
	tokens := strings.Split(u, "/")

	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i] == "" {
			tokens = append(tokens[:i], tokens[i+1:]...)
		}
	}

	return tokens
}

// Route call the routing function?
func Route(req *http.Request) (RouteResponse, error) {
	response, error := handlers[findRoute(req).Handler](req)

	return response, error
}
