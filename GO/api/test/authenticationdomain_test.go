package tests

import (
	"log"

	"justasking/GO/api/domain/authentication"
	"justasking/GO/api/startup/boot"
	"justasking/GO/api/startup/env"
	"testing"
)

func init() {
	// Load the configuration file
	config, err := env.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Register the services
	boot.RegisterServices(config)
}

// TestGetGoogleUser tests getting of Google User
func TestGetGoogleUser(t *testing.T) {
	googleID := "11887185481020715099"
	result, err := authenticationrepo.GetGoogleUserBySub(googleID)
	if err != nil {
		t.Error("error is ", err)
	} else if result.ID == 0 {
		t.Error("user does not exist", result)
	}
}

func TestTokenDeserialize(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZmlyc3RuYW1lIjoiU2ViYXN0aWFuIiwibGFzdG5hbWUiOiJDaGFuZGUiLCJlbWFpbCI6InNlYmFzdGlhbi5jaGFuZGVAZ21haWwuY29tIiwic3RhbmRhcmRjbGFpbXMiOnsiZXhwIjoxNDk1OTI4OTE0LCJpc3MiOiJqdXN0YXNraW5nIn19.YD2-lyLhOUrqIWr2KWm6SHQpYzv-Z9kOAfbV0yJ_WjA"
	result := tokendomain.DeserializeToken(token)
	if result == nil {
		t.Error("Deserialize didn't work")
	}
}
