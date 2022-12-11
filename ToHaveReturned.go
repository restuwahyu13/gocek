package gocek

import "reflect"

func (h *gocekAssertion) ToHaveReturned() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if h.o == "==" && reflect.TypeOf(input).Kind().String() == reflect.Func.String() && reflect.TypeOf(input).NumOut() != 1 {
		h.t.FailNow()
	} else if h.o == "!=" && reflect.TypeOf(input).Kind().String() == reflect.Func.String() && reflect.TypeOf(input).NumOut() == 1 {
		h.t.FailNow()
	}
}
