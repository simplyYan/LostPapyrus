package lostpapyrus

import (
    "net/http"
)

type App struct {
    Router      *Router
    middlewares []HandlerFunc
}

func New() *App {
    return &App{Router: NewRouter()}
}

func (app *App) Use(middleware HandlerFunc) {
    app.middlewares = append(app.middlewares, middleware)
}

func (app *App) Get(path string, handler HandlerFunc) {
    app.Router.Handle("GET", path, handler)
}

func (app *App) Post(path string, handler HandlerFunc) {
    app.Router.Handle("POST", path, handler)
}

func (app *App) Listen(addr string) {
    http.ListenAndServe(addr, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctx := &Context{
        ResponseWriter: w,
        Request:        req,
        Params:         make(map[string]string),
    }

    var next HandlerFunc
    next = func(c *Context, _ HandlerFunc) {
        app.Router.ServeHTTP(w, req)
    }

    for i := len(app.middlewares) - 1; i >= 0; i-- {
        current := app.middlewares[i]
        next = createNext(current, next)
    }

    next(ctx, nil)
}

func createNext(current, next HandlerFunc) HandlerFunc {
    return func(ctx *Context, _ HandlerFunc) {
        current(ctx, next)
    }
}
