package msa

type Location struct {
	Base
	Location            *SimpleLocation `json:"location"`
	Country             *string         `json:"country"`
	District            *string         `json:"district"`
	State               *string         `json:"state"`
	PostalCode          *string         `json:"postal_code"`
	City                *string         `json:"city"`
	StreetName          *string         `json:"street_name"`
	StreetNumber        *string         `json:"street_number"`
	CompleteAddress     *string         `json:"complete_address"`
	CompleteCityAddress *string         `json:"complete_city_address"`
}

func (l *Location) GetLatitude() *float64 {
	if l.Location != nil {
		return &l.Location.Latitude
	}
	return nil
}

func (l *Location) GetLongitude() *float64 {
	if l.Location != nil {
		return &l.Location.Longitude
	}
	return nil
}

func (l *Location) Equal(location SimpleLocation) bool {
	if l.Location == nil {
		return false
	}
	return l.Location.Longitude == location.Longitude && l.Location.Latitude == location.Latitude
}