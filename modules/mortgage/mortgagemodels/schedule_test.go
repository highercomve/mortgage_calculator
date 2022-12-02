package mortgagemodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	supportedError = "Schedule value is not supported. The supported schedules are: accelerated bi-weekly, bi-weekly, monthly"
)

func TestSchedule_IsValid(t *testing.T) {
	t.Run("Schedule is valid", func(t *testing.T) {
		s := Schedule("every 3 month")

		assert.Errorf(t, s.IsValid(), "validation should fail")
		assert.EqualError(t, s.IsValid(), supportedError, "schedule should fail")

		s = Schedule("accelerated bi-weekly")
		assert.Equal(t, s.IsValid(), nil, "schedule should be valid")

		s = Schedule("bi-weekly")
		assert.Equal(t, s.IsValid(), nil, "schedule should be valid")

		s = Schedule("monthly")
		assert.Equal(t, s.IsValid(), nil, "schedule should be valid")
	})
}
