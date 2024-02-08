package repository

var InsertUserQuery = `INSERT INTO users (first_name, last_name, email, password, date_of_birth, mobile_no, address, city, postal_code, role_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) `
