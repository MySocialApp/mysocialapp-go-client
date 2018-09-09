package msa

type FluentAccount struct {
	Session *Session
}

func (f *FluentAccount) Get() (*Account, *RestError) {
	return f.Session.ClientService.Account().Get()
}