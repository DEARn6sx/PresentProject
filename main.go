package main

import (
	crudpgdb "github.com/DearGo_lang/CRUD_PgDb"
	_ "github.com/lib/pq"
)

func main() {
	crudpgdb.ConnectionDB()
	//crudpgdb.CreateProduct("BENZ",1800000,2)
	//crudpgdb.GetProduct(6)
	//crudpgdb.UpdateProduct(6,"Orange",50,1)
	//crudpgdb.DeleteProduct(7)
	crudpgdb.GetAllProduct()
}
