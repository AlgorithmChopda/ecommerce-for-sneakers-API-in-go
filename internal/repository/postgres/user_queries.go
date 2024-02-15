package repository

var InsertUserQuery = `INSERT INTO users (first_name, last_name, email, password, date_of_birth, mobile_no, address, city, postal_code, role_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) `

var GetUserWithEmailAndPassword = `select email, password from users where email=$1 and password=$2`

var GetIdRoleAndPassword = `select id, role_id, password from users where email=$1`

var GetUserWithEmail = `select email from users where email=$1`

var GetBuyerOrSellerList = `SELECT u.first_name, u.last_name, u.email, u.date_of_birth, u.mobile_no, u.address, u.city, 
								   u.postal_code, u.created_at, u.updated_at, r.name  
						  FROM users as u 
						  JOIN role as r ON  r.id = u.role_id 
						  WHERE role_id = $1`

var GetBuyerAndSellerList = `SELECT u.first_name, u.last_name, u.email, u.date_of_birth, u.mobile_no, u.address, u.city, 
									u.postal_code, u.created_at, u.updated_at, r.name 
							 FROM users as u 
							 JOIN role as r ON  r.id = u.role_id 
							 WHERE role_id = $1 OR role_id = $2`

var GetUserWithId = `SELECT u.first_name, u.last_name, u.email, u.date_of_birth, u.mobile_no, u.address, u.city, 
					 u.postal_code, u.created_at, u.updated_at, r.name 
					 FROM users as u 
					 JOIN role as r ON  r.id = u.role_id 
					 WHERE u.id = $1`
