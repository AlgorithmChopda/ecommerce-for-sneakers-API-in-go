package repository

// 0 = false, 1 = true

const CreateOrder = `INSERT INTO orders (total_amount, buyer_id, is_cart) VALUES (0, $1, '0') Returning id`

const OrderWithUserId = `SELECT * FROM orders where buyer_id = $1 and is_cart = '0'`

const GetProductQuantityAndPrice = `SELECT quantity, price FROM product JOIN productdetail ON product.id = productdetail.product_id where productdetail.id = $1`

const UpdateOrderAmount = `update orders set total_amount = total_amount + $2 where id = $1`

const AddProductToOrder = `INSERT INTO orderitem (product_detail_id, order_id, price, quantity) VALUES ($1, $2, $3, $4)`

const GetBuyerIdOfOrder = `SELECT buyer_id FROM orders where orders.id = $1`

const UpdateOrderItem = `UPDATE orderitem SET quantity = $4, price = $3 WHERE order_id = $2 AND product_detail_id = $1`

const GetOrderItemPriceAndQuantity = `SELECT quantity, price FROM orderitem where order_id = $1 AND product_detail_id = $2`
