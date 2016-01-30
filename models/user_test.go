package models

import (
    "testing"

    _ "github.com/mattn/go-sqlite3"

    . "github.com/franela/goblin"
    . "github.com/onsi/gomega"
)


func TestUser(t *testing.T) {
    g := Goblin(t)

    //special hook for gomega
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    g.Describe("User model", func() {
        var (
            user User
            skill Skill
        )

        g.Before(func() {
            OpenDB("sqlite3", "./gblogtest.db")
            db.LogMode(true)

            user = User{Nickname: "Azul", Email: "chinakyc@qq.com"}
            skill = Skill{Name: "Python", Value: 80}
        })

        g.BeforeEach(func() {
            db.CreateTable(&User{})
            db.CreateTable(&Skill{})
        })

        g.AfterEach(func() {
            db.DropTable(&User{})
            db.DropTable(&Skill{})
        })

        g.It("should handle password", func() {
            user.SetPassword("123123123")
            user.Save()

            vaild := user.Validate("123123123")
            Expect(vaild).To(BeTrue())

            vaild = user.Validate("123")
            Expect(vaild).To(BeFalse())

            user2 := User{}
            user2.Find("chiasf@qq.com")
            vaild = user2.Validate("123123123")
            Expect(vaild).To(BeFalse())
        })

        g.It("should add and get skills", func() {
            skills := []Skill{}

            user.AddSkill(skill)

            // user.Create()
            db.Create(&user)

            skills = append(skills, skill)

            test_skills := user.GetSkills()

            Expect(skills[0].Name).To(Equal(test_skills[0].Name))
        })

        g.It("should get user by email", func() {
            db.Create(&user)

            user_test := User{}

            user_test.Find("chinakyc@qq.com")

            Expect(user.Nickname).To(Equal(user_test.Nickname))
        })
    })
}
