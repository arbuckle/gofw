//run.go

package gofw

import (
	"net/http"
)

var routes = Routes{}
var middleware = []Middleware{}

type Template string //TEMP

// Template work:  Using pongo as the default template writer, since conceptually it's very similar
// to django/handlebars templates.  An important challenge will be to loosely couple templating engines with
// the rest of the framework.  Alternative template engines should be able to expose an interface that the
// framework will connect with in order to manage template rendering.

// TemplateDirs []string > represents a list of directories to search for templates, in order
// TemplateMiddleware > Templating will be handled entirely by a middleware plugin.  The last middleware
//  in the list will be a template middleware, which will extract the rawHttpResponseWriter from the
//  gofw.Response object and write a template to it.

// TemplateDirs []string > list of templates
// TemplateLoader > func that turns templates into cached template objects.
// TemplateMiddleware > func that reads template

// gofw will handle template loading using the directories enumerated in TemplateDirs []string
// gofw will load and cache all templates when the application is bootstrapped.
// gofw.Response will accept a template_name string which will determine the template to be written.
// Templates will be handled by middleware, so anyone can improve or swap out the template engine easily.

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
	Lang     string

	// Server configuration settings
	BasePath string
	Port     string
}

//TODO:  create an initialization function that a client app can use to export a settings object.
func NewSettings() *Settings {
	return &Settings{
		BasePath: "/",
		Port:     "8081"}
}

func registerRoutes(r Routes) {
	for _, url := range r.routes {
		routes.routes = append(routes.routes, url) //TODO:  consider removing dependency on this project-level global
	}
}

func registerMiddleware(m []Middleware) {
	for _, mWare := range m {
		middleware = append(middleware, mWare) //TODO:  consider removing dependency on this project-level global
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
	http.ListenAndServe(":"+settings.Port, nil)
}
