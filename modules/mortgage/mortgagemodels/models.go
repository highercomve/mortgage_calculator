package mortgagemodels

type InputP struct {
	Schedule      Schedule `json:"schedule" validate:"required,is-schedule"`
	Period        Period   `json:"period" validate:"required,is-period"`
	Price         float64  `json:"price" validate:"required"`
	DownPayment   float64  `json:"downpayment" validate:"required,is-downpayment=Price"`
	AnualInterest float64  `json:"anual_interest" validate:"required,gte=0,lte=100"`
}

type PaymentPerSchedule struct {
	PeriodNumber int64   `json:"period_number"`
	Amount       float64 `json:"amount"`
}

type OutputP struct {
	Schedule         Schedule `json:"schedule"`
	NumberOfPayments int      `json:"number_of_payments"`
	Payment          float64  `json:"payment"`
}
