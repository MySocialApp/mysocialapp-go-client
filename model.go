package msa

type Model struct {
	Configuration *Configuration `json:"configuration"`
}

func (m *Model) SetConfiguration(conf *Configuration) {
	m.Configuration = conf
}

type ModelInterface interface {
	SetConfiguration(conf *Configuration)
}
