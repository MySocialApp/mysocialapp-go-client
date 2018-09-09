package msa

func NewClient(appId string, clientConf *ClientConfiguration) *MySocialApp {
	m := MySocialApp{
		Configuration: NewConfiguration(appId, clientConf),
	}
	return &m
}

type MySocialApp struct {
	Configuration *Configuration
}

func (m *MySocialApp) Connect(login string, password string) (*Session, *RestError) {
	session := newSession(m.Configuration)
	at, err := session.ClientService.Login().Login(&LoginCredentials{Username: login, Password: password})
	if err != nil {
		return nil, err
	}
	session.Configuration.ClientConfiguration.SetToken(at.AccessToken)
	return session, nil
}

func (m *MySocialApp) CreateAccount(email string, password string, firstName string) (*Session, *RestError) {
	session := newSession(m.Configuration)
	_, err := session.ClientService.Register().Create(User{Email: email, Password: &password, FirstName: firstName})
	if err != nil {
		return nil, err
	}
	return m.Connect(email, password)
}
