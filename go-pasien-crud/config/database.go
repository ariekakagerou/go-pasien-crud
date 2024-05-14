package config

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

// DBConnection mengembalikan koneksi ke database
func DBConnection() (*sql.DB, error) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "go_crud_rs"

    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    return db, err
}
