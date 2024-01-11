package crudpgdb

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Supplier_id   int	`json:"supplier_id"`
	Supplier_name string `json:"supplier_name"`
}
