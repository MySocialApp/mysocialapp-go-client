package msa

type Base struct {
	Model
	Id             *int64         `json:"id"`
	IdStr          *string        `json:"id_str"`
	Type           *string        `json:"type"`
	CreatedDate    *string        `json:"created_date"`
	DisplayedName  *string        `json:"displayed_name"`
	DisplayedPhoto *Photo         `json:"displayed_photo"`
	EntityType     *EntityType    `json:"entity_type"`
	AccessControl  *AccessControl `json:"access_control"`
	Owner          *User          `json:"owner"`
}

func Delete(b *Base) *RestError {
	// to override
	return nil
}

func Save(b *Base) *RestError {
	// to override
	return nil
}
