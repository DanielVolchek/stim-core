package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/danielvolchek/stim-utils/pkg/router"
)

var ItemRouter router.Route = router.Route{
	Route:        "/items",
	FinalHandler: ItemRouteHandler,
}

var ItemRouteHandler http.HandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", ":8081/db/items", nil)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error creating request:"))
		return
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error sending request:"))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Error reading response body:"))
		return
	}

	// Print the response
	w.Write(body)

})
