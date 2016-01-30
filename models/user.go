package models

import (
    "time"
    "errors"

    "github.com/richardbowden/passwordHash"
)


const (
    ADMIN = iota
    EDITOR
    GUEST
)

type Skill struct {
    ID     int     `gorm:"primary_key" json:"-"`
    UserID int     `sql:"index" json:"-"`
    Name   string  `sql:"unique" json:"name"`
    Value  int     `json:"value"`
}

type User struct {
    ID            int       `gorm:"primary_key" json:"-"`
    Role          int       `sql:"default:2" json:"role"`
    Nickname      string    `json:"nickname"`
    Email         string    `sql:"not null;unique" json:"email"`
    Password_hash string    `json:"-"`
    About_me      string    `json:"aboutMe"`
    Skills        []Skill   `json:"skills"`
    Articles      []Article `json:"articles"`
    Last_seen     time.Time `json:"last_seen"`
}

func (self *User) Create() {
    db.Save(self)
}

func (self *User) Save() {
    db.Save(self)
}

func (self *User) Find(email string) {
    db.Where("email = ?", email).First(self)
}

func (self *User) SetPassword(s string) {
    self.Password_hash ,_ = passwordHash.HashWithDefaults(s)
}

func (self *User) Validate(s string) bool{
    if self.ID == 0 {
        return false
    }

    if self.Password_hash == "" {
        return false
    }
    return passwordHash.Validate(s, self.Password_hash)
}

func (self *User) AddSkill(skill Skill) {
    self.Skills = append(self.Skills, skill)
}

func (self *User) AddArticle(article *Article) error {
    // manual set `UserName` `UserID`
    article.UserName = self.Nickname
    article.UserEmail = self.Email

    // when you add article to user must create user first
    if self.ID == 0 {
        return errors.New("Must invoke `.Create` before")
    }

    article.UserID = self.ID

    return nil

    // This can cause multiple insertions
    // should invokle `article.Create` insert article
    // self.Articles = append(self.Articles, *article)
}

func (self *User) GetSkills() []Skill {
    skills := []Skill{}
    db.Where("user_id = ?", self.ID).Find(&skills)

    return skills
}


func InitUserDB() {
    db.CreateTable(&Skill{})
    db.CreateTable(&User{})

    python_skill := Skill{Name: "Python", Value: 5}
    router_skill := Skill{Name: "Network", Value: 1}
    linux_skill := Skill{Name: "Linux", Value: 3}
    js_skill := Skill{Name: "NodeJs", Value: 1}

    user := User{Nickname: "Azul", Email: "chinakyc@qq.com", About_me: "", Skills: []Skill{python_skill, js_skill, linux_skill, router_skill}, Last_seen: time.Now()}
    user.SetPassword("123123123")

    db.Create(&user)
}
