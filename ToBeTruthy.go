package gocek

import "reflect"

func (h *assertion) ToBeTruthy() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if !reflect.DeepEqual(input, nil) {
		if h.o == "==" && input != true {
			h.t.FailNow()
		} else if h.o == "!=" && input == true {
			h.t.FailNow()
		}
	} else {
		h.t.FailNow()
	}
}
