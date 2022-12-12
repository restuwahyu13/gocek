package gocek

import (
	"reflect"
	"regexp"
)

func (h *gocekAssertion) ToMatch(regex string) {
	input := h.v
	pattern := regex

	defer h.t.Cleanup(func() {
		input = nil
		pattern = ``
	})

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(pattern, nil) {
		inputType := reflect.TypeOf(input).Kind().String()
		patternType := reflect.TypeOf(pattern).Kind().String()

		if inputType != reflect.String.String() && patternType != reflect.String.String() {
			h.t.FailNow()
		}

		regexExp, err := regexp.Compile(pattern)
		if err != nil {
			h.t.FailNow()
		}

		ok := regexExp.MatchString(input.(string))

		if h.o == "==" && !ok {
			h.t.FailNow()
		} else if h.o == "!=" && ok {
			h.t.FailNow()
		}
	} else {
		h.t.FailNow()
	}
}
