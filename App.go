package main

import (
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jmoiron/sqlx"
)

func main() {
	Server().Run()
}
