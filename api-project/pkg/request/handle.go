package request

import (
	"api-project/pkg/response"
	"errors"
	"fmt"
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
		stringValue = r.URL.Query().Get(paramName)
	case "path":
		stringValue = r.PathValue(paramName)
	default:
		response.JsonError(*w, fmt.Sprintf(ErrUnknownParamSource.Error()+": %s", paramSource), http.StatusBadRequest)
		return zeroValue, ErrUnknownParamSource
	}
	if stringValue == "" {
		if required {
			response.JsonError(*w, fmt.Sprintf(ErrEmptyRequired.Error()+": %s", paramName), http.StatusBadRequest)
			return zeroValue, ErrEmptyRequired
		}
		return zeroValue, nil
	}

	var untypedValue any
	var err error
	var convertTo string
	switch any(zeroValue).(type) {
	case string:
		return any(stringValue).(T), nil
	case uint:
		value, err := strconv.ParseUint(stringValue, 10, 32)
		if err != nil {
			response.JsonError(*w, fmt.Sprintf("Error convert %s to uint: %s", paramName, err.Error()), http.StatusBadRequest)
			return zeroValue, err
		}
		untypedValue = uint(value)
	case int:
		value, err := strconv.ParseInt(stringValue, 10, 32)
		if err != nil {
			response.JsonError(*w, fmt.Sprintf("Error convert %s to int: %s", paramName, err.Error()), http.StatusBadRequest)
			return zeroValue, err
		}
		untypedValue = int(value)
	case float64:
		convertTo = "float64"
		untypedValue, err = strconv.ParseFloat(stringValue, 32)
	case bool:
		convertTo = "bool"
		untypedValue, err = strconv.ParseBool(stringValue)
	default:
		response.JsonError(*w, fmt.Sprintf(ErrUnknownParamType.Error()+": %s", paramName), http.StatusBadRequest)
		return zeroValue, ErrUnknownParamType
	}
	if err != nil {
		response.JsonError(*w, fmt.Sprintf("Error convert %s to %s: %s", paramName, convertTo, err.Error()), http.StatusBadRequest)
		return zeroValue, err
	}

	return any(untypedValue).(T), nil
}
