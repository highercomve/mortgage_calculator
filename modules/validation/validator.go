package validation

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/highercomve/mortgage_calculator/modules/mortgage/mortgagemodels"
)

var (
	trans ut.Translator
)

// CustomValidator payload validation
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate payload
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func TranslateError(err error) error {
	e := err.(validator.ValidationErrors)
	msgs := []string{}
	for _, value := range e.Translate(trans) {
		msgs = append(msgs, value)
	}

	return errors.New(strings.Join(msgs, ", "))
}

func GetValidator() *CustomValidator {
	validate := &CustomValidator{
		Validator: validator.New(),
	}
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate.Validator, trans)

	validate.Validator.RegisterValidation("is-schedule", isValidSchedule)
	validate.Validator.RegisterValidation("is-period", isValidPeriod)
	validate.Validator.RegisterValidation("is-downpayment", isValidDownpayment)

	validate.Validator.RegisterTranslation("is-schedule", trans, func(ut ut.Translator) error {
		return ut.Add(
			"is-schedule",
			fmt.Sprintf(
				"Schedule value is not supported. The supported schedules are: %s",
				mortgagemodels.PosibleSchedules(),
			),
			true,
		)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is-schedule", fe.Field())
		return t
	})

	validate.Validator.RegisterTranslation("is-period", trans, func(ut ut.Translator) error {
		return ut.Add(
			"is-period",
			"period should be 5 year increments between 5 and 30 years",
			true,
		)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is-period", fe.Field())
		return t
	})

	validate.Validator.RegisterTranslation("is-downpayment", trans, func(ut ut.Translator) error {
		return ut.Add(
			"is-downpayment",
			"downpayment need to be more than 5% and less than 100%",
			true,
		)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("is-downpayment", fe.Field())
		return t
	})

	return validate
}

func isValidDownpayment(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()

	compareField, compareKind, _, ok := fl.GetStructFieldOK2()
	if !ok || compareKind != kind {
		return false
	}

	percentage := field.Float() * 100 / compareField.Float()
	percentage = math.Round(percentage*100) / 100

	return percentage >= 5 && percentage < 100
}

func isValidPeriod(field validator.FieldLevel) bool {
	if valuer, ok := field.Field().Interface().(mortgagemodels.Period); ok {
		err := valuer.IsValid()
		if err == nil {
			return true
		}
	}

	return false
}

func isValidSchedule(field validator.FieldLevel) bool {
	if valuer, ok := field.Field().Interface().(mortgagemodels.Schedule); ok {
		err := valuer.IsValid()
		if err == nil {
			return true
		}
	}

	return false
}
