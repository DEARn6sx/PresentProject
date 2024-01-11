package crudpgdb

import "log"

func CreateProduct(n string, p int, sp int) {
	err := createProducts(&Product{Name: n, Price: p,Supplier_id: sp})
	if err != nil {
		log.Fatal(err)
	}
	print("Create Success")
}
func createProducts(product *Product) error {
	_, err := DB.Exec(
		"INSERT INTO public.products(name, price, supplier_id)VALUES ($1, $2, $3);",
		product.Name,
		product.Price,
		product.Supplier_id,
	)
	return err
}
