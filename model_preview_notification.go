package msa

type PreviewNotification struct {
	ConfigId         string        `json:"config_id"`
	Type             string        `json:"type"`
	Id               int64         `json:"id"`
	IdStr            string        `json:"id_str"`
	Total            int           `json:"total"`
	LastNotification *Notification `json:"last_notification"`
}

func (p *PreviewNotification) Consume() (*PreviewNotification, *RestError){
	// TODO
	return nil, nil
}
