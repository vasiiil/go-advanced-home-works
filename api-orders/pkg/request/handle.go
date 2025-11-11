package request

import (
	"api-orders/pkg/response"
	"errors"
	"net/http"
	"strconv"
)

var ErrUnknownParamType = errors.New("unknown param type")
var ErrUnknownParamSource = errors.New("unknown param source")
var ErrEmptyRequired = errors.New("empty required param")

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		response.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}

func PrepareParam[T any](w *http.ResponseWriter, r *http.Request, paramSource string, paramName string, required bool) (T, error) {
	var stringValue string
	var zeroValue T
	switch paramSource {
	case "query":
		queryValues := r.URL.Query()
		stringValue = queryValues.Get(paramName)
	case "path":
		stringValue = r.PathValue(paramName)
	default:
		response.JsonError(*w, ErrUnknownParamSource.Error(), http.StatusBadRequest)
		return zeroValue, ErrUnknownParamSource
	}
	if stringValue == "" {
		if required {
			response.JsonError(*w, ErrEmptyRequired.Error(), http.StatusBadRequest)
			return zeroValue, ErrEmptyRequired
		}
		return zeroValue, nil
	}

	var untypedValue any
	var err error
	switch any(zeroValue).(type) {
	case string:
		return any(stringValue).(T), nil
	case uint:
		value, err := strconv.ParseUint(stringValue, 10, 64)
		if err != nil {
			response.JsonError(*w, err.Error(), http.StatusBadRequest)
			return zeroValue, err
		}
		untypedValue = uint(value)
	case int:
		value, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			response.JsonError(*w, err.Error(), http.StatusBadRequest)
			return zeroValue, err
		}
		untypedValue = int(value)
	case float64:
		untypedValue, err = strconv.ParseFloat(stringValue, 64)
	case bool:
		untypedValue, err = strconv.ParseBool(stringValue)
	default:
		response.JsonError(*w, ErrUnknownParamType.Error(), http.StatusBadRequest)
		return zeroValue, ErrUnknownParamType
	}
	if err != nil {
		response.JsonError(*w, err.Error(), http.StatusBadRequest)
		return zeroValue, err
	}

	return any(untypedValue).(T), nil
}
