# This is a GoLang demo app

In this app you can upload your file , save data (Movie and hero name) and you can also list your saved files and movies.

This app use MySql Database to store all info and it runs on port 3000.

You need to few values as env to run this code.

Set these as env variables to use in code Or you can use your own envs.

```sh
export MYSQL_USER="root"
export MYSQL_PASSWORD="password"
export MYSQL_DATABASE="db"
export DBNAME="hey"
export HSTNAME="127.0.0.1:3306"
```

`MYSQL_USER` is the username to access the DB.
`MYSQL_PASSWORD` is the password to access your mysql db.
`MYSQL_DATABASE` is the old db name which is created by the mysql by default.
`DBNAME` is the name of the DB which you want to use for this app.
`HSTNAME` is the hostname of the database. You can pass `HSTNAME` which is hostname as string also.