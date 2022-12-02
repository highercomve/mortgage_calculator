package mortgagemodels

import (
	"fmt"
	"strings"
)

type Schedule string

const (
	AcceleratedBiWeekly = Schedule("accelerated bi-weekly")
	BiWeekly            = Schedule("bi-weekly")
	Monthly             = Schedule("monthly")
)

var (
	validSchedules = map[Schedule]bool{
		AcceleratedBiWeekly: true,
		BiWeekly:            true,
		Monthly:             true,
	}
)

func PosibleSchedules() string {
	result := []string{}
	for k := range validSchedules {
		result = append(result, string(k))
	}

	return strings.Join(result, ", ")
}

func (s *Schedule) IsValid() error {
	_, ok := validSchedules[*s]
	if !ok {
		return fmt.Errorf(
			"Schedule value is not supported. The supported schedules are: %s",
			PosibleSchedules(),
		)
	}

	return nil
}

func (s *Schedule) PaymentByYear() float64 {
	switch *s {
	case AcceleratedBiWeekly:
		return 26
	case BiWeekly:
		return 26
	case Monthly:
		return 12
	}

	return 12
}
