package smappee_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gotest.tools/v3/assert"
	"thomascrmbz.com/smappee"
)

var (
	clientID, clientSecret, username, password string
	s                                          *smappee.Smappee
)

func init() {
	godotenv.Load()
	clientID = os.Getenv("SMAPPEE_CLIENT_ID")
	clientSecret = os.Getenv("SMAPPEE_CLIENT_SECRET")
	username = os.Getenv("SMAPPEE_USERNAME")
	password = os.Getenv("SMAPPEE_PASSWORD")
	s, _ = smappee.NewSmappee(clientID, clientSecret, username, password)
}

func TestSmappeeConnection(t *testing.T) {
	_, err := smappee.NewSmappee(clientID, clientSecret, username, password)
	assert.Assert(t, err)
}
