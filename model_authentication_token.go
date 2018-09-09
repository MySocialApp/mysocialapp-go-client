package msa

type AuthenticationToken struct {
	Model
	Nickname    string `json:"nickname"`
	AccessToken string `json:"access_token"`
}
