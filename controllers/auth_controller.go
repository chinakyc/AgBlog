package controllers

import (
    "fmt"
    "time"
    "strings"

    "github.com/dgrijalva/jwt-go"

    "github.com/chinakyc/AgBlog/app"
    "github.com/chinakyc/AgBlog/logging"
    "github.com/chinakyc/AgBlog/models"
)


type jsonLoginDate struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type responseUserData struct {
    User    string `json:"user"`
    Token   string `json:"token"`
    Role    int    `json:"role"`
}

type authMiddware struct {
    signingKey    []byte
    keyFunc jwt.Keyfunc
}

func AuthMiddware() *authMiddware{
    return &authMiddware{
        signingKey: []byte("MakeYourPrivateKey"),
        keyFunc: func(t *jwt.Token) (interface{}, error) {
            return []byte("MakeYourPrivateKey"), nil
        },
    }
}

func (self *authMiddware) LoginController(w app.ResponseWriter, r *app.Request) {
    var tokenString string

    data := jsonLoginDate{}

    // Decode Json from request
    err := r.DecodeJsonPayload(&data)

    if err != nil {
        logging.Logger.Error(fmt.Sprintf("Error: %s", err))
        w.WriteHeader(500)
        w.WriteJson(map[string]string{"error": fmt.Sprintf("Error: %s", err)})
        return
    }

    // extract 
    email := data.Email
    password := data.Password

    // use email get user
    user := models.User{}
    user.Find(email)

    // validate password generate jwt tokenString
    // user jwt we can ignore CRSF
    if user.Validate(password) {
        user.Last_seen = time.Now().UTC()
        user.Save()
        token := jwt.New(jwt.SigningMethodHS256)
        token.Claims["email"] = user.Email
        token.Claims["role"] = user.Role
        token.Claims["exp"] = time.Now().Add(time.Hour * 6).UTC().Unix()
        tokenString, err = token.SignedString(self.signingKey)
        if err != nil {
            logging.Logger.Error(fmt.Sprintf("Error: %s", err))
            w.WriteHeader(500)
            w.WriteJson(map[string]string{"error": fmt.Sprintf("Error: %s", err)})
        }
        w.WriteJson(responseUserData{user.Nickname, tokenString, user.Role})

    } else {
        w.WriteHeader(400)
        w.WriteJson(map[string]string{"error": "email or password incorrect"})
    }
}

func (self *authMiddware) ValidateUser(r *app.Request) bool {
    bearer := r.Header.Get("Authorization")

    if strings.HasPrefix(bearer, "Bearer ") {
        token, err := jwt.Parse(bearer[7:], self.keyFunc)

        if err == nil {
            email := token.Claims["email"].(string)

            user := models.User{}
            user.Find(email)
            r.Env["user"] = &user
            return true
        }

        logging.Logger.Error(fmt.Sprintf("Error: %s", err))
    }

    return false
}

func (self *authMiddware) ValidateAuthority(r *app.Request, role interface{}) bool {
    if r.Env["user"] == nil {
        return false
    }else{
        user := r.Env["user"].(*models.User)
        if user.Role <= role.(int) {
            return true
        }
    }
    return false
}
