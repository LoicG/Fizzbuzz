package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func WriteJson(w http.ResponseWriter, output interface{}) {
	bytes, err := json.Marshal(output)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func RequireParameter(values url.Values, param string) (string, error) {
	value := values.Get(param)
	if value != "" {
		return value, nil
	}
	return "", fmt.Errorf("missing %s parameter in request", param)
}

func RequireIntParameter(values url.Values, param string) (int, error) {
	value, err := RequireParameter(values, param)
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s parameter: %v", param, err)
	}
	return result, nil
}
