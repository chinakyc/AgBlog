package controllers


import (
    "github.com/chinakyc/AgBlog/app"
    "github.com/chinakyc/AgBlog/models"
)


func OwnerController(w app.ResponseWriter, r *app.Request) {
    user := models.User{}
    user.Find("chinakyc@qq.com")
    user.Skills = user.GetSkills()
    w.WriteJson(user)
}
