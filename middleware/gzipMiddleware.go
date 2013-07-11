// gzipMiddleware.go
// SCRAPPING THIS EFFORT FOR NOW TO FOCUST ON gofw.Request and gofw.Response objects.
// Using the gofw middleware framework, gzips the rendered template response before sending it to the client.
// approach courtesy of Andrew Gerrand: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/eVnTcMwNVjM


package middleware

import (
    "compress/gzip"
    "fmt"
    "io"
    "net/http"
    "github.com/arbuckle/gofw"
)


type gzipResponseWriter struct {
    io.Writer
    http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
    //fmt.Println(w.Writer.Write(b))
    return w.Writer.Write(b)
}


func NewGZipMiddleware () gofw.Middleware {
    return gofw.Middleware{ProcessRequest: gZipRequestProcessor, ProcessResponse: gZipResponseProcessor}
}


func gZipRequestProcessor(request gofw.Request, response gofw.Response, args gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
//   response.SetHeader("Content-Encoding", "gzip")
//   response.SetHeader("Content-Type", "text/html; charset-utf-8")
   fmt.Fprintf(response.RawHttpResponseWriter, "gZip Request Middleware (does nothing)\n") 
   return request, response, args
}

func gZipResponseProcessor(request gofw.Request, response gofw.Response, args gofw.URLArgs) (gofw.Request, gofw.Response, gofw.URLArgs) {
    fmt.Fprintf(response.RawHttpResponseWriter, "gZip Response Middleware\n") 

    gz := gzip.NewWriter(response.RawHttpResponseWriter)
    newResponse := gofw.Response{RawHttpResponseWriter: gzipResponseWriter{Writer: gz, ResponseWriter: response.RawHttpResponseWriter}}
    fmt.Fprintf(newResponse.RawHttpResponseWriter, "Watever")
    defer gz.Close()
    return request, newResponse, args
}


