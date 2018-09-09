package msa

type NotificationAck struct {
	DeviceId           *string     `json:"device_id"`
	AppPlatform        AppPlatform `json:"app_platform"`
	AppVersion         *string     `json:"app_version"`
	AdvertisingId      *string     `json:"advertising_id"`
	NotificationKey    *string     `json:"notification_key"`
	NotificationAction *string     `json:"notification_action"`
	Location           *Location   `json:"location"`
}

type AppPlatform string

const (
	AppPlatformAndroid AppPlatform = "ANDROID"
	AppPlatformSdk     AppPlatform = "SDK"
)
