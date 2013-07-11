//urls.go

package main

import (
	"github.com/arbuckle/gofw"
)

func makeRoutes() gofw.Routes {
	routes := gofw.Routes{}
	routes.Register("^/$", Sreegs, "root")
	routes.Register("^/test/(?P<id>[0-9]+)/(?P<oth_id>[0-9]+)/$", Sreegs, "testing")
	routes.Register("^/test1/$", Sreegs, "testing1")
	routes.Register("^/test2/$", Sreegs, "testing2")
	
	return routes
}
