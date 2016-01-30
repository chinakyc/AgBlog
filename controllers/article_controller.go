package controllers

import (
    "fmt"
    "strconv"

    "github.com/microcosm-cc/bluemonday"

    "github.com/chinakyc/AgBlog/app"
    "github.com/chinakyc/AgBlog/logging"
    "github.com/chinakyc/AgBlog/models"
)


// use for decode Json
type articleJsonBody struct {
    Title    string
    Content  string
    Markdown string
    Category string
}

func GetArticle(w app.ResponseWriter, r *app.Request) {
    article_id, err := strconv.Atoi(r.PathParams["article_id"])

    if err == nil {
        article := models.Article{}
        article.Find(article_id)

        if article.ID != 0 {
            w.WriteJson(article)
            return
        }
    }
    logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    w.WriteHeader(404)
    w.WriteJson(map[string]string{"error": "article no found"})
}

func generateArticleContent(r *app.Request) (string, string, models.Category, []byte){
    articleJson := articleJsonBody{}

    r.DecodeJsonPayload(&articleJson)

    // params
    markdown       := articleJson.Markdown
    title          := articleJson.Title
    category_name  := articleJson.Category
    unsafe         := articleJson.Content

    // find category
    category := models.Category{}
    if category_name != "" {
        category.FindByName(category_name)
        if category.ID == 0 {
            category.Name = category_name
            category.Create()
        }
    }

    // HTML sanitizer
    html := bluemonday.UGCPolicy().SanitizeBytes([]byte(unsafe))

    return title, markdown, category, html
}

func PostArticle(w app.ResponseWriter, r *app.Request) {
    user := r.Env["user"].(*models.User)

    title, markdown , category, html := generateArticleContent(r)

    if title == "" {
        w.WriteHeader(400)
        w.WriteJson(map[string]string{"error": "missing stuff"})
        return
    }

    // create article
    article := models.Article{Title: title, Content: string(html), Markdown: markdown}
    user.AddArticle(&article)
    category.AddArticle(&article)
    article.Create()

    w.WriteJson(article)
}

func DeleteArticle(w app.ResponseWriter, r *app.Request) {
    article_id, err := strconv.Atoi(r.PathParams["article_id"])

    if err == nil {
        article := models.Article{}
        article.Find(article_id)

        if article.ID != 0 {
            article.Delete()
            return
        }
    }

    logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    w.WriteHeader(404)
    w.WriteJson(map[string]string{"error": "article no found"})
}

func ModifyArticle(w app.ResponseWriter, r *app.Request) {
    article_id, err := strconv.Atoi(r.PathParams["article_id"])

    if err == nil {
        title, markdown, _, html := generateArticleContent(r)

        article := models.Article{}
        article.Find(article_id)

        if article.ID != 0 {
            article.Title    = title
            article.Markdown = markdown
            article.Content  = string(html)
            // category.AddArticle(&article)
            article.Save()
            w.WriteJson(article)
            return
        }
    }
    logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    w.WriteHeader(404)
    w.WriteJson(map[string]string{"error": "article no found"})
}
