package order

import (
	"net/http"

	"github.com/kopjenmbeng/kanggo_rest_test/internal/utility/validator"
)

type CreateOrderRequest struct {
	Charts []string `json:"charts"`
}

func (req *CreateOrderRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *CreateOrderRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("charts", req.Charts); err != nil {
		return err
	}

	return nil
}
