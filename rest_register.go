package msa

type RestRegister struct {
	RestBase
	Configuration *Configuration
}

func (r *RestRegister) Create(user User) (*User, *RestError) {
	var u User
	return &u, r.Configuration.Post("/register", &u, user)
}
