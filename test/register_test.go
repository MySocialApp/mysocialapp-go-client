package msatest

import (
	".."
	"testing"
	"log"
	"fmt"
)

func TestRegister(t *testing.T) {
	client := msa.NewClient("u470584465854a728453", nil)
	email := fmt.Sprintf("%s@mysocialapp.io", randStringRunes(10))
	session, err := client.CreateAccount(email, randStringRunes(10), randStringRunes(10))
	if err != nil {
		t.Errorf("invalid login response: %+v", err)
		return
	}
	log.Printf("account creation ok %+v", session)
}
