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

	if !reflect.DeepEqual(input, nil) {
		isError := strings.Contains(reflect.ValueOf(input).String(), "*errors")

		if h.o == "==" && !isError {
			h.t.FailNow()
		} else if h.o == "!=" && isError {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
