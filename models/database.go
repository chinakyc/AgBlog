package models

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)


var db gorm.DB

func OpenDB (dialect string, path string) {
    var err error

    db, err = gorm.Open(dialect, path)

    db.LogMode(true)

    if err != nil {
        log.Fatal(err)
    }
}

