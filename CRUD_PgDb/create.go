package crudpgdb

import "log"

func CreateProduct(n string, p int) {
	err := createProducts(&Product{Name: n, Price: p})
	if err != nil {
		log.Fatal(err)
	}
	print("Create Success")
}
func createProducts(product *Product) error {
	_, err := DB.Exec(
		"INSERT INTO public.products(name, price)VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}
