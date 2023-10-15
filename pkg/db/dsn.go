package db

import (
	"fmt"
	"os"
)

// Set these as env variables to use in code Or you can use your own envs.
// Below are the commands that I have used to set the env variabels
// export MYSQL_USER="root"
// export MYSQL_PASSWORD="password"
// export MYSQL_DATABASE="db"
// export DBNAME="citadel"
// export HSTNAME="127.0.0.1:3306"

var (
	username  = os.Getenv("MYSQL_USER")
	password  = os.Getenv("MYSQL_PASSWORD")
	dbname    = os.Getenv("DBNAME")
	olddbname = os.Getenv("MYSQL_DATABASE")
	hostname  = os.Getenv("HSTNAME")
)

// Creating mysql connection string "user:password@tcp(127.0.0.1:3306)/database_name"

func CreatDbDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, olddbname)
}

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
