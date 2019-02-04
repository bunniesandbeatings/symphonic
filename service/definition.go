package service

import "syscall"

type Port uint16
type Param string

const (
	ParamEnvironment Param = "ENV"
	ParamPort        Param = "PORT"
)

type Command []string

type Edge struct {
	Port Port
}

type Route struct {
	// This simplest service discovery to use with a PaaS. Simply have all your servers run on the same url,
	// with path based routing. The PaaS does the load balancing. Downside, forced path based routing for all services
	// If you provide a path for one of the services, then symphony will create an NGINX server when running locally and
	// will map routes by path to your app. Symphony won't rewrite the path, it's assumed
	// your service can be (and is) configured to use the path in request. To this end, the env var ROUTE is also set with
	// this value
	Path string `json:path`

	// Todo support this, and service discovery in general
	// Another way to seperate services is by using different host name routes. Symphony can mix and match with host and path routing
	// and can even handle both for a service, so that some services are on one host under different names, etc.
	// Use a fqdn, and this will cause Symphony to create different ports, to simulate different hosts.
	// HOSTPORT will be set to the edge router PORT used for this emulated host (and it'll appear in the Host header of requests)
	// The downside to using RouteByHost is there's no conventional approach to service discovery.
	//Host string
}

type Service struct {
	Name string `json:name`

	Route Route `json:route`

	// A path relative to the common root where all commands operate from.
	// This is a convention based solution, so you do need your services to be accessible from a common
	// root, and to establish that structure with your colleagues.
	WorkingPath string `json:workingPath`

	// *** COMMANDS ***
	// Note that any array entry that contains "((PARAM))" where PARAM is a known
	// template variable, will be replaced

	// All commands will receive the same Environment as the one that calls Symphony. Use this to your advantage if
	// you need to pass through special configurations that don't map to the 4 default ENV environments.

	// Also the following env vars will be set (where applicable)
	// PORT
	// ENV [test, acceptance, production, development]
	// ROUTE for path based routing, is the path component used in routing.

	// Command to ensure dependencies are installed
	// try to only install deps for the given ((ENV)) environment. This allows users to only get the deps needed, for say
	// deploying to production.
	Deps Command `json:deps`

	// Command to start the app
	Start Command `json:start`
	// Signal to end the running service, defaults to SIGKILL
	TermSignal syscall.Signal `json:signal`

	// Command to clean up state. This is only called for the TEST env to allow the
	// service to clear any state between feature tests
	Clean Command `json:clean`

	// Command that runs all the tests for this service, to prove that it's working correctly
	Test Command `json:test`
}
