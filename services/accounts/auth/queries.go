package auth_service

// continue with github
const checkIfGithubAccountExists = `select id from users where github_uid = $1`
const signUpGithubQuery = `
	insert into users (
		picture, full_name, email, github_uid, github_token
	) 
	values(@picture, @full_name, @email , @github_uid, @github_token) returning id;
`

// continue with google
