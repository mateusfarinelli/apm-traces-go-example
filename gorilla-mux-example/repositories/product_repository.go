package repositories

import (
	"apm-trace-with-gorilla-mux-example/interfaces"
	"apm-trace-with-gorilla-mux-example/model"
	"context"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository (conn *sql.DB) interfaces.ProductRepository {
	return &ProductRepository{
		connection: conn,
	}
}

func (pr *ProductRepository) GetProducts(ctx context.Context) ([]model.Product, error){	
	query:= "SELECT id, product_name, price FROM product"

	rows, err := pr.connection.QueryContext(ctx, query)

	if(err != nil) {
		fmt.Println(err)

		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next(){
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price)

		if err != nil {
			fmt.Println(err)

			return []model.Product{}, err
		}

		productList = append(productList, productObject)
	}
	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error){
	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id = $1")

	if(err != nil) {
		fmt.Println(err)

		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)

	if err != nil {
		fmt.Println(err)

		if(err == sql.ErrNoRows){
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil

}