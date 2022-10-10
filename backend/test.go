package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func main() {
	dataSource := fmt.Sprintf("%s:%s@(%s)/%s",
		getEnv("MYSQL_USER", ""),
		getEnv("MYSQL_ROOT_PASSWORD", ""),
		getEnv("MYSQL_HOSTNAME", ""),
		getEnv("MYSQL_DB_NAME", ""),
	)
	fmt.Println("hoge", dataSource)
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
