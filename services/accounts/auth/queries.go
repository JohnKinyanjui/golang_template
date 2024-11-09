package auth_service

const validateQuery = `
	(select count() as orgs from only orgs where <-user[where id = $auth] group all limit 1) ?? {
		orgs: 0
	};
`

// continue with github
const checkIfGithubAccountExists = `select id from users where github_uid = $1`
const signUpGithubQuery = `
	insert into users (
		picture, full_name, email, github_uid, github_token
	) 
	values($1, $2, $3 , $4, $5) returning id;
`

// continue with google
