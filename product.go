package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Id                 string `json:"id"`
	Eancode            string `json:"eancode"`
	Name               string `json:"name"`
	Amount_description string `json:"amount_description"`
	Min_amount         int    `json:"min_amount"`
	Count              int    `json:"count"`
}

func (p *Product) Fill() {
	err := dbsettings.DB.QueryRow("SELECT * FROM items WHERE eancode = ?", p.Eancode).Scan(&p.Id, &p.Eancode, &p.Name, &p.Amount_description, &p.Min_amount, &p.Count)
	if err != nil && err == sql.ErrNoRows {
		// TODO: Handle non extisting products
	} else if err != nil {
		fmt.Println("Error fetching product: ", err)
	}
}

func (p *Product) Subtract(amount int) {
	_, err := dbsettings.DB.Exec("UPDATE items SET count = count - ? WHERE id = ?", amount, p.Id)
	if err != nil {
		fmt.Println("Error updating product: ", err)
	}
	p.Count -= amount

}

func (p *Product) Add(amount int) {
	_, err := dbsettings.DB.Exec("UPDATE items SET count = count + ? WHERE id = ?", amount, p.Id)
	if err != nil {
		fmt.Println("Error updating product: ", err)
	}
	p.Count += amount
}
