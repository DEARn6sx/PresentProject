package main

import (
	crudpgdb "github.com/DearGo_lang/CRUD_PgDb"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	crudpgdb.ConnectionDB()
	//crudpgdb.CreateProduct("BENZ",1800000,2)
	//crudpgdb.GetProduct(6)
	//crudpgdb.UpdateProduct(6,"Orange",50,1)
	//crudpgdb.DeleteProduct(7)
	//crudpgdb.GetAllProduct()

	app := fiber.New()

	app.Get("/", crudpgdb.Hello)
	app.Get("/product/:id", crudpgdb.GetProductHandler)
	app.Get("/product", crudpgdb.GetallProductHandler)
	app.Post("/product", crudpgdb.CreateProductHandler)
	app.Put("/product/:id", crudpgdb.UpdateProductHandler)
	app.Delete("/product/:id", crudpgdb.DeleteProductHandler)

	app.Listen(":8080")

	defer crudpgdb.DB.Close()
}
