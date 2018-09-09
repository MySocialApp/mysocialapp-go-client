package msa

type RestError struct {
	Status int
	Body   *string
	Error  error
}

func (e *RestError) Message() string {
	return "Error"
}
