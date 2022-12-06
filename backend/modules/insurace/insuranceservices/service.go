package insuranceservices

import "errors"

type InsuranceService interface {
	Calculate(price float64, downPayment float64) (float64, error)
}

type CMHC struct{}

func NewCMHC() *CMHC {
	return &CMHC{}
}

func (s *CMHC) Calculate(price float64, downPayment float64) (float64, error) {
	percentage := downPayment * 100 / price
	if percentage < 5 {
		return 0, errors.New("downpayment can't be less than 5%")
	}
	if percentage >= 5 && percentage < 10 {
		return 4.00, nil
	}
	if percentage >= 10 && percentage < 15 {
		return 3.10, nil
	}
	if percentage >= 15 && percentage < 20 {
		return 2.80, nil
	}

	return 0, nil
}
