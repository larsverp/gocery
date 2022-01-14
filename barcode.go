package main

type Barcode struct {
	EAN string
}

func (b *Barcode) Scaned() Product {
	product := Product{Eancode: b.EAN}
	product.Fill()
	return product
}
