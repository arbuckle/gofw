// This is an example piece of middleware that:
// 1. appends a header to the incoming request,
// 2. writes a line to the end of the outgoing response.

package main

import (
	"fmt"
	"github.com/arbuckle/gofw"
	"time"
)

func timerMiddleware() gofw.Middleware {
	return gofw.Middleware{ProcessRequest: timerRequestProcessor, ProcessResponse: timerResponseProcessor}
}

func timerRequestProcessor(r gofw.Request, w gofw.Response, a gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
	fmt.Println("RequestMiddleware-Timer")
	w.Set("timerStart", time.Now())
	return r, w, a
}

func timerResponseProcessor(r gofw.Request, w gofw.Response, a gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
	fmt.Println("ResponseMiddleware-Timer")
	time.Sleep(56 * time.Millisecond)
	fmt.Fprintf(w.RawHttpResponseWriter, "\nPage rendered in: %s", time.Now().Sub(w.Get("timerStart").(time.Time)))

	return r, w, a
}
