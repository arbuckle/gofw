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
	W http.ResponseWriter
}

type Request struct {
	R *http.Request
}

type URLArgs map[string] string

type Handler func(Request, Response, URLArgs)

func processRequest (request *http.Request, response http.ResponseWriter, args URLArgs, handler func(Response, Request)) {
	fmt.Println("Processing Request:\t" + request.URL.Path)
	fmt.Println(args)
	handler(Response{W:response}, Request{R:request})
}

