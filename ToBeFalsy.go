package gocek

import "reflect"

func (h *gocekAssertion) ToBeFalsy() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if !reflect.DeepEqual(input, nil) {
		if h.o == "==" && !reflect.DeepEqual(input, false) && input != false {
			h.t.FailNow()
		} else if h.o == "!=" && !reflect.DeepEqual(input, false) && input == false {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
