package accounts_service

const getAccountQuery = `
	select picture, username, email from users where id = $1
`
