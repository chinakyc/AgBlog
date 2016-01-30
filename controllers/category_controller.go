package controllers

import (
    "fmt"
    "strconv"

    "github.com/chinakyc/AgBlog/app"
    "github.com/chinakyc/AgBlog/logging"
    "github.com/chinakyc/AgBlog/models"
)


type articlesController struct {
    PagePerNumberMAX int
}

func (c articlesController) ServeHTTP(w app.ResponseWriter, r *app.Request) {
    var err error

    // default args
    offset      := 0
    limit       := c.PagePerNumberMAX
    category_id := 0
    err         = nil

    str_category_id := r.PathParams["category_id"]
    if str_category_id != "" {
        category_id, err = strconv.Atoi(str_category_id)
        if err != nil {
            logging.Logger.Error(fmt.Sprintf("Error: %s", err))
            category_id = 0
        }
    }

    uri_params := r.URL.Query()

    if limit_str := uri_params["limit"]; limit_str != nil {
        limit, err = strconv.Atoi(limit_str[0])
        if err != nil {
            logging.Logger.Error(fmt.Sprintf("Error: %s", err))
            limit = c.PagePerNumberMAX
        }
    }

    if offset_str := uri_params["offset"]; offset_str != nil {
        offset,err = strconv.Atoi(offset_str[0])
        if err != nil {
            logging.Logger.Error(fmt.Sprintf("Error: %s", err))
            offset = 0
        }
    }

    category := models.Category{ID:category_id}
    articles := category.AllArticles(offset, limit)

    w.WriteJson(articles)
}


func ArticlesController (per_page int) articlesController{
    return articlesController{PagePerNumberMAX: per_page}
}

func CategorysController(w app.ResponseWriter, r *app.Request) {
    category  := models.Category{}
    categorys := category.All()

    w.WriteJson(categorys)
}

