package repository

const CreateProductQuery = `INSERT INTO product (name, description, seller_id, brand_id) 
							VALUES ($1, $2, $3, $4) Returning id`

const CreateProductDetailQuery = `INSERT INTO productdetail (product_id, size, quantity, price, color, image)
							 VALUES ($1, $2, $3, $4, $5, $6)`
