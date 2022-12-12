package gocek

import (
	"reflect"
	"strings"
)

func (h *gocekAssertion) ToBeError() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if h.o == "==" && !reflect.DeepEqual(input, nil) {
		inputType := reflect.TypeOf(input).Kind().String()
		isError := strings.Contains(reflect.ValueOf(input).Elem().Addr().String(), "errors")

		if !reflect.DeepEqual(inputType, reflect.Ptr.String()) || !isError {
			h.t.FailNow()
		}
	} else {
		if h.o == "!=" && input != nil {
			h.t.FailNow()
		}
	}
}
