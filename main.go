package main

import (
	crudpgdb "github.com/DearGo_lang/CRUD_PgDb"
	_ "github.com/lib/pq"
)

func main() {
	crudpgdb.ConnectionDB()
	crudpgdb.CreateProduct("BMW",1500000,2)
}
