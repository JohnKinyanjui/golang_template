package accounts_service

const getAccountQuery = `
	select picture, full_name, email from users where id = $1
`
