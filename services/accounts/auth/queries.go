package auth_service

// continue with google
const signInWithGoogleQuery = `
	select id from users where email = @email and google_uid = @google_uid
`

const updateWithGoogleQuery = `
	update users 
		set email = @email, google_uid = @google_uid where id = @id 
	returning id;
`

// continue with google
