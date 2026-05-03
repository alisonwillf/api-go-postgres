package repository

import (
	"database/sql"
	"fmt"

	"github.com/alisonferreira/APIGO/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	//implementar a query para buscar os produtos no banco de dados
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			continue
		}
		productList = append(productList, p)
	}

	if len(productList) == 0 {
		return []model.Product{}, nil
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	query.Close()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (pr *ProductRepository) GetProductsById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &produto, nil
}
