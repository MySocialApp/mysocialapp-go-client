package msa

type Account struct {
	Base                 `json:"-"`
	Id            int64  `json:"id"`
	IdStr         string `json:"id_str"`
	FirstName     string `json:"first_name"`
	UpdatedDate   string `json:"updated_date"`
	Email         string `json:"email"`
	DisplayedName string `json:"displayed_name"`
}

func (a *Account) Update() *RestError {
	return RestAccount{Configuration: a.Configuration}.Update(a)
}
