package repository

const CreateCompanyQuery = `insert into companydetail (company_name, company_address) values($1, $2) Returning id`

const InsertSellerQuery = `INSERT INTO users (first_name, last_name, email, password, date_of_birth, mobile_no, address, city, postal_code, role_id, company_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) `
