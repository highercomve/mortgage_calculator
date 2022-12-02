package mortgagemodels

import "fmt"

type Period int

func (p *Period) IsValid() error {
	if *p > 30 {
		return fmt.Errorf("Period should not be more than 30 years")
	}
	if *p < 5 {
		return fmt.Errorf("Period should not be less than 5 years")
	}

	if *p%5 > 0 {
		return fmt.Errorf("Period can be only increments of 5 years")
	}

	return nil
}
