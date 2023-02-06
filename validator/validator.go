package validator

import (
	"errors"
	"math"

	"github.com/sirupsen/logrus"
)

const _INCORRECT_QUANTIFIER = "incorrect Quantifier: Quantifier is greater than length of numbers array"
const _INCORRECT_PERCENTILE_QUANTIFIER = "incorrect Quantifier: Quantifier must be between 0 and 100"

func ValidateQuantifier(quantifier, numbersLength int, isPercentile bool) error {
	if isPercentile {
		return validatePercentileQuantifier(quantifier, numbersLength)
	}

	return validateQuantifier(quantifier, numbersLength)
}

func validateQuantifier(quantifier, numbersLength int) error {
	if quantifier > numbersLength {
		logrus.Errorf(_INCORRECT_QUANTIFIER)

		return errors.New(_INCORRECT_QUANTIFIER)
	}

	return nil
}

func validatePercentileQuantifier(quantifier, numbersLength int) error {
	index := int(math.Round((float64(quantifier) / 100) * float64(numbersLength)))
	if index < 0 || index > numbersLength {
		logrus.Errorf(_INCORRECT_PERCENTILE_QUANTIFIER)

		return errors.New(_INCORRECT_PERCENTILE_QUANTIFIER)
	}

	return nil
}
