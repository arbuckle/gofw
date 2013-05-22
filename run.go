//run.go

package gofw

import (
	"flag"
	"fmt"
	"runtime/pprof"
	"os"
	"os/signal"
	"log"

	"net/http"
)


var routes = Routes{}

func RegisterRoutes(r *Routes) {
	for _, url := range(r.routes) {
		routes.routes = append(routes.routes, url)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func Run() {

	/* Begin Profiling */
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        
    }

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			fmt.Println("Exiting", sig)
			pprof.StopCPUProfile()
			os.Exit(0)
		}
	}()
	/* End Profiling */


	// Run.go needs to expose an interface for a management script in a FW project to interact with.
	// The FW project file will read the settings file, extract URLs and other information, and send this
	// information back to the server, here, which will ???
	
	// To start:
	// Read a list of URLs.
	// Bind handlers on the basis of said list.
	
	// Now:  
	// Update URLs to accept a regexp pattern instead of a set string
	// And update handler function to inject named RE variables into context...
	// Then:
	// Split this code out into a separate library
	// And:
	// Import the library from the management script within a sample application.

	http.HandleFunc("/", routes.baseHandler)
	http.ListenAndServe(":8081", nil)
}