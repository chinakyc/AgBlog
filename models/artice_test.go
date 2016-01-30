package models

import (
    "testing"
    "time"

    _ "github.com/mattn/go-sqlite3"

    . "github.com/franela/goblin"
    . "github.com/onsi/gomega"
)


func TestArticle(t *testing.T) {
    g := Goblin(t)

    //special hook for gomega
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    g.Describe("Article and Category model", func() {
        var (
            category Category
            article1 Article
            article2 Article
            article3 Article
            user User
        )

        g.Before(func() {
            OpenDB("sqlite3", "./gblogtest.db")
            db.LogMode(true)

            user = User{Nickname:"Azul", Email:"chinakyc@qq.com"}
            category = Category{Name: "test"}
            article1 = Article{Title: "test1", Content: "test1", CreateTime: (time.Now().Add(1 * time.Hour).UTC())}
            article2 = Article{Title: "test2", Content: "test2", CreateTime: (time.Now().Add(3 * time.Hour).UTC())}
            article3 = Article{Title: "test3", Content: "test3", CreateTime: (time.Now().Add(3 * time.Hour).UTC())}
        })

        g.BeforeEach(func() {
            db.CreateTable(&User{})
            db.CreateTable(&Category{})
            db.CreateTable(&Article{})

            user.Create()
            _ = user.AddArticle(&article1)
            _ = user.AddArticle(&article2)
            _ = user.AddArticle(&article3)

            category.Create()
            _ = category.AddArticle(&article1)
            _ = category.AddArticle(&article2)
            _ = category.AddArticle(&article3)

            article1.Create()
            article2.Create()
            article3.Create()

        })

        g.AfterEach(func() {
            db.DropTable(&User{})
            db.DropTable(&Category{})
            db.DropTable(&Article{})
        })

        g.It("should get article by id", func() {
            article_test := Article{}

            article_test.Find(article1.ID)

            Expect(article_test).To(Equal(article1))
        })

        g.It("should get articles by Category", func() {
            articles := category.AllArticles(0,3)
            articles_test := []Article{article3, article2, article1}

            Expect(articles.TotalCount).To(Equal(len(articles_test)))
            Expect(articles.Articles).To(Equal(articles_test))

            articles = category.AllArticles(1,2)
            articles_test = []Article{article2, article1}

            Expect(articles.Articles).To(Equal(articles_test))
            Expect(articles.TotalCount).To(Equal(3))
        })

        g.It("should get all articles", func() {
            categorys := category.All()

            Expect(categorys).To(Equal([]Category{category}))
        })
    })
}
