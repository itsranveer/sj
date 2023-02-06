package routes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/itsranveer/sj/models"
	"github.com/itsranveer/sj/validator"
	"github.com/itsranveer/sj/worker"
	"github.com/sirupsen/logrus"
)

type Router struct {
	math *worker.Math
}

func NewRoutes(math *worker.Math) *Router {
	return &Router{
		math: math,
	}
}

func (rt *Router) Min(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Min API Call")

	input, err := rt.getInput(w, r)
	if err != nil {
		return
	}

	err = rt.validateInput(w, r, input, false)
	if err != nil {
		return
	}

	result := rt.math.Min(context.Background(), input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (rt *Router) Max(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Max API Call")

	input, err := rt.getInput(w, r)
	if err != nil {
		return
	}

	err = rt.validateInput(w, r, input, false)
	if err != nil {
		return
	}

	result := rt.math.Max(context.Background(), input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (rt *Router) Avg(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Avg API Call")

	input, err := rt.getInput(w, r)
	if err != nil {
		return
	}

	err = rt.validateInput(w, r, input, false)
	if err != nil {
		return
	}

	result := rt.math.Avg(context.Background(), input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (rt *Router) Median(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Median API Call")

	input, err := rt.getInput(w, r)
	if err != nil {
		return
	}

	err = rt.validateInput(w, r, input, false)
	if err != nil {
		return
	}

	result := rt.math.Median(context.Background(), input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (rt *Router) Percentile(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Percentile API Call")

	input, err := rt.getInput(w, r)
	if err != nil {
		return
	}

	err = rt.validateInput(w, r, input, true)
	if err != nil {
		return
	}

	result := rt.math.Percentile(context.Background(), input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (rt *Router) getInput(w http.ResponseWriter, r *http.Request) (*models.MathInput, error) {
	var input models.MathInput
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading input, %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MathResponseError{
			Results:     "failure",
			Description: "error reading input",
		})

		return nil, err
	}

	err = json.Unmarshal(reqBody, &input)
	if err != nil {
		logrus.Errorf("wrong input data, %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.MathResponseError{
			Results:     "failure",
			Description: "wrong input data",
		})

		return nil, err
	}

	return &input, nil
}

func (rt *Router) validateInput(w http.ResponseWriter, r *http.Request, input *models.MathInput, isPercentile bool) error {
	err := validator.ValidateQuantifier(input.Quantifier, len(input.Numbers), isPercentile)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MathResponseError{
			Results:     "failure",
			Description: err.Error(),
		})

		return err
	}

	return nil
}
