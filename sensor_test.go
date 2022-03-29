package smappee_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetSensorConsumptions(t *testing.T) {
	sl, err1 := s.GetServiceLocation(38248)

	err2 := sl.SetActuator(2, false, 300)

	assert.Assert(t, err1)
	assert.Assert(t, err2)
}
