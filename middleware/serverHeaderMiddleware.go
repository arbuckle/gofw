// This is an example piece of middleware that:
// 1. appends a header to the incoming request, 
// 2. writes a line to the end of the outgoing response.


package middleware


import (
    "fmt"
    "github.com/arbuckle/gofw"
)

func ServerHeaderMiddleware() gofw.Middleware {
    return gofw.Middleware{ProcessRequest: serverHeaderRequestProcessor, ProcessResponse: serverHeaderResponseProcessor}
}


func serverHeaderRequestProcessor(r gofw.Request, w gofw.Response, a gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
    fmt.Println("RequestMiddleware")
    w.SetHeader("Server", "Smile :)")
    return r, w, a
}


func serverHeaderResponseProcessor(r gofw.Request, w gofw.Response, a gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
    fmt.Println("ResponseMiddleware")
    fmt.Fprintf(w.RawHttpResponseWriter, "\nThis line written by middleware")
    return r, w, a
}

