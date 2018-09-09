package msa

import (
	"net/url"
	"fmt"
	"encoding/json"
)

type Notification struct {
	Model
	Id                     int64                  `json:"id"`
	ConfigId               string                 `json:"config_id"`
	Type                   string                 `json:"type"`
	CreatedDate            string                 `json:"created_date"`
	Title                  string                 `json:"title"`
	Description            string                 `json:"description"`
	Url                    *string                `json:"url"`
	ImageURL               *string                `json:"image_url"`
	NotificationKey        *string                `json:"notification_key"`
	RequestAck             bool                   `json:"request_ack"`
	ShowNotification       bool                   `json:"show_notification"`
	ForceNotificationSound bool                   `json:"force_notification_sound"`
	Payload                map[string]interface{} `json:"payload"`
}

func (n *Notification) GetRootUrl() string {
	if n.Url == nil {
		return ""
	}
	u, err := url.Parse(*n.Url)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Hostname())
}

func (n *Notification) GetOwner() *User {
	if base := n.GetPayloadBase(); base != nil {
		return base.Owner
	}
	return nil
}

func (n *Notification) getPayloadString(index string) *string {
	if n.Payload == nil {
		return nil
	}
	if _, ok := n.Payload[index]; !ok {
		return nil
	}
	if str, ok := n.Payload[index].(string); ok {
		return &str
	}
	return nil
}

func (n *Notification) GetRecipientUserId() *string {
	return n.getPayloadString("recipient_user_id")
}

func (n *Notification) GetRecipientDeviceId() *string {
	return n.getPayloadString("recipient_device_id")
}

func (n *Notification) GetPayloadBase() *Base {
	if n.Payload == nil {
		return nil
	}
	var base Base
	j, _ := json.Marshal(n.Payload)
	if err := json.Unmarshal(j, &base); err == nil {
		return &base
	}
	return nil
}

func (n *Notification) Ack() (*NotificationAck, *RestError) {
	// TODO
	return nil, nil
}
