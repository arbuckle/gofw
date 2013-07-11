// ProjectRegistry.go

package main

import (
	"github.com/arbuckle/gofw"
	middleware "github.com/arbuckle/gofw/middleware"
)

func makeSettings() *gofw.Settings {
	settings := gofw.NewSettings()

	settings.Debug = true
	settings.BasePath = "/"
	settings.Port = "8081"
	settings.URLHandler = makeRoutes
	settings.MiddlewareHandlers = []gofw.Middleware{
		timerMiddleware(),
		middleware.ServerHeaderMiddleware()}
	return settings
}
