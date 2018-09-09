package msa

type TextWallMessage struct {
	BaseWall
	Message *string `json:"message"`
}

func (t *TextWallMessage) Update() *RestError {
	// TODO
	return nil
}

func (t *TextWallMessage) Delete() *RestError {
	// TODO
	return nil
}

