package main

import (
	"errors"
)

type productModel struct {
	products []product
}

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func newProductModel() *productModel {
	m := new(productModel)
	return m
}

func (m *productModel) getProduct(p product) (product, error) {
	for i := 0; i < len(m.products); i++ {
		if m.products[i].ID == p.ID {
			return m.products[i], nil
		}
	}
	return p, errors.New("Not Found")
}

func (m *productModel) updateProduct(p product) error {
	for i := 0; i < len(m.products); i++ {
		if m.products[i].ID == p.ID {
			m.products[i].Name = p.Name
			m.products[i].Price = p.Price
			return nil
		}
	}
	return errors.New("Not Found")
}

func (m *productModel) deleteProduct(ID int) error {
	for i := 0; i < len(m.products); i++ {
		if m.products[i].ID == ID {
			m.products = append(m.products[:i], m.products[i+1:]...)
			return nil
		}
	}
	return errors.New("Not Found")
}

func (m *productModel) createProduct(p product) {
	newProduct := []product{p}
	m.products = append(m.products, newProduct...)
}

func (m *productModel) getProducts() ([]product, error) {
	return m.products, nil
}
