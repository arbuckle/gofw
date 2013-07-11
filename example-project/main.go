//go manage.go

package main

import (
	"github.com/arbuckle/gofw"
)


func main() {
    settings := makeSettings()
	gofw.Run(settings)
}
