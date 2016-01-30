package main

import (
    "os"
    "log"
    "time"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "github.com/gorilla/handlers"

    "github.com/chinakyc/AgBlog/app"
    "github.com/chinakyc/AgBlog/cache"
    "github.com/chinakyc/AgBlog/models"
    "github.com/chinakyc/AgBlog/logging"
    "github.com/chinakyc/AgBlog/controllers"
)


func main() {

    models.OpenDB("sqlite3", "./goblog.db")

    // My Customized httprouter
    webapp := app.NewApp()

    // `auth` group login and authentication methods
    auth := controllers.AuthMiddware()

    // simple cache
    redisPool := cache.NewCachePool("127.0.0.1:6379")
    cacheBackend := cache.NewCache("blog_", redisPool, time.Hour * 12)

    // static file
    webapp.ServeFiles("/static/*filepath", http.Dir("static/"))

    // template use for angularjs
    webapp.ServeFiles("/templates/*filepath", http.Dir("static/templates/"))

    // call origin `httprouter.GET`
    webapp.Router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        http.ServeFile(w, r, "static/templates/base.html")
    })

    // use email and password login return JWT token
    webapp.POST("/login", auth.LoginController)

    webapp.GET("/categorys", controllers.CategorysController)

    // return articles by category max limit 10
    webapp.Handler("GET", "/category/:category_id", controllers.ArticlesController(10))

    // return and cache owner data
    webapp.Handler(
        "GET",
        "/owner",
        app.CacheHanle(
            cacheBackend,
            app.MiddlewareFunc(controllers.OwnerController),
        ),
    )

    // return and cache article 
    webapp.Handler(
        "GET",
        "/article/:article_id",
        app.CacheHanle(
            cacheBackend,
            app.MiddlewareFunc(controllers.GetArticle)))

    // require user token and at least editor permissions
    // create article
    webapp.Handler(
        "POST",
        "/article",
        app.RequireUserHandle(
            auth.ValidateUser,
            app.AccessHandle(
                auth.ValidateAuthority,
                models.EDITOR,
                app.MiddlewareFunc(controllers.PostArticle)),
        ),
    )

    // require user token and at least admin permissions
    // delete article and cache
    webapp.Handler(
        "DELETE",
        "/article/:article_id",
        app.RequireUserHandle(
            auth.ValidateUser,
            app.AccessHandle(
                auth.ValidateAuthority,
                models.ADMIN,
                app.DeleteCacheHanle(
                    cacheBackend,
                    app.MiddlewareFunc(controllers.DeleteArticle)),
            ),
        ),
    )

    // require user token and at least admin permissions
    // update article and delete cache
    webapp.Handler(
        "PUT",
        "/article/:article_id",
        app.RequireUserHandle(
            auth.ValidateUser,
            app.AccessHandle(
                auth.ValidateAuthority,
                models.ADMIN,
                app.DeleteCacheHanle(
                    cacheBackend,
                app.MiddlewareFunc(controllers.ModifyArticle)),
            ),
        ),
    )

    logging.Logger.Info("Start listen...")
    log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, handlers.CompressHandler(webapp))))
}
