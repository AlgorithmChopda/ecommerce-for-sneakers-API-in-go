package repository

// 0 = false, 1 = true

const CreateOrder = `INSERT INTO orders (total_amount, buyer_id, is_cart) VALUES (0, $1, '0') Returning id`

const OrderWithUserId = `SELECT * FROM orders where buyer_id = $1 and is_cart = '0'`
