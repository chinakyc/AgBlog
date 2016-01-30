package app

import (
    "fmt"

    "github.com/chinakyc/AgBlog/logging"
)

// middleware use my customized `request` `response`
type Middleware interface {
    ServeHTTP(ResponseWriter, *Request)
}

type ValidateUserFunc func (*Request) bool

type ValidateAuthorityFunc func (*Request, interface{}) bool


// like http.HandleFunc
type MiddlewareFunc func(ResponseWriter, *Request)

func (m MiddlewareFunc) ServeHTTP(w ResponseWriter, r *Request) {
    m(w, r)
}


// login Middleware
type requireUserHandle struct {
    handle Middleware
    loadUser ValidateUserFunc
}

func (h requireUserHandle) ServeHTTP(w ResponseWriter, r *Request) {
    authenticated := h.loadUser(r)

    if authenticated {
        h.handle.ServeHTTP(w, r)
    } else {
        w.WriteHeader(401)
        w.WriteJson(map[string]string{"error": "authenticated fail"})
    }

}

func RequireUserHandle (v ValidateUserFunc, m Middleware) requireUserHandle {
    return requireUserHandle{m, v}
}


// access Middleware
type accessHandle struct {
    handle Middleware
    validate ValidateAuthorityFunc
    role interface{}
}

func (h accessHandle) ServeHTTP(w ResponseWriter, r *Request) {
    authorized := h.validate(r, h.role)

    if authorized {
       h.handle.ServeHTTP(w, r)
    } else {
        w.WriteHeader(403)
        w.WriteJson(map[string]string{"error": "authorized fail"})
    }
}

func AccessHandle (v ValidateAuthorityFunc, role interface{}, m Middleware) accessHandle {
    return accessHandle{m, v, role}
}


type Cache interface {
    Set(string, []byte) error
    Get(string) ([]byte, error)
    Del(string) error
}

type cacheHandle struct {
    Cache  Cache
    handle Middleware
}

func (self *cacheHandle) ServeHTTP(w ResponseWriter, r *Request) {
    content , err := self.Cache.Get(r.URL.String())
    if err != nil {
        logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    }

    if content != nil {
        w.WriteHeader(200)
        w.Write(content)
        return
    }
    self.handle.ServeHTTP(w, r)
    content = w.getContentCopy()

    err = self.Cache.Set(r.URL.String(), content)

    if err != nil{
        logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    }
}

func CacheHanle(cache Cache, handle Middleware) *cacheHandle {
    return &cacheHandle{
        cache,
        handle,
    }
}


type deleteCacheHandle struct {
    Cache  Cache
    handle Middleware
}

func (self *deleteCacheHandle) ServeHTTP(w ResponseWriter, r *Request) {
    err := self.Cache.Del(r.URL.String())
    if err != nil{
        logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    }

    self.handle.ServeHTTP(w, r)
}

func DeleteCacheHanle(cache Cache, handle Middleware) *deleteCacheHandle {
    return &deleteCacheHandle{
        cache,
        handle,
    }
}
