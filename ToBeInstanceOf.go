package gocek

import "reflect"

func (h *gocekAssertion) ToBeInstanceOf(value reflect.Kind) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if !reflect.DeepEqual(input, nil) {
		if h.o == "==" && reflect.TypeOf(input).Kind().String() != value.String() {
			h.t.FailNow()
		} else if h.o == "!=" && reflect.TypeOf(input).Kind().String() == value.String() {
			h.t.FailNow()
		}
	} else {
		if h.o == "!=" && !reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
