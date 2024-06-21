package lostpapyrus

import (
    "encoding/json"
    "net/http"
)

type Context struct {
    ResponseWriter http.ResponseWriter
    Request        *http.Request
    Params         map[string]string
}

func (ctx *Context) Send(body string) *Context {
    ctx.ResponseWriter.Write([]byte(body))
    return ctx
}

func (ctx *Context) Status(statusCode int) *Context {
    ctx.ResponseWriter.WriteHeader(statusCode)
    return ctx
}

func (ctx *Context) JSON(data interface{}) *Context {
    ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
    json.NewEncoder(ctx.ResponseWriter).Encode(data)
    return ctx
}

func (ctx *Context) BindJSON(obj interface{}) error {
    decoder := json.NewDecoder(ctx.Request.Body)
    return decoder.Decode(obj)
}
