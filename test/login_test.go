package msatest

import (
	"testing"
	"log"
	".."
)

func TestLogin(t *testing.T) {
	client := msa.NewClient("u470584465854a728453", nil)
	session, err := client.Connect("alice.jeith@mysocialapp.io", "myverysecretpassw0rd")
	if err != nil {
		t.Errorf("invalid login response: %+v", err)
		return
	}
	log.Printf("connect ok %+v", session)

	account, restErr := session.Account().Get()
	if restErr != nil {
		t.Errorf("invalid response from account: %+v", restErr)
	}
	log.Printf("account %+v", account)
}