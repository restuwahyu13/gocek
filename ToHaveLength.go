package gocek

func (h *gocekAssertion) ToHaveLength(value int) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = 0
	})

	if h.o == "==" && input != value {
		h.t.FailNow()
	} else if h.o == "!=" && input == value {
		h.t.FailNow()
	}
}
