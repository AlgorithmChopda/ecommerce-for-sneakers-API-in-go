package repository

import (
	"fmt"
	"strings"
)

const CreateProductQuery = `INSERT INTO product (name, description, seller_id, brand_id) 
							VALUES ($1, $2, $3, $4) Returning id`

const CreateProductDetailQuery = `INSERT INTO productdetail (product_id, size, quantity, price, color, image)
							 VALUES ($1, $2, $3, $4, $5, $6)`

const GetAllProductsAndDetail = `SELECT * FROM product JOIN productdetail ON product.id = productdetail.product_id`

const GetProductById = `SELECT p.*, pd.size, pd.color, pd.image, pd.price, pd.quantity,b.name AS brand_name
						FROM product p
						JOIN productdetail pd ON p.id = pd.product_id
						JOIN brand b ON p.brand_id = b.id
						where p.id = $1;`

const UpdateProduct = `UPDATE product
					   SET name = $2,description = $3, updated_at = CURRENT_TIMESTAMP
					   WHERE product.id = $1;`

const UpdateProductDetail = `UPDATE productdetail
							 SET quantity = $2
							 WHERE id = $1 AND seller_id = $4`

func getQueryForFilters(filters map[string]string, skip, limit int) string {
	var rawQuery string = `SELECT p.*, pd.size, pd.color, pd.image, pd.price, pd.quantity, pd.id, b.name AS brand_name 
						   from product as p 
						   JOIN productdetail as pd ON p.id = pd.product_id 
						   JOIN brand b ON p.brand_id = b.id `

	isFirstKey := true
	for key, value := range filters {
		if isFirstKey {
			rawQuery += "where "
			isFirstKey = false
		}
		if key == "name" {
			rawQuery += fmt.Sprintf("b.%s = '%s' AND ", key, value)
		}

		if key == "color" {
			rawQuery += fmt.Sprintf("pd.%s = '%s' AND ", key, value)
		}

		if key == "size" {
			rawQuery += fmt.Sprintf("pd.%s = %s AND ", key, value)
		}
	}
	rawQuery = strings.TrimSuffix(rawQuery, "AND ")
	rawQuery += fmt.Sprintf(" ORDER BY p.id ")
	// rawQuery += fmt.Sprintf(" ORDER BY id OFFSET %d LIMIT %d", skip, limit)

	return rawQuery
}
