package server


func retrieveInt(key string, queryValues url.Values, converterFunction func(s string) (int, error), handlerFunction func(int)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}

func retrieveAddress(key string, queryValues url.Values, converterFunction func(s string) (address.Address, error), handlerFunction func(address.Address)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}

func retrieveTime(key string, queryValues url.Values, converterFunction func(s string) (time.Time, error), handlerFunction func(time.Time)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}

func retrieveString(key string, queryValues url.Values, converterFunction func(s string) (string, error), handlerFunction func(string)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}

func retrieveSocialSecurityNumber(key string, queryValues url.Values, converterFunction func(s string) (socialSecurityNumber.SocialSecurityNumber, error), handlerFunction func(socialSecurityNumber.SocialSecurityNumber)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}

func retrieveIntSlice(key string, queryValues url.Values, converterFunction func(s string) ([]int, error), handlerFunction func([]int)) {
    if stringValue, present := retrieveIfPresent(key, queryValues); present {
        value, err := converterFunction(stringValue)
        if err != nil {
            panic(err)
        }
        handlerFunction(value)
    }
}


func retrieveIfPresent(key string, queryValues url.Values) (string, bool) {
	value := queryValues.Get(key)
	return value, value != ""
}
