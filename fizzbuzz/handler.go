package main

import (
	"fmt"
	"net/http"
	"strconv"
	"utils"
)

// Return a list of strings with numbers from 1 to 'limit', where:
//  all multiples of 'int1' are replaced by 'string1',
//  all multiples of 'int2' are replaced by 'string2',
//  all multiples of 'int1' and 'int2' are replaced by 'string1string2'
func getFizzBuzzList(int1, int2, limit int, string1, string2 string) []string {
	result := make([]string, limit)
	for i := 1; i <= limit; i++ {
		switch {
		case i%(int1*int2) == 0:
			result[i-1] = string1 + string2
		case i%int1 == 0:
			result[i-1] = string1
		case i%int2 == 0:
			result[i-1] = string2
		default:
			result[i-1] = strconv.FormatInt(int64(i), 10)
		}
	}
	return result
}

func FizzBuzz(request *http.Request) ([]string, error) {
	params := request.URL.Query()
	int1, err := utils.RequireIntParameter(params, "int1")
	if err != nil {
		return nil, err
	}
	int2, err := utils.RequireIntParameter(params, "int2")
	if err != nil {
		return nil, err
	}
	limit, err := utils.RequireIntParameter(params, "limit")
	if err != nil {
		return nil, err
	}
	string1, err := utils.RequireParameter(params, "string1")
	if err != nil {
		return nil, err
	}
	string2, err := utils.RequireParameter(params, "string2")
	if err != nil {
		return nil, err
	}
	if int1 < 1 || int2 < 1 {
		return nil, fmt.Errorf("int1/int2 parameter must be greater than zero")
	}
	return getFizzBuzzList(int1, int2, limit, string1, string2), nil
}
