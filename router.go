package lostpapyrus

import (
    "net/http"
    "strings"
)

type Router struct {
    routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
    return &Router{routes: make(map[string]map[string]HandlerFunc)}
}

type HandlerFunc func(*Context, HandlerFunc)

func (r *Router) Handle(method string, path string, handler HandlerFunc) {
    if _, exists := r.routes[method]; !exists {
        r.routes[method] = make(map[string]HandlerFunc)
    }
    r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    method := req.Method
    path := req.URL.Path

    for route, handler := range r.routes[method] {
        if match, params := r.matchRoute(route, path); match {
            ctx := &Context{
                ResponseWriter: w,
                Request:        req,
                Params:         params,
            }
            handler(ctx, nil)
            return
        }
    }

    http.NotFound(w, req)
}

func (r *Router) matchRoute(route, path string) (bool, map[string]string) {
    routeParts := strings.Split(route, "/")
    pathParts := strings.Split(path, "/")

    if len(routeParts) != len(pathParts) {
        return false, nil
    }

    params := make(map[string]string)

    for i := range routeParts {
        if strings.HasPrefix(routeParts[i], ":") {
            params[routeParts[i][1:]] = pathParts[i]
        } else if routeParts[i] != pathParts[i] {
            return false, nil
        }
    }

    return true, params
}
