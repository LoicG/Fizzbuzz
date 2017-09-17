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
	result := []string{}
	for i := 1; i <= limit; i++ {
		value := ""
		if i%int1 == 0 {
			value += string1
		}
		if i%int2 == 0 {
			value += string2
		}
		if value == "" {
			value = strconv.Itoa(i)
		}
		result = append(result, value)
	}
	return result
}

func FizzBuzz(request *http.Request) (interface{}, error) {
	int1, err := utils.RequireIntParameter(request, "int1")
	if err != nil {
		return nil, err
	}
	int2, err := utils.RequireIntParameter(request, "int2")
	if err != nil {
		return nil, err
	}
	limit, err := utils.RequireIntParameter(request, "limit")
	if err != nil {
		return nil, err
	}
	string1, err := utils.RequireParameter(request, "string1")
	if err != nil {
		return nil, err
	}
	string2, err := utils.RequireParameter(request, "string2")
	if err != nil {
		return nil, err
	}
	if int1 < 1 || int2 < 1 {
		return nil, fmt.Errorf("int1/int2 parameter must be greater than zero")
	}
	return getFizzBuzzList(int1, int2, limit, string1, string2), nil
}
