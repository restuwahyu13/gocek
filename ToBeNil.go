package gocek

import "reflect"

func (h *gocekAssertion) ToBeNil() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	inputType := reflect.DeepEqual(input, nil)

	if h.o == "==" && !inputType {
		h.t.FailNow()
	} else if h.o == "!=" && inputType {
		h.t.FailNow()
	}
}
