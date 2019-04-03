package schema

import (
	"encoding/json"
	"fmt"
	"webapp-course-be/db/dbmodel"
)

var ProducsContainer *Products
var prod map[int]dbmodel.Product

func NewProducts() *Products {
	if ProducsContainer == nil {
		prod = make(map[int]dbmodel.Product)
		ProducsContainer = &Products{
			products: prod,
			nextId:   1,
		}
		return ProducsContainer
	}
	return ProducsContainer
}

type Products struct {
	products map[int]dbmodel.Product
	nextId   int
}

func (p *Products) Insert(product dbmodel.Product) {
	product.Id = p.nextId
	p.products[product.Id] = product
	p.nextId++
	p.DebugInConsole()
}

func (p *Products) GetAll() (products []dbmodel.Product) {
	for _, product := range p.products {
		products = append(products, product)
	}
	return
}

func (p *Products) GetById(id *int) *dbmodel.Product {
	product, ok := p.products[*id]
	if !ok {
		return nil
	}
	return &product
}

func (p *Products) Update(product *dbmodel.Product) *int {
	_, ok := p.products[product.Id]
	if !ok {
		return nil
	}
	delete(p.products, product.Id)
	p.products[product.Id] = *product
	p.DebugInConsole()
	return &product.Id
}

func (p *Products) Delete(id *int) *int {
	_, ok := p.products[*id]
	if !ok {
		return nil
	}
	delete(p.products, *id)
	p.DebugInConsole()
	return id
}

func (p *Products) DebugInConsole() {
	jsonProds, _ := json.Marshal(p.products)
	fmt.Println("\n::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println(string(jsonProds))
}
