package repository

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
					   WHERE product.id = $1; `

const UpdateProductDetail = `UPDATE productdetail
							 SET quantity = $2
							 WHERE id = $1`
