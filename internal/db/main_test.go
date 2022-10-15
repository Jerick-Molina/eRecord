package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	cfg := mysql.Config{
		User:   "buggie",
		Passwd: "Mixon9090",
		Net:    "tcp",
		Addr:   "192.168.3.139",
		DBName: "BuggieDB"}
	conn, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("Cannot connect to dbL: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
