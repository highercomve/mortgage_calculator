package insuranceservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMHC_Calculate(t *testing.T) {
	t.Run("less than 5% should return error", func(t *testing.T) {
		s := NewCMHC()
		i, err := s.Calculate(100000, 1000)
		assert.Equal(t, err.Error(), "downpayment can't be less than 5%", "")
		assert.Equal(t, i, float64(0), "downpayment can't be less than 5%", "")
	})

	t.Run("5% to 9.99% should return 4%", func(t *testing.T) {
		s := NewCMHC()
		i, err := s.Calculate(100000, 5000)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(4), "downpayment can't be less than 5%", "")

		i, err = s.Calculate(100000, 9999)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(4), "downpayment can't be less than 5%", "")
	})

	t.Run("10% to 14.99% should return 3.1%", func(t *testing.T) {
		s := NewCMHC()
		i, err := s.Calculate(100000, 10000)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(3.1), "downpayment can't be less than 5%", "")

		i, err = s.Calculate(100000, 14999)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(3.1), "downpayment can't be less than 5%", "")
	})

	t.Run("15% to 19.99% should return 2.8%", func(t *testing.T) {
		s := NewCMHC()
		i, err := s.Calculate(100000, 15000)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(2.8), "downpayment can't be less than 5%", "")

		i, err = s.Calculate(100000, 19999)
		assert.Nil(t, err, "")
		assert.Equal(t, i, float64(2.8), "downpayment can't be less than 5%", "")
	})
}
