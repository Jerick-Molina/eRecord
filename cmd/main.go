package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB"}
	_, e := sql.Open("mysql", cfg.FormatDSN())

	if e != nil {

	}
}
