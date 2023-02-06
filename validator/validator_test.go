package validator_test

import (
	"testing"

	"github.com/itsranveer/sj/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidateQuantifier(t *testing.T) {
	assert.NoError(t, validator.ValidateQuantifier(1, 4, false))

	assert.Error(t, validator.ValidateQuantifier(5, 4, false))
}

func TestValidatePercentileQuantifier(t *testing.T) {
	assert.NoError(t, validator.ValidateQuantifier(60, 4, true))

	assert.Error(t, validator.ValidateQuantifier(120, 4, true))
	assert.Error(t, validator.ValidateQuantifier(-50, 4, true))
}
