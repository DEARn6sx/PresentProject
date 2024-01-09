package main

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"  // or the Docker service name if running in another container
  port     = 5432         // default PostgreSQL port
  user     = "mydearuser"     // as defined in docker-compose.yml
  password = "mydearpassword" // as defined in docker-compose.yml
  dbname   = "mydeardatabase" // as defined in docker-compose.yml
)

func main() {
  // Connection string
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // Open a connection
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // Check the connection
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Successfully connected!")
}