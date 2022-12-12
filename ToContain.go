package gocek

import (
	"reflect"
	"strings"
)

func (h *gocekAssertion) ToContain(value string) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = ""
	})

	if !reflect.DeepEqual(input, nil) && reflect.TypeOf(input).String() == reflect.TypeOf(value).String() {
		if h.o == "==" && !strings.Contains(input.(string), value) {
			h.t.FailNow()
		} else if h.o == "!=" && strings.Contains(input.(string), value) {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && !reflect.DeepEqual(input, value) {
			h.t.FailNow()
		}
	}
}
