package msa

type UserSettings struct {
	statStatusEnabled *bool                     `json:"stat_status_enabled"`
	notification      *UserSettingsNotification `json:"notification"`
	languageZone      *LanguageZone             `json:"language_zone"`
	interfaceLanguage *InterfaceLanguage        `json:"interface_language"`
}

func (u *UserSettings) IsStatStatusEnabled() bool {
	if u.statStatusEnabled == nil {
		return true
	}
	return *u.statStatusEnabled
}

func (u *UserSettings) GetUserSettingsNotification() UserSettingsNotification {
	if u.notification == nil {
		return UserSettingsNotification{
			Enabled:           true,
			PushEnabled:       true,
			MailEnabled:       true,
			MailFrequency:     MailFrequencyDaily,
			EventEnabled:      true,
			MaximumDistance:   75000,
			MentionEnabled:    true,
			MessagingEnabled:  true,
			CommentEnabled:    true,
			LikeEnabled:       true,
			OfferEnabled:      true,
			SoundEnabled:      true,
			NewsletterEnabled: true,
		}
	}
	return *u.notification
}

func (u *UserSettings) GetLanguageZone() LanguageZone {
	if u.languageZone == nil || *u.languageZone == "" {
		return LanguageZoneFR
	}
	return *u.languageZone
}

func (u *UserSettings) GetInterfaceLanguage() InterfaceLanguage {
	if u.languageZone == nil || *u.languageZone == "" {
		return InterfaceLanguageFR
	}
	return *u.interfaceLanguage
}

type InterfaceLanguage string
type LanguageZone string
type MailFrequency string

const (
	InterfaceLanguageFR InterfaceLanguage = "FR"
	InterfaceLanguageEN InterfaceLanguage = "EN"
	InterfaceLanguageDE InterfaceLanguage = "DE"

	LanguageZoneFR LanguageZone = "FR"
	LanguageZoneEN LanguageZone = "EN"
	LanguageZoneDE LanguageZone = "DE"

	MailFrequencyNever    MailFrequency = "NEVER"
	MailFrequencyRealTime MailFrequency = "REAL_TIME"
	MailFrequencyDaily    MailFrequency = "DAILY"
	MailFrequencyWeekly   MailFrequency = "WEEKLY"
)

var (
	DefaultInterfaceLanguage = InterfaceLanguageFR
	DefaultLanguageZone      = LanguageZoneFR
)

type UserSettingsNotification struct {
	Enabled           bool          `json:"enabled"`
	PushEnabled       bool          `json:"push_enabled"`
	MailEnabled       bool          `json:"mail_enabled"`
	MailFrequency     MailFrequency `json:"mail_frequency"`
	EventEnabled      bool          `json:"event_enabled"`
	MaximumDistance   int           `json:"maximum_distance"`
	MentionEnabled    bool          `json:"mention_enabled"`
	MessagingEnabled  bool          `json:"messaging_enabled"`
	CommentEnabled    bool          `json:"comment_enabled"`
	LikeEnabled       bool          `json:"like_enabled"`
	OfferEnabled      bool          `json:"offer_enabled"`
	SoundEnabled      bool          `json:"sound_enabled"`
	NewsletterEnabled bool          `json:"newsletter_enabled"`
}
