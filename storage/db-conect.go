package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	once   sync.Once
	dbprod string = "host=roundhouse.proxy.rlwy.net port=11307 user=postgres password=AdeGAcGd41C5EgDee5CA5gAf6Dbcc-6b dbname=railway sslmode=require"
	//dblocal string = "host=localhost port=5432 user=ceadl password=ceadl2023 dbname=ceadl_info sslmode=disable"
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", dbprod)
		if err != nil {
			log.Fatalf("can`t open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can`t do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

func Pool() *sql.DB {
	return db
}
