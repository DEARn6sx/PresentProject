package crudpgdb

import (
	"fmt"
	"log"
)

func DeleteProduct(id int) {
	err := deleteProducts(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete Product Successfull")
}
func deleteProducts(id int) error {
	_, err := DB.Exec(
		"DELETE FROM public.products WHERE id = $1;",
		id,
	)
	return err
}
