package gocek

import "reflect"

func (h *gocekAssertion) ToBeTruthy() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if !reflect.DeepEqual(input, nil) {
		inputType := reflect.DeepEqual(input, true)

		if h.o == "==" && !inputType && input != true {
			h.t.FailNow()
		} else if h.o == "!=" && inputType && input == true {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
