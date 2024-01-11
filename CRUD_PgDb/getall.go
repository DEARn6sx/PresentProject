package crudpgdb

import (
	"fmt"
	"log"
)

func GetAllProduct() {
	product, err := getAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)
}
func getAllProducts() ([]Product, error) {

	rows, err := DB.Query(`
		SELECT pd.id, pd.name, pd.price, s.name, s.id
		FROM public.products pd
		JOIN public.suppliers s ON pd.supplier_id = s.id;`)

	if err != nil {
		return nil , err
	}
	var products []Product

	for rows.Next(){
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Supplier_name, &p.Supplier_id)
		if err != nil {
			return nil , err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil ,err
	}
	return products, nil
}
