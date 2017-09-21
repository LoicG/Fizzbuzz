package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func RequireParameter(request *http.Request, param string) (string, error) {
	value := request.URL.Query().Get(param)
	if value != "" {
		return value, nil
	}
	return "", fmt.Errorf("missing %s parameter in request", param)
}

func RequireIntParameter(request *http.Request, param string) (int, error) {
	value, err := RequireParameter(request, param)
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s parameter: %v", param, err)
	}
	return result, nil
}
