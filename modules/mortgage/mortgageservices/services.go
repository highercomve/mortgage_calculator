package mortgageservices

import (
	"math"

	"github.com/highercomve/mortgage_calculator/modules/insurace/insuranceservices"
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgagemodels"
	"github.com/highercomve/mortgage_calculator/modules/validation"
)

type Service interface {
	Calculate(input *mortgagemodels.InputP) *mortgagemodels.OutputP
}

type DefaultService struct {
	insurance insuranceservices.InsuranceService
}

func NewDefaultService(i insuranceservices.InsuranceService) *DefaultService {
	return &DefaultService{
		insurance: i,
	}
}

func (s *DefaultService) Calculate(input *mortgagemodels.InputP) (*mortgagemodels.OutputP, error) {
	validate := validation.GetValidator()
	err := validate.Validate(input)
	if err != nil {
		return nil, validation.TranslateError(err)
	}

	iPercentage, err := s.insurance.Calculate(input.Price, input.DownPayment)
	if err != nil {
		return nil, err
	}

	r := input.AnualInterest / 12 / 100
	n := float64(input.Period) * 12
	restPrice := input.Price - input.DownPayment
	price := ((restPrice * iPercentage / 100) + restPrice)
	compoundInterest := math.Pow(1+r, n)
	monthlyPayment := price * ((r * compoundInterest) / (compoundInterest - 1))
	payment := s.getPayment(input.Schedule, monthlyPayment)
	numberOfpayments := s.getNumberOfPayments(
		monthlyPayment,
		float64(input.Period),
		payment,
	)

	payload := &mortgagemodels.OutputP{
		Schedule:         input.Schedule,
		Payment:          payment,
		NumberOfPayments: numberOfpayments,
	}

	return payload, nil
}

func (s *DefaultService) getNumberOfPayments(mp float64, p float64, payment float64) int {
	return int(math.Round((mp * 12 * p) / payment))
}

func (s *DefaultService) getPayment(schedule mortgagemodels.Schedule, monthly float64) float64 {
	switch schedule {
	case mortgagemodels.AcceleratedBiWeekly:
		return monthly / 2
	case mortgagemodels.BiWeekly:
		return (monthly * 12) / schedule.PaymentByYear()
	default:
		return monthly
	}
}
