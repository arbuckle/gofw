/*
	http.go
	Abstractions around the http.ResponseWriter and http.Request objects.
		
	Users of this library should not have to reference the golang documentation for 
	http.Request and http.ResponseWriter objects.  Instead, a basic interface for 
	setting context variables, validating forms, and rendering output should be created
	around these objects, with direct access to the raw http.ResponseWriter object
	reserved for power users.
	
	I think the idea here should be that this provides a thin wrapper around 
	the more powerful http.ResponseWriter / http.Request objects in order to 
	constrain the approaches people may take towards performing any particular action.
	
	
	
*/

package gofw

import (
	"fmt"
	"net/http"
)


type Response struct {
	RawHttpResponseWriter http.ResponseWriter
}

type Request struct {
	RawHttpRequest *http.Request
}

type URLArgs map[string] string

type Handler func(Request, Response, URLArgs)


// This won't compile.  Fix it.
type Middleware struct {
    ProcessRequest MiddlewarePreprocessing
    ProcessResponse MiddlewarePostprocessing  
}

type MiddlewarePreprocessing func(Request, Response) (Request, Response, URLArgs)
type MiddlewarePostprocessing func(Request, Response, URLArgs) (Request, Response, URLArgs)


// processRequest provides the core of the framework's functionality.  Here, we will
// apply middleware, append context arguments to the Request object, and possibly accomplish 
// some other necessary functions of response processing.
func processRequest (request *http.Request, response http.ResponseWriter, args URLArgs, handler Handler) {
    fmt.Println("Processing Request:\t" + request.URL.Path)
    fmt.Println(args)
	
    // get installed middleware
    // walk through list of installed middleware
    // apply middleware at each turn

	handler(Request{RawHttpRequest:request}, Response{RawHttpResponseWriter:response}, args)
}








