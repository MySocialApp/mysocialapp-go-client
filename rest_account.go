package msa

type RestAccount struct {
	RestBase
	Configuration *Configuration
}

func (r *RestAccount) Get() (*Account, *RestError) {
	var a Account
	return &a, r.Configuration.Get("/account", &a)
}

func (r *RestAccount) Update(account *Account) *RestError {
	return r.Configuration.Put("/account", account, account)
}

