package repository

// 0 = false, 1 = true

const CreateOrder = `INSERT INTO orders (total_amount, buyer_id, is_cart) VALUES (0, $1, '1') Returning id`

// checks if the buyer has a cart or not ?
const OrderWithUserId = `SELECT * FROM orders where buyer_id = $1 and is_cart = '1'`

const GetProductQuantityAndPrice = `SELECT quantity, price FROM product JOIN productdetail ON product.id = productdetail.product_id where productdetail.id = $1`

const UpdateOrderAmount = `update orders set total_amount = total_amount + $2 where id = $1`

const AddProductToOrder = `INSERT INTO orderitem (product_detail_id, order_id, price, quantity) VALUES ($1, $2, $3, $4)`

const GetBuyerIdOfOrder = `SELECT buyer_id FROM orders where orders.id = $1`

const UpdateOrderItem = `UPDATE orderitem SET quantity = $4, price = $3 WHERE order_id = $2 AND product_detail_id = $1`

const GetOrderItemPriceAndQuantity = `SELECT quantity, price FROM orderitem where order_id = $1 AND product_detail_id = $2`

const PlaceOrder = `UPDATE orders SET is_cart = '0', order_date = CURRENT_TIMESTAMP, shipping_address = $2 where id = $1`

const GetOrderItemCount = `SELECT COUNT(*) AS item_count FROM orderitem WHERE order_id = $1`

const OrderWithOrderId = `SELECT * FROM orders where id = $1 and is_cart = '1'`

const GetOrderItems = ` SELECT 	p.name, p.description, pd.size, pd.color, pd.image, pd.price, oi.quantity
						FROM orderitem AS oi
						JOIN productdetail AS pd
						ON oi.product_detail_id = pd.id
						JOIN product AS p
						ON pd.product_id = p.id
						WHERE oi.order_id = $1`

const GetOrderAmount = `SELECT total_amount FROM orders where id = $1`

const CheckOrderValid = `SELECT * FROM orders where buyer_id = $1 and id = $2 and is_cart = '1'`

const GetPlacedOrderDetails = `SELECT total_amount, order_date, shipping_address FROM orders where id = $1 AND buyer_id = $2 AND is_cart = '0'`

const GetUserOrders = `SELECT id, total_amount, order_date, shipping_address FROM orders where buyer_id = $1 AND is_cart = '0'`

const GetOrderItemProductIdAndQuantity = `SELECT product_detail_id, quantity 
										  FROM orderitem JOIN orders ON orderitem.order_id = orders.id 
										  WHERE orders.id = $1`

const GetProductQuantity = `SELECT quantity FROM productdetail where productdetail.id = $1`

const DeleteItemFromOrder = `DELETE FROM orderitem WHERE order_id = $1 AND product_detail_id = $2`
