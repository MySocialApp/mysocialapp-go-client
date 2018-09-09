package msa

type RestLogin struct {
	RestBase
	Configuration *Configuration
}

func (r *RestLogin) Login(creds *LoginCredentials) (*AuthenticationToken, *RestError) {
	var l AuthenticationToken
	return &l, r.Configuration.Post("/login", &l, creds)
}
