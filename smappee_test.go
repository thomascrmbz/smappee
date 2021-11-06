package smappee_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/thomascrmbz/smappee"
	"gotest.tools/v3/assert"
)

var (
	clientID, clientSecret, username, password string
)

func init() {
	godotenv.Load()
	clientID = os.Getenv("SMAPPEE_CLIENT_ID")
	clientSecret = os.Getenv("SMAPPEE_CLIENT_SECRET")
	username = os.Getenv("SMAPPEE_USERNAME")
	password = os.Getenv("SMAPPEE_PASSWORD")
}

func TestSmappeeConnection(t *testing.T) {
	_, err := smappee.NewSmappee(clientID, clientSecret, username, password)
	assert.Assert(t, err)
}
