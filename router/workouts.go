package router

import "net/http"
import "github.com/ktbartholomew/workout-builder-api/workouts"

func getWorkouts(req *http.Request) (RouteResponse, error) {
	response := NewResponse()

	response.Body = workouts.GetCollection()

	return response, nil
}

func getWorkout(req *http.Request) (RouteResponse, error) {
	response := NewResponse()

	response.Body = workouts.Get()

	return response, nil
}

func createWorkout(req *http.Request) (RouteResponse, error) {
	response := NewResponse()
	response.StatusCode = http.StatusAccepted

	return response, nil
}
