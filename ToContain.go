package gocek

import "strings"

func (h *assertion) ToContain(value string) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = ""
	})

	if h.o == "==" && !strings.Contains(input.(string), value) {
		h.t.FailNow()
	} else if h.o == "!=" && strings.Contains(input.(string), value) {
		h.t.FailNow()
	}
}
