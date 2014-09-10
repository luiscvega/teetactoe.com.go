package logic

import (
        "database/sql"
)

var db *sql.DB

func Init(database *sql.DB) (err error) {
        db = database
        return
}
