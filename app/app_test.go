package app

import (
    "testing"
    "net/http"
    "net/http/httptest"

    // "github.com/julienschmidt/httprouter"
    . "github.com/franela/goblin"
    . "github.com/onsi/gomega"
)


type testhandleEnv struct {
    handle Middleware
}

func (h testhandleEnv) ServeHTTP (w ResponseWriter, r *Request) {
    r.Env["name"] = "kyc"
    h.handle.ServeHTTP(w, r)
}

func testHandleEnv(m MiddlewareFunc) testhandleEnv{
    return testhandleEnv{m,}
}


func Test(t *testing.T) {
    g := Goblin(t)

    // special hook for gomega
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    g.Describe("httprouter normal test", func() {
        var app *App

        g.Before(func() {
            app = NewApp()
        })

        g.It("should like a normal httprouter", func() {

            routed := false

            app.Handle("GET", "/user/:name", func(w ResponseWriter, r *Request) {
                routed = true
                Expect(r.PathParams).To(Equal(map[string]string{"name": "gopher"}))
            })

            w := httptest.NewRecorder()
            req, _ := http.NewRequest("GET", "/user/gopher", nil)

            app.ServeHTTP(w, req)

            Expect(routed).To(BeTrue())

        })
    })

    g.Describe("Costom Rquest", func() {
        var (
            app *App
            req *http.Request
            w http.ResponseWriter
        )

        g.BeforeEach(func() {
            app    = NewApp()
            req, _ = http.NewRequest("GET", "/user/kyc", nil)
            w      = httptest.NewRecorder()
        })

        g.It("should carry path params", func() {
            app.GET("/user/:name", func (w ResponseWriter, r *Request) {
                Expect(r.PathParams["name"]).To(Equal("kyc"))
            })

            app.ServeHTTP(w, req)
        })

        g.It("should carry Env", func() {
            app.Handler("GET", "/user/:name", testHandleEnv(func (w ResponseWriter, r *Request) {
                Expect(r.Env["name"].(string)).To(Equal("kyc"))
            }))

            app.ServeHTTP(w, req)
        })
    })
}
