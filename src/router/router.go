package router

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

type RouteHandler = func(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type Router struct {
	request         *events.APIGatewayProxyRequest
	routes          map[string]RouteHandler
	notFoundHandler RouteHandler
}

func NewRouter(request *events.APIGatewayProxyRequest, notFoundHandler RouteHandler) Router {
	return Router{request, make(map[string]RouteHandler, 20), notFoundHandler}
}

func (r *Router) AddRoute(path string, handler RouteHandler) {
	r.routes[path] = handler
}

func (r Router) Handle(path string) (events.APIGatewayProxyResponse, error) {
	for route, closure := range r.routes {
		if route == path {
			return closure(r.request)
		}

		index := strings.Index(route, "{")
		if index == -1 {
			continue
		}

		baseUrl := route[0:index]
		if strings.Contains(path, baseUrl) {
			return closure(r.request)
		}
	}

	return r.notFoundHandler(r.request)
}
