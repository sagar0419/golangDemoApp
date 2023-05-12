package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CheckDb(moviename string) string {
	query := fmt.Sprintf("SELECT hero_name FROM movie WHERE movie_name REGEXP '%s';", moviename)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open DB to check Movies record \n")
		panic(err)
	}
	defer db.Close()
	var Movierecord string
	err = db.QueryRowContext(ctx, query).Scan(&Movierecord)
	if err != nil {
		log.Printf("Unable to exec query on db to check Movie record \n")
		return ""
	}
	return Movierecord
}
