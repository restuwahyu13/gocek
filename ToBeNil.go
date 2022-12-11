package gocek

import "reflect"

func (h *gocekAssertion) ToBeNil() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if h.o == "==" && !reflect.DeepEqual(input, nil) {
		h.t.FailNow()
	} else if h.o == "!=" && reflect.DeepEqual(input, nil) {
		h.t.FailNow()
	}
}
