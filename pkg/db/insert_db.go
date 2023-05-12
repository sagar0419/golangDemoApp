package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InsertMovie(name, hero string) {
	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open Db connection to add Movie details \n")
		panic(err)
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Query
	query := fmt.Sprintf("INSERT INTO movie(movie_name, hero_name) VALUES('%s','%s');", name, hero)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Unable to execute query to insert Movie detail in db \n")
		panic(err)
	}
	log.Printf("Movie added successfully \n")
}
