//run.go

package gofw

import (
	"net/http"
)


var routes = Routes{}

func RegisterRoutes(r *Routes) {
	for _, url := range(r.routes) {
		routes.routes = append(routes.routes, url)
	}
}

func Run() {
	// Run.go needs to expose an interface for a management script in a FW project to interact with.
	// The FW project file will read the settings file, extract URLs and other information, register
    // functionality within the application, and finally process the request.
	

	http.HandleFunc("/", routes.baseHandler)
	http.ListenAndServe(":8081", nil)
}
