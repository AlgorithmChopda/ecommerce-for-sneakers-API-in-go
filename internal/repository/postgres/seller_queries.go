package repository

const CreateCompanyQuery = `insert into companydetail (company_name, company_address) values($1, $2) Returning id`

const InsertSellerQuery = `INSERT INTO users (first_name, last_name, email, password, date_of_birth, mobile_no, address, city, postal_code, role_id, company_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) `

const GetAllSellers = `SELECT first_name, last_name, email, date_of_birth, mobile_no, address, city, postal_code, created_at, 
							  updated_at, cd.company_name, cd.company_address
						FROM users
						JOIN companydetail as cd
						ON users.company_id = cd.id
						WHERE role_id = $1;`

const DeleteSeller = `DELETE FROM users where id = $1`
