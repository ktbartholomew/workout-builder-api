package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ktbartholomew/workout-builder-api/config"
	"github.com/ktbartholomew/workout-builder-api/log"
	"github.com/ktbartholomew/workout-builder-api/router"
	"github.com/ktbartholomew/workout-builder-api/util"
)

// Request extends http.Request and adds a request ID field
type Request struct {
	*http.Request
	ID string
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, r *http.Request) {
		req := Request{Request: r, ID: util.NewUUID()}

		log.Log(log.Message{RequestID: req.ID, Event: fmt.Sprintf("%s %s", req.Request.Method, req.Request.URL)})
		routerResponse, err := router.Route(req.Request)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			io.WriteString(res, "{\"status\":\"error\"}")
		}

		res.Header().Set("content-type", "application/json")
		res.WriteHeader(routerResponse.StatusCode)
		io.WriteString(res, util.ToJSON(routerResponse.Body))
	})

	var port = ":" + config.Get("port")

	log.Log(log.Message{Event: "start-listening"})
	http.ListenAndServe(port, nil)
}
