package models

import (
    "time"
    "errors"
)


// use for write response json
type ArticlesResponse struct {
    TotalCount    int       `json:"total_count"`
    Articles      []Article `json:"articles"`
}

type Category struct {
    ID       int       `gorm:"primary_key" json:"id"`
    Name     string    `sql:"index" json:"name"`
    Articles []Article `json:"articles"`
}

type Article struct {
    ID           int        `gorm:"primary_key" json:"id"`
    CategoryID   int        `sql:"index" json:"-"`
    UserID       int        `sql:"index" json:"-"`
    // in my case basically impossible to modify the author so
    // cache author here better than use `Join Query`
    UserName     string     `json:"author_name"`
    UserEmail    string     `json:"author_email"`
    Title        string     `json:"title"`
    Content      string     `sql:"size:65535" json:"content"`
    Markdown     string     `sql:"size:65535" json:"markdown"`
    CreateTime   time.Time  `sql:"index; DEFAULT:current_timestamp" json:"create_time"`
}


func (self *Category) Create() {
    db.Create(self)
}

func (self *Category) FindByName(name string) {
    db.Where("name = ?", name).First(self)
}

// return all categorys
func (self *Category) All() []Category{
    categorys := []Category{}

    db.Find(&categorys)

    return categorys
}

// get all articles in this `category`
func (self *Category) AllArticles(offset int, limit int) ArticlesResponse{
    articles := []Article{}
    total_count := 0

    if self.ID == 0 {
        db.Model(Article{}).Count(&total_count).Order("create_time desc").Offset(offset).Limit(limit).Find(&articles)
        return ArticlesResponse{total_count, articles}
    }

    db.Model(Article{}).Where("category_id = ?", self.ID).Count(&total_count).Order("create_time desc").Offset(offset).Limit(limit).Find(&articles)

    return ArticlesResponse{total_count, articles}
}

func (self *Category) AddArticle(article *Article) error{
    // when you add article to `category` must create the `category` first
    if self.ID == 0 {
        return errors.New("Must invoke `.Create` before")
    }

    article.CategoryID = self.ID

    return nil

    // cause multiple insertions
    // should invokle `article.Create` insert article
    // self.Articles = append(self.Articles, *article)
}


func (self *Article) Create() {
    db.Create(self)
}

func (self *Article) Save() {
    db.Save(self)
}

func (self *Article) Delete() {
    db.Unscoped().Delete(self)
}

func (self *Article) Find(id int) {
    db.Where("id = ?", id).First(self)
}


func InitArticleDB() {
    db.CreateTable(&Article{})
    db.CreateTable(&Category{})
}
