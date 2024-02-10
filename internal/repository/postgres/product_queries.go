package repository

const CreateProductQuery = `INSERT INTO product (name, description, seller_id, brand_id) 
							VALUES ($1, $2, $3, $4) Returning id`

const CreateProductDetailQuery = `INSERT INTO productdetail (product_id, size, quantity, price, color, image)
							 VALUES ($1, $2, $3, $4, $5, $6)`

const GetAllProductsAndDetail = `select * from product inner join productdetail on product.id = productdetail.product_id`

const GetProductById = `SELECT p.*, pd.size, pd.color, pd.image, pd.price, pd.quantity,b.name AS brand_name
						FROM product p
						JOIN productdetail pd ON p.id = pd.product_id
						JOIN brand b ON p.brand_id = b.id
						where p.id = $1;`
