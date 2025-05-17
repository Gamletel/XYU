package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return err
	}
	return validate.Struct(dst)
}
