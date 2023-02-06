package models

type MathInput struct {
	Numbers    []float64 `json:"numbers"`
	Quantifier int       `json:"quantifier"`
}

type MathResponse struct {
	Description string    `json:"description"`
	Results     []float64 `json:"results"`
}

type MathResponseError struct {
	Description string `json:"description"`
	Results     string `json:"results"`
}
