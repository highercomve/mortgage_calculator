package validation

import (
	"fmt"
	"testing"

	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgagemodels"
	"github.com/stretchr/testify/assert"
)

const (
	supportedError = "Schedule value is not supported. The supported schedules are: accelerated bi-weekly, bi-weekly, monthly"
)

func Test_scheduleValidate(t *testing.T) {
	validate := GetValidator()

	t.Run("accelerated bi-weekly shoul.d be valid", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(5),
			Schedule:      mortgagemodels.AcceleratedBiWeekly,
		}
		err := validate.Validate(x)
		assert.Equal(t, err, nil, fmt.Sprintf("%s should be supported", mortgagemodels.AcceleratedBiWeekly))
	})

	t.Run("bi-weekly should be valid", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(5),
			Schedule:      mortgagemodels.BiWeekly,
		}
		err := validate.Validate(x)
		assert.Equal(t, err, nil, fmt.Sprintf("%s should be supported", mortgagemodels.BiWeekly))
	})

	t.Run("every 3 month should be invalid", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(5),
			Schedule:      mortgagemodels.Schedule("every 3 month"),
		}
		err := validate.Validate(x)

		assert.NotNil(t, err, "error should not be nil")
		assert.Error(t, err, supportedError, "Schedule shouldn't be supported")
	})

	t.Run("monthly should be valid", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(5),
			Schedule:      mortgagemodels.Monthly,
		}
		err := validate.Validate(x)
		assert.Equal(t, err, nil, fmt.Sprintf("%s should be supported", mortgagemodels.Monthly))
	})
}

func Test_isValidPeriod(t *testing.T) {
	validate := GetValidator()

	t.Run("period can't be less than 5 years", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(1),
			Schedule:      mortgagemodels.Monthly,
		}
		err := validate.Validate(x)
		assert.Error(t, err, "period should be 5 year increments between 5 and 30 years", "Period should be more than 5 years")
	})

	t.Run("period can't be more than 30 years", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(31),
			Schedule:      mortgagemodels.Monthly,
		}
		err := validate.Validate(x)
		assert.Error(t, err, "period should be 5 year increments between 5 and 30 years", "Period should be more than 5 years")
	})

	t.Run("period should be a 5 years increment", func(t *testing.T) {
		x := &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(6),
			Schedule:      mortgagemodels.Monthly,
		}
		err := validate.Validate(x)
		assert.Error(t, err, "period should be 5 year increments between 5 and 30 years", "Period should be more than 5 years")

		x = &mortgagemodels.InputP{
			Price:         200000,
			DownPayment:   10000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(10),
			Schedule:      mortgagemodels.Monthly,
		}
		err = validate.Validate(x)
		assert.Nil(t, err, "10 years is a valid period")
	})
}
