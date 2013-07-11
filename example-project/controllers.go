// controllers.go

package main

import (
	"fmt"
	"github.com/arbuckle/gofw"
)

// Sreegs test handler
func Sreegs(r gofw.Request, w gofw.Response, args gofw.URLArgs) {
	fmt.Fprintf(w.RawHttpResponseWriter, "Hello World\n")
	return 
}

