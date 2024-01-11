package crudpgdb

import (
	"fmt"
	"log"
)

func UpdateProduct(id int, n string, p int, sp int) {
	product, err := updateProducts(id, &Product{Name: n, Price: p, Supplier_id: sp})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Update Success : ", product)
}
func updateProducts(id int, product *Product) (Product, error) {
	var p Product
	row := DB.QueryRow(
		"UPDATE public.products SET name=$2, price=$3, supplier_id=$4 WHERE id=$1 RETURNING id, name, price, supplier_id, (SELECT name FROM public.suppliers WHERE id = $4) AS supplier_name;",
		id,
		product.Name,
		product.Price,
		product.Supplier_id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_id, &p.Supplier_name)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
