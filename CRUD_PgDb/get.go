package crudpgdb

import (
	"fmt"
	"log"
)

func GetProduct(id int) {
	product, err := getProducts(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)
}
func getProducts(id int) (Product, error) {
	var p Product
	row := DB.QueryRow(`
	SELECT pd.id, pd.name, pd.price, s.name, s.id
	FROM public.products pd
	JOIN public.suppliers s ON pd.supplier_id = s.id
	WHERE pd.id = $1;`,
		id)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_name, &p.Supplier_id)
	if err != nil {
		return Product{}, err //return ค่าว่าง product
	}
	return p, nil
}
