package server

import (
	"fmt"
	"net/url"
	"subLease/src/server/address"
	"subLease/src/server/socialSecurityNumber"
	"time"
)

func retrieveInt(key string, queryValues url.Values, converterFunction func(s string) (int, error), errorMessages *[]string, handlerFunction func(int)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveAddress(key string, queryValues url.Values, converterFunction func(s string) (address.Address, error), errorMessages *[]string, handlerFunction func(address.Address)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveTime(key string, queryValues url.Values, converterFunction func(s string) (time.Time, error), errorMessages *[]string, handlerFunction func(time.Time)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveString(key string, queryValues url.Values, converterFunction func(s string) (string, error), errorMessages *[]string, handlerFunction func(string)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveSocialSecurityNumber(key string, queryValues url.Values, converterFunction func(s string) (socialSecurityNumber.SocialSecurityNumber, error), errorMessages *[]string, handlerFunction func(socialSecurityNumber.SocialSecurityNumber)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveIntSlice(key string, queryValues url.Values, converterFunction func(s string) ([]int, error), errorMessages *[]string, handlerFunction func([]int)) {
	if stringValue, present := retrieveIfPresent(key, queryValues); present {
		value, err := converterFunction(stringValue)
		if err != nil {
			*errorMessages = append(*errorMessages, fmt.Sprintf("Error while parsing query parameter %s. Error: %s", key, err.Error()))
		}
		handlerFunction(value)
	}
}

func retrieveIfPresent(key string, queryValues url.Values) (string, bool) {
	value := queryValues.Get(key)
	return value, value != ""
}
