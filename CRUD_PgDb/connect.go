package crudpgdb

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"      // or the Docker service name if running in another container
	port     = 5432             // default PostgreSQL port
	user     = "dearuser"     // as defined in docker-compose.yml
	password = "dearpassword" // as defined in docker-compose.yml
	dbname   = "deardatabase" // as defined in docker-compose.yml
)

// ประกาศ db เป็น global เพื่อให้ db เป็นตัวแทนในการพูดคุย
var DB *sql.DB

func ConnectionDB() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	DB = sdb

	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}
