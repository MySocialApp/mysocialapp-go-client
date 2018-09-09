package msa

type Comment struct {
	Base
	Message *string `json:"message"`
	Photo   *Photo  `json:"photo"`
	Parent  *Base   `json:"parent"`
}

func (c *Comment) ReplyBack() (*Comment, *RestError) {
	// TODO
	return nil, nil
}

func (c *Comment) Update() *RestError {
	// TODO
	return nil
}

func (c *Comment) Delete() *RestError {
	// TODO
	return nil
}
