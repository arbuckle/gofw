//run.go

package gofw

import (
	"net/http"
)


var routes = Routes{}
var middleware = []*Middleware{}


type Template string //TEMP
type DatabaseConfig struct {
    Engine, Name, User, Password, Host, Port string
}

type Settings struct {
    // Various boolean settings
    Debug, I18N, L10N bool

    // Administrator information
    Admins [][2]string

    // IP Address whitelist for debug output
    InternalIP []string

    // Database info
    Databases map[string]DatabaseConfig

    // Middleware handling functions
    MiddlewareHandlers []Middleware

    // Template processors (should there be more than one??)
    TemplateProcessors []Template

    // Project handler func for Routes
    URLHandler func() Routes

    // Hostnames from which requests will be permitted
    AllowedHosts []string

    // Public content settings - directories to expose to public, hostname from which to serve public requests.
    PublicDirs []string
    PublicHost string

    // Time Zone
    TimeZone string
    Lang string

    // Server configuration settings
    BasePath string
    Port string
}

//TODO:  create an initialization function that a client app can use to export a settings object.
func NewSettings() *Settings {
    return &Settings{
        BasePath: "/", 
        Port: "8081"}
}


func registerRoutes(r Routes) {
	for _, url := range(r.routes) {
		routes.routes = append(routes.routes, url) //TODO:  consider removing dependency on this project-level global
	}
}

func registerMiddleware(m []Middleware) {
	for _, mWare := range(m) {
		middleware = append(middleware, &mWare) //TODO:  consider removing dependency on this project-level global
	}
}

func Run(settings *Settings) {
	// Run.go needs to expose an interface for a management script in a FW project to interact with.
	// The FW project file will read the settings file, extract URLs and other information, register
    // functionality within the application, and finally process the request.
	



    // Registering URLs declared in project base handler.
    registerRoutes(settings.URLHandler())
    registerMiddleware(settings.MiddlewareHandlers)

    

    // Starting webserver
	http.HandleFunc(settings.BasePath, routes.baseHandler)
	http.ListenAndServe(":" + settings.Port, nil)
}



