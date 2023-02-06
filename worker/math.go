package worker

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/itsranveer/sj/models"
)

type Math struct {
}

func NewMath() *Math {
	return &Math{}
}

func (m *Math) Min(ctx context.Context, input *models.MathInput) *models.MathResponse {
	desc := fmt.Sprintf("Min of %v with quantifier %v", input.Numbers, input.Quantifier)

	sort.Float64s(input.Numbers)

	return &models.MathResponse{
		Description: desc,
		Results:     input.Numbers[:input.Quantifier],
	}
}

func (m *Math) Max(ctx context.Context, input *models.MathInput) *models.MathResponse {
	desc := fmt.Sprintf("Max of %v with quantifier %v", input.Numbers, input.Quantifier)

	sort.Float64s(input.Numbers)

	return &models.MathResponse{
		Description: desc,
		Results:     input.Numbers[len(input.Numbers)-input.Quantifier:],
	}
}

func (m *Math) Avg(ctx context.Context, input *models.MathInput) *models.MathResponse {
	desc := fmt.Sprintf("Avg of %v", input.Numbers)

	var sum float64 = 0
	for _, v := range input.Numbers {
		sum += v
	}

	return &models.MathResponse{
		Description: desc,
		Results:     []float64{sum / float64(len(input.Numbers))},
	}
}

func (m *Math) Median(ctx context.Context, input *models.MathInput) *models.MathResponse {
	desc := fmt.Sprintf("Median of %v", input.Numbers)

	sort.Float64s(input.Numbers)

	if len(input.Numbers)%2 == 0 {
		return &models.MathResponse{
			Description: desc,
			Results:     []float64{(input.Numbers[len(input.Numbers)/2-1] + input.Numbers[len(input.Numbers)/2]) / 2},
		}
	}

	return &models.MathResponse{
		Description: desc,
		Results:     []float64{input.Numbers[len(input.Numbers)/2]},
	}
}

func (m *Math) Percentile(ctx context.Context, input *models.MathInput) *models.MathResponse {
	desc := fmt.Sprintf("%vth percentile of %v", input.Quantifier, input.Numbers)

	sort.Float64s(input.Numbers)
	index := int(math.Round((float64(input.Quantifier) / 100) * float64(len(input.Numbers))))

	if index == 0 {
		return &models.MathResponse{
			Description: desc,
			Results:     []float64{input.Numbers[index]},
		}
	}

	return &models.MathResponse{
		Description: desc,
		Results:     []float64{input.Numbers[index-1]},
	}
}
