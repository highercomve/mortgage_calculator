package mortgageservices

import (
	"testing"

	"github.com/highercomve/mortgage_calculator/modules/insurace/insuranceservices"
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgagemodels"
	"github.com/stretchr/testify/assert"
)

func TestDefaultService_Calculate(t *testing.T) {

	t.Run("Downpayment can't be less than 5%", func(t *testing.T) {
		input := &mortgagemodels.InputP{
			Price:         100000,
			DownPayment:   1000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(20),
			Schedule:      mortgagemodels.Monthly,
		}

		service := NewDefaultService(insuranceservices.NewCMHC())
		_, err := service.Calculate(input)

		assert.Equal(t, err.Error(), "downpayment need to be more than 5% and less than 100%", "Downpayment can't be less than 5%")
	})

	t.Run("Downpayment can't be more than 100%", func(t *testing.T) {
		input := &mortgagemodels.InputP{
			Price:         100000,
			DownPayment:   100000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(20),
			Schedule:      mortgagemodels.Monthly,
		}

		service := NewDefaultService(insuranceservices.NewCMHC())
		_, err := service.Calculate(input)

		assert.Equal(t, err.Error(), "downpayment need to be more than 5% and less than 100%", "Downpayment can't be less than 5%")
	})

	t.Run("Period can't be more than 30 years", func(t *testing.T) {
		input := &mortgagemodels.InputP{
			Price:         100000,
			DownPayment:   50000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(40),
			Schedule:      mortgagemodels.Monthly,
		}

		service := NewDefaultService(insuranceservices.NewCMHC())
		_, err := service.Calculate(input)

		assert.Equal(t, err.Error(), "period should be 5 year increments between 5 and 30 years", "Period can't be more than 30 years")
	})

	t.Run("Period can't be less than 5 years", func(t *testing.T) {
		input := &mortgagemodels.InputP{
			Price:         100000,
			DownPayment:   50000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(4),
			Schedule:      mortgagemodels.Monthly,
		}

		service := NewDefaultService(insuranceservices.NewCMHC())
		_, err := service.Calculate(input)

		assert.Equal(t, err.Error(), "period should be 5 year increments between 5 and 30 years", "Period can't be more than 30 years")
	})

	t.Run("Monthly, 5% Downpayment, 5% interest, 200K value, 20 years", func(t *testing.T) {
		input := &mortgagemodels.InputP{
			Price:         300000,
			DownPayment:   40000,
			AnualInterest: 5,
			Period:        mortgagemodels.Period(20),
			Schedule:      mortgagemodels.Monthly,
		}

		service := NewDefaultService(insuranceservices.NewCMHC())
		resp, err := service.Calculate(input)

		assert.Nil(t, err, "error should be nil")
		assert.Equal(t, resp.NumberOfPayments, 240, "number of payment should be 240")
		assert.Equal(t, resp.Payment, 1769.077354544168, "payment should be ")
		assert.Equal(t, string(resp.Schedule), "monthly", "payment should be monthly")
	})

}
