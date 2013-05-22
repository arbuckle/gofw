
/*
	Urls.go is the base URL handler for gofw.
	TODO:  Read up on GODOC and write documentation for the individual
	methods that comprise the interface exposed by this package.
	
	Briefly, a "Routes" type exposes an interface via its Register method which
	enables for the creation of simple bindings between URL regexp patterns and 
	event handlers.  
	
	As an abstraction between the project-level Routes.go file, and
	the framework-level handling of controllers, this module will
	need to create a harness that bridges the gap between the two concepts,
	so that "net/http" does not need to be imported within the user's project.
	
	As an aside, it should be easy for the user to swap out the urls.go
	frontend with one of their own, permitting multiple formats of urlconf.
*/

package gofw

import (
	"net/http"
	"regexp"
)


// Basic struct for defining a URL Route
type url struct {
	Pattern *regexp.Regexp
	Handler func(Response, Request)
	Name string
}
func (u *url) String() string {
	return u.Name
}


// Struct for representing a collection of URL Routes 
type Routes struct {
	routes []*url
}

// Register method on Routes provides means of adding URLs to a route.
func (routes *Routes) Register(pattern string, handler func(Response, Request), name string) {
	routes.routes = append(routes.routes, &url{regexp.MustCompile(pattern), handler, name})
}

// getKeywordArguments is called by the base handler to extract keyword arguments from a URL pattern.
func getKeywordArguments(path string, pattern *regexp.Regexp) (args URLArgs) {
	args = make(URLArgs)
	matches := pattern.FindAllStringSubmatch(path, -1)
	for i, key := range(pattern.SubexpNames()) {
		if i == 0 || key == "" {
			continue
		}
		args[key] = matches[0][i]
	}
	return args
}

// baseHandler is called by run.go to listen to all incoming HTTP traffic and route responses 
// according to the user-provided URLConf.
// This method is designed to be easily swapped out in favor of alternative URLConfs, 
// with the only requirement being that the replacement call processRequest with the necessary arguments.
func (routes *Routes) baseHandler (w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	for _, url := range(routes.routes) {
		if url.Pattern.MatchString(path) {
			args := getKeywordArguments(path, url.Pattern)
			processRequest(r, w, args, url.Handler)
			return
		}
	}
	http.NotFound(w, r)
	return 
}

