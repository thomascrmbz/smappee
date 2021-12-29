package smappee_test

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestGetSensorConsumptions(t *testing.T) {
	sl, err1 := s.GetServiceLocation(62957)
	_, err2 := sl.GetSensorConsumptions(2, 1, time.Now().Add(-15*time.Minute))

	assert.Assert(t, err1)
	assert.Assert(t, err2)
}
