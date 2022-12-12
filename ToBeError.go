package gocek

import (
	"reflect"
)

func (h *gocekAssertion) ToBeError() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if h.o == "==" && !reflect.DeepEqual(input, nil) {
		inputType := reflect.TypeOf(input).Kind().String()

		if !reflect.DeepEqual(inputType, reflect.Ptr.String()) {
			h.t.FailNow()
		}
	} else {
		if h.o == "!=" && input != nil {
			h.t.FailNow()
		}
	}
}
