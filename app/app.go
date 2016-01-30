package app

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)

type Handle func (w ResponseWriter, r *Request)


type App struct {
    httprouter.Router
}

// warp handle use my coustom Request and Response 
func (self *App) GET(path string, handle Handle) {
    self.Router.Handle("GET", path, self.WarpHandle(handle))
}

func (self *App) HEAD(path string, handle Handle) {
    self.Router.Handle("HEAD", path, self.WarpHandle(handle))
}

func (self *App) OPTIONS(path string, handle Handle) {
    self.Router.Handle("OPTIONS", path, self.WarpHandle(handle))
}

func (self *App) POST(path string, handle Handle) {
    self.Router.Handle("POST", path, self.WarpHandle(handle))
}

func (self *App) PUT(path string, handle Handle) {
    self.Router.Handle("PUT", path, self.WarpHandle(handle))
}

func (self *App) PATCH(path string, handle Handle) {
    self.Router.Handle("PATCH", path, self.WarpHandle(handle))
}

func (self *App) DELETE(path string, handle Handle) {
    self.Router.Handle("DELETE", path, self.WarpHandle(handle))
}

func (self *App) Handle(method, path string, handle Handle) {
    self.Router.Handle(method, path, self.WarpHandle(handle))
}

func (self *App) Handler(method, path string, middleware Middleware) {
    self.Router.Handle(method, path, self.WarpMiddleware(middleware))
}

func (self *App) WarpHandle(h Handle) httprouter.Handle {
    return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        h(AdaptResponse(w), AdaptRequest(r, ps))
    }
}

func (self *App) WarpMiddleware(h Middleware) httprouter.Handle {
    return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        h.ServeHTTP(AdaptResponse(w), AdaptRequest(r, ps))
    }
}

func NewApp() *App {
    return &App{
        httprouter.Router{
            RedirectTrailingSlash:  true,
            RedirectFixedPath:      true,
            HandleMethodNotAllowed: true,
        },
    }
}
