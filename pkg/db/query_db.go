package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	Movie string `json:"movie"`
	Hero  string `json:"hero"`
}

func QueryDb() ([]Movie, error) {
	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open db connection to quer db \n")
		panic(err)
	}
	defer db.Close()

	//DB Query
	query := `SELECT movie_name, hero_name FROM movie;`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Print("unable to prepare statement for query on Movie db \n ")
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("unable to execute query on Movie db \n")
	}

	defer rows.Close()

	var details []Movie
	for rows.Next() {
		var b Movie
		err := rows.Scan(&b.Movie, &b.Hero)
		if err != nil {
			log.Print("unable to scan row in db", err)
			return nil, err
		}
		details = append(details, b)
	}
	return details, nil
}
