package msa

type BaseImpl interface {
	delete() *RestError
	save() *RestError
}
