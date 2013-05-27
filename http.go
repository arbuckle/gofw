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

func (response *Response) SetHeader(key, value string) {
    fmt.Println("Setting Header", key, value)
    response.RawHttpResponseWriter.Header().Set(key, value)
}


type Request struct {
	RawHttpRequest *http.Request
}

type URLArgs map[string] string

type Handler func(Request, Response, URLArgs)

type PackageMiddleware []Middleware

type Middleware struct {
    ProcessRequest MiddlewareProcessor
    ProcessResponse MiddlewareProcessor
}

type MiddlewareProcessor func(Request, Response, URLArgs) (Request, Response, URLArgs)


// processRequest provides the core of the framework's functionality.  Here, we will
// apply middleware, append context arguments to the Request object, and possibly accomplish 
// some other necessary functions of response processing.
func processRequest (request *http.Request, response http.ResponseWriter, args URLArgs, handler Handler) {
    fmt.Println("Processing Request:\t" + request.URL.Path)
    fmt.Println(args)
	
    // get installed middleware
    // walk through list of installed middleware
    // apply middleware at each turn

    gofwRequest := Request{RawHttpRequest:request}
    gofwResponse := Response{RawHttpResponseWriter:response}

    // Applying Request Middleware
    for i := 0; i < len(middleware); i += 1 {
        gofwRequest, gofwResponse, args = middleware[i].ProcessRequest(gofwRequest, gofwResponse, args)
    }

    // Executing target handler
	handler(gofwRequest, gofwResponse, args)

    // Applying Response Middleware
    for i := len(middleware) - 1; i >= 0; i -= 1 {
        gofwRequest, gofwResponse, args = middleware[i].ProcessResponse(gofwRequest, gofwResponse, args)
        
    }
}








