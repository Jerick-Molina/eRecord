package main

import (
	"database/sql"
	"eRecord/cmd/server"
	"eRecord/internal/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB"}
	dbConn, e := sql.Open("mysql", cfg.FormatDSN())

	if e != nil {

	}
	record := db.NewRecord(dbConn)

	server := server.NewServer(record)

	server.Start("localhost:8080")

}
