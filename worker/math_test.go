package worker_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/itsranveer/sj/models"
	"github.com/itsranveer/sj/worker"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	app := worker.NewMath()
	ctx := context.Background()

	input := &models.MathInput{
		Numbers:    []float64{50, 20, 70, 30, 100},
		Quantifier: 1,
	}

	response := &models.MathResponse{
		Description: fmt.Sprintf("Min of %v with quantifier %v", input.Numbers, input.Quantifier),
		Results:     []float64{20},
	}

	assert.Equal(t, response, app.Min(ctx, input))
}

func TestMax(t *testing.T) {
	app := worker.NewMath()
	ctx := context.Background()

	input := &models.MathInput{
		Numbers:    []float64{50, 20, 70, 30, 100},
		Quantifier: 1,
	}

	response := &models.MathResponse{
		Description: fmt.Sprintf("Max of %v with quantifier %v", input.Numbers, input.Quantifier),
		Results:     []float64{100},
	}

	assert.Equal(t, response, app.Max(ctx, input))
}

func TestAvg(t *testing.T) {
	app := worker.NewMath()
	ctx := context.Background()

	input := &models.MathInput{
		Numbers: []float64{50, 20, 70, 30, 25},
	}

	response := &models.MathResponse{
		Description: fmt.Sprintf("Avg of %v", input.Numbers),
		Results:     []float64{39},
	}

	assert.Equal(t, response, app.Avg(ctx, input))
}

func TestMedian(t *testing.T) {
	app := worker.NewMath()
	ctx := context.Background()

	input := &models.MathInput{
		Numbers: []float64{50, 20, 70, 30, 25},
	}

	response := &models.MathResponse{
		Description: fmt.Sprintf("Median of %v", input.Numbers),
		Results:     []float64{30},
	}

	assert.Equal(t, response, app.Median(ctx, input))
}

func TestPercentile(t *testing.T) {
	app := worker.NewMath()
	ctx := context.Background()

	input := &models.MathInput{
		Numbers:    []float64{50, 20, 70, 30, 110},
		Quantifier: 80,
	}

	response := &models.MathResponse{
		Description: fmt.Sprintf("%vth percentile of %v", input.Quantifier, input.Numbers),
		Results:     []float64{70},
	}

	assert.Equal(t, response, app.Percentile(ctx, input))
}
